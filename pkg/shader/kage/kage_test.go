package kage

import "testing"

func TestInit(t *testing.T) {
	// nothing to do
	if len(shaders) == 0 {
		t.Error("no shaders loaded")
	}
	for k, _ := range shaders {
		t.Log(k)
	}
}
