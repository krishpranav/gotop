package utils

import (
	"math"
	"time"
)

var (
	K = math.Pow(10, 3)
	M = math.Pow(10, 6)
	G = math.Pow(10, 9)
	T = math.Pow(10, 12)
	Q = math.Pow(10, 15)
)

func roundOffNearestTen(num float64, divisor float64) float64 {
	x := num / divisor
	return math.Round(x*10) / 10
}

func RoundValues(num1, num2 float64, inBytes bool) ([]float64, string) {
	nums := []float64{}
	var units string
	var n float64
	if num1 > num2 {
		n = num1
	} else {
		n = num2
	}

	switch {
	case n < K:
		nums = append(nums, num1)
		nums = append(nums, num2)
		units = " "

	case n < M:
		nums = append(nums, roundOffNearestTen(num1, K))
		nums = append(nums, roundOffNearestTen(num2, K))
		units = " per thousand "

	case n < G:
		nums = append(nums, roundOffNearestTen(num1, M))
		nums = append(nums, roundOffNearestTen(num2, M))
		units = " per million "

	case n < T:
		nums = append(nums, roundOffNearestTen(num1, G))
		nums = append(nums, roundOffNearestTen(num2, G))
		units = " per billion "

	case n < Q:
		nums = append(nums, roundOffNearestTen(num1, T))
		nums = append(nums, roundOffNearestTen(num2, T))
		units = " per trillion "

	case n >= Q:
		nums = append(nums, roundOffNearestTen(num1, Q))
		nums = append(nums, roundOffNearestTen(num2, Q))
		units = " per quadrillion "
	}

	if inBytes {
		switch units {
		case " ":
			units = " B "
		case " per thousand ":
			units = " kB "
		case " per million ":
			units = " mB "
		case " per billion ":
			units = " gB "
		case " per trillion ":
			units = " tB "
		case " per quadrillion ":
			units = " pB "
		}
	}

	return nums, units

}
