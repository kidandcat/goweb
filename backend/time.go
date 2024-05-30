package backend

import "time"

var ClockClient struct {
	Now func() (time.Time, error)
}

type Clock struct{}

func (h *Clock) Now() time.Time {
	return time.Now()
}
