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
	conn, err := grpc.Dial("52.246.252.11:443", grpc.WithInsecure())
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
	// key := "CgObKRM3MjQzMDc0OTQwMDgwMDE0MjY3AAA="
	sDec, _ := base64.StdEncoding.DecodeString(key)
	req := &cache.Request{
		Key: sDec,
	}
	response, err := client.cacheServiceClient.ProcessRequest(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
