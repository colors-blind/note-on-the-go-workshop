package main

import "fmt"

func BubbleSort(nums []int, length int) {
	for i := 0; i < length; i++ {
		for j := length - 1; j > i; j-- {
			if nums[j] < nums[j-1] {
				var temp int
				temp = nums[j-1]
				nums[j-1] = nums[j]
				nums[j] = temp
			}
		}
	}
}

func main() {
	numbers := []int{95, 45, 15, 78, -1, 51, -19, 12, -999}
	length := 9
	BubbleSort(numbers, length)
	fmt.Println(numbers)
}
