package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jonathon-chew/GoSH/internal/builtins"
	"github.com/jonathon-chew/GoSH/internal/ls"
	"github.com/jonathon-chew/GoSH/internal/utils"
)

func getUserInput() (string, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		fmt.Printf("No user input, %s\n", err.Error())
		os.Exit(1)
	}
	var user_input = scanner.Text()

	split_input := strings.Split(user_input, " ")
	var command string = split_input[0]
	var args []string

	if len(split_input) > 1 {
		args = split_input[1:]
	}

	return command, args
}

func main() {

	for {
		os.Stdout.Write([]byte("$ "))
		command, args := getUserInput()
		utils.WriteHistory(command)

		switch command {
		case "exit":
			os.Exit(0)
		case "echo":
			fmt.Println(strings.Replace(strings.Join(args, " "), "echo", "", 1))
		case "cd":
			builtins.CD(strings.Join(args, " "))
		case "ls":
			var files []string
			files = ls.Ls(args)
			if len(files) == 0 {
				break
			}
			for _, file := range files {
				fmt.Println(file)
			}
		case "pwd":
			builtins.PWD()
		case "type":
			builtins.TYPE(strings.Join(args, " "))
		default:
			binary, binary_path, err := utils.Search_binaries(command)
			if err != nil {
				fmt.Printf("%s: command not found\n", command)
			}

			if binary != "" {
				fmt.Println("Binary found: ", binary, " ", binary_path)
			}

			utils.RunExternalBinary(command, args)
		}
	}
}
