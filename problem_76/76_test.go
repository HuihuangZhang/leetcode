package problem76

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinWindow(t *testing.T) {
	testCases := []struct{
		name string
		inputS string
		inputT string
		want string
	} {
		{
			name: "case 0",
			inputS: "ADOBECODEBANC",
			inputT: "ABC",
			want: "BANC",
		},
		{
			name: "case 1",
			inputS: "a",
			inputT: "a",
			want: "a",
		},
		{
			name: "case 2",
			inputS: "a",
			inputT: "aa",
			want: "",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := minWindow(tt.inputS, tt.inputT)
			assert.Equal(t, tt.want, result)
		})
	}
}