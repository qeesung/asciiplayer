// progress package show the progress bar when deal some busy
// things, and can register new bar then wait all bars finished.
package progress

import (
	"github.com/gosuri/uiprogress"
	"sync"
)

// MaxSteps define the max steps for a bar.
var MaxSteps = 10000

// WaitingBar extend the wait group, and register a new bar
// then call Wait() method to wait all bars to finished.
type WaitingBar struct {
	sync.WaitGroup
}

// Start method start show the progress bar
func (*WaitingBar) Start() {
	uiprogress.Start()
}

// Stop method stop the progress bar
func (*WaitingBar) Stop() {
	uiprogress.Stop()
}

// AddBar add a new bar, then return a notifer channel that can report the 
// progress back, then progress bar will be updated in another goroutine.
func (waitingBar *WaitingBar) AddBar(barName string, steps int) chan<- int {
	notifier := make(chan int)
	bar := uiprogress.AddBar(steps).AppendCompleted().PrependElapsed()
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return barName + ": "
	})

	go func() {
		waitingBar.Add(1)
		defer waitingBar.Done()
		for range notifier {
			bar.Incr()
		}
		bar.Set(MaxSteps)
	}()
	return notifier
}
