package sim

import (
	"fmt"
	"math"
)

// Vec2 represents a 2-dimensional vector.
type Vec2 struct {
	X float64
	Y float64
}

// String returns a string representing v
func (v Vec2) String() string {
	return fmt.Sprintf("<%.2f, %.2f>", v.X, v.Y)
}

// Add returns the sum of two Vec2s.
func (v Vec2) Add(b Vec2) Vec2 {
	return Vec2{
		X: v.X + b.X,
		Y: v.Y + b.Y,
	}
}

// Sub returns the difference of two Vec2s.
func (v Vec2) Sub(b Vec2) Vec2 {
	return Vec2{
		X: v.X - b.X,
		Y: v.Y - b.Y,
	}
}

// Scale multiplies each component of v by a scalar s.
func (v Vec2) Scale(s float64) Vec2 {
	return Vec2{
		X: v.X * s,
		Y: v.Y * s,
	}
}

// Div divides each component of v by a scalar s
func (v Vec2) Div(s float64) Vec2 {
	return Vec2{
		X: v.X / s,
		Y: v.Y / s,
	}
}

// Norm returns the norm (magnitude) of v.
func (v Vec2) Norm() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

// Normalize returns a unit vector at the same angle as v.
func (v Vec2) Normalize() Vec2 {
	return v.Scale(1 / v.Norm())
}
