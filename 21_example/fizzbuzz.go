package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("fizzbuzz\n")
		} else if i%3 == 0 {
			fmt.Print("fizz\n")
		} else if i%5 == 0 {
			fmt.Print("buzz\n")
		} else {
			fmt.Print(i, "\n")
		}
	}
}
