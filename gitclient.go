package main

import (
	"fmt"
	"net"
	// "log"
	"os"
	// "runtime"
	// "time"
	"encoding/json"
	// "strings"
	"bytes"
)

type GitClient struct {
	User        string `json:user`
	CommandLine string `json:commandline`
}

// var NetHandler *Client

// type Client struct {
//  l    *net.UnixAddr
//  r    *net.UnixAddr
//  conn *net.UnixConn
// }

// func initContract() {
//  var err error
//  NetHandler = &Client{}
//  NetHandler.l, err = net.ResolveUnixAddr("unixgram", "/tmp/gitclient.sock")
//  if err != nil {
//      fmt.Println("init gitlocal.sock error")
//  }
//  NetHandler.r, err = net.ResolveUnixAddr("unixgram", "/tmp/gitserver.sock")
//  if err != nil {
//      fmt.Println("init gitlocal.sock error")
//  }
//  defer os.Remove("/tmp/gitlocal.sock")
//  // func ListenUnixgram(net string, laddr *UnixAddr) (*UnixConn, error)

//  NetHandler.conn, err = net.DialUnix("unixgram", NetHandler.l, NetHandler.r)
//  if err != nil {
//      fmt.Println("init connection error")
//  }

//  for {

//      NetHandler.conn.Write([]byte("hello"))
//      defer func() {
//          NetHandler.conn.Close()
//      }()

//      time.Sleep(time.Second)
//      b := make([]byte, 64)
//      _, _, err2 := NetHandler.conn.ReadFrom(b)
//      fmt.Println(string(b) + "\n")
//      if err2 != nil {
//          fmt.Printf("ReadFrom error\n")
//      }
//      fmt.Printf("read:[%s]\n", string(b))
//  }

// }

// // func createContract() {

// // if err != nil {
// //   fmt.Println("ListenUnixgram err")
// // }

// // coreAddr, err := net.ResolveUnixAddr("unixgram", "/tmp/server.sock")
// // if err != nil {
// //   fmt.Println("ResolveUnixAddr err")
// // }

// // buf := make([]byte, 512*1024)
// // buf = []byte{'a', 'b', 'c', 'd'}
// // n, err := localConn.WriteTo(buf, coreAddr)
// // if err != nil {
// //   fmt.Printf("[core] write to unix socket err: %v", err)

// // }

// // fmt.Printf("write %d byte", n)
// // buf1 := make([]byte, 512*1024)
// // aa, bb, cc := localConn.ReadFrom(buf1)
// // fmt.Println("111111111111111", aa, bb, cc)
// // }
// //
func main() {
	// initContract()
	//
	//
	ta, err := net.ResolveUnixAddr("unixgram", "/tmp/unixdomaincli")
	c1, err := net.ListenUnixgram("unixgram", ta)
	if err != nil {
		fmt.Printf("DialUnix error")
	}
	defer os.Remove("/tmp/unixdomaincli2")

	fmt.Println("Start git child and waiting commands")

	for {
		body := make([]byte, 1000)
		_, from, err := c1.ReadFrom(body)
		if err != nil {
			fmt.Printf("ReadFrom error")
		}
		if from != nil {
			fmt.Println("from.name:", from)
		}

		index := bytes.IndexByte(body, 0)
		body1 := body[0:index]

		var gc *GitClient
		err = json.Unmarshal([]byte(string(body1)), &gc)
		if err != nil {
			fmt.Println("error:", err)
		}

		ret := fmt.Sprintf("%v run %v is Success", gc.User, gc.CommandLine)

		c1.WriteTo([]byte(ret), from)

	}
}
