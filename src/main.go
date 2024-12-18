package main

import (
	"fmt"
	"my-own-shell/tools"
	"github.com/ergochat/readline"
)

var (
	shell_prompt = "$ "
	history tools.History = tools.History{History: []string{},} // history is better if made here directly
)

func main() {
	var rl, err = readline.New(shell_prompt)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rl.Close()

	for {
		line, err := rl.ReadLine()
		if err != nil {
			fmt.Println(err)
			return
		}
		history.Append(line)

		if line == "history" {
			history.Show()
			continue
		}

		fmt.Fprintf(rl, "%s", Parse(line))
	}
}
