package task_443

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_compress(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		chars []byte
		want  int
	}{
		{
			name: "Case 1",
			chars: []byte("aabbccc"),
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := compress(tt.chars)
			require.Equal(t, tt.want, got)
		})
	}
}
