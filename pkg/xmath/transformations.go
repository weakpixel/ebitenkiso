package xmath

import (
	"math"
)

// ========================
// TRANSLATION
// ========================

// Translate moves a point by an offset.
// Example: newPos := Translate(oldPos, Vector2{1, 0})
func Translate(point, offset Vector2) Vector2 {
	return point.Add(offset)
}

// ========================
// SCALING
// ========================

// Scale scales a point relative to the origin (0,0).
// Example: bigger := Scale(point, 2, 2)
func Scale(point Vector2, scaleX, scaleY float64) Vector2 {
	return Vector2{point.X * scaleX, point.Y * scaleY}
}

// ScaleAround scales a point around a specific pivot.
// Example: bigger := ScaleAround(point, pivot, 2, 2)
func ScaleAround(point, pivot Vector2, scaleX, scaleY float64) Vector2 {
	// Move to pivot space, scale, then move back
	relative := point.Sub(pivot)
	scaled := Vector2{relative.X * scaleX, relative.Y * scaleY}
	return scaled.Add(pivot)
}

// ========================
// ROTATION
// ========================

// Rotate rotates a point around the origin by radians.
// Example: rotated := Rotate(point, math.Pi/4)
func Rotate(point Vector2, radians float64) Vector2 {
	c := math.Cos(radians)
	s := math.Sin(radians)
	return Vector2{
		point.X*c - point.Y*s,
		point.X*s + point.Y*c,
	}
}

// RotateAround rotates a point around a pivot by radians.
// Example: rotated := RotateAround(point, pivot, math.Pi/2)
func RotateAround(point, pivot Vector2, radians float64) Vector2 {
	relative := point.Sub(pivot)
	rotated := Rotate(relative, radians)
	return rotated.Add(pivot)
}

// ========================
// MATRIX MATH
// ========================

// Mat3 represents a 3x3 transformation matrix.
type Mat3 [3][3]float64

// Identity3 returns the identity matrix.
func Identity3() Mat3 {
	return Mat3{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
}

// TranslationMatrix creates a translation matrix.
func TranslationMatrix(tx, ty float64) Mat3 {
	return Mat3{
		{1, 0, tx},
		{0, 1, ty},
		{0, 0, 1},
	}
}

// ScalingMatrix creates a scaling matrix.
func ScalingMatrix(sx, sy float64) Mat3 {
	return Mat3{
		{sx, 0, 0},
		{0, sy, 0},
		{0, 0, 1},
	}
}

// RotationMatrix creates a rotation matrix (radians).
func RotationMatrix(radians float64) Mat3 {
	c := math.Cos(radians)
	s := math.Sin(radians)
	return Mat3{
		{c, -s, 0},
		{s, c, 0},
		{0, 0, 1},
	}
}

// MulMat3 multiplies two 3x3 matrices.
func MulMat3(a, b Mat3) Mat3 {
	var result Mat3
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			sum := 0.0
			for k := 0; k < 3; k++ {
				sum += a[row][k] * b[k][col]
			}
			result[row][col] = sum
		}
	}
	return result
}

// TransformPoint applies a 3x3 transformation matrix to a Vector2.
func TransformPoint(m Mat3, v Vector2) Vector2 {
	x := m[0][0]*v.X + m[0][1]*v.Y + m[0][2]
	y := m[1][0]*v.X + m[1][1]*v.Y + m[1][2]
	return Vector2{x, y}
}
