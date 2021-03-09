package timer

import (
	"time"
)

type Timer struct {
	ticker *time.Ticker
	done   chan bool
	f      handler
}

type handler func(t time.Time) bool

func New(d time.Duration) *Timer {
	return &Timer{
		ticker: time.NewTicker(d),
		done:   make(chan bool),
	}
}

func (timer *Timer) Start(f func(t time.Time) bool) {
	timer.f = f
	for {
		select {
		case <-timer.done:
			return
		case t := <-timer.ticker.C:
			if f(t) {
				return
			}
		}
	}
}

func (timer *Timer) Reset(d time.Duration)  {
	timer.ticker.Reset(d)
}

func (timer *Timer) Restart() {
	timer.Stop()
	timer.Start(timer.f)
}

func (timer *Timer) Stop() {
	timer.done <- true
}
