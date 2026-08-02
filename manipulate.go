package tinycolor

import (
	"math"
)

// applyModification writes new RGB(A) values back onto the receiver after a
// manipulation like Lighten/Darken/Saturate. This mirrors what Parse() does
// in color.go for the same reason: values under 1.0 get rounded immediately
// to resolve the 0-255-vs-0-1 ambiguity, but values >= 1.0 are left as
// floats so precision survives chained calls (Lighten().Darken()...), same
// as tinycolor2's constructor path.
func (c *Color) applyModification(r, g, b, a float64) *Color {
	if c == nil {
		return nil
	}
	if r < 1.0 {
		r = jsRound(r)
	}
	if g < 1.0 {
		g = jsRound(g)
	}
	if b < 1.0 {
		b = jsRound(b)
	}
	c.r = r
	c.g = g
	c.b = b
	c.a = boundAlpha(a)
	c.roundA = jsRound(c.a*100.0) / 100.0
	return c
}

// Lighten bumps up the L channel of the color's HSL representation. amount
// defaults to 10 and is a percentage, so Lighten(20) moves L up by 0.2.
// Negative amounts work too (same as calling Darken).
func (c *Color) Lighten(amount ...int) *Color {
	if c == nil {
		return nil
	}
	amt := 10
	if len(amount) > 0 {
		amt = amount[0]
	}
	hsl := c.HSL()
	hsl.L += float64(amt) / 100.0
	hsl.L = clamp01(hsl.L)
	s := convertToPercentage(hsl.S)
	l := convertToPercentage(hsl.L)
	rgb := hslToRgb(hsl.H, s, l)
	return c.applyModification(rgb.R, rgb.G, rgb.B, c.a)
}

// Darken is Lighten with the sign flipped.
func (c *Color) Darken(amount ...int) *Color {
	if c == nil {
		return nil
	}
	amt := 10
	if len(amount) > 0 {
		amt = amount[0]
	}
	hsl := c.HSL()
	hsl.L -= float64(amt) / 100.0
	hsl.L = clamp01(hsl.L)
	s := convertToPercentage(hsl.S)
	l := convertToPercentage(hsl.L)
	rgb := hslToRgb(hsl.H, s, l)
	return c.applyModification(rgb.R, rgb.G, rgb.B, c.a)
}

// Brighten works directly on RGB rather than going through HSL like the
// other adjustments - that's how tinycolor2 does it too, so keeping it here.
func (c *Color) Brighten(amount ...int) *Color {
	if c == nil {
		return nil
	}
	amt := 10
	if len(amount) > 0 {
		amt = amount[0]
	}
	factor := jsRound(255.0 * -(float64(amt) / 100.0))
	r := math.Max(0, math.Min(255, jsRound(c.r)-factor))
	g := math.Max(0, math.Min(255, jsRound(c.g)-factor))
	b := math.Max(0, math.Min(255, jsRound(c.b)-factor))
	return c.applyModification(r, g, b, c.a)
}

// Saturate pushes the S channel up by amount/100 (default 10%).
func (c *Color) Saturate(amount ...int) *Color {
	if c == nil {
		return nil
	}
	amt := 10
	if len(amount) > 0 {
		amt = amount[0]
	}
	hsl := c.HSL()
	hsl.S += float64(amt) / 100.0
	hsl.S = clamp01(hsl.S)
	s := convertToPercentage(hsl.S)
	l := convertToPercentage(hsl.L)
	rgb := hslToRgb(hsl.H, s, l)
	return c.applyModification(rgb.R, rgb.G, rgb.B, c.a)
}

// Desaturate is the inverse of Saturate.
func (c *Color) Desaturate(amount ...int) *Color {
	if c == nil {
		return nil
	}
	amt := 10
	if len(amount) > 0 {
		amt = amount[0]
	}
	hsl := c.HSL()
	hsl.S -= float64(amt) / 100.0
	hsl.S = clamp01(hsl.S)
	s := convertToPercentage(hsl.S)
	l := convertToPercentage(hsl.L)
	rgb := hslToRgb(hsl.H, s, l)
	return c.applyModification(rgb.R, rgb.G, rgb.B, c.a)
}

// Greyscale is just Desaturate(100).
func (c *Color) Greyscale() *Color {
	if c == nil {
		return nil
	}
	return c.Desaturate(100)
}

