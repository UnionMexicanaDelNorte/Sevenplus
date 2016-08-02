package models
var (
	PeriodosList map[string]*Periodo
)
func init() {
	PeriodosList = make(map[string]*Periodo)
}
type Periodo struct {
	PERIOD string
}
func AddPeriodo(u Periodo, index string)  {
	PeriodosList[index] = &u	
}
func GetAllPeriodos() map[string]*Periodo {
	return PeriodosList
}