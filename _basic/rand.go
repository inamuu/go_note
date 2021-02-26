package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	list := []string{
		"taro",
		"hanako",
		"yoko",
		"jiro",
		"hanao",
		"nao",
		"kaz",
		"yusuke",
		"you",
		"waka",
		"miu",
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1; i++ {
		num := rand.Intn(len(list))
		fmt.Println(list[num])
	}
}
