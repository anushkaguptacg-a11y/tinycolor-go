package tinycolor

import (
	"math"
	"testing"
)

func TestTinyColorInitialization(t *testing.T) {
	r := Parse("red")
	if r == nil {
		t.Fatal("expected color instance, got nil")
	}

	if Parse(r) != r {
		t.Error("when given a tinycolor instance, Parse() should return the same instance")
	}

	redHex := Parse("red", Options{Format: "hex"})
	if redHex.GetFormat() != "hex" {
		t.Errorf("expected format 'hex', got '%s'", redHex.GetFormat())
	}

	ratioColor := FromRatio(RGB{R: 1, G: 0, B: 0}, Options{Format: "hex"})
	if ratioColor.GetFormat() != "hex" {
		t.Errorf("expected format 'hex', got '%s'", ratioColor.GetFormat())
	}

	obj := HSL{H: 180, S: 0.5, L: 0.5}
	_ = Parse(obj)
	if obj.S != 0.5 {
		t.Errorf("original HSL object was modified: S is %f", obj.S)
	}
}

func TestOriginalInput(t *testing.T) {
	colorRgbUp := "RGB(39, 39, 39)"
	colorRgbLow := "rgb(39, 39, 39)"
	colorRgbMix := "RgB(39, 39, 39)"
	tinycolorObj := Parse(colorRgbMix)
	inputObj := RGB{R: 100, G: 100, B: 100}

	if Parse(colorRgbLow).GetOriginalInput() != colorRgbLow {
		t.Errorf("original lowercase input not returned: got %v", Parse(colorRgbLow).GetOriginalInput())
	}
	if Parse(colorRgbUp).GetOriginalInput() != colorRgbUp {
		t.Errorf("original uppercase input not returned: got %v", Parse(colorRgbUp).GetOriginalInput())
	}
	if Parse(colorRgbMix).GetOriginalInput() != colorRgbMix {
		t.Errorf("original mixed input not returned: got %v", Parse(colorRgbMix).GetOriginalInput())
	}
	if Parse(tinycolorObj).GetOriginalInput() != colorRgbMix {
		t.Errorf("when given tinycolor instance, original color string not returned: got %v", Parse(tinycolorObj).GetOriginalInput())
	}
	if Parse(inputObj).GetOriginalInput() != inputObj {
		t.Errorf("when given an object, the object is not returned: got %v", Parse(inputObj).GetOriginalInput())
	}
	if Parse("").GetOriginalInput() != "" {
		t.Errorf("when given empty string, empty string not returned: got %v", Parse("").GetOriginalInput())
	}
	if Parse(nil).GetOriginalInput() != "" {
		t.Errorf("when given nil, empty string not returned: got %v", Parse(nil).GetOriginalInput())
	}
}

func TestCloningColor(t *testing.T) {
	originalColor := Parse("red")
	clonedColor := originalColor.Clone()

	if clonedColor.r != originalColor.r || clonedColor.g != originalColor.g || clonedColor.b != originalColor.b || clonedColor.a != originalColor.a {
		t.Error("cloned color fields differ from original color")
	}

	clonedColor.SetAlpha(0.5)
	if clonedColor.GetAlpha() == originalColor.GetAlpha() {
		t.Error("cloned color alpha is not changing independently")
	}
}

func TestSettingAlpha(t *testing.T) {
	hexSetter := Parse("rgba(255, 0, 0, 1)")
	if hexSetter.GetAlpha() != 1 {
		t.Errorf("Alpha should start as 1, got %f", hexSetter.GetAlpha())
	}

	returnedFromSetAlpha := hexSetter.SetAlpha(0.9)
	if returnedFromSetAlpha != hexSetter {
		t.Error("SetAlpha return value should be the same color pointer")
	}
	if hexSetter.GetAlpha() != 0.9 {
		t.Errorf("SetAlpha(0.9) should change alpha to 0.9, got %f", hexSetter.GetAlpha())
	}

	hexSetter.SetAlpha(0.5)
	if hexSetter.GetAlpha() != 0.5 {
		t.Errorf("SetAlpha(0.5) should change alpha to 0.5, got %f", hexSetter.GetAlpha())
	}

	hexSetter.SetAlpha(0)
	if hexSetter.GetAlpha() != 0 {
		t.Errorf("SetAlpha(0) should change alpha to 0, got %f", hexSetter.GetAlpha())
	}

	hexSetter.SetAlpha(-1)
	if hexSetter.GetAlpha() != 1 {
		t.Errorf("SetAlpha(-1) should bound to 1, got %f", hexSetter.GetAlpha())
	}

	hexSetter.SetAlpha(2)
	if hexSetter.GetAlpha() != 1 {
		t.Errorf("SetAlpha(2) should bound to 1, got %f", hexSetter.GetAlpha())
	}

	hexSetter.SetAlpha(math.NaN())
	if hexSetter.GetAlpha() != 1 {
		t.Errorf("SetAlpha(NaN) should bound to 1, got %f", hexSetter.GetAlpha())
	}
}

