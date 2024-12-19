package main

import (
	"my-own-shell/tools"
)

type Commands struct {
	Cd func(string)string
	Pwd func()string
	Env func()string
	Ls func(bool)string
	Rename func(string, string)string
	Mkdir func(string)string
}

var Command = Commands{
	Cd: tools.Cd,
	Pwd: tools.Pwd,
	Env: tools.Env,
	Ls: tools.Ls,
	Rename: tools.Rename,
	Mkdir: tools.Mkdir,
}
