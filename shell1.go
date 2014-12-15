//run:
//go run shell1.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Welcome to Shell1 [Version 0.0.1]\n\n")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Shell1> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")

		if "exit" == input {
			break
		} else if "" == input {
		} else {
			fmt.Printf(input + "\n")
		}

	}

}
