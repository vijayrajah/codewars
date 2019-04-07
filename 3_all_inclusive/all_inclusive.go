// Input:

//     a string strng
//     an array of strings arr

// Output of function contain_all_rots(strng, arr) (or containAllRots or contain-all-rots):

//     a boolean true if all rotations of strng are included in arr (C returns 1)
//     false otherwise (C returns 0)

// contain_all_rots(
// 	"bsjq", ["bsjq", "qbsj", "sjqb", "twZNsslC", "jqbs"]) -> true

//   contain_all_rots(
// 	"Ajylvpy", ["Ajylvpy", "ylvpyAj", "jylvpyA", "lvpyAjy", "pyAjylv", "vpyAjyl", "ipywee"]) -> false)

package main

import (
	"fmt"
	// "strings"
)

func containAllRots(s string, a []string) bool {

	if s == "" {
		return true
	}

	r := genAllrot(s)
	fmt.Println("a is ", r)

	for _, i := range r {
		if !isin(i, a) {
			return false
		}
	}
	return true
}

func isin(s string, a []string) bool {

	for _, b := range a {
		if s == b {
			return true
		}
	}
	return false
}

func genAllrot(s string) []string {

	c := make([]string, 0, 0)
	l := len(s)
	for i := 0; i < l; i++ {

		var a string
		for j := 0; j < l; j++ {
			k := i + j
			if k < l {
				a = a + string(s[k])
			} else {
				a = a + string(s[k-l])
			}
		}
		c = append(c, a)
	}
	return c
}

func main() {

	if containAllRots("Ajylvpy", []string{"Ajylvpy", "ylvpyAj", "jylvpyA", "lvpyAjy", "pyAjylv", "vpyAjyl", "ipywee"}) {
		fmt.Println("failed :-(")

	} else {
		fmt.Println("Passed!!")
	}

}
