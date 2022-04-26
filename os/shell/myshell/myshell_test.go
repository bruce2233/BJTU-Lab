package myshell

import (
	"os"
	"syscall"
	"testing"
)

func TestGetenv(t *testing.T) {
	syscall.Setenv("PATH", "/etc/mypath")
	value, _ := syscall.Getenv("PATH")
	t.Log(value)
}
func TestFork(t *testing.T) {
	// r1, _, _ :=syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)
	// t.Log(r1)
	// t.Log("a process")
	sttyArgs := syscall.ProcAttr{
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
	}
	// sttyAr
	// t.Log(sttyArgs)
	// t.Log(os.Environ())
	// procattr := syscall.ProcAttr{}
	syscall.Setenv("PATH", "/etc/mypath")
	path, _ := syscall.Getenv("PATH")
	pid, err := syscall.ForkExec(path+"/"+"myapp", []string{}, &sttyArgs)
	if err != nil {
		t.Log("err")
	}
	t.Log(pid)
}
func TestMkdir(t *testing.T) {
	err := syscall.Mkdir("/home/zhangye/github/os/shell/myshell/testdir", 0)
	if err != nil {
		t.Log("mkdir err!")
	}
}

func TestOpen(t *testing.T) {
	fd, _ := syscall.Open("no", syscall.O_WRONLY, 0)
	t.Log(fd)

}
