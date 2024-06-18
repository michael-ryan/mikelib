package vec

import (
	"math"
	"testing"
)

func TestNewVec2(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		y    float64
		want Vec2
	}{
		{
			name: "(0,0)",
			x:    0,
			y:    0,
			want: Vec2{0, 0},
		},
		{
			name: "(1,2)",
			x:    1,
			y:    2,
			want: Vec2{1, 2},
		},
		{
			name: "(3.0134,-4.2162)",
			x:    3.0134,
			y:    -4.2162,
			want: Vec2{3.0134, -4.2162},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVec2(tt.x, tt.y); !got.Equals(tt.want) {
				t.Errorf("NewVec2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_Add(t *testing.T) {
	tests := []struct {
		name string
		v1   Vec2
		v2   Vec2
		want Vec2
	}{
		{
			name: "(0,0) + (0,0) = (0,0)",
			v1:   Vec2{0, 0},
			v2:   Vec2{0, 0},
			want: Vec2{0, 0},
		},
		{
			name: "(1,2) + (3,4) = (4,6)",
			v1:   Vec2{1, 2},
			v2:   Vec2{3, 4},
			want: Vec2{4, 6},
		},
		{
			name: "(-20, 3) + (14, -1) = (-6, 2)",
			v1:   Vec2{-20, 3},
			v2:   Vec2{14, -1},
			want: Vec2{-6, 2},
		},
		// can't really test the below thanks to rounding errors
		// unless we use AlmostEquals
		// {
		// 	name: "(-3, 2.01) + (2.05, 0.003) = (-0.95, 2.013)",
		// 	v1:   Vec2{-3, 2.01},
		// 	v2:   Vec2{2.05, 0.003},
		// 	want: Vec2{-0.95, 2.013},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.Add(tt.v2); !got.Equals(tt.want) {
				t.Errorf("Vec2.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_Subtract(t *testing.T) {
	tests := []struct {
		name string
		v1   Vec2
		v2   Vec2
		want Vec2
	}{
		{
			name: "(0,0) - (0,0) = (0,0)",
			v1:   Vec2{0, 0},
			v2:   Vec2{0, 0},
			want: Vec2{0, 0},
		},
		{
			name: "(1,2) - (3,4) = (-2,-2)",
			v1:   Vec2{1, 2},
			v2:   Vec2{3, 4},
			want: Vec2{-2, -2},
		},
		{
			name: "(-20, 3) - (14, -1) = (-34, 4)",
			v1:   Vec2{-20, 3},
			v2:   Vec2{14, -1},
			want: Vec2{-34, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.Subtract(tt.v2); !got.Equals(tt.want) {
				t.Errorf("v1.Subtract(v2) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_Dot(t *testing.T) {
	tests := []struct {
		name string
		v1   Vec2
		v2   Vec2
		want float64
	}{
		{
			name: "(0,0) . (0,0) = 0",
			v1:   Vec2{0, 0},
			v2:   Vec2{0, 0},
			want: 0,
		},
		{
			name: "(1,2) . (3,4) = 11",
			v1:   Vec2{1, 2},
			v2:   Vec2{3, 4},
			want: 11,
		},
		{
			name: "(-20, 3) . (14, -1) = -283",
			v1:   Vec2{-20, 3},
			v2:   Vec2{14, -1},
			want: -283,
		},
		{
			name: "(0, 0) . (1, 40) = 0",
			v1:   Vec2{0, 0},
			v2:   Vec2{1, 40},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.Dot(tt.v2); got != tt.want {
				t.Errorf("v1.Dot(v2) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_Multiply(t *testing.T) {
	tests := []struct {
		name string
		v    Vec2
		n    float64
		want Vec2
	}{
		{
			name: "(0,0) * 0 = (0,0)",
			v:    Vec2{0, 0},
			n:    0,
			want: Vec2{0, 0},
		},
		{
			name: "(1,2) * 3 = (3,6)",
			v:    Vec2{1, 2},
			n:    3,
			want: Vec2{3, 6},
		},
		{
			name: "(-20, 3) * 0.5 = (-10, 1.5)",
			v:    Vec2{-20, 3},
			n:    0.5,
			want: Vec2{-10, 1.5},
		},
		{
			name: "(21, 12) * 10 = (210, 120)",
			v:    Vec2{21, 12},
			n:    10,
			want: Vec2{210, 120},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Multiply(tt.n); !got.Equals(tt.want) {
				t.Errorf("v.Multiply(n) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_Divide(t *testing.T) {
	tests := []struct {
		name    string
		v       Vec2
		n       float64
		want    Vec2
		wantErr bool
	}{
		{
			name:    "(0,0) / 0 = error",
			v:       Vec2{0, 0},
			n:       0,
			want:    Vec2{0, 0},
			wantErr: true,
		},
		{
			name:    "(1, 3) / 0 = error",
			v:       Vec2{1, 3},
			n:       0,
			want:    Vec2{0, 0},
			wantErr: true,
		},
		{
			name:    "(1, 3) / -2 = (-0.5, -1.5)",
			v:       Vec2{1, 3},
			n:       -2,
			want:    Vec2{-0.5, -1.5},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Divide(tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("v.Divide(n) error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.Equals(tt.want) {
				t.Errorf("v.Divide(n) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_Magnitude(t *testing.T) {
	tests := []struct {
		name string
		v    Vec2
		want float64
	}{
		{
			name: "(0,0) = 0",
			v:    Vec2{0, 0},
			want: 0,
		},
		{
			name: "(3,4) = 5",
			v:    Vec2{3, 4},
			want: 5,
		},
		{
			name: "(-3,4) = 5",
			v:    Vec2{-3, 4},
			want: 5,
		},
		{
			name: "(-6,-8) = 10",
			v:    Vec2{-6, -8},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Magnitude(); got != tt.want {
				t.Errorf("v.Magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_Normalised(t *testing.T) {
	tests := []struct {
		name    string
		v       Vec2
		want    Vec2
		wantErr bool
	}{
		{
			name:    "(0,0) = error",
			v:       Vec2{0, 0},
			want:    Vec2{0, 0},
			wantErr: true,
		},
		{
			name:    "(3,4) = (0.6, 0.8)",
			v:       Vec2{3, 4},
			want:    Vec2{0.6, 0.8},
			wantErr: false,
		},
		{
			name:    "(-3,4) = (-0.6, 0.8)",
			v:       Vec2{-3, 4},
			want:    Vec2{-0.6, 0.8},
			wantErr: false,
		},
		{
			name:    "(-6,-8) = (-0.6, -0.8)",
			v:       Vec2{-6, -8},
			want:    Vec2{-0.6, -0.8},
			wantErr: false,
		},
		{
			name:    "(0, 1) = (0, 1)",
			v:       Vec2{0, 1},
			want:    Vec2{0, 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Normalised()
			if (err != nil) != tt.wantErr {
				t.Errorf("v.Normalised() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.Equals(tt.want) {
				t.Errorf("v.Normalised() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_Angle(t *testing.T) {
	tests := []struct {
		name    string
		v1      Vec2
		v2      Vec2
		want    float64
		wantErr bool
	}{
		{
			name:    "angle((0,0), (1,2)) = error",
			v1:      Vec2{0, 0},
			v2:      Vec2{1, 2},
			want:    0,
			wantErr: true,
		},
		{
			name:    "angle((0,1), (1, 0)) = pi/2",
			v1:      Vec2{0, 1},
			v2:      Vec2{1, 0},
			want:    math.Pi / 2,
			wantErr: false,
		},
		{
			name:    "angle((0,1), (0, -1)) = pi",
			v1:      Vec2{0, 1},
			v2:      Vec2{0, -1},
			want:    math.Pi,
			wantErr: false,
		},
		{
			name:    "angle((0,1), (1, 1)) = pi/4",
			v1:      Vec2{0, 1},
			v2:      Vec2{1, 1},
			want:    math.Pi / 4,
			wantErr: false,
		},
		{
			name:    "angle((0,1), (0, 0)) = error",
			v1:      Vec2{0, 1},
			v2:      Vec2{0, 0},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v1.Angle(tt.v2)
			if (err != nil) != tt.wantErr {
				t.Errorf("v1.Angle(v2) error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if math.Abs(got-tt.want) > 0.0000001 {
				t.Errorf("v1.Angle(v2) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_Lerp(t *testing.T) {
	tests := []struct {
		name string
		v1   Vec2
		v2   Vec2
		t    float64
		want Vec2
	}{
		{
			name: "(0,0).lerp((1,1), 0) = (0,0)",
			v1:   Vec2{0, 0},
			v2:   Vec2{1, 1},
			t:    0,
			want: Vec2{0, 0},
		},
		{
			name: "(0,0).lerp((2,2), 0.5) = (1,1)",
			v1:   Vec2{0, 0},
			v2:   Vec2{2, 2},
			t:    0.5,
			want: Vec2{1, 1},
		},
		{
			name: "(0,0).lerp((-2,2), 1) = (-2,2)",
			v1:   Vec2{0, 0},
			v2:   Vec2{-2, 2},
			t:    1,
			want: Vec2{-2, 2},
		},
		{
			name: "(1, 2).lerp((3, 4), 3) = (7, 8)",
			v1:   Vec2{1, 2},
			v2:   Vec2{3, 4},
			t:    3,
			want: Vec2{7, 8},
		},
		{
			name: "(1, 2).lerp((3, 4), 0.5) = (2, 3)",
			v1:   Vec2{1, 2},
			v2:   Vec2{3, 4},
			t:    0.5,
			want: Vec2{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.Lerp(tt.v2, tt.t); !got.Equals(tt.want) {
				t.Errorf("v1.Lerp(v2, t) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_LerpClamped(t *testing.T) {
	tests := []struct {
		name string
		v1   Vec2
		v2   Vec2
		t    float64
		want Vec2
	}{
		{
			name: "(0,0).lerp((1,1), 0) = (0,0)",
			v1:   Vec2{0, 0},
			v2:   Vec2{1, 1},
			t:    0,
			want: Vec2{0, 0},
		},
		{
			name: "(0,0).lerp((2,2), 0.5) = (1,1)",
			v1:   Vec2{0, 0},
			v2:   Vec2{2, 2},
			t:    0.5,
			want: Vec2{1, 1},
		},
		{
			name: "(0,0).lerp((-2,2), 1) = (-2,2)",
			v1:   Vec2{0, 0},
			v2:   Vec2{-2, 2},
			t:    1,
			want: Vec2{-2, 2},
		},
		{
			name: "(1, 2).lerp((3, 4), 3) = (3, 4)",
			v1:   Vec2{1, 2},
			v2:   Vec2{3, 4},
			t:    3,
			want: Vec2{3, 4},
		},
		{
			name: "(1, 2).lerp((3, 4), -40) = (1, 2)",
			v1:   Vec2{1, 2},
			v2:   Vec2{3, 4},
			t:    -40,
			want: Vec2{1, 2},
		},
		{
			name: "(1, 2).lerp((3, 4), 0.5) = (2, 3)",
			v1:   Vec2{1, 2},
			v2:   Vec2{3, 4},
			t:    0.5,
			want: Vec2{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.LerpClamped(tt.v2, tt.t); !got.Equals(tt.want) {
				t.Errorf("v1.LerpClamped(v2, t) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_Equals(t *testing.T) {
	tests := []struct {
		name string
		v1   Vec2
		v2   Vec2
		want bool
	}{
		{
			name: "(0,0).Equals((0,0)) = true",
			v1:   Vec2{0, 0},
			v2:   Vec2{0, 0},
			want: true,
		},
		{
			name: "(0,0).Equals((1,1)) = false",
			v1:   Vec2{0, 0},
			v2:   Vec2{1, 1},
			want: false,
		},
		{
			name: "(1,-1).Equals((1,1)) = false",
			v1:   Vec2{1, -1},
			v2:   Vec2{1, 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.Equals(tt.v2); got != tt.want {
				t.Errorf("v1.Equals(v2) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_AlmostEqual(t *testing.T) {
	tests := []struct {
		name      string
		v1        Vec2
		v2        Vec2
		threshold float64
		want      bool
	}{
		{
			name:      "(0,0).AlmostEqual((0,0), 0) = true",
			v1:        Vec2{0, 0},
			v2:        Vec2{0, 0},
			threshold: 0,
			want:      true,
		},
		{
			name:      "(0,0).AlmostEqual((1,1), 0) = false",
			v1:        Vec2{0, 0},
			v2:        Vec2{1, 1},
			threshold: 0,
			want:      false,
		},
		{
			name:      "(0,0).AlmostEqual((1,1), 1) = true",
			v1:        Vec2{0, 0},
			v2:        Vec2{1, 1},
			threshold: 1,
			want:      true,
		},
		{
			name:      "(0.001,0).AlmostEqual((0,0), 0.00001) = false",
			v1:        Vec2{0.001, 0},
			v2:        Vec2{0, 0},
			threshold: 0.00001,
			want:      false,
		},
		{
			name:      "(0.001,0).AlmostEqual((0,0), 0.01) = true",
			v1:        Vec2{0.001, 0},
			v2:        Vec2{0, 0},
			threshold: 0.01,
			want:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v1.AlmostEquals(tt.v2, tt.threshold); got != tt.want {
				t.Errorf("v1.AlmostEqual(v2, threshold) = %v, want %v", got, tt.want)
			}
		})
	}
}
