package xmath

import (
	"math"
	"testing"
)

func TestTranslate(t *testing.T) {
	p := Vector2{2, 3}
	offset := Vector2{1, -1}
	res := Translate(p, offset)
	expected := Vector2{3, 2}
	if res != expected {
		t.Errorf("Translate: expected %v, got %v", expected, res)
	}
}

func TestScale(t *testing.T) {
	p := Vector2{2, 3}
	res := Scale(p, 2, 3)
	expected := Vector2{4, 9}
	if res != expected {
		t.Errorf("Scale: expected %v, got %v", expected, res)
	}
}

func TestScaleAround(t *testing.T) {
	p := Vector2{3, 4}
	pivot := Vector2{1, 1}
	res := ScaleAround(p, pivot, 2, 3)
	expected := Vector2{5, 10}
	if res != expected {
		t.Errorf("ScaleAround: expected %v, got %v", expected, res)
	}
}

func TestRotate(t *testing.T) {
	p := Vector2{1, 0}
	rot := Rotate(p, math.Pi/2)
	if math.Abs(rot.X) > 1e-9 || math.Abs(rot.Y-1) > 1e-9 {
		t.Errorf("Rotate: expected {0,1}, got {%v,%v}", rot.X, rot.Y)
	}
}

func TestRotateAround(t *testing.T) {
	p := Vector2{2, 1}
	pivot := Vector2{1, 1}
	rot := RotateAround(p, pivot, math.Pi)
	expected := Vector2{0, 1}
	if math.Abs(rot.X-expected.X) > 1e-9 || math.Abs(rot.Y-expected.Y) > 1e-9 {
		t.Errorf("RotateAround: expected %v, got %v", expected, rot)
	}
}

func TestIdentity3(t *testing.T) {
	id := Identity3()
	expected := Mat3{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	if id != expected {
		t.Errorf("Identity3: expected %v, got %v", expected, id)
	}
}

func TestTranslationMatrix(t *testing.T) {
	m := TranslationMatrix(2, 3)
	expected := Mat3{
		{1, 0, 2},
		{0, 1, 3},
		{0, 0, 1},
	}
	if m != expected {
		t.Errorf("TranslationMatrix: expected %v, got %v", expected, m)
	}
}

func TestScalingMatrix(t *testing.T) {
	m := ScalingMatrix(2, 3)
	expected := Mat3{
		{2, 0, 0},
		{0, 3, 0},
		{0, 0, 1},
	}
	if m != expected {
		t.Errorf("ScalingMatrix: expected %v, got %v", expected, m)
	}
}

func TestRotationMatrix(t *testing.T) {
	angle := math.Pi / 2
	m := RotationMatrix(angle)
	c := math.Cos(angle)
	s := math.Sin(angle)
	expected := Mat3{
		{c, -s, 0},
		{s, c, 0},
		{0, 0, 1},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if math.Abs(m[i][j]-expected[i][j]) > 1e-9 {
				t.Errorf("RotationMatrix: expected %v, got %v", expected, m)
				return
			}
		}
	}
}

func TestMulMat3(t *testing.T) {
	a := Mat3{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	b := Mat3{
		{9, 8, 7},
		{6, 5, 4},
		{3, 2, 1},
	}
	res := MulMat3(a, b)
	expected := Mat3{
		{30, 24, 18},
		{84, 69, 54},
		{138, 114, 90},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if math.Abs(res[i][j]-expected[i][j]) > 1e-9 {
				t.Errorf("MulMat3: expected %v, got %v", expected, res)
				return
			}
		}
	}
}

func TestTransformPoint(t *testing.T) {
	m := TranslationMatrix(2, 3)
	v := Vector2{1, 1}
	res := TransformPoint(m, v)
	expected := Vector2{3, 4}
	if math.Abs(res.X-expected.X) > 1e-9 || math.Abs(res.Y-expected.Y) > 1e-9 {
		t.Errorf("TransformPoint: expected %v, got %v", expected, res)
	}
}
