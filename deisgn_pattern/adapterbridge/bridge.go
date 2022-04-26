package adapterbridge

type IDataBaseParser interface {
	SetLanguage(ILanguage)
	PrintGrade()
}

type ILanguage interface {
	Say()
}

type Chinese struct{}

func (chinese Chinese) Say() {
	println("in Chinese")
}

type English struct{}

func (english English) Say() {
	println("in English")

}

type WaterMark struct{}

func (waterMark WaterMark) Say() {
	println("in WaterMark")
}

type ILanguageSetter interface {
	SetLanguage(ILanguage)
}
type LanguageSetterImpl struct {
	ILanguage
}

func (languageSetterImpl *LanguageSetterImpl) SetLanguage(il ILanguage) {
	languageSetterImpl.ILanguage = il
}

type MySql struct {
	*LanguageSetterImpl
}

func (mysql MySql) PrintGrade() {
	println("data from Mysql")
	mysql.ILanguage.Say()
}

type Oracle struct {
	*LanguageSetterImpl
}

func (oracle Oracle) PrintGrade() {
	println("data from Oracle")
	oracle.ILanguage.Say()
}

type SqlServer struct {
	*LanguageSetterImpl
}

func (sqlServer SqlServer) PrintGrade() {
	NewSqlServer().GetGrade()
	sqlServer.ILanguage.Say()
}