func TestNamedColorsDatabase(t *testing.T) {
	if len(Names) != 149 {
		t.Errorf("expected 149 named colors, got %d", len(Names))
	}
	if Names["aliceblue"] != "f0f8ff" {
		t.Errorf("expected aliceblue to be f0f8ff, got %s", Names["aliceblue"])
	}
	if Names["red"] != "f00" {
		t.Errorf("expected red to be f00, got %s", Names["red"])
	}
	// Verify last-defined wins behavior for duplicate hex values
	if HexNames["0ff"] != "cyan" {
		t.Errorf("expected hex 0ff to map to 'cyan' (last-defined wins), got '%s'", HexNames["0ff"])
	}
	if HexNames["f0f"] != "magenta" {
		t.Errorf("expected hex f0f to map to 'magenta' (last-defined wins), got '%s'", HexNames["f0f"])
	}
	if HexNames["a9a9a9"] != "darkgrey" {
		t.Errorf("expected hex a9a9a9 to map to 'darkgrey', got '%s'", HexNames["a9a9a9"])
	}
}

func TestRGBTextParsing(t *testing.T) {
	tests := []struct {
		input      interface{}
		r, g, b, a float64
		ok         bool
	}{
		{"rgb 255 0 0", 255, 0, 0, 1.0, true},
		{"rgb(255, 0, 0)", 255, 0, 0, 1.0, true},
		{"rgb (255, 0, 0)", 255, 0, 0, 1.0, true},
		{RGB{R: 255, G: 0, B: 0}, 255, 0, 0, 1.0, true},
	}

	for _, tt := range tests {
		c := Parse(tt.input)
		if c.ok != tt.ok || c.r != tt.r || c.g != tt.g || c.b != tt.b || c.a != tt.a {
			t.Errorf("Parse(%v) = {r:%f, g:%f, b:%f, a:%f, ok:%t}, expected {r:%f, g:%f, b:%f, a:%f, ok:%t}",
				tt.input, c.r, c.g, c.b, c.a, c.ok, tt.r, tt.g, tt.b, tt.a, tt.ok)
		}
	}
}

func TestPercentageRGBTextParsing(t *testing.T) {
	tests := []struct {
		input      interface{}
		r, g, b, a float64
		ok         bool
	}{
		{"rgb 100% 0% 0%", 255, 0, 0, 1.0, true},
		{"rgb(100%, 0%, 0%)", 255, 0, 0, 1.0, true},
		{"rgb (100%, 0%, 0%)", 255, 0, 0, 1.0, true},
	}

	for _, tt := range tests {
		c := Parse(tt.input)
		if c.ok != tt.ok || c.r != tt.r || c.g != tt.g || c.b != tt.b || c.a != tt.a {
			t.Errorf("Parse(%v) = {r:%f, g:%f, b:%f, a:%f, ok:%t}, expected {r:%f, g:%f, b:%f, a:%f, ok:%t}",
				tt.input, c.r, c.g, c.b, c.a, c.ok, tt.r, tt.g, tt.b, tt.a, tt.ok)
		}
	}
}

func TestHSLParsing(t *testing.T) {
	// Note: in HSL parsing, we convert to RGB. Since convert.go is implemented in Phase 4,
	// HSL inputs will convert to RGB then. For Phase 3, we verify that HSL inputs are parsed successfully.
	c := Parse("hsl(251, 100%, 38%)")
	if !c.ok {
		t.Error("hsl(251, 100%, 38%) should be valid HSL input")
	}
	c2 := Parse("hsl 100 20 10")
	if !c2.ok {
		t.Error("hsl 100 20 10 should be valid HSL input")
	}
}

