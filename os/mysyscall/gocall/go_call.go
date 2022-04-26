package gocall

import (
	"path/filepath"
	"syscall"
)

func Getcwd() string {
	cwd := make([]byte, 1024)
	n, _ := syscall.Getcwd(cwd)
	cwdstr := string(cwd[0 : n-1])
	return cwdstr
}

func GetDirFileNames(myPath string) []string {
	const (
		bufSize = 2048
	)

	var fd int
	var err error
	if filepath.IsAbs(myPath) {
		// fmt.Println(true)
		fd, err = syscall.Open(myPath, syscall.O_RDONLY, 0)
	} else {
		fd, err = syscall.Open(Getcwd()+"/"+myPath, syscall.O_RDONLY, 0)
	}
	// println(fd)
	defer syscall.Close(fd)
	if err != nil {
		println("path doesn't exist")
	}

	direntBytes := make([]byte, bufSize)
	// fmt.Println(direntBytes)
	_, _ = syscall.ReadDirent(fd, direntBytes)
	fileNames := []string{}
	_, _, fileNames = syscall.ParseDirent(direntBytes, 100, fileNames)
	return fileNames
}
func ForkExec(file string, execPath string) {
	if filepath.IsAbs(file) {
		// syscall.ForkExec(file)
	}
}
func Getenv(key string) string {
	value, ok := syscall.Getenv(key)
	if !ok {
		println("env key not found!")
	}
	return value
}
