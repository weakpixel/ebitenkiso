package sprites

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	Frames []Frame
	Name   string
	Loop   bool

	curIndex int
	elapsed  time.Duration
}

func (a *Animation) Reset() {
	a.curIndex = 0
	a.elapsed = 0
}

func (a *Animation) Update(dt time.Duration) {
	if !a.Loop && a.curIndex == len(a.Frames)-1 {
		return
	}
	a.elapsed += dt
	frameDuration := a.Frames[a.curIndex].Duration
	if a.elapsed >= frameDuration {
		a.elapsed = a.elapsed - frameDuration
		newIndex := a.curIndex + 1
		if newIndex >= len(a.Frames) {
			newIndex = 0
		}
		a.curIndex = newIndex
	}
}

func (a *Animation) Image() *ebiten.Image {
	return a.Frames[a.curIndex].Image

}
