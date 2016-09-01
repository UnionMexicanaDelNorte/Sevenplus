package models
var (
	OtrosCamposList map[string]*OtrosCampos
)
func init() {
	OtrosCamposList = make(map[string]*OtrosCampos)
}
type OtrosCampos struct {
	IdCampo int
	Nombre string
	URL string
}

func AddOtrosCampos(u OtrosCampos, index string)  {
	OtrosCamposList[index] = &u	
}
func GetAllOtrosCampos() map[string]*OtrosCampos {
	return OtrosCamposList
}
func ClearOtrosCampos() {
	OtrosCamposList = make(map[string]*OtrosCampos)	
}
