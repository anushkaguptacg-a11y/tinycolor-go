package tinycolor

import (
	"strings"
	"testing"
)

func TestRandomColor(t *testing.T) {
	Seed(12345)

	randomColor := Random()
	if randomColor == nil {
		t.Fatal("expected random color, got nil")
	}

	if randomColor.GetAlpha() != 1.0 {
		t.Errorf("expected alpha 1.0, got %f", randomColor.GetAlpha())
	}

	if randomColor.GetFormat() != "prgb" {
		t.Errorf("expected format prgb, got %s", randomColor.GetFormat())
	}

	randomColor.SetAlpha(0.5)
	hex8 := randomColor.Hex8String(false)
	if !strings.HasSuffix(hex8, "80") {
		t.Errorf("expected hex8 suffix '80', got %s", hex8)
	}
}
