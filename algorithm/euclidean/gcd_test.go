package gcd

import (
	"fmt"
	"testing"
)

func TestGCD(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{10, 5, 5},
		{25, 5, 5},
		{30, 15, 15},
		{30, 9, 3},
		{100, 9, 1},
		{
			2 * 2 * 3 * 3 * 5,
			2 * 3 * 5 * 7 * 13,
			2 * 3 * 5,
		}, {
			2 * 2 * 3 * 3 * 13,
			2 * 3 * 5 * 7 * 13,
			2 * 3 * 13,
		}, {
			2 * 3 * 5 * 7 * 11 * 13 * 17 * 19,
			3 * 3 * 7 * 7 * 11 * 11 * 17 * 17,
			3 * 7 * 11 * 17,
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v,%v)", tc.a, tc.b), func(t *testing.T) {
			got := GCD(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("GCD() = %v; want %v", got, tc.want)
			}
		})
	}
}
