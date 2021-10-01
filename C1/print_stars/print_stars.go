package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	seed := time.Now().UnixNano()
	fmt.Println(seed)
	rand.Seed(seed)
	r := rand.Intn(5) + 1
	stars := strings.Repeat("*", r)
	fmt.Println(stars)
}

