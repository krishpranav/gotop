package utils_test

import (
	"testing"
	"time"

	"github.com/pesos/grofer/src/utils"
)

func TestRoundValues(t *testing.T) {
	tests := []struct {
		expectedUnit        string
		input               []float64
		expectedRoundedVals []float64
		inBytes             bool
	}{
		{
			expectedUnit:        " ",
			input:               []float64{999, 895},
			expectedRoundedVals: []float64{999, 895},
			inBytes:             false,
		},
		{
			expectedUnit:        " per thousand ",
			input:               []float64{100000, 1000},
			expectedRoundedVals: []float64{100, 1},
			inBytes:             false,
		},
		{
			expectedUnit:        " per million ",
			input:               []float64{10000000, 1000},
			expectedRoundedVals: []float64{10, 0},
			inBytes:             false,
		},
		{
			expectedUnit:        " per billion ",
			input:               []float64{100000000, 100000000000},
			expectedRoundedVals: []float64{0.1, 100},
			inBytes:             false,
		},
		{
			expectedUnit:        " per trillion ",
			input:               []float64{100000000000, 10000000000000},
			expectedRoundedVals: []float64{0.1, 10},
			inBytes:             false,
		},
		{
			expectedUnit:        " per quadrillion ",
			input:               []float64{100000000000000, 10000000000000000},
			expectedRoundedVals: []float64{0.1, 10},
			inBytes:             false,
		},
		{
			expectedUnit:        " B ",
			input:               []float64{999, 895},
			expectedRoundedVals: []float64{999, 895},
			inBytes:             true,
		},
		{
			expectedUnit:        " kB ",
			input:               []float64{100000, 1000},
			expectedRoundedVals: []float64{100, 1},
			inBytes:             true,
		},
		{
			expectedUnit:        " mB ",
			input:               []float64{10000000, 1000},
			expectedRoundedVals: []float64{10, 0},
			inBytes:             true,
		},
		{
			expectedUnit:        " gB ",
			input:               []float64{100000000, 100000000000},
			expectedRoundedVals: []float64{0.1, 100},
			inBytes:             true,
		},
		{
			expectedUnit:        " tB ",
			input:               []float64{100000000000, 10000000000000},
			expectedRoundedVals: []float64{0.1, 10},
			inBytes:             true,
		},
		{
			expectedUnit:        " pB ",
			input:               []float64{100000000000000, 10000000000000000},
			expectedRoundedVals: []float64{0.1, 10},
			inBytes:             true,
		},
	}

	for _, test := range tests {
		testRoundedVals, testUnit := utils.RoundValues(test.input[0], test.input[1], test.inBytes)
		utils.Equals(t, test.expectedRoundedVals, testRoundedVals)
		utils.Equals(t, test.expectedUnit, testUnit)
	}
}
