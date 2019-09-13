package main

import (
	"fmt"
	"log"
	"time"

	"github.com/qiushuiqs/idhash-benchmark/connection"
)

func main() {

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

}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s", name, elapsed)
}
