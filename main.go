package main

import (
	"flag"
	"log"
	"runtime"
	"strings"

	"github.com/jamesthecog/faker/generator"
	"github.com/jamesthecog/faker/senders"
)

func main() {
	dest := flag.String("dest", "", "The endpoint or URL to send fake data to.")
	sendMode := flag.String("mode", "single", "The kind of request to simulate ('single', 'batch', 'high_volume', and 'concurrent').")
	numRequests := flag.Int("num_requests", 5, "How many requests should be sent (only applicable for batch requests)?")
	requestDuration := flag.Int("duration", 10, "How long should requests be sent for (only applicable for streaming requests)?")
	numWorkers := flag.Int("num_workers", max(1, runtime.NumCPU()-2), "The amount of concurrent processes to spawn for concurrent sending")
	flag.Parse()

	sendChoice := strings.ToLower(strings.ToLower(*sendMode))
	switch sendChoice {
	case "single":
		payload, err := generator.FakeData()
		if err != nil {
			log.Fatalf("Could not send the payload for the following reason: %v", err)
		}
		senders.SendPayload(*dest, payload)
	case "batch":
		senders.SendBatch(*dest, *numRequests)
	case "high_volume":
		senders.HighVolume(*dest, *requestDuration)
	case "concurrent":
		senders.ConcurrentSending(*dest, *numWorkers, *requestDuration)
	}
}
