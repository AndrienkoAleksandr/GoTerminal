package main

import "fmt"
import (
	"os"
	"syscall"
)
//import "syscall"


func main() {
	cmdToRun := "/bin/sleep"
	args := []string{"--help"}
	procAttr := new(os.ProcAttr)
	procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
	process, err := os.StartProcess(cmdToRun, args, procAttr);
	//process.Signal(syscall.SIGTERM)
	if err == nil {
		fmt.Print("Running with pid ", process.Pid)

		// It is not clear from docs, but Realease actually detaches the process
		err = process.Release();
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("Erorr " + err.Error())
	}
}