// Function to calculate  cubic volume
// from small end diameter, large end diameter and log length
// by David R. Larsen, October 30, 2012
// Creative Commons,  http://creativecommons.org/licenses/by-nc/3.0/us/
package main

import (
	"fmt"
	"math"
)

func cubicvolume(sdia float64, ldia float64, length float64, equationtype string, unittype string, sameunits bool) float64 {
	var As float64
	var Al float64
	var value float64
	if unittype == "imperial" {
		if sameunits == true {
			As = math.Pi * math.Pow(sdia, 2.0)
			Al = math.Pi * math.Pow(ldia, 2.0)
		} else {
			As = math.Pi / (4.0 * 144.0) * math.Pow((sdia), 2.0)
			Al = math.Pi / (4.0 * 144.0) * math.Pow((ldia), 2.0)
		}

		if equationtype == "smalian" {
			value = length / 2.0 * (As + Al)
		} else if equationtype == "cone" {
			value = length / 3.0 * (As + math.Sqrt(As*Al) + Al)
		} else if equationtype == "neiloid" {
			value = length / 4.0 * (As + math.Cbrt(math.Pow(As, 2.0)*Al) + math.Cbrt(As*math.Pow(Al, 2.0)) + Al)
		}
	} else if unittype == "metric" {
		if sameunits == true {
			As = math.Pi * math.Pow(sdia, 2.0)
			Al = math.Pi * math.Pow(ldia, 2.0)
		} else {
			As = math.Pi / (4.0 * 10000.0) * math.Pow((sdia), 2.0)
			Al = math.Pi / (4.0 * 10000.0) * math.Pow((ldia), 2.0)
		}

		if equationtype == "smalian" {
			value = length / 2.0 * (As + Al)
		} else if equationtype == "cone" {
			value = length / 3.0 * (As + math.Sqrt(As*Al) + Al)
		} else if equationtype == "neiloid" {
			value = length / 4.0 * (As + math.Cbrt(math.Pow(As, 2.0)*Al) + math.Cbrt(As*math.Pow(Al, 2)) + Al)
		}
	} else {
		fmt.Println("\nUnknown unittype,", unittype, " options are imperial or metric\n")
	}
	return value
}

func main() {
	fmt.Println("smalian =", cubicvolume(10.0, 12.0, 16.0, "smalian", "imperial", false))
	fmt.Println("smalian =", cubicvolume(25.0, 29.0, 5.0, "smalian", "metric", false))
	fmt.Println("smalian =", cubicvolume(10.0, 12.0, 16.0, "smalian", "cunits", false))
	fmt.Println("cone =", cubicvolume(10.0, 12.0, 16.0, "cone", "imperial", false))
	fmt.Println("cone =", cubicvolume(25.0, 29.0, 5.0, "cone", "metric", false))
	fmt.Println("cone =", cubicvolume(10.0, 12.0, 16.0, "cone", "cunits", false))
	fmt.Println("neiloid =", cubicvolume(10.0, 12.0, 16.0, "neiloid", "imperial", false))
	fmt.Println("neiloid =", cubicvolume(25.0, 29.0, 5.0, "neiloid", "metric", false))
	fmt.Println("neiloid =", cubicvolume(10.0, 12.0, 16.0, "neiloid", "cunits", false))
}
