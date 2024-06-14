package streams

import (
	"testing"
)

func TestMapDoubles(t *testing.T) {
	var tests = []struct {
		xs   []int
		want []int
		name string
	}{
		{
			xs:   []int{1, 2, 3, 4, 5},
			want: []int{2, 4, 6, 8, 10},
			name: "1..5",
		},
	}

	doubleFunc := func(x int) int { return 2 * x }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ys := Map(tt.xs, doubleFunc)
			if len(ys) != len(tt.want) {
				t.Errorf("len(ys) = %v, got %v", len(ys), len(tt.want))
			}

			for i := range len(ys) {
				if ys[i] != tt.want[i] {
					t.Errorf("ys[%v] = %v, got %v", i, ys[i], tt.want[i])
				}
			}
		})
	}
}
