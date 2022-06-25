package lab2

import "fmt"

type TokenEntity struct {
	text      string
	tokenName string
}

//产生式
type Production struct {
	from string
	to   string
}

type LL1Analyzer struct {
	analysisStack string         //分析栈
	leftString    string         //输入串
	analysisTable [][]Production //分析表
}

func (analyzer *LL1Analyzer) analyze() {
	var tokenEntityList []TokenEntity
	fmt.Println(tokenEntityList)
	analyzer.buildTable()
}

func (analyzer *LL1Analyzer) buildTable() {
	analyzer.analysisTable['E']['i'] = Production{from: "E", to: "Te"}
	analyzer.analysisTable['E']['('] = Production{from: "E", to: "Te"}
	analyzer.analysisTable['e']['+'] = Production{from: "e", to: "ATe"}
	analyzer.analysisTable['e']['-'] = Production{from: "e", to: "ATe"}
	analyzer.analysisTable['e']['('] = Production{from: "e", to: ""}
	analyzer.analysisTable['e'][')'] = Production{from: "e", to: ""}
	analyzer.analysisTable['T']['i'] = Production{from: "T", to: "Ft"}
	analyzer.analysisTable['T']['('] = Production{from: "T", to: "Ft"}
	analyzer.analysisTable['t']['+'] = Production{from: "t", to: ""}
	analyzer.analysisTable['t']['-'] = Production{from: "t", to: ""}
	analyzer.analysisTable['t']['*'] = Production{from: "t", to: "MFt"}
	analyzer.analysisTable['t']['/'] = Production{from: "t", to: "MFt"}
	analyzer.analysisTable['t'][')'] = Production{from: "t", to: ""}
	analyzer.analysisTable['t']['#'] = Production{from: "t", to: ""}
}


