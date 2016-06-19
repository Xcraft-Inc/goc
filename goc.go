package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func main() {
	var bin = os.Args[0]
	var ext = filepath.Ext(bin)
	var name = bin[0 : len(bin)-len(ext)]

	cmd := exec.Command(name+".cmd", os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	var waitStatus syscall.WaitStatus

	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			os.Exit(waitStatus.ExitStatus())
		} else {
			os.Exit(1)
		}
	}
}
