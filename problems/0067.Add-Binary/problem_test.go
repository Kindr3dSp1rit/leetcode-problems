package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	input struct {
		a string
		b string
	}

	output string
)

func Test_AddBinary(t *testing.T) {

	testCases := []struct {
		in  input
		out output
	}{
		{
			in: input{
				a: "11",
				b: "1",
			},
			out: "100",
		},
		{
			in: input{
				a: "1010",
				b: "1011",
			},
			out: "10101",
		},
		{
			in: input{
				a: "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
				b: "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011",
			},
			out: "110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000",
		},
	}

	for i, test := range testCases {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, string(test.out), addBinary(test.in.a, test.in.b))
		})
	}
}
