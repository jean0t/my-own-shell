package main

import (
	"os"
	"os/exec"
)
//==================================================================

// Important structure to organization
type Arg struct {
	Command string // command is considered the first word being called
	Args []string // argument is considered everything after command
}

var arg Arg // it will hold the user input and separate in command/Args


func SanitizeInput(input string) []string {

  var result []string = []string{}
  var tmp string = ""
  var quoted bool = false
  for i, v := range input {
    switch {
      case v == ' ' && quoted == false :
        result = append(result, tmp)
        tmp = ""

      case i == len(input) - 1 && v != ' ':
        tmp += string(v)
        result = append(result, tmp)
        tmp = ""

      case v == '\'' || v == '"':
        if !quoted {
          quoted = true
          tmp += string(v)
          continue
        } else {
          quoted = false
          tmp += string(v)
          continue
        }

      default:
        tmp += string(v)
    }
  }
  return result
}


func Parse(input string) string {

	var parsedInput []string = SanitizeInput(input)

	if len(parsedInput) < 2 && len(parsedInput) > 0 {

		arg.Command = parsedInput[0]
		arg.Args = []string{}

	} else if len(parsedInput) >= 2 {

		arg.Command = parsedInput[0]
		arg.Args = parsedInput[1:]

	} else {

		arg.Command = ""
		arg.Args = []string{}

	}

	switch(arg.Command) {
		case "cd":
			return Command.Cd(arg.Args[0])

		case "pwd":
			return Command.Pwd()

		case "env":
			return Command.Env()

		case "ls":
			if len(arg.Args) > 0 {
        if arg.Args[0] == "-l" {
				  return Command.Ls(true)
        }
        return ""
			} else {
				return Command.Ls(false)
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
			if len(arg.Args) > 0 && !(arg.Args[0] == "") {
				return Command.Ping(arg.Args[0])
			} else {
				return "An error has occurred\n"
			}

		default:
      var cmd *exec.Cmd
      if len(arg.Args) > 0 {
			  cmd = exec.Command(arg.Command, arg.Args...)
      } else {
        cmd = exec.Command(arg.Command)
      }
      cmd.Stdout = os.Stdout
      cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				return ""
			}

			return ""
	}

}
