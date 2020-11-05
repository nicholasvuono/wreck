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

func Batch(options Options, requests []Request) []string {
	var responses []string

	if options.Iterations != 0 && options.Duration == 0 {
		var wg sync.WaitGroup
		for i := 0; i < options.Iterations; i++ {
			for i := 0; i < options.Vus; i++ {
				wg.Add(1)
				go func([]string) {
					defer wg.Done()
					batch, err := wreckhttp.Batch(requests)
					explain(err)
					response := batch.Send()
					responseString := fmt.Sprintf("%v", response)
					responses = append(responses, responseString)
				}(responses)
			}
		}
	} else if options.Iterations == 0 {
		now := time.Now()
		after := now.Add(time.Duration(options.Duration) * time.Second)

		for {
			var wg sync.WaitGroup
			now := time.Now()
			for i := 0; i < options.Vus; i++ {
				wg.Add(1)
				go func([]string) {
					defer wg.Done()
					batch, err := wreckhttp.Batch(requests)
					explain(err)
					response := batch.Send()
					responseString := fmt.Sprintf("%v", response)
					responses = append(responses, responseString)
				}(responses)
			}
			if now.After(after) {
				break
			}
			wg.Wait()
		}
	} else {
		err := errors.New("Error Options: Duration and Iteration cannot be used at the same time")
		explain(err)
	}
	return responses
}
