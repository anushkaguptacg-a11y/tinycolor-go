package tinycolor

import (
	"fmt"
	"strconv"
)

// RGBString returns a string representation of the color in rgb or rgba format
func (c *Color) RGBString() string {
	if c == nil {
		return ""
	}
	r := jsRound(c.r)
	g := jsRound(c.g)
	b := jsRound(c.b)
	if c.a == 1.0 {
		return fmt.Sprintf("rgb(%.0f, %.0f, %.0f)", r, g, b)
	}
	return fmt.Sprintf("rgba(%.0f, %.0f, %.0f, %s)", r, g, b, formatAlpha(c.a))
}

// PercentageRGB returns the percentage representation of each channel
func (c *Color) PercentageRGB() PercentageRGB {
	if c == nil {
		return PercentageRGB{}
	}
	return PercentageRGB{
		R: fmt.Sprintf("%.0f%%", jsRound(bound01(c.r, 255)*100.0)),
		G: fmt.Sprintf("%.0f%%", jsRound(bound01(c.g, 255)*100.0)),
		B: fmt.Sprintf("%.0f%%", jsRound(bound01(c.b, 255)*100.0)),
		A: c.a,
	}
}

// PercentageRGBString returns a percentage representation string
func (c *Color) PercentageRGBString() string {
	if c == nil {
		return ""
	}
	r := jsRound(bound01(c.r, 255) * 100.0)
	g := jsRound(bound01(c.g, 255) * 100.0)
	b := jsRound(bound01(c.b, 255) * 100.0)
	if c.a == 1.0 {
		return fmt.Sprintf("rgb(%.0f%%, %.0f%%, %.0f%%)", r, g, b)
	}
	return fmt.Sprintf("rgba(%.0f%%, %.0f%%, %.0f%%, %s)", r, g, b, formatAlpha(c.a))
}

// Hex returns the hex string representation of the color (without # prefix)
func (c *Color) Hex(allow3Char bool) string {
	if c == nil {
		return ""
	}
	return rgbToHex(c.r, c.g, c.b, allow3Char)
}

// HexString returns the hex representation with # prefix
func (c *Color) HexString(allow3Char bool) string {
	if c == nil {
		return ""
	}
	return "#" + c.Hex(allow3Char)
}

// Hex8 returns the hex8 representation of the color (without # prefix)
func (c *Color) Hex8(allow4Char bool) string {
	if c == nil {
		return ""
	}
	return rgbaToHex(c.r, c.g, c.b, c.a, allow4Char)
}

// Hex8String returns the hex8 representation with # prefix
func (c *Color) Hex8String(allow4Char bool) string {
	if c == nil {
		return ""
	}
	return "#" + c.Hex8(allow4Char)
}

// HSLString returns the HSL/HSLA string representation of the color
func (c *Color) HSLString() string {
	if c == nil {
		return ""
	}
	hsl := rgbToHsl(c.r, c.g, c.b)
	h := jsRound(hsl.H * 360.0)
	s := jsRound(hsl.S * 100.0)
	l := jsRound(hsl.L * 100.0)
	if c.a == 1.0 {
		return fmt.Sprintf("hsl(%.0f, %.0f%%, %.0f%%)", h, s, l)
	}
	return fmt.Sprintf("hsla(%.0f, %.0f%%, %.0f%%, %s)", h, s, l, formatAlpha(c.a))
}

// HSVString returns the HSV/HSVA string representation of the color
func (c *Color) HSVString() string {
	if c == nil {
		return ""
	}
	hsv := rgbToHsv(c.r, c.g, c.b)
	h := jsRound(hsv.H * 360.0)
	s := jsRound(hsv.S * 100.0)
	v := jsRound(hsv.V * 100.0)
	if c.a == 1.0 {
		return fmt.Sprintf("hsv(%.0f, %.0f%%, %.0f%%)", h, s, v)
	}
	return fmt.Sprintf("hsva(%.0f, %.0f%%, %.0f%%, %s)", h, s, v, formatAlpha(c.a))
}

// Name returns the matching named CSS color name
func (c *Color) Name() (string, bool) {
	if c == nil {
		return "", false
	}
	if c.a == 0 {
		return "transparent", true
	}
	if c.a < 1 {
		return "", false
	}
	h := rgbToHex(c.r, c.g, c.b, true)
	if name, found := HexNames[h]; found {
		return name, true
	}
	return "", false
}

// Filter returns Internet Explorer gradient filter syntax
func (c *Color) Filter(secondColor ...interface{}) string {
	if c == nil {
		return ""
	}
	hex8String := "#" + rgbaToArgbHex(c.r, c.g, c.b, c.a)
	secondHex8String := hex8String
	gradientType := ""
	if c.gradientType != "" {
		gradientType = "GradientType = 1, "
	}
	if len(secondColor) > 0 && secondColor[0] != nil {
		s := Parse(secondColor[0])
		if s != nil {
			secondHex8String = "#" + rgbaToArgbHex(s.r, s.g, s.b, s.a)
		}
	}
	return fmt.Sprintf("progid:DXImageTransform.Microsoft.gradient(%sstartColorstr=%s,endColorstr=%s)", gradientType, hex8String, secondHex8String)
}

// String returns formatted representation of the color according to target or format option
func (c *Color) String(format ...string) string {
	if c == nil {
		return ""
	}
	formatStr := ""
	formatSet := false
	if len(format) > 0 && format[0] != "" {
		formatStr = format[0]
		formatSet = true
	} else {
		formatStr = c.format
	}

	hasAlpha := c.a < 1.0 && c.a >= 0.0
	needsAlphaFormat := !formatSet && hasAlpha && (formatStr == "hex" || formatStr == "hex6" || formatStr == "hex3" || formatStr == "hex4" || formatStr == "hex8" || formatStr == "name")

	if needsAlphaFormat {
		if formatStr == "name" && c.a == 0 {
			if name, found := c.Name(); found {
				return name
			}
		}
		return c.RGBString()
	}

	var formattedString string
	foundFormat := false

	switch formatStr {
	case "rgb":
		formattedString = c.RGBString()
		foundFormat = true
	case "prgb":
		formattedString = c.PercentageRGBString()
		foundFormat = true
	case "hex", "hex6":
		formattedString = c.HexString(false)
		foundFormat = true
	case "hex3":
		formattedString = c.HexString(true)
		foundFormat = true
	case "hex4":
		formattedString = c.Hex8String(true)
		foundFormat = true
	case "hex8":
		formattedString = c.Hex8String(false)
		foundFormat = true
	case "name":
		if name, found := c.Name(); found {
			formattedString = name
			foundFormat = true
		}
	case "hsl":
		formattedString = c.HSLString()
		foundFormat = true
	case "hsv":
		formattedString = c.HSVString()
		foundFormat = true
	}

	if foundFormat {
		return formattedString
	}
	return c.HexString(false)
}

// formatAlpha formats the alpha float64 to match JS behavior
func formatAlpha(a float64) string {
	rounded := jsRound(100.0*a) / 100.0
	return strconv.FormatFloat(rounded, 'f', -1, 64)
}
