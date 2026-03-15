package main

import (
	"flag"
	"log"
	"runtime"
	"strings"

	"github.com/jamesthecog/faker/generator"
	"github.com/jamesthecog/faker/senders"
)

// NB: I've directly used generator.FakeData() for single-send modes - because there's
// literally no reason why we can't just send the request ourselves here (i.e., not make a
// helper function elsewhere) when the utilities needed to do so are already present PLUS
// it's not meaningfully different like the other modes.
func main() {
	dest := flag.String("dest", "", "The endpoint or URL to send fake data to.")
	sendMode := flag.String("mode", "single", "The kind of request to simulate ('single', 'batch', 'high_volume', and 'concurrent').")
	numRequests := flag.Int("num_requests", 5, "How many requests should be sent (only applicable for batch requests)?")
	requestDuration := flag.Int("duration", 10, "How long should requests be sent for (only applicable for streaming requests)?")
	numWorkers := flag.Int("num_workers", max(1, runtime.NumCPU()-3), "The amount of concurrent processes to spawn for concurrent sending")
	flag.Parse()

	sendChoice := strings.TrimSpace(strings.ToLower(*sendMode))

	if strings.TrimSpace(*dest) == "" {
		log.Fatalln("`dest` cannot be an empty string!")
	}
	if *numWorkers == 1 && sendChoice == "concurrent" {
		log.Fatalln("Your machine's too weak to be doing things concurrently - try one of the other modes instead!")
	}

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
	default:
		log.Fatalln("'mode' needs to be one of the following: 'single', 'batch', 'high_volume', or 'concurrent'!")
	}
}
