package wreck

import (
	"errors"

	"github.com/mxschmitt/playwright-go"
)

type WebTest struct {
	Test       func() map[string]float64
	Duration   int
	Iterations int
}

func (w *WebTest) Run() map[string]float64 {
	var results map[string]float64
	if w.Iterations != 0 && w.Duration == 0 {
		results = webTestIterations(w.Iterations, w.Test)
	} else if w.Iterations == 0 {
		results = webTestDuration(w.Duration, w.Test)
	} else {
		err := errors.New("Error Options: Duration and Iteration cannot be used at the same time")
		explain(err)
	}
	return results
}

func webTestIterations(iterations int, f func() map[string]float64) map[string]float64 {
	//Implement iteration logic similar to batch just without concurrency
	return nil
}

func webTestDuration(duration int, f func() map[string]float64) map[string]float64 {
	//Implement duration logic similar to batch just without concurrency
	return nil
}

func group(label string, f func(*playwright.Page) *playwright.Page) {
	//implment group logic to capture timings for different poage object or events
	//within the browser level user test
}
