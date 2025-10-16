package shader

import (
	_ "embed"
	"math/rand"
	"time"

	"github.com/weakpixel/ebitenkiso/pkg/timer"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed kage/noise.kage
	noiseRaw []byte
	noise    = mustNewShader(noiseRaw)
)

func NewNoise(delay time.Duration) *NoiseShader {
	return &NoiseShader{
		shader: noise,
		t:      timer.New(delay),
	}
}

type NoiseShader struct {
	shader *ebiten.Shader
	seed   float32
	Invert bool
	t      *timer.Timer
}

func (s *NoiseShader) Update(dt time.Duration) {
	done := s.t.Update(dt)
	if done {
		s.t.Reset()
		min, max := float32(1.0), float32(8000)
		s.seed = min + rand.Float32()*(max-min)
	}
}

func (s *NoiseShader) Draw(srcImage *ebiten.Image, screen *ebiten.Image, op *ebiten.DrawRectShaderOptions) {
	op.Images[0] = srcImage
	inv := 0.0
	if s.Invert {
		inv = 1.0
	}
	op.Uniforms = map[string]any{
		"Seed":   s.seed,
		"Invert": inv,
	}
	w, h := srcImage.Bounds().Dx(), srcImage.Bounds().Dy()
	screen.DrawRectShader(w, h, s.shader, op)
}
