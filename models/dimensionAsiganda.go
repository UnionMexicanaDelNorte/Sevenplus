package models
var (
	DimensionAList map[string]*DimensionAsignada
)
func init() {
	DimensionAList = make(map[string]*DimensionAsignada)
}
type DimensionAsignada struct {
	ANL_CAT_ID string
	ENTRY_NUM int64
}
func AddDimensionA(u DimensionAsignada, index string)  {
	DimensionAList[index] = &u	
}
func GetAllDimensionA() map[string]*DimensionAsignada {
	return DimensionAList
}
func ClearDimensionA() {
	//checar documentacion si es asi o no
	//AGREGAR A TODOS!
	DimensionAList = make(map[string]*DimensionAsignada)	
}