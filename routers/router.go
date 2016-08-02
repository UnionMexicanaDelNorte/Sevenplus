package routers

import (
	"sevenplus/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/app", &controllers.LoginController{})
    beego.Router("/estadosDeCuenta", &controllers.EstadosDeCuentaController{})
    beego.Router("/generarEstadoDeCuenta", &controllers.GenerarEstadosDeCuentaController{})
    beego.Router("/generarDiarioEnExcel", &controllers.GenerarDiarioEnExcelController{})
    beego.Router("/generarDiario", &controllers.GenerarDiarioEnPDFController{})
    beego.Router("/registraDiarioAPI", &controllers.RegistraDiarioAPIController{})
    beego.Router("/registraDiario", &controllers.RegistraDiarioController{})
    beego.Router("/cambiarUnidadDeNegocio", &controllers.CambiarUnidadDeNegocioController{})
    beego.Router("/guardarBUNIT", &controllers.GuardarBUNITController{})
    beego.Router("/tipoDeDimensiones", &controllers.TipoDeDimensionesController{})
    beego.Router("/generarBUNIT", &controllers.GenerarBUNITController{})
	beego.Router("/nuevoTipoDeDimensiones", &controllers.NuevoTipoDeDimensionesController{})
	beego.Router("/editarTipoDeDimensiones", &controllers.EditarTipoDeDimensionesController{})
	beego.Router("/nuevaDimension", &controllers.NuevaDimensionController{})
	beego.Router("/dimensiones", &controllers.DimensionesController{})
	beego.Router("/editarDimension", &controllers.EditarDimensionController{})
	beego.Router("/clasificacionDeDimensiones", &controllers.ClasificacionDimensionController{})
	beego.Router("/obtieneDimensionesAsignadasATipo", &controllers.ObtieneDimensionController{})
	beego.Router("/guardarClasificacionDimensiones", &controllers.GuardarClasificacionDimensionesController{})
	beego.Router("/nuevaCuenta", &controllers.NuevaCuentaController{})
	beego.Router("/listaCedulas", &controllers.ListaCedulasController{})
	beego.Router("/nuevoCedula", &controllers.NuevaCedulasController{})
	beego.Router("/configInicialLineasCedulas", &controllers.ConfigInicialLineasCedulasController{})
	beego.Router("/nuevoLineaCedula", &controllers.NuevoLineasCedulasController{})
	beego.Router("/listaLineasCedulas", &controllers.ListaLineasCedulasController{})
	beego.Router("/generaCedula", &controllers.GenerarCedulasController{})
	beego.Router("/guardarConceptoCedula", &controllers.GuardarConceptoCedulasController{})
	beego.Router("/listaConceptosCedulas", &controllers.ListaConceptosCedulasController{})
	beego.Router("/generaCedulaPorConcepto", &controllers.GenerarCedulaPorConceptoController{})
	beego.Router("/generaCedulaPorLinea", &controllers.GenerarCedulaPorLineaController{})
	beego.Router("/configInicialDiariosReversiados", &controllers.DiariosReversiadosController{})
	beego.Router("/nuevoTipoDeDiario", &controllers.NuevoTipoDeDiarioController{})
	beego.Router("/listaTiposDeDiario", &controllers.ListaTiposDeDiarioController{})
	beego.Router("/configInicialDeLineasDeTiposDeDiario", &controllers.ConfigInicialLineasTiposDeDiarioController{})
	beego.Router("/nuevoOModificaLineaTipoDeDiario", &controllers.NuevoLineasTiposDeDiarioController{})
	beego.Router("/listaLineasTiposDeDiario", &controllers.ListaLineasTiposDeDiarioController{})
	beego.Router("/dameLineaDelTipoDeDiario", &controllers.DameLineaDelTipoDeDiarioController{})
	beego.Router("/dameDimensionesDisponiblesSegunLaCuenta", &controllers.DameDimensionesDisponiblesSegunLaCuentaController{})
	beego.Router("/guardaLineaEnTemporalesDiario", &controllers.NuevoLineasDeDiarioController{})
	beego.Router("/contabilizaPorTimeStamp", &controllers.ContabilizaDiarioController{})
	beego.Router("/dashboard", &controllers.DashboardController{})
	beego.Router("/veDetallePrimerNivel", &controllers.VeDetallePrimerNivelController{})
	beego.Router("/veDetalle", &controllers.VeDetalleController{})
}






