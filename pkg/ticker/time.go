package ticker

import "time"

func Time() *TimeTicker {
	return &TimeTicker{
		last: time.Now(),
	}
}

type TimeTicker struct {
	last time.Time
}

func (d *TimeTicker) Tick() time.Duration {
	now := time.Now()
	dt := now.Sub(d.last)
	d.last = now
	return dt
}
