package xmath

import (
	"math"
	"testing"
)

func TestIsTargetInVisionCone(t *testing.T) {
	enemyPos := Vector2{0, 0}
	facingDir := Vector2{1, 0}.Normalize()
	maxDistance := 10.0

	// Target directly in front, within cone and distance
	targetPos := Vector2{5, 0}
	if !IsTargetInVisionCone(enemyPos, facingDir, targetPos, 90, maxDistance) {
		t.Errorf("IsTargetInVisionCone: expected true, got false (directly in front)")
	}

	// Target at edge of cone
	angle := 45.0 // half of 90 degree cone
	x := math.Cos(angle * math.Pi / 180)
	y := math.Sin(angle * math.Pi / 180)
	targetEdge := Vector2{5 * x, 5 * y}
	if !IsTargetInVisionCone(enemyPos, facingDir, targetEdge, 90, maxDistance) {
		t.Errorf("IsTargetInVisionCone: expected true, got false (at cone edge)")
	}

	// Target outside cone
	targetOutside := Vector2{0, 5}
	if IsTargetInVisionCone(enemyPos, facingDir, targetOutside, 60, maxDistance) {
		t.Errorf("IsTargetInVisionCone: expected false, got true (outside cone)")
	}

	// Target out of range
	targetFar := Vector2{20, 0}
	if IsTargetInVisionCone(enemyPos, facingDir, targetFar, 90, maxDistance) {
		t.Errorf("IsTargetInVisionCone: expected false, got true (out of range)")
	}

	// Target behind enemy
	targetBehind := Vector2{-5, 0}
	if IsTargetInVisionCone(enemyPos, facingDir, targetBehind, 120, maxDistance) {
		t.Errorf("IsTargetInVisionCone: expected false, got true (behind enemy)")
	}
}
func TestAngleBetween(t *testing.T) {
	v1 := Vector2{1, 0}
	v2 := Vector2{0, 1}
	angle := v1.AngleBetween(v2)
	if math.Abs(angle-math.Pi/2) > 1e-9 {
		t.Errorf("AngleBetween: expected %v, got %v", math.Pi/2, angle)
	}

	v3 := Vector2{1, 0}
	v4 := Vector2{1, 0}
	angle2 := v3.AngleBetween(v4)
	if math.Abs(angle2) > 1e-9 {
		t.Errorf("AngleBetween: expected 0, got %v", angle2)
	}
}

func TestIsFacingSameDirection(t *testing.T) {
	v1 := Vector2{1, 0}
	v2 := Vector2{1, 0.01}
	if !v1.IsFacingSameDirection(v2, 1) {
		t.Errorf("IsFacingSameDirection: expected true, got false")
	}

	v3 := Vector2{1, 0}
	v4 := Vector2{-1, 0}
	if v3.IsFacingSameDirection(v4, 10) {
		t.Errorf("IsFacingSameDirection: expected false, got true")
	}
}

func TestIsInFront(t *testing.T) {
	entityPos := Vector2{0, 0}
	targetPos := Vector2{1, 0}
	facingDir := Vector2{1, 0}
	if !IsInFront(entityPos, targetPos, facingDir) {
		t.Errorf("IsInFront: expected true, got false")
	}

	targetPos2 := Vector2{-1, 0}
	if IsInFront(entityPos, targetPos2, facingDir) {
		t.Errorf("IsInFront: expected false, got true")
	}
}