func TestHSVParsing(t *testing.T) {
	c := Parse("hsv 251.1 0.887 .918")
	if !c.ok {
		t.Error("hsv 251.1 0.887 .918 should be valid HSV input")
	}
	c2 := Parse("hsva 251.1 0.887 0.918 0.5")
	if !c2.ok {
		t.Error("hsva should be valid HSV input")
	}
	if c2.a != 0.5 {
		t.Errorf("expected alpha 0.5, got %f", c2.a)
	}
}

func TestInvalidParsing(t *testing.T) {
	tests := []interface{}{
		"this is not a color",
		"#red",
		"  #red",
		"##123456",
		"  ##123456",
		RGB{R: math.NaN(), G: 0, B: 0},
	}

	for _, input := range tests {
		c := Parse(input)
		if c.ok {
			t.Errorf("expected input %v to be invalid, but parsed successfully", input)
		}
		// Invalid colors default to black (#000000)
		if c.r != 0 || c.g != 0 || c.b != 0 {
			t.Errorf("invalid color should default to black (0,0,0), got (%f,%f,%f)", c.r, c.g, c.b)
		}
	}
}

func TestOracleRegressionTests(t *testing.T) {
	tests := []struct {
		input      string
		ok         bool
		r, g, b, a float64
	}{
		{"  rgb(255, 0, 0)  ", true, 255, 0, 0, 1.0},
		{"rgb(\t255\t0\t0\t)", true, 255, 0, 0, 1.0},
		{"rgb (255, 0, 0)", true, 255, 0, 0, 1.0},
		{"rgb(255|0|0)", true, 255, 0, 0, 1.0},
		{"rgb(255,0,0)", true, 255, 0, 0, 1.0},
		{"rgb 255 0 0", true, 255, 0, 0, 1.0},
		{"rgba(255, 0, 0, 0.5)", true, 255, 0, 0, 0.5},
		{"rgba 255 0 0 .5", true, 255, 0, 0, 0.5},
		{"rgba(255|0|0|0.5)", true, 255, 0, 0, 0.5},
		{"rgb(100%, 0%, 0%)", true, 255, 0, 0, 1.0},
		{"rgb(100.0%, 0.0%, 0.0%)", true, 255, 0, 0, 1.0},
		{"rgb(110%, -10%, 50%)", true, 25.5, 0, 127.5, 1.0},
		{"rgba(100%, 0%, 0%, 50%)", true, 255, 0, 0, 1.0}, // alpha percent parses as float 50 -> out of bounds -> set to 1.0
		{"RGB(255, 0, 0)", true, 255, 0, 0, 1.0},
		{"HEX8(#FF0000FF)", false, 0, 0, 0, 1.0},
		{"RED", true, 255, 0, 0, 1.0},
		{"  Red  ", true, 255, 0, 0, 1.0},
		{"transparent", true, 0, 0, 0, 0.0},
		{"TRANSPARENT", true, 0, 0, 0, 0.0},
		{"#f00", true, 255, 0, 0, 1.0},
		{"f00", true, 255, 0, 0, 1.0},
		{"#ff0000", true, 255, 0, 0, 1.0},
		{"ff0000", true, 255, 0, 0, 1.0},
		{"#ff0000ff", true, 255, 0, 0, 1.0},
		{"ff0000ff", true, 255, 0, 0, 1.0},
		{"#f00f", true, 255, 0, 0, 1.0},
		{"f00f", true, 255, 0, 0, 1.0},
		{"#f00g", false, 0, 0, 0, 1.0},
		{"#ff000", false, 0, 0, 0, 1.0},
		{"#ff00000", false, 0, 0, 0, 1.0},
		{"#fff000000", false, 0, 0, 0, 1.0},
		{"rgb(255, 0)", false, 0, 0, 0, 1.0},
		{"rgb(255, 0, 0, 0.5)", true, 255, 0, 0, 1.0}, // Matches rgb regex partially, ignoring last arg
		{"rgba(255, 0, 0)", false, 0, 0, 0, 1.0},
		{"rgb(300, -50, 256)", true, 255, 0, 255, 1.0},
		{"rgba(255, 0, 0, 2)", true, 255, 0, 0, 1.0},
		{"rgba(255, 0, 0, -0.5)", true, 255, 0, 0, 1.0},
	}

	for _, tt := range tests {
		c := Parse(tt.input)
		if c.ok != tt.ok || c.r != tt.r || c.g != tt.g || c.b != tt.b || c.a != tt.a {
			t.Errorf("Parse(%q) = {r:%f, g:%f, b:%f, a:%f, ok:%t}, expected {r:%f, g:%f, b:%f, a:%f, ok:%t}",
				tt.input, c.r, c.g, c.b, c.a, c.ok, tt.r, tt.g, tt.b, tt.a, tt.ok)
		}
	}
}

