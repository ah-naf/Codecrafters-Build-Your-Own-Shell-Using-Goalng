package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var SHELL_BUILTIN = []string{"echo", "type", "exit"}

func main() {
	for {
		// Uncomment this block to pass the first stage
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		message, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		message = strings.TrimSpace(message)
		commands := strings.Split(message, " ")
		

		switch commands[0] {
		case "exit":
			code, err := strconv.Atoi(commands[1])
			if err != nil {
				os.Exit(1)
			}
			os.Exit(code)
		case "echo":
			fmt.Println(strings.Join(commands[1:], " "))
		case "type":
			command := strings.Join(commands[1:], " ")
			if slices.Contains(SHELL_BUILTIN, command) {
				fmt.Printf("%s is a shell builtin\n", command)
			} else {
				fmt.Printf("%s: not found\n", command)
			}
		default:
			fmt.Printf("%s: command not found\n", commands[0])
		}
	}
}
