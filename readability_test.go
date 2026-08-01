package tinycolor

import (
	"math"
	"testing"
)

func TestReadabilityCalculations(t *testing.T) {
	// WCAG Examples
	// Black vs White = 21:1
	if r := Readability("#000", "#fff"); math.Abs(r-21.0) > 1e-5 {
		t.Errorf("expected 21:1 for black vs white, got %f", r)
	}
	// Same color = 1:1
	if r := Readability("#f00", "#f00"); math.Abs(r-1.0) > 1e-5 {
		t.Errorf("expected 1:1 for same color, got %f", r)
	}
	// Pure Red vs White
	// Red L = 0.2126, White L = 1.0
	// (1.0 + 0.05) / (0.2126 + 0.05) = 1.05 / 0.2626 = 3.99847677
	expectedRedWhite := 1.05 / 0.2626
	if r := Readability("#f00", "#fff"); math.Abs(r-expectedRedWhite) > 1e-5 {
		t.Errorf("expected %f for red vs white, got %f", expectedRedWhite, r)
	}
	// Pure Blue vs White
	// Blue L = 0.0722, White L = 1.0
	// (1.0 + 0.05) / (0.0722 + 0.05) = 1.05 / 0.1222 = 8.59247135
	expectedBlueWhite := 1.05 / 0.1222
	if r := Readability("#00f", "#fff"); math.Abs(r-expectedBlueWhite) > 1e-5 {
		t.Errorf("expected %f for blue vs white, got %f", expectedBlueWhite, r)
	}

	// Ported Readability function tests
	if r := Readability("#000", "#000"); r != 1.0 {
		t.Errorf("expected 1, got %f", r)
	}
	if r := Readability("#000", "#111"); math.Abs(r-1.1121078) > 1e-5 {
		t.Errorf("expected ~1.112, got %f", r)
	}
	if r := Readability("#000", "#fff"); r != 21.0 {
		t.Errorf("expected 21, got %f", r)
	}
}

func TestIsReadable(t *testing.T) {
	// level: AA, size: small (default) -> contrast >= 4.5
	if IsReadable("#ff0088", "#5c1a72") {
		t.Error("expected false for #ff0088 vs #5c1a72 (AA small)")
	}
	// level: AA, size: large -> contrast >= 3.0
	if !IsReadable("#ff0088", "#5c1a72", WCAG2Opts{Level: "AA", Size: "large"}) {
		t.Error("expected true for #ff0088 vs #5c1a72 (AA large)")
	}
	// level: AAA, size: small -> contrast >= 7.0
	if IsReadable("#ff0088", "#5c1a72", WCAG2Opts{Level: "AAA", Size: "small"}) {
		t.Error("expected false for #ff0088 vs #5c1a72 (AAA small)")
	}
	// level: AAA, size: large -> contrast >= 4.5
	if IsReadable("#ff0088", "#5c1a72", WCAG2Opts{Level: "AAA", Size: "large"}) {
		t.Error("expected false for #ff0088 vs #5c1a72 (AAA large)")
	}

	// #ff0088 vs #2e0c3a: contrast ratio 4.56
	if !IsReadable("#ff0088", "#2e0c3a") {
		t.Error("expected true for #ff0088 vs #2e0c3a (AA small)")
	}
	if !IsReadable("#ff0088", "#2e0c3a", WCAG2Opts{Level: "AA", Size: "large"}) {
		t.Error("expected true for #ff0088 vs #2e0c3a (AA large)")
	}
	if IsReadable("#ff0088", "#2e0c3a", WCAG2Opts{Level: "AAA", Size: "small"}) {
		t.Error("expected false for #ff0088 vs #2e0c3a (AAA small)")
	}
	if !IsReadable("#ff0088", "#2e0c3a", WCAG2Opts{Level: "AAA", Size: "large"}) {
		t.Error("expected true for #ff0088 vs #2e0c3a (AAA large)")
	}

	// #db91b8 vs #2e0c3a: contrast ratio 7.12
	if !IsReadable("#db91b8", "#2e0c3a") {
		t.Error("expected true for #db91b8 vs #2e0c3a (AA small)")
	}
	if !IsReadable("#db91b8", "#2e0c3a", WCAG2Opts{Level: "AA", Size: "large"}) {
		t.Error("expected true for #db91b8 vs #2e0c3a (AA large)")
	}
	if !IsReadable("#db91b8", "#2e0c3a", WCAG2Opts{Level: "AAA", Size: "small"}) {
		t.Error("expected true for #db91b8 vs #2e0c3a (AAA small)")
	}
	if !IsReadable("#db91b8", "#2e0c3a", WCAG2Opts{Level: "AAA", Size: "large"}) {
		t.Error("expected true for #db91b8 vs #2e0c3a (AAA large)")
	}
}

