package myshell

import (
	"mysyscall/gocall"
	"path/filepath"
	"syscall"
)

func ls(myPath string) []string {
	const (
		bufSize = 2048
	)

	var fd int
	var err error
	//判断路径
	if filepath.IsAbs(myPath) {
		fd, err = syscall.Open(myPath, syscall.O_RDONLY, 0)
	} else {
		fd, err = syscall.Open(gocall.Getcwd()+"/"+myPath, syscall.O_RDONLY, 0)
	}
	defer syscall.Close(fd)
	if err != nil {
		println("path doesn't exist")
	}

	direntBytes := make([]byte, bufSize)
	_, _ = syscall.ReadDirent(fd, direntBytes)
	fileNames := []string{}
	_, _, fileNames = syscall.ParseDirent(direntBytes, 100, fileNames)
	return fileNames
}

func run(file string) (int, error) {
	var pid int
	var err error
	procAttr := syscall.ProcAttr{
		Files: []uintptr{0, 1, 2},
	}
	path := gocall.Getenv("PATH")
	// println("run file :", file)
	// println("run path :", path)
	pid, err = syscall.ForkExec(path+"/"+file, nil, &procAttr)
	if err != nil {
		return pid, err
	} else {
		path = gocall.Getenv("HOME")
		pid, err = syscall.ForkExec(path+"/"+file, nil, &procAttr)
		if err != nil {
			return pid, err
		}
	}
	return pid, err
}
func creat(file string) error {

	fd, err := syscall.Creat(file, 0)
	// syscall.Read(fd,buf)
	// println(n)
	if err != nil {
		println("create file fail")
	}
	println(fd)
	return err
}
