package tools

import "os"

func Pwd() string {
	var Path, _ = os.Getwd()
	return Path + "\n"
}
