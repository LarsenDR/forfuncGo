// Function to calculate precent stocking
// from basal area  and tree per acre
// by David R. Larsen, September 12, 2013
// Creative Commons,  http://creativecommons.org/licenses/by-nc/3.0/us/
package main

import (
	"fmt"
	"math"
	"strings"
)

//  Percent stocking has only been worked out for imperial measurements

func parameters(group string) []float64 {
	value := []float64{ -0.00507, 0.01698, 0.00307 }
	if strings.Contains( group, "northern.red.oak" ) {
	   value = []float64{ 0.02476, 0.006272, 0.00469 }
	}
	return value
}

// Function to calculate a quadratic mean
func qmd(ba float64, tpa float64) float64 {
	val := math.Sqrt( ((ba / tpa) / 0.005454154) )
	return val
}

// percentStocking is a function that return the percent stocking
// using a Gingrich style tree area equation
func percentStocking(ba float64, tpa float64, group string) float64 {
	var percent float64
        dia := qmd( ba, tpa )
	amd := -0.259 + 0.973 * dia
	b := parameters( group )
	percent = (b[0] + b[1] * amd + b[2] * math.Pow(dia, 2.0)) * tpa
	if percent < 0.0 {
	    percent = 0.0
	}
	return percent
}

func main() {
	fmt.Println("percent stocking( 55, 100, \"upland.oak\") = ", percentStocking(55.0, 100.0, "upland.oak"))
	fmt.Println("percent stocking( 100, 200, \"upland.oak\") = ", percentStocking(100.0, 200.0, "upland.oak"))
	fmt.Println("percent stocking( 50, 100, \"northern.red.oak\") = ", percentStocking(50.0, 100.0, "northern.red.oak"))
	fmt.Println("percent stocking( 100, 50, \"upland.oak\") = ", percentStocking(100.0, 50.0, "upland.oak"))
	fmt.Println("percent stocking( 50, 400, \"upland.oak\") = ", percentStocking(50.0, 400.0, "upland.oak"))
	
}	
	
