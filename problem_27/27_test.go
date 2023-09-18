package problem27

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveInPlace(t *testing.T) {
	testCases := []struct {
		name string
		inputArray []int
		inputNum int
		expectLen int
		expectArray []int
	} {
		{
			name: "case0",
			inputArray: []int{},
			inputNum: 2,
			expectLen: 0,
			expectArray: []int{},
		},
		{
			name: "case1",
			inputArray: []int{3,2,2,3},
			inputNum: 3,
			expectLen: 2,
			expectArray: []int{2, 2},
		},
		{
			name: "case2",
			inputArray: []int{0,1,2,2,3,0,4,2},
			inputNum: 2,
			expectLen: 5,
			expectArray: []int{ 0, 1, 3, 0, 4},
		},
		{
			name: "case3",
			inputArray: []int{0,1,2,2,3,0,4,2},
			inputNum: 5,
			expectLen: 8,
			expectArray: []int{0,1,2,2,3,0,4,2},
		},
		{
			name: "case4",
			inputArray: []int{3, 3},
			inputNum: 3,
			expectLen: 0,
			expectArray: []int{},
		},
		{
			name: "case5",
			inputArray: []int{4, 5},
			inputNum: 4,
			expectLen: 1,
			expectArray: []int{5},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result, rArr := removeInPlaceV2(tt.inputArray, tt.inputNum)
			assert.Equal(t, tt.expectLen, result)
			fmt.Println(">>>", tt.expectArray, tt.inputArray, rArr)

			for i := 0; i < result; i++ {
				assert.Equal(t, tt.expectArray[i], rArr[i])
			}
		})
	}
}