package lua

import (
	lua "github.com/Shopify/go-lua"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func registerEbiten(vm *ScriptVM) {
	vm.SetInt("KeyArrowUp", int(ebiten.KeyArrowUp))
	vm.SetInt("KeyArrowDown", int(ebiten.KeyArrowDown))
	vm.SetInt("KeyArrowLeft", int(ebiten.KeyArrowLeft))
	vm.SetInt("KeyArrowRight", int(ebiten.KeyArrowRight))
	registerKeyFns(vm.state, map[string]keyFuncs{
		"isKeyPressed":      ebiten.IsKeyPressed,
		"isKeyJustPressed":  inpututil.IsKeyJustPressed,
		"isKeyJustReleased": inpututil.IsKeyJustReleased,
	})

	vm.SetInt("MouseButton0", int(ebiten.MouseButton0))
	vm.SetInt("MouseButton1", int(ebiten.MouseButton1))
	vm.SetInt("MouseButton2", int(ebiten.MouseButton2))
	vm.SetInt("MouseButton3", int(ebiten.MouseButton3))
	vm.SetInt("MouseButton4", int(ebiten.MouseButton4))
	vm.SetInt("MouseButtonLeft", int(ebiten.MouseButtonLeft))
	vm.SetInt("MouseButtonMiddle", int(ebiten.MouseButtonMiddle))
	vm.SetInt("MouseButtonRight", int(ebiten.MouseButtonRight))

	registerMouseFns(vm.state, map[string]mouseFuncs{
		"isMouseButtonPressed":      ebiten.IsMouseButtonPressed,
		"isMouseButtonJustPressed":  inpututil.IsMouseButtonJustPressed,
		"isMouseButtonJustReleased": inpututil.IsMouseButtonJustReleased,
	})

	vm.state.Register("cursorPosition", func(l *lua.State) int {
		x, y := ebiten.CursorPosition()
		l.PushInteger(x)
		l.PushInteger(y)
		return 2
	})
}

type keyFuncs func(key ebiten.Key) bool

func registerKeyFns(l *lua.State, fns map[string]keyFuncs) {
	for name, fn := range fns {
		l.Register(name, func(l *lua.State) int {
			pressed := false
			if v, ok := l.ToNumber(1); ok {
				k := ebiten.Key(v)
				pressed = fn(k)
			}
			l.PushBoolean(pressed)
			return 1
		})
	}
}

type mouseFuncs func(key ebiten.MouseButton) bool

func registerMouseFns(l *lua.State, fns map[string]mouseFuncs) {
	for name, fn := range fns {
		l.Register(name, func(l *lua.State) int {
			pressed := false
			if v, ok := l.ToNumber(1); ok {
				k := ebiten.MouseButton(v)
				pressed = fn(k)
			}
			l.PushBoolean(pressed)
			return 1
		})
	}
}
