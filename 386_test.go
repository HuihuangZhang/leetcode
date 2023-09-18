package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test386(t *testing.T) {
	testCases := []struct {
		input int
		expect []int
	} {
		{
			input: 13,
			expect: []int{1,10,11,12,13,2,3,4,5,6,7,8,9},
		},
		// {
		// 	input: 2,
		// 	expect: []int{1,2},
		// },
	}

	for i, tt := range testCases {
		t.Run(fmt.Sprintf("case_%v", i), func(t *testing.T) {
			result := lexicalOrder(tt.input)
			assert.Equal(t, tt.expect, result)
		})
	}
	
}