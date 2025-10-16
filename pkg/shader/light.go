package shader

import (
	_ "embed"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed kage/light2.kage
	lightRaw []byte
	light    = mustNewShader(lightRaw)
)

func NewLight(offset float32) *LightShader {
	return &LightShader{
		shader: light,
		Offset: offset,
	}
}

type LightShader struct {
	shader *ebiten.Shader
	PosX   float32
	PosY   float32
	Offset float32
}

func (s *LightShader) Update(dt time.Duration) {
	x, y := ebiten.CursorPosition()
	s.PosX = float32(x) + s.Offset
	s.PosY = float32(y) + s.Offset

}

func (s *LightShader) Draw(srcImage *ebiten.Image, screen *ebiten.Image, op *ebiten.DrawRectShaderOptions) {
	w, h := srcImage.Bounds().Dx(), srcImage.Bounds().Dy()
	op.Uniforms = map[string]any{
		"Value": []float32{s.PosX, s.PosY},
	}
	op.Images[0] = srcImage
	screen.DrawRectShader(w, h, s.shader, op)
}
