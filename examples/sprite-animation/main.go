package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/weakpixel/ebitenkiso/pkg/res"
	"github.com/weakpixel/ebitenkiso/pkg/shader"
	"github.com/weakpixel/ebitenkiso/pkg/sprites"
	"github.com/weakpixel/ebitenkiso/pkg/sprites/aseprite"
	"github.com/weakpixel/ebitenkiso/pkg/ticker"
	"github.com/weakpixel/ebitenkiso/pkg/tween"

	"image/color"
	_ "image/jpeg"
	_ "image/png"
)

func main() {
	start()
}

func start() {

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Sprite Animation Example")

	g := Game{
		// sprite: sprite,
		ticker: ticker.Time(),
	}
	err := ebiten.RunGame(&g)
	if err != nil {
		panic(err)
	}
}

var (
	pixels     = float64(350)
	moveTween  = tween.NewSeq(tween.New(0, pixels, time.Second, tween.InExpo), tween.New(pixels, 0, time.Second, tween.InExpo))
	moveTweenY = tween.NewSeq(tween.New(0, pixels, time.Second, tween.InExpo), tween.New(pixels, 0, time.Second, tween.InExpo))

	noise = shader.NewNoise(time.Millisecond * 50)
	other = shader.NewAbberation(10)
)

type Game struct {
	sprite *sprites.Sprite
	ticker ticker.Ticker
	x, y   float64
}

func (g *Game) Update() error {
	if g.sprite == nil {
		spriteUri := res.MustParse("examples/sprite-animation/run-cycle-48x48.json")
		sprite, err := aseprite.LoadSprite(spriteUri)
		if err != nil {
			panic(err)
		}
		moveTween.SetLoop(true)
		moveTweenY.SetLoop(true)
		sprite.Shader = shader.NewAbberation(3)

		g.sprite = sprite
	}

	dt := g.ticker.Tick()
	noise.Update(dt)
	other.Update(dt)
	g.x, _ = moveTween.Update(dt)
	g.y, _ = moveTweenY.Update(dt)
	g.sprite.Update(dt)
	return nil
}

var buff = ebiten.NewImage(800, 600)
var buff2 = ebiten.NewImage(800, 600)

func init() {
	buff.Fill(color.RGBA{0, 0, 0, 255})
}
func (g *Game) Draw(screen *ebiten.Image) {
	noise.Draw(buff, buff2, &ebiten.DrawRectShaderOptions{})
	other.Draw(buff2, screen, &ebiten.DrawRectShaderOptions{})
	g.sprite.Draw(g.x, g.y, false, screen, ebiten.ColorScale{})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}
