package algo

import "math"

// PrimesIn returns an array of all the primes between 1 - n
func PrimesIn(n int) []int {

	var primes []int
	// var composites []int

	for num := 2; num <= n; num++ {

		pCheck := true

		for i := 0; i < len(primes); i++ {
			if math.Mod(float64(num), float64(primes[i])) == 0 {
				pCheck = false
				break
			}
		}

		if pCheck {
			primes = append(primes, num)
		}
		// if math.Mod(float64(num), float64(2)) == 0 && num != 2 {
		// 	continue
		// }
		//
		// if math.Mod(float64(num), float64(3)) == 0 && num != 3 {
		// 	continue
		// }
		//
		// if math.Mod(float64(num), float64(5)) == 0 && num != 3 {
		// 	continue
		// }
		//
		// if !contains(composites, num) {
		//
		// 	primes = append(primes, num)
		// 	fmt.Println(num)
		//
		// 	var composite int
		//
		// 	for multi := 1; composite <= n; multi++ {
		//
		// 		composite = num * multi
		//
		// 		composites = append(composites, composite)
		//
		// 	}
		// }
	}

	return primes
}

func contains(slice []int, num int) bool {

	for x := 0; x < len(slice); x++ {

		if slice[x] == num {
			return true
		}

	}

	return false
}
