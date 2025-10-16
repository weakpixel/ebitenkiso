package shader

import (
	_ "embed"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed kage/default.kage
	defRaw []byte
	def    = mustNewShader(defRaw)
)

func mustNewShader(raw []byte) *ebiten.Shader {
	s, err := ebiten.NewShader(raw)
	if err != nil {
		panic(err)
	}
	return s
}

func NewDefault() *DefaultShader {
	return &DefaultShader{
		shader: def,
	}
}

type DefaultShader struct {
	shader *ebiten.Shader
}

func (s *DefaultShader) Update(dt time.Duration) {}

func (s *DefaultShader) Draw(srcImage *ebiten.Image, screen *ebiten.Image, op *ebiten.DrawRectShaderOptions) {
	w, h := srcImage.Bounds().Dx(), srcImage.Bounds().Dy()
	op.Images[0] = srcImage
	screen.DrawRectShader(w, h, s.shader, op)

}

func Group(s ...Shader) Shader {
	return &ShaderGroup{
		shaders: s,
	}
}

type ShaderGroup struct {
	shaders     []Shader
	bufferImage *ebiten.Image
	useBuffer   bool
}

func (s *ShaderGroup) bufferSize() (int, int) {
	if s.bufferImage == nil {
		return 0, 0
	}
	return s.bufferImage.Bounds().Dx(), s.bufferImage.Bounds().Dy()
}

func (s *ShaderGroup) buffer(srcImage *ebiten.Image) *ebiten.Image {
	w, h := srcImage.Bounds().Dx(), srcImage.Bounds().Dy()
	ww, hh := s.bufferSize()
	if w != ww || hh != h {
		s.bufferImage = ebiten.NewImage(w, h)

	}
	return s.bufferImage

}

func (s *ShaderGroup) Update(dt time.Duration) {
	for _, s := range s.shaders {
		s.Update(dt)
	}
}

func (s *ShaderGroup) Draw(srcImage *ebiten.Image, screen *ebiten.Image, op *ebiten.DrawRectShaderOptions) {
	if !s.useBuffer {
		for _, s := range s.shaders {
			s.Draw(srcImage, screen, op)
		}
		return
	}

	buffer := s.buffer(srcImage)
	for _, s := range s.shaders {
		s.Draw(srcImage, buffer, op)
		srcImage.DrawImage(buffer, nil)
	}
	screen.DrawImage(buffer, nil)
}

type Shader interface {
	Draw(srcImage *ebiten.Image, screen *ebiten.Image, op *ebiten.DrawRectShaderOptions)
	Update(dt time.Duration)
}
