// Function to calculate a simple taper equation
// returning the diameter athe height of interest h.
// by David R. Larsen, October 11, 2012
// Creative Commons,  http://creativecommons.org/licenses/by-nc/3.0/us/
package main

import (
	"fmt"
)

func simpleTaper(h float64, dbh float64, ht float64, htcb float64, stumpR float64, stemR float64, bh float64) float64 {
	diameterCrownBase := dbh + stemR*(htcb-bh)
	crownLength := ht - htcb
	topRate := diameterCrownBase / crownLength
	d := 0.0

	if h < bh {
		d = dbh + stumpR*(h-bh)
	} else if (h >= bh) && (h < htcb) {
		d = dbh + stemR*(h-bh)
	} else {
		d = (ht - h) * topRate
	}
	return d

}

func main() {
	fmt.Println("diameter =", simpleTaper(12, 16, 65, 30, -0.0713, -0.0169, 4.5))
}
