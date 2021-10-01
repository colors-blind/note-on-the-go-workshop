package main

import "fmt"

func main() {
	for i := -10; i < 10; i++ {
		if i&1 == 0 {
			fmt.Printf("%d is even\n", i)
		} else {
			fmt.Printf("%d is odd\n", i)
		}
	}
}
