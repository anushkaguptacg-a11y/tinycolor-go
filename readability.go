package tinycolor

import (
	"strings"
)

// WCAG2Opts defines level, size, and fallback configurations for readability checks.
type WCAG2Opts struct {
	Level                 string // "AA" or "AAA", defaults to "AA"
	Size                  string // "small" or "large", defaults to "small"
	IncludeFallbackColors bool   // defaults to false
}

// validateWCAG2Parms normalizes the WCAG2Opts parameters
func validateWCAG2Parms(opts ...WCAG2Opts) WCAG2Opts {
	opt := WCAG2Opts{
		Level: "AA",
		Size:  "small",
	}
	if len(opts) > 0 {
		opt.Level = opts[0].Level
		opt.Size = opts[0].Size
		opt.IncludeFallbackColors = opts[0].IncludeFallbackColors
	}

	opt.Level = strings.ToUpper(opt.Level)
	opt.Size = strings.ToLower(opt.Size)

	if opt.Level != "AA" && opt.Level != "AAA" {
		opt.Level = "AA"
	}
	if opt.Size != "small" && opt.Size != "large" {
		opt.Size = "small"
	}
	return opt
}

// Readability calculates the contrast ratio between two colors (returns 1.0 to 21.0)
func Readability(color1, color2 interface{}) float64 {
	c1 := Parse(color1)
	c2 := Parse(color2)
	if c1 == nil || c2 == nil || !c1.ok || !c2.ok {
		return 1.0
	}
	l1 := c1.GetLuminance()
	l2 := c2.GetLuminance()

	var maxL, minL float64
	if l1 > l2 {
		maxL = l1
		minL = l2
	} else {
		maxL = l2
		minL = l1
	}

	return (maxL + 0.05) / (minL + 0.05)
}

// IsReadable checks if the contrast ratio between color1 and color2 satisfies WCAG2 guidelines.
// It incorporates a floating-point tolerance of 1e-9.
func IsReadable(color1, color2 interface{}, opts ...WCAG2Opts) bool {
	readability := Readability(color1, color2)
	opt := validateWCAG2Parms(opts...)

	const epsilon = 1e-9
	out := false
	switch opt.Level + opt.Size {
	case "AAsmall", "AAAlarge":
		out = readability >= (4.5 - epsilon)
	case "AAlarge":
		out = readability >= (3.0 - epsilon)
	case "AAAsmall":
		out = readability >= (7.0 - epsilon)
	}
	return out
}

// MostReadable returns the most readable color from a list for a given base color.
// If IncludeFallbackColors is true and no candidate is readable, it returns black or white (whichever has higher contrast).
func MostReadable(baseColor interface{}, colorList []interface{}, opts ...WCAG2Opts) *Color {
	opt := validateWCAG2Parms(opts...)

	if len(colorList) == 0 {
		if opt.IncludeFallbackColors {
			return MostReadable(baseColor, []interface{}{"#fff", "#000"}, WCAG2Opts{
				Level:                 opt.Level,
				Size:                  opt.Size,
				IncludeFallbackColors: false,
			})
		}
		return nil
	}

	var bestColor *Color
	bestScore := -1.0

	for _, c := range colorList {
		score := Readability(baseColor, c)
		if score > bestScore {
			bestScore = score
			bestColor = Parse(c)
		}
	}

	if IsReadable(baseColor, bestColor, opt) || !opt.IncludeFallbackColors {
		return bestColor
	}

	// Fallback to black and white
	return MostReadable(baseColor, []interface{}{"#fff", "#000"}, WCAG2Opts{
		Level:                 opt.Level,
		Size:                  opt.Size,
		IncludeFallbackColors: false,
	})
}
