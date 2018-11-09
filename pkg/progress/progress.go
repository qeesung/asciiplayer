package progress

import (
	"github.com/gosuri/uiprogress"
	"sync"
)

var MaxSteps = 10000

type WaitingBar struct {
	sync.WaitGroup
}

func (*WaitingBar) Start() {
	uiprogress.Start()
}

func (*WaitingBar) Stop() {
	uiprogress.Stop()
}

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
