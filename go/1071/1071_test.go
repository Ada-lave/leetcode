package task_1071

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_gcdOfStrings(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		str1 string
		str2 string
		want string
	}{
		struct{name string; str1 string; str2 string; want string}{
			name: "case 1",
			str1: "ABCABC",
			str2: "ABC",
			want: "ABC",
		},
		struct{name string; str1 string; str2 string; want string}{
			name: "case 2",
			str1: "AAAAAB",
			str2: "AAA",
			want: "",			
		},
		struct{name string; str1 string; str2 string; want string}{
			name: "case 3",
			str1: "CXTXNCXTXNCXTXNCXTXNCXTXN",
			str2: "CXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXN",
			want: "CXTXN",			
		},
		struct{name string; str1 string; str2 string; want string}{
			name: "case 4",
			str1: "AA",
			str2: "A",
			want: "A",			
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gcdOfStrings(tt.str1, tt.str2)
			require.Equal(t, tt.want, got)
		})
	}
}
