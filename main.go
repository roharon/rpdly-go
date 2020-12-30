package main

import (
	"log"
	"net"

	"github.com/roharon/rpdly-go-url/config"
	handler "github.com/roharon/rpdly-go-url/handler"
	pb "github.com/roharon/rpdly-go-url/protobuf/uri/v1"
	"google.golang.org/grpc"
)

func main() {
	configuration := config.GetConfig()

	lis, err := net.Listen("tcp", configuration.SERVER_ADDRESS)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterUriExchangeServer(s, &handler.RouteServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
