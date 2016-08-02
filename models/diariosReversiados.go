package models
var (
	DiariosReversiados map[string]*DiarioRe
)
func init() {
	DiariosReversiados = make(map[string]*DiarioRe)
}
type DiarioRe struct {
	Diario int
	AllocRef int
	PERIOD int64
}
func AddDiarioReversiado(u DiarioRe, index string)  {
	DiariosReversiados[index] = &u	
}
func GetAllDiarioReversiados() map[string]*DiarioRe {
	return DiariosReversiados
}
func ClearDiariosReversiados() {
	DiariosReversiados = make(map[string]*DiarioRe)
}