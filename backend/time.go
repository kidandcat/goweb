package backend

import "time"

var ClockClient struct {
	Read func() (time.Time, error)
}

type Clock struct{}

func (h *Clock) Read() time.Time {
	return time.Now()
}
