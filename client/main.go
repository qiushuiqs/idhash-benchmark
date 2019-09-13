package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/qiushuiqs/idhash-benchmark/cache"
	"github.com/qiushuiqs/idhash-benchmark/connection"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("40.91.88.191:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	s := "CgObKRQxNTQ3MzQzMTc4MDU1NzM4MjU5MAAA"
	numOfError := 0

	for x := 0; x < 100; x++ {
		go func() {
			// defer timeTrack(time.Now(), "factorial")
			response, err := connection.Conn(s)
			if err != nil {
				log.Fatalf("Error : %s", err)
				numOfError++
			}
			fmt.Println(response.String())
		}()
	}

	// s := "CgObKRQxNTQ3MzQzMTc4MDU1NzM4MjU5MAAA"
	sDec, _ := base64.StdEncoding.DecodeString("CgObKRQxNTQ3MzQzMTc4MDU1NzM4MjU5MAAA")
	req := &cache.Request{
		MUID: string(sDec),
	}
	c := cache.NewCacheServiceClient(conn)
	response, err := c.ProcessRequest(context.Background(), req)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Println(response.String())
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s", name, elapsed)
}
