package adapterbridge

import (
	"testing"
)

func TestAB(t *testing.T) {
	var dataPaser IDataBaseParser
	var language ILanguage
	var languageSetterImpl LanguageSetterImpl
	dataPaser = &MySql{&languageSetterImpl}
	language = Chinese{}
	dataPaser.SetLanguage(language)
	dataPaser.PrintGrade()

	language = English{}
	dataPaser.SetLanguage(language)
	dataPaser.PrintGrade()

	language = WaterMark{}
	dataPaser.SetLanguage(language)
	dataPaser.PrintGrade()

	dataPaser = &Oracle{&languageSetterImpl}
	language = Chinese{}
	dataPaser.SetLanguage(language)
	dataPaser.PrintGrade()

	language = English{}
	dataPaser.SetLanguage(language)
	dataPaser.PrintGrade()

	language = WaterMark{}
	dataPaser.SetLanguage(language)
	dataPaser.PrintGrade()

	dataPaser = &SqlServer{&languageSetterImpl}
	language = Chinese{}
	dataPaser.SetLanguage(language)
	dataPaser.PrintGrade()

	language = English{}
	dataPaser.SetLanguage(language)
	dataPaser.PrintGrade()

	language = WaterMark{}
	dataPaser.SetLanguage(language)
	dataPaser.PrintGrade()
}

type D struct {
	E
}
type E struct {
	a int
}

func (e *E) setE() {
	e.a = 10
}
func TestAnother(t *testing.T) {
	d := D{}
	d.setE()
	t.Log(d)
}
