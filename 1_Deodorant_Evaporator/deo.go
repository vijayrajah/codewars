package main

import (
	"fmt"
)

func Evaporator(content float64, evapPerDay int, threshold int) int {
	// your code

	rem := content
	day := 0

	for ; ; day++ {
		rem = rem - (rem * (float64(evapPerDay) / float64(100)))

		t := (1 - (rem / content)) * 100
		// t := content - ((content - rem) / 100)

		// fmt.Println("Reduction is ", reduction, " Rem is ", rem, "t is ", t, "day is ", day)

		if int(100-t) < threshold {
			break
		}

	}

	return day + 1
}

func main() {

	if Evaporator(100, 5, 5) != 59 {
		fmt.Println("test Failed")
	} else {
		fmt.Println("test passed")
	}
}
