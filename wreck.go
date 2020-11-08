package wreck

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/nicholasvuono/wreckhttp"
)

type Request = wreckhttp.Request

type Options struct {
	Vus        int
	Duration   int
	Iterations int
}

var wg sync.WaitGroup

var responses []string

func Batch(options Options, requests []Request) []string {
	if options.Iterations != 0 && options.Duration == 0 {
		concurrrentBatchIterations(options, requests)
	} else if options.Iterations == 0 {
		concurrentBatchDuration(options, requests)
	} else {
		err := errors.New("Error Options: Duration and Iteration cannot be used at the same time")
		explain(err)
	}
	return responses
}

func concurrentBatchDuration(options Options, requests []Request) {
	after := time.Now().Add(time.Duration(options.Duration) * time.Second)
	for {
		now := time.Now()
		for i := 0; i < options.Vus; i++ {
			wg.Add(1)
			go sendBatch(requests)
		}
		if now.After(after) {
			break
		}
		wg.Wait()
	}
}

func concurrrentBatchIterations(options Options, requests []Request) {
	for i := 0; i < options.Iterations; i++ {
		for i := 0; i < options.Vus; i++ {
			wg.Add(1)
			go sendBatch(requests)
		}
		wg.Wait()
	}
}

func sendBatch(requests []Request) {
	defer wg.Done()
	batch, err := wreckhttp.Batch(requests)
	explain(err)
	responses = append(responses, fmt.Sprintf("%v", batch.Send()))
}
