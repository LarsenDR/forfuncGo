// Function to calculate the quadratic mean diameter 
// from basal area and tree per acre
// by David R. Larsen, August 15, 2013
// Creative Commons,  http://creativecommons.org/licenses/by-nc/3.0/us/
package main

import (
	"fmt"
	"math"
)

// Function to calculate a quadratic mean 
func Qmd( ba float64, tpa float64, unittype string) float64 {
	if unittype == "imperial" {
		return math.Sqrt( (ba/tpa)/0.005454154 )
	}else if unittype == "metric" {
		return math.Sqrt( (ba/tpa)/0.00007854 )
	}
	return 0.0
}

func main() {
	fmt.Println("qmd =", Qmd(80.0, 200.0, "imperial") )
	fmt.Println("qmd =", Qmd(18.3, 494.1, "metric") )
	fmt.Println("qmd =", Qmd(80.0, 200.0, "cunits") )
}	
