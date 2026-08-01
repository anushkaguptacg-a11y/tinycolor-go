package tinycolor

import (
	"math"
	"strconv"
	"strings"
)

type rgbConverted struct {
	R, G, B float64
}

type hslConverted struct {
	H, S, L float64
}

type hsvConverted struct {
	H, S, V float64
}

// hslToRgb converts an HSL color value to RGB.
// Assumes h in [0, 360], s and l in [0, 100] (or [0, 1]).
// Returns R, G, B in the set [0, 255].
func hslToRgb(h, s, l interface{}) rgbConverted {
	hVal := bound01(h, 360)
	sVal := bound01(s, 100)
	lVal := bound01(l, 100)

	hue2rgb := func(p, q, t float64) float64 {
		if t < 0 {
			t += 1.0
		}
		if t > 1 {
			t -= 1.0
		}
		if t < 1.0/6.0 {
			return p + (q-p)*6.0*t
		}
		if t < 1.0/2.0 {
			return q
		}
		if t < 2.0/3.0 {
			return p + (q-p)*(2.0/3.0-t)*6.0
		}
		return p
	}

	var r, g, b float64
	if sVal == 0 {
		r = lVal
		g = lVal
		b = lVal
	} else {
		var q float64
		if lVal < 0.5 {
			q = lVal * (1.0 + sVal)
		} else {
			q = lVal + sVal - lVal*sVal
		}
		p := 2.0*lVal - q
		r = hue2rgb(p, q, hVal+1.0/3.0)
		g = hue2rgb(p, q, hVal)
		b = hue2rgb(p, q, hVal-1.0/3.0)
	}

	return rgbConverted{
		R: r * 255.0,
		G: g * 255.0,
		B: b * 255.0,
	}
}

// hsvToRgb converts an HSV color value to RGB.
// Assumes h in [0, 360], s and v in [0, 100] (or [0, 1]).
// Returns R, G, B in the set [0, 255].
func hsvToRgb(h, s, v interface{}) rgbConverted {
	hVal := bound01(h, 360) * 6.0
	sVal := bound01(s, 100)
	vVal := bound01(v, 100)

	i := math.Floor(hVal)
	f := hVal - i
	p := vVal * (1.0 - sVal)
	q := vVal * (1.0 - f*sVal)
	t := vVal * (1.0 - (1.0-f)*sVal)
	mod := int(i) % 6

	var r, g, b float64
	switch mod {
	case 0:
		r, g, b = vVal, t, p
	case 1:
		r, g, b = q, vVal, p
	case 2:
		r, g, b = p, vVal, t
	case 3:
		r, g, b = p, q, vVal
	case 4:
		r, g, b = t, p, vVal
	case 5:
		r, g, b = vVal, p, q
	}

	return rgbConverted{
		R: r * 255.0,
		G: g * 255.0,
		B: b * 255.0,
	}
}

// rgbToHsl converts an RGB color value to HSL.
// Assumes r, g, and b are contained in the set [0, 255] or [0, 1].
// Returns h, s, l in [0, 1].
func rgbToHsl(r, g, b float64) hslConverted {
	rVal := bound01(r, 255)
	gVal := bound01(g, 255)
	bVal := bound01(b, 255)

	max := math.Max(rVal, math.Max(gVal, bVal))
	min := math.Min(rVal, math.Min(gVal, bVal))
	var h, s float64
	l := (max + min) / 2.0

	if max == min {
		h = 0.0
		s = 0.0 // achromatic
	} else {
		d := max - min
		if l > 0.5 {
			s = d / (2.0 - max - min)
		} else {
			s = d / (max + min)
		}
		switch max {
		case rVal:
			offset := 0.0
			if gVal < bVal {
				offset = 6.0
			}
			h = (gVal-bVal)/d + offset
		case gVal:
			h = (bVal-rVal)/d + 2.0
		case bVal:
			h = (rVal-gVal)/d + 4.0
		}
		h /= 6.0
	}

	return hslConverted{H: h, S: s, L: l}
}

// rgbToHsv converts an RGB color value to HSV.
// Assumes r, g, and b are contained in the set [0, 255] or [0, 1].
// Returns h, s, v in [0, 1].
func rgbToHsv(r, g, b float64) hsvConverted {
	rVal := bound01(r, 255)
	gVal := bound01(g, 255)
	bVal := bound01(b, 255)

	max := math.Max(rVal, math.Max(gVal, bVal))
	min := math.Min(rVal, math.Min(gVal, bVal))
	var h, s float64
	v := max
	d := max - min

	if max == 0 {
		s = 0.0
	} else {
		s = d / max
	}

	if max == min {
		h = 0.0 // achromatic
	} else {
		switch max {
		case rVal:
			offset := 0.0
			if gVal < bVal {
				offset = 6.0
			}
			h = (gVal-bVal)/d + offset
		case gVal:
			h = (bVal-rVal)/d + 2.0
		case bVal:
			h = (rVal-gVal)/d + 4.0
		}
		h /= 6.0
	}

	return hsvConverted{H: h, S: s, V: v}
}

// rgbToHex converts an RGB color to hex
func rgbToHex(r, g, b float64, allow3Char bool) string {
	hex := []string{
		pad2(strconv.FormatInt(int64(jsRound(r)), 16)),
		pad2(strconv.FormatInt(int64(jsRound(g)), 16)),
		pad2(strconv.FormatInt(int64(jsRound(b)), 16)),
	}

	if allow3Char && hex[0][0] == hex[0][1] && hex[1][0] == hex[1][1] && hex[2][0] == hex[2][1] {
		return string(hex[0][0]) + string(hex[1][0]) + string(hex[2][0])
	}
	return strings.Join(hex, "")
}

// rgbaToHex converts an RGBA color to 8-character hex
func rgbaToHex(r, g, b, a float64, allow4Char bool) string {
	hex := []string{
		pad2(strconv.FormatInt(int64(jsRound(r)), 16)),
		pad2(strconv.FormatInt(int64(jsRound(g)), 16)),
		pad2(strconv.FormatInt(int64(jsRound(b)), 16)),
		pad2(convertDecimalToHex(strconv.FormatFloat(a, 'f', -1, 64))),
	}

	if allow4Char && hex[0][0] == hex[0][1] && hex[1][0] == hex[1][1] && hex[2][0] == hex[2][1] && hex[3][0] == hex[3][1] {
		return string(hex[0][0]) + string(hex[1][0]) + string(hex[2][0]) + string(hex[3][0])
	}
	return strings.Join(hex, "")
}

// rgbaToArgbHex converts an RGBA color to an ARGB Hex8 string (required for toFilter)
func rgbaToArgbHex(r, g, b, a float64) string {
	hex := []string{
		pad2(convertDecimalToHex(strconv.FormatFloat(a, 'f', -1, 64))),
		pad2(strconv.FormatInt(int64(jsRound(r)), 16)),
		pad2(strconv.FormatInt(int64(jsRound(g)), 16)),
		pad2(strconv.FormatInt(int64(jsRound(b)), 16)),
	}
	return strings.Join(hex, "")
}
