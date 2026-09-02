package cat

import (
	"bytes"
	"log"
	"os"
	"os/exec"
)

const (
	testFile string = "/Users/hunteradder626/Projects/Go/Learning/bash_commands/tests/example.txt"
)

func runCommand(option string) {
	os.Args = append(os.Args, option, testFile)
	cmd(os.Args[1:])
	os.Args = []string{os.Args[0]}
}

func Test_cat() {

	options := []string{"-b", "-e", "-l", "-n", "-s", "-t", "-u", "-v"}

	for _, option := range options {
		runCommand(option)

		cmd := exec.Command("cat", option, testFile)

		var outStr bytes.Buffer
		cmd.Stdout = &outStr

		ErrRunningCommand := cmd.Run()
		if ErrRunningCommand != nil {
			log.Fatal(ErrRunningCommand)
			return
		}

	}
}
