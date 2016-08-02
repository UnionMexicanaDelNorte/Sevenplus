package models
var (
	DimensionesList map[string]*Dimensiones
)
func init() {
	DimensionesList = make(map[string]*Dimensiones)
}
type Dimensiones struct {
	ANL_CAT_ID string
	S_HEAD string
	DESCR string
}

func AddDimension(u Dimensiones, index string)  {
	DimensionesList[index] = &u	
}
func GetAllDimensiones() map[string]*Dimensiones {
	return DimensionesList
}
