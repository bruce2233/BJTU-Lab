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
	ruleByte, err := ioutil.ReadFile("rule.json") //打开json文件
	if err != nil {
		fmt.Println("open rule.json err")
	}
	var rule Rule
	err = json.Unmarshal(ruleByte, &rule) //解析json格式
	if err != nil {
		fmt.Println("unmarshal json err")
	}
	var tokensCompiled []TokenCompiled
	for _, item := range rule.Tokens { //将正则表达式和对应的token名称保存
		fmt.Println(item.ExpStr)
		exp := regexp.MustCompile(item.ExpStr)
		tokenCompiled := TokenCompiled{item, *exp}
		tokensCompiled = append(tokensCompiled, tokenCompiled)
	}
	lexScanner := &LexScanner{}
	lexScanner.TokensCompiled = tokensCompiled
	lexScanner.pos = 0

	context, err := ioutil.ReadFile("input.txt") //读取输入串
	if err != nil {
		fmt.Println("read file error!")
	}
	err = ioutil.WriteFile("output.txt", []byte{}, 0777) //打印输出二元式序列组
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
	TokenName string //单词名称
	text      []byte //匹配到的输入串
}

type Rule struct {
	Tokens        []Token  `json:"tokens"`        //单词token结构体数组
	ReservedWords []string `json:"reservedWords"` //保留字
}
type Token struct {
	ExpStr    string `json:"expStr"` //单词的正则表达式
	TokenName string `json:"token"`  //单词的名称
}

type TokenCompiled struct {
	Token                     //单词token
	ExpCompiled regexp.Regexp //编译过的正则表达式
}
