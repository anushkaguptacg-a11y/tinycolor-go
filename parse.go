package tinycolor

import (
	"math"
	"strconv"
	"strings"
)

type parsedInput struct {
	r, g, b interface{}
	h, s, l interface{}
	v       interface{}
	a       interface{}
	format  string
	ok      bool
}

// stringInputToObject parses a CSS/HTML string into raw string component tokens
func stringInputToObject(color string) parsedInput {
	color = trimLeftReg.ReplaceAllString(color, "")
	color = trimRightReg.ReplaceAllString(color, "")
	color = strings.ToLower(color)

	named := false
	if hex, found := Names[color]; found {
		color = hex
		named = true
	} else if color == "transparent" {
		return parsedInput{
			r:      "0",
			g:      "0",
			b:      "0",
			a:      "0",
			format: "name",
			ok:     true,
		}
	}

	// Match permissive rgb patterns (ignoring trailing arguments using submatch partial matching)
	if match := rgbReg.FindStringSubmatch(color); len(match) > 3 {
		format := "rgb"
		if strings.HasSuffix(match[1], "%") {
			format = "prgb"
		}
		return parsedInput{
			r:      match[1],
			g:      match[2],
			b:      match[3],
			a:      "1",
			format: format,
			ok:     true,
		}
	}
	if match := rgbaReg.FindStringSubmatch(color); len(match) > 4 {
		format := "rgb"
		if strings.HasSuffix(match[1], "%") {
			format = "prgb"
		}
		return parsedInput{
			r:      match[1],
			g:      match[2],
			b:      match[3],
			a:      match[4],
			format: format,
			ok:     true,
		}
	}
	if match := hslReg.FindStringSubmatch(color); len(match) > 3 {
		return parsedInput{
			h:      match[1],
			s:      match[2],
			l:      match[3],
			a:      "1",
			format: "hsl",
			ok:     true,
		}
	}
	if match := hslaReg.FindStringSubmatch(color); len(match) > 4 {
		return parsedInput{
			h:      match[1],
			s:      match[2],
			l:      match[3],
			a:      match[4],
			format: "hsl",
			ok:     true,
		}
	}
	if match := hsvReg.FindStringSubmatch(color); len(match) > 3 {
		return parsedInput{
			h:      match[1],
			s:      match[2],
			v:      match[3],
			a:      "1",
			format: "hsv",
			ok:     true,
		}
	}
	if match := hsvaReg.FindStringSubmatch(color); len(match) > 4 {
		return parsedInput{
			h:      match[1],
			s:      match[2],
			v:      match[3],
			a:      match[4],
			format: "hsv",
			ok:     true,
		}
	}

	// Match hex strings
	if match := hex8Reg.FindStringSubmatch(color); len(match) > 4 {
		format := "hex8"
		if named {
			format = "name"
		}
		return parsedInput{
			r:      float64(parseIntFromHex(match[1])),
			g:      float64(parseIntFromHex(match[2])),
			b:      float64(parseIntFromHex(match[3])),
			a:      convertHexToDecimal(match[4]),
			format: format,
			ok:     true,
		}
	}
	if match := hex6Reg.FindStringSubmatch(color); len(match) > 3 {
		format := "hex"
		if named {
			format = "name"
		}
		return parsedInput{
			r:      float64(parseIntFromHex(match[1])),
			g:      float64(parseIntFromHex(match[2])),
			b:      float64(parseIntFromHex(match[3])),
			a:      1.0,
			format: format,
			ok:     true,
		}
	}
	if match := hex4Reg.FindStringSubmatch(color); len(match) > 4 {
		format := "hex8"
		if named {
			format = "name"
		}
		return parsedInput{
			r:      float64(parseIntFromHex(match[1] + match[1])),
			g:      float64(parseIntFromHex(match[2] + match[2])),
			b:      float64(parseIntFromHex(match[3] + match[3])),
			a:      convertHexToDecimal(match[4] + match[4]),
			format: format,
			ok:     true,
		}
	}
	if match := hex3Reg.FindStringSubmatch(color); len(match) > 3 {
		format := "hex"
		if named {
			format = "name"
		}
		return parsedInput{
			r:      float64(parseIntFromHex(match[1] + match[1])),
			g:      float64(parseIntFromHex(match[2] + match[2])),
			b:      float64(parseIntFromHex(match[3] + match[3])),
			a:      1.0,
			format: format,
			ok:     true,
		}
	}

	return parsedInput{ok: false}
}

