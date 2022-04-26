package adapterbridge

type OriSqlServer struct{}

func (oriSqlServer OriSqlServer) OriGetGrade() {
	println("data from SqlServer")
}

type ISqlServer interface{
	GetGrade()
}
type SqlServerAdapter struct{
	OriSqlServer
}
func (sqlServerAdapter SqlServerAdapter)GetGrade(){
	sqlServerAdapter.OriSqlServer.OriGetGrade()
}
func NewSqlServer() ISqlServer{
	return SqlServerAdapter{}
}