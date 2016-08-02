package models
var (
	CedulasLineasList map[string]*CedulasLineas
)
func init() {
	CedulasLineasList = make(map[string]*CedulasLineas)
}
type CedulasLineas struct {
	IdLinea int
	IdConcepto int64
	Cuenta string
	TipoDC int64
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

func AddCedulasLineas(u CedulasLineas, index string)  {
	CedulasLineasList[index] = &u	
}
func GetAllCedulasLineas() map[string]*CedulasLineas {
	return CedulasLineasList
}
func ClearCedulasLineas() {
	CedulasLineasList = make(map[string]*CedulasLineas)	
}
