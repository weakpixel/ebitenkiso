//go:build js && wasm
// +build js,wasm

package logger

import (
	"fmt"
	"syscall/js"
)

var (
	info  js.Value
	err   js.Value
	debug js.Value
)

func init() {
	console := js.Global().Get("console")
	info = console.Get("log")
	if !info.Truthy() || info.Type() != js.TypeFunction {
		panic("console.log is not a function")
	}
	err = console.Get("error")
	if !err.Truthy() || err.Type() != js.TypeFunction {
		panic("console.error is not a function")
	}

	debug = console.Get("debug")
	if !debug.Truthy() || debug.Type() != js.TypeFunction {
		panic("console.debug is not a function")
	}
}

func Infof(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	info.Invoke("[WASM][INFO] " + msg)

}
func Errorf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	err.Invoke("[WASM][ERROR] " + msg)
}

func DebugF(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	debug.Invoke("[WASM][DEBUG] " + msg)
}
