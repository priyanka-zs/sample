package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_swap(t *testing.T) {
	testcases := []struct {
		desc    string
		inputA  int
		inputB  int
		outputA int
		outputB int
	}{
		{"success", 2, 3, 3, 2},
	}
	for _, tc := range testcases {
		out1, out2 := Swap(tc.inputA, tc.inputB)
		assert.Equal(t, tc.outputA, out1)
		assert.Equal(t, tc.outputB, out2)
	}
}

func BenchmarkSwap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Swap(2, 3)
	}
}
