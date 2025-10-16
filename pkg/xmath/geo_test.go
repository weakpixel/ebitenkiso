package xmath

import (
	"math"
	"testing"
)

func TestDegToRad(t *testing.T) {
	if math.Abs(DegToRad(180)-math.Pi) > 1e-9 {
		t.Errorf("DegToRad: expected %v, got %v", math.Pi, DegToRad(180))
	}

	for _, given := range []float64{0, 30, 45, 60, 90, 120, 180, 270, 360, 22, 36, 162, 199} {
		rad := DegToRad(given)
		result := RadToDeg(rad)

		if given != math.Round(result) {
			t.Errorf("DegToRad/RadToDeg: expected %v, got %v", given, result)
		}
	}

}

func TestRadToDeg(t *testing.T) {
	if math.Abs(RadToDeg(math.Pi)-180) > 1e-9 {
		t.Errorf("RadToDeg: expected 180, got %v", RadToDeg(math.Pi))
	}
}

func TestSinCos(t *testing.T) {
	s, c := SinCos(math.Pi / 2)
	if math.Abs(s-1) > 1e-9 || math.Abs(c) > 1e-9 {
		t.Errorf("SinCos: expected (1,0), got (%v,%v)", s, c)
	}
}

func TestCircularMovement(t *testing.T) {
	x, y := CircularMovement(0, 0, 1, math.Pi/2)
	if math.Abs(x) > 1e-9 || math.Abs(y-1) > 1e-9 {
		t.Errorf("CircularMovement: expected (0,1), got (%v,%v)", x, y)
	}
}

func TestEuclideanDistance(t *testing.T) {
	// Test with same points (distance should be 0)
	d0 := EuclideanDistance(0, 0, 0, 0)
	if math.Abs(d0) > 1e-9 {
		t.Errorf("EuclideanDistance: expected 0, got %v", d0)
	}

	// Test with negative coordinates
	d1 := EuclideanDistance(-1, -2, -4, -6)
	expected1 := 5.0
	if math.Abs(d1-expected1) > 1e-9 {
		t.Errorf("EuclideanDistance: expected %v, got %v", expected1, d1)
	}

	// Test with swapped points (should be same as above)
	d2 := EuclideanDistance(-4, -6, -1, -2)
	if math.Abs(d2-expected1) > 1e-9 {
		t.Errorf("EuclideanDistance: expected %v, got %v", expected1, d2)
	}

	// Test with floating point coordinates
	d3 := EuclideanDistance(1.5, 2.5, 4.5, 6.5)
	expected3 := 5.0
	if math.Abs(d3-expected3) > 1e-9 {
		t.Errorf("EuclideanDistance: expected %v, got %v", expected3, d3)
	}

}

func TestSquaredDistance(t *testing.T) {
	// Test with swapped coordinates
	d := SquaredDistance(3, 4, 0, 0)

	if math.Abs(d-25) > 1e-9 {
		t.Errorf("SquaredDistance: expected 25, got %v", d)
	}

	// Test with negative coordinates
	d2 := SquaredDistance(-1, -2, -4, -6)
	if math.Abs(d2-25) > 1e-9 {
		t.Errorf("SquaredDistance: expected 25, got %v", d2)
	}

	// Test with floating point coordinates
	d3 := SquaredDistance(1.5, 2.5, 4.5, 6.5)
	if math.Abs(d3-25) > 1e-9 {
		t.Errorf("SquaredDistance: expected 25, got %v", d3)
	}

	// Test with same points (should be zero)
	d4 := SquaredDistance(2.2, 3.3, 2.2, 3.3)
	if math.Abs(d4) > 1e-9 {
		t.Errorf("SquaredDistance: expected 0, got %v", d4)
	}

	// Test with large values
	d5 := SquaredDistance(1000, 2000, 1003, 2004)
	if math.Abs(d5-25) > 1e-9 {
		t.Errorf("SquaredDistance: expected 25, got %v", d5)
	}

}

func TestCircleVsCircle(t *testing.T) {
	c1 := Circle{0, 0, 1}
	c2 := Circle{1.5, 0, 1}
	if !CircleIntersect(c1, c2) {
		t.Error("CircleIntersect: expected true, got false")
	}
	c3 := Circle{3, 0, 1}
	if CircleIntersect(c1, c3) {
		t.Error("CircleIntersect: expected false, got true")
	}
}

func TestRectIntersect(t *testing.T) {
	r1 := Rect{0, 0, 2, 2}
	r2 := Rect{1, 1, 2, 2}
	if !RectIntersect(r1, r2) {
		t.Error("RectIntersect: expected true, got false")
	}
	r3 := Rect{3, 3, 1, 1}
	if RectIntersect(r1, r3) {
		t.Error("RectIntersect: expected false, got true")
	}
}

func TestOBBvsOBB(t *testing.T) {
	// Square at origin
	poly := Polygon{
		Points: []Vector2{
			{-1, -1},
			{1, -1},
			{1, 1},
			{-1, 1},
		},
	}
	pos1 := Vector2{0, 0}
	pos2 := Vector2{1, 0}
	rot1 := 0.0
	rot2 := 0.0
	if !OBBvsOBB(poly, poly, pos1, pos2, rot1, rot2) {
		t.Error("OBBvsOBB: expected true, got false")
	}
	// Move far apart
	pos3 := Vector2{5, 0}
	if OBBvsOBB(poly, poly, pos1, pos3, rot1, rot2) {
		t.Error("OBBvsOBB: expected false, got true")
	}
	// Rotated overlap
	pos4 := Vector2{0.5, 0}
	rot2 = math.Pi / 4
	if !OBBvsOBB(poly, poly, pos1, pos4, rot1, rot2) {
		t.Error("OBBvsOBB: expected true, got false (rotated overlap)")
	}
}

func TestTransformPolygon(t *testing.T) {
	poly := Polygon{
		Points: []Vector2{
			{1, 0},
		},
	}
	pos := Vector2{0, 0}
	rot := math.Pi / 2
	trans := transformPolygon(poly, pos, rot)
	if math.Abs(trans[0].X) > 1e-9 || math.Abs(trans[0].Y-1) > 1e-9 {
		t.Errorf("transformPolygon: expected (0,1), got (%v,%v)", trans[0].X, trans[0].Y)
	}
}

func TestHasSeparatingAxis(t *testing.T) {
	// Two squares, one at (0,0), one at (3,0)
	sq1 := []Vector2{{-1, -1}, {1, -1}, {1, 1}, {-1, 1}}
	sq2 := []Vector2{{2, -1}, {4, -1}, {4, 1}, {2, 1}}
	if !hasSeparatingAxis(sq1, sq2) {
		t.Error("hasSeparatingAxis: expected true (separated), got false")
	}
	if hasSeparatingAxis(sq1, sq1) {
		t.Error("hasSeparatingAxis: expected false (overlapping), got true")
	}
}

func TestProjectPolygon(t *testing.T) {
	axis := Vector2{1, 0}
	poly := []Vector2{{1, 2}, {3, 4}, {5, 6}}
	min, max := projectPolygon(axis, poly)
	if math.Abs(min-1) > 1e-9 || math.Abs(max-5) > 1e-9 {
		t.Errorf("projectPolygon: expected min=1, max=5, got min=%v, max=%v", min, max)
	}
}
