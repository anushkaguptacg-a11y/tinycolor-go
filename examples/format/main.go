package main

import (
	"fmt"
	"tinycolor"
)

func main() {
	c := tinycolor.Parse("rgba(255, 0, 0, 0.5)")

	fmt.Println("RGBString:", c.RGBString())                       // rgba(255, 0, 0, 0.5)
	fmt.Println("PercentageRGBString:", c.PercentageRGBString())   // rgba(100%, 0%, 0%, 0.5)
	fmt.Println("HexString:", c.HexString(false))                  // #ff0000 (ignores alpha)
	fmt.Println("Hex8String:", c.Hex8String(true))                  // #f008 (shortened Hex8)
	fmt.Println("HSLString:", c.HSLString())                       // hsla(0, 100%, 50%, 0.5)
	fmt.Println("Filter:", c.Filter())                             // progid:DXImageTransform.Microsoft.gradient(startColorstr=#80ff0000,endColorstr=#80ff0000)
}
