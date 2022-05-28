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
	err = ioutil.WriteFile("output.txt", []byte{}, 0777)
	outPutFile, err := os.OpenFile("output.txt", os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("open output err")
	}

	for {
		if lexScanner.pos == len(context) {
			break
		}
		res := lexScanner.findFromExps(context)
		outPutFile.Write([]byte(string(res.text) + ": " + res.TokenName + ";\n"))
	}
}

type LexScanner struct {
	pos            int
	TokensCompiled []TokenCompiled
}

func (lexScanner *LexScanner) findFromExps(context []byte) TokenMatchText {
	for _, item := range lexScanner.TokensCompiled {
		res := item.ExpCompiled.FindSubmatch(context[lexScanner.pos:])
		if res == nil {
			continue
		}
		var matchedRes []byte
		if len(res) == 1 {
			matchedRes = res[0]
		} else if len(res) == 2 {
			matchedRes = res[1]
		}
		lexScanner.pos += len(matchedRes)
		fmt.Println(string(matchedRes), ": ", item.TokenName, ";")
		tokenMatchText := TokenMatchText{
			TokenName: item.TokenName,
			text:      matchedRes}
		return tokenMatchText
	}
	panic("invalid lex position: " + string(context[lexScanner.pos:lexScanner.pos+10]))
}

func (lexScanner *LexScanner) replaceReservedWord(reservedWords *[]string) {
	_, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("read file error!")
	}
	fmt.Println(reservedWords)
}

type TokenMatchText struct {
	TokenName string
	text      []byte
}

type Rule struct {
	Tokens        []Token  `json:"tokens"`
	ReservedWords []string `json:"reservedWords"`
}
type Token struct {
	ExpStr    string `json:"expStr"`
	TokenName string `json:"token"`
}

type TokenCompiled struct {
	Token
	ExpCompiled regexp.Regexp
}
