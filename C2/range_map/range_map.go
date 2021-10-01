package main

import "fmt"

func main() {
	config := map[string]string{
		"debug":    "1",
		"logLevel": "info",
		"version":  "1.1.1",
	}
	for key, value := range config {
		fmt.Printf("%s -> %s\n", key, value)
	}
}
