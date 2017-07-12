package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	max  = 100
	fizz = 3
	buzz = 5
)

func main() {
	var i int
	var list []string = make([]string, max+1)
	for i = 1; i <= max; i++ {
		list[i] = strconv.Itoa(i)
	}

	for i = 0; i < max; i = i + fizz {
		list[i] = "Fizz"
	}

	for i = 0; i <= max; i = i + buzz {
		var pre string

		_, err := strconv.Atoi(list[i])
		if err != nil {
			pre = list[i]
		} else {
			pre = ""
		}

		list[i] = pre + "Buzz"
	}

	fmt.Printf("%s", strings.Join(list, "\n"))
}
