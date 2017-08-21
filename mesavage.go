// Function to calculate Mesavage and Girard tree Volumes
// from dbh and merchantable tree height, and Girard form class
// by David R. Larsen, June 21, 2015
// Creative Commons,  http://creativecommons.org/licenses/by-nc/3.0/us/
package main

import (
	"fmt"
	"math"
)

// Function to calculate a quadratic mean
func Mesavage(dbh float64, mht float64, voltype string, girard float64) float64 {
	var a, b, c []float64
	L := mht / 16.0
	cor := (1.0 + ((girard - 78.0) * 0.03))

	if voltype == "Int1/4" {
		a = []float64{-13.35212, 9.58615, 1.52968}
		b = []float64{1.79620, -2.59995, -0.27465}
		c = []float64{0.04482, 0.45997, -0.00961}
	} else if voltype == "Scribner" {
		a = []float64{-22.50365, 17.53508, -0.59242}
		b = []float64{3.02888, -4.34381, -0.02302}
		c = []float64{-0.01969, 0.51593, -0.02035}
	} else if voltype == "Doyle" {
		a = []float64{-29.37337, 41.51275, 0.55743}
		b = []float64{2.78043, -8.77272, -0.04516}
		c = []float64{0.04177, 0.59042, -0.01578}
	}
	v1 := (a[0] + a[1]*L + a[2]*math.Pow(L, 2))
	v2 := (b[0] + b[1]*L + b[2]*math.Pow(L, 2)) * dbh
	v3 := (c[0] + c[1]*L + c[2]*math.Pow(L, 2)) * math.Pow(dbh, 2)
	volume := (v1 + v2 + v3) * cor
	return volume
}

func main() {
	fmt.Println("Int 1/4 (78)=", Mesavage(24.0, 40.0, "Int1/4", 78.0))
	fmt.Println("Int 1/4 (82)=", Mesavage(24.0, 40.0, "Int1/4", 82.0))
	fmt.Println("Scribner (78)=", Mesavage(24.0, 40.0, "Scribner", 78.0))
	fmt.Println("Scribner (82)=", Mesavage(24.0, 40.0, "Scribner", 82.0))
	fmt.Println("Doyle (78)=", Mesavage(24.0, 40.0, "Doyle", 78.0))
	fmt.Println("Doyle (82)=", Mesavage(24.0, 40.0, "Doyle", 82.0))
}