type rgbResult struct {
	r, g, b, a float64
	ok         bool
	format     string
}

// inputToRGB parses raw input and normalizes it to bounds-checked RGB values.
func inputToRGB(input interface{}) rgbResult {
	rgb := rgbResult{r: 0, g: 0, b: 0, a: 1.0, ok: false}

	// Helper to check valid CSS unit
	isValidCSSUnit := func(val interface{}) bool {
		if val == nil {
			return false
		}
		sVal := ""
		switch v := val.(type) {
		case string:
			sVal = v
		case float64:
			if math.IsNaN(v) {
				return false
			}
			sVal = strconv.FormatFloat(v, 'f', -1, 64)
		case float32:
			if math.IsNaN(float64(v)) {
				return false
			}
			sVal = strconv.FormatFloat(float64(v), 'f', -1, 32)
		case int:
			sVal = strconv.Itoa(v)
		case int64:
			sVal = strconv.FormatInt(v, 10)
		}
		return cssUnitReg.MatchString(sVal)
	}

	var parsed parsedInput
	isStr := false
	if s, ok := input.(string); ok {
		parsed = stringInputToObject(s)
		isStr = true
	}

	if isStr {
		if !parsed.ok {
			return rgbResult{r: 0, g: 0, b: 0, a: 1.0, ok: false}
		}
		rgb.format = parsed.format

		if parsed.format == "rgb" || parsed.format == "prgb" || parsed.format == "name" || parsed.format == "hex" || parsed.format == "hex8" {
			rgb.r = bound01(parsed.r, 255) * 255.0
			rgb.g = bound01(parsed.g, 255) * 255.0
			rgb.b = bound01(parsed.b, 255) * 255.0

			var aVal float64
			if s, ok := parsed.a.(string); ok {
				aVal = parseFloat(s)
			} else if f, ok := parsed.a.(float64); ok {
				aVal = f
			}
			rgb.a = boundAlpha(aVal)
			rgb.ok = true
			return rgbResult{
				r:      math.Min(255, math.Max(rgb.r, 0)),
				g:      math.Min(255, math.Max(rgb.g, 0)),
				b:      math.Min(255, math.Max(rgb.b, 0)),
				a:      rgb.a,
				ok:     true,
				format: rgb.format,
			}
		}

		if parsed.format == "hsl" {
			s := convertToPercentage(parsed.s)
			l := convertToPercentage(parsed.l)
			converted := hslToRgb(parsed.h, s, l)

			var aVal float64
			if s, ok := parsed.a.(string); ok {
				aVal = parseFloat(s)
			} else if f, ok := parsed.a.(float64); ok {
				aVal = f
			}
			rgb.r = converted.R
			rgb.g = converted.G
			rgb.b = converted.B
			rgb.a = boundAlpha(aVal)
			rgb.ok = true
			return rgbResult{
				r:      math.Min(255, math.Max(rgb.r, 0)),
				g:      math.Min(255, math.Max(rgb.g, 0)),
				b:      math.Min(255, math.Max(rgb.b, 0)),
				a:      rgb.a,
				ok:     true,
				format: "hsl",
			}
		}

		if parsed.format == "hsv" {
			s := convertToPercentage(parsed.s)
			v := convertToPercentage(parsed.v)
			converted := hsvToRgb(parsed.h, s, v)

			var aVal float64
			if s, ok := parsed.a.(string); ok {
				aVal = parseFloat(s)
			} else if f, ok := parsed.a.(float64); ok {
				aVal = f
			}
			rgb.r = converted.R
			rgb.g = converted.G
			rgb.b = converted.B
			rgb.a = boundAlpha(aVal)
			rgb.ok = true
			return rgbResult{
				r:      math.Min(255, math.Max(rgb.r, 0)),
				g:      math.Min(255, math.Max(rgb.g, 0)),
				b:      math.Min(255, math.Max(rgb.b, 0)),
				a:      rgb.a,
				ok:     true,
				format: "hsv",
			}
		}
	}

	var hasA bool
	var a float64 = 1.0

	if m, ok := input.(map[string]interface{}); ok {
		if isValidCSSUnit(m["r"]) && isValidCSSUnit(m["g"]) && isValidCSSUnit(m["b"]) {
			rVal := bound01(m["r"], 255) * 255.0
			gVal := bound01(m["g"], 255) * 255.0
			bVal := bound01(m["b"], 255) * 255.0
			rgb.r = rVal
			rgb.g = gVal
			rgb.b = bVal
			rgb.ok = true
			if s, ok := m["r"].(string); ok && strings.HasSuffix(s, "%") {
				rgb.format = "prgb"
			} else {
				rgb.format = "rgb"
			}
		} else if isValidCSSUnit(m["h"]) && isValidCSSUnit(m["s"]) && isValidCSSUnit(m["v"]) {
			s := convertToPercentage(m["s"])
			v := convertToPercentage(m["v"])
			hVal := bound01(m["h"], 360) * 360.0
			sVal := bound01(s, 100) * 100.0
			vVal := bound01(v, 100) * 100.0
			converted := hsvToRgb(hVal, sVal, vVal)
			rgb.r = converted.R
			rgb.g = converted.G
			rgb.b = converted.B
			rgb.ok = true
			rgb.format = "hsv"
		} else if isValidCSSUnit(m["h"]) && isValidCSSUnit(m["s"]) && isValidCSSUnit(m["l"]) {
			s := convertToPercentage(m["s"])
			l := convertToPercentage(m["l"])
			hVal := bound01(m["h"], 360) * 360.0
			sVal := bound01(s, 100) * 100.0
			lVal := bound01(l, 100) * 100.0
			converted := hslToRgb(hVal, sVal, lVal)
			rgb.r = converted.R
			rgb.g = converted.G
			rgb.b = converted.B
			rgb.ok = true
			rgb.format = "hsl"
		}
		if val, exists := m["a"]; exists && val != nil {
			hasA = true
			switch v := val.(type) {
			case float64:
				a = v
			case float32:
				a = float64(v)
			case int:
				a = float64(v)
			case string:
				a = parseFloat(v)
			}
		}
	} else {
		switch v := input.(type) {
		case RGB:
			if isValidCSSUnit(v.R) && isValidCSSUnit(v.G) && isValidCSSUnit(v.B) {
				rgb.r = bound01(v.R, 255) * 255.0
				rgb.g = bound01(v.G, 255) * 255.0
				rgb.b = bound01(v.B, 255) * 255.0
				rgb.ok = true
				rgb.format = "rgb"
				if v.A != nil {
					hasA = true
					switch va := v.A.(type) {
					case float64:
						a = va
					case float32:
						a = float64(va)
					case int:
						a = float64(va)
					case string:
						a = parseFloat(va)
					}
				}
			}
		case HSL:
			if isValidCSSUnit(v.H) && isValidCSSUnit(v.S) && isValidCSSUnit(v.L) {
				s := convertToPercentage(v.S)
				l := convertToPercentage(v.L)
				converted := hslToRgb(v.H, s, l)
				rgb.r = converted.R
				rgb.g = converted.G
				rgb.b = converted.B
				rgb.ok = true
				rgb.format = "hsl"
				if v.A != nil {
					hasA = true
					switch va := v.A.(type) {
					case float64:
						a = va
					case float32:
						a = float64(va)
					case int:
						a = float64(va)
					case string:
						a = parseFloat(va)
					}
				}
			}
		case HSV:
			if isValidCSSUnit(v.H) && isValidCSSUnit(v.S) && isValidCSSUnit(v.V) {
				s := convertToPercentage(v.S)
				vVal := convertToPercentage(v.V)
				converted := hsvToRgb(v.H, s, vVal)
				rgb.r = converted.R
				rgb.g = converted.G
				rgb.b = converted.B
				rgb.ok = true
				rgb.format = "hsv"
				if v.A != nil {
					hasA = true
					switch va := v.A.(type) {
					case float64:
						a = va
					case float32:
						a = float64(va)
					case int:
						a = float64(va)
					case string:
						a = parseFloat(va)
					}
				}
			}
		}
	}

	if rgb.ok {
		if hasA {
			rgb.a = boundAlpha(a)
		} else {
			rgb.a = 1.0
		}
		return rgbResult{
			r:      math.Min(255, math.Max(rgb.r, 0)),
			g:      math.Min(255, math.Max(rgb.g, 0)),
			b:      math.Min(255, math.Max(rgb.b, 0)),
			a:      rgb.a,
			ok:     true,
			format: rgb.format,
		}
	}

	return rgbResult{r: 0, g: 0, b: 0, a: 1.0, ok: false}
}
