package main

import (
	"fmt"
	"net"
)

func main() {

	la, err := net.ResolveUnixAddr("unixgram", "/tmp/client.sock")
	localConn, err := net.ListenUnixgram("unixgram", la)
	if err != nil {
		fmt.Println("ListenUnixgram err")
	}
	// defer os.Remove(path)
	// defer conn.Close()

	coreAddr, err := net.ResolveUnixAddr("unixgram", "/tmp/server.sock")
	if err != nil {
		fmt.Println("ResolveUnixAddr err")
	}

	buf := make([]byte, 512*1024)
	buf = []byte{'a', 'b', 'c', 'd'}
	n, err := localConn.WriteTo(buf, coreAddr)
	if err != nil {
		fmt.Println("[core] write to unix socket err: %v", err)

	}

	fmt.Printf("write %d byte", n)

}
