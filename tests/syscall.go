package main

// package tests
import (
	"fmt"
	// "log"
	// "os"
	// "runtime"
	"os"
	"syscall"
	"testing"
	"time"
)

//将整个用户变成另外一个用户执行程序
// func daemon() {
// 	time.Sleep(1000)
// 	fmt.Println("start daemon")
// 	fmt.Println(syscall.Getuid())
// 	syscall.RawSyscall(syscall.SYS_SETUID, uintptr(18), 0, 0)
// 	//fmt.Println(syscall.Setuid(14))
// 	fmt.Println(syscall.Getuid())
// 	os.MkdirAll("/tmp/111", 644)
// 	err := os.Remove("/tmp/111")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	// syscall.RawSyscall(syscall.SYS_SETUID, uintptr(0), 0, 0)
// 	err = os.Remove("/tmp/111")
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("remove 111 success")
// 	}

// }

func daemon() {
	var name string
	name = "/project/src/jd.com/cdn/gateway/gateway"
	var args []string = make([]string, 5)
	args[0] = "gateway"
	args[1] = "-conf"
	args[2] = "/project/src/jd.com/cdn/gateway/config.json"
	args[3] = "-imp"
	args[4] = "/project/src/jd.com/cdn/gateway/important.properties"
	cred := &syscall.Credential{}
	cred.Uid = uint32(12)

	attrs := &os.ProcAttr{
		Env:   os.Environ(),
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Sys:   &syscall.SysProcAttr{Ptrace: true, Credential: cred},
	}
	if proc, err := os.StartProcess(name, args, attrs); err ==
		// if proc, err := os.StartProcess("/bin/ls", []string{"/"}, attrs); err ==
		nil {

		// proc.Wait() // <- tried with and without this Wait()
		_, err := proc.Wait()
		if err != nil {
			fmt.Println(err)
		}

		foo := syscall.PtraceAttach(proc.Pid)
		fmt.Printf("Started: %v\n", proc.Pid)
		fmt.Printf("PtraceAttach result: %v\n", foo)

	}
}

///project/src/jd.com/cdn/gateway/gateway -conf /project/src/jd.com/cdn/gateway/config.json -imp /project/src/jd.com/cdn/gateway/important.properties
func TestSyscall(t *testing.T) {
	// daemon(0, 1)
	go daemon()

	fmt.Println("start")

	// _, err = proc.Wait()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// foo := syscall.PtraceAttach(proc.Pid)
	// fmt.Printf("Started: %v\n", proc.Pid)
	// fmt.Printf("PtraceAttach result: %v\n", foo)

}

///project/src/jd.com/cdn/gateway/gateway -conf /project/src/jd.com/cdn/gateway/config.json -imp /project/src/jd.com/cdn/gateway/important.properties
func main() {
	// daemon(0, 1)
	go daemon()

	fmt.Println("start")
	time.Sleep(10000 * time.Millisecond)
}
