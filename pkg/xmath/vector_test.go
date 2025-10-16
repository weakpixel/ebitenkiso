package xmath

import (
	"math"
	"testing"
)

func TestVector2_Add(t *testing.T) {
	a := Vector2{1, 2}
	b := Vector2{3, 4}
	res := a.Add(b)
	if res != (Vector2{4, 6}) {
		t.Errorf("Add: expected {4,6}, got {%v,%v}", res.X, res.Y)
	}
}

func TestVector2_Sub(t *testing.T) {
	a := Vector2{5, 7}
	b := Vector2{2, 3}
	res := a.Sub(b)
	if res != (Vector2{3, 4}) {
		t.Errorf("Sub: expected {3,4}, got {%v,%v}", res.X, res.Y)
	}
}

func TestVector2_Mul(t *testing.T) {
	v := Vector2{2, 3}
	res := v.Mul(2)
	if res != (Vector2{4, 6}) {
		t.Errorf("Mul: expected {4,6}, got {%v,%v}", res.X, res.Y)
	}
}

func TestVector2_Div(t *testing.T) {
	v := Vector2{4, 6}
	res := v.Div(2)
	if res != (Vector2{2, 3}) {
		t.Errorf("Div: expected {2,3}, got {%v,%v}", res.X, res.Y)
	}
}

func TestVector2_Length(t *testing.T) {
	v := Vector2{3, 4}
	if v.Length() != 5 {
		t.Errorf("Length: expected 5, got %v", v.Length())
	}
}

func TestVector2_LengthSquared(t *testing.T) {
	v := Vector2{3, 4}
	if v.LengthSquared() != 25 {
		t.Errorf("LengthSquared: expected 25, got %v", v.LengthSquared())
	}
}

func TestVector2_Normalize(t *testing.T) {
	v := Vector2{3, 4}
	n := v.Normalize()
	if math.Abs(n.Length()-1) > 1e-9 {
		t.Errorf("Normalize: expected length 1, got %v", n.Length())
	}
}

func TestVector2_Dot(t *testing.T) {
	a := Vector2{1, 2}
	b := Vector2{3, 4}
	if a.Dot(b) != 11 {
		t.Errorf("Dot: expected 11, got %v", a.Dot(b))
	}
}

func TestVector2_Distance(t *testing.T) {
	a := Vector2{0, 0}
	b := Vector2{3, 4}
	if a.Distance(b) != 5 {
		t.Errorf("Distance: expected 5, got %v", a.Distance(b))
	}
}

func TestVector2_DistanceSquared(t *testing.T) {
	a := Vector2{0, 0}
	b := Vector2{3, 4}
	if a.DistanceSquared(b) != 25 {
		t.Errorf("DistanceSquared: expected 25, got %v", a.DistanceSquared(b))
	}
}

func TestLerp(t *testing.T) {
	a := Vector2{0, 0}
	b := Vector2{10, 10}
	res := Lerp(a, b, 0.5)
	if res != (Vector2{5, 5}) {
		t.Errorf("Lerp: expected {5,5}, got {%v,%v}", res.X, res.Y)
	}
}

func TestVector2_Rotate(t *testing.T) {
	v := Vector2{1, 0}
	rot := v.Rotate(math.Pi / 2)
	if math.Abs(rot.X) > 1e-9 || math.Abs(rot.Y-1) > 1e-9 {
		t.Errorf("Rotate: expected {0,1}, got {%v,%v}", rot.X, rot.Y)
	}
}

func TestVector2_Angle(t *testing.T) {
	v := Vector2{0, 1}
	if math.Abs(v.Angle()-math.Pi/2) > 1e-9 {
		t.Errorf("Angle: expected %v, got %v", math.Pi/2, v.Angle())
	}
}

func TestFromAngle(t *testing.T) {
	v := FromAngle(math.Pi / 2)
	if math.Abs(v.X) > 1e-9 || math.Abs(v.Y-1) > 1e-9 {
		t.Errorf("FromAngle: expected {0,1}, got {%v,%v}", v.X, v.Y)
	}
}

func TestVector2_Reflect(t *testing.T) {
	v := Vector2{1, -1}
	n := Vector2{0, 1} // y-axis normal
	ref := v.Reflect(n)
	if math.Abs(ref.X-1) > 1e-9 || math.Abs(ref.Y-1) > 1e-9 {
		t.Errorf("Reflect: expected {1,1}, got {%v,%v}", ref.X, ref.Y)
	}
}
