// Package randstr provides you the generator of random strings.
package randstr

import (
	"math"
	"math/rand"
	"time"
)

var (
	defaultCharacters []rune      = []rune("abcdefgihijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	randSrc           rand.Source = rand.NewSource(time.Now().UnixNano())
)

// New generates a random string which has the specified length.
//
// You can specify characters used to generate it.
// The default characters is alphabets (Uppercase & Lowercase) and numbers.
func New(l int, opts ...func(*Config)) string {
	conf := &Config{
		chars: defaultCharacters,
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
	cache, remain := randSrc.Int63(), maxIdx
	for i := l - 1; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), maxIdx
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
	bits := math.Ceil(math.Log(float64(len(chars)))) + 1
	return uint(bits)
}

func maskIndex(chars []rune) int64 {
	mask := 1<<bitsIndex(chars) - 1
	return int64(mask)
}

func maxIndex(chars []rune) uint {
	return 63 / bitsIndex(chars)
}
