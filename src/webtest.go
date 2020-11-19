package wreck

import (
	"github.com/mxschmitt/playwright-go"
)

type WebTest struct {
	Test     func() map[string]float64
	Duration int
	Results  map[string]float64
}

func (w *WebTest) Run() {
	results := w.Test()
	w.Results = results
}

func group(label string, f func(*playwright.Page) *playwright.Page) {

}
