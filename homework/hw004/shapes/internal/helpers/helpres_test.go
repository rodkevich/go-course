package helpers

import (
	"math"
	"testing"
)

func TestUsedArgsIncludeInvalid(t *testing.T) {
	type args struct {
		args []float64
	}
	tests := []struct {
		name  string
		args  args
		wantB bool
	}{
		{"nil", args{nil}, true},
		{"empty []", args{[]float64{}}, true},
		{"first 0 & pos", args{[]float64{0, math.MaxFloat64}}, true},
		{"first neg & 0", args{[]float64{-math.MaxFloat64, 0}}, true},
		{"first pos & -0", args{[]float64{43, -0}}, true},
		{"pos & pos", args{[]float64{1, 1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := UsedArgsIncludeInvalid(tt.args.args); gotB != tt.wantB {
				t.Errorf("UsedArgsIncludeInvalid() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}
