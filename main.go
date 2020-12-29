package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	handler "github.com/roharon/rpdly-go-url/handler"
	pb "github.com/roharon/rpdly-go-url/protobuf/uri"

	"google.golang.org/grpc"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env %v", err)
	}

	lis, err := net.Listen("tcp", os.Getenv("address"))
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUriExchangeServer(s, &handler.RouteServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
