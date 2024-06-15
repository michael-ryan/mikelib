package vec

import (
	"errors"
	"math"
)

// vec3 represents a vector in 3D space.
// Many receivers are provided - note these are value receivers,
// and therefore never modify the vec3 being operated upon.
//
// vec3s may hold any float or integer type, but note that you may only operate on like-for-like.
// That is to say, for example, you may only multiply a vec3[float64] with another vec3[float64].
//
// It is recommended that you choose one underlying numeric type, and stick to it.
type vec3[T number] struct {
	X, Y, Z T
}

// NewVec3 returns a new 3D vector.
func NewVec3[T number](x T, y T, z T) vec3[T] {
	return vec3[T]{
		X: x,
		Y: y,
		Z: z,
	}
}

// Add computes v1 + v2.
func (v1 vec3[T]) Add(v2 vec3[T]) vec3[T] {
	return NewVec3(
		v1.X+v2.X,
		v1.Y+v2.Y,
		v1.Z+v2.Z,
	)
}

// Subtract computes v1 - v2.
func (v1 vec3[T]) Subtract(v2 vec3[T]) vec3[T] {
	return NewVec3(
		v1.X-v2.X,
		v1.Y-v2.Y,
		v1.Z-v2.Z,
	)
}

// Dot computes the dot product between v1 and v2.
func (v1 vec3[T]) Dot(v2 vec3[T]) T {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

// Multiply returns this vector multiplied by a scalar value.
func (v vec3[T]) Multiply(n T) vec3[T] {
	return NewVec3(
		v.X*n,
		v.Y*n,
		v.Z*n,
	)
}

// Divide returns this vector divided by a scalar value.
func (v vec3[T]) Divide(n T) vec3[T] {
	return NewVec3(
		v.X/n,
		v.Y/n,
		v.Z/n,
	)
}

// Magnitude returns the length of this vector.
func (v vec3[T]) Magnitude() float64 {
	x := float64(v.X)
	y := float64(v.Y)
	z := float64(v.Z)
	return math.Sqrt(x*x + y*y + z*z)
}

// Normalised returns the vector in the same direction as this vector with a length of 1.
//
// Since a 0-length array has no direction, if |v| = 0 then this function will return an error.
func (v vec3[T]) Normalised() (vec3[float64], error) {
	magnitude := v.Magnitude()

	if magnitude == 0 {
		return vec3[float64]{}, errors.New("tried to normalise a 0-length vector")
	}

	return NewVec3(
		float64(v.X)/magnitude,
		float64(v.Y)/magnitude,
		float64(v.Z)/magnitude,
	), nil
}

// Angle computes the angle between v1 and v2, in radians.
func (v1 vec3[T]) Angle(v2 vec3[T]) (float64, error) {
	if v1.Magnitude() == 0 {
		return 0, errors.New("v1 length is 0, cannot compute angle")
	}

	if v2.Magnitude() == 0 {
		return 0, errors.New("v2 length is 0, cannot compute angle")
	}

	return math.Acos(float64(v1.Dot(v2)) / (v1.Magnitude() * v2.Magnitude())), nil
}

// Lerp linearly interpolates between v1 and v2 by factor t.
// Note that this function will always return a vec2[float64].
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
func (v1 vec3[T]) Lerp(v2 vec3[T], t float64) vec3[float64] {
	return NewVec3(
		float64(v1.X)+t*float64(v2.X-v1.X),
		float64(v1.Y)+t*float64(v2.Y-v1.Y),
		float64(v1.Z)+t*float64(v2.Z-v1.Z),
	)
}

func (v1 vec3[T]) Cross(v2 vec3[T]) vec3[T] {
	return NewVec3(
		v1.Y*v2.Z-v1.Z*v2.Y,
		v1.Z*v2.X-v1.X*v2.Z,
		v1.X*v2.Y-v1.Y*v2.X,
	)
}
