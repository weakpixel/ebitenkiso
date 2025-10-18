package main

import (
	"bytes"
	"fmt"
	"image"
	"os"
	"path/filepath"
	"time"

	_ "image/jpeg"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 320
	screenHeight = 180
)

var (
	runnerImage   *ebiten.Image
	curShader     *ebiten.Shader
	mixColor      = false
	enabledShader = true
	threshold     = 0.1
	radius        = 12
	blurRadius    = 6
)

type Game struct {
}

func (g *Game) Update() error {

	if inpututil.IsKeyJustReleased(ebiten.Key1) {
		enabledShader = !enabledShader
	}

	if inpututil.IsKeyJustReleased(ebiten.Key2) {
		mixColor = !mixColor
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyQ) {
		threshold -= 0.02
		fmt.Println("threshold: ", threshold)
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyW) {
		threshold += 0.02
		fmt.Println("threshold: ", threshold)
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyE) {
		radius -= 1
		fmt.Println("radius: ", radius)
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyR) {
		radius += 1
		fmt.Println("radius: ", radius)
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyD) {
		blurRadius -= 1
		fmt.Println("blurRadius: ", blurRadius)
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyF) {
		blurRadius += 1
		fmt.Println("blurRadius: ", blurRadius)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	if curShader != nil {
		op := &ebiten.DrawRectShaderOptions{}
		x, y := ebiten.CursorPosition()

		mix := 0
		if mixColor {
			mix = 1
		}

		if enabledShader {
			op.Uniforms = map[string]any{
				// "Seed":   s.seed,
				// "Invert": inv,
				"Value":          []float32{float32(x), float32(y)},
				"MixColor":       mix,
				"ColorThreshold": threshold,
				"Radius":         radius,
				"BlurRadius":     blurRadius,
			}
			op.Images[0] = runnerImage
			w, h := screen.Bounds().Dx(), screen.Bounds().Dy()
			screen.DrawRectShader(w, h, curShader, op)
		} else {
			screen.DrawImage(runnerImage, nil)
		}

		ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))

	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

var basePath = filepath.Join("pkg", "shader")

func main() {

	go func() {
		mod := time.Now()
		for {
			shaderFile := filepath.Join(basePath, "kage", "test.kage")
			st, err := os.Stat(shaderFile)
			if err != nil {
				panic(err)
			}
			if st.ModTime() != mod {
				mod = st.ModTime()
				f, err := os.ReadFile(shaderFile)
				if err != nil {
					panic(err)
				}
				curShader, err = ebiten.NewShader(f)
				if err != nil {
					fmt.Println("shader error: " + err.Error())
				} else {
					fmt.Println("shader reloaded")
				}
			}
			time.Sleep(time.Second)
		}
	}()

	f, err := os.ReadFile(filepath.Join(basePath, "test", "shader-test.png"))
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(bytes.NewReader(f))
	if err != nil {
		panic(err)
	}
	runnerImage = ebiten.NewImageFromImage(img)

	factor := 3
	ebiten.SetWindowSize(screenWidth*factor, screenHeight*factor)
	ebiten.SetWindowTitle("Animation (Ebitengine Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
