package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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
	var (
		dest, mode                                           string
		numRequests, requestDuration, numWorkers, defWorkers int
	)
	defWorkers = max(1, runtime.NumCPU()-2)
	flag.StringVar(&dest, "l", "", "")
	flag.StringVar(&dest, "location", "", "")
	flag.StringVar(&mode, "m", "single", "")
	flag.StringVar(&mode, "mode", "single", "")
	flag.IntVar(&numRequests, "nr", 5, "")
	flag.IntVar(&numRequests, "num_reqs", 5, "")
	flag.IntVar(&requestDuration, "rd", 10, "")
	flag.IntVar(&requestDuration, "duration", 10, "")
	flag.IntVar(&numWorkers, "nw", defWorkers, "")
	flag.IntVar(&numWorkers, "num_workers", defWorkers, "")

	flag.Usage = func() {
		programName := strings.Split(os.Args[0], "\\")
		fmt.Fprintf(os.Stderr, "Usage instructions for %s:\n", strings.TrimSpace(programName[len(programName)-1]))
		fmt.Fprint(os.Stderr, "  -l, -location string\n    \tThe destination address you want the payload sent to.\n")
		fmt.Fprint(os.Stderr, "  -m, -mode string\n    \tThe sending mode you want the CLI app to simulate: `single`, `batch`, `high_vol`, or "+
			"`concurrent` (default: single).\n")
		fmt.Fprint(os.Stderr, "  -nr, -num_reqs integer\n    \tThe amount of requests to batch-send during `batch` sending mode (default: 5).\n")
		fmt.Fprint(os.Stderr, "  -d, -duration string\n    \tHow many seconds should the `high_vol` or `concurrent` mode be sending for (default: 10 seconds)?\n")
		fmt.Fprintf(os.Stderr, "  -nw, -num_workers integer\n    \tThe amount of workers to use during `concurrent` sending mode (default: %d workers).\n",
			defWorkers)
	}
	flag.Parse()

	sendChoice := strings.TrimSpace(strings.ToLower(mode))

	if dest == "" || mode == "" || numRequests == 0 || requestDuration == 0 || numWorkers == 0 {
		flag.Usage()
		os.Exit(0)
	}
	if numWorkers == 1 && sendChoice == "concurrent" {
		log.Fatalln("Your machine's too weak to be doing things concurrently - try one of the other modes instead!")
	}

	switch sendChoice {
	case "single":
		payload, err := generator.FakeData()
		if err != nil {
			log.Fatalf("Could not send the payload for the following reason: %v", err)
		}
		senders.SendPayload(dest, payload)
	case "batch":
		senders.SendBatch(dest, numRequests)
	case "high_vol":
		senders.HighVolume(dest, requestDuration)
	case "concurrent":
		senders.ConcurrentSending(dest, numWorkers, requestDuration)
	default:
		log.Fatalln("`mode` or `m` needs to be one of the following: `single`, `batch`, `high_vol`, or `concurrent`!")
	}
}
