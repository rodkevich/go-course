package helpers

import (
	"math"
	"testing"
)

func TestUsedArgsIncludeInvalid(t *testing.T) {
	tests := []struct {
		name  string
		args  []float64
		wantB bool
	}{
		{"nil", nil, true},
		{"empty []", []float64{}, true},
		{"first 0 & pos", []float64{0, math.MaxFloat64}, true},
		{"first neg & 0", []float64{-math.MaxFloat64, 0}, true},
		{"first pos & -0", []float64{43, -0}, true},
		{"pos & pos", []float64{1, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := UsedArgsIncludeInvalid(tt.args); gotB != tt.wantB {
				t.Errorf("UsedArgsIncludeInvalid() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}
