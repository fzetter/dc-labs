package main

import (
	"fmt"
  "math"
	"math/rand"
)

// ADD
func add(x int, y int) int {
	return x + y
}

// SWAP
func swap(x, y string) (string, string) {
	return y, x
}

// SPLIT
func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}

// VARIABLES
var c, python, java bool
var i, j int = 1, 2

// CONSTANTS
// Constants cannot be declared using the := syntax.
const Pi = 3.14

// MAIN
func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
  fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
  fmt.Println(math.Pi)

  // ADD
  fmt.Println(add(42, 13))

  // SWAP
  a, b := swap("hello", "world")
	fmt.Println(a, b)

  // SPLIT
  fmt.Println(split(17))

  // VARIABLES
  var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

  // Inside a function, the := short assignment statement can be
  // used in place of a var declaration with implicit type.
  k, l, m := 3, 4, 5
  fmt.Printf("Hello %v %v %T \n", k, l, m)

  // CONSTANTS
  fmt.Println("Happy", Pi, "Day")

}
