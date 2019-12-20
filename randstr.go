// Package randstr provides you the generator of random strings.
package randstr

import (
	"math"
	"math/rand"
	"time"
)

var (
	defaultChars []rune     = []rune(`abcdefgihijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890`)
	defaultRand  *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

var globalRand *rand.Rand = defaultRand

// Seed uses the provided seed value to initialize the default Source to a
// deterministic state. If Seed is not called, the generator behaves as
// if seeded by Seed(time.Now().UnixNano()).
func Seed(seed int64) {
	globalRand = rand.New(rand.NewSource(seed))
}

// New generates a random string which has the specified length.
//
// You can specify characters used to generate it.
// The default characters is alphabets (Uppercase & Lowercase) and numbers.
func New(l int, opts ...func(*Config)) string {
	conf := &Config{
		chars: defaultChars,
	}
	for _, opt := range opts {
		opt(conf)
	}

	return newWithConfig(l, conf)
}

func newWithConfig(l int, conf *Config) string {
	bitsIdx := bitsIndex(conf.chars)
	maskIdx := maskIndex(conf.chars)
	maxIdx := maxIndex(conf.chars)

	chars := make([]rune, l)
	cache, remain := globalRand.Int63(), maxIdx
	for i := l - 1; i >= 0; {
		if remain == 0 {
			cache, remain = globalRand.Int63(), maxIdx
		}
		idx := int(cache & maskIdx)
		if idx < len(conf.chars) {
			chars[i] = conf.chars[idx]
			i--
		}
		cache >>= bitsIdx
		remain--
	}
	return string(chars)
}

func bitsIndex(chars []rune) uint {
	bits := math.Trunc(math.Log2(float64(len(chars)))) + 1
	return uint(bits)
}

func maskIndex(chars []rune) int64 {
	mask := 1<<bitsIndex(chars) - 1
	return int64(mask)
}

func maxIndex(chars []rune) uint {
	return 63 / bitsIndex(chars)
}
