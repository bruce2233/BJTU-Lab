package lab1

import "io/ioutil"

var (
	ruleFile  = "rule.txt"
	inputFile = "code.txt"
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

func (fsm *FSM) loadNodes() {
	bytes, err := ioutil.ReadFile(inputFile)
	if err !=nil{
		panic("nodes file err")
	}
	inputString := string(bytes)
	println(inputString)
}