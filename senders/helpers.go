package senders

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"math/rand/v2"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/jamesthecog/faker/generator"
)

// Sends the map to be sent to the destination address with Go's
// native packages.
func SendPayload(dest string, payload map[string]any) {
	jsonedData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln(err)
	}

	response, err := http.Post(dest, "application/json", bytes.NewBuffer(jsonedData))
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	log.Printf("Message sent! (status %d)\n", response.StatusCode)
}

// The worker for the concurrent sending mode - separated out
// so that the main function of interest doesn't become too bloated.
// Like yo momma's stomach when yous was inside her.
func worker(ctx context.Context, wg *sync.WaitGroup, dest string, counter *atomic.Int64) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			payload, err := generator.FakeData()
			if err != nil {
				log.Printf("Failed to send a response for the following reason: %v", err)
				continue
			}
			SendPayload(dest, payload)
			counter.Add(1)

			jitter := time.Duration(latency + rand.IntN(latency))
			time.Sleep(time.Millisecond * jitter)
		}
	}
}
