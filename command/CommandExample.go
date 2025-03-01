package main

import (
	"os/exec"
	"syscall"
	"fmt"
	"time"
	"os"
)

func main() {
	cmd := exec.Command("/home/user/GoWorkSpace/src/github.com/AndrienkoAleksandr/GoTerminal/command/myscript.sh")
	//cmd.SysProcAttr.Setpgid = true
	cmd.SysProcAttr =  &syscall.SysProcAttr{
		//Setsid:     true,
		//Setctty:    true,
		//Foreground: false,
		//Noctty:     true,
		Setpgid:    true,
	}
	var waitStatus syscall.WaitStatus
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		// Did the command fail because of an unsuccessful exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			fmt.Sprintf("%d", waitStatus.ExitStatus())
		}
	} else {
		// Command was successful

		timer := time.NewTimer(time.Second * 5)
		<-timer.C
		//waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
		//cmd.Process.Release();
		fmt.Sprintf("%d", waitStatus.ExitStatus())

		//fmt.Println(cmd.Process.Release())

		//SIGINT 
		//cmd.Process.Signal(syscall.SIGTERM)
		//syscall.Kill(cmd.Process.Pid, syscall.SIGTERM)
		err := cmd.Process.Kill()
		//syscall.Kill(cmd.Process.Pid, syscall.SIGKILL)
		if err != nil {
			fmt.Print(err)
		}
	}
}
