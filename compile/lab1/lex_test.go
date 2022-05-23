package lab1

import (
	"fmt"
	"regexp"
	"testing"
)

func TestScan(t *testing.T) {
	t.Log("testing...")
	scan()
}

func TestReg(t *testing.T) {
	s := []byte("/* \nad*/")
	re := regexp.MustCompile("\\A/\\*[\\s\\S]*\\*/")
	res := re.FindSubmatch(s)
	fmt.Println(res)
	t.Log(res)
	for _, item := range res {
		fmt.Println(string(item))
	}
}
