package tinycolor

import (
	"math/rand"
	"sync"
	"time"
)

// rand.Rand isn't safe for concurrent use, so we keep a mutex alongside the
// shared source rather than handing out a fresh *rand.Rand per call (which
// would need a fresh seed each time, and be needlessly wasteful for this).
var (
	rngMu sync.Mutex
	rng   = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// Seed sets the seed for the random color generator
func Seed(seed int64) {
	rngMu.Lock()
	defer rngMu.Unlock()
	rng.Seed(seed)
}

// Random returns a random Color instance with format "prgb" and alpha 1.0,
// matching the JavaScript implementation.
func Random() *Color {
	rngMu.Lock()
	r, g, b := rng.Float64(), rng.Float64(), rng.Float64()
	rngMu.Unlock()

	return FromRatio(RGB{R: r, G: g, B: b}, Options{Format: "prgb"})
}
