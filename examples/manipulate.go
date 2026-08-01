package main

import (
	"fmt"
	"tinycolor"
)

func main() {
	// Chainable in-place manipulations
	c := tinycolor.Parse("#6699cc")
	c.Lighten(10).Saturate(20).Spin(30)
	fmt.Println("Chained Manipulation Result:", c.HexString(false)) // #7d7de8

	// Complement (returns a new instance)
	cOrig := tinycolor.Parse("red")
	cComp := cOrig.Complement()
	fmt.Println("Original:", cOrig.HexString(false))   // #ff0000
	fmt.Println("Complement:", cComp.HexString(false)) // #00ffff

	// Triad combinations (returns new instances)
	triad := cOrig.Triad()
	fmt.Println("Triad colors:")
	for i, color := range triad {
		fmt.Printf("  %d: %s\n", i+1, color.HexString(false))
	}
}
