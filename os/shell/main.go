package main

import (
	"os/user"
	"shellmod/myshell"
	"syscall"
)

func main() {
	syscall.Setenv("PATH", "/etc/mypath")
	u, _ := user.Current()
	if u.Username == "root" {
		syscall.Setenv("HOME", "/root")
	} else {
		syscall.Setenv("HOME", "/home/"+u.Username)
	}
	println("A shell by 19281030-张云鹏")
	app := myshell.InitCmd()
	err := app.Run()
	print(err)
}
