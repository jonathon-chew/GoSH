package builtins

import (
	"fmt"
	"os"
	"strings"
)

func CD(user_input string) {
	var desired_directory string = strings.Replace(user_input, "cd ", "", 1)
	file_info, ok := os.Lstat(desired_directory)
	if ok != nil {
		fmt.Println("Directory does not exist: ", desired_directory, ok.Error())
		return
	}

	if !file_info.IsDir() {
		fmt.Println(desired_directory, " is not a directory")
	} else {
		os.Chdir(desired_directory)
	}
}
