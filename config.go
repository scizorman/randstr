package randstr

// Config is the configuration to generate a random string.
type Config struct {
	chars []rune
}

// WithChars sets characters to use as a optional parameters.
func WithChars(str string) func(*Config) {
	return func(conf *Config) {
		chars := removeDuplicates([]rune(str))
		conf.chars = chars
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
