/**
* Program to display the points generated on elliptic curve y^2 = x^3+7 with
* finite field field F17
 */
package main

import (
	"fmt"
)

func main() {
	p := 17
	for x := 0; x < p; x++ {
		for y := 0; y < p; y++ {
			if (((x*x*x)+7)-(y*y))%p == 0 {
				fmt.Printf("x = %d, y=%d \n", x, y)
			}
		}
	}
}
