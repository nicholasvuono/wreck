package wreck

import (
	"errors"
	"time"

	"github.com/mxschmitt/playwright-go"
)

type WebTest struct {
	Test       func() map[string]int64
	Duration   int
	Iterations int
}

func (w *WebTest) Run() []map[string]int64 {
	var results []map[string]int64
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

func webTestIterations(iterations int, f func() map[string]int64) map[string]int64 {
	var results map[string]int64
	for i := 0; i < iterations; i++ {
		results = f()
	}
	return results
}

func webTestDuration(duration int, f func() map[string]int64) map[string]int64 {
	var results map[string]int64
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

func step(label string, f func(*playwright.Page) *playwright.Page, page *playwright.Page) (string, int64, *playwright.Page) {
	start := time.Now()
	page = f(page)
	return label, time.Since(start).Milliseconds(), page
}