type conversionCase struct {
	hex  string
	hex8 string
	rgb  map[string]interface{}
	hsv  map[string]interface{}
	hsl  map[string]interface{}
}

var conversionCases = []conversionCase{
	{
		hex:  "#ffffff",
		hex8: "#ffffffff",
		rgb:  map[string]interface{}{"r": "100.0%", "g": "100.0%", "b": "100.0%"},
		hsv:  map[string]interface{}{"h": "0", "s": "0.000", "v": "1.000"},
		hsl:  map[string]interface{}{"h": "0", "s": "0.000", "l": "1.000"},
	},
	{
		hex:  "#808080",
		hex8: "#808080ff",
		rgb:  map[string]interface{}{"r": "050.0%", "g": "050.0%", "b": "050.0%"},
		hsv:  map[string]interface{}{"h": "0", "s": "0.000", "v": "0.500"},
		hsl:  map[string]interface{}{"h": "0", "s": "0.000", "l": "0.500"},
	},
	{
		hex:  "#000000",
		hex8: "#000000ff",
		rgb:  map[string]interface{}{"r": "000.0%", "g": "000.0%", "b": "000.0%"},
		hsv:  map[string]interface{}{"h": "0", "s": "0.000", "v": "0.000"},
		hsl:  map[string]interface{}{"h": "0", "s": "0.000", "l": "0.000"},
	},
	{
		hex:  "#ff0000",
		hex8: "#ff0000ff",
		rgb:  map[string]interface{}{"r": "100.0%", "g": "000.0%", "b": "000.0%"},
		hsv:  map[string]interface{}{"h": "0.0", "s": "1.000", "v": "1.000"},
		hsl:  map[string]interface{}{"h": "0.0", "s": "1.000", "l": "0.500"},
	},
	{
		hex:  "#bfbf00",
		hex8: "#bfbf00ff",
		rgb:  map[string]interface{}{"r": "075.0%", "g": "075.0%", "b": "000.0%"},
		hsv:  map[string]interface{}{"h": "60.0", "s": "1.000", "v": "0.750"},
		hsl:  map[string]interface{}{"h": "60.0", "s": "1.000", "l": "0.375"},
	},
	{
		hex:  "#008000",
		hex8: "#008000ff",
		rgb:  map[string]interface{}{"r": "000.0%", "g": "050.0%", "b": "000.0%"},
		hsv:  map[string]interface{}{"h": "120.0", "s": "1.000", "v": "0.500"},
		hsl:  map[string]interface{}{"h": "120.0", "s": "1.000", "l": "0.250"},
	},
	{
		hex:  "#80ffff",
		hex8: "#80ffffff",
		rgb:  map[string]interface{}{"r": "050.0%", "g": "100.0%", "b": "100.0%"},
		hsv:  map[string]interface{}{"h": "180.0", "s": "0.500", "v": "1.000"},
		hsl:  map[string]interface{}{"h": "180.0", "s": "1.000", "l": "0.750"},
	},
	{
		hex:  "#8080ff",
		hex8: "#8080ffff",
		rgb:  map[string]interface{}{"r": "050.0%", "g": "050.0%", "b": "100.0%"},
		hsv:  map[string]interface{}{"h": "240.0", "s": "0.500", "v": "1.000"},
		hsl:  map[string]interface{}{"h": "240.0", "s": "1.000", "l": "0.750"},
	},
	{
		hex:  "#bf40bf",
		hex8: "#bf40bfff",
		rgb:  map[string]interface{}{"r": "075.0%", "g": "025.0%", "b": "075.0%"},
		hsv:  map[string]interface{}{"h": "300.0", "s": "0.667", "v": "0.750"},
		hsl:  map[string]interface{}{"h": "300.0", "s": "0.500", "l": "0.500"},
	},
	{
		hex:  "#a0a424",
		hex8: "#a0a424ff",
		rgb:  map[string]interface{}{"r": "062.8%", "g": "064.3%", "b": "014.2%"},
		hsv:  map[string]interface{}{"h": "61.8", "s": "0.779", "v": "0.643"},
		hsl:  map[string]interface{}{"h": "61.8", "s": "0.638", "l": "0.393"},
	},
	{
		hex:  "#1eac41",
		hex8: "#1eac41ff",
		rgb:  map[string]interface{}{"r": "011.6%", "g": "067.5%", "b": "025.5%"},
		hsv:  map[string]interface{}{"h": "134.9", "s": "0.828", "v": "0.675"},
		hsl:  map[string]interface{}{"h": "134.9", "s": "0.707", "l": "0.396"},
	},
	{
		hex:  "#b430e5",
		hex8: "#b430e5ff",
		rgb:  map[string]interface{}{"r": "070.4%", "g": "018.7%", "b": "089.7%"},
		hsv:  map[string]interface{}{"h": "283.7", "s": "0.792", "v": "0.897"},
		hsl:  map[string]interface{}{"h": "283.7", "s": "0.775", "l": "0.542"},
	},
	{
		hex:  "#fef888",
		hex8: "#fef888ff",
		rgb:  map[string]interface{}{"r": "099.8%", "g": "097.4%", "b": "053.2%"},
		hsv:  map[string]interface{}{"h": "56.9", "s": "0.467", "v": "0.998"},
		hsl:  map[string]interface{}{"h": "56.9", "s": "0.991", "l": "0.765"},
	},
	{
		hex:  "#19cb97",
		hex8: "#19cb97ff",
		rgb:  map[string]interface{}{"r": "009.9%", "g": "079.5%", "b": "059.1%"},
		hsv:  map[string]interface{}{"h": "162.4", "s": "0.875", "v": "0.795"},
		hsl:  map[string]interface{}{"h": "162.4", "s": "0.779", "l": "0.447"},
	},
	{
		hex:  "#362698",
		hex8: "#362698ff",
		rgb:  map[string]interface{}{"r": "021.1%", "g": "014.9%", "b": "059.7%"},
		hsv:  map[string]interface{}{"h": "248.3", "s": "0.750", "v": "0.597"},
		hsl:  map[string]interface{}{"h": "248.3", "s": "0.601", "l": "0.373"},
	},
	{
		hex:  "#7e7eb8",
		hex8: "#7e7eb8ff",
		rgb:  map[string]interface{}{"r": "049.5%", "g": "049.3%", "b": "072.1%"},
		hsv:  map[string]interface{}{"h": "240.5", "s": "0.316", "v": "0.721"},
		hsl:  map[string]interface{}{"h": "240.5", "s": "0.290", "l": "0.607"},
	},
}

