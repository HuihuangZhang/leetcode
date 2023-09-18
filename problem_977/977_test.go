package problem977

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquareOrderedList(t *testing.T) {
	testCases := []struct {
		name string
		input []int
		expect []int
	} {
		{
			name: "case_0",
			input: []int{},
			expect: []int{},
		},
		{
			name: "case_1",
			input: []int{-4,-1,0,3,10},
			expect: []int{0,1,9,16,100},
		},
		{
			name: "case_2",
			input: []int{-7,-3,2,3,11},
			expect: []int{4,9,9,49,121},
		},
	}
	
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := SquareOrderedList(tt.input)
			assert.Equal(t, tt.expect, result)
		})
	}
}