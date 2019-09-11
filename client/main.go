package main

import (
	"log"

	"github.com/benchmark/cache"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	req := &cache.Request{
		MUID: "05CADF6776056BC81AD2D2EC72056851",
	}

	c := cache.NewCacheServiceClient(conn)
	response, err := c.ProcessRequest(context.Background(), req)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Greeting)
}
