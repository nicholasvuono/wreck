package main

import (
	"fmt"
	"wreckhttp"
)

func main() {

	requests := []wreckhttp.Request{
		wreckhttp.Request{
			Method:  "GET",
			URL:     "http://github.com/nicholasvuono",
			Headers: nil,
			Body:    nil,
		},
		wreckhttp.Request{
			Method: "POST",
			URL:    "https://httpbin.org/post",
			Headers: map[string][]string{
				"Accept": []string{"application/json"},
			},
			Body: map[string]string{
				"name":  "Test API Guy",
				"email": "testapiguy@email.com",
			},
		},
	}
	batch := wreckhttp.NewBatch(requests)
	responses := batch.Send()
	for _, res := range responses {
		fmt.Println(res.GetResponse())
	}
}
