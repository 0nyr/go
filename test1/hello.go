package main

// list of imports
import (
	"fmt"

	"rsc.io/quote"
)

// ommit all but last argument type if similar
func add(x, y int) int {
	return x + y
}

// A function can return any number of results
func swap(x, y string) (string, string) {
	return y, x
}

// Named return values ("naked return")
func split(sum int) (x, y int) { // Go's return values may be named.
	x = sum * 4 / 9
	y = sum - x
	return // A return statement without arguments returns the named return values.
}

// list of global variables
var (
	ToBe   bool   = false
	MaxInt uint64 = 1<<64 - 1
	i, j   int    = 1, 2
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	fmt.Println(ToBe, i, j)

	// short variable declaration
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	// type inference (only for local variables)
	var k = 3 // long declaration
	l := 4    // short declaration
	fmt.Println(k, l)

	// zero values by default
	var i1 int
	var f1 float64
	var b1 bool
	var s1 string
	fmt.Printf("%v %v %v %q\n", i1, f1, b1, s1)

	// constants
	// WARN: cannot use := for constants
	const Pi = 3.14
	const Truth = true
	fmt.Println("Happy", Pi, "Day", Truth)
}
