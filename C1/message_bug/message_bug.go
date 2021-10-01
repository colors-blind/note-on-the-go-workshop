package main

import "fmt"

func main() {
	var msg string
	count := 5
	if count > 5 {
		msg = "Greater than 5"
	} else {
		msg = "Not greater than 5"
	}
	fmt.Println(msg)
}
