package main

import (
  "fmt"
  "math"
  "runtime"
  "time"
)

// POW
func pow(x, n, lim float64) float64 {

  // Like for, the if statement can start with
  // a short statement to execute before the condition.

  // Variables declared inside an if short statement
  // are also available inside any of the else blocks.
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// MAIN
func main() {

  // LOOPS
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

  // OPTIONAL INIT AND POST-STATEMENT
  sum = 1
	for ; sum < 1000; {
		sum += sum
	}

	fmt.Println(sum)

  // WHILE LOOP IS A FOR
  sum = 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)

  // POW
  fmt.Println(
    pow(3, 2, 10),
    pow(3, 3, 20),
  )

  // SWITCH
  // the break statement that is needed at the end of each
  // case in those languages is provided automatically in Go.
  fmt.Print("Go runs on ")

  switch os := runtime.GOOS; os {
    case "darwin": fmt.Println("OS X.")
    case "linux": fmt.Println("Linux.")
    default: fmt.Printf("%s.\n", os)
  }

  // ***

  fmt.Print("When's Saturday? ")
  today := time.Now().Weekday()

  switch time.Saturday {
    case today + 0: fmt.Println("Today.")
    case today + 1: fmt.Println("Tomorrow.")
    case today + 2: fmt.Println("In two days.")
    default: fmt.Println("Too far away.")
  }

  // ***

  t := time.Now()

  switch {
    case t.Hour() < 12: fmt.Println("Good morning!")
    case t.Hour() < 17: fmt.Println("Good afternoon.")
    default: fmt.Println("Good evening.")
  }

  // FOR EACH
  // for <key>, <value> := range <container>{}
  mapp := map[int]string {1: "one", 2: "two", 3: "three"}

  for integ, spell := range mapp {
    fmt.Println(integ, " = ", spell)
  } 

  // DEFER
  // A defer statement defers the execution of a
  // function until the surrounding function returns.
  defer fmt.Println("world")
	fmt.Print("hello ")

}
