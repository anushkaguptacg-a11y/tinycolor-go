package main

import (
	"fmt"
	"tinycolor"
)

func main() {
	// Parse from hex strings (3, 4, 6, and 8 chars)
	c1 := tinycolor.Parse("#f00")
	c2 := tinycolor.Parse("#f00f")
	c3 := tinycolor.Parse("#ff0000")
	c4 := tinycolor.Parse("#ff0000ff")

	fmt.Println("c1 Hex:", c1.HexString(false)) // #ff0000
	fmt.Println("c2 Hex8:", c2.Hex8String(false)) // #ff0000ff
	fmt.Println("c3 Hex:", c3.HexString(false)) // #ff0000
	fmt.Println("c4 Hex8:", c4.Hex8String(false)) // #ff0000ff

	// Parse from CSS color names
	c5 := tinycolor.Parse("saddlebrown")
	fmt.Println("c5 Name:", c5.Hex8String(false)) // #8b4513ff

	// Parse from HSL, HSV, RGB structs
	c6 := tinycolor.Parse(tinycolor.HSL{H: 120, S: 0.5, L: 0.5})
	c7 := tinycolor.Parse(tinycolor.HSV{H: 240, S: 0.6, V: 0.8})
	c8 := tinycolor.Parse(tinycolor.RGB{R: 128, G: 128, B: 128})

	fmt.Println("c6 RGB:", c6.RGBString()) // rgb(64, 191, 64)
	fmt.Println("c7 Hex:", c7.HexString(false)) // #5151cc
	fmt.Println("c8 HSL:", c8.HSLString()) // hsl(0, 0%, 50%)
}
