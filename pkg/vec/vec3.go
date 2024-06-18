package vec

import (
	"errors"
	"math"
)

// Vec3 represents a vector in 3D space.
// Many receivers are provided - note these are value receivers,
// and therefore never modify the Vec3 being operated upon.
type Vec3 struct {
	X, Y, Z float64
}

// NewVec3 returns a new 3D vector.
func NewVec3(x, y, z float64) Vec3 {
	return Vec3{
		X: x,
		Y: y,
		Z: z,
	}
}

// Add computes v1 + v2.
func (v1 Vec3) Add(v2 Vec3) Vec3 {
	return NewVec3(
		v1.X+v2.X,
		v1.Y+v2.Y,
		v1.Z+v2.Z,
	)
}

// Subtract computes v1 - v2.
func (v1 Vec3) Subtract(v2 Vec3) Vec3 {
	return NewVec3(
		v1.X-v2.X,
		v1.Y-v2.Y,
		v1.Z-v2.Z,
	)
}

// Dot computes the dot product between v1 and v2.
func (v1 Vec3) Dot(v2 Vec3) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

// Multiply returns this vector multiplied by a scalar value.
func (v Vec3) Multiply(n float64) Vec3 {
	return NewVec3(
		v.X*n,
		v.Y*n,
		v.Z*n,
	)
}

// Divide returns this vector divided by a scalar value.
func (v Vec3) Divide(n float64) Vec3 {
	return NewVec3(
		v.X/n,
		v.Y/n,
		v.Z/n,
	)
}

// Magnitude returns the length of this vector.
func (v Vec3) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Normalised returns the vector in the same direction as this vector with a length of 1.
//
// Since a 0-length array has no direction, if |v| = 0 then this function will return an error.
func (v Vec3) Normalised() (Vec3, error) {
	magnitude := v.Magnitude()

	if magnitude == 0 {
		return Vec3{}, errors.New("tried to normalise a 0-length vector")
	}

	return NewVec3(
		v.X/magnitude,
		v.Y/magnitude,
		v.Z/magnitude,
	), nil
}

// Angle computes the angle between v1 and v2, in radians.
func (v1 Vec3) Angle(v2 Vec3) (float64, error) {
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
func (v1 Vec3) Lerp(v2 Vec3, t float64) Vec3 {
	return NewVec3(
		v1.X+t*(v2.X-v1.X),
		v1.Y+t*(v2.Y-v1.Y),
		v1.Z+t*(v2.Z-v1.Z),
	)
}

// LerpClamped linearly interpolates between v1 and v2 by factor t, clamping the result between v1 and v2.
//
// Mathematically, this computes v1 * (1 - t) + v2 * t for 0 <= t <= 1.
//
// At t <= 0, the result of this function is equal to v1.
//
// At t >= 1, the result of this function is equal to v2.
func (v1 Vec3) LerpClamped(v2 Vec3, t float64) Vec3 {
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

// Cross returns the cross product of v1 and v2.
func (v1 Vec3) Cross(v2 Vec3) Vec3 {
	return NewVec3(
		v1.Y*v2.Z-v1.Z*v2.Y,
		v1.Z*v2.X-v1.X*v2.Z,
		v1.X*v2.Y-v1.Y*v2.X,
	)
}

// Equals returns true if the two vectors are equal.
func (v1 Vec3) Equals(v2 Vec3) bool {
	return v1.X == v2.X && v1.Y == v2.Y && v1.Z == v2.Z
}

// AlmostEqual returns true if the two vectors are almost equal, within some tolerance threshold.
func (v1 Vec3) AlmostEquals(v2 Vec3, threshold float64) bool {
	return math.Abs(v1.X-v2.X) <= threshold && math.Abs(v1.Y-v2.Y) <= threshold && math.Abs(v1.Z-v2.Z) <= threshold
}
