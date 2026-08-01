package tinycolor

import (
	"math"
	"strconv"
	"strings"
	"testing"
)

var DESATURATIONS = []string{
	"ff0000", "fe0101", "fc0303", "fb0404", "fa0505", "f90606", "f70808", "f60909", "f50a0a", "f40b0b",
	"f20d0d", "f10e0e", "f00f0f", "ee1111", "ed1212", "ec1313", "eb1414", "e91616", "e81717", "e71818",
	"e61919", "e41b1b", "e31c1c", "e21d1d", "e01f1f", "df2020", "de2121", "dd2222", "db2424", "da2525",
	"d92626", "d72828", "d62929", "d52a2a", "d42b2b", "d22d2d", "d12e2e", "d02f2f", "cf3030", "cd3232",
	"cc3333", "cb3434", "c93636", "c83737", "c73838", "c63939", "c43b3b", "c33c3c", "c23d3d", "c13e3e",
	"bf4040", "be4141", "bd4242", "bb4444", "ba4545", "b94646", "b84747", "b64949", "b54a4a", "b44b4b",
	"b34d4d", "b14e4e", "b04f4f", "af5050", "ad5252", "ac5353", "ab5454", "aa5555", "a85757", "a75858",
	"a65959", "a45b5b", "a35c5c", "a25d5d", "a15e5e", "9f6060", "9e6161", "9d6262", "9c6363", "9a6565",
	"996666", "986767", "966969", "956a6a", "946b6b", "936c6c", "916e6e", "906f6f", "8f7070", "8e7171",
	"8c7373", "8b7474", "8a7575", "887777", "877878", "867979", "857a7a", "837c7c", "827d7d", "817e7e",
	"808080",
}

var SATURATIONS = []string{
	"ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000",
	"ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000",
	"ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000",
	"ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000",
	"ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000",
	"ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000",
	"ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000",
	"ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000",
	"ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000",
	"ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000", "ff0000",
	"ff0000",
}

var LIGHTENS = []string{
	"ff0000", "ff0505", "ff0a0a", "ff0f0f", "ff1414", "ff1a1a", "ff1f1f", "ff2424", "ff2929", "ff2e2e",
	"ff3333", "ff3838", "ff3d3d", "ff4242", "ff4747", "ff4d4d", "ff5252", "ff5757", "ff5c5c", "ff6161",
	"ff6666", "ff6b6b", "ff7070", "ff7575", "ff7a7a", "ff8080", "ff8585", "ff8a8a", "ff8f8f", "ff9494",
	"ff9999", "ff9e9e", "ffa3a3", "ffa8a8", "ffadad", "ffb3b3", "ffb8b8", "ffbdbd", "ffc2c2", "ffc7c7",
	"ffcccc", "ffd1d1", "ffd6d6", "ffdbdb", "ffe0e0", "ffe5e5", "ffebeb", "fff0f0", "fff5f5", "fffafa",
	"ffffff", "ffffff", "ffffff", "ffffff", "ffffff", "ffffff", "ffffff", "ffffff", "ffffff", "ffffff",
	"ffffff", "ffffff", "ffffff", "ffffff", "ffffff", "ffffff", "ffffff", "ffffff", "ffffff", "ffffff",
	"ffffff", "ffffff", "ffffff", "ffffff", "ffffff", "ffffff", "ffffff", "ffffff", "ffffff", "ffffff",
	"ffffff", "ffffff", "ffffff", "ffffff", "ffffff", "ffffff", "ffffff", "ffffff", "ffffff", "ffffff",
	"ffffff", "ffffff", "ffffff", "ffffff", "ffffff", "ffffff", "ffffff", "ffffff", "ffffff", "ffffff",
	"ffffff",
}

