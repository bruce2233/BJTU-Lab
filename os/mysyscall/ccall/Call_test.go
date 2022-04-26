package ccall

import (
	"mysyscall/gocall"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)
import "testing"
import "syscall"
import "fmt"
import "reflect"
import "bytes"

// import "unsafe"

func TestCallPid(t *testing.T) {
	pid, _, _ := syscall.Syscall(39, 0, 0, 0) // 用不到的就补上 0
	fmt.Println(reflect.TypeOf(pid))
	fmt.Println("Process id: ", pid)
	a := 10
	b := &a
	print(b)
}

func TestCallGetCWD(t *testing.T) {
	cwd := make([]byte, 0, 100)
	syscall.Getcwd(cwd)
	fmt.Println("cwdstr: ", string(cwd))

}
func TestCallChdir(t *testing.T) {
	cdpath := "../"
	syscall.Chdir(cdpath)
}

func TestOpen(t *testing.T) {
	const (
		bufSize = 2048
	)
	cwd := gocall.Getcwd()
	dstPath := filepath.Clean(cwd + "/../")
	t.Log(dstPath)
	t.Log(cwd == "/home/zhangye/github/os/mysyscall/ccall")
	t.Log(dstPath == "/home/zhangye/github/os/mysyscall")
	fd1, _ := syscall.Open(dstPath, syscall.O_RDONLY, 0)
	t.Logf("%d", fd1)
	direntBytes := make([]byte, bufSize)
	// direntBytes := bytes.Repeat([]byte{0xCD}, bufSize)
	n, _ := syscall.ReadDirent(fd1, direntBytes)
	t.Logf("%d", n)

	fileNames := []string{}
	consumed, count, fileNames := syscall.ParseDirent(direntBytes, 100, fileNames)
	t.Logf("%d", consumed)
	t.Logf("%d", count)
	t.Logf("%q", fileNames)
}

func TestDirent(t *testing.T) {
	const (
		direntBufSize   = 2048 // arbitrary? See https://go.dev/issue/37323.
		filenameMinSize = 11
	)

	d := t.TempDir()
	t.Logf("tmpdir: %s", d)

	for i, c := range []byte("0123456789") {
		name := string(bytes.Repeat([]byte{c}, filenameMinSize+i))
		err := os.WriteFile(filepath.Join(d, name), nil, 0644)
		if err != nil {
			t.Fatalf("writefile: %v", err)
		}
	}

	names := make([]string, 0, 10)

	fd, err := syscall.Open(d, syscall.O_RDONLY, 0)
	if err != nil {
		t.Fatalf("syscall.open: %v", err)
	}
	defer syscall.Close(fd)

	buf := bytes.Repeat([]byte{0xCD}, direntBufSize)
	for {
		n, err := syscall.ReadDirent(fd, buf)
		if err == syscall.EINVAL {
			// On linux, 'man getdents64' says that EINVAL indicates “result buffer is too small”.
			// Try a bigger buffer.
			t.Logf("ReadDirent: %v; retrying with larger buffer", err)
			buf = bytes.Repeat([]byte{0xCD}, len(buf)*2)
			continue
		}
		if err != nil {
			t.Fatalf("syscall.readdir: %v", err)
		}
		t.Logf("ReadDirent: read %d bytes", n)
		if n == 0 {
			break
		}

		var consumed, count int
		consumed, count, names = syscall.ParseDirent(buf[:n], -1, names)
		t.Logf("ParseDirent: %d new name(s)", count)
		if consumed != n {
			t.Fatalf("ParseDirent: consumed %d bytes; expected %d", consumed, n)
		}
	}

	sort.Strings(names)
	t.Logf("names: %q", names)

	if len(names) != 10 {
		t.Errorf("got %d names; expected 10", len(names))
	}
	for i, name := range names {
		ord, err := strconv.Atoi(name[:1])
		if err != nil {
			t.Fatalf("names[%d] is non-integer %q: %v", i, names[i], err)
		}
		if expected := string(strings.Repeat(name[:1], filenameMinSize+ord)); name != expected {
			t.Errorf("names[%d] is %q (len %d); expected %q (len %d)", i, name, len(name), expected, len(expected))
		}
	}
}

func TestSlice(t *testing.T) {
	a := bytes.Repeat([]byte{255}, 100)
	t.Log(len(a))
	t.Log(cap(a))

	b := []byte{1, 2, 3, 4, 5}
	copy(b, a)
	t.Log(len(b))
	t.Log(b)
}

func TestLs(t *testing.T) {
	names := gocall.Ls("../")
	// fmt.Printf("%q",names)
	fmt.Println(names)
}