func TestColorEquality(t *testing.T) {
	for _, c := range conversionCases {
		if !Equals(c.rgb, c.hex) {
			t.Errorf("expected RGB %v to equal Hex %s", c.rgb, c.hex)
		}
		if !Equals(c.rgb, c.hex8) {
			t.Errorf("expected RGB %v to equal Hex8 %s", c.rgb, c.hex8)
		}
		if !Equals(c.rgb, c.hsl) {
			t.Errorf("expected RGB %v to equal HSL %v", c.rgb, c.hsl)
		}
		if !Equals(c.rgb, c.hsv) {
			t.Errorf("expected RGB %v to equal HSV %v", c.rgb, c.hsv)
		}
		if !Equals(c.rgb, c.rgb) {
			t.Errorf("expected RGB %v to equal itself", c.rgb)
		}

		if !Equals(c.hex, c.hex) {
			t.Errorf("expected Hex %s to equal itself", c.hex)
		}
		if !Equals(c.hex, c.hex8) {
			t.Errorf("expected Hex %s to equal Hex8 %s", c.hex, c.hex8)
		}
		if !Equals(c.hex, c.hsl) {
			t.Errorf("expected Hex %s to equal HSL %v", c.hex, c.hsl)
		}
		if !Equals(c.hex, c.hsv) {
			t.Errorf("expected Hex %s to equal HSV %v", c.hex, c.hsv)
		}

		if !Equals(c.hsl, c.hsv) {
			t.Errorf("expected HSL %v to equal HSV %v", c.hsl, c.hsv)
		}
	}
}

func TestConversionObjects(t *testing.T) {
	for _, c := range conversionCases {
		color := Parse(c.hex)
		if !color.ok {
			t.Errorf("expected color %s to be valid", c.hex)
		}

		hslObj := color.HSL()
		parsedHsl := Parse(hslObj)
		if !Equals(color, parsedHsl) {
			t.Errorf("expected parsed HSL %v to equal original color %s", hslObj, c.hex)
		}

		hsvObj := color.HSV()
		parsedHsv := Parse(hsvObj)
		if !Equals(color, parsedHsv) {
			t.Errorf("expected parsed HSV %v to equal original color %s", hsvObj, c.hex)
		}

		rgbObj := color.RGB()
		parsedRgb := Parse(rgbObj)
		if !Equals(color, parsedRgb) {
			t.Errorf("expected parsed RGB %v to equal original color %s", rgbObj, c.hex)
		}
	}
}

