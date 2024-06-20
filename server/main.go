package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	pb "github.com/federicotdn/slowpizza/slowpizza"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	defaultPort = 50051
)

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)

type server struct {
	pb.UnimplementedDeliveryServer
}

func (s *server) logHeaders(ctx context.Context) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Printf("no context metadata")
	} else {
		log.Printf("metadata debug:")
		for k, vs := range md {
			log.Printf("key: %v", k)
			for _, v := range vs {
				log.Printf("value: %v", v)
			}
		}
	}
}

func (s *server) replyForOrderRequest(req *pb.OrderRequest) *pb.OrderReply {
	item := strings.ToLower(req.Item)
	message := "Sorry, we do not serve that item."
	if strings.Index(item, "pizza") != -1 {
		message = fmt.Sprintf("Added %v to order.", req.Item)
	}

	return &pb.OrderReply{Message: message}
}

func (s *server) OrderItem(ctx context.Context, in *pb.OrderRequest) (*pb.OrderReply, error) {
	log.Printf("called: OrderItem")
	s.logHeaders(ctx)

	log.Printf("received order item: %v", in.Item)
	return s.replyForOrderRequest(in), nil
}

func (s *server) OrderMultipleItems(client pb.Delivery_OrderMultipleItemsServer) error {
	log.Printf("called: OrderMultipleItems")
	s.logHeaders(client.Context())

	for {
		req, err := client.Recv()
		if err != nil {
			return err
		}

		log.Printf("received order item: %v", req.Item)
		err = client.Send(s.replyForOrderRequest(req))
		if err != nil {
			return err
		}
	}
}

func validateContextToken(ctx context.Context) error {
	token := os.Getenv("SLOWPIZZA_AUTH_TOKEN")
	if token == "disabled" {
		log.Printf("auth token disabled")
		return nil
	} else if token == "" {
		log.Printf("auth token incorrectly configured")
		return errInvalidToken
	}

	log.Printf("performing auth token authentication")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errMissingMetadata
	}
	value := md["authorization"]
	if len(value) == 0 {
		return errInvalidToken
	}
	auth := value[0]
	if auth != "Bearer "+token {
		return errInvalidToken
	}

	return nil
}

func ensureValidTokenStream(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if err := validateContextToken(ss.Context()); err != nil {
		return err
	}
	return handler(srv, ss)
}

func ensureValidToken(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	if err := validateContextToken(ctx); err != nil {
		return nil, err
	}
	return handler(ctx, req)
}

func main() {
	portStr := os.Getenv("SLOWPIZZA_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = defaultPort
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(ensureValidToken),
		grpc.StreamInterceptor(ensureValidTokenStream),
	}

	s := grpc.NewServer(opts...)
	pb.RegisterDeliveryServer(s, &server{})
	log.Printf("SlowPizza server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
