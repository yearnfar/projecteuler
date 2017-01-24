package problem1

import (
	"fmt"
	"math"
)

func init() {
	fmt.Println(problem1(1000))
}

func problem1(n float64) float64 {
	m := float64(0)

	for i := float64(0); i < n; i++ {
		if math.Mod(i, 3) == 0 || math.Mod(i, 5) == 0 {
			m += i
		}
	}

	return m
}
