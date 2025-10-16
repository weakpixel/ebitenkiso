package shader

import (
	_ "embed"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed kage/glitch.kage
	glitch1Raw []byte
	glitch1    = mustNewShader(glitch1Raw)
)

func NewGlitchShader() *GlitchShader {
	return &GlitchShader{
		shader: glitch1,
		MinVal: -200,
		MaxVal: 200,
		values: [2]float32{},
	}
}

type GlitchShader struct {
	shader *ebiten.Shader
	values [2]float32
	MinVal float32
	MaxVal float32
	cout   int
}

func (s *GlitchShader) Update(dt time.Duration) {
	s.cout += 1
	for i := 0; i < 2; i++ {
		s.values[i] = s.MinVal + rand.Float32()*(s.MaxVal-s.MinVal)
	}
}

func (s *GlitchShader) Draw(srcImage *ebiten.Image, screen *ebiten.Image, op *ebiten.DrawRectShaderOptions) {
	w, h := srcImage.Bounds().Dx(), srcImage.Bounds().Dy()
	op.Uniforms = map[string]any{
		"Value": s.values,
	}
	op.Images[0] = srcImage
	screen.DrawRectShader(w, h, s.shader, op)
}
