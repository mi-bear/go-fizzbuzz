package main

import "fmt"

var (
	max  = 100
	fizz = 3
	buzz = 5
)

func main() {
	i := 1
	f := fizz
	b := buzz

	for i <= max {
		var fb bool

		if i == f {
			f += fizz
			fb = true
			fmt.Printf("%s", "Fizz")
		}

		if i == b {
			b += buzz
			fb = true
			fmt.Printf("%s", "Buzz")
		}

		if !fb {
			fmt.Printf("%d", i)
		}

		fmt.Printf("%s", "\n")
		i++
	}
}
