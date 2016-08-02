package models
var (
	TiposDeDiarioList map[string]*TiposDD
)
func init() {
	TiposDeDiarioList = make(map[string]*TiposDD)
}
type TiposDD struct {
	IdTipoDeDiario int
	Codigo string
	Nombre string
}
func AddTiposDeDiariop(u TiposDD, index string)  {
	TiposDeDiarioList[index] = &u	
}
func GetAllTiposDeDiario() map[string]*TiposDD {
	return TiposDeDiarioList
}
func ClearTiposDeDiario() {
	TiposDeDiarioList = make(map[string]*TiposDD)
}