func TestMostReadable(t *testing.T) {
	// Candidate present and readable
	best1 := MostReadable("#000", []interface{}{"#111", "#222"})
	if best1.Hex(false) != "222222" {
		t.Errorf("expected 222222, got %s", best1.Hex(false))
	}

	best2 := MostReadable("#f00", []interface{}{"#d00", "#0d0"})
	if best2.Hex(false) != "00dd00" {
		t.Errorf("expected 00dd00, got %s", best2.Hex(false))
	}

	best3 := MostReadable("#fff", []interface{}{"#fff", "#fff"})
	if best3.Hex(false) != "ffffff" {
		t.Errorf("expected ffffff, got %s", best3.Hex(false))
	}

	// includeFallbackColors
	bestFallback := MostReadable("#fff", []interface{}{"#fff", "#fff"}, WCAG2Opts{IncludeFallbackColors: true})
	if bestFallback.Hex(false) != "000000" {
		t.Errorf("expected fallback 000000, got %s", bestFallback.Hex(false))
	}

	bestNoFallback := MostReadable("#123", []interface{}{"#124", "#125"}, WCAG2Opts{IncludeFallbackColors: false})
	if bestNoFallback.Hex(false) != "112255" {
		t.Errorf("expected 112255, got %s", bestNoFallback.Hex(false))
	}

	bestFallback2 := MostReadable("#123", []interface{}{"#124", "#125"}, WCAG2Opts{IncludeFallbackColors: true})
	if bestFallback2.Hex(false) != "ffffff" {
		t.Errorf("expected fallback ffffff, got %s", bestFallback2.Hex(false))
	}

	bestFallback3 := MostReadable("#ff0088", []interface{}{"#000", "#fff"}, WCAG2Opts{IncludeFallbackColors: false})
	if bestFallback3.Hex(false) != "000000" {
		t.Errorf("expected 000000, got %s", bestFallback3.Hex(false))
	}

	bestFallback4 := MostReadable("#ff0088", []interface{}{"#2e0c3a"}, WCAG2Opts{
		IncludeFallbackColors: true,
		Level:                 "AAA",
		Size:                  "large",
	})
	if bestFallback4.Hex(false) != "2e0c3a" {
		t.Errorf("expected 2e0c3a, got %s", bestFallback4.Hex(false))
	}

	bestFallback5 := MostReadable("#ff0088", []interface{}{"#2e0c3a"}, WCAG2Opts{
		IncludeFallbackColors: true,
		Level:                 "AAA",
		Size:                  "small",
	})
	if bestFallback5.Hex(false) != "000000" {
		t.Errorf("expected fallback 000000, got %s", bestFallback5.Hex(false))
	}
}

func TestMostReadableEdgeCases(t *testing.T) {
	// 1. Empty candidate list
	// includeFallbackColors = false -> returns nil
	if res := MostReadable("#fff", []interface{}{}, WCAG2Opts{IncludeFallbackColors: false}); res != nil {
		t.Errorf("expected nil for empty candidate list, got %v", res)
	}
	// includeFallbackColors = true -> falls back to black/white
	if res := MostReadable("#fff", []interface{}{}, WCAG2Opts{IncludeFallbackColors: true}); res.Hex(false) != "000000" {
		t.Errorf("expected fallback 000000 for empty candidate list, got %s", res.Hex(false))
	}

	// 2. Ties (should return the first one with the best score)
	// Both "#111" and "#111" have the same readability vs white. The first one should be chosen.
	first := "#111"
	candidates := []interface{}{first, "#111"}
	res := MostReadable("#fff", candidates)
	if res.Hex(false) != "111111" {
		t.Errorf("expected 111111, got %s", res.Hex(false))
	}

	// 3. Alpha handling
	// Readability should only use R,G,B channels. Let's pass colors with alpha
	c1 := RGB{R: 0, G: 0, B: 0, A: 0.5}
	c2 := RGB{R: 255, G: 255, B: 255, A: 0.1}
	if r := Readability(c1, c2); r != 21.0 {
		t.Errorf("expected 21:1 regardless of alpha, got %f", r)
	}

	// 4. No readable color available (IncludeFallbackColors: false)
	// Base is white. Candidate list is only white. Contrast is 1.0 (unreadable).
	// Should return white (the candidate) because fallback is disabled.
	resUnreadable := MostReadable("#fff", []interface{}{"#fff"}, WCAG2Opts{IncludeFallbackColors: false})
	if resUnreadable.Hex(false) != "ffffff" {
		t.Errorf("expected ffffff, got %s", resUnreadable.Hex(false))
	}
}
