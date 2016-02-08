package main

import (
	"fmt"
	"maths/algo"
	"os"
	"strconv"
	"time"
)

func main() {
	// a := 1022
	// b := 356

	// t := time.Now().UnixNano()
	// hcf := algo.Hcf(a, b)

	// fmt.Printf("highest common factor: %d in %d nano seconds\n", hcf, dt)

	p, err := strconv.Atoi(os.Args[1])
	if err == nil {
		t := time.Now().UnixNano()
		ps := algo.PrimesIn(p)
		dt := time.Now().UnixNano() - t

		for p := 0; p < len(ps); p++ {
			fmt.Println(ps[p])
		}

		fmt.Printf("time %d \n", dt)
	}

}
