package randstr

import (
	"math/rand"
)

// Config is the configuration to generate a random string.
type Config struct {
	chars []rune
	rand  *rand.Rand
}

// WithChars sets characters to use as a optional parameters.
func WithChars(str string) func(*Config) {
	return func(conf *Config) {
		chars := removeDuplicates([]rune(str))
		conf.chars = chars
	}
}

// WithRand sets a source of random numbers.
func WithRand(r *rand.Rand) func(*Config) {
	return func(conf *Config) {
		conf.rand = r
	}
}

func removeDuplicates(chars []rune) []rune {
	res := make([]rune, 0, len(chars))
	enc := map[rune]bool{}
	for _, c := range chars {
		if !enc[c] {
			enc[c] = true
			res = append(res, c)
		}
	}
	return res
}
