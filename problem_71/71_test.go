package problem71

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimplifyPath(t *testing.T) {
	testCases := []struct {
		name string
		inputPath string
		want string
	} {
		{
			name: "case_0",
			inputPath: "/home/",
			want: "/home",
		},
		{
			name: "case_1",
			inputPath: "/../",
			want: "/",
		},
		{
			name: "case_2",
			inputPath: "/home//foo/",
			want: "/home/foo",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := simplifyPath(tt.inputPath)
			assert.Equal(t, tt.want, result)
		})
	}
}