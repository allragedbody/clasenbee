package main

import (
	"fmt"
	"net"
	//"os"
)

func main() {
	ta, err := net.ResolveUnixAddr("unixgram", "/tmp/unixdomaincli")
	c1, err := net.ListenUnixgram("unixgram", ta)
	if err != nil {
		fmt.Printf("DialUnix error")
	}
	fmt.Printf("-------------------")

	for {
		b := make([]byte, 64)
		_, from, err := c1.ReadFrom(b)
		fmt.Println(string(b) + "\n")
		if err != nil {
			fmt.Printf("ReadFrom error")
		}

		if from != nil {
			fmt.Println("from.name:", from)
		}
		fmt.Printf("before write\n")
		c1.WriteTo([]byte("response"), from)
		fmt.Printf("end write\n")
	}
}
