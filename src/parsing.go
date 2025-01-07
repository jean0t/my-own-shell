package main

import (
	"strings"
	"os"
	"os/exec"
)

type Arg struct {
	Command string
	Args []string
}

type ShellFunction func([]byte) string

var arg Arg

func emptyArgument(argument []string) bool {
	for _, v := range argument {
		if string(v) == "" {
			return true
		}
	}

	return false
}

func Parse(input string) string {

	var parsedInput = strings.Fields(input)

	if len(parsedInput) < 2 && len(parsedInput) > 0 {

		arg.Command = parsedInput[0]
		arg.Args = []string{""}

	} else if len(parsedInput) >= 2 {

		arg.Command = parsedInput[0]
		arg.Args = parsedInput[1:]

	} else {

		arg.Command = ""
		arg.Args = []string{""}

	}

	switch(arg.Command) {
		case "cd":
			return Command.Cd(arg.Args[0])

		case "pwd":
			return Command.Pwd()

		case "env":
			return Command.Env()

		case "ls":
			if emptyArgument(arg.Args) || arg.Args[0] != "-l" {
				return Command.Ls(false)
			} else {
				return Command.Ls(true)
			}

		case "rename":
			if len(arg.Args) == 2 {
				return Command.Rename(arg.Args[0], arg.Args[1])
			} else {
				return "Not enough arguments\n"
			}

		case "":
			return ""

		case "exit":
			os.Exit(0)
			return ""

		case "mkdir":
			if len(arg.Args) == 1 {
				return Command.Mkdir(arg.Args[0])
			}
			return "Pass a name for the directory\n"

		case "ping":
			if len(arg.Args) > 0 && !emptyArgument(arg.Args) {
				return Command.Ping(arg.Args[0])
			} else {
				return "An error has occurred\n"
			}

		default:
			var cmd = exec.Command(arg.Command, arg.Args...)
			err := cmd.Run()
			if err != nil {
				return "An error occurred\n"
			}

			return ""
	}

}
