package lab1

import (
	"fmt"
	"regexp"
	"testing"
)

func TestScan(t *testing.T) {
	t.Log("testing...")
	// exp1 := regexp.MustCompile("^[a-zA-Z]+\\w")
	// fmt.Println(exp1)
	scan()
}

func TestReg(t *testing.T) {
	str := "<"
	matched, err := regexp.MatchString("<>(?:>)", str)
	fmt.Println(matched, err)
}
