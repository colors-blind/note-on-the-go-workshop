package main

import "fmt"

var level = "PKG"

func main() {
	fmt.Println("vim-go")
	fmt.Println("Main start:", level)
	level = "changed!"
	if true {
		var level string
		level = "BLOCK"
		fmt.Println("Block start:", level)
	}
	funcA()
}

func funcA() {
	fmt.Println("funcA start:", level)
}
