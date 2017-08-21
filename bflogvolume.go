// Function to calculate Board foot volume
// from small end diameter and log length
// by David R. Larsen, October 11, 2012
// Creative Commons,  http://creativecommons.org/licenses/by-nc/3.0/us/
package main

import (
	"fmt"
	"math"
)

// Function to calculate the Doyle Board foot volume
// from small end diameter and log length
func doyle(sdia float64, length float64) float64 {
	value := math.Pow(((sdia-4.0)/4.0), 2.0) * length
	return value
}

// Function to calculate the Scribner Board foot volume
// from small end diameter and log length
func scribner(sdia float64, length float64) float64 {
	value := (0.79*math.Pow(sdia, 2.0) - 2*sdia - 4) * (length / 16)
	return value
}

// Function to calculate the International Board foot volume
// from small end diameter and log length
func intVolume(sdia float64, length float64) float64 {
	var value float64
	if length == 4.0 {
		value = 0.22*math.Pow(sdia, 2.0) - 0.71*sdia
	} else if length == 8.0 {
		value = 0.44*math.Pow(sdia, 2.0) - 1.20*sdia - 0.30
	} else if length == 12.0 {
		value = 0.66*math.Pow(sdia, 2.0) - 1.47*sdia - 0.79
	} else if length == 16.0 {
		value = 0.88*math.Pow(sdia, 2.0) - 1.52*sdia - 1.36
	} else if length == 20.0 {
		value = 1.10*math.Pow(sdia, 2.0) - 1.35*sdia - 1.90
	} else if length == 24.0 {
		value = 1.10*math.Pow(sdia, 2.0) - 1.35*sdia - 1.90 + 0.22*math.Pow(sdia, 2.0) - 0.71*sdia
	} else if length == 28.0 {
		value = 1.10*math.Pow(sdia, 2.0) - 1.35*sdia - 1.90 + 0.44*math.Pow(sdia, 2.0) - 1.20*sdia - 0.30
	} else if length == 32.0 {
		value = 1.10*math.Pow(sdia, 2.0) - 1.35*sdia - 1.90 + 0.66*math.Pow(sdia, 2.0) - 1.47*sdia - 0.79
	} else if length == 36.0 {
		value = 1.10*math.Pow(sdia, 2.0) - 1.35*sdia - 1.90 + 0.88*math.Pow(sdia, 2.0) - 1.52*sdia - 1.36
	} else if length == 40.0 {
		value = (1.10*math.Pow(sdia, 2.0) - 1.35*sdia - 1.90) * 2
	}
	return value
}

func main() {
	fmt.Println("doyle =", doyle(10, 16))
	fmt.Println("doyle =", doyle(28, 16))
	fmt.Println("scribner =", scribner(10, 16))
	fmt.Println("scribner =", scribner(28, 16))
	fmt.Println("International volume =", intVolume(10, 16))
	fmt.Println("International volume =", intVolume(28, 16))
}
