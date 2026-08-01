package tinycolor

import (
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// Seed sets the seed for the random color generator
func Seed(seed int64) {
	rng.Seed(seed)
}

// Random returns a random Color instance with format "prgb" and alpha 1.0,
// matching the JavaScript implementation.
func Random() *Color {
	return FromRatio(RGB{
		R: rng.Float64(),
		G: rng.Float64(),
		B: rng.Float64(),
	}, Options{Format: "prgb"})
}
