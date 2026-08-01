package tinycolor

import (
	"regexp"
)

// Regex patterns exactly matching the JavaScript implementation
const (
	cssInteger = `[-\+]?\d+%?`
	cssNumber  = `[-\+]?\d*\.\d+%?`
	cssUnit    = `(?:` + cssNumber + `)|(?:` + cssInteger + `)`

	// permissiveMatch3 matches functions with 3 arguments (e.g. rgb, hsl, hsv)
	permissiveMatch3 = `[\s|\(]+(` + cssUnit + `)[,|\s]+(` + cssUnit + `)[,|\s]+(` + cssUnit + `)\s*\)?`
	// permissiveMatch4 matches functions with 4 arguments (e.g. rgba, hsla, hsva)
	permissiveMatch4 = `[\s|\(]+(` + cssUnit + `)[,|\s]+(` + cssUnit + `)[,|\s]+(` + cssUnit + `)[,|\s]+(` + cssUnit + `)\s*\)?`
)

var (
	trimLeftReg  = regexp.MustCompile(`^\s+`)
	trimRightReg = regexp.MustCompile(`\s+$`)

	cssUnitReg = regexp.MustCompile(cssUnit)
	rgbReg     = regexp.MustCompile(`(?i)rgb` + permissiveMatch3)
	rgbaReg    = regexp.MustCompile(`(?i)rgba` + permissiveMatch4)
	hslReg     = regexp.MustCompile(`(?i)hsl` + permissiveMatch3)
	hslaReg    = regexp.MustCompile(`(?i)hsla` + permissiveMatch4)
	hsvReg     = regexp.MustCompile(`(?i)hsv` + permissiveMatch3)
	hsvaReg    = regexp.MustCompile(`(?i)hsva` + permissiveMatch4)

	hex3Reg = regexp.MustCompile(`^#?([0-9a-fA-F]{1})([0-9a-fA-F]{1})([0-9a-fA-F]{1})$`)
	hex6Reg = regexp.MustCompile(`^#?([0-9a-fA-F]{2})([0-9a-fA-F]{2})([0-9a-fA-F]{2})$`)
	hex4Reg = regexp.MustCompile(`^#?([0-9a-fA-F]{1})([0-9a-fA-F]{1})([0-9a-fA-F]{1})([0-9a-fA-F]{1})$`)
	hex8Reg = regexp.MustCompile(`^#?([0-9a-fA-F]{2})([0-9a-fA-F]{2})([0-9a-fA-F]{2})([0-9a-fA-F]{2})$`)
)
