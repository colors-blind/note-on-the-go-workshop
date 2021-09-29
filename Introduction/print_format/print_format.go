package main

import "time"
import "fmt"

type Person struct {
	Name string
	Age int
}

func PrintDiff(value interface{}) {
	fmt.Printf("%v %+v %#v %T\n", value, value, value, value)
}

func main() {
	var name string
	name = "test"
	PrintDiff(name)

	var tom = Person{
		Name: "tom",
		Age: 5,
	}

	PrintDiff(tom)

	var ctime time.Time
	PrintDiff(ctime)
}
