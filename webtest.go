package wreck

import (
	"errors"
	"time"

	"github.com/mxschmitt/playwright-go"
)

type WebTest struct {
	Test       func() map[string]float64
	Duration   int
	Iterations int
}

func (w *WebTest) Run() []map[string]float64 {
	var results []map[string]float64
	if w.Iterations != 0 && w.Duration == 0 {
		results = append(results, webTestIterations(w.Iterations, w.Test))
	} else if w.Iterations == 0 {
		results = append(results, webTestDuration(w.Duration, w.Test))
	} else {
		err := errors.New("Error Options: Duration and Iteration cannot be used at the same time")
		explain(err)
	}
	return results
}

func webTestIterations(iterations int, f func() map[string]float64) map[string]float64 {
	var results map[string]float64
	for i := 0; i < iterations; i++ {
		results = f()
	}
	return results
}

func webTestDuration(duration int, f func() map[string]float64) map[string]float64 {
	var results map[string]float64
	after := time.Now().Add(time.Duration(duration) * time.Second)
	for {
		now := time.Now()
		results = f()
		if now.After(after) {
			break
		}
	}
	return results
}

func step(label string, f func(*playwright.Page) *playwright.Page, page *playwright.Page) *playwright.Page {
	defer timeIt(time.Now(), label)
	page = f(page)
	return page
}
