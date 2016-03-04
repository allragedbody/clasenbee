package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	ta, err := net.ResolveUnixAddr("unixgram", "/tmp/unixdomaincli")
	if err != nil {
		fmt.Printf("ResolveUnixAddr error\n")
	}

	sa, err1 := net.ResolveUnixAddr("unixgram", "/tmp/unixdomaincli2")
	if err1 != nil {
		fmt.Printf("ResolveUnixAddr error\n")
	}

	c2, err := net.DialUnix("unixgram", sa, ta)
	if err != nil {
		fmt.Printf("DialUnix error\n")
	}
	for {
		c2.Write([]byte("hello"))
		defer func() {
			c2.Close()
		}()

		time.Sleep(time.Second)
		b := make([]byte, 64)
		_, _, err2 := c2.ReadFrom(b)
		fmt.Println(string(b) + "\n")
		if err2 != nil {
			fmt.Printf("ReadFrom error\n")
		}
		fmt.Printf("read:[%s]\n", string(b))
	}
	os.Remove("/tmp/unixdomaincli2")

}
