//go:build !js && !wasm

package logger

import "fmt"

func Infof(format string, v ...interface{}) {
	fmt.Println("[INFO] " + fmt.Sprintf(format, v...))
}

func Errorf(format string, v ...interface{}) {
	fmt.Println("[ERROR] " + fmt.Sprintf(format, v...))
}

func DebugF(format string, v ...interface{}) {
	fmt.Println("[DEBUG] " + fmt.Sprintf(format, v...))
}
