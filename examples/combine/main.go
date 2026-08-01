package main

import (
	"fmt"
	"tinycolor"
)

func main() {
	c := tinycolor.Parse("red")

	// Analogous colors
	analogous := c.Analogous()
	fmt.Println("Analogous colors:")
	for _, color := range analogous {
		fmt.Println("  ", color.HexString(false))
	}

	// Monochromatic colors
	monochromatic := c.Monochromatic()
	fmt.Println("Monochromatic colors:")
	for _, color := range monochromatic {
		fmt.Println("  ", color.HexString(false))
	}

	// SplitComplement
	split := c.SplitComplement()
	fmt.Println("SplitComplement colors:")
	for _, color := range split {
		fmt.Println("  ", color.HexString(false))
	}
}
