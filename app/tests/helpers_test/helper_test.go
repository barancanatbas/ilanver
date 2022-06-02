package helperstest

import (
	"ilanver/internal/helpers"
	"testing"
)

func TestRandNumber(t *testing.T) {

	tests := []struct {
		start int
		end   int
	}{
		{start: 1, end: 10},
		{start: 10, end: 20},
		{start: 20, end: 30},
		{start: 30, end: 40},
	}

	for _, test := range tests {
		result := helpers.RandNumber(test.start, test.end)
		if result < test.start || result > test.end {
			t.Errorf("RandNumber(%d, %d) = %d, want %d", test.start, test.end, result, test.start)
		}
	}
}
