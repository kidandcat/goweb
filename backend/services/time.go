package services

import "time"

var ClockClient struct {
	Read func() (time.Time, error)
}

type Clock struct{}

func NewClockService() *Clock {
	return &Clock{}
}

func (h *Clock) Read() time.Time {
	return time.Now()
}
