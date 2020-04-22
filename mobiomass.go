// Function to calculate biomass in a tree based on data collected in Missouri.
// Data for only black, post and white oaks and hickories
// dbh is in inches
// mht is in feet
// the result is in pound per tree in the merchantable log.

// Equations from Goerndt, M. E., D. R. Larsen, C. D. Keating. 2014. Evaluation of
// aboveground biomass weight and merchantable biomass weight for four hardwood
// species in Missouri. In the proceedings of the Central Hardwood Conference
// March 10-12, 2014, Carbondale Illinois, USDA Forest Service GTR-NRS-P-142.
// Pages 234-250
//
// by David R. Larsen, Copyright January 21, 2019
// Creative Commons http://creativecommons.org/licenses/by-nc/3.0/us/
package main

import (
	"fmt"
	"math"
)

func mobiomass(dbh, mht float64, species string) float64 {
	var mobio float64
	if (species == "black") && (dbh > 9.9) && (dbh < 36.1) {
		mobio = 1.67079 + 0.04796*math.Pow(dbh, 2) + 0.81549*math.Log(math.Pow(dbh, 2)*mht)
	} else if (species == "post") && (dbh > 9.9) && (dbh < 36.1) {
		mobio = -0.50714 + 0.01655*math.Pow(dbh, 2) + 0.81549*math.Log(math.Pow(dbh, 2)*mht)
	} else if (species == "hickory") && (dbh > 9.9) && (dbh < 36.1) {
		mobio = 0.70177 + 0.05791*math.Pow(dbh, 2) + 0.60755*math.Log(math.Pow(dbh, 2)*mht)
	} else if (species == "white") && (dbh > 9.9) && (dbh < 36.1) {
		mobio = 0.61557 + 0.00373*math.Pow(dbh, 2) + 0.71159*math.Log(math.Pow(dbh, 2)*mht)
	} else {
		mobio = 0.0
	}
	return mobio
}

func main() {
	fmt.Printf("black oak, 14.3, 24.0, imperial = %f\n", mobiomass(14.3, 24.0, "black"))
	fmt.Printf("white oak, 22.6, 32.0, imperial = %f\n", mobiomass(22.6, 32.0, "white"))
	fmt.Printf("scarlet oak, 18.9, 18.0 imperial = %f\n", mobiomass(18.9, 18.0, "hickory"))
}
