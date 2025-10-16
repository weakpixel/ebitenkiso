package sprites

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewSprite(sheet *SpriteSheet) *Sprite {
	return &Sprite{
		sheet: sheet,
	}
}

type Sprite struct {
	sheet  *SpriteSheet
	anim   *Animation
	Src    string
	Shader Shader
}

func (s *Sprite) SpriteSheet() *SpriteSheet {
	return s.sheet
}

func (s *Sprite) Animation() *Animation {
	return s.anim
}

func (s *Sprite) SetAnimation(name string, loop bool) {
	if s.anim == nil || name != s.anim.Name {
		s.anim = s.sheet.Animation(name)
		s.anim.Loop = loop
	}
}

func (s *Sprite) Update(dt time.Duration) {
	if s.anim != nil {
		s.anim.Update(dt)
		if s.Shader != nil {
			s.Shader.Update(dt)
		}
	}

}

func (s *Sprite) Draw(x float64, y float64, flipH bool, screen *ebiten.Image, colorScale *ebiten.ColorScale) {
	if s.anim != nil {
		img := s.anim.Image()

		if colorScale == nil {
			colorScale = &ebiten.ColorScale{}
		}

		geom := ebiten.GeoM{}
		if flipH {
			geom.Scale(-1, 1)
			geom.Translate(float64(img.Bounds().Dx()), 0)
		}
		geom.Translate(x, y)

		if s.Shader == nil {
			opt := &ebiten.DrawImageOptions{
				GeoM:       geom,
				ColorScale: *colorScale,
			}
			screen.DrawImage(img, opt)
		} else {
			// screen.DrawImage(img, &ebiten.DrawImageOptions{
			// 	GeoM:       geom,
			// 	ColorScale: *colorScale,
			// })

			opt := &ebiten.DrawRectShaderOptions{
				GeoM:       geom,
				ColorScale: *colorScale,
			}
			s.Shader.Draw(img, screen, opt)
		}
	}
}

type Shader interface {
	Draw(srcImage *ebiten.Image, screen *ebiten.Image, op *ebiten.DrawRectShaderOptions)
	Update(dt time.Duration)
}
