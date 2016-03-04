package main

import (
	"fmt"
	"net"
)

func main() {
	readConn, err := net.ListenUnixgram("unixgram", &net.UnixAddr{"/tmp/server.sock", "unixgram"})
	if err != nil {
		fmt.Println("ListenUnixgram error")
	}

	buf := make([]byte, 512*1024)

	n, _, err := readConn.ReadFrom(buf)

	fmt.Println("buf=", string(buf[:n]))
}
