package tinycolor

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

var floatReg = regexp.MustCompile(`^\s*([-\+]?(?:\d*\.\d+|\d+\.?))`)

// parseFloat mimics JS parseFloat behavior. It trims whitespace and parses
// float prefix, returning math.NaN() if no float was matched.
func parseFloat(s string) float64 {
	m := floatReg.FindStringSubmatch(s)
	if len(m) < 2 {
		return math.NaN()
	}
	val, err := strconv.ParseFloat(m[1], 64)
	if err != nil {
		return math.NaN()
	}
	return val
}

// boundAlpha bounds an alpha value to [0,1], with NaN, < 0, or > 1 mapping to 1.0.
func boundAlpha(a float64) float64 {
	if math.IsNaN(a) || a < 0.0 || a > 1.0 {
		return 1.0
	}
	return a
}

// bound01 takes an input representation in [0, max] and returns it mapped to [0, 1].
// It handles string inputs with "%" symbols, "1.0" edge cases, and floating roundoffs.
func bound01(n interface{}, max float64) float64 {
	if isOnePointZero(n) {
		n = "100%"
	}
	processPercent := isPercentage(n)

	var fVal float64
	switch v := n.(type) {
	case string:
		fVal = parseFloat(v)
	case float64:
		fVal = v
	case float32:
		fVal = float64(v)
	case int:
		fVal = float64(v)
	case int64:
		fVal = float64(v)
	default:
		fVal = 0.0
	}

	if math.IsNaN(fVal) {
		return 0.0
	}

	// Clamp to [0, max]
	if fVal < 0 {
		fVal = 0
	}
	if fVal > max {
		fVal = max
	}

	if processPercent {
		fVal = float64(int(fVal*max)) / 100.0
	}

	if math.Abs(fVal-max) < 0.000001 {
		return 1.0
	}

	return math.Mod(fVal, max) / max
}

// clamp01 bounds a value between 0 and 1.
func clamp01(val float64) float64 {
	if val < 0 {
		return 0
	}
	if val > 1 {
		return 1
	}
	return val
}

// parseIntFromHex parses a hex string chunk into base-10 integer.
func parseIntFromHex(val string) int64 {
	v, err := strconv.ParseInt(val, 16, 64)
	if err != nil {
		return 0
	}
	return v
}

// isOnePointZero detects if input is a string representing "1" or "1.0" with a decimal point.
func isOnePointZero(n interface{}) bool {
	if s, ok := n.(string); ok {
		return strings.Contains(s, ".") && parseFloat(s) == 1.0
	}
	return false
}

// isPercentage checks if input string represents percentage.
func isPercentage(n interface{}) bool {
	if s, ok := n.(string); ok {
		return strings.Contains(s, "%")
	}
	return false
}

// pad2 ensures a hex component has at least 2 characters.
func pad2(c string) string {
	if len(c) == 1 {
		return "0" + c
	}
	return c
}

// convertToPercentage converts a decimal <= 1 to its percentage string value.
func convertToPercentage(n interface{}) interface{} {
	switch v := n.(type) {
	case float64:
		if v <= 1.0 {
			return strconv.FormatFloat(v*100.0, 'f', -1, 64) + "%"
		}
		return v
	case float32:
		if v <= 1.0 {
			return strconv.FormatFloat(float64(v)*100.0, 'f', -1, 32) + "%"
		}
		return float64(v)
	case int:
		if v <= 1 {
			return strconv.Itoa(v*100) + "%"
		}
		return float64(v)
	case string:
		if strings.Contains(v, "%") {
			return v
		}
		fVal := parseFloat(v)
		if !math.IsNaN(fVal) && fVal <= 1.0 {
			return strconv.FormatFloat(fVal*100.0, 'f', -1, 64) + "%"
		}
		return v
	}
	return n
}

// convertDecimalToHex converts a decimal in [0,1] to a hex integer representation.
func convertDecimalToHex(d string) string {
	val := jsRound(parseFloat(d) * 255.0)
	if math.IsNaN(val) {
		val = 255.0
	}
	h := strconv.FormatInt(int64(val), 16)
	return h
}

// convertHexToDecimal converts a hex string back to a float in [0,1].
func convertHexToDecimal(h string) float64 {
	return float64(parseIntFromHex(h)) / 255.0
}


// jsRound emulates JavaScript Math.round exactly
func jsRound(x float64) float64 {
	return math.Floor(x + 0.5)
}
