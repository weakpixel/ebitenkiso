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
	Shader Shader
	geoM   ebiten.GeoM
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
	}
	if s.Shader != nil {
		s.Shader.Update(dt)
	}

}

func (s *Sprite) Draw(x float64, y float64, flipH bool, screen *ebiten.Image, colorScale ebiten.ColorScale) {
	if s.anim != nil {
		img := s.anim.Image()

		s.geoM.Reset()
		if flipH {
			w := float64(img.Bounds().Dx())
			s.geoM.Scale(-1, 1)
			s.geoM.Translate(w, 0)
		}

		s.geoM.Translate(x, y)

		if s.Shader == nil {
			screen.DrawImage(img, &ebiten.DrawImageOptions{
				GeoM:       s.geoM,
				ColorScale: colorScale,
			})
		} else {
			s.Shader.Draw(img, screen, &ebiten.DrawRectShaderOptions{
				GeoM:       s.geoM,
				ColorScale: colorScale,
			})
		}
	}
}

type Shader interface {
	Draw(srcImage *ebiten.Image, screen *ebiten.Image, op *ebiten.DrawRectShaderOptions)
	Update(dt time.Duration)
}
