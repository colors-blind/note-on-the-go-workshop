package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int
	for i = 1; i <= 100; i++ {
		var res string
		res = ""
		if i%3 == 0 {
			res += "Fizz"
		}
		if i%5 == 0 {
			res += "Buzz"
		}
		if res == "" {
			res = strconv.Itoa(i)
		}
		fmt.Printf("%d->%s\n", i, res)
	}
}
