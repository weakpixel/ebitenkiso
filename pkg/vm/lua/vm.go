package lua

import (
	"fmt"
	"sort"
	"strings"

	lua "github.com/Shopify/go-lua"
)

func NewVM() *ScriptVM {
	state := lua.NewState()
	lua.BaseOpen(state)
	lua.TableOpen(state)
	lua.StringOpen(state)
	lua.MathOpen(state)

	vm := &ScriptVM{state: state}

	registerEbiten(vm)

	vm.state.Register("inspectTable", func(l *lua.State) int {
		if l.IsTable(1) {
			l.PushNil()
			buf := strings.Builder{}
			buf.WriteString("{ ")
			list := []string{}
			for l.Next(-2) {
				// key at -2, value at -1
				key, _ := l.ToString(-2)
				val := l.ToValue(-1)

				list = append(list, fmt.Sprintf("%s=%v", key, val))
				// buf.WriteString(fmt.Sprintf("%s = %v, ", key, val))
				l.Pop(1) // pop value, keep key for next iteration
			}

			sort.Strings(list)
			fmt.Printf("[%s]\n", strings.Join(list, ", "))

		}
		return 0
	})

	return vm
}

type ScriptVM struct {
	state *lua.State
}

func (vm *ScriptVM) SetInt(name string, value int) {
	vm.state.PushInteger(value)
	vm.state.SetGlobal(name)
}

func (vm *ScriptVM) NewEnv() *Env {
	vm.state.NewTable()              // Create new environment table
	vm.state.NewTable()              // metatable
	vm.state.Global("_G")            // Push _G
	vm.state.SetField(-2, "__index") // metatable.__index = _G
	vm.state.SetMetaTable(-2)        // setmetatable(env, metatable)
	return &Env{
		idx:   vm.state.Top(),
		state: vm.state,
	}
}

type Env struct {
	idx   int
	state *lua.State
}

func (e *Env) RegisterGetterSetterNumber(name string, value *float64) {
	e.state.PushValue(e.idx)
	defer e.state.Pop(1)
	e.RegisterFn("set_"+name, func(l *lua.State) int {
		if v, ok := l.ToNumber(1); ok {
			*value = float64(v)
		}
		return 0
	})
	e.RegisterFn("get_"+name, func(l *lua.State) int {
		l.PushNumber(*value)
		return 1
	})
}

func (e *Env) RegisterMap(name string, data map[string]any) {
	e.state.PushValue(e.idx)

	// name of the map
	e.state.PushString(name)
	e.state.NewTable() // proxy table
	e.state.NewTable() // metatable with __index and __newindex calling Go functions
	e.state.PushString("__index")

	// GET
	e.state.PushGoFunction(func(l *lua.State) int {
		key, ok := e.state.ToString(2)
		if !ok {
			e.state.PushNil()
			return 1
		}
		val := data[key]
		switch v := val.(type) {
		case nil:
			l.PushNil()
		case bool:
			l.PushBoolean(v)
		case int:
			l.PushInteger(v)
		case int64:
			l.PushInteger(int(v))
		case float64:
			l.PushNumber(v)
		case string:
			l.PushString(v)
		// you can add more cases: map[string]any -> table, etc.
		default:
			l.PushString(fmt.Sprintf("%v", v)) // fallback
		}
		return 1
	})
	e.state.SetTable(-3)

	// __newindex
	e.state.PushString("__newindex")

	// SET
	e.state.PushGoFunction(func(l *lua.State) int {
		key, ok := e.state.ToString(2)
		if !ok {
			return 0
		}
		switch e.state.TypeOf(3) {
		case lua.TypeNil:
			delete(data, key)
		case lua.TypeBoolean:
			data[key] = l.ToBoolean(3)
		case lua.TypeNumber:
			data[key], _ = l.ToNumber(3) // float64
		case lua.TypeString:
			data[key], _ = l.ToString(3)
		default:
			data[key], _ = l.ToString(3)
		}
		return 0
	})
	e.state.SetTable(-3)
	// set metatable
	e.state.SetMetaTable(-2)

	e.state.SetTable(-3)

	e.state.Pop(1)

}
func (e *Env) RegisterFn(name string, fn lua.Function) {
	e.state.PushValue(e.idx)
	e.state.PushString(name)
	e.state.PushGoFunction(fn)
	e.state.SetTable(-3)
	e.state.Pop(1)
}

func (e *Env) SetInt(name string, value int) {
	e.state.PushValue(e.idx)
	e.state.PushString(name)
	e.state.PushInteger(value)
	e.state.SetTable(-3)
	e.state.Pop(1)
}

func (e *Env) LoadScript(name, script string) error {
	err := lua.LoadBuffer(e.state, script, name, "text")
	if err != nil {
		return err
	}
	e.state.PushValue(e.idx)
	lua.SetUpValue(e.state, -2, 1)
	return e.state.ProtectedCall(0, 0, 0)
}

func (e *Env) Call(name string) error {
	e.state.PushValue(e.idx)
	defer e.state.Pop(1)
	e.state.Field(-1, name)
	if !e.state.IsFunction(-1) {
		return fmt.Errorf("%s function not found", name)
	}
	e.state.PushValue(e.idx)
	lua.SetUpValue(e.state, -2, 1)
	return e.state.ProtectedCall(0, 0, 0)
}
