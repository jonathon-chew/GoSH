package utils

import (
	"fmt"
	"os"
	"os/exec"
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
