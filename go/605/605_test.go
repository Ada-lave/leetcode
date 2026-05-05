package task_605

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canPlaceFlowers(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		flowerbed []int
		n         int
		want      bool
	}{
		{
			name: "case 1",
			flowerbed: []int{1, 0, 0, 0, 1},
			n: 1,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canPlaceFlowers(tt.flowerbed, tt.n)
			require.Equal(t, tt,tt.want, got)
		})
	}
}
