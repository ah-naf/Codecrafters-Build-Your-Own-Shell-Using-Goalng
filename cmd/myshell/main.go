package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		// Uncomment this block to pass the first stage
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		cmd = strings.TrimSpace(cmd)

		switch cmd {
		case "exit 0":
			os.Exit(0)
		default:
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}
