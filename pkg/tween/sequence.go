package tween

import (
	"time"
)

func NewSeq(tweens ...*Tween) *Sequence {
	done := make([]bool, len(tweens))
	return &Sequence{
		tweens: tweens,
		done:   done,
	}
}

type Sequence struct {
	tweens []*Tween
	done   []bool
	loop   bool
	// lastVal float64
}

func (s *Sequence) Reset() {
	for idx := range s.done {
		s.done[idx] = false
		s.tweens[idx].Reset()
	}
}

func (s *Sequence) SetLoop(loop bool) {
	s.loop = loop
}

func (s *Sequence) Update(dt time.Duration) (float64, bool) {
	last := len(s.tweens) - 1
	for idx, t := range s.tweens {
		if s.done[idx] {
			continue
		}
		val, done := t.Update(dt)
		if done {
			s.done[idx] = true
		}
		return val, done && idx == last
	}
	if s.loop {
		s.Reset()
		return s.tweens[0].Update(dt)
	}
	return s.tweens[last].Update(dt)
}
