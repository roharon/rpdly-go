package main

import (
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
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

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
