package xmath

import (
	"math"
)

// Vector2 represents a 2D vector with X and Y components.
type Vector2 struct {
	X, Y float64
}

// Add returns the sum of two vectors.
// Example: playerPos = playerPos.Add(playerVelocity)
func (v Vector2) Add(o Vector2) Vector2 {
	return Vector2{v.X + o.X, v.Y + o.Y}
}

// Sub returns the difference of two vectors.
// Example: toTarget := enemyPos.Sub(playerPos)
func (v Vector2) Sub(o Vector2) Vector2 {
	return Vector2{v.X - o.X, v.Y - o.Y}
}

// Mul returns the vector scaled by a scalar.
// Example: fastBulletVelocity := bulletDirection.Mul(500)
func (v Vector2) Mul(scalar float64) Vector2 {
	return Vector2{v.X * scalar, v.Y * scalar}
}

// Div returns the vector divided by a scalar.
// Example: slowVelocity := velocity.Div(2)
func (v Vector2) Div(scalar float64) Vector2 {
	return Vector2{v.X / scalar, v.Y / scalar}
}

// Length returns the magnitude (length) of the vector.
// Example: speed := velocity.Length()
func (v Vector2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// LengthSquared returns the squared magnitude (faster, no sqrt).
// Example: if pos.Sub(target).LengthSquared() < 25 { /* close enough */ }
func (v Vector2) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y
}

// Normalize returns the unit vector in the same direction.
// Example: moveDir := inputVector.Normalize()
func (v Vector2) Normalize() Vector2 {
	len := v.Length()
	if len == 0 {
		return Vector2{0, 0}
	}
	return Vector2{v.X / len, v.Y / len}
}

// Dot returns the dot product of two vectors.
// Example: if velocity.Dot(wallNormal) < 0 { /* moving into wall */ }
func (v Vector2) Dot(o Vector2) float64 {
	return v.X*o.X + v.Y*o.Y
}

// Distance returns the distance between two points.
// Example: dist := playerPos.Distance(enemyPos)
func (v Vector2) Distance(o Vector2) float64 {
	return v.Sub(o).Length()
}

// DistanceSquared returns squared distance between two points.
// Example: if playerPos.DistanceSquared(enemyPos) < 10000 { /* in range */ }
func (v Vector2) DistanceSquared(o Vector2) float64 {
	return v.Sub(o).LengthSquared()
}

// Lerp linearly interpolates between two vectors by t (0 to 1).
// Example: smoothPos := Lerp(playerPos, targetPos, 0.1)
func Lerp(a, b Vector2, t float64) Vector2 {
	return Vector2{
		a.X + (b.X-a.X)*t,
		a.Y + (b.Y-a.Y)*t,
	}
}

// Rotate rotates the vector by an angle in radians.
// Example: rotatedDir := direction.Rotate(math.Pi / 4) // 45° turn
func (v Vector2) Rotate(angle float64) Vector2 {
	cosA := math.Cos(angle)
	sinA := math.Sin(angle)
	return Vector2{
		v.X*cosA - v.Y*sinA,
		v.X*sinA + v.Y*cosA,
	}
}

// Angle returns the angle (in radians) of the vector from the X-axis.
// Example: bulletAngle := bulletDir.Angle()
func (v Vector2) Angle() float64 {
	return math.Atan2(v.Y, v.X)
}

// FromAngle creates a unit vector from an angle in radians.
// Example: dir := FromAngle(playerAimAngle).Mul(speed)
func FromAngle(angle float64) Vector2 {
	return Vector2{math.Cos(angle), math.Sin(angle)}
}

// Reflect returns the vector reflected across a normal.
// Example: bounce := velocity.Reflect(surfaceNormal)
func (v Vector2) Reflect(normal Vector2) Vector2 {
	// Assumes normal is normalized
	return v.Sub(normal.Mul(2 * v.Dot(normal)))
}
