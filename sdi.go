// Function to calculate stand density index 
// from quadratic mean diameter and tree per acre
// by David R. Larsen, September 9, 2013
// Creative Commons,  http://creativecommons.org/licenses/by-nc/3.0/us/
package main

import (
	"fmt"
	"math"
)

// Function to calculate a quadratic mean 
func Sdi( tpa float64, qmd float64, unittype string) float64 {
	if unittype == "imperial" {
		return tpa * math.Pow(( qmd / 10.0 ),1.605)
	}else if unittype == "metric" {
		return tpa * math.Pow(( qmd / 25.4 ),1.605)
	}
	return 0.0
}

func main() {
	fmt.Println("sdi =", Sdi(200.0, 12.3, "imperial") )
	fmt.Println("sdi =", Sdi(494.1, 31.0, "metric") )
	fmt.Println("sdi =", Sdi(200.0, 1.0, "cunits") )
}	
