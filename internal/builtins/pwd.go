package builtins

import (
	"fmt"
	"os"
)

func PWD() {
	file_path, err := os.Getwd()
	if err != nil {
		fmt.Println("[Error]: Getting working directory: ", file_path)
	} else {
		fmt.Println(file_path)
	}
}
