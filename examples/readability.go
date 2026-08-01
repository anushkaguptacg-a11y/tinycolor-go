package main

import (
	"fmt"
	"tinycolor"
)

func main() {
	// Readability score calculation (Contrast ratio)
	ratio := tinycolor.Readability("#000", "#fff")
	fmt.Println("Contrast Ratio (Black vs White):", ratio) // 21.0

	// Check if two colors are readable under WCAG guidelines
	c1 := "#ff0088"
	c2 := "#5c1a72"
	fmt.Println("Readable (AA, small text):", tinycolor.IsReadable(c1, c2)) // false
	fmt.Println("Readable (AA, large text):", tinycolor.IsReadable(c1, c2, tinycolor.WCAG2Opts{
		Level: "AA",
		Size:  "large",
	})) // true

	// Choose the most readable color from a list
	candidates := []interface{}{"#124", "#125", "#fff"}
	best := tinycolor.MostReadable("#123", candidates, tinycolor.WCAG2Opts{
		IncludeFallbackColors: true,
	})
	fmt.Println("Most Readable Color:", best.HexString(false)) // #ffffff
}
