package lab1

import (
	"fmt"
	"io/ioutil"
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

func TestReservedWords(t *testing.T) {
	context, err := ioutil.ReadFile("rule.txt")
	if err != nil {
		t.Log("error!")
	}


}
