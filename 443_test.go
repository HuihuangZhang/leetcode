package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test443(t *testing.T) {
	testCases := []struct {
		input []byte
		expectChars []byte
	} {
		{
			input: []byte{'a','a','b','b','c','c','c'},
			expectChars: []byte{'a','2','b','2','c','3'},
		},
		{
			input: []byte{'a'},
			expectChars: []byte{'a'},
		},
		{
			input: []byte{'a'},
			expectChars: []byte{'a'},
		},
		{
			input: []byte{'a','b','b','b','b','b','b','b','b','b','b','b','b'},
			expectChars: []byte{'a','b','1','2'},
		},
	}


	for i, tt := range testCases {
		t.Run(fmt.Sprintf("case_%v", i), func(t *testing.T) {
			result := compress(tt.input)
			assert.Equal(t, len(tt.expectChars), result)
			assert.Equal(t, tt.expectChars, tt.input[:len(tt.expectChars)])
		})
	}
}