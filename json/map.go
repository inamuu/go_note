package main

import (
	"fmt"
)

func main() {
	m := map[string]string{
		"key1" : "apple",
		"key2" : "banana",
		"key3" : "grape",
	}

	for key, value := range m {
		if key == "key3" {
			fmt.Println(value)
		}
	}
}