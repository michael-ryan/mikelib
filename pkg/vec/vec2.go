package vec

import (
	"errors"
	"math"

	"golang.org/x/exp/constraints"
)

type number interface {
	constraints.Float | constraints.Integer
}

// vec2 represents a vector in 2D space.
// Many receivers are provided - note these are value receivers,
// and therefore never modify the vec2 being operated upon.
//
// vec2s may hold any float or integer type, but note that you may only operate on like-for-like.
// That is to say, for example, you may only multiply a vec2[float64] with another vec2[float64].
//
// It is recommended that you choose one underlying numeric type, and stick to it.
type vec2[T number] struct {
	X, Y T
}

// NewVec2 returns a new 2D vector.
func NewVec2[T number](x T, y T) vec2[T] {
	return vec2[T]{
		X: x,
		Y: y,
	}
}

// Add computes v1 + v2.
func (v1 vec2[T]) Add(v2 vec2[T]) vec2[T] {
	return NewVec2(
		v1.X+v2.X,
		v1.Y+v2.Y,
	)
}

// Subtract computes v1 - v2.
func (v1 vec2[T]) Subtract(v2 vec2[T]) vec2[T] {
	return NewVec2(
		v1.X-v2.X,
		v1.Y-v2.Y,
	)
}

// Dot computes the dot product between v1 and v2.
func (v1 vec2[T]) Dot(v2 vec2[T]) T {
	return v1.X*v2.X + v1.Y*v2.Y
}

// Multiply returns this vector multiplied by a scalar value.
func (v vec2[T]) Multiply(n T) vec2[T] {
	return NewVec2(
		v.X*n,
		v.Y*n,
	)
}

// Divide returns this vector divided by a scalar value.
func (v vec2[T]) Divide(n T) vec2[T] {
	return NewVec2(
		v.X/n,
		v.Y/n,
	)
}

// Magnitude returns the length of this vector.
func (v vec2[T]) Magnitude() float64 {
	x := float64(v.X)
	y := float64(v.Y)
	return math.Sqrt(x*x + y*y)
}

// Normalised returns the vector in the same direction as this vector with a length of 1.
//
// Since a 0-length array has no direction, if |v| = 0 then this function will return an error.
func (v vec2[T]) Normalised() (vec2[float64], error) {
	magnitude := v.Magnitude()

	if magnitude == 0 {
		return vec2[float64]{}, errors.New("tried to normalise a 0-length vector")
	}

	return NewVec2(
		float64(v.X)/magnitude,
		float64(v.Y)/magnitude,
	), nil
}

// Angle computes the angle between v1 and v2, in radians.
func (v1 vec2[T]) Angle(v2 vec2[T]) (float64, error) {
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
func (v1 vec2[T]) Lerp(v2 vec2[T], t float64) vec2[float64] {
	return NewVec2(
		float64(v1.X)+t*float64(v2.X-v1.X),
		float64(v1.Y)+t*float64(v2.Y-v1.Y),
	)
	// return v1.Multiply((1 - T(t))).Add(v2.Multiply(T(t)))
}
