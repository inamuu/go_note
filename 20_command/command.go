package main

import (
	"fmt"
	"os/exec"
)

func main() {
	//fmt.Println("処理開始:", time.Now())
	out, err := exec.Command("ls", "-la").Output()
	if err != nil {
		return
	}
	fmt.Println(string(out))
}
