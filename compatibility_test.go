package tinycolor

import (
	_ "embed"
	"encoding/json"
	"math"
	"strings"
	"testing"
)

//go:embed testdata/oracle_results.json
var oracleResultsJSON []byte

type oracleTestCase struct {
	Input  string  `json:"input"`
	Ok     bool    `json:"ok"`
	R      float64 `json:"r"`
	G      float64 `json:"g"`
	B      float64 `json:"b"`
	A      float64 `json:"a"`
	Hex    string  `json:"hex"`
	RgbStr string  `json:"rgbStr"`
}

func TestGoldenOracleCompatibility(t *testing.T) {
	var testCases []oracleTestCase
	err := json.Unmarshal(oracleResultsJSON, &testCases)
	if err != nil {
		t.Fatalf("failed to unmarshal oracle results JSON: %v", err)
	}

	const floatEpsilon = 1e-4 // allow minor floating point deviations

	for _, tc := range testCases {
		t.Run(tc.Input, func(t *testing.T) {
			c := Parse(tc.Input)
			if c.IsValid() != tc.Ok {
				t.Fatalf("IsValid() parity mismatch: expected %t, got %t", tc.Ok, c.IsValid())
			}

			if !tc.Ok {
				return
			}

			// Validate R, G, B, A channels
			if math.Abs(c.r-tc.R) > floatEpsilon {
				t.Errorf("R channel mismatch: expected %f, got %f", tc.R, c.r)
			}
			if math.Abs(c.g-tc.G) > floatEpsilon {
				t.Errorf("G channel mismatch: expected %f, got %f", tc.G, c.g)
			}
			if math.Abs(c.b-tc.B) > floatEpsilon {
				t.Errorf("B channel mismatch: expected %f, got %f", tc.B, c.b)
			}
			if math.Abs(c.GetAlpha()-tc.A) > floatEpsilon {
				t.Errorf("Alpha channel mismatch: expected %f, got %f", tc.A, c.GetAlpha())
			}

			// Validate Hex
			expectedHex := strings.TrimPrefix(strings.ToLower(tc.Hex), "#")
			gotHex := strings.ToLower(c.Hex(false))
			if gotHex != expectedHex {
				t.Errorf("Hex output mismatch: expected %s, got %s", expectedHex, gotHex)
			}

			// Validate RGB String representation
			expectedRgbStr := strings.ReplaceAll(tc.RgbStr, " ", "")
			gotRgbStr := strings.ReplaceAll(c.RGBString(), " ", "")
			if gotRgbStr != expectedRgbStr {
				t.Errorf("RGB String mismatch: expected %s, got %s", expectedRgbStr, gotRgbStr)
			}
		})
	}
}
