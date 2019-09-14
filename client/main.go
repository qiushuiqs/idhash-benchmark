package main

import (
	_ "bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/qiushuiqs/idhash-benchmark/connection"
	"github.com/rcrowley/go-metrics"

)

func main() {
	// workers, err := strconv.Atoi(os.Args[1])
	// if err != nil {
	// 	workers = 10
	// }

	workers := 10
	if len(os.Args) > 1 {
		workers, _ = strconv.Atoi(os.Args[1])
		// do something with command

	}
	runtime.GOMAXPROCS(workers)
	// read file
	tsvFile, err := os.Open("haha.txt")
	if err != nil {
		log.Println("error")
	}
	defer tsvFile.Close()
	reader := csv.NewReader(tsvFile)

	reader.Comma = '\t' // Use tab-delimited instead of comma <---- here!

	reader.FieldsPerRecord = -1
	csvData, _ := reader.ReadAll()

	start := time.Now()
	jobs := make(chan int, len(csvData))
	results := make(chan string, len(csvData))

	registry := metrics.NewRegistry()
	latencyTimer := metrics.NewRegisteredTimer("latency-timer", registry)
	reqCounter := metrics.NewRegisteredCounter("request-counter", registry)
	go runPrintMetics(&latencyTimer, &reqCounter)
	client, err := connection.Conn()
	// s := "CgObKRQxNTQ3MzQzMTc4MDU1NzM4MjU5MAAA"
	// numOfError := 0

	for x := 0; x < workers; x++ {
		go worker(x, jobs, results, csvData, client, &latencyTimer, &reqCounter)
	}

	for idx := range csvData {
		jobs <- idx

	}
	close(jobs)
	for a := 1; a <= len(csvData); a++ {
		// log.Println(<-results)
		<-results
	}
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func worker(id int, jobs <-chan int, results chan<- string, csvData [][]string, client *connection.IdHashClient, latencyTimer *metrics.Timer, reqCounter *metrics.Counter) {
	for j := range jobs {
		start := time.Now()
		response, err := client.Exec(csvData[j][0])
		if err != nil {
			log.Fatalf("Error : %s", err)
		}
		results <- response.String()
		_ = time.Since(start)
		(*latencyTimer).UpdateSince(start)
		(*reqCounter).Inc(1)
		// log.Printf("%s take %s", strconv.Itoa(j), elapsed)
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s", name, elapsed)
}


func runPrintMetics(latencyTimer *metrics.Timer, reqCounter *metrics.Counter) {
	ticker := time.Tick(10 * time.Second)
	runSecs := 0
	for now := range ticker {
		runSecs += 10
		latencyMetric := (*latencyTimer).Snapshot()
		counterMetric := (*reqCounter).Snapshot()
		// errMetrics := (*errCounter).Snapshot()
		// noMuidMetric := (*noMuidCounter).Snapshot()

		fmt.Printf("%s latency  %f,%f,%f, mean %f, num of requests %d", now, latencyMetric.Percentile(0.50)/1000000, latencyMetric.Percentile(0.95)/1000000, latencyMetric.Percentile(0.99)/1000000, latencyMetric.Mean()/1000000, counterMetric.Count() )
		fmt.Println()
	}
}
