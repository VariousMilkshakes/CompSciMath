package algo

import "math"

// Hcf Finds the highest common factor between a and b
func Hcf(a, b int) int {
	var c int
	if b < a {
		c = b
		b = a
		a = c
	}

	r := math.Mod(float64(b), float64(a))

	if r == 0 {
		return int(a)
	}

	return Hcf(a, int(r))
}

// func Naive(a, b int) int {
//
// }
//
// func breakDown(value int) (factor, result int) {
//
// }

//110180800
