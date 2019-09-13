package connection

import (
	"encoding/base64"
	"log"

	"github.com/qiushuiqs/idhash-benchmark/cache"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func Conn(key string) (*cache.Response, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("40.91.88.191:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
		return nil, err
	}
	defer conn.Close()

	// s := "CgObKRQxNTQ3MzQzMTc4MDU1NzM4MjU5MAAA"
	sDec, _ := base64.StdEncoding.DecodeString(key)
	req := &cache.Request{
		MUID: string(sDec),
	}

	c := cache.NewCacheServiceClient(conn)
	response, err := c.ProcessRequest(context.Background(), req)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
		return nil, err
	}
	return response, nil
	// if response != nil {
	// 	log.Printf("Response from server: %s", response.String())
	// 	log.Println(response.LinkedinID)

	// } else {
	// 	log.Println("nil response")
	// }
}
