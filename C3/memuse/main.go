package main

import (
	"fmt"
	"runtime"
)

func main() {
	var numList []int

	for i := 0; i < 1000000; i++ {
		numList = append(numList, 100)
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("TotalAlloc(Heat) %v MiB\n", m.TotalAlloc/1024/1024)
}
