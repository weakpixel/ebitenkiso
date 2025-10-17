package effect

import (
	"github.com/weakpixel/ebitenkiso/pkg/sprites"

	"github.com/hajimehoshi/ebiten/v2"
)

type Ghost struct {
	instances []instance
}

type instance struct {
	x   float64
	y   float64
	dir bool
}

func (g *Ghost) Reset() {
	g.instances = []instance{}
}

func (g *Ghost) Add(x, y float64, dir bool) {
	g.instances = append(g.instances, instance{
		x:   x,
		y:   y,
		dir: dir,
	})

}

func (g *Ghost) Draw(sprite *sprites.Sprite, screen *ebiten.Image) {
	colorScale := ebiten.ColorScale{}
	colorScale.ScaleAlpha(0.3)
	for _, pos := range g.instances {
		sprite.Draw(pos.x, pos.y, pos.dir, screen, colorScale)
	}
}