var BRIGHTENS = []string{
	"ff0000", "ff0303", "ff0505", "ff0808", "ff0a0a", "ff0d0d", "ff0f0f", "ff1212", "ff1414", "ff1717",
	"ff1919", "ff1c1c", "ff1f1f", "ff2121", "ff2424", "ff2626", "ff2929", "ff2b2b", "ff2e2e", "ff3030",
	"ff3333", "ff3636", "ff3838", "ff3b3b", "ff3d3d", "ff4040", "ff4242", "ff4545", "ff4747", "ff4a4a",
	"ff4c4c", "ff4f4f", "ff5252", "ff5454", "ff5757", "ff5959", "ff5c5c", "ff5e5e", "ff6161", "ff6363",
	"ff6666", "ff6969", "ff6b6b", "ff6e6e", "ff7070", "ff7373", "ff7575", "ff7878", "ff7a7a", "ff7d7d",
	"ff7f7f", "ff8282", "ff8585", "ff8787", "ff8a8a", "ff8c8c", "ff8f8f", "ff9191", "ff9494", "ff9696",
	"ff9999", "ff9c9c", "ff9e9e", "ffa1a1", "ffa3a3", "ffa6a6", "ffa8a8", "ffabab", "ffadad", "ffb0b0",
	"ffb2b2", "ffb5b5", "ffb8b8", "ffbaba", "ffbdbd", "ffbfbf", "ffc2c2", "ffc4c4", "ffc7c7", "ffc9c9",
	"ffcccc", "ffcfcf", "ffd1d1", "ffd4d4", "ffd6d6", "ffd9d9", "ffdbdb", "ffdede", "ffe0e0", "ffe3e3",
	"ffe5e5", "ffe8e8", "ffebeb", "ffeded", "fff0f0", "fff2f2", "fff5f5", "fff7f7", "fffafa", "fffcfc",
	"ffffff",
}

var DARKENS = []string{
	"ff0000", "fa0000", "f50000", "f00000", "eb0000", "e60000", "e00000", "db0000", "d60000", "d10000",
	"cc0000", "c70000", "c20000", "bd0000", "b80000", "b30000", "ad0000", "a80000", "a30000", "9e0000",
	"990000", "940000", "8f0000", "8a0000", "850000", "800000", "7a0000", "750000", "700000", "6b0000",
	"660000", "610000", "5c0000", "570000", "520000", "4d0000", "470000", "420000", "3d0000", "380000",
	"330000", "2e0000", "290000", "240000", "1f0000", "190000", "140000", "0f0000", "0a0000", "050000",
	"000000", "000000", "000000", "000000", "000000", "000000", "000000", "000000", "000000", "000000",
	"000000", "000000", "000000", "000000", "000000", "000000", "000000", "000000", "000000", "000000",
	"000000", "000000", "000000", "000000", "000000", "000000", "000000", "000000", "000000", "000000",
	"000000", "000000", "000000", "000000", "000000", "000000", "000000", "000000", "000000", "000000",
	"000000", "000000", "000000", "000000", "000000", "000000", "000000", "000000", "000000", "000000",
	"000000",
}

func TestModifications(t *testing.T) {
	for i := 0; i <= 100; i++ {
		c := Parse("red")
		if c.Desaturate(i).Hex(false) != DESATURATIONS[i] {
			t.Errorf("Desaturate %d expected %s, got %s", i, DESATURATIONS[i], c.Hex(false))
		}
	}
	for i := 0; i <= 100; i++ {
		c := Parse("red")
		if c.Saturate(i).Hex(false) != SATURATIONS[i] {
			t.Errorf("Saturate %d expected %s, got %s", i, SATURATIONS[i], c.Hex(false))
		}
	}
	for i := 0; i <= 100; i++ {
		c := Parse("red")
		if c.Lighten(i).Hex(false) != LIGHTENS[i] {
			t.Errorf("Lighten %d expected %s, got %s", i, LIGHTENS[i], c.Hex(false))
		}
	}
	for i := 0; i <= 100; i++ {
		c := Parse("red")
		if c.Brighten(i).Hex(false) != BRIGHTENS[i] {
			t.Errorf("Brighten %d expected %s, got %s", i, BRIGHTENS[i], c.Hex(false))
		}
	}
	for i := 0; i <= 100; i++ {
		c := Parse("red")
		if c.Darken(i).Hex(false) != DARKENS[i] {
			t.Errorf("Darken %d expected %s, got %s", i, DARKENS[i], c.Hex(false))
		}
	}

	cGrey := Parse("red")
	if cGrey.Greyscale().Hex(false) != "808080" {
		t.Errorf("Greyscale expected 808080, got %s", cGrey.Hex(false))
	}
}

