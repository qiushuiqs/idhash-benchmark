package main

import (
	_ "bufio"
	_ "bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
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

	// runtime.GOMAXPROCS(workers)
	// read file
	tsvFile, err := os.Open("haha.txt")
	if err != nil {
		log.Println("error")
	}
	defer tsvFile.Close()
	reader := csv.NewReader(tsvFile)
	reader.Comma = '\t' // Use tab-delimited instead of comma <---- here!
	reader.FieldsPerRecord = -1

	// br := bufio.NewReader(tsvFile)
	// bom, _, err := br.ReadRune()
	// if bom != '\uFEFF' {
	// 	br.UnreadRune() // Not a BOM -- put the rune back
	// }
	// line, err := br.ReadSlice('\n')
	// tokens := bytes.Split(line, []byte("\t"))

	csvData, _ := reader.ReadAll()

	lookups := len(csvData)
	workers := 8

	if len(os.Args) > 2 {
		workers, _ = strconv.Atoi(os.Args[1])
		lookups, _ = strconv.Atoi(os.Args[2])
	}

	jobs := make(chan int, lookups)
	results := make(chan string, lookups)

	registry := metrics.NewRegistry()
	latencyTimer := metrics.NewRegisteredTimer("latency-timer", registry)
	reqCounter := metrics.NewRegisteredCounter("request-counter", registry)
	clients := make([]*connection.IdHashClient, workers)
	for i := 0; i < workers; i++ {
		client, _ := connection.Conn()
		clients[i] = client
	}

	start := time.Now()
	go runPrintMetics(&latencyTimer, &reqCounter)
	// s := "CgObKRQxNTQ3MzQzMTc4MDU1NzM4MjU5MAAA"
	// numOfError := 0

	// for x := 0; x < workers; x++ {
	// 	go request(x, jobs, results, csvData, clients[1], &latencyTimer, &reqCounter)
	// }
	go request(jobs, results, csvData, clients, &latencyTimer, &reqCounter)

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

func request(jobs <-chan int, results chan<- string, csvData [][]string, clients []*connection.IdHashClient, latencyTimer *metrics.Timer, reqCounter *metrics.Counter) {

	if len(clients) == 0 {
		return
	}
	for j := range jobs {
		start := time.Now()
		index := j % len(clients)
		response, err := clients[index].Exec(csvData[j][0])
		if err != nil {
			log.Printf("Error : %s", err)
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

		fmt.Printf("%s latency  %f,%f,%f, mean %f, num of requests %d", now, latencyMetric.Percentile(0.50)/1000000, latencyMetric.Percentile(0.95)/1000000, latencyMetric.Percentile(0.99)/1000000, latencyMetric.Mean()/1000000, counterMetric.Count())
		fmt.Println()
	}
}
