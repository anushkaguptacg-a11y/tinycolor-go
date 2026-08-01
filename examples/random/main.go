package main

import (
	"fmt"
	"tinycolor"
)

func main() {
	// Seed the random color generator for determinism
	tinycolor.Seed(42)

	// Generate a random color
	c := tinycolor.Random()
	fmt.Println("Random Color Format:", c.GetFormat()) // prgb
	fmt.Println("Random Color Hex:", c.HexString(false))

	// Adjust alpha independently
	c.SetAlpha(0.5)
	fmt.Println("Random Color with 50% Alpha:", c.Hex8String(false))
}
