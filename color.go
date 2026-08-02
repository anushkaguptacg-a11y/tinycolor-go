package tinycolor

import (
	"math"
)

// RGB represents an RGB color input/output with values in [0, 255] (or floats in [0, 255])
type RGB struct {
	R float64
	G float64
	B float64
	A interface{}
}

// PercentageRGB represents a percentage-based RGB color representation
type PercentageRGB struct {
	R string
	G string
	B string
	A float64
}

// HSL represents HSL color format
type HSL struct {
	H float64
	S float64
	L float64
	A interface{}
}

// HSV represents HSV color format
type HSV struct {
	H float64
	S float64
	V float64
	A interface{}
}

// Options contains instantiation configuration
type Options struct {
	Format       string
	GradientType string
}

// Color is the main color struct preserving intermediate floating point values
type Color struct {
	r             float64
	g             float64
	b             float64
	a             float64
	roundA        float64
	format        string
	gradientType  string
	ok            bool
	originalInput interface{}
}

// Parse parses a given interface input into a Color instance
func Parse(input interface{}, opts ...Options) *Color {
	if input == nil {
		c := &Color{
			r:             0,
			g:             0,
			b:             0,
			a:             1.0,
			roundA:        1.0,
			ok:            false,
			originalInput: "",
		}
		if len(opts) > 0 {
			c.format = opts[0].Format
			c.gradientType = opts[0].GradientType
		}
		return c
	}

	// If input is already *Color, hand back the same instance - this matches
	// tinycolor2's `instanceof` check in the JS source, which also returns
	// the original reference rather than a copy. Not a missed Clone().
	if color, ok := input.(*Color); ok {
		return color
	}

	var format string
	var gradientType string
	if len(opts) > 0 {
		format = opts[0].Format
		gradientType = opts[0].GradientType
	}

	rgb := inputToRGB(input)

	c := &Color{
		r:             rgb.r,
		g:             rgb.g,
		b:             rgb.b,
		a:             rgb.a,
		roundA:        jsRound(100.0*rgb.a) / 100.0,
		format:        rgb.format,
		gradientType:  gradientType,
		ok:            rgb.ok,
		originalInput: input,
	}

	if format != "" {
		c.format = format
	}

	// Don't let the range of [0,255] come back in [0,1].
	// Potentially lose a little bit of precision here, but will fix issues where
	// .5 gets interpreted as half of the total, instead of half of 1
	// If it was supposed to be 128, this was already taken care of by `inputToRgb`
	if c.r < 1.0 {
		c.r = jsRound(c.r)
	}
	if c.g < 1.0 {
		c.g = jsRound(c.g)
	}
	if c.b < 1.0 {
		c.b = jsRound(c.b)
	}

	return c
}

// FromRatio initializes a Color from ratio inputs (all in [0,1]).
// Only RGB needs special handling here: inputToRGB's HSL/HSV branches
// already run S/L/V through convertToPercentage, so ratio values there
// resolve correctly through a plain Parse() call. RGB's branch doesn't do
// that conversion, so the *255 below is what actually makes ratios work.
func FromRatio(input interface{}, opts ...Options) *Color {
	if rgb, ok := input.(RGB); ok {
		return Parse(RGB{
			R: rgb.R * 255,
			G: rgb.G * 255,
			B: rgb.B * 255,
			A: rgb.A,
		}, opts...)
	}
	return Parse(input, opts...)
}

// IsValid checks if the color was parsed successfully
func (c *Color) IsValid() bool {
	if c == nil {
		return false
	}
	return c.ok
}

// GetOriginalInput returns the original input used to create the Color
func (c *Color) GetOriginalInput() interface{} {
	if c == nil {
		return nil
	}
	return c.originalInput
}

// GetFormat returns the format of the color (e.g. hex, rgb, name, etc.)
func (c *Color) GetFormat() string {
	if c == nil {
		return ""
	}
	return c.format
}

// GetAlpha returns the alpha value of the color
func (c *Color) GetAlpha() float64 {
	if c == nil {
		return 0.0
	}
	return c.a
}

