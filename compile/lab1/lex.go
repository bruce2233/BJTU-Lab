package lab1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func scan() {
	// context := []byte("abc 12 d34 {")
	// token := []string{"IDENTIFIER"}
	// var context []byte
	// exp1 := regexp.MustCompile(`^[a-zA-Z]+\w`)
	// exp2 := regexp.MustCompile(`^[0-9]+`)
	// exp3 := regexp.MustCompile("^[\\+\\-\\*\\;\\(\\)\\,\\{\\}]+")
	// exp4 := regexp.MustCompile(`^\s+`)
	// exps := []regexp.Regexp{}
	// exps = append(exps, *exp1)
	// exps = append(exps, *exp2)
	// exps = append(exps, *exp3)
	// exps = append(exps, *exp4)
	ruleByte, err := ioutil.ReadFile("rule.json")
	if err != nil {
		fmt.Println("open rule.json err")
	}
	var rule Rule
	err = json.Unmarshal(ruleByte, &rule)
	if err != nil {
		fmt.Println("unmarshal json err")
	}
	var tokensCompiled []TokenCompiled
	for _, item := range rule.Tokens {
		fmt.Println(item.ExpStr)
		exp := regexp.MustCompile(item.ExpStr)
		// tokensCompiled[index].ExpCompiled = *exp
		tokenCompiled := TokenCompiled{item, *exp}
		tokensCompiled = append(tokensCompiled, tokenCompiled)
	}
	lexScanner := &LexScanner{}
	lexScanner.TokensCompiled = tokensCompiled
	lexScanner.pos = 0

	context, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("read file error!")
	}
	outPutFile, err := os.OpenFile("output.txt", os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("open output err")
	}

	for {
		if lexScanner.pos == len(context) {
			break
		}
		res := lexScanner.findFromExps(context)
		outPutFile.Write(res)
	}
}

type LexScanner struct {
	pos            int
	TokensCompiled []TokenCompiled
}

func (lexScanner *LexScanner) findFromExps(context []byte) []byte {
	for _, item := range lexScanner.TokensCompiled {
		res := item.ExpCompiled.Find(context[lexScanner.pos:])
		if res != nil {
			lexScanner.pos += len(res)
			fmt.Println(string(res), ": ", item.TokenName)
			return res
		}
	}
	panic("no")
	return nil
}

type TokenMatchText struct {
	token []byte
	text  []byte
}

type Rule struct {
	Tokens []Token `json:"tokens"`
}
type Token struct {
	ExpStr    string `json:"expStr"`
	TokenName string `json:"token"`
}

type TokenCompiled struct {
	Token
	ExpCompiled regexp.Regexp
}
