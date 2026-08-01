# TinyColor-Go

![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green.svg)
[![CI](https://github.com/anushkaguptacg-a11y/tinycolor-go/actions/workflows/go.yml/badge.svg)](https://github.com/anushkaguptacg-a11y/tinycolor-go/actions/workflows/go.yml)

A feature-complete Go port of the original JavaScript **TinyColor** library with strict behavioral compatibility, comprehensive testing, WCAG accessibility utilities, and GitHub Actions CI.

This project preserves TinyColor's parsing quirks, conversion logic, formatting behavior, manipulation semantics, and edge cases while providing an idiomatic Go API.

---

## Features

- ✅ Complete TinyColor API implemented in Go
- ✅ CSS color parsing
  - RGB / RGBA
  - HSL / HSLA
  - HSV / HSVA
  - Hex (#RGB, #RGBA, #RRGGBB, #RRGGBBAA)
  - CSS named colors
  - `transparent`
- ✅ RGB ↔ HSL ↔ HSV conversions
- ✅ Color formatting
  - Hex / Hex8
  - RGB / Percentage RGB
  - HSL / HSV
  - CSS color names
- ✅ Color manipulation
  - Lighten
  - Darken
  - Brighten
  - Saturate
  - Desaturate
  - Greyscale
  - Spin
  - Mix
- ✅ Color combinations
  - Complement
  - Triad
  - Tetrad
  - Analogous
  - Split Complement
  - Monochromatic
- ✅ WCAG 2.0 Readability utilities
- ✅ GitHub Actions CI
- ✅ JavaScript-compatible rounding behavior
- ✅ Extensive unit tests and compatibility regression tests

---

## Installation

```bash
go get github.com/anushkaguptacg-a11y/tinycolor-go
```

---

# Usage

## Parse Colors

```go
package main

import (
	"fmt"

	"github.com/anushkaguptacg-a11y/tinycolor-go"
)

func main() {
	color := tinycolor.Parse("#3498db")

	fmt.Println(color.HexString(false))
	fmt.Println(color.RGBString())
	fmt.Println(color.HSLString())
}
```

---

## Color Manipulation

```go
c := tinycolor.Parse("#6699cc")

c.Lighten(10).
	Saturate(20).
	Spin(30)

fmt.Println(c.HexString(false))

// Output:
// #7d7de8
```

---

## Color Combinations

```go
c := tinycolor.Parse("red")

triad := c.Triad()

for _, color := range triad {
	fmt.Println(color.HexString(false))
}
```

Output

```
#ff0000
#00ff00
#0000ff
```

---

## WCAG Readability

```go
ratio := tinycolor.Readability("#000", "#fff")

fmt.Println(ratio)

// Output:
// 21
```

Checking accessibility:

```go
ok := tinycolor.IsReadable(
	"#ff0088",
	"#5c1a72",
	tinycolor.WCAG2Opts{
		Level: "AA",
		Size: "large",
	},
)

fmt.Println(ok)
```

Finding the most readable color:

```go
best := tinycolor.MostReadable(
	"#123",
	[]interface{}{
		"#124",
		"#125",
		"#fff",
	},
	tinycolor.WCAG2Opts{
		IncludeFallbackColors: true,
	},
)

fmt.Println(best.HexString(false))
```

---

## Formatting

```go
c := tinycolor.Parse("rgba(255,0,0,0.5)")

fmt.Println(c.RGBString())
fmt.Println(c.PercentageRGBString())
fmt.Println(c.HexString(false))
fmt.Println(c.Hex8String(true))
fmt.Println(c.HSLString())
fmt.Println(c.Filter())
```

Produces

```
rgba(255, 0, 0, 0.5)
rgba(100%, 0%, 0%, 0.5)
#ff0000
#f008
hsla(0, 100%, 50%, 0.5)
progid:DXImageTransform.Microsoft.gradient(...)
```

---

# Testing

Run all tests

```bash
go test ./...
```

Run verbose tests

```bash
go test -v ./...
```

Run benchmarks

```bash
go test -bench=. -benchmem
```

Run static analysis

```bash
go vet ./...
```

---

# Benchmarks

Benchmarks executed on a **13th Gen Intel Core i7-1360P**.

| Benchmark                  | Performance |
| -------------------------- | ----------: |
| Parse                      | ~18.9 µs/op |
| RGB ↔ HSL / HSV Conversion |  ~172 ns/op |
| Formatting                 |  ~1.2 µs/op |

---

# Project Highlights

- Complete TinyColor API
- Strict JavaScript behavioral compatibility
- Extensive compatibility regression tests
- JavaScript rounding emulation (`Math.round`)
- WCAG accessibility utilities
- GitHub Actions CI
- MIT Licensed

---

# Credits

This project is a Go port of the original **TinyColor** JavaScript library.

Original project:

https://github.com/bgrins/TinyColor

Full credit goes to **Brian Grinstead** and all TinyColor contributors for the original implementation and algorithms.

---

# License

This project is licensed under the MIT License.

See the [LICENSE](LICENSE) file for details.
