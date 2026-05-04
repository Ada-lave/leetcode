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
		struct{name string; str1 string; str2 string; want string}{
			name: "case 5",
			str1: "ABABABAB",
			str2: "ABAB",
			want: "ABAB",			
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gcdOfStrings(tt.str1, tt.str2)
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_gcd(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		a    int
		b    int
		want int
	}{
		struct{name string; a int; b int; want int}{
			name: "case 1",
			a: 18,
			b: 12,
			want: 6,
		},
		struct{name string; a int; b int; want int}{
			name: "case 2",
			a: 12,
			b: 18,
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gcd(tt.a, tt.b)
			// TODO: update the condition below to compare got with tt.want.
			require.Equal(t, got, tt.want)
		})
	}
}
