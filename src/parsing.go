package main

import (
	"strings"
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
			if emptyArgument(arg.Args) {
				return Command.Ls(false)
			} else {
				return Command.Ls(true)
			}

		case "":
			return ""

		default:
			var cmd = exec.Command(arg.Command, arg.Args...)
			err := cmd.Run()
			if err != nil {
				return "An error occurred\n"
			}

			return ""
	}

}
