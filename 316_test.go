package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFn(t *testing.T) {
	testCases := []struct {
		input string
		expect string
	} {
		{
			input: "bcabc",
			expect: "abc",
		},
		{
			input: "cbacdcbc",
			expect: "acdb",
		},
		{
			input: "abacb",
			expect: "abc",
		},
		{
			input: "bbcaac",
			expect: "bac",
		},
	}
	for i, tt := range testCases {
		t.Run(fmt.Sprintf("case_%v", i), func(t *testing.T) {
			result := removeDuplicateLetters(tt.input)
			assert.Equal(t, tt.expect, result)
		})
	}
	
}