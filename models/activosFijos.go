package models
var (
	ActivosFijosList map[string]*ActivoFijo
)
func init() {
	ActivosFijosList = make(map[string]*ActivoFijo)
}
type ActivoFijo struct {
	ASSET_CODE string
	STATUS int
	ASSET_STATUS int
	DESCR string
	START_PERIOD int
	END_PERIOD int
	ULTIMO_PERIOD int
	DISPOSAL_PERIOD int
	DISPOSED int
	BASE_GROSS float64
	BASE_DEP float64
	BASE_NET float64
	BASE_PCENT float64
	TXN_GROSS float64
	TXN_DEP float64
	TXN_NET float64
	TXN_PCENT float64
}
func AddActivoFijo(u ActivoFijo, index string)  {
	ActivosFijosList[index] = &u	
}
func GetAllActivosFijos() map[string]*ActivoFijo {
	return ActivosFijosList
}
func ClearActivosFijos() {
	ActivosFijosList = make(map[string]*ActivoFijo)
}