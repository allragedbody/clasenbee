package main

import (
	"fmt"
	"net"
	// "log"
	// "os"
	// "runtime"
	"os"
	"syscall"
	// "time"
)

var NetHandler *Server

type Server struct {
	r    *net.UnixAddr
	conn *net.UnixConn
}

func createChild() {
	var name string
	name = "gitclient"
	var args []string = make([]string, 4)
	args[0] = name

	cred := &syscall.Credential{}
	cred.Uid = uint32(1000)

	attrs := &os.ProcAttr{
		Env:   os.Environ(),
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Sys:   &syscall.SysProcAttr{Ptrace: false, Credential: cred, Setsid: true, Setpgid: false, Setctty: false},
	}

	if proc, err := os.StartProcess(name, args, attrs); err ==
		// if proc, err := os.StartProcess("/bin/ls", []string{"/"}, attrs); err ==
		nil {
		// proc.Wait() // <- tried with and without this Wait()
		_, err := proc.Wait()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func initContract() {
	var err error
	NetHandler = &Server{}
	NetHandler.r, err = net.ResolveUnixAddr("unixgram", "/tmp/gitremote.sock")
	if err != nil {
		fmt.Println("init gitlocal.sock error")
	}
	NetHandler.conn, err = net.ListenUnixgram("unixgram", NetHandler.r)
	if err != nil {
		fmt.Printf("DialUnix error")
	}

}

func main() {
	initContract()
	createChild()
	for {
		b := make([]byte, 64)
		_, from, err := NetHandler.conn.ReadFrom(b)
		fmt.Println(string(b) + "\n")
		if err != nil {
			fmt.Printf("ReadFrom error")
		}

		if from != nil {
			fmt.Println("from.name:", from)
		}
		fmt.Printf("before write\n")
		NetHandler.conn.WriteTo([]byte("response"), from)
		fmt.Printf("end write\n")
	}
}
