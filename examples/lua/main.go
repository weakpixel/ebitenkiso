package main

import (
	"image/color"

	"github.com/weakpixel/ebitenkiso/examples/assets"
	vm "github.com/weakpixel/ebitenkiso/pkg/vm/lua"

	"github.com/Shopify/go-lua"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	data map[string]any
	env  *vm.Env
	X, Y float64
	draw DrawContext
}

type DrawContext struct {
	Screen *ebiten.Image
	color  color.RGBA
}

func (c *DrawContext) Register(env *vm.Env) {
	env.RegisterFn("drawFilledCircle", c.DrawFilledCircle)
}

func (c *DrawContext) DrawFilledCircle(l *lua.State) int {
	if c.Screen == nil {
		return 0
	}
	x, _ := l.ToNumber(1)
	y, _ := l.ToNumber(2)
	radius, _ := l.ToNumber(3)
	vector.FillCircle(c.Screen, float32(x), float32(y), float32(radius), c.color, false)
	return 0
}

func (p *Player) Update() error {
	return p.env.Call("update")
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.draw.Screen = screen
	p.env.Call("draw")
	p.draw.Screen = nil
}

type Game struct {
	vm      *vm.ScriptVM
	player1 *Player
	player2 *Player
}

func (g *Game) Update() error {
	must(g.player1.Update())
	must(g.player2.Update())
	return nil
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player1.Draw(screen)
	g.player2.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func newPlayer(vm *vm.ScriptVM, name string, script string) *Player {

	env := vm.NewEnv()

	player := &Player{
		data: map[string]any{
			"name": name,
		},
		env: env,
		X:   320,
		Y:   240,
		draw: DrawContext{
			color: color.RGBA{0, 0xff, 0, 0xff},
		},
	}

	player.draw.Register(env)

	env.RegisterFn("move", func(l *lua.State) int {
		x, _ := l.ToNumber(1)
		y, _ := l.ToNumber(2)
		player.X += float64(x)
		player.Y += float64(y)
		return 0
	})

	env.RegisterGetterSetterNumber("x", &player.X)
	env.RegisterGetterSetterNumber("y", &player.Y)
	env.RegisterMap("data", player.data)

	raw, err := assets.FS.ReadFile(script)
	if err != nil {
		panic(err)
	}
	err = env.LoadScript(script, string(raw))
	if err != nil {
		panic(err)
	}

	return player
}
func main() {
	vm := vm.NewVM()
	player1 := newPlayer(vm, "player1", "player-mouse.lua")
	player1.env.SetInt("PlayerLeft", int(ebiten.KeyArrowLeft))
	player1.env.SetInt("PlayerRight", int(ebiten.KeyArrowRight))
	player1.env.SetInt("PlayerUp", int(ebiten.KeyArrowUp))
	player1.env.SetInt("PlayerDown", int(ebiten.KeyArrowDown))

	player2 := newPlayer(vm, "player2", "player-keyboard.lua")
	player2.env.SetInt("PlayerLeft", int(ebiten.KeyA))
	player2.env.SetInt("PlayerRight", int(ebiten.KeyD))
	player2.env.SetInt("PlayerUp", int(ebiten.KeyW))
	player2.env.SetInt("PlayerDown", int(ebiten.KeyS))

	game := Game{
		vm:      vm,
		player1: player1,
		player2: player2,
	}

	ebiten.SetWindowSize(640, 480)
	err := ebiten.RunGame(&game)
	if err != nil {
		panic(err)
	}
}
