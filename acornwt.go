// Function to calculate mean acorn weight (in lb) from dbh and species and trees per acre
// by David R. Larsen, Copyright January 21, 2019
// Creative Commons http://creativecommons.org/licenses/by-nc/3.0/us/
package main

import (
	"fmt"
	"math"
)

func acornwt(dbh float64, species string) float64 {
	var awt float64
	if (species == "black") && (dbh > 9.9) && (dbh < 36.1) {
		awt = -1.9065 + 0.2973*dbh
	} else if (species == "chestnut") && (dbh > 9.9) && (dbh < 36.1) {
		awt = 0.0008271*math.Pow(dbh, 3) - 0.08157*math.Pow(dbh, 2) + 2.692*dbh - 18.85
	} else if (species == "red") && (dbh > 9.9) && (dbh < 36.1) {
		awt = 0.0004016*math.Pow(dbh, 4) - 0.0349937*math.Pow(dbh, 3) + 0.9864357*math.Pow(dbh, 2) - 9.5233885*dbh + 27.32720516
	} else if (species == "scarlet") && (dbh > 9.9) && (dbh < 36.1) {
		awt = 0.0005975*math.Pow(dbh, 4) - 0.05325*math.Pow(dbh, 3) + 1.651*math.Pow(dbh, 2) - 19.97*dbh + 84.71
	} else if (species == "white") && (dbh > 9.9) && (dbh < 36.1) {
		awt = 0.0001987*math.Pow(dbh, 4) - 0.02027*math.Pow(dbh, 3) + 0.694*math.Pow(dbh, 2) - 8.768*dbh + 37.36
	} else {
		awt = 0.0
	}
	return awt
}

func main() {
	fmt.Printf("black oak, 14.3, imperial = %f\n", acornwt(14.3, "black"))
	fmt.Printf("white oak, 22.6, imperial = %f\n", acornwt(22.6, "white"))
	fmt.Printf("scarlet oak, 18.9, imperial = %f\n", acornwt(18.9, "scarlet"))
}
