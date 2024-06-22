package vec

import (
	"errors"
	"math"
)

// Vec2 represents a vector in 2D space.
// Many methods are provided - note these are value receivers,
// and therefore never modify the Vec2 being operated upon.
type Vec2 struct {
	X, Y float64
}

// Add computes v1 + v2.
func (v1 Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{
		v1.X + v2.X,
		v1.Y + v2.Y,
	}
}

// Subtract computes v1 - v2.
func (v1 Vec2) Subtract(v2 Vec2) Vec2 {
	return Vec2{
		v1.X - v2.X,
		v1.Y - v2.Y,
	}
}

// Dot computes the dot product between v1 and v2.
func (v1 Vec2) Dot(v2 Vec2) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

// Multiply returns this vector multiplied by a scalar value.
func (v Vec2) Multiply(n float64) Vec2 {
	return Vec2{
		v.X * n,
		v.Y * n,
	}
}

// Divide returns this vector divided by a scalar value.
func (v Vec2) Divide(n float64) (Vec2, error) {
	if n == 0 {
		return Vec2{}, errors.New("tried to divide by 0")
	}

	return Vec2{
		v.X / n,
		v.Y / n,
	}, nil
}

// Magnitude returns the length of this vector.
func (v Vec2) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Normalised returns the vector in the same direction as this vector with a length of 1.
//
// Since a 0-length array has no direction, if |v| = 0 then this function will return an error.
func (v Vec2) Normalised() (Vec2, error) {
	magnitude := v.Magnitude()

	if magnitude == 0 {
		return Vec2{}, errors.New("tried to normalise a 0-length vector")
	}

	return Vec2{
		v.X / magnitude,
		v.Y / magnitude,
	}, nil
}

// Angle computes the angle between v1 and v2, in radians.
func (v1 Vec2) Angle(v2 Vec2) (float64, error) {
	if v1.Magnitude() == 0 {
		return 0, errors.New("v1 length is 0, cannot compute angle")
	}

	if v2.Magnitude() == 0 {
		return 0, errors.New("v2 length is 0, cannot compute angle")
	}

	return math.Acos(v1.Dot(v2) / (v1.Magnitude() * v2.Magnitude())), nil
}

// Lerp linearly interpolates between v1 and v2 by factor t.
//
// Mathematically, this computes v1 * (1 - t) + v2 * t.
//
// At t = 0, the result of this function is equal to v1.
//
// At t = 1, the result of this function is equal to v2.
//
// No safeguards are in place for vales of t that do not satisfy 0 <= t <= 1.
// Instead, this will extrapolate beyond v1 or v2. For example, values of t > 1 will be in the direction of
// (v2 - v1). t = -2 will result in a vector equal to v1 * 3 + v2 * -2.
//
// If you wish to have the result clamped between v1 and v2, use [LerpClamped].
func (v1 Vec2) Lerp(v2 Vec2, t float64) Vec2 {
	return Vec2{
		v1.X + t*(v2.X-v1.X),
		v1.Y + t*(v2.Y-v1.Y),
	}
}

// LerpClamped linearly interpolates between v1 and v2 by factor t, clamping the result between v1 and v2.
//
// Mathematically, this computes v1 * (1 - t) + v2 * t for 0 <= t <= 1.
//
// At t <= 0, the result of this function is equal to v1.
//
// At t >= 1, the result of this function is equal to v2.
func (v1 Vec2) LerpClamped(v2 Vec2, t float64) Vec2 {
	var t_clamped float64

	if t < 0 {
		t_clamped = 0
	} else if t > 1 {
		t_clamped = 1
	} else {
		t_clamped = t
	}

	return v1.Lerp(v2, t_clamped)
}

// Equals returns true if the two vectors are equal.
func (v1 Vec2) Equals(v2 Vec2) bool {
	return v1.X == v2.X && v1.Y == v2.Y
}

// AlmostEqual returns true if the two vectors are almost equal, within some tolerance threshold.
func (v1 Vec2) AlmostEquals(v2 Vec2, threshold float64) bool {
	return math.Abs(v1.X-v2.X) <= threshold && math.Abs(v1.Y-v2.Y) <= threshold
}
