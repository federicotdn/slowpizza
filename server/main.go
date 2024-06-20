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
	"google.golang.org/grpc/metadata"
)

const (
	defaultPort = 50051
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
	log.Printf("received order item: %v", in.Item)
	s.logHeaders(ctx)
	return s.replyForOrderRequest(in), nil
}

func (s *server) OrderMultipleItems(client pb.Delivery_OrderMultipleItemsServer) error {
	log.Printf("called: OrderMultipleItems")
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
	s := grpc.NewServer()
	pb.RegisterDeliveryServer(s, &server{})
	log.Printf("SlowPizza server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
