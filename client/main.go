package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

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
	return !*unsafeToken
}

const itemPrompt = "Enter item name: "

var (
	addr            = flag.String("addr", "localhost:50051", "The address to connect to")
	plaintext       = flag.Bool("plaintext", false, "Disable TLS")
	authToken       = flag.String("token", "", "Authorization token")
	cert            = flag.String("cert", "", "Path to PEM-encoded server certificate (.crt)")
	unsafeToken     = flag.Bool("unsafetoken", false, "Allow sending authorization token over plaintext")
	interactive     = flag.Bool("i", false, "Use interactive prompt to send items")
	confirmCount    = flag.Int("confirmcount", 1, "Number of times the server must confirm the order")
	confirmInterval = flag.Int("confirminterval", 0, "Number of seconds between order confirmations")
	headers         arrayFlags
	items           arrayFlags
)

func promptInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func main() {
	flag.Var(&headers, "H", "Extra headers ('Key: Value')")
	flag.Var(&items, "item", "Item to order")
	flag.Parse()

	if len(items) == 0 && !*interactive {
		log.Fatalf("no items to order were specified (use -item)")
	}

	var opts []grpc.DialOption

	if *plaintext {
		opts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	} else {
		if *cert != "" {
			creds, err := credentials.NewClientTLSFromFile(*cert, "")
			if err != nil {
				log.Fatalf("could not create client tls: %v", err)
			}

			opts = []grpc.DialOption{grpc.WithTransportCredentials(creds)}
		} else {
			opts = []grpc.DialOption{grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
				MinVersion: tls.VersionTLS13,
			}))}
		}
	}

	if *authToken != "" {
		token := grpc.WithPerRPCCredentials(tokenGRPCAuth(*authToken))
		opts = append(opts, token)
	}

	conn, err := grpc.NewClient(*addr, opts...)
	if err != nil {
		log.Fatalf("grpc connection error: %v", err)
	}
	defer conn.Close()
	c := pb.NewAgentClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, elem := range headers {
		key, value, found := strings.Cut(elem, ":")
		key = strings.TrimSpace(key)
		if !found || key == "" {
			log.Fatalf("invalid header specified")
		}
		ctx = metadata.AppendToOutgoingContext(ctx, key, strings.TrimSpace(value))
	}

	if *interactive {
		item := promptInput(itemPrompt)
		items = append(items, item)
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

		for i := 0; i < len(items); i++ {
			item := items[i]

			log.Printf("adding item to order: %v", item)
			err = order.Send(&pb.OrderRequest{
				Item:             item,
				ConfirmCount:     int32(*confirmCount),
				ConfirmIntervalS: int64(*confirmInterval),
			})
			if err != nil {
				log.Fatalf("grpc error: %v", err)
			}

			for i := 0; i < *confirmCount; i++ {
				resp, err := order.Recv()
				if err != nil {
					log.Fatalf("grpc error: %v", err)
				}

				log.Printf("response %v: %v", i+1, resp.Message)
			}

			if *interactive {
				newItem := promptInput(itemPrompt)
				if newItem == "" {
					break
				}

				items = append(items, newItem)
			}
		}
	}
}
