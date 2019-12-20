package randstr_test

import (
	"fmt"

	"github.com/scizorman/randstr"
)

func Example() {
	// Each run gives different results.
	s := randstr.New(10)
	fmt.Println(s)
}

func Example_number() {
	// Each run gives different results.
	s := randstr.New(10, randstr.WithCharacters("1234567890"))
	fmt.Println(s)
}
