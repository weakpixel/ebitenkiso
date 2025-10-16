package sprites

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Frame struct {
	Image    *ebiten.Image
	Duration time.Duration
	Width    int
	Height   int
}
