package randstr_test

import (
	"fmt"
	"math/rand"

	"github.com/scizorman/randstr"
)

func Example() {
	// For testing, set the seed to 1.
	// So please use it without `randstr.WithRand()` if you do not set a specific seed.
	r := rand.New(rand.NewSource(1))
	s := randstr.New(20, randstr.WithRand(r))
	fmt.Println(s)
	// Output: CGHLF9EoWyo1KFHeio1r
}

func Example_number() {
	// For testing, set the seed to 1.
	// So please use it without `randstr.WithRand()` if you do not set a specific seed.
	r := rand.New(rand.NewSource(1))
	s := randstr.New(20, randstr.WithChars("1234567890"), randstr.WithRand(r))
	fmt.Println(s)
	// Output: 97301642757693321863
}

func Example_with_special_characters() {
	// For testing, set the seed to 1.
	// So please use it without `randstr.WithRand()` if you do not set a specific seed.
	chars := `abcdefgihijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890@%+\/'!#$^?:(){}[]~-_`
	r := rand.New(rand.NewSource(1))
	s := randstr.New(20, randstr.WithChars(chars), randstr.WithRand(r))
	fmt.Println(s)
	// Output: x42dht51DXM\S]}Y!q%_
}
