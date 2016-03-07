package gitcontrol

import (
	"fmt"
	"net"
	// "log"
	// "os"
	// "runtime"
	"encoding/json"
	"os"
	"syscall"
	"time"
)

var NetHandler *Server

type Server struct {
	r    *net.UnixAddr
	conn *net.UnixConn
}

type GitClient struct {
	User        string `json:user`
	CommandLine string `json:commandline`
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
		Sys:   &syscall.SysProcAttr{Ptrace: false, Credential: cred, Setsid: false, Setpgid: false, Setctty: false},
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

func InitGitServe() {
	// initContract()
	//
	//
	os.Remove("/tmp/unixdomaincli")
	os.Remove("/tmp/unixdomaincli2")

	go createChild()
	time.Sleep(2000 * time.Millisecond)
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
	os.Chown("/tmp/unixdomaincli2", 1000, 0)

	for {
		time.Sleep(time.Second)
		gc := &GitClient{}
		gc.User = "dcms"
		gc.CommandLine = "git pull origin master"

		body, err := json.Marshal(gc)
		if err != nil {
			panic(err.Error())
		}

		c2.Write(body)
		defer func() {
			c2.Close()
		}()

		b := make([]byte, 1000)
		_, _, err2 := c2.ReadFrom(b)

		if err2 != nil {
			fmt.Printf("ReadFrom error\n")
		}
		fmt.Printf("read:[%s]\n", string(b))
	}

	// time.Sleep(22000 * time.Millisecond)
	// for {

}
