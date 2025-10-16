package xmath

import "math"

// AngleBetween returns the angle (in radians) between two vectors.
// Example: angle := v1.AngleBetween(v2)
func (v Vector2) AngleBetween(o Vector2) float64 {
	v1Norm := v.Normalize()
	v2Norm := o.Normalize()
	dot := v1Norm.Dot(v2Norm)
	// Clamp dot to avoid floating-point errors outside [-1, 1]
	if dot > 1 {
		dot = 1
	} else if dot < -1 {
		dot = -1
	}
	return math.Acos(dot) // result in radians
}

// IsFacingSameDirection returns true if the angle between the vectors is small.
// Useful for checking if two entities are aligned in direction.
func (v Vector2) IsFacingSameDirection(o Vector2, toleranceDegrees float64) bool {
	angle := v.AngleBetween(o) * (180.0 / math.Pi) // convert to degrees
	return angle <= toleranceDegrees
}

// IsInFront checks if the target position is in front of the entity's facing direction.
// facingDir should be normalized.
func IsInFront(entityPos, targetPos, facingDir Vector2) bool {
	toTarget := targetPos.Sub(entityPos).Normalize()
	return facingDir.Dot(toTarget) > 0
}

// IsTargetInVisionCone checks if target is inside enemy's vision cone.
// facingDir should be normalized.
// coneAngleDegrees is the full width of the vision cone (e.g., 90 means 45° to each side).
func IsTargetInVisionCone(enemyPos, facingDir, targetPos Vector2, coneAngleDegrees float64, maxDistance float64) bool {
	toTarget := targetPos.Sub(enemyPos)
	distance := toTarget.Length()
	if distance > maxDistance {
		return false // Out of range
	}

	toTargetNorm := toTarget.Normalize()
	dot := facingDir.Dot(toTargetNorm)

	// Convert cone half-angle to radians
	halfAngleRad := (coneAngleDegrees / 2) * (math.Pi / 180)

	// Threshold for dot product
	// cos(half angle) is the minimum dot value for the target to be inside the cone
	threshold := math.Cos(halfAngleRad)

	return dot >= threshold
}
