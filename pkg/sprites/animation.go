package sprites

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	Frames []Frame
	Name   string
	Loop   bool

	frameIndex int
	elapsed    time.Duration
}

func (a *Animation) Done() bool {
	return !a.Loop && a.frameIndex == len(a.Frames)-1 && a.elapsed == 0
}

func (a *Animation) Reset() {
	a.frameIndex = 0
	a.elapsed = 0
}

func (a *Animation) Update(dt time.Duration) {
	if !a.Loop && a.frameIndex == len(a.Frames)-1 {
		return
	}
	a.elapsed += dt
	for {
		frameDuration := a.Frames[a.frameIndex].Duration
		if a.elapsed < frameDuration {
			break
		}
		a.elapsed -= frameDuration
		a.frameIndex++

		if a.frameIndex >= len(a.Frames) {
			if a.Loop {
				a.frameIndex = 0
			} else {
				a.frameIndex = len(a.Frames) - 1
				a.elapsed = 0
				return
			}
		}
	}
}

func (a *Animation) Image() *ebiten.Image {
	if len(a.Frames) == 0 {
		return nil
	}
	return a.Frames[a.frameIndex].Image

}
