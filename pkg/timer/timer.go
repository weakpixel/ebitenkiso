package timer

import (
	"time"
)

func New(duration time.Duration) *Timer {
	return &Timer{
		duration: duration,
		elapsed:  0,
	}
}

type Timer struct {
	duration time.Duration
	elapsed  time.Duration
}

func (t *Timer) Reset() {
	t.elapsed = 0
}

func (t *Timer) Duration() time.Duration {
	return t.duration
}

func (t *Timer) Done() bool {
	return t.elapsed >= t.duration
}

func (t *Timer) Update(dt time.Duration) bool {
	t.elapsed = t.elapsed + dt
	return t.Done()
}
