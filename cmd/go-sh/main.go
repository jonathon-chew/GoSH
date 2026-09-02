package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

func pathExists(path string) bool {
	_, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return true
}

func pathJoin(curr_path string, add_path string) string {
	return fmt.Sprintf("%s%s%s", curr_path, string(os.PathSeparator), add_path)
}

func IsExecAny(mode os.FileMode) bool {
	return mode&0111 != 0
}

func search_binaries(binary string) (string, string, error) {

	var path = os.Getenv("PATH")
	if path == "" {
		return "", "", fmt.Errorf("Unable to find the PATH env variable")
	}

	for folder := range strings.SplitSeq(path, string(os.PathListSeparator)) {
		if pathExists(pathJoin(folder, binary)) {
			file_info, err := os.Lstat(pathJoin(folder, binary))
			if err != nil {
				return "", "", fmt.Errorf("Unable to get information on the binary")
			}
			if IsExecAny(file_info.Mode()) {
				return binary, pathJoin(folder, binary), nil
			}
		}
	}
	return "", "", fmt.Errorf("Unable to find the binary")
}

func writeHistory(command string) {
	fp, err := os.Open("$HOME/.gosh_history")
	if err != nil {
		os.Exit(1)
	}

	fp.Write([]byte(command))
}

func runExternalBinary(executable string, args []string) {
	cmd := exec.Command(executable, args...)
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("[Error]: Getting working directroy", err.Error())
	}
	cmd.Dir = cwd

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("[Error]: from binary", err.Error())
	}
	if len(output) > 0 {
		fmt.Println(string(output))
	}
}

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
			var desired_directory string = strings.Replace(user_input, "cd ", "", 1)
			file_info, ok := os.Lstat(desired_directory)
			if ok != nil {
				fmt.Println("Directory does not exist: ", desired_directory, ok.Error())
				break
			}

			if !file_info.IsDir() {
				fmt.Println(desired_directory, " is not a directory")
			} else {
				os.Chdir(desired_directory)
			}
		case "type":
			var binary = strings.Split(user_input, " ")[1]
			binary_look_up, binary_path, err := search_binaries(binary)
			if err != nil {
				fmt.Printf("%s: not found\n", binary)
			}

			if len(strings.Split(user_input, " ")) <= 1 {
				fmt.Println("no command found")
			} else if slices.Contains([]string{"exit", "type", "echo", "pwd"}, binary) {
				fmt.Printf("%s is a shell builtin\n", binary)
			} else if binary == binary_look_up {
				fmt.Printf("%s is in %s\n", binary, binary_path)
			}
		case "pwd":
			file_path, err := os.Getwd()
			if err != nil {
				fmt.Println("[Error]: Getting working directory: ", file_path)
			} else {
				fmt.Println(file_path)
			}
		default:
			var args = strings.Split(user_input, " ")
			command = args[0]

			binary, binary_path, err := search_binaries(command)
			if err != nil {
				fmt.Printf("%s: command not found\n", command)
			}

			if binary != "" {
				fmt.Println("Binary found: ", binary, " ", binary_path)
			}

			runExternalBinary(command, args[1:])
		}
	}
}
