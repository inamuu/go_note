package main

import "fmt"

func testfunc1(x int) string {
	y := x + 1
	return fmt.Sprint(y)
}

func main() {
	var example = 1
	fmt.Println(testfunc1(example))
}
