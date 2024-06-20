package main

import (
	"context"
	"crypto/x509"
	"flag"
	"log"
	"strings"
	"time"

	pb "github.com/federicotdn/slowpizza/slowpizza"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return strings.Join(*i, ", ")
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

type tokenGRPCAuth string

var _ credentials.PerRPCCredentials = tokenGRPCAuth("")

func (t tokenGRPCAuth) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	return map[string]string{
		"Authorization": "Bearer " + string(t),
	}, nil
}

func (tokenGRPCAuth) RequireTransportSecurity() bool {
	return true
}

var (
	addr      = flag.String("addr", "localhost:50051", "The address to connect to")
	plaintext = flag.Bool("plaintext", false, "Disable TLS")
	authToken = flag.String("token", "", "Authorization token")
	headers arrayFlags
	items arrayFlags
)

func main() {
	flag.Var(&headers, "H", "Extra headers ('Key: Value')")
	flag.Var(&items, "item", "Item to order")
	flag.Parse()

	if len(items) == 0 {
		log.Fatalf("no items to order were specified (use -item)")
	}

	token := grpc.WithPerRPCCredentials(tokenGRPCAuth(*authToken))
	pool := x509.NewCertPool()
	tc := credentials.NewClientTLSFromCert(pool, "")

	var opts []grpc.DialOption

	if *plaintext {
		opts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	} else {
		opts = []grpc.DialOption{grpc.WithTransportCredentials(tc)}
	}

	if *authToken != "" {
		opts = append(opts, token)
	}

	conn, err := grpc.NewClient(*addr, opts...)
	if err != nil {
		log.Fatalf("grpc connection error: %v", err)
	}
	defer conn.Close()
	c := pb.NewDeliveryClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for _, elem := range headers {
		key, value, found := strings.Cut(elem, ":")
		key = strings.TrimSpace(key)
		if !found || key == "" {
			log.Fatalf("invalid header specified")
		}
		ctx = metadata.AppendToOutgoingContext(ctx, key, strings.TrimSpace(value))
	}

	if len(items) == 1 {
		log.Printf("adding item to order: %v", items[0])
		resp, err := c.OrderItem(ctx, &pb.OrderRequest{Item: items[0]})
		if err != nil {
			log.Fatalf("grpc error: %v", err)
		}

		log.Printf("response: %v", resp.Message)
	} else {
		order, err := c.OrderMultipleItems(ctx)
		if err != nil {
			log.Fatalf("grpc error: %v", err)
		}

		for _, item := range items {
			log.Printf("adding item to order: %v", item)
			err = order.Send(&pb.OrderRequest{Item: item})
			if err != nil {
				log.Fatalf("grpc error: %v", err)
			}

			resp, err := order.Recv()
			if err != nil {
				log.Fatalf("grpc error: %v", err)
			}

			log.Printf("response: %v", resp.Message)
		}
	}
}
