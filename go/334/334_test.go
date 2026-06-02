package task_334

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_increasingTriplet(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want bool
	}{
		{
			name: "Case 1",
			nums: []int{1, 0, 3, 2, 5},
			want: true,
		},
		{
			name: "Case 2",
			nums: []int{0,4,2,1,0,-1,-3},
			want: false,
		},
		{
			name: "Case 3",
			nums: []int{1,6,2,5,1},
			want: true,
		},
		{
			name: "Case 4",
			nums: []int{20,100,10,12,5,13},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := increasingTriplet(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