// SetAlpha sets the alpha value of the color
func (c *Color) SetAlpha(value float64) *Color {
	if c == nil {
		return nil
	}
	c.a = boundAlpha(value)
	c.roundA = jsRound(c.a*100.0) / 100.0
	return c
}

// Clone creates a new identical copy of the Color instance
func (c *Color) Clone() *Color {
	if c == nil {
		return nil
	}
	return &Color{
		r:             c.r,
		g:             c.g,
		b:             c.b,
		a:             c.a,
		roundA:        c.roundA,
		format:        c.format,
		gradientType:  c.gradientType,
		ok:            c.ok,
		originalInput: c.originalInput,
	}
}

// Equals compares two colors for equality
func Equals(color1, color2 interface{}) bool {
	c1 := Parse(color1)
	c2 := Parse(color2)
	if c1 == nil || c2 == nil || !c1.ok || !c2.ok {
		return false
	}
	r1 := jsRound(c1.r)
	g1 := jsRound(c1.g)
	b1 := jsRound(c1.b)
	a1 := jsRound(c1.a*100.0) / 100.0

	r2 := jsRound(c2.r)
	g2 := jsRound(c2.g)
	b2 := jsRound(c2.b)
	a2 := jsRound(c2.a*100.0) / 100.0

	return r1 == r2 && g1 == g2 && b1 == b2 && a1 == a2
}

// RGB returns the rounded RGB values and alpha of the color
func (c *Color) RGB() RGB {
	if c == nil {
		return RGB{}
	}
	return RGB{
		R: jsRound(c.r),
		G: jsRound(c.g),
		B: jsRound(c.b),
		A: c.a,
	}
}

// HSL returns the HSL representation of the color (h in [0, 360], s and l in [0, 1])
func (c *Color) HSL() HSL {
	if c == nil {
		return HSL{}
	}
	hsl := rgbToHsl(c.r, c.g, c.b)
	return HSL{
		H: hsl.H * 360.0,
		S: hsl.S,
		L: hsl.L,
		A: c.a,
	}
}

// HSV returns the HSV representation of the color (h in [0, 360], s and v in [0, 1])
func (c *Color) HSV() HSV {
	if c == nil {
		return HSV{}
	}
	hsv := rgbToHsv(c.r, c.g, c.b)
	return HSV{
		H: hsv.H * 360.0,
		S: hsv.S,
		V: hsv.V,
		A: c.a,
	}
}

// GetBrightness returns the brightness of the color (0-255)
func (c *Color) GetBrightness() float64 {
	if c == nil {
		return 0.0
	}
	rgb := c.RGB()
	return (rgb.R*299.0 + rgb.G*587.0 + rgb.B*114.0) / 1000.0
}

// GetLuminance returns the relative luminance of the color (0-1)
func (c *Color) GetLuminance() float64 {
	if c == nil {
		return 0.0
	}
	rgb := c.RGB()
	rSRGB := rgb.R / 255.0
	gSRGB := rgb.G / 255.0
	bSRGB := rgb.B / 255.0

	var r, g, b float64
	if rSRGB <= 0.03928 {
		r = rSRGB / 12.92
	} else {
		r = math.Pow((rSRGB+0.055)/1.055, 2.4)
	}

	if gSRGB <= 0.03928 {
		g = gSRGB / 12.92
	} else {
		g = math.Pow((gSRGB+0.055)/1.055, 2.4)
	}

	if bSRGB <= 0.03928 {
		b = bSRGB / 12.92
	} else {
		b = math.Pow((bSRGB+0.055)/1.055, 2.4)
	}

	return 0.2126*r + 0.7152*g + 0.0722*b
}

// IsDark returns true if the color is dark
func (c *Color) IsDark() bool {
	if c == nil {
		return false
	}
	return c.GetBrightness() < 128.0
}

// IsLight returns true if the color is light
func (c *Color) IsLight() bool {
	if c == nil {
		return false
	}
	return !c.IsDark()
}
