package main

import "fmt"

func multipleArgs(arg1, arg2 string) (string, string) {
	return arg1, arg2
}

func main() {
	a, b := multipleArgs("Hello", "World")
	fmt.Println(a, b)
}
