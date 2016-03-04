package main

import (
	"fmt"
	// "log"
	// "os"
	// "runtime"
	"os"
	"syscall"
	"time"
)

var ch chan int

func createChild() {
	var name string
	name = "newprocess"
	var args []string = make([]string, 4)
	args[0] = name
	args[1] = "child"
	// attrs := &os.ProcAttr{Sys: &syscall.SysProcAttr{Ptrace: true}}
	// proc, err := os.StartProcess(name, args, attrs)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	cred := &syscall.Credential{}
	cred.Uid = uint32(12)

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

		// foo := syscall.PtraceAttach(proc.Pid)
		// fmt.Printf("Started: %v\n", proc.Pid)
		// fmt.Printf("PtraceAttach result: %v\n", foo)
		<-ch
	}
}

func ChildProcess() {
	time.Sleep(5000 * time.Millisecond)
	fmt.Printf("Child Pid %v \n", os.Getpid())
	fmt.Println(proc)
	os.MkdirAll("/tmp/112112", 644)
}

func ParentProcess() {
	time.Sleep(5000 * time.Millisecond)
	fmt.Printf("Parent Pid %v\n", os.Getpid())

	os.MkdirAll("/tmp/1121121111", 644)
}

func main() {
	// daemon(0, 1)
	// go daemon()
	//
	//
	if syscall.Getuid() == 0 {
		fmt.Println("Start Primary Process!!!")
		go createChild()
	} else {
		fmt.Println("Start Child Process!!!")
		for {
			ChildProcess()
		}
		ch <- 1
	}
	for {
		ParentProcess()
	}
}
