package lab1

import (
	"regexp"
)

var (
	ruleFile  = "rule.txt"
	inputFile = "code.txt"
	regs      = []string{"a+"}
)

type FSM struct {
	inputString string
	nodeList    []Node
	curNode     Node
	curPosition int
}

type Node struct {
	state int
}

func (node Node) transfer(inputString string) {
	input := "test input"
	println("get", input, "transfer from a to b")
}

func (fsm *FSM) loadRegs() {
	matched, err := regexp.Match(regs[0], []byte("aaa"))
	if err != nil {
		println("compile reg err")
	}
	println(matched)
}
