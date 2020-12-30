package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/roharon/rpdly-go-url/config"
	pb "github.com/roharon/rpdly-go-url/protobuf/uri/v1"
	"google.golang.org/grpc"
)

func Run() error {
	conf := config.GetConfig()

	grpcServerEnpoint := flag.String("grpc_endpoint", conf.SERVER_ADDRESS, "gRPC server endpoint")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterUriExchangeHandlerFromEndpoint(ctx, mux, *grpcServerEnpoint, opts)

	if err != nil {
		return err
	}

	return http.ListenAndServe(conf.PROXY_PORT, mux)
}

func main() {
	flag.Parse()

	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
