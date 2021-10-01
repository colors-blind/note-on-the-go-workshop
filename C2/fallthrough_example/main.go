package main

import "fmt"

func main() {
	num := 10
	gtTen := num > 10
	switch gtTen {
	case false:
		fmt.Println("num > 5")
		fallthrough
	case false:
		fmt.Println("num > 6")
		fallthrough
	case true:
		fmt.Println("num > 7")
		fallthrough
	case false:
		fmt.Println("num > 8")
	}
}
