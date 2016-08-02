package models
var (
	UnidadDeNegocioList map[string]*Unidad
)
func init() {
	UnidadDeNegocioList = make(map[string]*Unidad)
}
type Unidad struct {
	BUNIT string
}
func AddUnidadDeNegocio(u Unidad, index string)  {
	UnidadDeNegocioList[index] = &u	
}
func GetAllUnidadDeNegocio() map[string]*Unidad {
	return UnidadDeNegocioList
}