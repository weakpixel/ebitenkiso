package tween

import (
	"time"
)

func New(begin, end float64, duration time.Duration, easing TweenFunc) *Tween {
	return &Tween{
		begin:    begin,
		end:      end,
		change:   end - begin,
		duration: duration,
		easing:   easing,
	}
}

type Updater interface {
	Update(dt time.Duration) (val float64, done bool)
}

type Tween struct {
	begin    float64
	end      float64
	change   float64
	duration time.Duration
	elapsed  time.Duration
	easing   TweenFunc
	loop     bool
}

func (t *Tween) SetLoop(loop bool) {
	t.loop = loop
}

func (t *Tween) Reset() {
	t.elapsed = 0
}

func (t *Tween) Update(dt time.Duration) (float64, bool) {
	t.elapsed += dt
	done := t.elapsed > t.duration
	var curTime float64
	if done {
		curTime = t.duration.Seconds()
		if t.loop {
			t.elapsed = t.elapsed - t.duration
		}
	} else {
		curTime = t.elapsed.Seconds()
	}
	val := t.easing(curTime, t.begin, t.change, t.duration.Seconds())
	return val, done
}
