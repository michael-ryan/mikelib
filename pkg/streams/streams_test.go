package streams

import (
	"fmt"
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
			name: "1..5 *2",
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
		{
			name: "Empty input",
			args: args{
				xs:      []int{},
				mapFunc: func(i int) int { return i },
			},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ys := Map(tt.args.xs, tt.args.mapFunc)
			if len(ys) != len(tt.want) {
				t.Errorf("len(Map()) = %v, want %v", len(ys), len(tt.want))
			}

			for i := range len(ys) {
				if ys[i] != tt.want[i] {
					t.Errorf("Map()[%v] = %v, want %v", i, ys[i], tt.want[i])
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
		{
			name: "1..10 >5",
			args: args{
				xs:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				predicate: func(i int) bool { return i > 5 },
			},
			want: []int{6, 7, 8, 9, 10},
		},
		{
			name: "1..3 =5",
			args: args{
				xs:        []int{1, 2, 3},
				predicate: func(i int) bool { return i == 5 },
			},
			want: []int{},
		},
		{
			name: "1..3 =2",
			args: args{
				xs:        []int{1, 2, 3},
				predicate: func(i int) bool { return i == 2 },
			},
			want: []int{2},
		},
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
		want []int
	}{
		{
			name: "Empty",
			args: args{
				xs: []int{},
			},
			want: []int{},
		},
		{
			name: "1..5",
			args: args{
				xs: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]int, 0)
			for g := range ToGenerator(tt.args.xs) {
				got = append(got, g)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToGenerator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollect(t *testing.T) {
	type args struct {
		xs []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Empty",
			args: args{
				xs: []int{},
			},
			want: []int{},
		},
		{
			name: "1..5",
			args: args{
				xs: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := make(chan int)

			go func(xs []int, c chan<- int) {
				for _, x := range xs {
					c <- x
				}
				close(c)
			}(tt.args.xs, c)

			if got := Collect(c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Collect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFold(t *testing.T) {
	type args struct {
		xs           []int
		folder       func(int, int) int
		initialValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1..5 + 0",
			args: args{
				xs:           []int{1, 2, 3, 4, 5},
				folder:       func(x, y int) int { return x + y },
				initialValue: 0,
			},
			want: 15,
		},
		{
			name: "1..5 * 1",
			args: args{
				xs:           []int{1, 2, 3, 4, 5},
				folder:       func(x, y int) int { return x * y },
				initialValue: 1,
			},
			want: 120,
		},
		{
			name: "1..5 * 2",
			args: args{
				xs:           []int{1, 2, 3, 4, 5},
				folder:       func(x, y int) int { return x * y },
				initialValue: 2,
			},
			want: 240,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fold(tt.args.xs, tt.args.folder, tt.args.initialValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fold() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Don't really want to bother with mocking as we are just looking for side effects here
// Keep this test here to catch runtime errors.
func TestForeach(t *testing.T) {
	type args struct {
		xs []int
		f  func(int)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Sprint 1..5",
			args: args{
				xs: []int{1, 2, 3, 4, 5},
				f:  func(i int) { fmt.Sprintln(i) },
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Foreach(tt.args.xs, tt.args.f)
		})
	}
}
