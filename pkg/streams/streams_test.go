package streams

import (
	"testing"
)

func TestMap(t *testing.T) {
	var tests = []struct {
		xs      []int
		want    []int
		mapFunc func(int) int
		name    string
	}{
		{
			xs:      []int{1, 2, 3, 4, 5},
			want:    []int{2, 4, 6, 8, 10},
			mapFunc: func(i int) int { return 2 * i },
			name:    "1..5 double",
		},
		{
			xs:      []int{3, 2, 1},
			want:    []int{0, -1, -2},
			mapFunc: func(i int) int { return i - 3 },
			name:    "3..1 -3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ys := Map(tt.xs, tt.mapFunc)
			if len(ys) != len(tt.want) {
				t.Errorf("len(ys) = %v, expected %v", len(ys), len(tt.want))
			}

			for i := range len(ys) {
				if ys[i] != tt.want[i] {
					t.Errorf("ys[%v] = %v, expected %v", i, ys[i], tt.want[i])
				}
			}
		})
	}
}
