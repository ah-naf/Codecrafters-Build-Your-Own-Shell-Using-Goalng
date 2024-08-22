package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

var SHELL_BUILTIN = []string{"echo", "type", "exit", "pwd"}

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
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				os.Exit(1)
			}
			fmt.Println(dir)
		case "type":
			command := strings.Join(commands[1:], " ")
			if slices.Contains(SHELL_BUILTIN, command) {
				fmt.Printf("%s is a shell builtin\n", command)
			} else {
				found := false
				paths := strings.Split(os.Getenv("PATH"), ":")
				for _, path := range paths {
					fp := filepath.Join(path, command)
					if _, err := os.Stat(fp); err == nil {
						fmt.Printf("%s is %s\n", command, fp)
						found = true
						break
					}
				}
				if !found {
					fmt.Printf("%s: not found\n", command)
				}
			}
		default:
			command := exec.Command(commands[0], strings.Join(commands[1:], " "))
			command.Stdin = os.Stdin
			command.Stdout = os.Stdout
			command.Stderr = os.Stderr

			if err := command.Run(); err != nil {
				fmt.Printf("%s: command not found\n", commands[0])
			}
	
		}
	}
}
