package models
var (
	IglesiasList map[string]*Iglesias
)
func init() {
	IglesiasList = make(map[string]*Iglesias)
}
type PorCuenta struct {
	ACNT_CODE string
	Saldo float64
	DESCR string
}
type PorPeriodo struct {
	PERIOD int
	Distrito string
	CuentasList map[string]*PorCuenta
}
type Iglesias struct {
	Nombre string
	Codigo string
	PeriodosList map[string]*PorPeriodo
}
func AddIglesias(u Iglesias, index string)  {
	IglesiasList[index] = &u	
}
func DameIglesiaPorCodigo(index string) Iglesias {
	return *IglesiasList[index]
}
func GetAllIglesias() map[string]*Iglesias {
	return IglesiasList
}
func ClearIglesias() {
	IglesiasList = make(map[string]*Iglesias)
}