func TestFormattingAndSerialization(t *testing.T) {
	// 1. WithRatio
	cRatio1 := FromRatio(RGB{R: 1, G: 1, B: 1})
	if cRatio1.HexString(false) != "#ffffff" {
		t.Errorf("expected #ffffff, got %s", cRatio1.HexString(false))
	}
	cRatio2 := FromRatio(RGB{R: 1, G: 0, B: 0, A: 0.5})
	if cRatio2.RGBString() != "rgba(255, 0, 0, 0.5)" {
		t.Errorf("expected rgba(255, 0, 0, 0.5), got %s", cRatio2.RGBString())
	}
	cRatio3 := FromRatio(RGB{R: 1, G: 0, B: 0, A: 1})
	if cRatio3.RGBString() != "rgb(255, 0, 0)" {
		t.Errorf("expected rgb(255, 0, 0), got %s", cRatio3.RGBString())
	}

	// 2. WithoutRatio
	cNoRatio1 := Parse(RGB{R: 1, G: 1, B: 1})
	if cNoRatio1.HexString(false) != "#010101" {
		t.Errorf("expected #010101, got %s", cNoRatio1.HexString(false))
	}
	cNoRatio2 := Parse(RGB{R: 0.1, G: 0.1, B: 0.1})
	if cNoRatio2.HexString(false) != "#000000" {
		t.Errorf("expected #000000, got %s", cNoRatio2.HexString(false))
	}

	// 3. Hex Formatting
	cRed := Parse("rgb 255 0 0")
	if cRed.HexString(false) != "#ff0000" {
		t.Errorf("expected #ff0000, got %s", cRed.HexString(false))
	}
	if cRed.HexString(true) != "#f00" {
		t.Errorf("expected #f00, got %s", cRed.HexString(true))
	}
	cRedAlpha := Parse("rgba 255 0 0 0.5")
	if cRedAlpha.Hex8String(false) != "#ff000080" {
		t.Errorf("expected #ff000080, got %s", cRedAlpha.Hex8String(false))
	}
	if Parse("rgba 255 0 0 0").Hex8String(false) != "#ff000000" {
		t.Errorf("expected #ff000000, got %s", Parse("rgba 255 0 0 0").Hex8String(false))
	}
	if Parse("rgba 255 0 0 1").Hex8String(false) != "#ff0000ff" {
		t.Errorf("expected #ff0000ff, got %s", Parse("rgba 255 0 0 1").Hex8String(false))
	}
	if Parse("rgba 255 0 0 1").Hex8String(true) != "#f00f" {
		t.Errorf("expected #f00f, got %s", Parse("rgba 255 0 0 1").Hex8String(true))
	}
	if cRed.Hex(false) != "ff0000" {
		t.Errorf("expected ff0000, got %s", cRed.Hex(false))
	}
	if cRed.Hex(true) != "f00" {
		t.Errorf("expected f00, got %s", cRed.Hex(true))
	}
	if cRedAlpha.Hex8(false) != "ff000080" {
		t.Errorf("expected ff000080, got %s", cRedAlpha.Hex8(false))
	}

	// 4. HSV String
	cHSV1 := Parse("hsv 251.1 0.887 .918")
	if cHSV1.HSVString() != "hsv(251, 89%, 92%)" {
		t.Errorf("expected hsv(251, 89%%, 92%%), got %s", cHSV1.HSVString())
	}
	cHSV2 := Parse("hsva 251.1 0.887 0.918 0.5")
	if cHSV2.HSVString() != "hsva(251, 89%, 92%, 0.5)" {
		t.Errorf("expected hsva(251, 89%%, 92%%, 0.5), got %s", cHSV2.HSVString())
	}

	// 5. Named Colors
	if name, _ := Parse("#f00").Name(); name != "red" {
		t.Errorf("expected red, got %s", name)
	}
	if _, found := Parse("#fa0a0a").Name(); found {
		t.Error("expected false for #fa0a0a")
	}

	// 6. Invalid Alpha Normalization
	cAlphaNeg := Parse(map[string]interface{}{"r": 255.0, "g": 20.0, "b": 10.0, "a": -1.0})
	if cAlphaNeg.RGBString() != "rgb(255, 20, 10)" {
		t.Errorf("expected rgb(255, 20, 10), got %s", cAlphaNeg.RGBString())
	}
	cAlphaNegZero := Parse(map[string]interface{}{"r": 255.0, "g": 20.0, "b": 10.0, "a": -0.0})
	if cAlphaNegZero.RGBString() != "rgba(255, 20, 10, 0)" {
		t.Errorf("expected rgba(255, 20, 10, 0), got %s", cAlphaNegZero.RGBString())
	}
	cAlphaZero := Parse(map[string]interface{}{"r": 255.0, "g": 20.0, "b": 10.0, "a": 0.0})
	if cAlphaZero.RGBString() != "rgba(255, 20, 10, 0)" {
		t.Errorf("expected rgba(255, 20, 10, 0), got %s", cAlphaZero.RGBString())
	}
	cAlphaHalf := Parse(map[string]interface{}{"r": 255.0, "g": 20.0, "b": 10.0, "a": 0.5})
	if cAlphaHalf.RGBString() != "rgba(255, 20, 10, 0.5)" {
		t.Errorf("expected rgba(255, 20, 10, 0.5), got %s", cAlphaHalf.RGBString())
	}
	cAlphaOne := Parse(map[string]interface{}{"r": 255.0, "g": 20.0, "b": 10.0, "a": 1.0})
	if cAlphaOne.RGBString() != "rgb(255, 20, 10)" {
		t.Errorf("expected rgb(255, 20, 10), got %s", cAlphaOne.RGBString())
	}
	cAlphaHuge := Parse(map[string]interface{}{"r": 255.0, "g": 20.0, "b": 10.0, "a": 100.0})
	if cAlphaHuge.RGBString() != "rgb(255, 20, 10)" {
		t.Errorf("expected rgb(255, 20, 10), got %s", cAlphaHuge.RGBString())
	}
	cAlphaStringInvalid := Parse(map[string]interface{}{"r": 255.0, "g": 20.0, "b": 10.0, "a": "asdfasd"})
	if cAlphaStringInvalid.RGBString() != "rgb(255, 20, 10)" {
		t.Errorf("expected rgb(255, 20, 10), got %s", cAlphaStringInvalid.RGBString())
	}

	// 7. toString() with Alpha Set
	redNamed := FromRatio(RGB{R: 255, G: 0, B: 0, A: 0.6}, Options{Format: "name"})
	redHex := FromRatio(RGB{R: 255, G: 0, B: 0, A: 0.4}, Options{Format: "hex"})
	transparentNamed := FromRatio(RGB{R: 255, G: 0, B: 0, A: 0}, Options{Format: "name"})

	if redNamed.GetFormat() != "name" {
		t.Errorf("expected format name, got %s", redNamed.GetFormat())
	}
	if redHex.GetFormat() != "hex" {
		t.Errorf("expected format hex, got %s", redHex.GetFormat())
	}

	if redNamed.String() != "rgba(255, 0, 0, 0.6)" {
		t.Errorf("expected rgba(255, 0, 0, 0.6), got %s", redNamed.String())
	}
	if redHex.String() != "rgba(255, 0, 0, 0.4)" {
		t.Errorf("expected rgba(255, 0, 0, 0.4), got %s", redHex.String())
	}

	if redNamed.String("hex") != "#ff0000" {
		t.Errorf("expected #ff0000, got %s", redNamed.String("hex"))
	}
	if redNamed.String("hex6") != "#ff0000" {
		t.Errorf("expected #ff0000, got %s", redNamed.String("hex6"))
	}
	if redNamed.String("hex3") != "#f00" {
		t.Errorf("expected #f00, got %s", redNamed.String("hex3"))
	}
	if redNamed.String("hex8") != "#ff000099" {
		t.Errorf("expected #ff000099, got %s", redNamed.String("hex8"))
	}
	if redNamed.String("hex4") != "#f009" {
		t.Errorf("expected #f009, got %s", redNamed.String("hex4"))
	}
	if redNamed.String("name") != "#ff0000" {
		t.Errorf("expected #ff0000, got %s", redNamed.String("name"))
	}

	if _, found := redNamed.Name(); found {
		t.Error("expected redNamed.Name() to be false")
	}
	if transparentNamed.String() != "transparent" {
		t.Errorf("expected transparent, got %s", transparentNamed.String())
	}

	redHex.SetAlpha(0)
	if redHex.String() != "rgba(255, 0, 0, 0)" {
		t.Errorf("expected rgba(255, 0, 0, 0), got %s", redHex.String())
	}

	// 8. setting alpha
	hexSetter := Parse("rgba(255, 0, 0, 1)")
	if hexSetter.GetAlpha() != 1.0 {
		t.Errorf("expected alpha 1, got %f", hexSetter.GetAlpha())
	}
	returnedFromSetAlpha := hexSetter.SetAlpha(0.9)
	if returnedFromSetAlpha != hexSetter {
		t.Error("setAlpha should return the color instance")
	}
	if hexSetter.GetAlpha() != 0.9 {
		t.Errorf("expected alpha 0.9, got %f", hexSetter.GetAlpha())
	}
	hexSetter.SetAlpha(0.5)
	if hexSetter.GetAlpha() != 0.5 {
		t.Errorf("expected alpha 0.5, got %f", hexSetter.GetAlpha())
	}
	hexSetter.SetAlpha(0)
	if hexSetter.GetAlpha() != 0.0 {
		t.Errorf("expected alpha 0, got %f", hexSetter.GetAlpha())
	}
	hexSetter.SetAlpha(-1)
	if hexSetter.GetAlpha() != 1.0 {
		t.Errorf("expected alpha 1.0, got %f", hexSetter.GetAlpha())
	}
	hexSetter.SetAlpha(2)
	if hexSetter.GetAlpha() != 1.0 {
		t.Errorf("expected alpha 1.0, got %f", hexSetter.GetAlpha())
	}
	hexSetter.SetAlpha(math.NaN())
	if hexSetter.GetAlpha() != 1.0 {
		t.Errorf("expected alpha 1.0, got %f", hexSetter.GetAlpha())
	}

	// 9. transparent named toName()
	if name, _ := Parse(RGB{R: 255, G: 20, B: 10, A: 0}).Name(); name != "transparent" {
		t.Errorf("expected transparent, got %s", name)
	}
	if Parse("transparent").String() != "transparent" {
		t.Errorf("expected transparent, got %s", Parse("transparent").String())
	}
	if Parse("transparent").Hex(false) != "000000" {
		t.Errorf("expected 000000, got %s", Parse("transparent").Hex(false))
	}

	// 10. getBrightness, isDark, isLight
	cDark := Parse("#000")
	cLight := Parse("#fff")
	if cDark.GetBrightness() != 0 {
		t.Errorf("expected brightness 0, got %f", cDark.GetBrightness())
	}
	if cLight.GetBrightness() != 255 {
		t.Errorf("expected brightness 255, got %f", cLight.GetBrightness())
	}
	if !cDark.IsDark() {
		t.Error("expected #000 to be dark")
	}
	if cLight.IsDark() {
		t.Error("expected #fff to not be dark")
	}
	if cDark.IsLight() {
		t.Error("expected #000 to not be light")
	}
	if !cLight.IsLight() {
		t.Error("expected #fff to be light")
	}

	// 11. toFilter
	cFilter1 := Parse("rgba(255, 0, 0, 1)")
	expectedFilter1 := "progid:DXImageTransform.Microsoft.gradient(startColorstr=#ffff0000,endColorstr=#ffff0000)"
	if cFilter1.Filter() != expectedFilter1 {
		t.Errorf("expected %s, got %s", expectedFilter1, cFilter1.Filter())
	}
	expectedFilter2 := "progid:DXImageTransform.Microsoft.gradient(startColorstr=#ffff0000,endColorstr=#ff00ff00)"
	if cFilter1.Filter("rgba(0, 255, 0, 1)") != expectedFilter2 {
		t.Errorf("expected %s, got %s", expectedFilter2, cFilter1.Filter("rgba(0, 255, 0, 1)"))
	}
}

func BenchmarkParse(b *testing.B) {
	inputs := []string{
		"#ff0000",
		"rgb(255, 0, 0)",
		"rgba(255, 0, 0, 0.5)",
		"hsl(251, 89%, 92%)",
		"hsla(251, 89%, 92%, 0.5)",
		"red",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, input := range inputs {
			_ = Parse(input)
		}
	}
}

func BenchmarkConversion(b *testing.B) {
	c := Parse("#6699cc")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.HSL()
		_ = c.HSV()
	}
}

func BenchmarkFormatting(b *testing.B) {
	c := Parse("#6699cc")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.RGBString()
		_ = c.HexString(false)
		_ = c.HSLString()
	}
}
