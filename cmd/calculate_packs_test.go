package calculatePack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type structure[T interface{}, V interface{}] struct {
	Inputs   T
	Packs    T
	Expected V
}

func TestPackCalculator(t *testing.T) {
	tests := []structure[[]int, map[int]int]{
		{
			Inputs:   []int{501},
			Packs:    []int{250, 500, 1000, 2000, 5000},
			Expected: map[int]int{250: 1, 500: 1},
		},

		{
			Inputs:   []int{251},
			Packs:    []int{250, 500, 1000, 2000, 5000},
			Expected: map[int]int{500: 1},
		},

		{
			Inputs:   []int{12001},
			Packs:    []int{250, 500, 1000, 2000, 5000},
			Expected: map[int]int{5000: 2, 2000: 1, 250: 1},
		},

		{
			Inputs:   []int{1001},
			Packs:    []int{250, 500, 1000, 2000, 5000},
			Expected: map[int]int{1000: 1, 250: 1},
		},

		{
			Inputs:   []int{1},
			Packs:    []int{250, 500, 1000, 2000, 5000},
			Expected: map[int]int{250: 1},
		},

		{
			Inputs:   []int{250},
			Packs:    []int{250, 500, 1000, 2000, 5000},
			Expected: map[int]int{250: 1},
		},

		{
			Inputs:   []int{501},
			Packs:    []int{250, 500, 1000, 2000, 5000},
			Expected: map[int]int{500: 1, 250: 1},
		},
	}

	for _, test := range tests {

		packCalculator := NewPackCalculator(
			test.Packs,
			NewBinarySearcher(test.Packs),
		)

		for _, input := range test.Inputs {
			result := packCalculator.CalculatePacks(input)
			for k, v := range result {
				assert.Equal(t, test.Expected[k], v)
			}
		}
	}
}