// Spin rotates hue by the given number of degrees, wrapping into [0, 360).
func (c *Color) Spin(amount float64) *Color {
	if c == nil {
		return nil
	}
	hsl := c.HSL()
	hue := math.Mod(hsl.H+amount, 360.0)
	if hue < 0 {
		hue = 360.0 + hue
	}
	hsl.H = hue
	s := convertToPercentage(hsl.S)
	l := convertToPercentage(hsl.L)
	rgb := hslToRgb(hsl.H, s, l)
	return c.applyModification(rgb.R, rgb.G, rgb.B, c.a)
}

// Mix blends two colors by a given percentage weight (0-100)
func Mix(color1, color2 interface{}, amount ...int) *Color {
	amt := 50
	if len(amount) > 0 {
		amt = amount[0]
	}
	c1 := Parse(color1)
	c2 := Parse(color2)
	p := float64(amt) / 100.0

	rgb1 := c1.RGB()
	rgb2 := c2.RGB()

	r := (rgb2.R-rgb1.R)*p + rgb1.R
	g := (rgb2.G-rgb1.G)*p + rgb1.G
	b := (rgb2.B-rgb1.B)*p + rgb1.B

	c1a := 0.0
	if c1 != nil {
		c1a = c1.a
	}
	c2a := 0.0
	if c2 != nil {
		c2a = c2.a
	}
	a := (c2a-c1a)*p + c1a

	return Parse(RGB{
		R: r,
		G: g,
		B: b,
		A: a,
	})
}

// Complement returns the complementary color (rotated 180 degrees)
func (c *Color) Complement() *Color {
	if c == nil {
		return nil
	}
	hsl := c.HSL()
	hsl.H = math.Mod(hsl.H+180.0, 360.0)
	return Parse(hsl)
}

// Analogous returns a list of analogous colors
func (c *Color) Analogous(resultsAndSlices ...int) []*Color {
	if c == nil {
		return nil
	}
	res := 6
	if len(resultsAndSlices) > 0 {
		res = resultsAndSlices[0]
	}
	slices := 30
	if len(resultsAndSlices) > 1 {
		slices = resultsAndSlices[1]
	}
	hsl := c.HSL()
	part := 360.0 / float64(slices)
	ret := []*Color{c.Clone()}

	h := math.Mod(hsl.H-part*float64(res/2)+720.0, 360.0)
	for res > 1 {
		h = math.Mod(h+part, 360.0)
		ret = append(ret, Parse(HSL{H: h, S: hsl.S, L: hsl.L, A: c.a}))
		res--
	}
	return ret
}

// Monochromatic returns a list of monochromatic colors
func (c *Color) Monochromatic(results ...int) []*Color {
	if c == nil {
		return nil
	}
	res := 6
	if len(results) > 0 {
		res = results[0]
	}
	hsv := c.HSV()
	ret := []*Color{}
	modification := 1.0 / float64(res)
	v := hsv.V
	for res > 0 {
		ret = append(ret, Parse(HSV{H: hsv.H, S: hsv.S, V: v, A: c.a}))
		v = math.Mod(v+modification, 1.0)
		res--
	}
	return ret
}

// SplitComplement returns a list of split complementary colors
func (c *Color) SplitComplement() []*Color {
	if c == nil {
		return nil
	}
	hsl := c.HSL()
	return []*Color{
		c.Clone(),
		Parse(HSL{H: math.Mod(hsl.H+72.0, 360.0), S: hsl.S, L: hsl.L, A: c.a}),
		Parse(HSL{H: math.Mod(hsl.H+216.0, 360.0), S: hsl.S, L: hsl.L, A: c.a}),
	}
}

// polyad is helper for triad and tetrad
func (c *Color) polyad(number int) []*Color {
	if c == nil {
		return nil
	}
	hsl := c.HSL()
	result := []*Color{c.Clone()}
	step := 360.0 / float64(number)
	for i := 1; i < number; i++ {
		result = append(result, Parse(HSL{
			H: math.Mod(hsl.H+float64(i)*step, 360.0),
			S: hsl.S,
			L: hsl.L,
			A: c.a,
		}))
	}
	return result
}

// Triad returns a list of triad colors
func (c *Color) Triad() []*Color {
	if c == nil {
		return nil
	}
	return c.polyad(3)
}

// Tetrad returns a list of tetrad colors
func (c *Color) Tetrad() []*Color {
	if c == nil {
		return nil
	}
	return c.polyad(4)
}
