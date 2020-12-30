package coverage

import (
	"context"
	"log"
	"net"
	"os"
	"strconv"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/roharon/rpdly-go-url/handler"
	pb "github.com/roharon/rpdly-go-url/protobuf/uri/v1"
	redis "github.com/roharon/rpdly-go-url/utils/redis"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	pb.RegisterUriExchangeServer(server, &handler.RouteServer{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestGetSetUri(t *testing.T) {
	mr, err := miniredis.Run()

	os.Setenv("REDIS_ADDRESS", mr.Addr())
	os.Setenv("REDIS_PASSWORD", "")
	// set env for using redis on test

	if err != nil {
		// set redis server
		t.Error(err, mr)
	}

	test := []struct {
		uri string
	}{
		{
			"https://naver.com",
		},
		{
			"https://google.com",
		},
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer((dialer())))

	if err != nil {
		t.Error(err)
	}
	defer conn.Close()

	client := pb.NewUriExchangeClient(conn)

	for _, tt := range test {
		t.Run(tt.uri, func(t *testing.T) {
			request := &pb.Request{Uri: tt.uri}
			response, err := client.SetUri(ctx, request)

			rds := redis.RedisClient()
			// for check uri DEFAULT_LENGTH

			if response != nil {
				length, err := rds.Get(handler.RANDOM_LENGTH)

				if err != nil {
					t.Error(err)
				}

				val, err := strconv.Atoi(length)
				if err != nil {
					t.Error(err)
				}

				if len(response.GetUri()) != val {
					t.Errorf("Url LENGTH error ||Uri: %s expected Len:%d||", response.GetUri(), val)
				}

				if key, err := rds.Get(handler.URL_PREFIX + response.GetUri()); key != tt.uri || err != nil {
					t.Error("redis Key is wrong")
				}
			}

			if err != nil {
				t.Error("error occur", err)
				// if add testcase that occur error, should change it.
			}
		})
	}
}
