# Go TinyColor

A strict, fast, and feature-complete Go port of the popular JavaScript [TinyColor](https://github.com/bgrins/TinyColor) color manipulation and parsing library. 

Developed to achieve 100% behavioral parity with the JavaScript implementation, this Go library handles all inputs, CSS formatting quirks, mathematical conversions, and color manipulations exactly like the original.

## Features

- **Strict Parser Parity**: Supports all CSS color notations (`rgb`, `rgba`, `hsl`, `hsla`, `hsv`, `hsva`, 3/4/6/8-character hex, CSS named colors, and `transparent`) with loose spacing and wrapping matches.
- **Conversion Symmetry**: RGB, HSL, and HSV conversion logic preserved exactly, maintaining raw float precision.
- **In-place Mutation**: Modifying methods (`Lighten`, `Darken`, `Saturate`, `Desaturate`, `Spin`, `Brighten`) mutate the instance and return the same pointer for chaining, matching JavaScript's mutation semantics.
- **WCAG 2.0 Readability**: Contrast ratio calculations (`Readability`), standard WCAG check combinations (`IsReadable`), and optimal color selections (`MostReadable`) with custom fallbacks and float-point tolerance.
- **JavaScript Rounding Emulation**: Built-in `jsRound` helper matching JavaScript's `Math.round` rounding logic (rounding `.5` towards positive infinity) rather than Go's default round-away-from-zero logic.

---

## Installation

```bash
go get github.com/yourusername/tinycolor
```

*(Note: Recommend checking out this project as your Go package import path.)*

---

## Usage Examples

### 1. Initialization and Parsing

Initialize colors from strings, structs, or ratios:

```go
package main

import (
	"fmt"
	"github.com/yourusername/tinycolor"
)

func main() {
	// From hex
	c1 := tinycolor.Parse("#f00")
	fmt.Println(c1.HexString(false)) // "#ff0000"

	// From CSS named color
	c2 := tinycolor.Parse("saddlebrown")
	fmt.Println(c2.Hex8String(false)) // "#8b4513ff"

	// From HSL / HSV struct
	c3 := tinycolor.Parse(tinycolor.HSL{H: 120, S: 0.5, L: 0.5})
	fmt.Println(c3.RGBString()) // "rgb(64, 191, 64)"

	// From Ratio (0.0 to 1.0 values)
	c4 := tinycolor.FromRatio(tinycolor.RGB{R: 0.5, G: 0.5, B: 0.5})
	fmt.Println(c4.HexString(false)) // "#808080"
}
```

### 2. Output Formatting

Format colors to multiple standard representations:

```go
c := tinycolor.Parse("rgba(255, 0, 0, 0.5)")

fmt.Println(c.RGBString())            // "rgba(255, 0, 0, 0.5)"
fmt.Println(c.PercentageRGBString())  // "rgba(100%, 0%, 0%, 0.5)"
fmt.Println(c.HexString(false))       // "#ff0000" (automatically ignores alpha)
fmt.Println(c.Hex8String(true))       // "#f008" (shortened Hex8)
fmt.Println(c.HSLString())            // "hsla(0, 100%, 50%, 0.5)"
fmt.Println(c.Filter())               // "progid:DXImageTransform.Microsoft.gradient(startColorstr=#80ff0000,endColorstr=#80ff0000)"
```

### 3. Modifications (Chainable)

Modifying operations change the instance in place and return the same pointer:

```go
c := tinycolor.Parse("#6699cc")

// Chained manipulations mutating the instance in place
c.Lighten(10).Saturate(20).Spin(30)
fmt.Println(c.HexString(false)) // "#7d7de8"
```

### 4. Combinations

Combinations return slices of *new* color instances without mutating the original:

```go
c := tinycolor.Parse("red")

// Complementary color
comp := c.Complement()
fmt.Println(comp.HexString(false)) // "#00ffff"
fmt.Println(c.HexString(false))    // "#ff0000" (original remains unchanged)

// Triad combination
triad := c.Triad()
for _, color := range triad {
	fmt.Println(color.HexString(false))
}
// Outputs: #ff0000, #00ff00, #0000ff
```

### 5. WCAG 2 Readability Calculations

```go
// Readability contrast ratio
ratio := tinycolor.Readability("#000", "#fff")
fmt.Println(ratio) // 21.0

// AA check for small text
fmt.Println(tinycolor.IsReadable("#ff0088", "#5c1a72")) // false

// AA check for large text
fmt.Println(tinycolor.IsReadable("#ff0088", "#5c1a72", tinycolor.WCAG2Opts{
	Level: "AA",
	Size:  "large",
})) // true

// Optimal foreground selection
best := tinycolor.MostReadable("#123", []interface{}{"#124", "#125", "#fff"}, tinycolor.WCAG2Opts{
	IncludeFallbackColors: true,
})
fmt.Println(best.HexString(false)) // "#ffffff" (falls back to white for best contrast)
```

---

## Benchmarks

High-performance benchmarks performed on a **13th Gen Intel Core i7-1360P** running Windows 11:

```
BenchmarkParse-16         	   81492	     18918 ns/op	    2953 B/op	      96 allocs/op
BenchmarkConversion-16    	 9023144	       172.2 ns/op	      16 B/op	       2 allocs/op
BenchmarkFormatting-16    	  828511	      1247 ns/op	     112 B/op	      12 allocs/op
```

---

## License

This project is licensed under the MIT License - see the `LICENSE` file for details.
