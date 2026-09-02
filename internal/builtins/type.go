package builtins

import (
	"fmt"
	"slices"
	"strings"
)

func TYPE(user_input string) {
	var binary = strings.Split(user_input, " ")[1]
	binary_look_up, binary_path, err := utils.search_binaries(binary)
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
}
