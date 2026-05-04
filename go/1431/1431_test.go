package task_1431

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kidsWithCandies(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		candies      []int
		extraCandies int
		want         []bool
	}{
		struct {
			name         string
			candies      []int
			extraCandies int
			want         []bool
		}{
			name:         "case 1",
			candies:      []int{2, 3, 5, 1, 3},
			extraCandies: 3,
			want:         []bool{true, true, true, false, true},
		},
		struct {
			name         string
			candies      []int
			extraCandies int
			want         []bool
		}{
			name:         "case 2",
			candies:      []int{4, 2, 1, 1, 2},
			extraCandies: 1,
			want:         []bool{true, false, false, false, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kidsWithCandies(tt.candies, tt.extraCandies)
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_findKidWithGreatestCandies(t *testing.T) {
	tests := []struct {
		name string
		candies []int
		want int
	}{
		{
			name: "case 1",
			candies: []int{5,4,1,5,8},
			want: 4,
		},
		{
			name: "case 2",
			candies: []int{5,4,100,5,8},
			want: 2,
		},
		{
			name: "case 2",
			candies: []int{5,4,100,-5,8},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findKidWithGreatestCandies(tt.candies)
			require.Equal(t, got, tt.want)
		})
	}
}
