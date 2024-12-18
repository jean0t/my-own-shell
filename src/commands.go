package main

import (
	"my-own-shell/tools"
)

type Commands struct {
	Cd func(string)string
	Pwd func()string
	Env func()string
	Ls func(bool)string
}

var Command = Commands{
	Cd: tools.Cd,
	Pwd: tools.Pwd,
	Env: tools.Env,
	Ls: tools.Ls,
}
