package main

import {
	"fmt"
	"net"
	"os"
}

func main()  {
	addr, err := net.ResolvIPAddr("ip", "inamuu.com")
	if err != nil {
		fmt.Println("error")
		os.Exit(1)
	}
	fmt.Println(addr)
}