func TestSpin(t *testing.T) {
	if h := math.Round(Parse("#f00").Spin(-1234).HSL().H); h != 206 {
		t.Errorf("Spin -1234 expected 206, got %f", h)
	}
	if h := math.Round(Parse("#f00").Spin(-360).HSL().H); h != 0 {
		t.Errorf("Spin -360 expected 0, got %f", h)
	}
	if h := math.Round(Parse("#f00").Spin(-120).HSL().H); h != 240 {
		t.Errorf("Spin -120 expected 240, got %f", h)
	}
	if h := math.Round(Parse("#f00").Spin(0).HSL().H); h != 0 {
		t.Errorf("Spin 0 expected 0, got %f", h)
	}
	if h := math.Round(Parse("#f00").Spin(10).HSL().H); h != 10 {
		t.Errorf("Spin 10 expected 10, got %f", h)
	}
	if h := math.Round(Parse("#f00").Spin(360).HSL().H); h != 0 {
		t.Errorf("Spin 360 expected 0, got %f", h)
	}
	if h := math.Round(Parse("#f00").Spin(2345).HSL().H); h != 185 {
		t.Errorf("Spin 2345 expected 185, got %f", h)
	}

	deltas := []float64{-360, 0, 360}
	for _, delta := range deltas {
		for name := range Names {
			c := Parse(name)
			orig := c.Hex(false)
			spun := c.Clone().Spin(delta).Hex(false)
			if spun != orig {
				t.Errorf("Spin %f on %s should have no effect, expected %s, got %s", delta, name, orig, spun)
			}
		}
	}
}

func TestMix(t *testing.T) {
	if l := Mix("#000", "#fff").HSL().L; l != 0.5 {
		t.Errorf("Mix #000 #fff expected lightness 0.5, got %f", l)
	}
	if h := Mix("#f00", "#000", 0).Hex(false); h != "ff0000" {
		t.Errorf("Mix #f00 #000 0 expected ff0000, got %s", h)
	}
	if h := Mix("#fff", "#000", 90).Hex(false); h != "1a1a1a" {
		t.Errorf("Mix #fff #000 90 expected 1a1a1a, got %s", h)
	}

	for i := 0; i < 100; i++ {
		expectedL := float64(i) / 100.0
		l := math.Round(Mix("#000", "#fff", i).HSL().L*100.0) / 100.0
		if l != expectedL {
			t.Errorf("Mix #000 #fff %d expected lightness %f, got %f", i, expectedL, l)
		}
	}

	for i := 0; i < 100; i++ {
		newHex := strconv.FormatInt(int64(math.Round(255.0*float64(100-i)/100.0)), 16)
		if len(newHex) == 1 {
			newHex = "0" + newHex
		}

		if h := Mix("#f00", "#000", i).Hex(false); h != newHex+"0000" {
			t.Errorf("Mix #f00 #000 %d expected %s, got %s", i, newHex+"0000", h)
		}
		if h := Mix("#0f0", "#000", i).Hex(false); h != "00"+newHex+"00" {
			t.Errorf("Mix #0f0 #000 %d expected %s, got %s", i, "00"+newHex+"00", h)
		}
		if h := Mix("#00f", "#000", i).Hex(false); h != "0000"+newHex {
			t.Errorf("Mix #00f #000 %d expected %s, got %s", i, "0000"+newHex, h)
		}
		if a := Mix("transparent", "#000", i).GetAlpha(); a != float64(i)/100.0 {
			t.Errorf("Mix transparent #000 %d expected alpha %f, got %f", i, float64(i)/100.0, a)
		}
	}
}

func helperColorsToHex(colors []*Color) string {
	hexes := []string{}
	for _, c := range colors {
		hexes = append(hexes, c.Hex(false))
	}
	return strings.Join(hexes, ",")
}

