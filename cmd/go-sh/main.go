package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jonathon-chew/go-sh/internal/builtins"
	"github.com/jonathon-chew/go-sh/internal/utils"
)

func main() {
	for {
		os.Stdout.Write([]byte("$ "))
		// _, err := fmt.Scanln(&user_input)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			fmt.Printf("No user input, %s\n", err.Error())
			os.Exit(1)
		}
		var user_input = scanner.Text()

		var command string
		if strings.Contains(user_input, " ") {
			command = strings.Split(user_input, " ")[0]
		} else {
			command = user_input
		}

		// writeHistory(command)

		switch command {
		case "exit":
			os.Exit(0)
		case "echo":
			fmt.Println(strings.Replace(user_input, "echo", "", 1))
		case "cd":
			builtins.CD(user_input)
		case "type":
			builtins.TYPE(user_input)
		case "pwd":
			builtins.PWD()
		default:
			var args = strings.Split(user_input, " ")
			command = args[0]

			binary, binary_path, err := utils.Search_binaries(command)
			if err != nil {
				fmt.Printf("%s: command not found\n", command)
			}

			if binary != "" {
				fmt.Println("Binary found: ", binary, " ", binary_path)
			}

			utils.RunExternalBinary(command, args[1:])
		}
	}
}
