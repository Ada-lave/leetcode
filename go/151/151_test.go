package task_151

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverseWords(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want string
	}{
		{
			name: "case 1",
			s: "abc got string",
			want: "string got abc",
		},
		{
			name: "case 2",
			s: "a good   example",
			want: "example good a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseWords(tt.s)
			require.Equal(t, tt.want, got)
		})
	}
}
