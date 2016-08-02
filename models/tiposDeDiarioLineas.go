package models
var (
	TiposDeDiariosLineasList map[string]*TipoDeDiariosLineas
)
func init() {
	TiposDeDiariosLineasList = make(map[string]*TipoDeDiariosLineas)
}
type TipoDeDiariosLineas struct {
	IdLinea int
	Linea int
	DeboFacturar int
	Cuenta string
	Cliente string
	Servicio string
	Concepto string
	DESCRIPTN string
	D_C string
	ANAL_T0 string
	ANAL_T1 string
	ANAL_T2 string
	ANAL_T3 string
	ANAL_T4 string
	ANAL_T5 string
	ANAL_T6 string
	ANAL_T7 string
	ANAL_T8 string
	ANAL_T9 string
}

func AddTiposDeDiarioLineas(u TipoDeDiariosLineas, index string)  {
	TiposDeDiariosLineasList[index] = &u	
}
func GetAllTiposDeDiariosLineas() map[string]*TipoDeDiariosLineas {
	return TiposDeDiariosLineasList
}
func ClearTiposDeDiarioLineas() {
	TiposDeDiariosLineasList = make(map[string]*TipoDeDiariosLineas)	
}
