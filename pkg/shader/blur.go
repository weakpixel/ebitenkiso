package shader

import (
	_ "embed"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed kage/blur.kage
	blurRaw []byte
	blur    = mustNewShader(blurRaw)

	//go:embed kage/abberation.kage
	abbrRaw []byte
	abbr    = mustNewShader(abbrRaw)

	//go:embed kage/shape-renderer.kage
	shapeRaw []byte
	shape    = mustNewShader(shapeRaw)

	//go:embed kage/blur-radial.kage
	blurRadialRaw []byte
	blurRadial    = mustNewShader(blurRadialRaw)
)

func NewShapeRenderer() *BlurShader {
	return &BlurShader{
		shader: shape,
	}
}

func NewAbberation(offset float32) *BlurShader {
	return &BlurShader{
		shader:      abbr,
		Offset:      offset,
		PosX:        offset,
		PosY:        offset,
		mouseUpdate: false,
	}
}

func NewBlurRadial(offset float32) *BlurShader {
	return &BlurShader{
		shader:      blurRadial,
		Offset:      offset,
		mouseUpdate: true,
	}
}

func NewBlur(offset float32) *BlurShader {
	return &BlurShader{
		shader:      blur,
		Offset:      offset,
		mouseUpdate: true,
	}
}

type BlurShader struct {
	shader      *ebiten.Shader
	PosX        float32
	PosY        float32
	Offset      float32
	mouseUpdate bool
}

func (s *BlurShader) Update(dt time.Duration) {
	if s.mouseUpdate {
		x, y := ebiten.CursorPosition()
		s.PosX = float32(x) + s.Offset
		s.PosY = float32(y) + s.Offset
	}
}

func (s *BlurShader) Draw(srcImage *ebiten.Image, screen *ebiten.Image, op *ebiten.DrawRectShaderOptions) {
	w, h := srcImage.Bounds().Dx(), srcImage.Bounds().Dy()
	op.Uniforms = map[string]any{
		"Value": []float32{s.PosX, s.PosY},
	}
	op.Images[0] = srcImage
	screen.DrawRectShader(w, h, s.shader, op)
}
