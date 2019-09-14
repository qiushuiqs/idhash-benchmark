package connection

import (
	"context"
	"encoding/base64"
	"log"

	"github.com/qiushuiqs/idhash-benchmark/cache"
	"google.golang.org/grpc"
)

type IdHashClient struct {
	cacheServiceClient cache.CacheServiceClient
}

func Conn() (*IdHashClient, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("40.91.88.191:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
		return nil, err
	}
	// defer conn.Close()

	// s := "CgObKRQxNTQ3MzQzMTc4MDU1NzM4MjU5MAAA"
	// sDec, _ := base64.StdEncoding.DecodeString(key)
	// req := &cache.Request{
	// 	MUID: string(sDec),
	// }

	c := cache.NewCacheServiceClient(conn)
	return &IdHashClient{cacheServiceClient: c}, nil

}

func (client *IdHashClient) Exec(key string) (*cache.Response, error) {
	// s := "CgObKRQxNTQ3MzQzMTc4MDU1NzM4MjU5MAAA"
	sDec, _ := base64.StdEncoding.DecodeString(key)
	req := &cache.Request{
		MUID: string(sDec),
	}
	response, err := client.cacheServiceClient.ProcessRequest(context.Background(), req)
	if err != nil {
		log.Fatalf("Error when calling Exec: %s", err)
		return nil, err
	}
	return response, nil
}
