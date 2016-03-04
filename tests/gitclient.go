package main

import (
	"fmt"
	"net"
	// "log"
	"os"
	// "runtime"
	"time"
)

var NetHandler *Client

type Client struct {
	l    *net.UnixAddr
	r    *net.UnixAddr
	conn *net.UnixConn
}

func initContract() {
	var err error
	NetHandler = &Client{}
	NetHandler.l, err = net.ResolveUnixAddr("unixgram", "/tmp/gitlocal.sock")
	if err != nil {
		fmt.Println("init gitlocal.sock error")
	}
	NetHandler.r, err = net.ResolveUnixAddr("unixgram", "/tmp/gitremote.sock")
	if err != nil {
		fmt.Println("init gitlocal.sock error")
	}

	// func ListenUnixgram(net string, laddr *UnixAddr) (*UnixConn, error)

	NetHandler.conn, err = net.DialUnix("unixgram", NetHandler.l, NetHandler.r)
	if err != nil {
		fmt.Println("init connection error")
	}
}

// func createContract() {

// if err != nil {
// 	fmt.Println("ListenUnixgram err")
// }

// coreAddr, err := net.ResolveUnixAddr("unixgram", "/tmp/server.sock")
// if err != nil {
// 	fmt.Println("ResolveUnixAddr err")
// }

// buf := make([]byte, 512*1024)
// buf = []byte{'a', 'b', 'c', 'd'}
// n, err := localConn.WriteTo(buf, coreAddr)
// if err != nil {
// 	fmt.Printf("[core] write to unix socket err: %v", err)

// }

// fmt.Printf("write %d byte", n)
// buf1 := make([]byte, 512*1024)
// aa, bb, cc := localConn.ReadFrom(buf1)
// fmt.Println("111111111111111", aa, bb, cc)
// }
//
func main() {
	initContract()
	// createContract()
	for {
		time.Sleep(1000 * time.Millisecond)
		_, err := os.Stat("/tmp/gitserver.sock")
		if err != nil {
			continue
		} else {
			break
		}
	}

	for {

		NetHandler.conn.Write([]byte("hello"))
		defer func() {
			NetHandler.conn.Close()
		}()

		time.Sleep(time.Second)
		b := make([]byte, 64)
		_, _, err2 := NetHandler.conn.ReadFrom(b)
		fmt.Println(string(b) + "\n")
		if err2 != nil {
			fmt.Printf("ReadFrom error\n")
		}
		fmt.Printf("read:[%s]\n", string(b))
	}
	os.Remove("/tmp/gitlocal.sock")
}
