package randstr_test

import (
	"fmt"

	"github.com/scizorman/randstr"
)

func Example() {
	// Seeding with the same value results in the same random string each run.
	// The default seed is time.Now().UnixNano(), so you can get a different
	// strings if you do not call Seed.
	randstr.Seed(1)
	s := randstr.New(20)
	fmt.Println(s)
	// Output: CGHLF9EoWyo1KFHeio1r
}

func Example_number() {
	// Seeding with the same value results in the same random string each run.
	// The default seed is time.Now().UnixNano(), so you can get a different
	// strings if you do not call Seed.
	randstr.Seed(1)
	s := randstr.New(20, randstr.WithCharacters("1234567890"))
	fmt.Println(s)
	// Output: 97301642757693321863
}

func Example_withSpecialCharacters() {
	// Seeding with the same value results in the same random string each run.
	// The default seed is time.Now().UnixNano(), so you can get a different
	// strings if you do not call Seed.
	chars := `abcdefgihijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890@%+\/'!#$^?:(){}[]~-_`
	randstr.Seed(1)
	s := randstr.New(20, randstr.WithCharacters(chars))
	fmt.Println(s)
	// Output: x42dht51DXM\S]}Y!q%_
}
