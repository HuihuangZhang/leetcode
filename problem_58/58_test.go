package problem58

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseLeftWords(t *testing.T) {
	testCases := []struct {
		name string
		inputS string
		inputN int
		wantResult string
	} {
		{
			name: "case_0",
			inputS: "abcdefg",
			inputN: 2,
			wantResult: "cdefgab",
		},
		{
			name: "case_1",
			inputS: "lrloseumgh",
			inputN: 6,
			wantResult: "umghlrlose",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := reverseLeftWords(tt.inputS, tt.inputN)
			assert.Equal(t, tt.wantResult, result)
		})
	}
}