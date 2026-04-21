package task_1768

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mergeAlternately(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		word1 string
		word2 string
		want  string
	}{
		struct{name string; word1 string; word2 string; want string}{
			name: "case 1",
			word1: "abc",
			word2: "ptra",
			want: "apbtcra",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeAlternately(tt.word1, tt.word2)
			require.Equal(t, got, tt.want)
		})
	}
}
