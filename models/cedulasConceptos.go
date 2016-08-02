package models
var (
	CedulasConceptosList map[string]*CedulasConceptos
)
func init() {
	CedulasConceptosList = make(map[string]*CedulasConceptos)
}
type CedulasConceptos struct {
	IdConcepto int64
	IdCedula int64
	Nombre string
}
func AddCedulasConceptos(u CedulasConceptos, index string)  {
	CedulasConceptosList[index] = &u	
}
func GetAllCedulasConceptos() map[string]*CedulasConceptos {
	return CedulasConceptosList
}
