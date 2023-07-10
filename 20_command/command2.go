package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("ls", "noexists").CombinedOutput()
	if err != nil {
		fmt.Printf("error!  ")
	}
	fmt.Printf("%s", string(out))
}
