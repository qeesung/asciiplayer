package progress

import (
	"github.com/gosuri/uiprogress"
)

var MaxSteps = 10000

type WaitingBar struct {
}

func (*WaitingBar) Start() {
	uiprogress.Start()
}

func (waitingBar *WaitingBar) AddBar(barName string, steps int) chan<- int {
	notifier := make(chan int)
	bar := uiprogress.AddBar(steps).AppendCompleted().PrependElapsed()
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return barName + ": "
	})

	go func() {
		for range notifier {
			bar.Incr()
		}
		bar.Set(MaxSteps)
	}()
	return notifier
}
