package xmath

import (
	"math"
)

// ========================
// ANGLES & ROTATIONS
// ========================

// DegToRad converts degrees to radians.
// Example: angleRad := DegToRad(90)
func DegToRad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

// RadToDeg converts radians to degrees.
// Example: angleDeg := RadToDeg(math.Pi)
func RadToDeg(rad float64) float64 {
	return rad * (180 / math.Pi)
}

// SinCos returns sin(angle), cos(angle) for given radians.
// Useful for circular motion, aiming, oscillations.
func SinCos(rad float64) (float64, float64) {
	return math.Sin(rad), math.Cos(rad)
}

// CircularMovement returns a point on a circle given center, radius, and angle in radians.
// Example: enemyPos := CircularMovement(centerX, centerY, 5, timeElapsed)
func CircularMovement(cx, cy, radius, angleRad float64) (float64, float64) {
	s, c := SinCos(angleRad)
	return cx + radius*c, cy + radius*s
}

// ========================
// DISTANCE FORMULAS
// ========================

// EuclideanDistance returns the distance between two points.
// Euclidean Distance:
// - Use when actual distances matter (e.g., clustering, nearest-neighbor search, measuring geometric distance).
// - Needed when comparisons must be interpretable in the original units.
//
// Example: dist := EuclideanDistance(p1x, p1y, p2x, p2y)
func EuclideanDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(SquaredDistance(x1, y1, x2, y2))
}

// SquaredDistance returns the squared distance between two points (faster).
// Squared Distance:
// - Often used in optimization because it is easier to compute and differentiate (no square root).
// - Common in k-means clustering (objective function minimizes squared distances, though nearest-centroid assignment still uses Euclidean distance).
// - Used in many machine learning loss functions (like Mean Squared Error).
//
// Example: if SquaredDistance(x1, y1, x2, y2) < 25 { /* close enough */ }
func SquaredDistance(x1, y1, x2, y2 float64) float64 {
	return (x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)
}

// ========================
// SHAPES
// ========================

// Circle represents a circle shape.
type Circle struct {
	X, Y   float64
	Radius float64
}

// Rect represents an axis-aligned rectangle.
type Rect struct {
	X, Y          float64 // Top-left corner
	Width, Height float64
}

// Polygon represents a convex polygon with points in local space.
type Polygon struct {
	Points []Vector2 // Using Vector2 from other file
}

// ========================
// COLLISION DETECTION
// ========================

// CircleIntersect checks if two circles intersect.
func CircleIntersect(c1, c2 Circle) bool {
	rSum := c1.Radius + c2.Radius
	return SquaredDistance(c1.X, c1.Y, c2.X, c2.Y) <= rSum*rSum
}

// RectIntersect checks if two axis-aligned rectangles intersect.
func RectIntersect(r1, r2 Rect) bool {
	return !(r1.X > r2.X+r2.Width ||
		r1.X+r1.Width < r2.X ||
		r1.Y > r2.Y+r2.Height ||
		r1.Y+r1.Height < r2.Y)
}

// OBBvsOBB checks if two oriented bounding boxes intersect.
// This uses the Separating Axis Theorem (SAT) for convex polygons.
func OBBvsOBB(p1, p2 Polygon, pos1, pos2 Vector2, rot1, rot2 float64) bool {
	// Transform points to world space
	world1 := transformPolygon(p1, pos1, rot1)
	world2 := transformPolygon(p2, pos2, rot2)

	// Check both sets of edges for separating axis
	if hasSeparatingAxis(world1, world2) {
		return false
	}
	if hasSeparatingAxis(world2, world1) {
		return false
	}
	return true
}

// ========================
// HELPER FUNCTIONS FOR OBB
// ========================

func transformPolygon(poly Polygon, pos Vector2, rot float64) []Vector2 {
	cosA := math.Cos(rot)
	sinA := math.Sin(rot)
	transformed := make([]Vector2, len(poly.Points))
	for i, p := range poly.Points {
		// Rotate
		rotX := p.X*cosA - p.Y*sinA
		rotY := p.X*sinA + p.Y*cosA
		// Translate
		transformed[i] = Vector2{rotX + pos.X, rotY + pos.Y}
	}
	return transformed
}

func hasSeparatingAxis(poly1, poly2 []Vector2) bool {
	count := len(poly1)
	for i := 0; i < count; i++ {
		// Get current edge
		p1 := poly1[i]
		p2 := poly1[(i+1)%count]
		edge := Vector2{p2.X - p1.X, p2.Y - p1.Y}

		// Get perpendicular axis
		axis := Vector2{-edge.Y, edge.X}.Normalize()

		// Project both polygons
		min1, max1 := projectPolygon(axis, poly1)
		min2, max2 := projectPolygon(axis, poly2)

		// Check gap
		if max1 < min2 || max2 < min1 {
			return true
		}
	}
	return false
}

func projectPolygon(axis Vector2, poly []Vector2) (float64, float64) {
	min := axis.Dot(poly[0])
	max := min
	for _, p := range poly[1:] {
		proj := axis.Dot(p)
		if proj < min {
			min = proj
		}
		if proj > max {
			max = proj
		}
	}
	return min, max
}
