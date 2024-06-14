package streams

import (
	"testing"
)

func TestMapDoubles(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}
	doubleFunc := func(x int) int { return 2 * x }

	ys := Map(xs, doubleFunc)

	expected := []int{2, 4, 6, 8, 10}
	if len(ys) != len(expected) {
		t.Errorf("len(ys) = %v, got %v", len(ys), len(expected))
	}

	for i := range len(ys) {
		if ys[i] != expected[i] {
			t.Errorf("ys[%v] = %v, got %v", i, ys[i], expected[i])
		}
	}
}
