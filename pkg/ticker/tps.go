package ticker

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func TPS() *TPSTicker {
	return &TPSTicker{}
}

type TPSTicker struct {
}

func (d *TPSTicker) Tick() time.Duration {
	dt := float64(1) / float64(ebiten.TPS()) * 1000
	return time.Duration(dt) * time.Millisecond
}
