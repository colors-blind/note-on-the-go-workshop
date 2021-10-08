package main

import "fmt"

func main() {
	array := []int{5, 4, 3, 2, 1}
	for index := range array {
		fmt.Println(index, array[index])
	}
	fmt.Println("************")
	slice := make([]int, 3, 5)
	slice = append(slice, 100)
	for index := range slice {
		fmt.Println(index, slice[index])
	}
}
