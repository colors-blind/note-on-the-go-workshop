package main

import "fmt"
import "time"

func FakeChangeValue(value int64) {
	value = 0
}

func RealChangeValue(value *int64) {
	*value = 0
}

func main() {
	var p_int1 *int64
	var num int64
	num, p_int1 = 10, &num
	fmt.Printf("p_int1 is %#v type is %T\n", *p_int1, p_int1)

	p_int2 := new(float64)
	fmt.Printf("p_int2 is %#v type is %T\n", *p_int2, p_int2)

	time_p := &time.Time{}
	fmt.Printf("time_p is %#v\n", time_p)

	FakeChangeValue(num)
	fmt.Printf("num not change after call FakeChangeValue %v\n", num)

	RealChangeValue(p_int1)
	fmt.Printf("num changed after call RealChangeValue %v\n", num)
}

