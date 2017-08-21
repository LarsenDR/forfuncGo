// Function to calculate the density of a cubic unit of wood.
// waterdensity must in in the unit desired for output.
// by David R. Larsen, Copyright June 2, 2015
// Creative Commons http://creativecommons.org/licenses/by-nc/3.0/us/
//
// "Specific gravity G  is defined as the ratio of the density of a substance
// to the density of water pw at a specified reference temperature, typically
// 4 C (39 F), where pw  is 1.000 g cm-3(1,000 kg m-3 or 62.43 lb ft-3)." (Wood Handbook)
// Equation from Wood Handbook Chapter 4, page 10-12,  General Technical Report FPL-GTR 190
package main

import (
	"fmt"
)

func density(sg float64, mc float64, wdensity float64) float64 {
	density := wdensity * sg * (1 + (mc / 100.0))
	return density
}

func main() {
	fmt.Println("imperial =", density(0.4, 12.0, 62.43))
	fmt.Println("si =", density(0.4, 12.0, 1000.0))
}
