package streams

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	type args struct {
		xs      []int
		mapFunc func(int) int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1..5 double",
			args: args{
				xs:      []int{1, 2, 3, 4, 5},
				mapFunc: func(i int) int { return 2 * i },
			},
			want: []int{2, 4, 6, 8, 10},
		},
		{
			name: "3..1 -3",
			args: args{
				xs:      []int{3, 2, 1},
				mapFunc: func(i int) int { return i - 3 },
			},
			want: []int{0, -1, -2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ys := Map(tt.args.xs, tt.args.mapFunc)
			if len(ys) != len(tt.want) {
				t.Errorf("len(ys) = %v, want %v", len(ys), len(tt.want))
			}

			for i := range len(ys) {
				if ys[i] != tt.want[i] {
					t.Errorf("ys[%v] = %v, want %v", i, ys[i], tt.want[i])
				}
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		xs        []int
		predicate func(int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.xs, tt.args.predicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToGenerator(t *testing.T) {
	type args struct {
		xs []int
	}
	tests := []struct {
		name string
		args args
		want <-chan int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToGenerator(tt.args.xs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToGenerator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollect(t *testing.T) {
	type args struct {
		xs <-chan int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Collect(tt.args.xs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Collect() = %v, want %v", got, tt.want)
			}
		})
	}
}