func TestCombinations(t *testing.T) {
	cComplement := Parse("red")
	if cComplement.Complement().Hex(false) != "00ffff" {
		t.Errorf("Complement expected 00ffff, got %s", cComplement.Complement().Hex(false))
	}
	if cComplement.Hex(false) != "ff0000" {
		t.Errorf("Complement mutated instance! Expected ff0000, got %s", cComplement.Hex(false))
	}

	cAnalogous := Parse("red")
	expectedAnalogous := "ff0000,ff0066,ff0033,ff0000,ff3300,ff6600"
	if h := helperColorsToHex(cAnalogous.Analogous()); h != expectedAnalogous {
		t.Errorf("Analogous expected %s, got %s", expectedAnalogous, h)
	}

	cMono := Parse("red")
	expectedMono := "ff0000,2a0000,550000,800000,aa0000,d40000"
	if h := helperColorsToHex(cMono.Monochromatic()); h != expectedMono {
		t.Errorf("Monochromatic expected %s, got %s", expectedMono, h)
	}

	cSplit := Parse("red")
	expectedSplit := "ff0000,ccff00,0066ff"
	if h := helperColorsToHex(cSplit.SplitComplement()); h != expectedSplit {
		t.Errorf("SplitComplement expected %s, got %s", expectedSplit, h)
	}

	cTriad := Parse("red")
	expectedTriad := "ff0000,00ff00,0000ff"
	if h := helperColorsToHex(cTriad.Triad()); h != expectedTriad {
		t.Errorf("Triad expected %s, got %s", expectedTriad, h)
	}

	cTetrad := Parse("red")
	expectedTetrad := "ff0000,80ff00,00ffff,7f00ff"
	if h := helperColorsToHex(cTetrad.Tetrad()); h != expectedTetrad {
		t.Errorf("Tetrad expected %s, got %s", expectedTetrad, h)
	}
}

func TestMutationSemanticsAndNegativeValues(t *testing.T) {
	// Mutation semantics
	c := Parse("#f00")
	d := c.Lighten(10)
	if c != d {
		t.Error("Lighten did not return same pointer")
	}
	if c.Hex(false) != "ff3333" {
		t.Errorf("original color not mutated, expected ff3333, got %s", c.Hex(false))
	}

	// Negative values
	cNeg1 := Parse("#f00")
	cNeg1.Lighten(-10) // should act like Darken(10) -> red's lightness is 50%, goes to 40% -> #cc0000
	if cNeg1.Hex(false) != "cc0000" {
		t.Errorf("Lighten(-10) expected cc0000, got %s", cNeg1.Hex(false))
	}

	cNeg2 := Parse("#f00")
	cNeg2.Darken(-10) // should act like Lighten(10) -> lightness goes to 60% -> #ff3333
	if cNeg2.Hex(false) != "ff3333" {
		t.Errorf("Darken(-10) expected ff3333, got %s", cNeg2.Hex(false))
	}

	cNeg3 := Parse("#f00")
	cNeg3.Spin(-180) // should rotate red (hue 0) to hue 180 (cyan) -> #00ffff
	if cNeg3.Hex(false) != "00ffff" {
		t.Errorf("Spin(-180) expected 00ffff, got %s", cNeg3.Hex(false))
	}

	cNeg4 := Parse("#f00")
	cNeg4.Saturate(-50) // should act like Desaturate(50) -> saturation goes to 50% -> #bf4040
	if cNeg4.Hex(false) != "bf4040" {
		t.Errorf("Saturate(-50) expected bf4040, got %s", cNeg4.Hex(false))
	}
}

func TestChainableManipulations(t *testing.T) {
	c := Parse("#6699cc")
	d := c.Lighten(10).Saturate(20).Spin(30)
	// Let's verify result against JS
	// In JS, tinycolor("#6699cc").lighten(10).saturate(20).spin(30).toHex() -> "7d7de8"
	expectedHex := "7d7de8"
	if d.Hex(false) != expectedHex {
		t.Errorf("chained operations expected %s, got %s", expectedHex, d.Hex(false))
	}
	if c != d {
		t.Error("chaining did not preserve pointer identity")
	}
}
