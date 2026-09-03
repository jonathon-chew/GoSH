package cat

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func moreCommandLineFlags(array []string) bool {

	if len(array) > 2 {
		for _, str := range array[1:] {
			if strings.Contains(str, "-") {
				// log.Print(str)
				return true
			}
		}
	}

	return false
}

func cmd(arguments []string) {
	/*
	   		-b      Number the non-blank output lines, starting at 1.
	        -e      Display non-printing characters (see the -v option), and display a dollar sign
	                (‘$’) at the end of each line.
	        -l      Set an exclusive advisory lock on the standard output file descriptor.  This
	                lock is set using fcntl(2) with the F_SETLKW command.  If the output file is
	                already locked, cat will block until the lock is acquired.
	        -n      Number the output lines, starting at 1.
	        -s      Squeeze multiple adjacent empty lines, causing the output to be single spaced.
	        -t      Display non-printing characters (see the -v option), and display tab
	                characters as ‘^I’.
	        -u      Disable output buffering.
	        -v      Display non-printing characters so they are visible.  Control characters print
	                as ‘^X’ for control-X; the delete character (octal 0177) prints as ‘^?’.
	                Non-ASCII characters (with the high bit set) are printed as ‘M-’ (for meta)
	                followed by the character for the low 7 bits.
	*/

	var files = []string{}

	for numberOfArguments, argument := range arguments {
		switch argument {
		case "-b":
			// log.Print("Argument was -b...")
			if numberOfArguments == len(arguments) {
				for _, file := range files {
					fileContents, ErrReadFile := os.ReadFile(file)
					if ErrReadFile != nil {
						fmt.Printf("[ERROR]: %s\n", ErrReadFile)
					}
					newFileContents := printWithLineNumbersNoBlanks(string(fileContents))
					fmt.Print(strings.Join(newFileContents, "\n"))
				}
			} else if numberOfArguments != len(arguments) && !moreCommandLineFlags(arguments[numberOfArguments:]) {
				for _, file := range arguments[numberOfArguments+1:] {
					files = append(files, file)
				}

				for _, file := range files {
					fileContents, ErrReadFile := os.ReadFile(file)
					if ErrReadFile != nil {
						fmt.Printf("[ERROR]: %s\n", ErrReadFile)
					}
					newFileContents := printWithLineNumbersNoBlanks(string(fileContents))
					fmt.Print(strings.Join(newFileContents, "\n"))
				}
			}

		case "-e":
			// log.Print("Argument was -e...")
			if numberOfArguments == len(arguments) {
				// log.Print("Number of arguments seen matches total arguments...")
				for _, file := range files {
					fileContents, ErrReadFile := os.ReadFile(file)
					if ErrReadFile != nil {
						fmt.Printf("[ERROR]: %s\n", ErrReadFile)
					}
					newFileContents := printWithDollarAtTheEnd(string(fileContents))
					fmt.Print(newFileContents)
				}
			} else if numberOfArguments != len(arguments) && !moreCommandLineFlags(arguments[numberOfArguments:]) {
				// log.Print("Number of arguments seen does not matches total arguments and there are no more command line arguments...")
				for _, file := range arguments[numberOfArguments+1:] {
					files = append(files, file)
				}

				for _, file := range files {
					fileContents, ErrReadFile := os.ReadFile(file)
					if ErrReadFile != nil {
						fmt.Printf("[ERROR]: %s\n", ErrReadFile)
					}
					newFileContents := printWithDollarAtTheEnd(string(fileContents))
					fmt.Print(newFileContents)
				}
			} else if moreCommandLineFlags(arguments[numberOfArguments:]) {
				// log.Print("There are more command line flags...")
			}

		case "-l":
			fmt.Println("Argument was -l...")
			continue
		case "-n":
			// log.Print("Argument was -n...")
			if numberOfArguments == len(arguments) {
				for _, file := range files {
					fileContents, ErrReadFile := os.ReadFile(file)
					if ErrReadFile != nil {
						fmt.Printf("[ERROR]: %s\n", ErrReadFile)
					}
					newFileContents := printWithLineNumbers(string(fileContents))
					fmt.Print(strings.Join(newFileContents, "\n"))
				}
			}

		case "-s":
			// log.Print("Argument was -s...")
			if numberOfArguments == len(arguments) {
				for _, file := range files {
					fileContents, ErrReadFile := os.ReadFile(file)
					if ErrReadFile != nil {
						fmt.Printf("[ERROR]: %s\n", ErrReadFile)
					}
					newFileContents := printExcludeMultipleEmptyLines(string(fileContents))
					fmt.Print(strings.Join(newFileContents, "\n"))
				}
			}

		case "-t":
			continue
		case "-u":
			continue
		case "-v":
			continue
		default:
			files = append(files, argument)

			if numberOfArguments != len(arguments) && !moreCommandLineFlags(arguments[numberOfArguments:]) {
				for _, file := range arguments[numberOfArguments+1:] {
					files = append(files, file)
				}

				for _, file := range files {
					fileContents, ErrReadFile := os.ReadFile(file)
					if ErrReadFile != nil {
						fmt.Printf("[ERROR]: %s\n", ErrReadFile)
					}
					fmt.Print(string(fileContents))
				}
			}
		}
	}

}

func printWithLineNumbersNoBlanks(file string) []string {

	var filePrint []string
	splitFile := strings.Split(file, "\n")

	for index, line := range splitFile {
		if line != "\n" && line != "" {
			filePrint = append(filePrint, strconv.Itoa(index)+" "+line)
		}
	}

	return filePrint
}

func printWithDollarAtTheEnd(file string) string {

	dollarLineFile := strings.ReplaceAll(file, "\n", "$\n")

	return dollarLineFile

}

func printWithLineNumbers(file string) []string {

	var filePrint []string
	splitFile := strings.Split(file, "\n")

	for index, line := range splitFile {
		filePrint = append(filePrint, strconv.Itoa(index)+" "+line)
	}

	return filePrint
}

func printExcludeMultipleEmptyLines(file string) []string {

	var filePrint []string
	splitFile := strings.Split(file, "\n")

	for _, line := range splitFile {
		if line != "\n" && line != "" {
			filePrint = append(filePrint, line)
		}
	}

	return filePrint

}

func Cat() {
	// log.Print("Running cat...")
	if len(os.Args) > 1 {
		// log.Print("More than one argument found...")
		cmd(os.Args[1:])
	} else {
		fmt.Println("[ERROR]: No file")
	}
}
