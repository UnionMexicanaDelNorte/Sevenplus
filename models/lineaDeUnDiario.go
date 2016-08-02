package models
var (
	LineaDeUnDiarioList map[string]*LineaDeUnDiario
)
func init() {
	LineaDeUnDiarioList = make(map[string]*LineaDeUnDiario)
}

type LineaDeUnDiario struct {
	JRNAL_NO int
	JRNAL_LINE int
	AMOUNT float64
	D_C string
	ACNT_CODE string
	DESCRIPTN string
	PERIOD int
	JRNAL_SRCE string
	JRNAL_TYPE string
	TRANS_DATETIME string
}

func AddLineaDeUnDiario(u LineaDeUnDiario, index string)  {
	LineaDeUnDiarioList[index] = &u	
}
func GetAllLineaDeUnDiario() map[string]*LineaDeUnDiario {
	return LineaDeUnDiarioList
}
func ClearLineaDeUnDiario() {
	LineaDeUnDiarioList = make(map[string]*LineaDeUnDiario)	
}
