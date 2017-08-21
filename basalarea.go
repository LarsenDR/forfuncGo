// Function to calculate the basal area per acre
// from diameter and weight
// by David R. Larsen, November 20, 2013
// Creative Commons,  http://creativecommons.org/licenses/by-nc/3.0/us/

package main

import (
	"fmt"
)

func basalarea(dia []float64, wt []float64, unittype string) (ba float64) {
	cst := 0.005454154
	if unittype == "metric" {
		cst = 0.00007856
	}

	for i := range dia {
		ba = ba + cst*dia[i]*dia[i]*wt[i]
	}
	return
}

func main() {
	dia := []float64{8, 6, 8, 5, 4, 6, 7}
	wt := []float64{10, 10, 10, 10, 10, 10, 10}
	wt2 := []float64{25, 25, 25, 25, 25, 25, 25}
	fmt.Println("basalarea imperial =", basalarea(dia, wt, "imperial"))
	fmt.Println("basalarea metric =", basalarea(dia, wt2, "metric"))
}
