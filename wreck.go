package wreck

import (
	"fmt"
	"time"

	"github.com/nicholasvuono/wreckhttp"
)

type Request = wreckhttp.Request

type Options struct {
	Vus      int
	Duration int
}

func Batch(options Options, requests []Request) []string {
	var responses []string
	now := time.Now()
	after := now.Add(time.Duration(options.Duration) * time.Second)
	for {
		now = time.Now()
		semaphoreChan := make(chan struct{}, options.Vus)
		responsesChan := make(chan string)

		defer func() {
			close(semaphoreChan)
			close(responsesChan)
		}()

		for i := 0; i < options.Vus; i++ {
			go func() {
				batch, err := wreckhttp.Batch(requests)
				explain(err)
				response := batch.Send()
				responseString := fmt.Sprintf("%v", response)
				responsesChan <- responseString
				<-semaphoreChan
			}()
		}

		for {
			response := <-responsesChan
			responses = append(responses, response)
			if len(responsesChan) == 0 {
				break
			}
		}
		if now.After(after) {
			break
		}
	}
	return responses
}
