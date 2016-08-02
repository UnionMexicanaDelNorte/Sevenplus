package models
var (
	CedulasConceptosLineasList map[string]*CedulasConceptoLineas
)
func init() {
	CedulasConceptosLineasList = make(map[string]*CedulasConceptoLineas)
}
type CedulasConceptoLineas struct {
	IdConceptoOLinea int64
	NombreOCuenta string
	Cantidad float64
}
func AddCedulasConceptosLineas(u CedulasConceptoLineas, index string)  {
	CedulasConceptosLineasList[index] = &u	
}
func GetAllCedulasConceptosLineas() map[string]*CedulasConceptoLineas {
	return CedulasConceptosLineasList
}
func ClearCedulasConceptosLineas() {
	CedulasConceptosLineasList = make(map[string]*CedulasConceptoLineas)	
}