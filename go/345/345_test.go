package task_345

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverseVowels(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want string
	}{
		{
			name: "case1",
			s: "leetcode",
			want: "leotcede",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseVowels(tt.s)
			require.Equal(t, tt.want, got)
		})
	}
}
