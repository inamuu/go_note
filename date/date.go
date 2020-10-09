package main

import (
	"fmt"
	"time"

	"github.com/leekchan/timeutil"
)

func main() {
	n := time.Now()
	fmt.Println(timeutil.Strftime(&n, "%Y-%m-%d"))
}