package problem209

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinSubArrayLen(t *testing.T) {
	testCases := []struct {
		name string
		inputArr []int
		inputS int
		expectLen int
	} {
		{
			name: "case_0",
			inputArr: []int{2,3,1,2,4,3},
			inputS: 7,
			expectLen: 2,
		},
		{
			name: "case_1",
			inputArr: []int{1,4,4},
			inputS: 4,
			expectLen: 1,
		},
		{
			name: "case_2",
			inputArr: []int{1,1,1,1,1,1,1,1},
			inputS: 11,
			expectLen: 0,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := minSubArrayLen(tt.inputArr, tt.inputS)
			assert.Equal(t, tt.expectLen, result)
		})
	}
}
