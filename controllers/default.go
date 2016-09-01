package controllers
import (
	"github.com/astaxie/beego"
	_ "github.com/denisenkom/go-mssqldb"
	"database/sql"
	"log"
	"fmt"
	"flag"	
	"bytes"
	"bufio"
	"encoding/base64"
	"crypto/sha512"
	"strconv"
	"sevenplus/models"
	"github.com/jung-kurt/gofpdf"
	"math"
	"github.com/dimdin/decimal"
	"time"
	"github.com/leekchan/accounting"
	"encoding/csv"
	"strings"
	"encoding/hex"
	"encoding/json"
	"net/http"
	//"net/url"
	"io/ioutil"
	
)



type MainController struct {
	beego.Controller
}
type DashboardController struct {
	beego.Controller
}
type LoginController struct {
	beego.Controller
}
type GenerarDiarioEnExcelController struct {
	beego.Controller
}
type GenerarDiarioEnPDFController struct {
	beego.Controller
}
type GenerarDiarioController struct {
	beego.Controller
}

type RegistraDiarioAPIController struct {
	beego.Controller
}

type GenerarEstadosDeCuentaController struct {
	beego.Controller
}

type CambiarUnidadDeNegocioController struct {
	beego.Controller
}

type GuardarBUNITController struct {
	beego.Controller
}

type GenerarBUNITController struct {
	beego.Controller
}

type EstadosDeCuentaController struct {
	beego.Controller
}

type TipoDeDimensionesController struct {
	beego.Controller
}

type NuevoTipoDeDimensionesController struct {
	beego.Controller
}

type EditarTipoDeDimensionesController struct {
	beego.Controller
}
type NuevaDimensionController struct {
	beego.Controller
}
type DimensionesController struct {
	beego.Controller
}
type EditarDimensionController struct {
	beego.Controller
}
type ClasificacionDimensionController struct {
	beego.Controller
}
type ObtieneDimensionController struct {
	beego.Controller
}
type GuardarClasificacionDimensionesController struct {
	beego.Controller
}
type NuevaCuentaController struct {
	beego.Controller
}
type ListaConceptosCedulasController struct {
	beego.Controller
}
type ListaLineasCedulasController struct {
	beego.Controller
}
type ChecaOtrosCamposController struct {
	beego.Controller
}
type ListaLineasTiposDeDiarioController struct {
	beego.Controller
}
type ListaCedulasController struct {
	beego.Controller
}
type NuevaCedulasController struct {
	beego.Controller
}
type ConfigInicialLineasCedulasController struct {
	beego.Controller
}
type ConfigInicialLineasTiposDeDiarioController struct {
	beego.Controller
}
type NuevoLineasCedulasController struct {
	beego.Controller
}
type NuevoLineasTiposDeDiarioController struct {
	beego.Controller
}
type NuevoLineasDeDiarioController struct {
	beego.Controller
}
type GuardarConfigOpcionesController struct {
	beego.Controller
}
type ContabilizaDiarioController struct {
	beego.Controller
}
type GenerarCedulasController struct {
	beego.Controller
}
type GuardarConceptoCedulasController struct {
	beego.Controller
}
type GenerarCedulaPorConceptoController struct {
	beego.Controller
}
type GenerarCedulaPorLineaController struct {
	beego.Controller
}
type RegistraDiarioController struct {
	beego.Controller
}
type DiariosReversiadosController struct {
	beego.Controller
}
type NuevoTipoDeDiarioController struct {
	beego.Controller
}
type ListaTiposDeDiarioController struct {
	beego.Controller
}
type ListaActivosFijosController struct {
	beego.Controller
}
type DameLineaDelTipoDeDiarioController struct {
	beego.Controller
}
type DameDimensionesDisponiblesSegunLaCuentaController struct {
	beego.Controller
}
type VeDetalleController struct {
	beego.Controller
}
type VeDetallePrimerNivelController struct {
	beego.Controller
}
type GetConfigController struct {
	beego.Controller
}
type GenerarReporteDeIglesiasController struct {
	beego.Controller
}
type GenerarReporteDeMATController struct {
	beego.Controller
}




var debug = flag.Bool("debug", false, "enable debugging")
var password = flag.String("password", "SunPlus7!", "the database password")
var port *int = flag.Int("port", 1433, "the database port")
var server = flag.String("server", "localhost", "the database server")
var user = flag.String("user", "sa", "the database user")

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["mensaje"] = "Sevenplus"
	c.TplName = "index.tpl"
}
		
func (c *GenerarDiarioController) Get() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		//BUNIT := c.GetSession("BUNIT")
	//	diario := c.GetString("diario")	
	}
}
type respuesta struct {
     hash string 
     poliza []linea
}
type linea struct {
     PERIOD string 
     treference string
     Fecha string 
     cuenta string
     fondo string
     funcion string
     reestriccion string
     orgid string
     who string
     detalle string
     proyecto string 
     descripcion string
     cantidad float64
}
/*
type Pair struct {
  Key string
  Value interface{}
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { 
	num1, _  := strconv.Atoi(p[i].Key)
	num2, _  := strconv.Atoi(p[j].Key)
	return num1 > num2
}

// A function to turn a map into a PairList, then sort and return it. 
func sortMapByValue(m map[string]interface{}) PairList {
   p := make(PairList, len(m))
   i := 0
   for k, v := range m {
      p[i] = Pair{k, v}
   }
   sort.Sort(p)
   return p
}
*/


func (c * RegistraDiarioController) Post() {
	alias := c.GetSession("alias")
	if alias == nil {
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, errS := sql.Open("mssql", connString2)
		if errS != nil {
			log.Fatal("Perdon! Open connection failed:", errS.Error())
		}
		var JRNAL_NO int64
		query1 := "SELECT ISNULL((MAX(JRNAL_NO)+1),1) as JRNAL_NO FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_A_SALFLDG]"
   		rows1, err1 := conn.Query(query1)
   		if err1 != nil {
   			defer conn.Close()
			log.Fatal("Perdon! Error en querys hardcodeados: ", err1.Error())
		}
   		for rows1.Next()  {
			rows1.Scan(&JRNAL_NO)
		}
		jsonVariable := c.GetString("lineas")
		b := []byte(jsonVariable)
		var m map[string]interface{}
		err := json.Unmarshal(b, &m)
		if err != nil {
			log.Fatal("JSON Parse failed (perdon):", err.Error())
		}
		//var poliza PairList 
		//poliza =sortMapByValue(m)
		//generar un array map asociativo, ponerle los keys, values como vayan llegando y luego iterarlo en orden!
		var linea map[string]interface{}
		var D_C string
		D_C="C"
   		for k, v := range m { 
			JRNAL_LINE,error1:=strconv.Atoi(k);
			if error1 != nil {
				defer conn.Close()
				log.Fatal("Perdon! : ", error1.Error())
			}
			JRNAL_LINE=JRNAL_LINE+1
			linea = v.(map[string]interface{});
			cantidad := linea["cantidad"].(float64)
			if cantidad < 0 {
				D_C="D";
			} else {
				D_C="C";
			}
			amountString := fmt.Sprintf("%.2f", cantidad)
			diario := strconv.FormatInt(JRNAL_NO, 10)
			lineaDelDiario := strconv.Itoa(JRNAL_LINE)
			treference := linea["treference"]
			cuenta := linea["cuenta"]
			query := "INSERT INTO [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_A_SALFLDG] (TREFERENCE, ACCNT_CODE, ANAL_T0, ANAL_T1, ANAL_T2, ANAL_T3, ANAL_T4, ANAL_T5, ANAL_T6, ANAL_T7, ANAL_T8, ANAL_T9, DESCRIPTN, AMOUNT, OTHER_AMT, CONV_CODE, D_C, PERIOD, TRANS_DATETIME, JRNAL_NO, JRNAL_LINE, JRNAL_TYPE, JRNAL_SRCE, ENTRY_DATETIME, ENTRY_PRD, DUE_DATETIME, CONV_RATE, POSTING_DATETIME, CV4_CONV_CODE, CV4_AMT, CV4_CONV_RATE, CV4_OPERATOR, ORIGINATOR_ID, POSTER_ID, JNL_REVERSAL_TYPE, SPLIT_IN_PROGRESS, AUTHORISTN_IN_PROGRESS, MAN_PAY_OVER, APRVLS_EXTSN, SUPPLMNTRY_EXTSN, TRUE_RATED, AGREED_STATUS, BINDER_STATUS, PRINCIPAL_REQD, SPLIT_ORIG_LINE, CV5_DP, CV5_OPERATOR, CV5_CONV_RATE, CV5_AMT, CV5_CONV_CODE, CV4_DP, CONSUMED_BDGT_ID, LE_DETAILS_IND, EXCLUDE_BAL, MEMO_AMT, REPORT_AMT, REPORT_OPERATOR, REPORT_RATE, CONV_OPERATOR, BASE_OPERATOR, BASE_RATE, HOLD_OP_ID, HOLD_REF, ALLOC_IN_PROGRESS, ALLOCATION, ALLOC_REF, ALLOC_PERIOD, ASSET_IND, ASSET_CODE, ASSET_SUB, OTHER_DP, CLEARDOWN, REVERSAL, LOSS_GAIN, ROUGH_FLAG, IN_USE_FLAG, ORIGINATED_DATETIME) VALUES ('"+treference.(string)+"','"+cuenta.(string)+"', '', '', '"+linea["fondo"].(string)+"', '"+linea["funcion"].(string)+"', '"+linea["reestriccion"].(string)+"', '"+linea["orgid"].(string)+"', '"+linea["who"].(string)+"', '', '"+linea["proyecto"].(string)+"', '"+linea["detalle"].(string)+"', '"+linea["descripcion"].(string)+"', "+amountString+", "+amountString+", 'MXP1', '"+D_C+"', "+linea["PERIOD"].(string)+", '"+linea["fecha"].(string)+"', "+diario+", "+lineaDelDiario+", 'JV', 'MPJ', '"+linea["fecha"].(string)+"', "+linea["PERIOD"].(string)+", '"+linea["fecha"].(string)+"', 1.0000000, '"+linea["fecha"].(string)+"', 'MXP1', "+amountString+", 1.0000000, '*', 'MPJ', 'MPJ', 0, 0, 0, 0 ,0 ,0, 0, 0, '', 0, 0, '2', '*', 0.00, 0.00, '', '2', 0, '', '', 0.0, 0.0, '*', 0.0, '*', '*', 1.0000, '', 0, '', '', 0, 0, '', '', '', '2', '00000', '', '', '', '', '"+linea["fecha"].(string)+"')"
   			result , err2 := conn.Exec(query)
			if err2 != nil {
   				defer conn.Close()
				log.Fatal("Perdon 5! : ", err2.Error())
			}
			fmt.Println(result.RowsAffected())
		}
		example := map[string]interface{}{ "success":1, "diario" : JRNAL_NO}
		c.Data["json"] = &example
		c.ServeJSON()
		defer conn.Close()
	}
}
func (c * RegistraDiarioAPIController) Post() {
	hash := c.GetString("hash")
	currentTime := int64(time.Now().Unix())
	tm := time.Unix(currentTime, 0)
	dateString := tm.String() 
	substring := string(dateString[0:10])
   	byteArray := []byte(substring)
	hasher := sha512.New()
    hasher.Write(byteArray)
	cryptoText := base64.StdEncoding.EncodeToString(  []byte(hex.EncodeToString(hasher.Sum(nil))))
	if Compare(cryptoText,hash)==0 {
		BUNIT := c.GetString("BUNIT")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, errS := sql.Open("mssql", connString2)
		if errS != nil {
			log.Fatal("Perdon! Open connection failed:", errS.Error())
		}
		var JRNAL_NO int64
		query1 := "SELECT ISNULL((MAX(JRNAL_NO)+1),1) as JRNAL_NO FROM [SUNPLUSADV].[dbo].["+BUNIT+"_A_SALFLDG]"
   		rows1, err1 := conn.Query(query1)
   		if err1 != nil {
   			defer conn.Close()
			log.Fatal("Perdon! Error en querys hardcodeados: ", err1.Error())
		}
   		for rows1.Next()  {
			rows1.Scan(&JRNAL_NO)
		}
		fmt.Println(JRNAL_NO)
		jsonVariable := c.GetString("lineas")
		b := []byte(jsonVariable)
		var m map[string]interface{}
		err := json.Unmarshal(b, &m)
		if err != nil {
			log.Fatal("JSON Parse failed (perdon):", err.Error())
		}
		//var poliza PairList 
		//poliza =sortMapByValue(m)
		//generar un array map asociativo, ponerle los keys, values como vayan llegando y luego iterarlo en orden!
		var linea map[string]interface{}
		var D_C string
		D_C="C"
   		for k, v := range m { 
			JRNAL_LINE,error1:=strconv.Atoi(k);
			if error1 != nil {
				defer conn.Close()
				log.Fatal("Perdon! : ", error1.Error())
			}
			JRNAL_LINE=JRNAL_LINE+1
			linea = v.(map[string]interface{});
			cantidad := linea["cantidad"].(float64)
			if cantidad < 0 {
				D_C="D";
			} else {
				D_C="C";
			}
			amountString := fmt.Sprintf("%.2f", cantidad)
			diario := strconv.FormatInt(JRNAL_NO, 10)
			lineaDelDiario := strconv.Itoa(JRNAL_LINE)
			treference := linea["treference"]
			cuenta := linea["cuenta"]
			query := "INSERT INTO ["+beego.AppConfig.String("serverAux")+"].[SUNPLUSADV].[dbo].["+BUNIT+"_A_SALFLDG] (TREFERENCE, ACCNT_CODE, ANAL_T0, ANAL_T1, ANAL_T2, ANAL_T3, ANAL_T4, ANAL_T5, ANAL_T6, ANAL_T7, ANAL_T8, ANAL_T9, DESCRIPTN, AMOUNT, OTHER_AMT, CONV_CODE, D_C, PERIOD, TRANS_DATETIME, JRNAL_NO, JRNAL_LINE, JRNAL_TYPE, JRNAL_SRCE, ENTRY_DATETIME, ENTRY_PRD, DUE_DATETIME, CONV_RATE, POSTING_DATETIME, CV4_CONV_CODE, CV4_AMT, CV4_CONV_RATE, CV4_OPERATOR, ORIGINATOR_ID, POSTER_ID, JNL_REVERSAL_TYPE, SPLIT_IN_PROGRESS, AUTHORISTN_IN_PROGRESS, MAN_PAY_OVER, APRVLS_EXTSN, SUPPLMNTRY_EXTSN, TRUE_RATED, AGREED_STATUS, BINDER_STATUS, PRINCIPAL_REQD, SPLIT_ORIG_LINE, CV5_DP, CV5_OPERATOR, CV5_CONV_RATE, CV5_AMT, CV5_CONV_CODE, CV4_DP, CONSUMED_BDGT_ID, LE_DETAILS_IND, EXCLUDE_BAL, MEMO_AMT, REPORT_AMT, REPORT_OPERATOR, REPORT_RATE, CONV_OPERATOR, BASE_OPERATOR, BASE_RATE, HOLD_OP_ID, HOLD_REF, ALLOC_IN_PROGRESS, ALLOCATION, ALLOC_REF, ALLOC_PERIOD, ASSET_IND, ASSET_CODE, ASSET_SUB, OTHER_DP, CLEARDOWN, REVERSAL, LOSS_GAIN, ROUGH_FLAG, IN_USE_FLAG, ORIGINATED_DATETIME) VALUES ('"+treference.(string)+"','"+cuenta.(string)+"', '', '', '"+linea["fondo"].(string)+"', '"+linea["funcion"].(string)+"', '"+linea["reestriccion"].(string)+"', '"+linea["orgid"].(string)+"', '"+linea["who"].(string)+"', '', '"+linea["proyecto"].(string)+"', '"+linea["detalle"].(string)+"', '"+linea["descripcion"].(string)+"', "+amountString+", "+amountString+", 'MXP1', '"+D_C+"', "+linea["PERIOD"].(string)+", '"+linea["fecha"].(string)+"', "+diario+", "+lineaDelDiario+", 'JV', 'MPJ', '"+linea["fecha"].(string)+"', "+linea["PERIOD"].(string)+", '"+linea["fecha"].(string)+"', 1.0000000, '"+linea["fecha"].(string)+"', 'MXP1', "+amountString+", 1.0000000, '*', 'MPJ', 'MPJ', 0, 0, 0, 0 ,0 ,0, 0, 0, '', 0, 0, '2', '*', 0.00, 0.00, '', '2', 0, '', '', 0.0, 0.0, '*', 0.0, '*', '*', 1.0000, '', 0, '', '', 0, 0, '', '', '', '2', '00000', '', '', '', '', '"+linea["fecha"].(string)+"')"
   			fmt.Println(query)
			result , err2 := conn.Exec(query)
			if err2 != nil {
   				defer conn.Close()
				log.Fatal("Perdon 5! : ", err2.Error())
			}
			fmt.Println(result.RowsAffected())
		}
		defer conn.Close()
	}
}
func tienePermisosContador(tipo int) bool {
	if tipo == 3{
		return true
	}
	return false
}



func (c *GuardarClasificacionDimensionesController) Post() {
	alias := c.GetSession("alias")
	if alias == nil {
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		TipoDondeEstaranLasDimensiones := c.GetString("TipoDondeEstaranLasDimensiones")
		entry1 := c.GetString("entry1")
		entry2 := c.GetString("entry2")
		entry3 := c.GetString("entry3")
		entry4 := c.GetString("entry4")
		entry5 := c.GetString("entry5")
		entry6 := c.GetString("entry6")
		entry7 := c.GetString("entry7")
		entry8 := c.GetString("entry8")
		entry9 := c.GetString("entry9")
		entry10 := c.GetString("entry10")
		S_HEAD1 := c.GetString("S_HEAD1")
		S_HEAD2 := c.GetString("S_HEAD2")
		S_HEAD3 := c.GetString("S_HEAD3")
		S_HEAD4 := c.GetString("S_HEAD4")
		S_HEAD5 := c.GetString("S_HEAD5")
		S_HEAD6 := c.GetString("S_HEAD6")
		S_HEAD7 := c.GetString("S_HEAD7")
		S_HEAD8 := c.GetString("S_HEAD8")
		S_HEAD9 := c.GetString("S_HEAD9")
		S_HEAD10 := c.GetString("S_HEAD10")
		//primero debo de checar si existe la combinacion Tipo- entryNum
		//si existe y entryX es 0, debo borrarla
		//si no existe Tipo-entryNum y entryX!=0hago un insert
		//si existe y entryX !=0, hago un update!
		var ANL_CAT_ID string
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, errS := sql.Open("mssql", connString2)
		if errS != nil {
			log.Fatal("Perdon! Open connection failed:", errS.Error())
		}
		currentTime := int64(time.Now().Unix())
		tm := time.Unix(currentTime, 0)
		dateString := tm.String() 
		substring := string(dateString[0:10])	

		//1 {
		queryDireccion := "SELECT ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_ENT_DEFN] WHERE ENTRY_NUM = 1 AND ANL_ENT_ID = "+TipoDondeEstaranLasDimensiones
		rowsDireccion, _ := conn.Query(queryDireccion)
		if rowsDireccion.Next()  {
			rowsDireccion.Scan(&ANL_CAT_ID)
			if Compare(entry1,"0")==0 {
				query := `DELETE FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN] WHERE ENTRY_NUM = 1 AND ANL_ENT_ID = `+TipoDondeEstaranLasDimensiones
				conn.Exec(query)
			} else {
				if Compare(ANL_CAT_ID,entry1)==0 {
				} else {
					query1 := `UPDATE [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN] SET LAST_CHANGE_DATETIME = '`+substring+`',LAST_CHANGE_USER_ID = '`+alias.(string)+`', ANL_CAT_ID = '`+entry1+`', S_HEAD = '`+S_HEAD1+`' WHERE ENTRY_NUM = 1 AND ANL_ENT_ID = `+TipoDondeEstaranLasDimensiones
					conn.Exec(query1)
				}
			}
		}
		query := `INSERT INTO [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN]
			(ANL_ENT_ID,ANL_CAT_ID,UPDATE_COUNT,LAST_CHANGE_USER_ID,LAST_CHANGE_DATETIME,ENTRY_NUM,VALIDATE_IND,S_HEAD)
	VALUES(`+TipoDondeEstaranLasDimensiones+`,'`+entry1+`',1,'`+alias.(string)+`','`+substring+`',1,1,'`+S_HEAD1+`')`
		conn.Exec(query)
		//1 }
		//2 {
		queryDireccion = "SELECT ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_ENT_DEFN] WHERE ENTRY_NUM = 2 AND ANL_ENT_ID = "+TipoDondeEstaranLasDimensiones
		rowsDireccion, _ = conn.Query(queryDireccion)
		if rowsDireccion.Next()  {
			rowsDireccion.Scan(&ANL_CAT_ID)
			if Compare(entry2,"0")==0 {
				query := `DELETE FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN] WHERE ENTRY_NUM = 2 AND ANL_ENT_ID = `+TipoDondeEstaranLasDimensiones
				conn.Exec(query)
			} else {
				if Compare(ANL_CAT_ID,entry2)==0 {
				} else {
					query1 := `UPDATE [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN] SET LAST_CHANGE_DATETIME = '`+substring+`',LAST_CHANGE_USER_ID = '`+alias.(string)+`', ANL_CAT_ID = '`+entry2+`', S_HEAD = '`+S_HEAD2+`' WHERE ENTRY_NUM = 2 AND ANL_ENT_ID = `+TipoDondeEstaranLasDimensiones
					conn.Exec(query1)
				}
			}
		}
		query = `INSERT INTO [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN]
			(ANL_ENT_ID,ANL_CAT_ID,UPDATE_COUNT,LAST_CHANGE_USER_ID,LAST_CHANGE_DATETIME,ENTRY_NUM,VALIDATE_IND,S_HEAD)
	VALUES(`+TipoDondeEstaranLasDimensiones+`,'`+entry2+`',1,'`+alias.(string)+`','`+substring+`',2,1,'`+S_HEAD2+`')`
		conn.Exec(query)
		//2 }
		//3 {
		queryDireccion = "SELECT ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_ENT_DEFN] WHERE ENTRY_NUM = 3 AND ANL_ENT_ID = "+TipoDondeEstaranLasDimensiones
		rowsDireccion, _ = conn.Query(queryDireccion)
		if rowsDireccion.Next()  {
			rowsDireccion.Scan(&ANL_CAT_ID)
			if Compare(entry3,"0")==0 {
				query := `DELETE FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN] WHERE ENTRY_NUM = 3 AND ANL_ENT_ID = `+TipoDondeEstaranLasDimensiones
				conn.Exec(query)
			} else {
				if Compare(ANL_CAT_ID,entry3)==0 {
				} else {
					query1 := `UPDATE [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN] SET LAST_CHANGE_DATETIME = '`+substring+`',LAST_CHANGE_USER_ID = '`+alias.(string)+`', ANL_CAT_ID = '`+entry3+`', S_HEAD = '`+S_HEAD3+`' WHERE ENTRY_NUM = 3 AND ANL_ENT_ID = `+TipoDondeEstaranLasDimensiones
					conn.Exec(query1)
				}
			}
		}
		query = `INSERT INTO [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN]
			(ANL_ENT_ID,ANL_CAT_ID,UPDATE_COUNT,LAST_CHANGE_USER_ID,LAST_CHANGE_DATETIME,ENTRY_NUM,VALIDATE_IND,S_HEAD)
	VALUES(`+TipoDondeEstaranLasDimensiones+`,'`+entry3+`',1,'`+alias.(string)+`','`+substring+`',3,1,'`+S_HEAD3+`')`
		conn.Exec(query)
		//3 }
		//4 {
		queryDireccion = "SELECT ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_ENT_DEFN] WHERE ENTRY_NUM = 4 AND ANL_ENT_ID = "+TipoDondeEstaranLasDimensiones
		rowsDireccion, _ = conn.Query(queryDireccion)
		if rowsDireccion.Next()  {
			rowsDireccion.Scan(&ANL_CAT_ID)
			if Compare(entry4,"0")==0 {
				query := `DELETE FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN] WHERE ENTRY_NUM = 4 AND ANL_ENT_ID = `+TipoDondeEstaranLasDimensiones
				conn.Exec(query)
			} else {
				if Compare(ANL_CAT_ID,entry4)==0 {
				} else {
					query1 := `UPDATE [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN] SET LAST_CHANGE_DATETIME = '`+substring+`',LAST_CHANGE_USER_ID = '`+alias.(string)+`', ANL_CAT_ID = '`+entry4+`', S_HEAD = '`+S_HEAD4+`' WHERE ENTRY_NUM = 4 AND ANL_ENT_ID = `+TipoDondeEstaranLasDimensiones
					conn.Exec(query1)
				}
			}
		}
		query = `INSERT INTO [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN]
			(ANL_ENT_ID,ANL_CAT_ID,UPDATE_COUNT,LAST_CHANGE_USER_ID,LAST_CHANGE_DATETIME,ENTRY_NUM,VALIDATE_IND,S_HEAD)
	VALUES(`+TipoDondeEstaranLasDimensiones+`,'`+entry4+`',1,'`+alias.(string)+`','`+substring+`',4,1,'`+S_HEAD4+`')`
		conn.Exec(query)
		//4 }
		//5 {
		queryDireccion = "SELECT ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_ENT_DEFN] WHERE ENTRY_NUM = 5 AND ANL_ENT_ID = "+TipoDondeEstaranLasDimensiones
		rowsDireccion, _ = conn.Query(queryDireccion)
		if rowsDireccion.Next()  {
			rowsDireccion.Scan(&ANL_CAT_ID)
			if Compare(entry5,"0")==0 {
				query := `DELETE FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN] WHERE ENTRY_NUM = 5 AND ANL_ENT_ID = `+TipoDondeEstaranLasDimensiones
				conn.Exec(query)
			} else {
				if Compare(ANL_CAT_ID,entry5)==0 {
				} else {
					query1 := `UPDATE [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN] SET LAST_CHANGE_DATETIME = '`+substring+`',LAST_CHANGE_USER_ID = '`+alias.(string)+`', ANL_CAT_ID = '`+entry5+`', S_HEAD = '`+S_HEAD5+`' WHERE ENTRY_NUM = 5 AND ANL_ENT_ID = `+TipoDondeEstaranLasDimensiones
					conn.Exec(query1)
				}
			}
		}
		query = `INSERT INTO [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN]
			(ANL_ENT_ID,ANL_CAT_ID,UPDATE_COUNT,LAST_CHANGE_USER_ID,LAST_CHANGE_DATETIME,ENTRY_NUM,VALIDATE_IND,S_HEAD)
	VALUES(`+TipoDondeEstaranLasDimensiones+`,'`+entry5+`',1,'`+alias.(string)+`','`+substring+`',5,1,'`+S_HEAD5+`')`
		conn.Exec(query)
		//5 }
		//6 {
		queryDireccion = "SELECT ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_ENT_DEFN] WHERE ENTRY_NUM = 6 AND ANL_ENT_ID = "+TipoDondeEstaranLasDimensiones
		rowsDireccion, _ = conn.Query(queryDireccion)
		if rowsDireccion.Next()  {
			rowsDireccion.Scan(&ANL_CAT_ID)
			if Compare(entry6,"0")==0 {
				query := `DELETE FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN] WHERE ENTRY_NUM = 6 AND ANL_ENT_ID = `+TipoDondeEstaranLasDimensiones
				conn.Exec(query)
			} else {
				if Compare(ANL_CAT_ID,entry6)==0 {
				} else {
					query1 := `UPDATE [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN] SET LAST_CHANGE_DATETIME = '`+substring+`',LAST_CHANGE_USER_ID = '`+alias.(string)+`', ANL_CAT_ID = '`+entry6+`', S_HEAD = '`+S_HEAD6+`' WHERE ENTRY_NUM = 6 AND ANL_ENT_ID = `+TipoDondeEstaranLasDimensiones
					conn.Exec(query1)
				}
			}
		}
		query = `INSERT INTO [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN]
			(ANL_ENT_ID,ANL_CAT_ID,UPDATE_COUNT,LAST_CHANGE_USER_ID,LAST_CHANGE_DATETIME,ENTRY_NUM,VALIDATE_IND,S_HEAD)
	VALUES(`+TipoDondeEstaranLasDimensiones+`,'`+entry6+`',1,'`+alias.(string)+`','`+substring+`',6,1,'`+S_HEAD6+`')`
		conn.Exec(query)
		//6 }
		//7 {
		queryDireccion = "SELECT ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_ENT_DEFN] WHERE ENTRY_NUM = 7 AND ANL_ENT_ID = "+TipoDondeEstaranLasDimensiones
		rowsDireccion, _ = conn.Query(queryDireccion)
		if rowsDireccion.Next()  {
			rowsDireccion.Scan(&ANL_CAT_ID)
			if Compare(entry7,"0")==0 {
				query := `DELETE FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN] WHERE ENTRY_NUM = 7 AND ANL_ENT_ID = `+TipoDondeEstaranLasDimensiones
				conn.Exec(query)
			} else {
				if Compare(ANL_CAT_ID,entry7)==0 {
				} else {
					query1 := `UPDATE [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN] SET LAST_CHANGE_DATETIME = '`+substring+`',LAST_CHANGE_USER_ID = '`+alias.(string)+`', ANL_CAT_ID = '`+entry7+`', S_HEAD = '`+S_HEAD7+`' WHERE ENTRY_NUM = 7 AND ANL_ENT_ID = `+TipoDondeEstaranLasDimensiones
					conn.Exec(query1)
				}
			}
		}
		query = `INSERT INTO [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN]
			(ANL_ENT_ID,ANL_CAT_ID,UPDATE_COUNT,LAST_CHANGE_USER_ID,LAST_CHANGE_DATETIME,ENTRY_NUM,VALIDATE_IND,S_HEAD)
	VALUES(`+TipoDondeEstaranLasDimensiones+`,'`+entry7+`',1,'`+alias.(string)+`','`+substring+`',7,1,'`+S_HEAD7+`')`
		conn.Exec(query)
		//7 }
		//8 {
		queryDireccion = "SELECT ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_ENT_DEFN] WHERE ENTRY_NUM = 8 AND ANL_ENT_ID = "+TipoDondeEstaranLasDimensiones
		rowsDireccion, _ = conn.Query(queryDireccion)
		if rowsDireccion.Next()  {
			rowsDireccion.Scan(&ANL_CAT_ID)
			if Compare(entry8,"0")==0 {
				query := `DELETE FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN] WHERE ENTRY_NUM = 8 AND ANL_ENT_ID = `+TipoDondeEstaranLasDimensiones
				conn.Exec(query)
			} else {
				if Compare(ANL_CAT_ID,entry8)==0 {
				} else {
					query1 := `UPDATE [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN] SET LAST_CHANGE_DATETIME = '`+substring+`',LAST_CHANGE_USER_ID = '`+alias.(string)+`', ANL_CAT_ID = '`+entry8+`', S_HEAD = '`+S_HEAD8+`' WHERE ENTRY_NUM = 8 AND ANL_ENT_ID = `+TipoDondeEstaranLasDimensiones
					conn.Exec(query1)
				}
			}
		}
		query = `INSERT INTO [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN]
			(ANL_ENT_ID,ANL_CAT_ID,UPDATE_COUNT,LAST_CHANGE_USER_ID,LAST_CHANGE_DATETIME,ENTRY_NUM,VALIDATE_IND,S_HEAD)
	VALUES(`+TipoDondeEstaranLasDimensiones+`,'`+entry8+`',1,'`+alias.(string)+`','`+substring+`',8,1,'`+S_HEAD8+`')`
		conn.Exec(query)
		//8 }
		//9 {
		queryDireccion = "SELECT ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_ENT_DEFN] WHERE ENTRY_NUM = 9 AND ANL_ENT_ID = "+TipoDondeEstaranLasDimensiones
		rowsDireccion, _ = conn.Query(queryDireccion)
		if rowsDireccion.Next()  {
			rowsDireccion.Scan(&ANL_CAT_ID)
			if Compare(entry9,"0")==0 {
				query := `DELETE FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN] WHERE ENTRY_NUM = 9 AND ANL_ENT_ID = `+TipoDondeEstaranLasDimensiones
				conn.Exec(query)
			} else {
				if Compare(ANL_CAT_ID,entry9)==0 {
				} else {
					query1 := `UPDATE [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN] SET LAST_CHANGE_DATETIME = '`+substring+`',LAST_CHANGE_USER_ID = '`+alias.(string)+`', ANL_CAT_ID = '`+entry9+`', S_HEAD = '`+S_HEAD9+`' WHERE ENTRY_NUM = 9 AND ANL_ENT_ID = `+TipoDondeEstaranLasDimensiones
					conn.Exec(query1)
				}
			}
		}
		query = `INSERT INTO [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN]
			(ANL_ENT_ID,ANL_CAT_ID,UPDATE_COUNT,LAST_CHANGE_USER_ID,LAST_CHANGE_DATETIME,ENTRY_NUM,VALIDATE_IND,S_HEAD)
	VALUES(`+TipoDondeEstaranLasDimensiones+`,'`+entry9+`',1,'`+alias.(string)+`','`+substring+`',9,1,'`+S_HEAD9+`')`
		conn.Exec(query)
		//9 }
		//10 {
		queryDireccion = "SELECT ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_ENT_DEFN] WHERE ENTRY_NUM = 10 AND ANL_ENT_ID = "+TipoDondeEstaranLasDimensiones
		rowsDireccion, _ = conn.Query(queryDireccion)
		if rowsDireccion.Next()  {
			rowsDireccion.Scan(&ANL_CAT_ID)
			if Compare(entry10,"0")==0 {
				query := `DELETE FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN] WHERE ENTRY_NUM = 10 AND ANL_ENT_ID = `+TipoDondeEstaranLasDimensiones
				conn.Exec(query)
			} else {
				if Compare(ANL_CAT_ID,entry10)==0 {
				} else {
					query1 := `UPDATE [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN] SET LAST_CHANGE_DATETIME = '`+substring+`',LAST_CHANGE_USER_ID = '`+alias.(string)+`', ANL_CAT_ID = '`+entry10+`', S_HEAD = '`+S_HEAD10+`' WHERE ENTRY_NUM = 10 AND ANL_ENT_ID = `+TipoDondeEstaranLasDimensiones
					conn.Exec(query1)
				}
			}
		}
		query = `INSERT INTO [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_ENT_DEFN]
			(ANL_ENT_ID,ANL_CAT_ID,UPDATE_COUNT,LAST_CHANGE_USER_ID,LAST_CHANGE_DATETIME,ENTRY_NUM,VALIDATE_IND,S_HEAD)
	VALUES(`+TipoDondeEstaranLasDimensiones+`,'`+entry10+`',1,'`+alias.(string)+`','`+substring+`',10,1,'`+S_HEAD10+`')`
		conn.Exec(query)
		//10 }
		example := map[string]interface{}{ "success":1}
		c.Data["json"] = &example
		c.ServeJSON()
	}
} 

func (c *EditarDimensionController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		tipo := c.GetString("tipo")
		codigo := c.GetString("codigo")
		status := c.GetString("status")
		nombre := c.GetString("nombre")
		prohibido := c.GetString("prohibido")
		var ANL_CAT_ID string
		ANL_CAT_ID = tipo
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, errS := sql.Open("mssql", connString2)
		if errS != nil {
			log.Fatal("Perdon! Open connection failed:", errS.Error())
		}
		currentTime := int64(time.Now().Unix())
		tm := time.Unix(currentTime, 0)
		dateString := tm.String() 
		substring := string(dateString[0:10])	


		queryDireccion := "SELECT TOP 1 ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_CODE] WHERE ANL_CODE = '"+codigo+"' AND ANL_CAT_ID = '"+ANL_CAT_ID+"' order by ANL_CAT_ID desc"
		rowsDireccion, _ := conn.Query(queryDireccion)
		if rowsDireccion.Next()  {
			maximo := len(nombre)
			var LOOKUP string
			if maximo > 14 {
				maximo = 14
			}
			LOOKUP = string(nombre[0:maximo])
			query := `UPDATE [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_CODE]
			SET LAST_CHANGE_USER_ID = '`+alias.(string)+`', LAST_CHANGE_DATETIME = '`+substring+`', STATUS = `+status+`, LOOKUP = '`+LOOKUP+`', NAME =  '`+nombre+`', PROHIBIT_POSTING = `+prohibido+`  
	WHERE ANL_CAT_ID = '`+ANL_CAT_ID+`' AND ANL_CODE = '`+codigo+`'`
			result , errx := conn.Exec(query)
			if errx != nil {
				defer conn.Close()
				fmt.Println(query)
				log.Fatal("Perdon updatex! : ", errx.Error())
			}
			afectados, _ := result.RowsAffected()
			if afectados == 1 {
				example := map[string]interface{}{ "success":1}
				c.Data["json"] = &example
				c.ServeJSON()
				return
			}
			example := map[string]interface{}{ "success":3}
			c.Data["json"] = &example
			c.ServeJSON()
			return
		}
		example := map[string]interface{}{ "success":2}
		c.Data["json"] = &example
		c.ServeJSON()
	}
} 
func (c *NuevaDimensionController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		tipo := c.GetString("tipo")
		codigo := c.GetString("codigo")
		status := c.GetString("status")
		nombre := c.GetString("nombre")
		prohibido := c.GetString("prohibido")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, errS := sql.Open("mssql", connString2)
		if errS != nil {
			log.Fatal("Perdon! Open connection failed:", errS.Error())
		}
		currentTime := int64(time.Now().Unix())
		tm := time.Unix(currentTime, 0)
		dateString := tm.String() 
		substring := string(dateString[0:10])	
		queryDireccion1 := "SELECT TOP 1 ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_CAT] WHERE S_HEAD = '"+tipo+"' order by ANL_CAT_ID desc"
		rowsDireccion1, _ := conn.Query(queryDireccion1)
		var ANL_CAT_ID string
		if rowsDireccion1.Next()  {
			rowsDireccion1.Scan(&ANL_CAT_ID)
		}
		queryDireccion := "SELECT TOP 1 ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_CODE] WHERE ANL_CODE = '"+codigo+"' AND ANL_CAT_ID = '"+ANL_CAT_ID+"' order by ANL_CAT_ID desc"
		rowsDireccion, _ := conn.Query(queryDireccion)
		if rowsDireccion.Next()  {
			example := map[string]interface{}{ "success":2}
			c.Data["json"] = &example
			c.ServeJSON()
			return
		}
		maximo := len(nombre)
		var LOOKUP string
		if maximo > 14 {
			maximo = 14
		}
		LOOKUP = string(nombre[0:maximo])
		query := `INSERT INTO [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_CODE]
(ANL_CAT_ID,ANL_CODE,UPDATE_COUNT,LAST_CHANGE_USER_ID,LAST_CHANGE_DATETIME,STATUS,LOOKUP,NAME,DAG_CODE,BDGT_CHECK,BDGT_STOP,PROHIBIT_POSTING,NAVIGATION_OPTION,COMBINED_BDGT_CHCK)
VALUES
('`+ANL_CAT_ID+`','`+codigo+`',1,'`+alias.(string)+`','`+substring+`',`+status+`,'`+LOOKUP+`', '`+nombre+`',NULL,0,0,`+prohibido+`,99,0)`
		result , errx := conn.Exec(query)
		if errx != nil {
			defer conn.Close()
			fmt.Println(query)
			log.Fatal("Perdon insertx! : ", errx.Error())
		}
		afectados, _ := result.RowsAffected()
		if afectados == 1 {
			example := map[string]interface{}{ "success":1}
			c.Data["json"] = &example
			c.ServeJSON()
			return
		}
		example := map[string]interface{}{ "success":3}
		c.Data["json"] = &example
		c.ServeJSON()
		return
	}
} 




func (c *NuevaCuentaController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		

		ACNT_CODE := c.GetString("ACNT_CODE")
		DESCR := c.GetString("DESCR")

		ACNT_TYPE := c.GetString("ACNT_TYPE")
		BAL_TYPE := c.GetString("BAL_TYPE")
		STATUS := c.GetString("STATUS")
		LINK_ACNT := c.GetString("LINK_ACNT")
		ENTER_ANL_1 := c.GetString("ENTER_ANL_1")
		ENTER_ANL_2 := c.GetString("ENTER_ANL_2")
		ENTER_ANL_3 := c.GetString("ENTER_ANL_3")
		ENTER_ANL_4 := c.GetString("ENTER_ANL_4")
		ENTER_ANL_5 := c.GetString("ENTER_ANL_5")
		ENTER_ANL_6 := c.GetString("ENTER_ANL_6")
		ENTER_ANL_7 := c.GetString("ENTER_ANL_7")
		ENTER_ANL_8 := c.GetString("ENTER_ANL_8")
		ENTER_ANL_9 := c.GetString("ENTER_ANL_9")
		ENTER_ANL_10 := c.GetString("ENTER_ANL_10")
 		D1 := c.GetString("D1")
 		D2 := c.GetString("D2")
 		D3 := c.GetString("D3")
 		D4 := c.GetString("D4")
 		D5 := c.GetString("D5")
 		D6 := c.GetString("D6")
 		D7 := c.GetString("D7")
 		D8 := c.GetString("D8")
 		D9 := c.GetString("D9")
 		D10 := c.GetString("D10")
		
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, errS := sql.Open("mssql", connString2)
		if errS != nil {
			log.Fatal("Perdon! Open connection failed:", errS.Error())
		}
		currentTime := int64(time.Now().Unix())
		tm := time.Unix(currentTime, 0)
		dateString := tm.String() 
		substring := string(dateString[0:10])	
		//ANL_CAT_ID := "01"
	
		queryDireccion0 := "SELECT TOP 1 ACNT_CODE FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ACNT] WHERE ACNT_CODE = '"+ACNT_CODE+"'"
		rowsDireccion0, _ := conn.Query(queryDireccion0)
		if rowsDireccion0.Next()  {
			example := map[string]interface{}{ "success":2}
			c.Data["json"] = &example
			c.ServeJSON()
			return
		}
		DESCR = strings.Replace(DESCR, "'", "", -1)
		maximoAux := 49
		lenDESCR := len(DESCR)
		if lenDESCR > maximoAux {
			DESCR = string(DESCR[0:maximoAux])
		}

		var LOOKUP string
		if lenDESCR > 14 {
			lenDESCR = 14
		}
		LOOKUP = string(DESCR[0:lenDESCR])

		
		query := `INSERT INTO [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ACNT]
  (ACNT_CODE,UPDATE_COUNT,LAST_CHANGE_USER_ID,LAST_CHANGE_DATETIME,DESCR,S_HEAD,LOOKUP,DAG_CODE,ACNT_TYPE,BAL_TYPE,STATUS,SUPPRESS_REVAL,LONG_DESCR,CONV_CODE_CTRL,DFLT_CURR_CODE,ALLOCN_IN_PROGRESS, LINK_ACNT, RPT_CONV_CTRL, USER_AREA, CR_LIMIT,ENTER_ANL_1,ENTER_ANL_2,ENTER_ANL_3,ENTER_ANL_4,ENTER_ANL_5,ENTER_ANL_6,ENTER_ANL_7,ENTER_ANL_8,ENTER_ANL_9,ENTER_ANL_10, OIL, CV4_DFLT_CURR_CODE,  CV4_CONV_CODE_CTRL,CV5_DFLT_CURR_CODE,CV5_CONV_CODE_CTRL,BANK_CURR_REQD,ACNT_LINKS_ALLOWED,PASP_ACNT_TYPE)
  VALUES ('`+ACNT_CODE+`',1,'`+alias.(string)+`','`+substring+`','`+DESCR+`','`+LOOKUP+`', '`+LOOKUP+`', NULL, `+ACNT_TYPE+`,  `+BAL_TYPE+`, `+STATUS+`, 0, '`+DESCR+`', 2, 'MXP1',  0,'`+LINK_ACNT+`',  0 , NULL, 0, `+ENTER_ANL_1+`, `+ENTER_ANL_2+`, `+ENTER_ANL_3+`, `+ENTER_ANL_4+`, `+ENTER_ANL_5+`, `+ENTER_ANL_6+`, `+ENTER_ANL_7+`, `+ENTER_ANL_8+`, `+ENTER_ANL_9+`, `+ENTER_ANL_10+`, NULL, 'MXP1', 2, NULL, 2, 0, 0, 0)`
		result , err := conn.Exec(query)
		if err != nil {
			defer conn.Close()
			log.Fatal("Perdon insert cuenta! : ", err.Error())
		}
		afectados, _ := result.RowsAffected()
		if afectados == 1 {
			var numeroDimension string
			//Dimension 1
			if Compare(D1,"")!=0{//diferente a cadena vacia
				queryDireccion1 := "SELECT TOP 1 ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_ENT_DEFN] WHERE ANL_ENT_ID = 152 AND ENTRY_NUM = 1"
				rowsDireccion1, _ := conn.Query(queryDireccion1)
				if rowsDireccion1.Next()  {
					rowsDireccion1.Scan(&numeroDimension)
					queryDireccion11 := "SELECT TOP 1 ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ACNT_ANL_CAT] WHERE ANL_CAT_ID = '"+numeroDimension+"' AND ACNT_CODE = '"+ACNT_CODE+"'"
					rowsDireccion11, _ := conn.Query(queryDireccion11)
					if rowsDireccion11.Next()  {
					}else{
						query1 := `INSERT INTO [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ACNT_ANL_CAT]
						 (ANL_CAT_ID,ACNT_CODE,UPDATE_COUNT,LAST_CHANGE_USER_ID,LAST_CHANGE_DATETIME,ANL_CODE)
						  VALUES ('`+numeroDimension+`','`+ACNT_CODE +`',1,'`+alias.(string)+`','`+substring+`','`+D1+`')`
						conn.Exec(query1)
					}
				}
			}
			//Dimension 2
			if Compare(D2,"")!=0{//diferente a cadena vacia
				queryDireccion2 := "SELECT TOP 1 ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_ENT_DEFN] WHERE ANL_ENT_ID = 152 AND ENTRY_NUM = 2"
				rowsDireccion2, _ := conn.Query(queryDireccion2)
				if rowsDireccion2.Next()  {
					rowsDireccion2.Scan(&numeroDimension)
					queryDireccion22 := "SELECT TOP 1 ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ACNT_ANL_CAT] WHERE ANL_CAT_ID = '"+numeroDimension+"' AND ACNT_CODE = '"+ACNT_CODE+"'"
					rowsDireccion22, _ := conn.Query(queryDireccion22)
					if rowsDireccion22.Next()  {
					}else{
						query2 := `INSERT INTO [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ACNT_ANL_CAT]
						 (ANL_CAT_ID,ACNT_CODE,UPDATE_COUNT,LAST_CHANGE_USER_ID,LAST_CHANGE_DATETIME,ANL_CODE)
						  VALUES ('`+numeroDimension+`','`+ACNT_CODE +`',1,'`+alias.(string)+`','`+substring+`','`+D2+`')`
						conn.Exec(query2)
					}
				}
			}
			//Dimension 3
			if Compare(D3,"")!=0{//diferente a cadena vacia
				queryDireccion3 := "SELECT TOP 1 ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_ENT_DEFN] WHERE ANL_ENT_ID = 152 AND ENTRY_NUM = 3"
				rowsDireccion3, _ := conn.Query(queryDireccion3)
				if rowsDireccion3.Next()  {
					rowsDireccion3.Scan(&numeroDimension)
					queryDireccion33 := "SELECT TOP 1 ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ACNT_ANL_CAT] WHERE ANL_CAT_ID = '"+numeroDimension+"' AND ACNT_CODE = '"+ACNT_CODE+"'"
					rowsDireccion33, _ := conn.Query(queryDireccion33)
					if rowsDireccion33.Next()  {
					}else{
						query3 := `INSERT INTO [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ACNT_ANL_CAT]
						 (ANL_CAT_ID,ACNT_CODE,UPDATE_COUNT,LAST_CHANGE_USER_ID,LAST_CHANGE_DATETIME,ANL_CODE)
						  VALUES ('`+numeroDimension+`','`+ACNT_CODE +`',1,'`+alias.(string)+`','`+substring+`','`+D3+`')`
						conn.Exec(query3)
					}
				}
			}
			//Dimension 4
			if Compare(D4,"")!=0{//diferente a cadena vacia
				queryDireccion4 := "SELECT TOP 1 ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_ENT_DEFN] WHERE ANL_ENT_ID = 152 AND ENTRY_NUM = 4"
				rowsDireccion4, _ := conn.Query(queryDireccion4)
				if rowsDireccion4.Next()  {
					rowsDireccion4.Scan(&numeroDimension)
					queryDireccion44 := "SELECT TOP 1 ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ACNT_ANL_CAT] WHERE ANL_CAT_ID = '"+numeroDimension+"' AND ACNT_CODE = '"+ACNT_CODE+"'"
					rowsDireccion44, _ := conn.Query(queryDireccion44)
					if rowsDireccion44.Next()  {
					}else{
						query4 := `INSERT INTO [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ACNT_ANL_CAT]
						 (ANL_CAT_ID,ACNT_CODE,UPDATE_COUNT,LAST_CHANGE_USER_ID,LAST_CHANGE_DATETIME,ANL_CODE)
						  VALUES ('`+numeroDimension+`','`+ACNT_CODE +`',1,'`+alias.(string)+`','`+substring+`','`+D4+`')`
						conn.Exec(query4)
					}
				}
			}
			//Dimension 5
			if Compare(D5,"")!=0{//diferente a cadena vacia
				queryDireccion5 := "SELECT TOP 1 ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_ENT_DEFN] WHERE ANL_ENT_ID = 152 AND ENTRY_NUM = 5"
				rowsDireccion5, _ := conn.Query(queryDireccion5)
				if rowsDireccion5.Next()  {
					rowsDireccion5.Scan(&numeroDimension)
					queryDireccion55 := "SELECT TOP 1 ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ACNT_ANL_CAT] WHERE ANL_CAT_ID = '"+numeroDimension+"' AND ACNT_CODE = '"+ACNT_CODE+"'"
					rowsDireccion55, _ := conn.Query(queryDireccion55)
					if rowsDireccion55.Next()  {
					}else{
						query5 := `INSERT INTO [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ACNT_ANL_CAT]
						 (ANL_CAT_ID,ACNT_CODE,UPDATE_COUNT,LAST_CHANGE_USER_ID,LAST_CHANGE_DATETIME,ANL_CODE)
						  VALUES ('`+numeroDimension+`','`+ACNT_CODE +`',1,'`+alias.(string)+`','`+substring+`','`+D5+`')`
						conn.Exec(query5)
					}
				}
			}
			//Dimension 6
			if Compare(D6,"")!=0{//diferente a cadena vacia
				queryDireccion6 := "SELECT TOP 1 ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_ENT_DEFN] WHERE ANL_ENT_ID = 152 AND ENTRY_NUM = 6"
				rowsDireccion6, _ := conn.Query(queryDireccion6)
				if rowsDireccion6.Next()  {
					rowsDireccion6.Scan(&numeroDimension)
					queryDireccion66 := "SELECT TOP 1 ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ACNT_ANL_CAT] WHERE ANL_CAT_ID = '"+numeroDimension+"' AND ACNT_CODE = '"+ACNT_CODE+"'"
					rowsDireccion66, _ := conn.Query(queryDireccion66)
					if rowsDireccion66.Next()  {
					}else{
						query6 := `INSERT INTO [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ACNT_ANL_CAT]
						 (ANL_CAT_ID,ACNT_CODE,UPDATE_COUNT,LAST_CHANGE_USER_ID,LAST_CHANGE_DATETIME,ANL_CODE)
						  VALUES ('`+numeroDimension+`','`+ACNT_CODE +`',1,'`+alias.(string)+`','`+substring+`','`+D6+`')`
						conn.Exec(query6)
					}
				}
			}
			//Dimension 7
			if Compare(D7,"")!=0{//diferente a cadena vacia
				queryDireccion7 := "SELECT TOP 1 ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_ENT_DEFN] WHERE ANL_ENT_ID = 152 AND ENTRY_NUM = 7"
				rowsDireccion7, _ := conn.Query(queryDireccion7)
				if rowsDireccion7.Next()  {
					rowsDireccion7.Scan(&numeroDimension)
					queryDireccion77 := "SELECT TOP 1 ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ACNT_ANL_CAT] WHERE ANL_CAT_ID = '"+numeroDimension+"' AND ACNT_CODE = '"+ACNT_CODE+"'"
					rowsDireccion77, _ := conn.Query(queryDireccion77)
					if rowsDireccion77.Next()  {
					}else{
						query7 := `INSERT INTO [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ACNT_ANL_CAT]
						 (ANL_CAT_ID,ACNT_CODE,UPDATE_COUNT,LAST_CHANGE_USER_ID,LAST_CHANGE_DATETIME,ANL_CODE)
						  VALUES ('`+numeroDimension+`','`+ACNT_CODE +`',1,'`+alias.(string)+`','`+substring+`','`+D7+`')`
						conn.Exec(query7)
					}
				}
			}
			//Dimension 8
			if Compare(D8,"")!=0{//diferente a cadena vacia
				queryDireccion8 := "SELECT TOP 1 ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_ENT_DEFN] WHERE ANL_ENT_ID = 152 AND ENTRY_NUM = 8"
				rowsDireccion8, _ := conn.Query(queryDireccion8)
				if rowsDireccion8.Next()  {
					rowsDireccion8.Scan(&numeroDimension)
					queryDireccion88 := "SELECT TOP 1 ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ACNT_ANL_CAT] WHERE ANL_CAT_ID = '"+numeroDimension+"' AND ACNT_CODE = '"+ACNT_CODE+"'"
					rowsDireccion88, _ := conn.Query(queryDireccion88)
					if rowsDireccion88.Next()  {
					}else{
						query8 := `INSERT INTO [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ACNT_ANL_CAT]
						 (ANL_CAT_ID,ACNT_CODE,UPDATE_COUNT,LAST_CHANGE_USER_ID,LAST_CHANGE_DATETIME,ANL_CODE)
						  VALUES ('`+numeroDimension+`','`+ACNT_CODE +`',1,'`+alias.(string)+`','`+substring+`','`+D8+`')`
						conn.Exec(query8)
					}
				}
			}
			//Dimension 9
			if Compare(D9,"")!=0{//diferente a cadena vacia
				queryDireccion9 := "SELECT TOP 1 ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_ENT_DEFN] WHERE ANL_ENT_ID = 152 AND ENTRY_NUM = 9"
				rowsDireccion9, _ := conn.Query(queryDireccion9)
				if rowsDireccion9.Next()  {
					rowsDireccion9.Scan(&numeroDimension)
					queryDireccion99 := "SELECT TOP 1 ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ACNT_ANL_CAT] WHERE ANL_CAT_ID = '"+numeroDimension+"' AND ACNT_CODE = '"+ACNT_CODE+"'"
					rowsDireccion99, _ := conn.Query(queryDireccion99)
					if rowsDireccion99.Next()  {
					}else{
						query9 := `INSERT INTO [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ACNT_ANL_CAT]
						 (ANL_CAT_ID,ACNT_CODE,UPDATE_COUNT,LAST_CHANGE_USER_ID,LAST_CHANGE_DATETIME,ANL_CODE)
						  VALUES ('`+numeroDimension+`','`+ACNT_CODE +`',1,'`+alias.(string)+`','`+substring+`','`+D9+`')`
						conn.Exec(query9)
					}
				}
			}
			//Dimension 10
			if Compare(D10,"")!=0{//diferente a cadena vacia
				queryDireccion10 := "SELECT TOP 1 ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_ENT_DEFN] WHERE ANL_ENT_ID = 152 AND ENTRY_NUM = 10"
				fmt.Println(queryDireccion10)
				rowsDireccion10, _ := conn.Query(queryDireccion10)
				if rowsDireccion10.Next()  {
					rowsDireccion10.Scan(&numeroDimension)
					queryDireccion110 := "SELECT TOP 1 ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ACNT_ANL_CAT] WHERE ANL_CAT_ID = '"+numeroDimension+"' AND ACNT_CODE = '"+ACNT_CODE+"'"
					rowsDireccion110, _ := conn.Query(queryDireccion110)
					if rowsDireccion110.Next()  {
					}else{
						query10 := `INSERT INTO [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ACNT_ANL_CAT]
						 (ANL_CAT_ID,ACNT_CODE,UPDATE_COUNT,LAST_CHANGE_USER_ID,LAST_CHANGE_DATETIME,ANL_CODE)
						  VALUES ('`+numeroDimension+`','`+ACNT_CODE +`',1,'`+alias.(string)+`','`+substring+`','`+D10+`')`
						  fmt.Println(query10)
						conn.Exec(query10)
					}
				}
			}
			example := map[string]interface{}{ "success":1}
			c.Data["json"] = &example
			c.ServeJSON()
			return
		}
		example := map[string]interface{}{ "success":3}
		c.Data["json"] = &example
		c.ServeJSON()
		return
	}
} 




func (c *GuardarConceptoCedulasController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		idCedula := c.GetString("idCedula")
		nombre := c.GetString("nombre")
		
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, errS := sql.Open("mssql", connString2)
		if errS != nil {
			log.Fatal("Perdon! Open connection failed:", errS.Error())
		}
		query := `INSERT INTO [SUNPLUSADV].[dbo].[CedulasConceptos] (idCedula,nombre) VALUES (`+idCedula+`, '`+nombre+`')`
		result , err := conn.Exec(query)
		if err != nil {
			defer conn.Close()
			log.Fatal("Perdon insert! : ", err.Error())
		}
		afectados, _ := result.RowsAffected()
		if afectados == 1 {
			example := map[string]interface{}{ "success":1}
			c.Data["json"] = &example
			c.ServeJSON()
			return
		}
		example := map[string]interface{}{ "success":0}
		c.Data["json"] = &example
		c.ServeJSON()
		return
	}
} 



func (c *ContabilizaDiarioController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		correo := c.GetSession("usuario")
		timestamp := c.GetString("timestamp")
		idTipoDeDiario := c.GetString("idTipoDeDiario")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, errS := sql.Open("mssql", connString2)
		if errS != nil {
			log.Fatal("Perdon! Open connection failed:", errS.Error())
		}

		var debitos float64
		var creditos float64
		
		queryDebitos := "SELECT ISNULL(SUM(AMOUNT),0) as amount FROM [SUNPLUSADV].[dbo].[DiariosRetenidos] WHERE idTipoDeDiario = "+idTipoDeDiario+" AND usuario = '"+alias.(string)+"' AND timestamp = "+timestamp+" AND D_C = 'D'"
		rowsDebitos, _ := conn.Query(queryDebitos)
		if rowsDebitos.Next()  {
			rowsDebitos.Scan(&debitos)
		}

		queryCreditos := "SELECT ISNULL(SUM(AMOUNT),0) as amount FROM [SUNPLUSADV].[dbo].[DiariosRetenidos] WHERE idTipoDeDiario = "+idTipoDeDiario+" AND usuario = '"+alias.(string)+"' AND timestamp = "+timestamp+" AND D_C = 'C'"
		rowsCreditos, _ := conn.Query(queryCreditos)
		if rowsCreditos.Next()  {
			rowsCreditos.Scan(&creditos)
		}
		if debitos!=creditos {
			example := map[string]interface{}{ "success":2, "debitos" : debitos, "creditos" : creditos}
			c.Data["json"] = &example
			c.ServeJSON()
			return;
		}
		var codigo string

		queryCodigo := "SELECT Codigo FROM [SUNPLUSADV].[dbo].[TiposDeDiario] WHERE idTipoDeDiario = "+idTipoDeDiario+" AND BUNIT = '"+BUNIT.(string)+"'"
		rowsCodigo, _ := conn.Query(queryCodigo)
		if rowsCodigo.Next()  {
			rowsCodigo.Scan(&codigo)
		}

		var diario int
		query1 := "SELECT ISNULL((MAX(JRNAL_NO)+1),1) as JRNAL_NO FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_A_SALFLDG]"
   		rows1, err1 := conn.Query(query1)
   		if err1 != nil {
   			defer conn.Close()
			log.Fatal("Perdon! Error en querys hardcodeados: ", err1.Error())
		}
   		if rows1.Next()  {
			rows1.Scan(&diario)
		}
		JRNAL_NO := strconv.Itoa(diario)

		queryAver := `SELECT idLinea,lineaG,lineaT,referencia,DESCRIPTN,ACNT_CODE,D_C,AMOUNT,ANAL_T0,ANAL_T1,ANAL_T2,ANAL_T3,ANAL_T4,ANAL_T5,ANAL_T6,ANAL_T7,ANAL_T8,ANAL_T9,fecha
  FROM [SUNPLUSADV].[dbo].[DiariosRetenidos]
  WHERE BUNIT = '`+BUNIT.(string)+`' AND usuario = '`+alias.(string)+`' AND timestamp = `+timestamp+`
  order by lineaG asc`
		rowsAver, errA := conn.Query(queryAver)
		if errA != nil {
			fmt.Println(queryAver)
			defer conn.Close()
			log.Fatal("Perdon select Diario! : ", errA.Error())
		}
		todos := 0
		cuantos := 0
		for rowsAver.Next()  {
			var idLinea int
			var lineaG int
			var lineaT int
			var referencia string
			var DESCRIPTN string
			var ACNT_CODE string
			var fecha string
			var D_C string
			var AMOUNT float64
			var ANAL_T0 string
			var ANAL_T1 string
			var ANAL_T2 string
			var ANAL_T3 string
			var ANAL_T4 string
			var ANAL_T5 string
			var ANAL_T6 string
			var ANAL_T7 string
			var ANAL_T8 string
			var ANAL_T9 string
			rowsAver.Scan(&idLinea,&lineaG, &lineaT, &referencia, &DESCRIPTN, &ACNT_CODE, &D_C, &AMOUNT, &ANAL_T0, &ANAL_T1, &ANAL_T2, &ANAL_T3, &ANAL_T4, &ANAL_T5, &ANAL_T6, &ANAL_T7, &ANAL_T8, &ANAL_T9, &fecha)
			amountString := fmt.Sprintf("%.2f", AMOUNT)
			mes:=string(fecha[0:2])
			dia:=string(fecha[3:5])
			anio:=string(fecha[6:10])
			fechaBuena:=anio+"-"+mes+"-"+dia
			PERIOD:=anio+"0"+mes
			lineaGG := strconv.Itoa(lineaG)

			queryFactura := "SELECT deboFacturar, servicio, cliente, concepto FROM [SUNPLUSADV].[dbo].[TiposDeDiarioLineas] WHERE idTipoDeDiario = "+idTipoDeDiario+" AND Linea = "+strconv.Itoa(lineaT)
			rowsFactura, errB := conn.Query(queryFactura)
				
			if errB != nil {
				defer conn.Close()
				log.Fatal("Perdon select Diario! : ", errB.Error())
			}

			if rowsFactura.Next()  {
				var deboFacturar int
				var servicio string
				var cliente string
				var concepto string
				rowsFactura.Scan(&deboFacturar, &servicio, &cliente, &concepto)

				if deboFacturar == 1 {
					Empresa := ""
					pass := ""

					queryOpciones := `SELECT Empresa, pass
					FROM [SUNPLUSADV].[dbo].[zConfig] WHERE BUNIT = '`+BUNIT.(string)+`'`
					rowsOpciones, errOpciones := conn.Query(queryOpciones)
					if errOpciones != nil {
						defer conn.Close()
						log.Fatal("Perdon opciones! : ", errOpciones.Error())
					}
					if rowsOpciones.Next()  {
						rowsOpciones.Scan(&Empresa, &pass)
						uDec, errorx := base64.StdEncoding.DecodeString(pass)
						if errorx != nil {
							panic(errorx)
						}
						uDecS := string(uDec) 
						resp, err := http.Get("http://umn.redirectme.net:755/factura?Empresa="+Empresa+"&cliente="+cliente+"&concepto="+concepto+"&servicio="+servicio+"&precio="+amountString+"&contra="+uDecS+"&correoReceptor="+correo.(string)+"&diario="+strconv.Itoa(diario)+"&linea="+lineaGG+"&BUNIT="+BUNIT.(string))
						if err != nil {
							panic(errorx)
						}
						defer resp.Body.Close()
						body, err := ioutil.ReadAll(resp.Body)
						fmt.Printf("%s", body)
					}
				}
				
			}




			queryInsert := `INSERT INTO [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
			 (TREFERENCE, ACCNT_CODE, ANAL_T0, ANAL_T1, ANAL_T2, ANAL_T3, ANAL_T4, ANAL_T5, ANAL_T6, ANAL_T7, ANAL_T8, ANAL_T9, DESCRIPTN,
			  AMOUNT, OTHER_AMT, CONV_CODE, D_C, PERIOD, TRANS_DATETIME, JRNAL_NO, JRNAL_LINE, JRNAL_TYPE, JRNAL_SRCE, 
			  ENTRY_DATETIME, ENTRY_PRD, DUE_DATETIME, CONV_RATE, POSTING_DATETIME, CV4_CONV_CODE, CV4_AMT, CV4_CONV_RATE, CV4_OPERATOR, ORIGINATOR_ID, POSTER_ID, JNL_REVERSAL_TYPE, SPLIT_IN_PROGRESS, AUTHORISTN_IN_PROGRESS, MAN_PAY_OVER, APRVLS_EXTSN, SUPPLMNTRY_EXTSN, TRUE_RATED, AGREED_STATUS, BINDER_STATUS, PRINCIPAL_REQD, SPLIT_ORIG_LINE, CV5_DP, CV5_OPERATOR, CV5_CONV_RATE, CV5_AMT, CV5_CONV_CODE, CV4_DP, CONSUMED_BDGT_ID, LE_DETAILS_IND, EXCLUDE_BAL, MEMO_AMT, REPORT_AMT, REPORT_OPERATOR, REPORT_RATE, CONV_OPERATOR, BASE_OPERATOR, BASE_RATE, HOLD_OP_ID, HOLD_REF, ALLOC_IN_PROGRESS, ALLOCATION, ALLOC_REF, ALLOC_PERIOD, ASSET_IND, ASSET_CODE, ASSET_SUB, OTHER_DP, CLEARDOWN, REVERSAL, LOSS_GAIN, ROUGH_FLAG, IN_USE_FLAG, ORIGINATED_DATETIME) VALUES 
('`+referencia+`','`+ACNT_CODE+`', '`+ANAL_T0+`', '`+ANAL_T1+`', '`+ANAL_T2+`', '`+ANAL_T3+`', '`+ANAL_T4+`', '`+ANAL_T5+`', '`+ANAL_T6+`', '`+ANAL_T7+`', '`+ANAL_T8+`', '`+ANAL_T9+`', '`+DESCRIPTN+`',
 `+amountString+`, `+amountString+`, 'MXP1', '`+D_C+`', `+PERIOD+`, '`+fechaBuena+`', `+JRNAL_NO+`, `+lineaGG+`, '`+codigo+`', '`+alias.(string)+`', 
 '`+fechaBuena+`', `+PERIOD+`, '`+fechaBuena+`', 1.0000000, '`+fechaBuena+`', 'MXP1', `+amountString+`, 1.0000000, '*', '`+alias.(string)+`', '`+alias.(string)+`', 0, 0, 0, 0 ,0 ,0, 0, 0, '', 0, 0, '2', '*', 0.00, 0.00, '', '2', 0, '', '', 0.0, 0.0, '*', 0.0, '*', '*', 1.0000, '', 0, '', '', 0, 0, '', '', '', '2', '00000', '', '', '', '', '`+fechaBuena+`')`
   			todos = todos + 1
   			result , errI := conn.Exec(queryInsert)
   			if errI != nil {
				fmt.Println(queryInsert)
				defer conn.Close()
				log.Fatal("Perdon insert o updateY! : ", errI.Error())
			}
			afectados, _ := result.RowsAffected()
			if afectados == 1 {
				cuantos = cuantos + 1
			}
		} 
		if cuantos == todos {
			if cuantos != 0{
				queryDELETE := `DELETE FROM [SUNPLUSADV].[dbo].[DiariosRetenidos]
  				WHERE BUNIT = '`+BUNIT.(string)+`' AND usuario = '`+alias.(string)+`' AND timestamp = `+timestamp
				conn.Exec(queryDELETE)


				example := map[string]interface{}{ "success":1, "diario" : JRNAL_NO}
				c.Data["json"] = &example
				c.ServeJSON()
			}
			return
		}
		example := map[string]interface{}{ "success":0}
		c.Data["json"] = &example
		c.ServeJSON()
		return
	}
} 


func (c *GuardarConfigOpcionesController) Post() {
  alias := c.GetSession("alias")
  if alias == nil{
    return
  }
  tipoDeUsuario := c.GetSession("tipoDeUsuario")
  if tienePermisosContador(tipoDeUsuario.(int)) {
    BUNIT := c.GetSession("BUNIT")
    pass := c.GetString("pass")
    Empresa := c.GetString("Empresa")
    tipoDimension := c.GetString("tipoDimension")
    cryptoText := base64.StdEncoding.EncodeToString(  []byte(pass))
  

    connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
    conn, errS := sql.Open("mssql", connString2)
    if errS != nil {
      log.Fatal("Perdon! Open connection failed:", errS.Error())
    }


    query := `INSERT INTO [SUNPLUSADV].[dbo].[zConfig] (BUNIT, pass, Empresa , tipoDimension) VALUES
     ( '`+BUNIT.(string)+`', '`+cryptoText+`', '`+Empresa+`',`+tipoDimension+`)`
    queryAver := "SELECT idConfig FROM [SUNPLUSADV].[dbo].[zConfig] WHERE BUNIT = '"+BUNIT.(string)+"'"
    rowsAver, _ := conn.Query(queryAver)
    if rowsAver.Next()  {
      var idConfig int
      rowsAver.Scan(&idConfig)
      idConfigS := strconv.Itoa(idConfig)
      query = `UPDATE [SUNPLUSADV].[dbo].[zConfig] SET Empresa = '`+Empresa+`', pass = '`+cryptoText+`', tipoDimension = `+tipoDimension+`
      WHERE idConfig =  `+ idConfigS
	} 
	result , err := conn.Exec(query)
	if err != nil {
		fmt.Println(query)
		defer conn.Close()
		log.Fatal("Perdon insert o updateX! : ", err.Error())
	}
	afectados, _ := result.RowsAffected()
	if afectados == 1 {
      example := map[string]interface{}{ "success":1}
      c.Data["json"] = &example
      c.ServeJSON()
      return
    }
    example := map[string]interface{}{ "success":0}
    c.Data["json"] = &example
    c.ServeJSON()
    return
  }
} 

func (c *NuevoLineasDeDiarioController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		timestamp := c.GetString("timestamp")
		lineaG := c.GetString("lineaG")
		lineaT := c.GetString("lineaT")
		ref := c.GetString("ref")
		fecha := c.GetString("fecha")
		DESCRIPTN := c.GetString("descripcion")
		AMOUNT := c.GetString("AMOUNT")
		ACNT_CODE := c.GetString("ACNT_CODE")
		D_C := c.GetString("D_C")
		idTipoDeDiario := c.GetString("idTipoDeDiario")
		ANAL_T0 := c.GetString("ANAL_T0")
		ANAL_T1 := c.GetString("ANAL_T1")
		ANAL_T2 := c.GetString("ANAL_T2")
		ANAL_T3 := c.GetString("ANAL_T3")
		ANAL_T4 := c.GetString("ANAL_T4")
		ANAL_T5 := c.GetString("ANAL_T5")
		ANAL_T6 := c.GetString("ANAL_T6")
		ANAL_T7 := c.GetString("ANAL_T7")
		ANAL_T8 := c.GetString("ANAL_T8")
		ANAL_T9 := c.GetString("ANAL_T9")
		
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, errS := sql.Open("mssql", connString2)
		if errS != nil {
			log.Fatal("Perdon! Open connection failed:", errS.Error())
		}


		query := `INSERT INTO [SUNPLUSADV].[dbo].[DiariosRetenidos] (idTipoDeDiario,BUNIT,lineaG, lineaT , timestamp , usuario, referencia, fecha, ACNT_CODE, DESCRIPTN, D_C, AMOUNT,  ANAL_T0, ANAL_T1, ANAL_T2, ANAL_T3, ANAL_T4, ANAL_T5, ANAL_T6, ANAL_T7, ANAL_T8, ANAL_T9 ) VALUES
		 ( `+idTipoDeDiario+`, '`+BUNIT.(string)+`',`+lineaG+`,`+lineaT+`,`+timestamp+`,'`+alias.(string)+`' ,'`+ref+`' ,'`+fecha+`','`+ACNT_CODE+`' ,'`+DESCRIPTN+`' ,'`+D_C+`' , `+AMOUNT+`, '`+ANAL_T0+`' ,'`+ANAL_T1+`' ,'`+ANAL_T2+`' ,'`+ANAL_T3+`' ,'`+ANAL_T4+`' ,'`+ANAL_T5+`' ,'`+ANAL_T6+`' ,'`+ANAL_T7+`' ,'`+ANAL_T8+`' ,'`+ANAL_T9+`' )`
		queryAver := "SELECT idLinea FROM [SUNPLUSADV].[dbo].[DiariosRetenidos] WHERE idTipoDeDiario = "+idTipoDeDiario+" AND usuario = '"+alias.(string)+"' AND timestamp = "+timestamp+" AND lineaG = "+lineaG+""
		rowsAver, _ := conn.Query(queryAver)
		if rowsAver.Next()  {
			var idLinea int
			rowsAver.Scan(&idLinea)
			idLineaS := strconv.Itoa(idLinea)
			query = `UPDATE [SUNPLUSADV].[dbo].[DiariosRetenidos] SET ACNT_CODE = '`+ACNT_CODE+`', referencia = '`+ref+`', fecha = '`+fecha+`', DESCRIPTN = '`+DESCRIPTN+`', AMOUNT = '`+AMOUNT+`', D_C = '`+D_C+`',
			ANAL_T0 = '`+ANAL_T0+`',
			ANAL_T1 = '`+ANAL_T1+`',
			ANAL_T2 = '`+ANAL_T2+`',
			ANAL_T3 = '`+ANAL_T3+`',
			ANAL_T4 = '`+ANAL_T4+`',
			ANAL_T5 = '`+ANAL_T5+`',
			ANAL_T6 = '`+ANAL_T6+`',
			ANAL_T7 = '`+ANAL_T7+`',
			ANAL_T8 = '`+ANAL_T8+`',
			ANAL_T9 = '`+ANAL_T9+`' WHERE idLinea =  `+ idLineaS
		} 
		result , err := conn.Exec(query)
		if err != nil {
			fmt.Println(query)
			defer conn.Close()
			log.Fatal("Perdon insert o updateX! : ", err.Error())
		}
		afectados, _ := result.RowsAffected()
		if afectados == 1 {
			var debitos float64
			var creditos float64
			
			queryDebitos := "SELECT ISNULL(SUM(AMOUNT),0) as amount FROM [SUNPLUSADV].[dbo].[DiariosRetenidos] WHERE idTipoDeDiario = "+idTipoDeDiario+" AND usuario = '"+alias.(string)+"' AND timestamp = "+timestamp+" AND D_C = 'D' AND referencia = '"+ref+"'"
			rowsDebitos, _ := conn.Query(queryDebitos)
			if rowsDebitos.Next()  {
				rowsDebitos.Scan(&debitos)
			}

			queryCreditos := "SELECT ISNULL(SUM(AMOUNT),0) as amount FROM [SUNPLUSADV].[dbo].[DiariosRetenidos] WHERE idTipoDeDiario = "+idTipoDeDiario+" AND usuario = '"+alias.(string)+"' AND timestamp = "+timestamp+" AND D_C = 'C' AND referencia = '"+ref+"'"
			rowsCreditos, _ := conn.Query(queryCreditos)
			if rowsCreditos.Next()  {
				rowsCreditos.Scan(&creditos)
			}

			example := map[string]interface{}{ "success":1, "debitos" : debitos, "creditos" : creditos}
			c.Data["json"] = &example
			c.ServeJSON()
			return
		}
		example := map[string]interface{}{ "success":0}
		c.Data["json"] = &example
		c.ServeJSON()
		return
	}
} 

func (c *NuevoLineasTiposDeDiarioController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		IdTipoDeDiario := c.GetString("IdTipoDeDiario")
		linea := c.GetString("linea")
		cliente := c.GetString("cliente")
		servicio := c.GetString("servicio")
		concepto := c.GetString("concepto")
		descripcion := c.GetString("descripcion")
		debofacturar := c.GetString("debofacturar")
		cuenta := c.GetString("cuenta")
		tipoDC := c.GetString("tipoDC")
		ANAL_T0 := c.GetString("ANAL_T0")
		ANAL_T1 := c.GetString("ANAL_T1")
		ANAL_T2 := c.GetString("ANAL_T2")
		ANAL_T3 := c.GetString("ANAL_T3")
		ANAL_T4 := c.GetString("ANAL_T4")
		ANAL_T5 := c.GetString("ANAL_T5")
		ANAL_T6 := c.GetString("ANAL_T6")
		ANAL_T7 := c.GetString("ANAL_T7")
		ANAL_T8 := c.GetString("ANAL_T8")
		ANAL_T9 := c.GetString("ANAL_T9")
		
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, errS := sql.Open("mssql", connString2)
		if errS != nil {
			log.Fatal("Perdon! Open connection failed:", errS.Error())
		}


		query := `INSERT INTO [SUNPLUSADV].[dbo].[TiposDeDiarioLineas] (idTipoDeDiario,ACNT_CODE, Linea , deboFacturar , servicio, cliente, concepto, DESCRIPTN, D_C,   ANAL_T0, ANAL_T1, ANAL_T2, ANAL_T3, ANAL_T4, ANAL_T5, ANAL_T6, ANAL_T7, ANAL_T8, ANAL_T9 ) VALUES
		 (`+IdTipoDeDiario+`,'`+cuenta+`',`+linea+`,`+debofacturar+`,'`+servicio+`' ,'`+cliente+`' ,'`+concepto+`','`+descripcion+`' ,'`+tipoDC+`' ,'`+ANAL_T0+`' ,'`+ANAL_T1+`' ,'`+ANAL_T2+`' ,'`+ANAL_T3+`' ,'`+ANAL_T4+`' ,'`+ANAL_T5+`' ,'`+ANAL_T6+`' ,'`+ANAL_T7+`' ,'`+ANAL_T8+`' ,'`+ANAL_T9+`' )`
		queryAver := "SELECT idLinea FROM [SUNPLUSADV].[dbo].[TiposDeDiarioLineas] WHERE idTipoDeDiario = "+IdTipoDeDiario+" AND Linea = "+linea+""
		rowsAver, _ := conn.Query(queryAver)
		if rowsAver.Next()  {
			var idLinea int
			rowsAver.Scan(&idLinea)
			idLineaS := strconv.Itoa(idLinea)
			query = `UPDATE [SUNPLUSADV].[dbo].[TiposDeDiarioLineas] SET ACNT_CODE = '`+cuenta+`', deboFacturar = `+debofacturar+`, servicio = '`+servicio+`', cliente = '`+cliente+`', concepto = '`+concepto+`', DESCRIPTN = '`+descripcion+`', D_C = '`+tipoDC+`',
			ANAL_T0 = '`+ANAL_T0+`',
			ANAL_T1 = '`+ANAL_T1+`',
			ANAL_T2 = '`+ANAL_T2+`',
			ANAL_T3 = '`+ANAL_T3+`',
			ANAL_T4 = '`+ANAL_T4+`',
			ANAL_T5 = '`+ANAL_T5+`',
			ANAL_T6 = '`+ANAL_T6+`',
			ANAL_T7 = '`+ANAL_T7+`',
			ANAL_T8 = '`+ANAL_T8+`',
			ANAL_T9 = '`+ANAL_T9+`' WHERE idLinea =  `+ idLineaS
		} 
		result , err := conn.Exec(query)
		if err != nil {
			defer conn.Close()
			log.Fatal("Perdon insert o update! : ", err.Error())
		}
		afectados, _ := result.RowsAffected()
		if afectados == 1 {
			example := map[string]interface{}{ "success":1}
			c.Data["json"] = &example
			c.ServeJSON()
			return
		}
		example := map[string]interface{}{ "success":0}
		c.Data["json"] = &example
		c.ServeJSON()
		return
	}
} 
func (c *NuevoLineasCedulasController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		IdConcepto := c.GetString("IdConcepto")
		cuenta := c.GetString("cuenta")
		tipoDC := c.GetString("tipoDC")
		ANAL_T0 := c.GetString("ANAL_T0")
		ANAL_T1 := c.GetString("ANAL_T1")
		ANAL_T2 := c.GetString("ANAL_T2")
		ANAL_T3 := c.GetString("ANAL_T3")
		ANAL_T4 := c.GetString("ANAL_T4")
		ANAL_T5 := c.GetString("ANAL_T5")
		ANAL_T6 := c.GetString("ANAL_T6")
		ANAL_T7 := c.GetString("ANAL_T7")
		ANAL_T8 := c.GetString("ANAL_T8")
		ANAL_T9 := c.GetString("ANAL_T9")
		
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, errS := sql.Open("mssql", connString2)
		if errS != nil {
			log.Fatal("Perdon! Open connection failed:", errS.Error())
		}
		query := `INSERT INTO [SUNPLUSADV].[dbo].[CedulasLinea] (idConcepto,cuenta, D_C_Tipo, ANAL_T0, ANAL_T1, ANAL_T2, ANAL_T3, ANAL_T4, ANAL_T5, ANAL_T6, ANAL_T7, ANAL_T8, ANAL_T9 ) VALUES
		 (`+IdConcepto+`,'`+cuenta+`',`+tipoDC+`,'`+ANAL_T0+`' ,'`+ANAL_T1+`' ,'`+ANAL_T2+`' ,'`+ANAL_T3+`' ,'`+ANAL_T4+`' ,'`+ANAL_T5+`' ,'`+ANAL_T6+`' ,'`+ANAL_T7+`' ,'`+ANAL_T8+`' ,'`+ANAL_T9+`' )`
		result , err := conn.Exec(query)
		if err != nil {
			defer conn.Close()
			log.Fatal("Perdon insert! : ", err.Error())
		}
		afectados, _ := result.RowsAffected()
		if afectados == 1 {
			example := map[string]interface{}{ "success":1}
			c.Data["json"] = &example
			c.ServeJSON()
			return
		}
		example := map[string]interface{}{ "success":0}
		c.Data["json"] = &example
		c.ServeJSON()
		return
	}
} 

func Dash(c *DashboardController )  {
	
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	hash := c.GetString("hash")
	currentTime := int64(time.Now().Unix())
	tm := time.Unix(currentTime, 0)
	dateString := tm.String() 
	substring := string(dateString[0:10])
   	byteArray := []byte(substring)
	hasher := sha512.New()
    hasher.Write(byteArray)
	cryptoText := base64.StdEncoding.EncodeToString(  []byte(hex.EncodeToString(hasher.Sum(nil))))
	esAPI := 0
	if Compare(cryptoText,hash)==0 {
		tipoDeUsuario = 3
		esAPI = 1
	}
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		if BUNIT == nil {
			BUNIT = c.GetString("BUNIT")
		}	
		PERIOD := c.GetString("PERIOD")
		anio:=string(PERIOD[0:4])	
		mes:=string(PERIOD[5:7])	
		PERIOD_ANTERIOR :=""
		mesNumero, _ := strconv.Atoi(mes)
		if mesNumero == 1 {
			mesS := "12"
			anioAnterior, _ := strconv.Atoi(anio)
			anioAnterior--
			PERIOD_ANTERIOR = strconv.Itoa(anioAnterior)+"0"+mesS
		} else {
			mesNumero--
			mesS := strconv.Itoa(mesNumero)
			if mesNumero < 10 {
				mesS = "0"+mesS
			}
			PERIOD_ANTERIOR = anio+"0"+mesS
		}



		//cuenta := c.GetString("cuenta")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, errS := sql.Open("mssql", connString2)
		if errS != nil {
			log.Fatal("Perdon! Open connection failed:", errS.Error())
		}
			
		//1 - 109, fondo 10, CAJAS Y BANCOS
		//11 - 119, fondo 10, INVERSIONES
		//139, fondo 10, CUENTAS X COBRAR
		//12, 14, 131,132,133,134,135,136,137,138, CUENTAS X COBRAR
		//3 - 349, fondo 10, CUENTAS X COBRAR
		//15- 159, fondo10, CUENTAS X COBRAR
		//16- 169, fondo10, IINVENTARIOS
		//17- 189, fondo10, Pagos anticipados
		//19, fondo 10, cuentas por cobrar entre fondos

		corrientes := 0.0
		var AMOUNT_ACTIVOS_CORRIENTES decimal.Dec
		

		query1 := `SELECT ISNULL(SUM(AMOUNT),0) as amount
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
		  WHERE SUBSTRING(ACCNT_CODE,1,3) in('100','101','102','103','104','105','106','107','108','109')
		  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
		rows1, err1 := conn.Query(query1)
		if err1 != nil {
			fmt.Println(query1)
			panic(err1)
		}
		
		if rows1.Next()  {
			rows1.Scan(&AMOUNT_ACTIVOS_CORRIENTES)
			amountPrima := math.Abs(AMOUNT_ACTIVOS_CORRIENTES.Float64())
			corrientes = corrientes + amountPrima
		}

		//11 - 119, fondo 10, INVERSIONES
		query2 := `SELECT ISNULL(SUM(AMOUNT),0) as amount
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
		  WHERE SUBSTRING(ACCNT_CODE,1,3) in('110','111','112','113','114','115','116','117','118','119')
		  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
		rows2, err2 := conn.Query(query2)
		if err2 != nil {
			fmt.Println(query2)
			panic(err2)
		}
		if rows2.Next()  {
			rows2.Scan(&AMOUNT_ACTIVOS_CORRIENTES)
			amountPrima := math.Abs(AMOUNT_ACTIVOS_CORRIENTES.Float64())
			corrientes = corrientes + amountPrima
		}

		//139, fondo 10, CUENTAS X COBRAR
		//12, 14, 131,132,133,134,135,136,137,138, CUENTAS X COBRAR
		//3 - 349, fondo 10, CUENTAS X COBRAR	
		
		query3_A := `SELECT ISNULL(SUM(AMOUNT),0) as amount
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
		  WHERE SUBSTRING(ACCNT_CODE,1,3) in('139')
		  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
		rows3_A, err3_A := conn.Query(query3_A)
		if err3_A != nil {
			fmt.Println(query3_A)
			panic(err3_A)
		}
		if rows3_A.Next()  {
			rows3_A.Scan(&AMOUNT_ACTIVOS_CORRIENTES)
			amountPrima := math.Abs(AMOUNT_ACTIVOS_CORRIENTES.Float64())
			corrientes = corrientes + amountPrima
		}
		var ACNT_CODE string

		query3_BB := `SELECT ACNT_CODE FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ACNT]
		 WHERE ISNUMERIC( SUBSTRING(ACNT_CODE,1,1))<>1 OR (SUBSTRING(ACNT_CODE,1,2) in('12','14','30','31','32','33','34') OR SUBSTRING(ACNT_CODE,1,3) in('130','131','132','133','134','135','136','137','138'))`
		rows3_BB, err3_BB := conn.Query(query3_BB)
		if err3_BB != nil {
			fmt.Println(query3_BB)
			panic(err3_BB)
		}
		for rows3_BB.Next()  {
			rows3_BB.Scan(&ACNT_CODE)
			query3_B := `SELECT ISNULL(SUM(AMOUNT),0) as amount
			  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
			  WHERE ACCNT_CODE = '`+ACNT_CODE+`'
			  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
			rows3_B, err3_B := conn.Query(query3_B)
			if err3_B != nil {
				fmt.Println(query3_B)
				panic(err3_B)
			}
			if rows3_B.Next()  {
				rows3_B.Scan(&AMOUNT_ACTIVOS_CORRIENTES)
				amountPrima1 := AMOUNT_ACTIVOS_CORRIENTES.Float64()
				if amountPrima1 < 0 {
					amountPrima := math.Abs(AMOUNT_ACTIVOS_CORRIENTES.Float64())
					corrientes = corrientes + amountPrima
				}
			}
		}

		//16- 169, fondo10, INVENTARIOS
		

		query4 := `SELECT ISNULL(SUM(AMOUNT),0) as amount
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
		  WHERE SUBSTRING(ACCNT_CODE,1,2) in('16')
		  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
		rows4, err4 := conn.Query(query4)
		if err4 != nil {
			fmt.Println(query4)
			panic(err4)
		}
		if rows4.Next()  {
			rows4.Scan(&AMOUNT_ACTIVOS_CORRIENTES)
			amountPrima := math.Abs(AMOUNT_ACTIVOS_CORRIENTES.Float64())
			corrientes = corrientes + amountPrima
		}
		//17- 189, fondo10, Pagos anticipados
		
		query5 := `SELECT ISNULL(SUM(AMOUNT),0) as amount
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
		  WHERE SUBSTRING(ACCNT_CODE,1,2) in('17','18')
		  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
		rows5, err5 := conn.Query(query5)
		if err5 != nil {
			fmt.Println(query5)
			panic(err5)
		}
		
		if rows5.Next()  {
			rows5.Scan(&AMOUNT_ACTIVOS_CORRIENTES)
			amountPrima := math.Abs(AMOUNT_ACTIVOS_CORRIENTES.Float64())
			corrientes = corrientes + amountPrima
		}

		//19, fondo 10, cuentas por cobrar entre fondos
		
		query6 := `SELECT ISNULL(SUM(AMOUNT),0) as amount
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
		  WHERE SUBSTRING(ACCNT_CODE,1,2) in('19')
		  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
		rows6, err6 := conn.Query(query6)
		if err6 != nil {
			fmt.Println(query6)
			panic(err6)
		}
		
		if rows6.Next()  {
			rows6.Scan(&AMOUNT_ACTIVOS_CORRIENTES)
			amountPrima := math.Abs(AMOUNT_ACTIVOS_CORRIENTES.Float64())
			corrientes = corrientes + amountPrima
		}

		//15- 159, fondo10, Documentos y prestamos x pagar
		

		query7 := `SELECT ISNULL(SUM(AMOUNT),0) as amount
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
		  WHERE SUBSTRING(ACCNT_CODE,1,2) in('15')
		  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
		rows7, err7 := conn.Query(query7)
		if err7 != nil {
			fmt.Println(query7)
			panic(err7)
		}
		
		if rows7.Next()  {
			rows7.Scan(&AMOUNT_ACTIVOS_CORRIENTES)
			amountPrima := math.Abs(AMOUNT_ACTIVOS_CORRIENTES.Float64())
			corrientes = corrientes + amountPrima
		}




		
		corrientesPasivos := 0.0
		var AMOUNT_CUENTAS_POR_PAGAR decimal.Dec
				
		query3_BB = `SELECT ACNT_CODE FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ACNT]
			 WHERE ISNUMERIC( SUBSTRING(ACNT_CODE,1,1))<>1 OR (SUBSTRING(ACNT_CODE,1,2) in('12','14','30','31','32','33','34') OR SUBSTRING(ACNT_CODE,1,3) in('130','131','132','133','134','135','136','137','138'))`
		rows3_BB, err3_BB = conn.Query(query3_BB)
		if err3_BB != nil {
			fmt.Println(query3_BB)
			panic(err3_BB)
		}
		for rows3_BB.Next()  {
			rows3_BB.Scan(&ACNT_CODE)
			query3_B := `SELECT ISNULL(SUM(AMOUNT),0) as amount
			  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
			  WHERE ACCNT_CODE = '`+ACNT_CODE+`'
			  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
			rows3_B, err3_B := conn.Query(query3_B)
			if err3_B != nil {
				fmt.Println(query3_B)
				panic(err3_B)
			}
			if rows3_B.Next()  {
				rows3_B.Scan(&AMOUNT_CUENTAS_POR_PAGAR)
				amountPrima1 := AMOUNT_CUENTAS_POR_PAGAR.Float64()
				if amountPrima1 > 0 {
					amountPrima := math.Abs(AMOUNT_CUENTAS_POR_PAGAR.Float64())
					corrientesPasivos=corrientesPasivos+amountPrima	
				}
			}
		}

				
		query7 = `SELECT ISNULL(SUM(AMOUNT),0) as amount
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
		  WHERE SUBSTRING(ACCNT_CODE,1,2) in('35')
		  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
		rows7, err7 = conn.Query(query7)
		if err7 != nil {
			fmt.Println(query7)
			panic(err7)
		}
		var AMOUNT_DOCUMENTOS decimal.Dec
		if rows7.Next()  {
			rows7.Scan(&AMOUNT_DOCUMENTOS)
			amountPrima := math.Abs(AMOUNT_DOCUMENTOS.Float64())
			corrientesPasivos=corrientesPasivos+amountPrima	
		}

		query1 = `SELECT ISNULL(SUM(AMOUNT),0) as amount
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
		  WHERE SUBSTRING(ACCNT_CODE,1,2) in('36')
		  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
		rows1, err1 = conn.Query(query1)
		if err1 != nil {
			fmt.Println(query1)
			panic(err1)
		}
		var AMOUNT_FONDOS_CONFIADOS decimal.Dec
		if rows1.Next()  {
			rows1.Scan(&AMOUNT_FONDOS_CONFIADOS)
			amountPrima := math.Abs(AMOUNT_FONDOS_CONFIADOS.Float64())
			corrientesPasivos=corrientesPasivos+amountPrima	
		}

		query2 = `SELECT ISNULL(SUM(AMOUNT),0) as amount
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
		  WHERE SUBSTRING(ACCNT_CODE,1,2) in('37','38')
		  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
		rows2, err2 = conn.Query(query2)
		if err2 != nil {
			fmt.Println(query2)
			panic(err2)
		}
		var AMOUNT_OTROS_PASIVOS_CORRIENTES decimal.Dec
		if rows2.Next()  {
			rows2.Scan(&AMOUNT_OTROS_PASIVOS_CORRIENTES)
			amountPrima := math.Abs(AMOUNT_OTROS_PASIVOS_CORRIENTES.Float64())
			corrientesPasivos=corrientesPasivos+amountPrima	
		}

			
		query4_BB := `SELECT ACNT_CODE FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ACNT]
			 WHERE SUBSTRING(ACNT_CODE,1,2) in('19')`
		rows4_BB, err4_BB := conn.Query(query4_BB)
		if err4_BB != nil {
			fmt.Println(query4_BB)
			panic(err4_BB)
		}
		for rows4_BB.Next()  {
			rows4_BB.Scan(&ACNT_CODE)
			query4_B := `SELECT ISNULL(SUM(AMOUNT),0) as amount
			  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
			  WHERE ACCNT_CODE = '`+ACNT_CODE+`'
			  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
			rows4_B, err4_B := conn.Query(query4_B)
			if err4_B != nil {
				fmt.Println(query4_B)
				panic(err4_B)
			}
			if rows4_B.Next()  {
				rows4_B.Scan(&AMOUNT_CUENTAS_POR_PAGAR)
				amountPrima1 := AMOUNT_CUENTAS_POR_PAGAR.Float64()
				if amountPrima1 > 0 {
					amountPrima := math.Abs(AMOUNT_CUENTAS_POR_PAGAR.Float64())
					corrientesPasivos=corrientesPasivos+amountPrima	
				}
			}
		}

				
				


		
		gastosOperativos := 0.0		
		esteAnio:=string(PERIOD[0:4])	
		PERIOD_INICIO := esteAnio+"001"

		anteriorAnio, _ := strconv.Atoi(esteAnio)
		anteriorAnio--
		anteriorAnioS := strconv.Itoa(anteriorAnio)

		//PERIOD_INICIO_ANTERIOR := anteriorAnioS+"001"
		PERIOD_ANO_ANTERIOR := anteriorAnioS+string(PERIOD[4:7])
			       
						
		//AND  SUBSTRING( CAST(PERIOD AS NVARCHAR(11)),1,4)  = '`+esteAnio+`'

		query7 = `SELECT ISNULL(SUM(AMOUNT),0) as amount
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
		  WHERE (SUBSTRING(ACCNT_CODE,1,2) in('80','81','82','85','86','87','88','89','90','91','92','93','94','95','98') OR
		  SUBSTRING(ACCNT_CODE,1,3) in('978') )
		  AND PERIOD >= `+PERIOD_INICIO+` AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10','20') 
		  AND ANAL_T3 >= 'AFO'`
		rows7, err7 = conn.Query(query7)
		if err7 != nil {
			fmt.Println(query7)
			panic(err7)
		}
		var AMOUNT_ESTE_ANO decimal.Dec
		if rows7.Next()  {
			rows7.Scan(&AMOUNT_ESTE_ANO)
			amountPrima := math.Abs(AMOUNT_ESTE_ANO.Float64())
			gastosOperativos = gastosOperativos + amountPrima
		}
						
		query1 = `SELECT ISNULL(SUM(AMOUNT),0) as amount
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
		  WHERE (SUBSTRING(ACCNT_CODE,1,2) in('80','81','82','85','86','87','88','89','90','91','92','93','94','95','98') OR
		  SUBSTRING(ACCNT_CODE,1,3) in('978') )
			 AND PERIOD >= `+PERIOD_ANO_ANTERIOR+` AND PERIOD < `+PERIOD_INICIO+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10','20') 
		   AND ANAL_T3 >= 'AFO'`
			
		
		rows1, err1 = conn.Query(query1)
		if err1 != nil {
			fmt.Println(query1)
			panic(err1)
		}
		var AMOUNT_FONDOS_CONFIADOS_2 decimal.Dec
		if rows1.Next()  {
			rows1.Scan(&AMOUNT_FONDOS_CONFIADOS_2)
			gastosOperativos = gastosOperativos + math.Abs(AMOUNT_FONDOS_CONFIADOS_2.Float64())
		}

		query8 := `SELECT ISNULL(SUM(AMOUNT),0) as amount
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
		  WHERE ACCNT_CODE = '878777'
		  AND PERIOD >= `+PERIOD_INICIO+` AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10','20') 
		  AND ANAL_T3 >= 'AFO'`
		rows8, err8 := conn.Query(query8)
		if err8 != nil {
			fmt.Println(query8)
			panic(err8)
		}
		var AMOUNT_DIEZMO_UNION decimal.Dec
		if rows8.Next()  {
			rows8.Scan(&AMOUNT_DIEZMO_UNION)
			amountPrima := math.Abs(AMOUNT_DIEZMO_UNION.Float64())
			gastosOperativos = gastosOperativos - amountPrima
		}

		query88 := `SELECT ISNULL(SUM(AMOUNT),0) as amount
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
		  WHERE ACCNT_CODE = '878777'
		  AND PERIOD >= `+PERIOD_ANO_ANTERIOR+` AND PERIOD < `+PERIOD_INICIO+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10','20') 
		  AND ANAL_T3 >= 'AFO'`
		rows88, err88 := conn.Query(query88)
		if err88 != nil {
			fmt.Println(query88)
			panic(err88)
		}
		if rows88.Next()  {
			rows88.Scan(&AMOUNT_DIEZMO_UNION)
			amountPrima := math.Abs(AMOUNT_DIEZMO_UNION.Float64())
			gastosOperativos = gastosOperativos - amountPrima
		}

		query9 := `SELECT ISNULL(SUM(AMOUNT),0) as amount
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
		  WHERE ACCNT_CODE = '878555'
		  AND PERIOD >= `+PERIOD_INICIO+` AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10','20') 
		  AND ANAL_T3 >= 'AFO'`
		rows9, err9 := conn.Query(query9)
		if err9 != nil {
			fmt.Println(query9)
			panic(err9)
		}
		var AMOUNT_DIEZMO_TAM decimal.Dec
		if rows9.Next()  {
			rows9.Scan(&AMOUNT_DIEZMO_TAM)
			amountPrima := math.Abs(AMOUNT_DIEZMO_TAM.Float64())
			gastosOperativos = gastosOperativos - amountPrima
		}

		query99 := `SELECT ISNULL(SUM(AMOUNT),0) as amount
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
		  WHERE ACCNT_CODE = '878555'
		  AND PERIOD >= `+PERIOD_ANO_ANTERIOR+` AND PERIOD < `+PERIOD_INICIO+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10','20') 
		  AND ANAL_T3 >= 'AFO'`
		rows99, err99 := conn.Query(query99)
		if err99 != nil {
			fmt.Println(query99)
			panic(err99)
		}
		if rows99.Next()  {
			rows99.Scan(&AMOUNT_DIEZMO_TAM)
			amountPrima := math.Abs(AMOUNT_DIEZMO_TAM.Float64())
			gastosOperativos = gastosOperativos - amountPrima
		}

		
		queryA := `SELECT ISNULL(SUM(AMOUNT),0) as amount
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
		  WHERE (SUBSTRING(ACCNT_CODE,1,2) in('80','81','82','83','84','85','86','87','88','89','90','91','92','93','94','95','98') OR
		  SUBSTRING(ACCNT_CODE,1,3) in('978') )
			 AND PERIOD >= `+PERIOD_INICIO+` AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10','20') 
		   AND ANAL_T3 = 'AFOMISIO01'`
			
		rowsA, errA := conn.Query(queryA)
		if errA != nil {
			fmt.Println(queryA)
			panic(errA)
		}
		var AMOUNT_AFOMISIO01 decimal.Dec
		if rowsA.Next()  {
			rowsA.Scan(&AMOUNT_AFOMISIO01)
			amountPrima := math.Abs(AMOUNT_AFOMISIO01.Float64())
			gastosOperativos = gastosOperativos - amountPrima
		}

		queryAA := `SELECT ISNULL(SUM(AMOUNT),0) as amount
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
		  WHERE (SUBSTRING(ACCNT_CODE,1,2) in('80','81','82','83','84','85','86','87','88','89','90','91','92','93','94','95','98') OR
		  SUBSTRING(ACCNT_CODE,1,3) in('978') )
			 AND PERIOD >= `+PERIOD_ANO_ANTERIOR+` AND PERIOD < `+PERIOD_INICIO+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10','20') 
		   AND ANAL_T3 = 'AFOMISIO01'`
			
		rowsAA, errAA := conn.Query(queryAA)
		if errAA != nil {
			fmt.Println(queryAA)
			panic(errAA)
		}
		if rowsAA.Next()  {
			rowsAA.Scan(&AMOUNT_AFOMISIO01)
			amountPrima := math.Abs(AMOUNT_AFOMISIO01.Float64())
			gastosOperativos = gastosOperativos - amountPrima
		}

		queryB := `SELECT ISNULL(SUM(AMOUNT),0) as amount
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
		  WHERE (SUBSTRING(ACCNT_CODE,1,2) in('80','81','82','83','84','85','86','87','88','89','90','91','92','93','94','95','98') OR
		  SUBSTRING(ACCNT_CODE,1,3) in('978') )
			 AND PERIOD >= `+PERIOD_INICIO+` AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10','20') 
		   AND ANAL_T3 = 'AFOAG1111G'`
			
		rowsB, errB := conn.Query(queryB)
		if errB != nil {
			fmt.Println(queryB)
			panic(errB)
		}
		var AMOUNT_AFOAG1111G decimal.Dec
		if rowsB.Next()  {
			rowsB.Scan(&AMOUNT_AFOAG1111G)
			amountPrima := math.Abs(AMOUNT_AFOAG1111G.Float64())
			gastosOperativos = gastosOperativos - amountPrima
		}

		queryBB := `SELECT ISNULL(SUM(AMOUNT),0) as amount
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
		  WHERE (SUBSTRING(ACCNT_CODE,1,2) in('80','81','82','83','84','85','86','87','88','89','90','91','92','93','94','95','98') OR
		  SUBSTRING(ACCNT_CODE,1,3) in('978') )
			 AND PERIOD >= `+PERIOD_ANO_ANTERIOR+` AND PERIOD < `+PERIOD_INICIO+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10','20') 
		   AND ANAL_T3 = 'AFOAG1111G'`
			
		rowsBB, errBB := conn.Query(queryBB)
		if errBB != nil {
			fmt.Println(queryBB)
			panic(errBB)
		}
		if rowsBB.Next()  {
			rowsBB.Scan(&AMOUNT_AFOAG1111G)
			amountPrima := math.Abs(AMOUNT_AFOAG1111G.Float64())
			gastosOperativos = gastosOperativos - amountPrima
		}

		queryC := `SELECT ISNULL(SUM(AMOUNT),0) as amount
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
		  WHERE (SUBSTRING(ACCNT_CODE,1,2) in('80','81','82','83','84','85','86','87','88','89','90','91','92','93','94','95','98') OR
		  SUBSTRING(ACCNT_CODE,1,3) in('978') )
			 AND PERIOD >= `+PERIOD_INICIO+` AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10','20') 
		   AND ANAL_T3 = 'AFOAG1111D'`
			
		rowsC, errC := conn.Query(queryC)
		if errC != nil {
			fmt.Println(queryC)
			panic(errC)
		}
		var AMOUNT_AFOAG1111GD decimal.Dec
		if rowsC.Next()  {
			rowsC.Scan(&AMOUNT_AFOAG1111GD)
			amountPrima := math.Abs(AMOUNT_AFOAG1111GD.Float64())
			gastosOperativos = gastosOperativos - amountPrima
		}

		queryCC := `SELECT ISNULL(SUM(AMOUNT),0) as amount
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
		  WHERE (SUBSTRING(ACCNT_CODE,1,2) in('80','81','82','83','84','85','86','87','88','89','90','91','92','93','94','95','98') OR
		  SUBSTRING(ACCNT_CODE,1,3) in('978') )
			 AND PERIOD >= `+PERIOD_ANO_ANTERIOR+` AND PERIOD < `+PERIOD_INICIO+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10','20') 
		   AND ANAL_T3 = 'AFOAG1111D'`
			
		rowsCC, errCC := conn.Query(queryCC)
		if errCC != nil {
			fmt.Println(queryCC)
			panic(errCC)
		}
		if rowsCC.Next()  {
			rowsCC.Scan(&AMOUNT_AFOAG1111GD)
			amountPrima := math.Abs(AMOUNT_AFOAG1111GD.Float64())
			gastosOperativos = gastosOperativos - amountPrima
		}
		pcent := 0.2
		if Compare(beego.AppConfig.String("TipoCampo"),"2")==0 {
			pcent = 0.3
		}
		if Compare(beego.AppConfig.String("TipoCampo"),"1")==0 {
			pcent = 0.2
		}
		gastosOperativos=gastosOperativos*pcent

		activosNetosAsignados := 0.0

		queryC = `SELECT ISNULL(SUM(AMOUNT),0) as amount
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
		  WHERE SUBSTRING(ACCNT_CODE,1,1) in('5','6','7','8','9')
			 AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10') 
		   AND SUBSTRING(ANAL_T3,1,2)  in ('AF')`
			
		rowsC, errC = conn.Query(queryC)
		if errC != nil {
			fmt.Println(queryC)
			panic(errC)
		}
		var AMOUNT_FONDOS_ASIGNADOS decimal.Dec
		if rowsC.Next()  {
			rowsC.Scan(&AMOUNT_FONDOS_ASIGNADOS)
			activosNetosAsignados = math.Abs(AMOUNT_FONDOS_ASIGNADOS.Float64())
		}

		corrientes = math.Abs(corrientes)
		activosNetosAsignados = math.Abs(activosNetosAsignados)

		//////
		///    AQUI EMPIEZA LA LIQUIDEZ!
		/////
		efectivo := 0.0
		queryC = `SELECT ISNULL(SUM(AMOUNT),0) as amount
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
		  WHERE SUBSTRING(ACCNT_CODE,1,2) in('10')
			 AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10') 
		   `
			
		rowsC, errC = conn.Query(queryC)
		if errC != nil {
			fmt.Println(queryC)
			panic(errC)
		}
		var AMOUNT_EFECTIVO decimal.Dec
		if rowsC.Next()  {
			rowsC.Scan(&AMOUNT_EFECTIVO)
			efectivo = math.Abs(AMOUNT_EFECTIVO.Float64())
		}

		inversiones := 0.0
		queryC = `SELECT ISNULL(SUM(AMOUNT),0) as amount
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
		  WHERE SUBSTRING(ACCNT_CODE,1,2) in('11')
			 AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10') 
		   `
			
		rowsC, errC = conn.Query(queryC)
		if errC != nil {
			fmt.Println(queryC)
			panic(errC)
		}
		if rowsC.Next()  {
			rowsC.Scan(&AMOUNT_EFECTIVO)
			inversiones = math.Abs(AMOUNT_EFECTIVO.Float64())
		}

		cuentasPorCobrarOrganizacionesSuperiores := 0.0

		query3_BB = `SELECT ACNT_CODE FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ACNT]
			 WHERE ACNT_CODE = 'AGL111'`
		rows3_BB, err3_BB = conn.Query(query3_BB)
		if err3_BB != nil {
			fmt.Println(query3_BB)
			panic(err3_BB)
		}
		var AMOUNT_CUENTAS_POR_COBRAR_ORG_SUPERIOR decimal.Dec
		for rows3_BB.Next()  {
			rows3_BB.Scan(&ACNT_CODE)
			query3_B := `SELECT ISNULL(SUM(AMOUNT),0) as amount
			  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
			  WHERE ACCNT_CODE = '`+ACNT_CODE+`'
			  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
			rows3_B, err3_B := conn.Query(query3_B)
			if err3_B != nil {
				fmt.Println(query3_B)
				panic(err3_B)
			}
			if rows3_B.Next()  {
				rows3_B.Scan(&AMOUNT_CUENTAS_POR_COBRAR_ORG_SUPERIOR)
				amountPrima1 := AMOUNT_CUENTAS_POR_COBRAR_ORG_SUPERIOR.Float64()
				if amountPrima1 < 0 {
					cuentasPorCobrarOrganizacionesSuperiores = math.Abs(AMOUNT_CUENTAS_POR_COBRAR_ORG_SUPERIOR.Float64())
				}
			}
		}






		queryPeriodos := "SELECT DISTINCT PERIOD FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_A_SALFLDG] order by PERIOD asc"
		rowsPeriodos, _ := conn.Query(queryPeriodos)
		var periodo models.Periodo
        _ = periodo
        for rowsPeriodos.Next()  {
			rowsPeriodos.Scan(&PERIOD)
			periodo := models.Periodo{PERIOD}
			models.AddPeriodo(periodo, PERIOD)
		}


		if esAPI == 0 {
			example := map[string]interface{}{ "success":1, "pcent" : pcent, "cuentaCobrarOrgSuperior" : cuentasPorCobrarOrganizacionesSuperiores, "inversiones" : inversiones, "efectivo" : efectivo, "periodos" : models.GetAllPeriodos(), "periodoAnterior":PERIOD_ANTERIOR, "activosNetosAsignados" : activosNetosAsignados, "corrientesActivos" : corrientes, "corrientesPasivos" : corrientesPasivos, "gastosOperativos" : gastosOperativos}
			c.Data["json"] = &example
			c.ServeJSON()
		} else {
			example := map[string]interface{}{ "success":1, "idCampo" : c.GetString("idCampo"), "nombre" : c.GetString("nombre"), "pcent" : pcent, "cuentaCobrarOrgSuperior" : cuentasPorCobrarOrganizacionesSuperiores, "inversiones" : inversiones, "efectivo" : efectivo, "periodos" : models.GetAllPeriodos(), "periodoAnterior":PERIOD_ANTERIOR, "activosNetosAsignados" : activosNetosAsignados, "corrientesActivos" : corrientes, "corrientesPasivos" : corrientesPasivos, "gastosOperativos" : gastosOperativos}
			c.Data["jsonp"] = &example
			c.ServeJSONP()
		}
		
		return
	}
}

func (c *DashboardController) Post() {
	Dash(c)
} 

func (c *DashboardController) Get() {
	Dash(c)
} 

func (c *DameDimensionesDisponiblesSegunLaCuentaController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		cuenta := c.GetString("cuenta")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, errS := sql.Open("mssql", connString2)
		if errS != nil {
			log.Fatal("Perdon! Open connection failed:", errS.Error())
		}
		
		tipoDimension := 0

		queryOpciones := `SELECT tipoDimension
		FROM [SUNPLUSADV].[dbo].[zConfig] WHERE BUNIT = '`+BUNIT.(string)+`'`
		rowsOpciones, errOpciones := conn.Query(queryOpciones)
		if errOpciones != nil {
			defer conn.Close()
			log.Fatal("Perdon opciones! : ", errOpciones.Error())
		}
		if rowsOpciones.Next()  {
			rowsOpciones.Scan(&tipoDimension)
		}

		query := `SELECT STATUS, ENTER_ANL_1, ENTER_ANL_2, ENTER_ANL_3, ENTER_ANL_4, ENTER_ANL_5, ENTER_ANL_6, ENTER_ANL_7, ENTER_ANL_8, ENTER_ANL_9, ENTER_ANL_10 
		FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ACNT] WHERE ACNT_CODE = '`+cuenta+`'`
		rowsAver, err := conn.Query(query)
		if err != nil {
			defer conn.Close()
			log.Fatal("Perdon insert! : ", err.Error())
		}
		if rowsAver.Next()  {

			var STATUS int
			var ENTER_ANL_1 int
			var ENTER_ANL_2 int
			var ENTER_ANL_3 int
			var ENTER_ANL_4 int
			var ENTER_ANL_5 int
			var ENTER_ANL_6 int
			var ENTER_ANL_7 int
			var ENTER_ANL_8 int
			var ENTER_ANL_9 int
			var ENTER_ANL_10 int

			rowsAver.Scan(&STATUS, &ENTER_ANL_1, &ENTER_ANL_2, &ENTER_ANL_3, &ENTER_ANL_4, &ENTER_ANL_5, &ENTER_ANL_6, &ENTER_ANL_7, &ENTER_ANL_8, &ENTER_ANL_9, &ENTER_ANL_10)
			

			example := map[string]interface{}{ "success":1, "tipoDimension" : tipoDimension, "status" : STATUS, "D0" : ENTER_ANL_1, "D1" : ENTER_ANL_2, "D2" : ENTER_ANL_3, "D3" : ENTER_ANL_4, "D4" : ENTER_ANL_5, "D5" : ENTER_ANL_6, "D6" : ENTER_ANL_7, "D7" : ENTER_ANL_8, "D8" : ENTER_ANL_9, "D9" : ENTER_ANL_10 }
			c.Data["json"] = &example
			c.ServeJSON()
			return
		}
		example := map[string]interface{}{ "success":0}
		c.Data["json"] = &example
		c.ServeJSON()
		return
	}
} 
func (c *DameLineaDelTipoDeDiarioController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		idTipoDeDiario := c.GetString("idTipoDeDiario")
		linea := c.GetString("linea")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, errS := sql.Open("mssql", connString2)
		if errS != nil {
			log.Fatal("Perdon! Open connection failed:", errS.Error())
		}
		query := `SELECT ACNT_CODE, D_C, DESCRIPTN, ANAL_T0, ANAL_T1, ANAL_T2, ANAL_T3, ANAL_T4, ANAL_T5, ANAL_T6, ANAL_T7, ANAL_T8, ANAL_T9 FROM [SUNPLUSADV].[dbo].[TiposDeDiarioLineas] WHERE Linea = `+linea+` AND idTipoDeDiario = `+idTipoDeDiario
		rowsAver, err := conn.Query(query)
		if err != nil {
			defer conn.Close()
			log.Fatal("Perdon insert! : ", err.Error())
		}
		if rowsAver.Next()  {
			var cuenta string
			var D_C string
			var DESCRIPTN string
			var ANAL_T0 string
			var ANAL_T1 string
			var ANAL_T2 string
			var ANAL_T3 string
			var ANAL_T4 string
			var ANAL_T5 string
			var ANAL_T6 string
			var ANAL_T7 string
			var ANAL_T8 string
			var ANAL_T9 string
			rowsAver.Scan(&cuenta, &D_C, &DESCRIPTN, &ANAL_T0, &ANAL_T1, &ANAL_T2, &ANAL_T3, &ANAL_T4, &ANAL_T5, &ANAL_T6, &ANAL_T7, &ANAL_T8, &ANAL_T9)
			example := map[string]interface{}{ "success":1, "cuenta" : cuenta, "D_C" : D_C, "DESCRIPTN" : DESCRIPTN, "D0" : ANAL_T0, "D1" : ANAL_T1, "D2" : ANAL_T2, "D3" : ANAL_T3, "D4" : ANAL_T4, "D5" : ANAL_T5, "D6" : ANAL_T6, "D7" : ANAL_T7, "D8" : ANAL_T8, "D9" : ANAL_T9 }
			c.Data["json"] = &example
			c.ServeJSON()
			return
		}
		example := map[string]interface{}{ "success":0}
		c.Data["json"] = &example
		c.ServeJSON()
		return
	}
} 
func (c *NuevoTipoDeDiarioController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		nombre := c.GetString("nombre")
		codigo := c.GetString("codigo")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, errS := sql.Open("mssql", connString2)
		if errS != nil {
			log.Fatal("Perdon! Open connection failed:", errS.Error())
		}
		query := `INSERT INTO [SUNPLUSADV].[dbo].[TiposDeDiario] (Codigo,nombre,BUNIT) VALUES ('`+codigo+`','`+nombre+`','`+BUNIT.(string)+`')`
		result , err := conn.Exec(query)
		if err != nil {
			defer conn.Close()
			log.Fatal("Perdon insert! : ", err.Error())
		}
		afectados, _ := result.RowsAffected()
		if afectados == 1 {
			example := map[string]interface{}{ "success":1}
			c.Data["json"] = &example
			c.ServeJSON()
			return
		}
		example := map[string]interface{}{ "success":0}
		c.Data["json"] = &example
		c.ServeJSON()
		return
	}
} 

func (c *NuevaCedulasController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		nombre := c.GetString("nombre")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, errS := sql.Open("mssql", connString2)
		if errS != nil {
			log.Fatal("Perdon! Open connection failed:", errS.Error())
		}
		query := `INSERT INTO [SUNPLUSADV].[dbo].[Cedulas](nombre,BUNIT) VALUES ('`+nombre+`','`+BUNIT.(string)+`')`
		result , err := conn.Exec(query)
		if err != nil {
			defer conn.Close()
			log.Fatal("Perdon insert! : ", err.Error())
		}
		afectados, _ := result.RowsAffected()
		if afectados == 1 {
			example := map[string]interface{}{ "success":1}
			c.Data["json"] = &example
			c.ServeJSON()
			return
		}
		example := map[string]interface{}{ "success":0}
		c.Data["json"] = &example
		c.ServeJSON()
		return
	}
} 


func (c *NuevoTipoDeDimensionesController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		S_HEAD := c.GetString("S_HEAD")
		DESCR := c.GetString("DESCR")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, errS := sql.Open("mssql", connString2)
		if errS != nil {
			log.Fatal("Perdon! Open connection failed:", errS.Error())
		}
		currentTime := int64(time.Now().Unix())
		tm := time.Unix(currentTime, 0)
		dateString := tm.String() 
		substring := string(dateString[0:10])	
		ANL_CAT_ID := "01"
		queryDireccion := "SELECT TOP 1 ANL_CAT_ID FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_CAT] order by ANL_CAT_ID desc"
		rowsDireccion, err := conn.Query(queryDireccion)
		var previoANAL string
		if rowsDireccion.Next()  {
			rowsDireccion.Scan(&previoANAL)
			numeroAux, _ := strconv.Atoi(previoANAL);
			numeroAux = numeroAux + 1
			if numeroAux > 9 {
				ANL_CAT_ID = strconv.Itoa(numeroAux);
			} else {
				ANL_CAT_ID = "0"+ strconv.Itoa(numeroAux);
			}
		}
		query := `INSERT INTO [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ANL_CAT](ANL_CAT_ID,UPDATE_COUNT,LAST_CHANGE_USER_ID,LAST_CHANGE_DATETIME,STATUS,LOOKUP,USEABLE_ANL_ENT_ID,S_HEAD,DESCR,DAG_CODE,AMEND_CODE,VALIDATE_IND,LNGTH,LINKED,IBUS_CODE_DIM_ID)
		VALUES ('`+ANL_CAT_ID+`',1,'`+alias.(string)+`','`+substring+`',0,'`+S_HEAD+`',NULL,'`+S_HEAD+`','`+DESCR+`',NULL,1,1,15,0,NULL)`
		result , err := conn.Exec(query)
		if err != nil {
			defer conn.Close()
			log.Fatal("Perdon insert! : ", err.Error())
		}
		afectados, _ := result.RowsAffected()
		if afectados == 1 {
			example := map[string]interface{}{ "success":1}
			c.Data["json"] = &example
			c.ServeJSON()
			return
		}
		example := map[string]interface{}{ "success":2}
		c.Data["json"] = &example
		c.ServeJSON()
		return
	}
} 
func (c *EditarTipoDeDimensionesController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		ANL_CAT_ID := c.GetString("ANL_CAT_ID")
		S_HEAD := c.GetString("S_HEAD")
		DESCR := c.GetString("DESCR")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, errS := sql.Open("mssql", connString2)
		if errS != nil {
			log.Fatal("Perdon! Open connection failed:", errS.Error())
		}
		currentTime := int64(time.Now().Unix())
		tm := time.Unix(currentTime, 0)
		dateString := tm.String() 
		substring := string(dateString[0:10])	   
		query := "UPDATE [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_CAT] SET LOOKUP = '"+S_HEAD+"', S_HEAD = '"+S_HEAD+"', DESCR = '"+DESCR+"', LAST_CHANGE_DATETIME = '"+substring+"', LAST_CHANGE_USER_ID = '"+alias.(string)+"' WHERE ANL_CAT_ID = '"+ANL_CAT_ID+"'"
		result , err := conn.Exec(query)
		if err != nil {
			defer conn.Close()
			log.Fatal("Perdon insert! : ", err.Error())
		}
		afectados, _ := result.RowsAffected()
		if afectados == 1 {
			example := map[string]interface{}{ "success":1}
			c.Data["json"] = &example
			c.ServeJSON()
			return
		}
		example := map[string]interface{}{ "success":2}
		c.Data["json"] = &example
		c.ServeJSON()
		return
	}
} 
func (c *GenerarBUNITController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetString("BUNIT")
		cuadre := c.GetString("cuadre")
		descr := c.GetString("descr")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, errS := sql.Open("mssql", connString2)
		if errS != nil {
			log.Fatal("Perdon! Open connection failed:", errS.Error())
		}
		//PRIMERO TENGO QUE HACER UN SELECT PARA VER SI YA EXISTE!! SI YA EXISTE, PUES NO SE CREA NADA!
		queryDireccion := "SELECT BU_CODE FROM [SU_DOMAINDB].[dbo].[DOMN_BU_DSRCE_LINK] WHERE BU_CODE = '"+BUNIT+"'"
		rowsDireccion, err := conn.Query(queryDireccion)
		if rowsDireccion.Next()  {
			example := map[string]interface{}{ "success":2}
			c.Data["json"] = &example
			c.ServeJSON()
			return
		}
		
		//["+beego.AppConfig.String("serverAux")+"].
		//INSERT INTO
		currentTime := int64(time.Now().Unix())
		tm := time.Unix(currentTime, 0)
		dateString := tm.String() 
		substring := string(dateString[0:10])	   
		query := "INSERT INTO [SU_DOMAINDB].[dbo].[DOMN_BU_DSRCE_LINK] (BU_CODE,UPDATE_COUNT,LAST_CHANGE_USER_ID,LAST_CHANGE_DATETIME,DSRCE_ID,ESB_ENABLED) VALUES ('"+BUNIT+"',0,'"+alias.(string)+"','"+substring+"',1,0)"
		result , err := conn.Exec(query)
		if err != nil {
			defer conn.Close()
			log.Fatal("Perdon insert! : ", err.Error())
		}
		afectados, _ := result.RowsAffected()
		if afectados == 1 {
			queryZ := `INSERT INTO [SUNPLUSADV].[dbo].[DB_DEFN] (DB_CODE,UPDATE_COUNT,LAST_CHANGE_USER_ID,LAST_CHANGE_DATETIME,DESCR,DATE_FORMAT,DEC_SEP,THSND_SEP,PRIMARY_BUDGET,P_COMMITMENT_LEDGER,FILLER_113,SRCE_DB_CODE,SRCE_VAL1_CONV_CODE,SRCE_VAL3_CONV_CODE,SRCE_PIVOT_VAL,TRGT_VAL1_CONV_CODE,TRGT_VAL3_CONV_CODE,TRGT_PIVOT_VAL,TRGT_SRCE_VAL1,TRGT_SRCE_VAL2,TRGT_SRCE_VAL3,TRGT_SRCE_VALM,UPGRADE_DATETIME,PAYMENT_DRIVE,SB_DEC_PL,PIVOT_CURR_CODE,VAL1_BASE_NAME,VAL1_CONV_CODE,VAL1_BASE_POST_RULE,VAL1_BASE_AMT_DEC_PL,VAL1_BASE_AMT_BAL,VAL1_BASE_AMT_BAL_TOLR,VAL1_BASE_BAL_ACNT_DR,VAL1_BASE_BAL_ACNT_CR,VAL2_OTHER_NAME,VAL2_OTHER_POST_RULE,VAL2_OTHER_AMT_BAL,VAL2_OTHER_BAL_ACNT_DR,VAL2_OTHER_BAL_ACNT_CR,VAL3_CURR_TYPE,VAL3_RPTG_NAME,VAL3_RPTG_CURR,VAL3_RPTG_POST_RULE,VAL3_RPTG_AMT_DEC_PL,VAL3_RPTG_AMT_BAL,VAL3_RPTG_AMT_BAL_TOLR,VAL3_BAL_ACNT_DR,VAL3_BAL_ACNT_CR,VAL4_MEMO_NAME,VAL4_MEMO_POST_RULE,RESPONSE_DATE_SEP,MAX_NUM_OF_PERDS,DAGS_IN_USE_ITEM,DAGS_IN_USE_SUPP,DAGS_IN_USE_CUST,DAGS_IN_USE_DEFN,DAGS_IN_USE_WHSE,DAGS_IN_USE_FIN,DAG_CODE,LA_EXPTD_REVN_LDG,VAL4_NAME,VAL4_CURR_CODE,VAL4_POST_RULE,VAL4_AMT_BAL,VAL4_BAL_ACNT_DR,VAL4_BAL_ACNT_CR,VAL4_CALC_FROM,VAL5_NAME,VAL5_CURR_CODE,VAL5_POST_RULE,VAL5_AMT_BAL,VAL5_BAL_ACNT_DR,VAL5_BAL_ACNT_CR,VAL5_CALC_FROM,ALLOCN_LOCKING_ALLOWED,ALLOW_SIGNED_SPLITS,ALLOW_MULTIPLE_SPLITS,VAL4_AMT_BAL_TOLR,VAL5_AMT_BAL_TOLR,PST_DIFF_LRGST_CARRIER,CORRECT_DIFF,FUNDING_CURR,VAL1_BASE_WO_TOLR,VAL1_BASE_WO_ACNT_DR,VAL1_BASE_WO_ACNT_CR,MONTH_FORMAT,FIN_BU,DAGS_IN_USE_ZONE,DAGS_IN_USE_LOCN,FORM_BACKGROUND_COLOUR,DAGS_IN_USE_ACNT,DAGS_IN_USE_JNL_TYPE,DAGS_IN_USE_ANL_T0,DAGS_IN_USE_ANL_T1,DAGS_IN_USE_ANL_T2,DAGS_IN_USE_ANL_T3,DAGS_IN_USE_ANL_T4,DAGS_IN_USE_ANL_T5,DAGS_IN_USE_ANL_T6,DAGS_IN_USE_ANL_T7,DAGS_IN_USE_ANL_T8,DAGS_IN_USE_ANL_T9,DATA_ARCHIVED,ARCHIVE_BU_CODE,VAL4_CV4_AMT_DEC_PL,VAL5_CV4_AMT_DEC_PL)
			VALUES ('`+BUNIT+`',1,'`+alias.(string)+`','`+substring+`','`+descr+`',0,'.','.','B','C',NULL,'',NULL,NULL,0,NULL,NULL,0,0,0,0,0,NULL,'\payments\',2,1,'Pesos Mexicanos','MXP1',3,2,2,0.001,'`+cuadre+`','`+cuadre+`','Pesos',3,3,NULL,NULL,99,NULL,'',0,0,3,0.000,NULL,NULL,'Units',1,'/',12,0,0,0,0,0,0,NULL,NULL,'Pesos Mexicanos','MXP1',3,1,'`+cuadre+`','`+cuadre+`',1,NULL,NULL,0,3,NULL,NULL,0,0,0,0,0.001,0.000,0,0,99,0.000,NULL,NULL,0,1,0,0,99,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL,2,0)`
			_ , errZ := conn.Exec(queryZ)
			if errZ != nil {
				defer conn.Close()
				log.Fatal("Perdon insert Z! : ", errZ.Error())
			}
			//afectadosZ, _ := result.RowsAffected()
			
			query1 := `CREATE TABLE [SUNPLUSADV].[dbo].[`+BUNIT+`_A_SALFLDG](
			[ACCNT_CODE] [dbo].[REF_MAINT_CODE] NOT NULL,
			[PERIOD] [dbo].[FIN_PERD_INT] NOT NULL,
			[TRANS_DATETIME] [dbo].[FIN_DATETIME] NOT NULL,
			[JRNAL_NO] [dbo].[FIN_JNL_NUM] NOT NULL,
			[JRNAL_LINE] [dbo].[FIN_JNL_LINE] NOT NULL,
			[AMOUNT] [dbo].[FIN_AMOUNT_BASE] NOT NULL,
			[D_C] [dbo].[FIN_FLAG_D_C] NOT NULL,
			[ALLOCATION] [dbo].[FIN_FLAG_ALLOCN_BR] NOT NULL,
			[JRNAL_TYPE] [dbo].[FIN_JNL_TYPE] NULL,
			[JRNAL_SRCE] [dbo].[FIN_JNL_SRCE] NOT NULL,
			[TREFERENCE] [dbo].[TRANSACTION_REF_MIXED] NULL,
			[DESCRIPTN] [dbo].[FIN_DESCR_50] NULL,
			[ENTRY_DATETIME] [dbo].[FIN_DATETIME] NULL,
			[ENTRY_PRD] [dbo].[FIN_PERD_INT] NOT NULL,
			[DUE_DATETIME] [dbo].[FIN_DATETIME] NULL,
			[ALLOC_REF] [dbo].[FIN_ALLOCN_REF] NOT NULL,
			[ALLOC_DATETIME] [dbo].[FIN_DATETIME] NULL,
			[ALLOC_PERIOD] [dbo].[FIN_PERD_INT] NOT NULL,
			[ASSET_IND] [dbo].[FIN_FLAG_ASSET] NOT NULL,
			[ASSET_CODE] [dbo].[FIN_ASSET_CODE] NOT NULL,
			[ASSET_SUB] [dbo].[FIN_ASSET_SUB_CODE] NOT NULL,
			[CONV_CODE] [dbo].[FIN_CONV_CODE] NOT NULL,
			[CONV_RATE] [dbo].[FIN_CONV_RATE] NOT NULL,
			[OTHER_AMT] [dbo].[FIN_OTHER_AMT_ACC] NOT NULL,
			[OTHER_DP] [dbo].[FIN_FLAG_OTHER_DP] NOT NULL,
			[CLEARDOWN] [dbo].[FIN_SEQ] NOT NULL,
			[REVERSAL] [dbo].[FIN_FLAG_Y_N_NA] NOT NULL,
			[LOSS_GAIN] [dbo].[FIN_LOSS_GAIN] NOT NULL,
			[ROUGH_FLAG] [dbo].[FIN_FLAG_ROUGH] NOT NULL,
			[IN_USE_FLAG] [dbo].[FIN_FLAG_Y_N_NA] NOT NULL,
			[ANAL_T0] [dbo].[FIN_ANL_T0] NOT NULL,
			[ANAL_T1] [dbo].[FIN_ANL_T1] NOT NULL,
			[ANAL_T2] [dbo].[FIN_ANL_T2] NOT NULL,
			[ANAL_T3] [dbo].[FIN_ANL_T3] NOT NULL,
			[ANAL_T4] [dbo].[FIN_ANL_T4] NOT NULL,
			[ANAL_T5] [dbo].[FIN_ANL_T5] NOT NULL,
			[ANAL_T6] [dbo].[FIN_ANL_T6] NOT NULL,
			[ANAL_T7] [dbo].[FIN_ANL_T7] NOT NULL,
			[ANAL_T8] [dbo].[FIN_ANL_T8] NOT NULL,
			[ANAL_T9] [dbo].[FIN_ANL_T9] NOT NULL,
			[POSTING_DATETIME] [dbo].[FIN_DATETIME] NULL,
			[ALLOC_IN_PROGRESS] [dbo].[FIN_ALLOCN_IN_PROG] NOT NULL,
			[HOLD_REF] [dbo].[FIN_JNL_REF] NOT NULL,
			[HOLD_OP_ID] [dbo].[FIN_OPR_ID] NOT NULL,
			[BASE_RATE] [dbo].[FIN_CONV_RATE] NOT NULL,
			[BASE_OPERATOR] [dbo].[FIN_FLAG_OPN] NOT NULL,
			[CONV_OPERATOR] [dbo].[FIN_FLAG_OPN] NOT NULL,
			[REPORT_RATE] [dbo].[FIN_CONV_RATE] NOT NULL,
			[REPORT_OPERATOR] [dbo].[FIN_FLAG_OPN] NOT NULL,
			[REPORT_AMT] [dbo].[FIN_AMT_RPTG_ACC] NOT NULL,
			[MEMO_AMT] [dbo].[FIN_AMT_MEMO_ACC] NOT NULL,
			[EXCLUDE_BAL] [dbo].[FIN_FLAG_Y_N_NA] NOT NULL,
			[LE_DETAILS_IND] [dbo].[FIN_FLAG_LE_DETAILS] NOT NULL,
			[CONSUMED_BDGT_ID] [dbo].[NUM_INT] NOT NULL,
			[CV4_CONV_CODE] [dbo].[FIN_CONV_CODE] NOT NULL,
			[CV4_AMT] [dbo].[FIN_AMOUNT_BASE] NOT NULL,
			[CV4_CONV_RATE] [dbo].[FIN_CONV_RATE] NOT NULL,
			[CV4_OPERATOR] [dbo].[FIN_FLAG_OPN] NOT NULL,
			[CV4_DP] [dbo].[FIN_FLAG_OTHER_DP] NOT NULL,
			[CV5_CONV_CODE] [dbo].[FIN_CONV_CODE] NOT NULL,
			[CV5_AMT] [dbo].[FIN_AMOUNT_BASE] NOT NULL,
			[CV5_CONV_RATE] [dbo].[FIN_CONV_RATE] NOT NULL,
			[CV5_OPERATOR] [dbo].[FIN_FLAG_OPN] NOT NULL,
			[CV5_DP] [dbo].[FIN_FLAG_OTHER_DP] NOT NULL,
			[LINK_REF_1] [dbo].[CHAR_ALPHA_V15] NULL,
			[LINK_REF_2] [dbo].[CHAR_ALPHA_V15] NULL,
			[LINK_REF_3] [dbo].[CHAR_ALPHA_V15] NULL,
			[ALLOCN_CODE] [dbo].[CHAR_CODE_5] NULL,
			[ALLOCN_STMNTS] [dbo].[NUM_SMALLINT] NULL,
			[OPR_CODE] [dbo].[OPERATOR_CODE] NULL,
			[SPLIT_ORIG_LINE] [dbo].[NUM_INT] NOT NULL,
			[VAL_DATETIME] [dbo].[FIN_DATETIME] NULL,
			[SIGNING_DETAILS] [dbo].[STRING_V30] NULL,
			[INSTLMT_DATETIME] [dbo].[FIN_DATETIME] NULL,
			[PRINCIPAL_REQD] [dbo].[FLAG_Y_N] NOT NULL,
			[BINDER_STATUS] [dbo].[FIN_FLAG_BINDER_STATUS] NOT NULL,
			[AGREED_STATUS] [dbo].[FLAG_Y_N] NOT NULL,
			[SPLIT_LINK_REF] [dbo].[CHAR_ALPHA_V15] NULL,
			[PSTG_REF] [dbo].[CHAR_ALPHA_V15] NULL,
			[TRUE_RATED] [dbo].[FLAG_TRUE_RATED] NOT NULL,
			[HOLD_DATETIME] [dbo].[FIN_DATETIME] NULL,
			[HOLD_TEXT] [dbo].[STRING_V30] NULL,
			[INSTLMT_NUM] [dbo].[NUM_SMALLINT] NULL,
			[SUPPLMNTRY_EXTSN] [dbo].[FLAG_Y_N] NOT NULL,
			[APRVLS_EXTSN] [dbo].[FLAG_Y_N] NOT NULL,
			[REVAL_LINK_REF] [dbo].[NUM_INT] NULL,
			[SAVED_SET_NUM] [dbo].[NUM_18_0] NULL,
			[AUTHORISTN_SET_REF] [int] NULL,
			[PYMT_AUTHORISTN_SET_REF] [int] NULL,
			[MAN_PAY_OVER] [dbo].[FLAG_Y_N] NOT NULL,
			[PYMT_STAMP] [dbo].[STRING_V10] NULL,
			[AUTHORISTN_IN_PROGRESS] [dbo].[FLAG_Y_N] NOT NULL,
			[SPLIT_IN_PROGRESS] [dbo].[FLAG_Y_N] NOT NULL,
			[VCHR_NUM] [dbo].[CHAR_ALPHA_V30] NULL,
			[JNL_CLASS_CODE] [dbo].[REF_MAINT_CODE] NULL,
			[ORIGINATOR_ID] [dbo].[OPERATOR_CODE] NULL,
			[ORIGINATED_DATETIME] [datetime] NULL,
			[LAST_CHANGE_USER_ID] [dbo].[LAST_CHANGE_USER_ID] NULL,
			[LAST_CHANGE_DATETIME] [datetime] NULL,
			[AFTER_PSTG_ID] [dbo].[OPERATOR_CODE] NULL,
			[AFTER_PSTG_DATETIME] [datetime] NULL,
			[POSTER_ID] [dbo].[OPERATOR_CODE] NULL,
			[ALLOC_ID] [dbo].[OPERATOR_CODE] NULL,
			[JNL_REVERSAL_TYPE] [dbo].[FLAG_JNL_REVERSAL_TYPE] NOT NULL,
			PRIMARY KEY CLUSTERED 
			(
				[ACCNT_CODE] ASC,
				[PERIOD] ASC,
				[TRANS_DATETIME] ASC,
				[JRNAL_NO] ASC,
				[JRNAL_LINE] ASC
			)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
			) ON [PRIMARY]`
			_ , err1 := conn.Exec(query1)
			if err1 != nil {
				defer conn.Close()
				log.Fatal("Perdon 1! : ", err1.Error())
			}
			query2 := `CREATE TABLE [dbo].[`+BUNIT+`_B_SALFLDG](
			[ACCNT_CODE] [dbo].[REF_MAINT_CODE] NOT NULL,
			[PERIOD] [dbo].[FIN_PERD_INT] NOT NULL,
			[TRANS_DATETIME] [dbo].[FIN_DATETIME] NOT NULL,
			[JRNAL_NO] [dbo].[FIN_JNL_NUM] NOT NULL,
			[JRNAL_LINE] [dbo].[FIN_JNL_LINE] NOT NULL,
			[AMOUNT] [dbo].[FIN_AMOUNT_BASE] NOT NULL,
			[D_C] [dbo].[FIN_FLAG_D_C] NOT NULL,
			[ALLOCATION] [dbo].[FIN_FLAG_ALLOCN_BR] NOT NULL,
			[JRNAL_TYPE] [dbo].[FIN_JNL_TYPE] NULL,
			[JRNAL_SRCE] [dbo].[FIN_JNL_SRCE] NOT NULL,
			[TREFERENCE] [dbo].[TRANSACTION_REF_MIXED] NULL,
			[DESCRIPTN] [dbo].[FIN_DESCR_50] NULL,
			[ENTRY_DATETIME] [dbo].[FIN_DATETIME] NULL,
			[ENTRY_PRD] [dbo].[FIN_PERD_INT] NOT NULL,
			[DUE_DATETIME] [dbo].[FIN_DATETIME] NULL,
			[ALLOC_REF] [dbo].[FIN_ALLOCN_REF] NOT NULL,
			[ALLOC_DATETIME] [dbo].[FIN_DATETIME] NULL,
			[ALLOC_PERIOD] [dbo].[FIN_PERD_INT] NOT NULL,
			[ASSET_IND] [dbo].[FIN_FLAG_ASSET] NOT NULL,
			[ASSET_CODE] [dbo].[FIN_ASSET_CODE] NOT NULL,
			[ASSET_SUB] [dbo].[FIN_ASSET_SUB_CODE] NOT NULL,
			[CONV_CODE] [dbo].[FIN_CONV_CODE] NOT NULL,
			[CONV_RATE] [dbo].[FIN_CONV_RATE] NOT NULL,
			[OTHER_AMT] [dbo].[FIN_OTHER_AMT_ACC] NOT NULL,
			[OTHER_DP] [dbo].[FIN_FLAG_OTHER_DP] NOT NULL,
			[CLEARDOWN] [dbo].[FIN_SEQ] NOT NULL,
			[REVERSAL] [dbo].[FIN_FLAG_Y_N_NA] NOT NULL,
			[LOSS_GAIN] [dbo].[FIN_LOSS_GAIN] NOT NULL,
			[ROUGH_FLAG] [dbo].[FIN_FLAG_ROUGH] NOT NULL,
			[IN_USE_FLAG] [dbo].[FIN_FLAG_Y_N_NA] NOT NULL,
			[ANAL_T0] [dbo].[FIN_ANL_T0] NOT NULL,
			[ANAL_T1] [dbo].[FIN_ANL_T1] NOT NULL,
			[ANAL_T2] [dbo].[FIN_ANL_T2] NOT NULL,
			[ANAL_T3] [dbo].[FIN_ANL_T3] NOT NULL,
			[ANAL_T4] [dbo].[FIN_ANL_T4] NOT NULL,
			[ANAL_T5] [dbo].[FIN_ANL_T5] NOT NULL,
			[ANAL_T6] [dbo].[FIN_ANL_T6] NOT NULL,
			[ANAL_T7] [dbo].[FIN_ANL_T7] NOT NULL,
			[ANAL_T8] [dbo].[FIN_ANL_T8] NOT NULL,
			[ANAL_T9] [dbo].[FIN_ANL_T9] NOT NULL,
			[POSTING_DATETIME] [dbo].[FIN_DATETIME] NULL,
			[ALLOC_IN_PROGRESS] [dbo].[FIN_ALLOCN_IN_PROG] NOT NULL,
			[HOLD_REF] [dbo].[FIN_JNL_REF] NOT NULL,
			[HOLD_OP_ID] [dbo].[FIN_OPR_ID] NOT NULL,
			[BASE_RATE] [dbo].[FIN_CONV_RATE] NOT NULL,
			[BASE_OPERATOR] [dbo].[FIN_FLAG_OPN] NOT NULL,
			[CONV_OPERATOR] [dbo].[FIN_FLAG_OPN] NOT NULL,
			[REPORT_RATE] [dbo].[FIN_CONV_RATE] NOT NULL,
			[REPORT_OPERATOR] [dbo].[FIN_FLAG_OPN] NOT NULL,
			[REPORT_AMT] [dbo].[FIN_AMT_RPTG_ACC] NOT NULL,
			[MEMO_AMT] [dbo].[FIN_AMT_MEMO_ACC] NOT NULL,
			[EXCLUDE_BAL] [dbo].[FIN_FLAG_Y_N_NA] NOT NULL,
			[LE_DETAILS_IND] [dbo].[FIN_FLAG_LE_DETAILS] NOT NULL,
			[CONSUMED_BDGT_ID] [dbo].[NUM_INT] NOT NULL,
			[CV4_CONV_CODE] [dbo].[FIN_CONV_CODE] NOT NULL,
			[CV4_AMT] [dbo].[FIN_AMOUNT_BASE] NOT NULL,
			[CV4_CONV_RATE] [dbo].[FIN_CONV_RATE] NOT NULL,
			[CV4_OPERATOR] [dbo].[FIN_FLAG_OPN] NOT NULL,
			[CV4_DP] [dbo].[FIN_FLAG_OTHER_DP] NOT NULL,
			[CV5_CONV_CODE] [dbo].[FIN_CONV_CODE] NOT NULL,
			[CV5_AMT] [dbo].[FIN_AMOUNT_BASE] NOT NULL,
			[CV5_CONV_RATE] [dbo].[FIN_CONV_RATE] NOT NULL,
			[CV5_OPERATOR] [dbo].[FIN_FLAG_OPN] NOT NULL,
			[CV5_DP] [dbo].[FIN_FLAG_OTHER_DP] NOT NULL,
			[LINK_REF_1] [dbo].[CHAR_ALPHA_V15] NULL,
			[LINK_REF_2] [dbo].[CHAR_ALPHA_V15] NULL,
			[LINK_REF_3] [dbo].[CHAR_ALPHA_V15] NULL,
			[ALLOCN_CODE] [dbo].[CHAR_CODE_5] NULL,
			[ALLOCN_STMNTS] [dbo].[NUM_SMALLINT] NULL,
			[OPR_CODE] [dbo].[OPERATOR_CODE] NULL,
			[SPLIT_ORIG_LINE] [dbo].[NUM_INT] NOT NULL,
			[VAL_DATETIME] [dbo].[FIN_DATETIME] NULL,
			[SIGNING_DETAILS] [dbo].[STRING_V30] NULL,
			[INSTLMT_DATETIME] [dbo].[FIN_DATETIME] NULL,
			[PRINCIPAL_REQD] [dbo].[FLAG_Y_N] NOT NULL,
			[BINDER_STATUS] [dbo].[FIN_FLAG_BINDER_STATUS] NOT NULL,
			[AGREED_STATUS] [dbo].[FLAG_Y_N] NOT NULL,
			[SPLIT_LINK_REF] [dbo].[CHAR_ALPHA_V15] NULL,
			[PSTG_REF] [dbo].[CHAR_ALPHA_V15] NULL,
			[TRUE_RATED] [dbo].[FLAG_TRUE_RATED] NOT NULL,
			[HOLD_DATETIME] [dbo].[FIN_DATETIME] NULL,
			[HOLD_TEXT] [dbo].[STRING_V30] NULL,
			[INSTLMT_NUM] [dbo].[NUM_SMALLINT] NULL,
			[SUPPLMNTRY_EXTSN] [dbo].[FLAG_Y_N] NOT NULL,
			[APRVLS_EXTSN] [dbo].[FLAG_Y_N] NOT NULL,
			[REVAL_LINK_REF] [dbo].[NUM_INT] NULL,
			[SAVED_SET_NUM] [dbo].[NUM_18_0] NULL,
			[AUTHORISTN_SET_REF] [int] NULL,
			[PYMT_AUTHORISTN_SET_REF] [int] NULL,
			[MAN_PAY_OVER] [dbo].[FLAG_Y_N] NOT NULL,
			[PYMT_STAMP] [dbo].[STRING_V10] NULL,
			[AUTHORISTN_IN_PROGRESS] [dbo].[FLAG_Y_N] NOT NULL,
			[SPLIT_IN_PROGRESS] [dbo].[FLAG_Y_N] NOT NULL,
			[VCHR_NUM] [dbo].[CHAR_ALPHA_V30] NULL,
			[JNL_CLASS_CODE] [dbo].[REF_MAINT_CODE] NULL,
			[ORIGINATOR_ID] [dbo].[OPERATOR_CODE] NULL,
			[ORIGINATED_DATETIME] [datetime] NULL,
			[LAST_CHANGE_USER_ID] [dbo].[LAST_CHANGE_USER_ID] NULL,
			[LAST_CHANGE_DATETIME] [datetime] NULL,
			[AFTER_PSTG_ID] [dbo].[OPERATOR_CODE] NULL,
			[AFTER_PSTG_DATETIME] [datetime] NULL,
			[POSTER_ID] [dbo].[OPERATOR_CODE] NULL,
			[ALLOC_ID] [dbo].[OPERATOR_CODE] NULL,
			[JNL_REVERSAL_TYPE] [dbo].[FLAG_JNL_REVERSAL_TYPE] NOT NULL,
			PRIMARY KEY CLUSTERED 
			(
				[ACCNT_CODE] ASC,
				[PERIOD] ASC,
				[TRANS_DATETIME] ASC,
				[JRNAL_NO] ASC,
				[JRNAL_LINE] ASC
			)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
			) ON [PRIMARY]`
			_ , err2 := conn.Exec(query2)
			if err2 != nil {
				defer conn.Close()
				log.Fatal("Perdon 2! : ", err2.Error())
			}
			query3 := `CREATE TABLE [dbo].[`+BUNIT+`_ACNT](
			[ACNT_CODE] [dbo].[REF_MAINT_CODE] NOT NULL,
			[UPDATE_COUNT] [dbo].[UPDATE_CNT] NOT NULL,
			[LAST_CHANGE_USER_ID] [dbo].[LAST_CHANGE_USER_ID] NOT NULL,
			[LAST_CHANGE_DATETIME] [dbo].[LAST_CHANGE_DATETIME] NOT NULL,
			[DESCR] [dbo].[DESCRIPTION_50] NOT NULL,
			[S_HEAD] [dbo].[SHORT_HEADING] NULL,
			[LOOKUP] [dbo].[LOOKUP] NOT NULL,
			[DAG_CODE] [dbo].[DATA_ACCESS_GROUP] NULL,
			[ACNT_TYPE] [dbo].[FLAG_ACNT_TYPE] NOT NULL,
			[BAL_TYPE] [dbo].[FLAG_BAL_TYPE] NOT NULL,
			[STATUS] [dbo].[STATUS] NOT NULL,
			[SUPPRESS_REVAL] [dbo].[FLAG_Y_N] NOT NULL,
			[LONG_DESCR] [dbo].[LONG_DESCRIPTION] NULL,
			[CONV_CODE_CTRL] [dbo].[FLAG_CONV_CODE_CTRL] NOT NULL,
			[DFLT_CURR_CODE] [dbo].[CURRENCY_CODE] NULL,
			[ALLOCN_IN_PROGRESS] [dbo].[FLAG_Y_N_NA] NOT NULL,
			[LINK_ACNT] [dbo].[REF_MAINT_CODE] NOT NULL,
			[RPT_CONV_CTRL] [dbo].[FLAG_Y_N] NOT NULL,
			[USER_AREA] [dbo].[CHAR_ALPHA_V30] NULL,
			[CR_LIMIT] [dbo].[NUM_18_0_SILS] NULL,
			[ENTER_ANL_1] [dbo].[FLAG_ENTER_ANL] NOT NULL,
			[ENTER_ANL_2] [dbo].[FLAG_ENTER_ANL] NOT NULL,
			[ENTER_ANL_3] [dbo].[FLAG_ENTER_ANL] NOT NULL,
			[ENTER_ANL_4] [dbo].[FLAG_ENTER_ANL] NOT NULL,
			[ENTER_ANL_5] [dbo].[FLAG_ENTER_ANL] NOT NULL,
			[ENTER_ANL_6] [dbo].[FLAG_ENTER_ANL] NOT NULL,
			[ENTER_ANL_7] [dbo].[FLAG_ENTER_ANL] NOT NULL,
			[ENTER_ANL_8] [dbo].[FLAG_ENTER_ANL] NOT NULL,
			[ENTER_ANL_9] [dbo].[FLAG_ENTER_ANL] NOT NULL,
			[ENTER_ANL_10] [dbo].[FLAG_ENTER_ANL] NOT NULL,
			[OIL] [dbo].[NUM_18_0_SILS] NULL,
			[CV4_DFLT_CURR_CODE] [dbo].[CURRENCY_CODE] NULL,
			[CV4_CONV_CODE_CTRL] [dbo].[FLAG_CONV_CODE_CTRL] NOT NULL,
			[CV5_DFLT_CURR_CODE] [dbo].[CURRENCY_CODE] NULL,
			[CV5_CONV_CODE_CTRL] [dbo].[FLAG_CONV_CODE_CTRL] NOT NULL,
			[BANK_CURR_REQD] [dbo].[FLAG_Y_N] NOT NULL,
			[ACNT_LINKS_ALLOWED] [dbo].[FLAG_Y_N] NOT NULL,
			[PASP_ACNT_TYPE] [dbo].[FLAG_PASP_ACNT_TYPE] NOT NULL,
			PRIMARY KEY CLUSTERED 
			(
				[ACNT_CODE] ASC
			)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
			) ON [PRIMARY]`
			_ , err3 := conn.Exec(query3)
			if err3 != nil {
				defer conn.Close()
				log.Fatal("Perdon 3! : ", err3.Error())
			}
			query4 := `CREATE TABLE [dbo].[`+BUNIT+`_ADDR](
			[ADDR_CODE] [dbo].[REF_MAINT_CODE] NOT NULL,
			[UPDATE_COUNT] [dbo].[UPDATE_CNT] NOT NULL,
			[LAST_CHANGE_USER_ID] [dbo].[LAST_CHANGE_USER_ID] NOT NULL,
			[LAST_CHANGE_DATETIME] [dbo].[LAST_CHANGE_DATETIME] NOT NULL,
			[STATUS] [dbo].[STATUS] NOT NULL,
			[S_HEAD] [dbo].[SHORT_HEADING] NULL,
			[LOOKUP] [dbo].[LOOKUP] NOT NULL,
			[TELEPHONE_NO] [dbo].[PHONE_FAX] NULL,
			[TELEX_FAX_NO] [dbo].[PHONE_FAX] NULL,
			[ADDR_LINE_1] [dbo].[STRING_V50] NULL,
			[ADDR_LINE_2] [dbo].[STRING_V50] NULL,
			[ADDR_LINE_3] [dbo].[STRING_V50] NULL,
			[ADDR_LINE_4] [dbo].[STRING_V50] NULL,
			[ADDR_LINE_5] [dbo].[STRING_V50] NULL,
			[TOWN_CITY] [dbo].[STRING_V50] NULL,
			[STATE] [dbo].[STRING_V50] NULL,
			[POST_CODE] [dbo].[POSTCODE] NULL,
			[COUNTRY] [dbo].[STRING_V50] NULL,
			[AREA] [dbo].[CHAR_ALPHA_V40] NULL,
			[CMMNT] [dbo].[LONG_DESCRIPTION] NULL,
			[EDI_PROFILE_ID] [int] NULL,
			[TIME_ZONE] [dbo].[NUM_SMALLINT_V2] NULL,
			[LANG_CODE] [dbo].[LANGUAGE_CODE] NULL,
			[COUNTRY_CODE] [dbo].[CHAR_CODE_3] NULL,
			[STATE_CODE] [dbo].[CHAR_CODE_2] NULL,
			[LINK_ADDR_CODE] [dbo].[REF_MAINT_CODE] NULL,
			[USER_AREA] [dbo].[CHAR_CODE_3] NULL,
			[WEB_PAGE_ADDR] [dbo].[WEB_ADDRESS] NULL,
			[TEMP_ADDR] [dbo].[FLAG_Y_N] NOT NULL,
			[DAG_CODE] [dbo].[DATA_ACCESS_GROUP] NULL,
			PRIMARY KEY CLUSTERED 
			(
				[ADDR_CODE] ASC
			)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
			) ON [PRIMARY]`
			_ , err4 := conn.Exec(query4)
			if err4 != nil {
				defer conn.Close()
				log.Fatal("Perdon 4! : ", err4.Error())
			}
			query5 := `CREATE TABLE [dbo].[`+BUNIT+`_ANL_CAT](
			[ANL_CAT_ID] [dbo].[ANALYSIS_DIMENSION] NOT NULL,
			[UPDATE_COUNT] [dbo].[UPDATE_CNT] NOT NULL,
			[LAST_CHANGE_USER_ID] [dbo].[LAST_CHANGE_USER_ID] NOT NULL,
			[LAST_CHANGE_DATETIME] [dbo].[LAST_CHANGE_DATETIME] NOT NULL,
			[STATUS] [dbo].[STATUS] NOT NULL,
			[LOOKUP] [dbo].[LOOKUP] NOT NULL,
			[USEABLE_ANL_ENT_ID] [dbo].[TBL_NUM] NULL,
			[S_HEAD] [dbo].[SHORT_HEADING_UPPER] NOT NULL,
			[DESCR] [dbo].[CHAR_ALPHA_V40] NOT NULL,
			[DAG_CODE] [dbo].[DATA_ACCESS_GROUP] NULL,
			[AMEND_CODE] [dbo].[FLAG_Y_N] NOT NULL,
			[VALIDATE_IND] [dbo].[FLAG_VALIDATION_IND] NOT NULL,
			[LNGTH] [dbo].[NUM_SMALLINT_NO_SIGN] NOT NULL,
			[LINKED] [dbo].[FLAG_Y_N] NOT NULL,
			[IBUS_CODE_DIM_ID] [dbo].[NUM_INT] NULL,
			PRIMARY KEY CLUSTERED 
			(
				[ANL_CAT_ID] ASC
			)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
			) ON [PRIMARY]`
			_ , err5 := conn.Exec(query5)
			if err5 != nil {
				defer conn.Close()
				log.Fatal("Perdon 5! : ", err5.Error())
			}
			query6 := `CREATE TABLE [dbo].[`+BUNIT+`_ANL_CODE](
			[ANL_CAT_ID] [dbo].[ANALYSIS_DIMENSION] NOT NULL,
			[ANL_CODE] [dbo].[REF_MAINT_CODE] NOT NULL,
			[UPDATE_COUNT] [dbo].[UPDATE_CNT] NOT NULL,
			[LAST_CHANGE_USER_ID] [dbo].[LAST_CHANGE_USER_ID] NOT NULL,
			[LAST_CHANGE_DATETIME] [dbo].[LAST_CHANGE_DATETIME] NOT NULL,
			[STATUS] [dbo].[STATUS] NOT NULL,
			[LOOKUP] [dbo].[LOOKUP] NOT NULL,
			[NAME] [dbo].[NAME] NULL,
			[DAG_CODE] [dbo].[DATA_ACCESS_GROUP] NULL,
			[BDGT_CHECK] [dbo].[FLAG_BDGT_CHECK] NOT NULL,
			[BDGT_STOP] [dbo].[FLAG_Y_N] NOT NULL,
			[PROHIBIT_POSTING] [dbo].[FLAG_Y_N] NOT NULL,
			[NAVIGATION_OPTION] [dbo].[FLAG_NAVIGATION_OPTION] NOT NULL,
			[COMBINED_BDGT_CHCK] [dbo].[FLAG_Y_N] NOT NULL,
			PRIMARY KEY CLUSTERED 
			(
				[ANL_CAT_ID] ASC,
				[ANL_CODE] ASC
			)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
			) ON [PRIMARY]`
			_ , err6 := conn.Exec(query6)
			if err6 != nil {
				defer conn.Close()
				log.Fatal("Perdon 6! : ", err6.Error())
			}
			query7 := `CREATE TABLE [dbo].[`+BUNIT+`_ASSET](
			[ASSET_CODE] [dbo].[CHAR_CODE_10_UPPER] NOT NULL,
			[UPDATE_COUNT] [dbo].[UPDATE_CNT] NOT NULL,
			[LAST_CHANGE_USER_ID] [dbo].[LAST_CHANGE_USER_ID] NOT NULL,
			[LAST_CHANGE_DATETIME] [dbo].[LAST_CHANGE_DATETIME] NOT NULL,
			[STATUS] [dbo].[STATUS] NOT NULL,
			[ASSET_STATUS] [dbo].[STATUS_ASSET] NOT NULL,
			[S_HEAD] [dbo].[SHORT_HEADING] NULL,
			[LOOKUP] [dbo].[LOOKUP] NOT NULL,
			[DESCR] [dbo].[DESCRIPTION_50] NOT NULL,
			[START_PERD] [dbo].[PERD] NULL,
			[END_PERD] [dbo].[PERD] NULL,
			[LAST_PERD] [dbo].[PERD] NULL,
			[DISPOSAL_PERD] [dbo].[PERD] NULL,
			[BAL_SHEET_ACNT] [dbo].[REF_MAINT_CODE] NULL,
			[PROFIT_LOSS_ACNT] [dbo].[REF_MAINT_CODE] NULL,
			[DISPOSED] [dbo].[FLAG_Y_N] NOT NULL,
			[BASE_GROSS] [dbo].[NUM_18_3_SILS_ACC] NOT NULL,
			[BASE_DEP] [dbo].[NUM_18_3_SILS_ACC] NOT NULL,
			[BASE_NET] [dbo].[NUM_18_3_SILS_ACC] NOT NULL,
			[BASE_DEP_MTHD] [dbo].[FLAG_DEP_MTHD] NOT NULL,
			[BASE_PCENT] [dbo].[PERCENTAGE_FIN] NULL,
			[BASE_FINAL] [dbo].[NUM_18_3_SILS_ACC] NULL,
			[BASE_TBL_CODE] [dbo].[CHAR_CODE_5] NULL,
			[BASE_TBL_COL] [dbo].[CHAR_CODE_2] NULL,
			[ASSET_CONV_CODE] [dbo].[CURRENCY_CODE] NULL,
			[TXN_GROSS] [dbo].[NUM_18_3_SILS_ACC] NOT NULL,
			[TXN_DEP] [dbo].[NUM_18_3_SILS_ACC] NOT NULL,
			[TXN_NET] [dbo].[NUM_18_3_SILS_ACC] NOT NULL,
			[TXN_DEP_MTHD] [dbo].[FLAG_DEP_MTHD] NOT NULL,
			[TXN_PCENT] [dbo].[PERCENTAGE_FIN] NULL,
			[TXN_FINAL] [dbo].[NUM_18_3_SILS_ACC] NULL,
			[TXN_TBL_CODE] [dbo].[CHAR_CODE_5] NULL,
			[TXN_TBL_COL] [dbo].[CHAR_CODE_2] NULL,
			[FINAL_VAL_OVER] [dbo].[FLAG_Y_N] NOT NULL,
			[BASE_DAYS_FIRST_PERD] [dbo].[NUM_SMALLINT] NULL,
			[BASE_ACTIVE_DAYS] [dbo].[NUM_SMALLINT] NULL,
			[TXN_DAYS_FIRST_PERD] [dbo].[NUM_SMALLINT] NULL,
			[TXN_ACTIVE_DAYS] [dbo].[NUM_SMALLINT] NULL,
			[LAST_DEP_AMT] [dbo].[NUM_18_3_SILS_ACC] NULL,
			[LAST_DEP_TXN_AMT] [dbo].[NUM_18_3_SILS_ACC] NULL,
			[BASE_POST_FINAL_VAL] [dbo].[NUM_18_3_SILS_ACC] NOT NULL,
			[TXN_POST_FINAL_VAL] [dbo].[NUM_18_3_SILS_ACC] NOT NULL,
			[RPT_ACCUM_GROSS] [dbo].[NUM_18_3_SILS_ACC] NOT NULL,
			[RPT_DEP] [dbo].[NUM_18_3_SILS_ACC] NOT NULL,
			[RPT_NET] [dbo].[NUM_18_3_SILS_ACC] NOT NULL,
			[RPT_CONV_CODE_CTRL] [dbo].[FLAG_CONV_CODE_CTRL] NOT NULL,
			[CV4_CONV_CODE] [dbo].[FIN_CONV_CODE] NULL,
			[CV4_GROSS] [dbo].[NUM_18_3_SILS_ACC] NOT NULL,
			[CV4_DEP] [dbo].[NUM_18_3_SILS_ACC] NOT NULL,
			[CV4_NET] [dbo].[NUM_18_3_SILS_ACC] NOT NULL,
			[CV4_DEP_MTHD] [dbo].[FLAG_DEP_MTHD] NOT NULL,
			[CV4_PCENT] [dbo].[PERCENTAGE_FIN] NULL,
			[CV4_FINAL] [dbo].[NUM_18_3_SILS_ACC] NULL,
			[CV4_TBL_CODE] [dbo].[CHAR_CODE_5] NULL,
			[CV4_TBL_COL] [dbo].[CHAR_CODE_2] NULL,
			[CV4_DAYS_FIRST_PERD] [dbo].[NUM_SMALLINT] NULL,
			[CV4_LAST_DEP_AMT] [dbo].[NUM_18_3_ACC] NULL,
			[CV4_ACTIVE_DAYS] [dbo].[NUM_SMALLINT] NULL,
			[CV4_POST_FINAL_VAL] [dbo].[NUM_18_3_SILS_ACC] NOT NULL,
			[RPT_DEP_MTHD] [dbo].[FLAG_DEP_MTHD] NOT NULL,
			[RPT_PCENT] [dbo].[PERCENTAGE_FIN] NULL,
			[RPT_FINAL] [dbo].[NUM_18_3_SILS_ACC] NULL,
			[RPT_TBL_CODE] [dbo].[CHAR_CODE_5] NULL,
			[RPT_TBL_COL] [dbo].[CHAR_CODE_2] NULL,
			[RPT_DAYS_FIRST_PERD] [dbo].[NUM_SMALLINT] NULL,
			[RPT_ACTIVE_DAYS] [dbo].[NUM_SMALLINT] NULL,
			[RPT_LAST_DEP_AMT] [dbo].[NUM_18_3_ACC] NULL,
			[RPT_POST_FINAL_VAL] [dbo].[NUM_18_3_SILS_ACC] NOT NULL,
			[ASSET_QTY] [dbo].[NUM_18_5_ACC] NULL,
			[ANTICIPATED_BS_ACNT] [dbo].[REF_MAINT_CODE] NULL,
			[ANTICIPATED_PL_ACNT] [dbo].[REF_MAINT_CODE] NULL,
			[BASE_ANTICIPATED_DEP] [dbo].[NUM_18_3_SILS_ACC] NOT NULL,
			[BASE_ANTICIPATED_FROM] [dbo].[NUM_SMALLINT_NO_SIGN] NULL,
			[BASE_ANTICIPATED_YEARS] [dbo].[NUM_SMALLINT] NULL,
			[BASE_ANTICIPATED_FACTOR] [dbo].[NUM_5_3] NULL,
			[BASE_ANTICIPATED_CUR_YR] [dbo].[FLAG_Y_N] NOT NULL,
			[BASE_ANTICIPATED_LST_YR] [dbo].[NUM_SMALLINT_NO_SIGN] NULL,
			[BASE_FIRST_YEAR_PCENT] [dbo].[NUM_6_3] NULL,
			[BASE_REDUCTION_PCENT] [dbo].[NUM_6_3] NULL,
			[BASE_REDUCTION_LST_YR] [dbo].[NUM_SMALLINT_NO_SIGN] NULL,
			[TXN_ANTICIPATED_DEP] [dbo].[NUM_18_3_SILS_ACC] NOT NULL,
			[TXN_ANTICIPATED_FROM] [dbo].[NUM_SMALLINT_NO_SIGN] NULL,
			[TXN_ANTICIPATED_YEARS] [dbo].[NUM_SMALLINT] NULL,
			[TXN_ANTICIPATED_FACTOR] [dbo].[NUM_5_3] NULL,
			[TXN_ANTICIPATED_CUR_YR] [dbo].[FLAG_Y_N] NOT NULL,
			[TXN_ANTICIPATED_LST_YR] [dbo].[NUM_SMALLINT_NO_SIGN] NULL,
			[TXN_FIRST_YEAR_PCENT] [dbo].[NUM_6_3] NULL,
			[TXN_REDUCTION_PCENT] [dbo].[NUM_6_3] NULL,
			[TXN_REDUCTION_LST_YR] [dbo].[NUM_SMALLINT_NO_SIGN] NULL,
			[RPT_ANTICIPATED_DEP] [dbo].[NUM_18_3_SILS_ACC] NOT NULL,
			[RPT_ANTICIPATED_FROM] [dbo].[NUM_SMALLINT_NO_SIGN] NULL,
			[RPT_ANTICIPATED_YEARS] [dbo].[NUM_SMALLINT] NULL,
			[RPT_ANTICIPATED_FACTOR] [dbo].[NUM_5_3] NULL,
			[RPT_ANTICIPATED_CUR_YR] [dbo].[FLAG_Y_N] NOT NULL,
			[RPT_ANTICIPATED_LST_YR] [dbo].[NUM_SMALLINT_NO_SIGN] NULL,
			[RPT_FIRST_YEAR_PCENT] [dbo].[NUM_6_3] NULL,
			[RPT_REDUCTION_PCENT] [dbo].[NUM_6_3] NULL,
			[RPT_REDUCTION_LST_YR] [dbo].[NUM_SMALLINT_NO_SIGN] NULL,
			[ASSET_CLASS_CODE] [dbo].[CHAR_CODE_10_UPPER] NULL,
			[PART_DISPOSED] [dbo].[FLAG_Y_N] NOT NULL,
			[PART_DISPOSAL_PERD] [dbo].[PERD] NULL,
			PRIMARY KEY CLUSTERED 
			(
				[ASSET_CODE] ASC
			)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
			) ON [PRIMARY]`
			_ , err7 := conn.Exec(query7)
			if err7 != nil {
				defer conn.Close()
				log.Fatal("Perdon 7! : ", err7.Error())
			}
			query8 := `CREATE TABLE [dbo].[`+BUNIT+`_ASSET_ANL_CAT](
			[ANL_CAT_ID] [dbo].[ANALYSIS_DIMENSION] NOT NULL,
			[ASSET_CODE] [dbo].[CHAR_CODE_10_UPPER] NOT NULL,
			[UPDATE_COUNT] [dbo].[UPDATE_CNT] NOT NULL,
			[LAST_CHANGE_USER_ID] [dbo].[LAST_CHANGE_USER_ID] NOT NULL,
			[LAST_CHANGE_DATETIME] [dbo].[LAST_CHANGE_DATETIME] NOT NULL,
			[ANL_CODE] [dbo].[ANL_CODE_0] NOT NULL,
			PRIMARY KEY CLUSTERED 
			(
				[ANL_CAT_ID] ASC,
				[ASSET_CODE] ASC
			)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
			) ON [PRIMARY]`
			_ , err8 := conn.Exec(query8)
			if err8 != nil {
				defer conn.Close()
				log.Fatal("Perdon 8! : ", err8.Error())
			}
			query9 := `CREATE TABLE [dbo].[`+BUNIT+`_NOTE](
			[NOTE_ID] [int] IDENTITY(1,1) NOT NULL,
			[UPDATE_COUNT] [dbo].[UPDATE_CNT] NOT NULL,
			[LAST_CHANGE_USER_ID] [dbo].[LAST_CHANGE_USER_ID] NOT NULL,
			[LAST_CHANGE_DATETIME] [dbo].[LAST_CHANGE_DATETIME] NOT NULL,
			PRIMARY KEY CLUSTERED 
			(
				[NOTE_ID] ASC
			)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
			) ON [PRIMARY]`
			_ , err9 := conn.Exec(query9)
			if err9 != nil {
				defer conn.Close()
				log.Fatal("Perdon 9! : ", err9.Error())
			}
			queryA := `CREATE TABLE [dbo].[`+BUNIT+`_NOTE_DETAIL](
			[NOTE_ID] [int] NOT NULL,
			[NOTE_SEQ] [dbo].[NUM_SMALLINT] NOT NULL,
			[UPDATE_COUNT] [dbo].[UPDATE_CNT] NOT NULL,
			[LAST_CHANGE_USER_ID] [dbo].[LAST_CHANGE_USER_ID] NOT NULL,
			[LAST_CHANGE_DATETIME] [dbo].[LAST_CHANGE_DATETIME] NOT NULL,
			[CREATE_DATETIME] [datetime] NULL,
			[CREATE_OPR] [dbo].[OPERATOR_CODE] NOT NULL,
			[INTERNAL_ONLY] [dbo].[FLAG_Y_N] NOT NULL,
			[NOTE_CLASS_CODE] [dbo].[REF_MAINT_CODE] NULL,
			[NOTE_TEXT] [dbo].[STRING_V1000] NOT NULL,
			PRIMARY KEY CLUSTERED 
			(
				[NOTE_ID] ASC,
				[NOTE_SEQ] ASC
			)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
			) ON [PRIMARY]`
			_ , errA := conn.Exec(queryA)
			if errA != nil {
				defer conn.Close()
				log.Fatal("Perdon A! : ", errA.Error())
			}
			queryB := `CREATE TABLE [dbo].[`+BUNIT+`_NOTE_LDG](
			[ACNT_CODE] [dbo].[REF_MAINT_CODE] NOT NULL,
			[PERD] [dbo].[FIN_PERD_INT] NOT NULL,
			[TXN_DATETIME] [dbo].[FIN_DATETIME] NOT NULL,
			[JNL_NUM] [dbo].[FIN_JNL_NUM] NOT NULL,
			[JNL_LINE] [dbo].[FIN_JNL_LINE] NOT NULL,
			[NOTE_ID] [int] NOT NULL,
			[UPDATE_COUNT] [dbo].[UPDATE_CNT] NOT NULL,
			[LAST_CHANGE_USER_ID] [dbo].[LAST_CHANGE_USER_ID] NOT NULL,
			[LAST_CHANGE_DATETIME] [dbo].[LAST_CHANGE_DATETIME] NOT NULL,
			[LDG_TYPE] [dbo].[LEDGER_CODE] NOT NULL,
			PRIMARY KEY CLUSTERED 
			(
				[ACNT_CODE] ASC,
				[PERD] ASC,
				[TXN_DATETIME] ASC,
				[JNL_NUM] ASC,
				[JNL_LINE] ASC,
				[NOTE_ID] ASC
			)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
			) ON [PRIMARY]`
			_ , errB := conn.Exec(queryB)
			if errB != nil {
				defer conn.Close()
				log.Fatal("Perdon B!: ", errB.Error())
			}
			queryC := `CREATE TABLE [dbo].[`+BUNIT+`_OWN_CO](
			[OWN_CO_CODE] [dbo].[REF_MAINT_CODE] NOT NULL,
			[UPDATE_COUNT] [dbo].[UPDATE_CNT] NOT NULL,
			[LAST_CHANGE_USER_ID] [dbo].[LAST_CHANGE_USER_ID] NOT NULL,
			[LAST_CHANGE_DATETIME] [dbo].[LAST_CHANGE_DATETIME] NOT NULL,
			[PYMT_TERMS_GRP_CODE] [dbo].[REF_MAINT_CODE] NULL,
			[DESCR] [dbo].[DESCRIPTION_50] NULL,
			[S_HEAD] [dbo].[SHORT_HEADING] NULL,
			[LOOKUP] [dbo].[LOOKUP] NOT NULL,
			[PYMT_RCPT_MTHD] [dbo].[FLAG_PYMT_MTHD] NOT NULL,
			[PREF_PYMT_MTHD] [dbo].[FLAG_PYMT_MTHD] NOT NULL,
			[INV_TO_ADDR_CODE] [dbo].[REF_MAINT_CODE] NULL,
			[WEB_PAGE_ADDR] [dbo].[WEB_ADDRESS] NULL,
			[EMAIL] [dbo].[EMAIL] NULL,
			[NAME] [dbo].[NAME] NULL,
			PRIMARY KEY CLUSTERED 
			(
				[OWN_CO_CODE] ASC
			)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
			) ON [PRIMARY]`
			_ , errC := conn.Exec(queryC)
			if errC != nil {
				defer conn.Close()
				log.Fatal("Perdon C!: ", errC.Error())
			}
			queryD := `CREATE TABLE [dbo].[`+BUNIT+`_PSTG_DETAIL](
			[PSTG_HDR_ID] [int] NOT NULL,
			[LINE_NUM] [dbo].[NUM_INT] NOT NULL,
			[UPDATE_COUNT] [dbo].[UPDATE_CNT] NOT NULL,
			[LAST_CHANGE_USER_ID] [dbo].[LAST_CHANGE_USER_ID] NOT NULL,
			[LAST_CHANGE_DATETIME] [dbo].[LAST_CHANGE_DATETIME] NOT NULL,
			[ACNT_CODE] [dbo].[REF_MAINT_CODE] NOT NULL,
			[PERD] [dbo].[PERD] NOT NULL,
			[TXN_DATETIME] [datetime] NULL,
			[JNL_NUM] [dbo].[NUM_INT] NOT NULL,
			[JNL_LINE_NUM] [dbo].[NUM_INT] NOT NULL,
			[JNL_TYPE] [dbo].[CHAR_CODE_5] NOT NULL,
			[JNL_SRCE] [dbo].[CHAR_CODE_5] NOT NULL,
			[TXN_REF] [dbo].[TRANSACTION_REF_MIXED] NOT NULL,
			[DESCR] [dbo].[STRING_V255] NOT NULL,
			[AMT] [dbo].[NUM_18_3_ACC] NOT NULL,
			[DR_CR_IND] [dbo].[CHAR_CODE_1] NOT NULL,
			[CONV_CODE] [dbo].[CHAR_CODE_5] NOT NULL,
			[CONV_RATE] [dbo].[NUM_18_9] NOT NULL,
			[TXN_AMT] [dbo].[NUM_18_3_ACC] NOT NULL,
			[TXN_DEC_PL] [dbo].[CHAR_CODE_1] NOT NULL,
			[BASE_RATE] [dbo].[NUM_18_9] NOT NULL,
			[BASE_OPR] [dbo].[CHAR_CODE_1] NOT NULL,
			[CONV_OPR] [dbo].[CHAR_CODE_1] NOT NULL,
			[RPT_RATE] [dbo].[NUM_18_9] NOT NULL,
			[RPT_OPR] [dbo].[CHAR_CODE_1] NOT NULL,
			[RPT_AMT] [dbo].[NUM_18_3_ACC] NOT NULL,
			[MEMO_AMT] [dbo].[NUM_18_5_ACC_FIN] NOT NULL,
			[ALLOCN_IND] [dbo].[NUM_SMALLINT] NOT NULL,
			[ALLOCN_REF] [dbo].[NUM_INT] NOT NULL,
			[ALLOCN_DATETIME] [datetime] NULL,
			[ALLOCN_PERD] [dbo].[PERD] NOT NULL,
			[ALLOCN_IN_PROGRESS] [dbo].[NUM_SMALLINT] NOT NULL,
			[ENTRY_DATETIME] [datetime] NULL,
			[ENTRY_PERD] [dbo].[PERD] NOT NULL,
			[DUE_DATETIME] [datetime] NULL,
			[PSTG_DATETIME] [datetime] NULL,
			[ASSET_IND] [dbo].[NUM_SMALLINT] NOT NULL,
			[ASSET_CODE] [dbo].[CHAR_CODE_10] NOT NULL,
			[ASSET_SUB_CODE] [dbo].[CHAR_CODE_5] NOT NULL,
			[CLEARDOWN] [dbo].[CHAR_CODE_5] NOT NULL,
			[REVERSAL] [dbo].[NUM_SMALLINT] NOT NULL,
			[LOSS_GAIN] [dbo].[NUM_SMALLINT] NOT NULL,
			[ROUGH_FLAG] [dbo].[NUM_SMALLINT] NOT NULL,
			[IN_USE_FLAG] [dbo].[NUM_SMALLINT] NOT NULL,
			[EXCL_BAL] [dbo].[NUM_SMALLINT] NOT NULL,
			[ANL_CODE_T0] [dbo].[REF_MAINT_CODE] NOT NULL,
			[ANL_CODE_T1] [dbo].[REF_MAINT_CODE] NOT NULL,
			[ANL_CODE_T2] [dbo].[REF_MAINT_CODE] NOT NULL,
			[ANL_CODE_T3] [dbo].[REF_MAINT_CODE] NOT NULL,
			[ANL_CODE_T4] [dbo].[REF_MAINT_CODE] NOT NULL,
			[ANL_CODE_T5] [dbo].[REF_MAINT_CODE] NOT NULL,
			[ANL_CODE_T6] [dbo].[REF_MAINT_CODE] NOT NULL,
			[ANL_CODE_T7] [dbo].[REF_MAINT_CODE] NOT NULL,
			[ANL_CODE_T8] [dbo].[REF_MAINT_CODE] NOT NULL,
			[ANL_CODE_T9] [dbo].[REF_MAINT_CODE] NOT NULL,
			[HOLD_REF] [dbo].[NUM_INT] NOT NULL,
			[HOLD_OPR_CODE] [dbo].[OPERATOR_CODE] NOT NULL,
			[DOC_1_DATETIME] [datetime] NULL,
			[DOC_2_DATETIME] [datetime] NULL,
			[DOC_3_DATETIME] [datetime] NULL,
			[DOC_4_DATETIME] [datetime] NULL,
			[DOC_NUM_PRFX_1] [dbo].[CHAR_ALPHA_V15] NULL,
			[DOC_NUM_1] [dbo].[NUM_INT] NOT NULL,
			[DOC_NUM_PRFX_2] [dbo].[CHAR_ALPHA_V15] NULL,
			[DOC_NUM_2] [dbo].[NUM_INT] NOT NULL,
			[DOC_NUM_PRFX_3] [dbo].[CHAR_ALPHA_V15] NULL,
			[DOC_NUM_3] [dbo].[NUM_INT] NOT NULL,
			[DOC_NUM_PRFX_4] [dbo].[CHAR_ALPHA_V15] NULL,
			[DOC_NUM_4] [dbo].[NUM_INT] NOT NULL,
			[DISC_1_DATETIME] [datetime] NULL,
			[DISC_PCENT_1] [dbo].[PERCENTAGE_PM] NOT NULL,
			[DISC_2_DATETIME] [datetime] NULL,
			[DISC_PCENT_2] [dbo].[PERCENTAGE_PM] NOT NULL,
			[INTEREST_DATETIME] [datetime] NULL,
			[INTEREST_PCENT] [dbo].[PERCENTAGE_PM] NOT NULL,
			[LATE_PYMT_DATETIME] [datetime] NULL,
			[LATE_PYMT_PCENT] [dbo].[PERCENTAGE_PM] NOT NULL,
			[PYMT_REF] [dbo].[CHAR_ALPHA_30] NOT NULL,
			[BANK_CODE] [dbo].[BANK_SORT_CODE] NOT NULL,
			[SRCE_REF] [dbo].[CHAR_CODE_10] NOT NULL,
			[MODULE_CODE] [dbo].[CHAR_CODE_1] NOT NULL,
			[PYMT_TERMS_GRP_CODE] [dbo].[REF_MAINT_CODE] NULL,
			[STD_TEXT_CLASS_CODE] [dbo].[REF_MAINT_CODE] NOT NULL,
			[STD_TEXT_CODE] [dbo].[REF_MAINT_CODE] NOT NULL,
			[CONSUMED_BDGT_ID] [dbo].[NUM_INT] NULL,
			[CV4_CONV_CODE] [dbo].[CHAR_CODE_5] NOT NULL,
			[CV4_AMT] [dbo].[NUM_18_3_ACC] NOT NULL,
			[CV4_CONV_RATE] [dbo].[NUM_18_9] NOT NULL,
			[CV4_OPERATOR] [dbo].[CHAR_CODE_1] NOT NULL,
			[CV4_DP] [dbo].[CHAR_CODE_1] NOT NULL,
			[CV5_CONV_CODE] [dbo].[CHAR_CODE_5] NOT NULL,
			[CV5_AMT] [dbo].[NUM_18_3_ACC] NOT NULL,
			[CV5_CONV_RATE] [dbo].[NUM_18_9] NOT NULL,
			[CV5_OPERATOR] [dbo].[CHAR_CODE_1] NOT NULL,
			[CV5_DP] [dbo].[CHAR_CODE_1] NOT NULL,
			[LINK_REF_1] [dbo].[CHAR_ALPHA_V15] NULL,
			[LINK_REF_2] [dbo].[CHAR_ALPHA_V15] NULL,
			[LINK_REF_3] [dbo].[CHAR_ALPHA_V15] NULL,
			[PRINCIPAL_CODE_1] [dbo].[CHAR_ALPHA_V15] NULL,
			[PRINCIPAL_CODE_2] [dbo].[CHAR_ALPHA_V15] NULL,
			[PRINCIPAL_CODE_3] [dbo].[CHAR_ALPHA_V15] NULL,
			[PRINCIPAL_CODE_4] [dbo].[CHAR_ALPHA_V15] NULL,
			[PRINCIPAL_CODE_5] [dbo].[CHAR_ALPHA_V15] NULL,
			[PRINCIPAL_CODE_6] [dbo].[CHAR_ALPHA_V15] NULL,
			[PRINCIPAL_CODE_7] [dbo].[CHAR_ALPHA_V15] NULL,
			[PRINCIPAL_CODE_8] [dbo].[CHAR_ALPHA_V15] NULL,
			[PRINCIPAL_CODE_9] [dbo].[CHAR_ALPHA_V15] NULL,
			[PRINCIPAL_CODE_10] [dbo].[CHAR_ALPHA_V15] NULL,
			[PRINCIPAL_CODE_11] [dbo].[CHAR_ALPHA_V15] NULL,
			[PRINCIPAL_CODE_12] [dbo].[CHAR_ALPHA_V15] NULL,
			[PRINCIPAL_CODE_13] [dbo].[CHAR_ALPHA_V15] NULL,
			[PRINCIPAL_CODE_14] [dbo].[CHAR_ALPHA_V15] NULL,
			[PRINCIPAL_CODE_15] [dbo].[CHAR_ALPHA_V15] NULL,
			[PRINCIPAL_CODE_16] [dbo].[CHAR_ALPHA_V15] NULL,
			[PRINCIPAL_CODE_17] [dbo].[CHAR_ALPHA_V15] NULL,
			[PRINCIPAL_CODE_18] [dbo].[CHAR_ALPHA_V15] NULL,
			[PRINCIPAL_CODE_19] [dbo].[CHAR_ALPHA_V15] NULL,
			[PRINCIPAL_CODE_20] [dbo].[CHAR_ALPHA_V15] NULL,
			[ALLOCN_CODE] [dbo].[CHAR_CODE_5] NULL,
			[ALLOCN_STMNTS] [dbo].[NUM_SMALLINT] NULL,
			[ALLOCN_USER_ID] [dbo].[OPERATOR_CODE] NULL,
			[SPLIT_ORIG_LINE] [dbo].[NUM_INT] NOT NULL,
			[VAL_DATETIME] [datetime] NULL,
			[SIGNING_DETAILS] [dbo].[CHAR_ALPHA_V30] NULL,
			[INSTLMT_DATETIME] [datetime] NULL,
			[BINDER_STATUS] [dbo].[FIN_FLAG_BINDER_STATUS] NOT NULL,
			[AGREED_STATUS] [dbo].[FLAG_Y_N] NOT NULL,
			[SPLIT_LINK_REF] [dbo].[CHAR_ALPHA_V15] NULL,
			[PSTG_REF] [dbo].[CHAR_ALPHA_V15] NULL,
			[TRUE_RATED] [dbo].[FLAG_TRUE_RATED] NOT NULL,
			[HOLD_DATETIME] [datetime] NULL,
			[HOLD_TEXT] [dbo].[STRING_V30] NULL,
			[INSTLMT_NUM] [dbo].[NUM_SMALLINT] NULL,
			[SUPPLMNTRY_EXTSN] [dbo].[FLAG_Y_N] NOT NULL,
			[APRVLS_EXTSN] [dbo].[FLAG_Y_N] NOT NULL,
			[REVAL_LINK_REF] [dbo].[NUM_INT] NULL,
			[MAN_PAY_OVER] [dbo].[FLAG_Y_N] NOT NULL,
			[PYMT_STAMP] [dbo].[STRING_V10] NULL,
			[AUTHORISTN_IN_PROGRESS] [dbo].[FLAG_Y_N] NOT NULL,
			[SPLIT_IN_PROGRESS] [dbo].[FLAG_Y_N] NOT NULL,
			[VCHR_NUM] [dbo].[CHAR_ALPHA_V30] NULL,
			[ORIGINATOR_ID] [dbo].[OPERATOR_CODE] NULL,
			[ORIGINATED_DATETIME] [datetime] NULL,
			[JNL_CLASS_CODE] [dbo].[REF_MAINT_CODE] NULL,
			[ALLOC_ID] [dbo].[OPERATOR_CODE] NULL,
			[JNL_REVERSAL_TYPE] [dbo].[FLAG_JNL_REVERSAL_TYPE] NOT NULL,
			PRIMARY KEY CLUSTERED 
			(
				[PSTG_HDR_ID] ASC,
				[LINE_NUM] ASC
			)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
			) ON [PRIMARY]`
			_ , errD := conn.Exec(queryD)
			if errD != nil {
				defer conn.Close()
				log.Fatal("Perdon D!: ", errD.Error())
			}
			queryE := `CREATE TABLE [dbo].[`+BUNIT+`_SSRFMSC](
			[SUN_TB] [dbo].[CHAR_CODE_3] NOT NULL,
			[KEY_FIELDS] [dbo].[CHAR_ALPHA_61] NOT NULL,
			[LOOKUP] [dbo].[FIN_LOOKUP] NOT NULL,
			[LAST_CHANGE_DATETIME] [dbo].[LAST_CHANGE_DATETIME] NOT NULL,
			[DAG] [dbo].[DATA_ACCESS_GROUP] NULL,
			[UPDATE_COUNT] [dbo].[UPDATE_CNT] NOT NULL,
			[LAST_CHANGE_USER_ID] [dbo].[LAST_CHANGE_USER_ID] NOT NULL,
			[RECORD_STATUS] [dbo].[FIN_NUM_1] NOT NULL,
			[SUN_DATA] [dbo].[STRING_V927] NOT NULL,
			PRIMARY KEY CLUSTERED 
			(
				[SUN_TB] ASC,
				[KEY_FIELDS] ASC
			)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
			) ON [PRIMARY]`
			_ , errE := conn.Exec(queryE)
			if errE != nil {
				defer conn.Close()
				log.Fatal("Perdon E!: ", errE.Error())
			}
			queryF := `CREATE TABLE [dbo].[`+BUNIT+`_ACNT_ANL_CAT](
			[ANL_CAT_ID] [dbo].[ANALYSIS_DIMENSION] NOT NULL,
			[ACNT_CODE] [dbo].[REF_MAINT_CODE] NOT NULL,
			[UPDATE_COUNT] [dbo].[UPDATE_CNT] NOT NULL,
			[LAST_CHANGE_USER_ID] [dbo].[LAST_CHANGE_USER_ID] NOT NULL,
			[LAST_CHANGE_DATETIME] [dbo].[LAST_CHANGE_DATETIME] NOT NULL,
			[ANL_CODE] [dbo].[ANL_CODE_0] NOT NULL,
			PRIMARY KEY CLUSTERED 
			(
				[ANL_CAT_ID] ASC,
				[ACNT_CODE] ASC
			)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
			) ON [PRIMARY]`
			_ , errF := conn.Exec(queryF)
			if errF != nil {
				defer conn.Close()
				log.Fatal("Perdon F!: ", errF.Error())
			}
			queryG := `CREATE TABLE [dbo].[`+BUNIT+`_ANL_ENT_DEFN](
				[ANL_ENT_ID] [dbo].[TBL_NUM] NOT NULL,
				[ANL_CAT_ID] [dbo].[ANALYSIS_DIMENSION] NOT NULL,
				[UPDATE_COUNT] [dbo].[UPDATE_CNT] NOT NULL,
				[LAST_CHANGE_USER_ID] [dbo].[LAST_CHANGE_USER_ID] NOT NULL,
				[LAST_CHANGE_DATETIME] [dbo].[LAST_CHANGE_DATETIME] NOT NULL,
				[ENTRY_NUM] [dbo].[NUM_SMALLINT] NOT NULL,
				[VALIDATE_IND] [dbo].[FLAG_VALIDATION_IND] NOT NULL,
				[S_HEAD] [dbo].[SHORT_HEADING_UPPER] NOT NULL,
			PRIMARY KEY CLUSTERED 
			(
				[ANL_ENT_ID] ASC,
				[ANL_CAT_ID] ASC
			)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
			) ON [PRIMARY]`
			_ , errG := conn.Exec(queryG)
			if errG != nil {
				defer conn.Close()
				log.Fatal("Perdon G!: ", errG.Error())
			}
		}
		example := map[string]interface{}{ "success":1}
		c.Data["json"] = &example
		c.ServeJSON()
	}
}
func (c *GuardarBUNITController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		idUsuario := c.GetSession("idUsuario")
		BUNITNUEVO := c.GetString("BUNIT")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		fmt.Println(connString2)
		conn, errS := sql.Open("mssql", connString2)
		if errS != nil {
			log.Fatal("Perdon! Open connection failed:", errS.Error())
		}
		//["+beego.AppConfig.String("serverAux")+"].
		//var idU string
		idU := strconv.Itoa(idUsuario.(int))
		query := "UPDATE [SUNPLUSADV].[dbo].[users] set BUNIT = '"+BUNITNUEVO+"' WHERE idUsuario = "+idU+""
			fmt.Println(query)
		result , err2 := conn.Exec(query)
		if err2 != nil {
				defer conn.Close()
			log.Fatal("Perdon 5! : ", err2.Error())
		}
		c.SetSession("BUNIT", BUNITNUEVO)
		c.SessionRegenerateID()
		fmt.Println(result.RowsAffected())
		example := map[string]interface{}{ "success":1}
		c.Data["json"] = &example
		c.ServeJSON()
	}
}


func (c *GenerarDiarioEnPDFController) Get() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		diario := c.GetString("diario")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
		queryDiario := "SELECT DESCR FROM [SUNPLUSADV].[dbo].[DB_DEFN] WHERE DB_CODE = '"+BUNIT.(string)+"'"
		rowsDiario, _ := conn.Query(queryDiario)
		var bunitD string
		for rowsDiario.Next()  {
			rowsDiario.Scan(&bunitD)
		}
		var PERIOD int
		var JRNAL_SRCE string
		var TREFERENCE_GLOBAL string
		var ANAL_T0_GLOBAL string
		var JRNAL_TYPE string
		query := "SELECT TOP 1 PERIOD, JRNAL_SRCE, TREFERENCE, ANAL_T0, JRNAL_TYPE FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_A_SALFLDG] WHERE JRNAL_NO = "+diario+" order by JRNAL_LINE asc"
   		rowsAntes, err1 := conn.Query(query)
   		if err1 != nil {
			log.Fatal("Perdón! :", err1.Error())
		}
		if rowsAntes.Next()  {
			rowsAntes.Scan(&PERIOD, &JRNAL_SRCE, &TREFERENCE_GLOBAL, &ANAL_T0_GLOBAL, &JRNAL_TYPE)
		}
		pdf := gofpdf.New("L", "mm", "A4", "")
		tr := pdf.UnicodeTranslatorFromDescriptor("")
	    pdf.SetHeaderFunc(func() {
			pdf.Image("logo.jpg", 10, 3, 30, 0, false, "", 0, "")
 			pdf.SetFont("Arial", "", 18)
 			pdf.SetXY(40, 10)
 			pdf.Cell(0,0,tr(bunitD))
			//pdf.CellFormat(0,0,tr(bunitD),"",1,"C",false,0,"")	
			pdf.SetFont("Arial", "B", 12)			
			periodoActual := strconv.Itoa(PERIOD)		
			pdf.SetXY(40, 22)
			pdf.Cell(0,0,tr("Periodo contable: "+periodoActual))
			pdf.SetXY(40, 30)
			pdf.Cell(0, 0, tr("Tipo de diario: "+JRNAL_TYPE))
			pdf.SetXY(105, 22)
			pdf.Cell(0, 0, tr("Referencia de Transacción: "+TREFERENCE_GLOBAL))
			pdf.SetXY(105, 30)
			pdf.Cell(0, 0, "Creado por: "+JRNAL_SRCE)
			pdf.SetFont("Arial", "B", 16)
			pdf.SetXY(215, 22)
			pdf.Cell(0, 0, tr("Diario N°: "+diario))
			pdf.Ln(20)
		})
		pdf.SetFooterFunc(func() {
			pdf.SetY(-15)
			pdf.SetFont("Arial", "I", 8)
			pdf.CellFormat(0, 10, fmt.Sprintf("Page %d", pdf.PageNo()),
				"", 0, "C", false, 0, "")
		})
		//pdf.AddPage()
	    pdf.SetFont("Arial", "", 8)
	   

		/////////////




		pdf.AddPage()
		
		queryX := "SELECT c.TREFERENCE, c.ACCNT_CODE, SUBSTRING( a.DESCR,0,25) as DESCR , c.ANAL_T1, c.ANAL_T2, c.ANAL_T3, c.ANAL_T4, c.ANAL_T5, c.ANAL_T6, c.ANAL_T9, c.ANAL_T8, SUBSTRING( c.DESCRIPTN,0,35) as DESCRIPTN, c.AMOUNT, c.CONV_CODE, c.ANAL_T7, c.D_C, c.ALLOCATION, c.ANAL_T0, c.JRNAL_LINE FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_A_SALFLDG] c INNER JOIN [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ACNT] a on a.ACNT_CODE = c.ACCNT_CODE WHERE c.JRNAL_NO = "+diario+" order by c.JRNAL_LINE asc"
   		rows, _ := conn.Query(queryX)
   		var D_C string
		var DESCRIPTN string
		var descripcionCuenta string
		var tfww string
		var fondo string
		var funcion string
		var moneda string
		var restriccion string
		var orgId string
		var who string
		var detalle string
		var proyecto string
		var flag string	
		var anal_t0 string	
		var cuenta string
		var allocation string
		var treference string	
		var linea int
		var amount decimal.Dec
		//records := [][]string{}
		//var i int
		//i = 0
		dGlobales := 0.0
		cGlobales := 0.0
		dRef := 0.0
		cRef := 0.0
		y := 50.0
		x := 0.0
		tope := 180.0
		ac := accounting.Accounting{Symbol: "$", Precision: 2}
		TREFERENCE_ANTERIOR := "-"
		for rows.Next()  {
			rows.Scan(&treference, &cuenta, &descripcionCuenta, &tfww, &fondo, &funcion, &restriccion, &orgId, &who, &detalle, &proyecto, &DESCRIPTN, &amount, &moneda, &flag, &D_C, &allocation, &anal_t0, &linea)
			treference = strings.TrimSpace(treference)
			anal_t0 = strings.TrimSpace(anal_t0)
			cuenta = strings.TrimSpace(cuenta)
			descripcionCuenta = strings.TrimSpace(descripcionCuenta)
			descripcionCuenta = strings.Replace(descripcionCuenta, "á", "a", -1)
			descripcionCuenta = strings.Replace(descripcionCuenta, "é", "e", -1)
			descripcionCuenta = strings.Replace(descripcionCuenta, "í", "i", -1)
			descripcionCuenta = strings.Replace(descripcionCuenta, "ó", "o", -1)
			descripcionCuenta = strings.Replace(descripcionCuenta, "ú", "u", -1)
			descripcionCuenta = strings.Replace(descripcionCuenta, "ñ", "n", -1)
			tfww = strings.TrimSpace(tfww)
			fondo = strings.TrimSpace(fondo)
			funcion = strings.TrimSpace(funcion)
			restriccion = strings.TrimSpace(restriccion)
			orgId = strings.TrimSpace(orgId)
			who = strings.TrimSpace(who)
			detalle = strings.TrimSpace(detalle)
			proyecto = strings.TrimSpace(proyecto)
			DESCRIPTN = strings.TrimSpace(DESCRIPTN)
			DESCRIPTN = strings.Replace(DESCRIPTN, "á", "a", -1)
			DESCRIPTN = strings.Replace(DESCRIPTN, "é", "e", -1)
			DESCRIPTN = strings.Replace(DESCRIPTN, "í", "i", -1)
			DESCRIPTN = strings.Replace(DESCRIPTN, "ó", "o", -1)
			DESCRIPTN = strings.Replace(DESCRIPTN, "ú", "u", -1)
			DESCRIPTN = strings.Replace(DESCRIPTN, "ñ", "n", -1)
			moneda = strings.TrimSpace(moneda)
			flag = strings.TrimSpace(flag)		
				

			if Compare(TREFERENCE_ANTERIOR,treference) == 0 {

			} else {
				if Compare(TREFERENCE_ANTERIOR,"-")==0 {//primera vez, no hacer nada
				} else {
					pdf.SetFont("Arial", "", 8)
					pdf.SetXY(170, y)
					pdf.Cell(0, 0,  "Total por referencia:")
					pdf.SetFont("Arial", "", 10)
					pdf.SetXY(240, y)
					pdf.Cell(0, 0,  ac.FormatMoney(math.Abs(dRef)))
					pdf.SetXY(260, y)
					pdf.Cell(0, 0, ac.FormatMoney(math.Abs(cRef)))
					y = y + 5
					dRef = 0
					cRef = 0
				}

				if y < (tope-20.0) {
				} else {
					y = 50
					pdf.AddPage()
				}


				pdf.SetFont("Arial", "B", 12)
				pdf.SetXY(10, y)
				pdf.Cell(0, 0, anal_t0)
				y = y + 6 //6 x q 12 es la font! ¿?
				pdf.ClipPolygon([]gofpdf.PointType{{10, y -2.5}, {280, y-2.5},{280, y-2.5 }}, true)
				pdf.ClipEnd()
				
				pdf.SetFont("Arial", "", 10)
				pdf.SetXY(10, y)
				pdf.Cell(0, 0, "Ln")
				pdf.SetXY(22, y)
				pdf.Cell(0, 0, "Cuenta")
				pdf.SetXY(50, y)
				pdf.Cell(0, 0, "Nombre de la cuenta")
				pdf.SetXY(95, y)
				pdf.Cell(0, 0, tr("Descripción"))
				pdf.SetXY(170, y)
				pdf.Cell(0, 0, "Dimensiones")
				pdf.SetXY(240, y)
				pdf.Cell(0, 0, tr("Débito"))
				pdf.SetXY(260, y)
				pdf.Cell(0, 0, tr("Crédito"))	
				y = y + 5//5 x la font??
				pdf.ClipPolygon([]gofpdf.PointType{{10, y -2.5}, {280, y-2.5},{280, y-2.5 }}, true)
				pdf.ClipEnd()
				TREFERENCE_ANTERIOR = treference
				
			}
			amountPrima := math.Abs(amount.Float64())
			if Compare(D_C,"C")==0 {
				amountPrima = amountPrima*-1.0
				cGlobales += amountPrima
				cRef += amountPrima
			} else {
				dGlobales += amountPrima;
				dRef +=amountPrima;
			}
		//	amountString := fmt.Sprintf("%.2f", amountPrima)// strconv.FormatFloat(amountPrima, 'E', 2, 64)
		
			lineaActual := strconv.Itoa(linea)
		
			pdf.SetXY(10, y)
			pdf.Cell(0, 0, lineaActual)
			pdf.SetXY(22, y)
			pdf.Cell(0, 0, tr(cuenta))
			pdf.SetXY(50, y)
			pdf.Cell(0, 0, tr(descripcionCuenta))
			pdf.SetXY(95, y)
			pdf.Cell(0, 0, tr(DESCRIPTN))
			pdf.SetXY(170, y)
			pdf.SetFont("Arial", "", 7)
			pdf.Cell(0, 0, tr(tfww)+"-"+tr(fondo)+"-"+tr(funcion)+"-"+tr(restriccion)+"-"+tr(orgId)+"-"+tr(who)+"-"+tr(flag)+"-"+tr(proyecto)+"-"+tr(detalle))
			pdf.SetFont("Arial", "", 10)
			if Compare(D_C,"C")==0 {
				x = 260
			} else {
				x = 240
			}
			pdf.SetXY(x, y)
			pdf.Cell(0, 0, ac.FormatMoney(math.Abs(amountPrima)))			
			//row1 := []string{treference,cuenta,descripcionCuenta,tfww,fondo,funcion,restriccion,orgId,who,detalle,proyecto,DESCRIPTN,amountString,moneda,flag,allocation}
			//records = append(records, row1)
			if y < tope{
				y = y + 5
			} else {
				y = 50
				pdf.AddPage()
		
			}
		}
		pdf.SetFont("Arial", "", 8)
		pdf.SetXY(170, y)
		pdf.Cell(0, 0,  "Total por referencia:")
		pdf.SetFont("Arial", "", 10)
		pdf.SetXY(240, y)
		pdf.Cell(0, 0,  ac.FormatMoney(math.Abs(dRef)))
		pdf.SetXY(260, y)
		pdf.Cell(0, 0, ac.FormatMoney(math.Abs(cRef)))
		y = y + 5
		pdf.ClipPolygon([]gofpdf.PointType{{10, y -2.5}, {280, y-2.5},{280, y-2.5 }}, true)
		pdf.ClipEnd()
		pdf.SetFont("Arial", "", 8)
		pdf.SetXY(170, y)
		pdf.Cell(0, 0,  "Total del diario:")
		pdf.SetFont("Arial", "", 10)
		pdf.SetXY(240, y)
		pdf.Cell(0, 0,  ac.FormatMoney(math.Abs(dGlobales)))
		pdf.SetXY(260, y)
		pdf.Cell(0, 0, ac.FormatMoney(math.Abs(cGlobales)))


 		var b bytes.Buffer
	    w := bufio.NewWriter(&b)
	    pdf.Output(w)
	    pdf.Close()
	    w.Flush()
	    c.Ctx.Output.ContentType("application/pdf")
	    c.Ctx.Output.Body(b.Bytes())
	}
}

func (c *GenerarDiarioEnExcelController) Get() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		diario := c.GetString("diario")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
		query := "SELECT c.TREFERENCE, c.ACCNT_CODE, a.DESCR, c.ANAL_T1, c.ANAL_T2, c.ANAL_T3, c.ANAL_T4, c.ANAL_T5, c.ANAL_T6, c.ANAL_T9, c.ANAL_T8,c.DESCRIPTN, c.AMOUNT, c.CONV_CODE, c.ANAL_T7, c.D_C, c.ALLOCATION FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_A_SALFLDG] c INNER JOIN [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ACNT] a on a.ACNT_CODE = c.ACCNT_CODE WHERE c.JRNAL_NO = "+diario+" order by c.JRNAL_LINE asc"
   		rows, err := conn.Query(query)
   		var D_C string
		var DESCRIPTN string
		var descripcionCuenta string
		var tfww string
		var fondo string
		var funcion string
		var moneda string
		var restriccion string
		var orgId string
		var who string
		var detalle string
		var proyecto string
		var flag string	
		var cuenta string
		var allocation string
		var treference string	
		var amount decimal.Dec
		records := [][]string{}
		var i int
		i = 0
		for rows.Next()  {
			rows.Scan(&treference, &cuenta, &descripcionCuenta, &tfww, &fondo, &funcion, &restriccion, &orgId, &who, &detalle, &proyecto, &DESCRIPTN, &amount, &moneda, &flag, &D_C, &allocation)
			treference = strings.TrimSpace(treference)
			cuenta = strings.TrimSpace(cuenta)
			descripcionCuenta = strings.TrimSpace(descripcionCuenta)
			descripcionCuenta = strings.Replace(descripcionCuenta, "á", "a", -1)
			descripcionCuenta = strings.Replace(descripcionCuenta, "é", "e", -1)
			descripcionCuenta = strings.Replace(descripcionCuenta, "í", "i", -1)
			descripcionCuenta = strings.Replace(descripcionCuenta, "ó", "o", -1)
			descripcionCuenta = strings.Replace(descripcionCuenta, "ú", "u", -1)
			descripcionCuenta = strings.Replace(descripcionCuenta, "ñ", "n", -1)
			tfww = strings.TrimSpace(tfww)
			fondo = strings.TrimSpace(fondo)
			funcion = strings.TrimSpace(funcion)
			restriccion = strings.TrimSpace(restriccion)
			orgId = strings.TrimSpace(orgId)
			who = strings.TrimSpace(who)
			detalle = strings.TrimSpace(detalle)
			proyecto = strings.TrimSpace(proyecto)
			DESCRIPTN = strings.TrimSpace(DESCRIPTN)
			DESCRIPTN = strings.Replace(DESCRIPTN, "á", "a", -1)
			DESCRIPTN = strings.Replace(DESCRIPTN, "é", "e", -1)
			DESCRIPTN = strings.Replace(DESCRIPTN, "í", "i", -1)
			DESCRIPTN = strings.Replace(DESCRIPTN, "ó", "o", -1)
			DESCRIPTN = strings.Replace(DESCRIPTN, "ú", "u", -1)
			DESCRIPTN = strings.Replace(DESCRIPTN, "ñ", "n", -1)
			moneda = strings.TrimSpace(moneda)
			flag = strings.TrimSpace(flag)		
			amountPrima := math.Abs(amount.Float64())
			if Compare(D_C,"C")==0 {
				amountPrima = amountPrima*-1.0
			}
			amountString := fmt.Sprintf("%.2f", amountPrima)// strconv.FormatFloat(amountPrima, 'E', 2, 64)
			row1 := []string{treference,cuenta,descripcionCuenta,tfww,fondo,funcion,restriccion,orgId,who,detalle,proyecto,DESCRIPTN,amountString,moneda,flag,allocation}
			records = append(records, row1)
			i = i + 1
		}
 		c.Ctx.Output.ContentType("text/csv")
 		c.Ctx.Output.Header("Content-Disposition", "attachment; filename=diario"+diario+".csv")
		writer := csv.NewWriter(c.Ctx.ResponseWriter)
		for _, record := range records {
		    err := writer.Write(record)
		    if err != nil {
		        fmt.Println("Error:", err)
		        return
		    }
		}
		writer.Flush()
	}
}
func (c *GenerarEstadosDeCuentaController) Get() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		acnt := c.GetString("cuenta")
		descr := c.GetString("descr")
		busca := c.GetString("busca")

		BUNIT := c.GetSession("BUNIT")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}

		if Compare(busca,"1")==0 {

			queryCuenta := "SELECT ACNT_CODE FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ACNT] WHERE UPPER(DESCR) like '%"+acnt+"%'"
			fmt.Println(queryCuenta)
		 	rowsCuenta, _ := conn.Query(queryCuenta)
			for rowsCuenta.Next()  {
				rowsCuenta.Scan(&acnt)
			}
			fmt.Println(acnt)
		}



		delPeriodo := c.GetString("delPeriodo")
		alPeriodo := c.GetString("alPeriodo")
		var ANAL_T0 string
		var ANAL_T1 string
		var ANAL_T2 string
		var ANAL_T3 string
		var ANAL_T4 string
		var ANAL_T5 string
		var ANAL_T6 string
		var ANAL_T7 string
		var ANAL_T8 string
		var ANAL_T9 string
		ANAL_T0_START := c.GetString("d0s")
		ANAL_T0_END := c.GetString("d0e")

		ANAL_T1_START := c.GetString("d1s")
		ANAL_T1_END := c.GetString("d1e")
		
		ANAL_T2_START := c.GetString("d2s")
		ANAL_T2_END := c.GetString("d2e")
		
		ANAL_T3_START := c.GetString("d3s")
		ANAL_T3_END := c.GetString("d3e")
		
		ANAL_T4_START := c.GetString("d4s")
		ANAL_T4_END := c.GetString("d4e")
		
		ANAL_T5_START := c.GetString("d5s")
		ANAL_T5_END := c.GetString("d5e")
		
		ANAL_T6_START := c.GetString("d6s")
		ANAL_T6_END := c.GetString("d6e")
		
		ANAL_T7_START := c.GetString("d7s")
		ANAL_T7_END := c.GetString("d7e")
		
		ANAL_T8_START := c.GetString("d8s")
		ANAL_T8_END := c.GetString("d8e")
		
		ANAL_T9_START := c.GetString("d9s")
		ANAL_T9_END := c.GetString("d9e")
		
		
		

		queryDireccion := "SELECT ADDR_LINE_1,ADDR_LINE_2,ADDR_LINE_3,ADDR_LINE_4,ADDR_LINE_5 FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ADDR] WHERE ADDR_CODE = '0000000000'"
		rowsDireccion, err := conn.Query(queryDireccion)
		var addr1 string
		var addr2 string
		var addr3 string
		var addr4 string
		var addr5 string
		for rowsDireccion.Next()  {
			rowsDireccion.Scan(&addr1,&addr2,&addr3,&addr4,&addr5)
		}
		currentTime := int64(time.Now().Unix())
		tm := time.Unix(currentTime, 0)
		dateString := tm.String() 
		substring := string(dateString[0:10])
		pdf := gofpdf.New("P", "mm", "A4", "")
		tr := pdf.UnicodeTranslatorFromDescriptor("")
	    pdf.SetHeaderFunc(func() {
			pdf.Image("logo.jpg", 10, 3, 30, 0, false, "", 0, "")
 			pdf.SetFont("Arial", "", 16)
 			if  pdf.PageNo() == 1 {
 				pdf.SetXY(0, 8)
			} else {
				pdf.SetXY(0, 18)
			}
			pdf.CellFormat(0,0,tr(addr1),"",1,"C",false,0,"")
			
			pdf.SetFont("Arial", "", 10)
			if  pdf.PageNo() == 1 {
 				pdf.SetXY(40, 28)
			} else {
				pdf.SetXY(140, 12)
			}
			pdf.Cell(0,0,tr(descr))
			pdf.SetFont("Arial", "", 12)
			if  pdf.PageNo() == 1 {
				pdf.SetXY(0, 14)
				pdf.CellFormat(0,0,tr(addr2),"",1,"C",false,0,"")
				pdf.SetXY(0, 18)
				pdf.CellFormat(0,0,tr(addr3),"",1,"C",false,0,"")
				pdf.SetXY(0, 22)
				pdf.CellFormat(0,0,tr(addr4),"",1,"C",false,0,"")
				pdf.SetXY(0, 26)
				pdf.CellFormat(0,0,tr(addr5),"",1,"C",false,0,"")
				pdf.SetFont("Arial", "", 10)
				pdf.SetXY(0, 18)
				pdf.CellFormat(0,0,tr("Fecha de generación: "+substring),"",1,"R",false,0,"")
				pdf.SetFont("Arial", "", 10)
				pdf.SetXY(0, 26)
				pdf.CellFormat(0,0,tr("Código de cuenta: "+acnt),"",1,"R",false,0,"")
				pdf.SetFont("Arial", "", 10)
				pdf.SetXY(0, 22)
				pdf.CellFormat(0,0,tr("Periodo final: "+alPeriodo),"",1,"R",false,0,"")
			} else {
				pdf.SetFont("Arial", "", 10)
				pdf.SetXY(140, 7)
				pdf.Cell(0,0,tr(acnt))
			}
			//pdf.SetY(5)
			pdf.SetFont("Arial", "B", 9)
			pdf.SetXY(10, 40)
			pdf.Cell(0, 0, "Fecha")
			pdf.SetXY(30, 40)
			pdf.Cell(0, 0, "Poliza")
			pdf.SetXY(45, 40)
			pdf.Cell(0, 0, "Reference 2")
			pdf.SetXY(70, 40)
			pdf.Cell(0, 0, "Transaction")
			pdf.SetXY(95, 40)
			pdf.Cell(0, 0, "Descripcion")
			pdf.SetXY(145, 40)
			pdf.Cell(0, 0, "Debito")
			pdf.SetXY(165, 40)
			pdf.Cell(0, 0, "Credito")
			pdf.SetXY(187, 40)
			pdf.Cell(0, 0, "Balanza")
			//pdf.CellFormat(30, 10, "Title", "1", 0, "C", false, 0, "")
			pdf.Ln(20)
		})
		pdf.SetFooterFunc(func() {
			pdf.SetY(-15)
			pdf.SetFont("Arial", "I", 8)
			///{nb}
			pdf.CellFormat(0, 10, fmt.Sprintf("Page %d", pdf.PageNo()),
				"", 0, "C", false, 0, "")
		})
		//pdf.AddPage()
	    pdf.SetFont("Arial", "", 8)
	   

		
		

//cambiar
		queryAnterior := "SELECT ISNULL( SUM(AMOUNT),0)  as AMOUNT FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_A_SALFLDG] WHERE ACCNT_CODE = '"+acnt+"' AND PERIOD < "+delPeriodo+" AND ALLOCATION != 'C' "
		
	
		
		query := "SELECT D_C, SUBSTRING( DESCRIPTN,0,29) as DESCRIPTN, JRNAL_NO, JRNAL_LINE, ANAL_T0, PERIOD, TRANS_DATETIME, ALLOCATION, TREFERENCE, AMOUNT FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_A_SALFLDG] WHERE ACCNT_CODE = '"+acnt+"' AND PERIOD >= "+delPeriodo+" AND PERIOD <= "+alPeriodo+" AND ALLOCATION != 'C' "
		if Compare(ANAL_T0_START,"")!= 0  {
			queryDimension := "SELECT DISTINCT ANAL_T0 FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_A_SALFLDG] WHERE ANAL_T0 >= '"+ANAL_T0_START+"' AND ANAL_T0 <= '"+ANAL_T0_END+"' order by ANAL_T0 asc"
			rowsDimension, _ := conn.Query(queryDimension)
			primeraVez := true
			auxS := " AND ANAL_T0 IN ("
			for rowsDimension.Next()  {
				rowsDimension.Scan(&ANAL_T0)
				if primeraVez {
					primeraVez = false
					auxS=auxS+"'"+strings.TrimSpace(ANAL_T0)+"'"
				} else {
					auxS=auxS+",'"+strings.TrimSpace(ANAL_T0)+"'"
				}
			}
			auxS = auxS + ") "
			query = query + auxS
			queryAnterior = queryAnterior + auxS
		}


		if Compare(ANAL_T1_START,"")!= 0  {
			queryDimension := "SELECT DISTINCT ANAL_T1 FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_A_SALFLDG] WHERE ANAL_T1 >= '"+ANAL_T1_START+"' AND ANAL_T1 <= '"+ANAL_T1_END+"' order by ANAL_T1 asc"
			rowsDimension, _ := conn.Query(queryDimension)
			primeraVez := true
			auxS := " AND ANAL_T1 IN ("
			for rowsDimension.Next()  {
				rowsDimension.Scan(&ANAL_T1)
				if primeraVez {
					primeraVez = false
					auxS=auxS+"'"+strings.TrimSpace(ANAL_T1)+"'"
				} else {
					auxS=auxS+",'"+strings.TrimSpace(ANAL_T1)+"'"
				}
			}
			auxS = auxS + ") "
			query = query + auxS
			queryAnterior = queryAnterior + auxS
		}

		if Compare(ANAL_T2_START,"")!= 0  {
			queryDimension := "SELECT DISTINCT ANAL_T2 FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_A_SALFLDG] WHERE ANAL_T2 >= '"+ANAL_T2_START+"' AND ANAL_T2 <= '"+ANAL_T2_END+"' order by ANAL_T2 asc"
			rowsDimension, _ := conn.Query(queryDimension)
			primeraVez := true
			auxS := " AND ANAL_T2 IN ("
			for rowsDimension.Next()  {
				rowsDimension.Scan(&ANAL_T2)
				if primeraVez {
					primeraVez = false
					auxS=auxS+"'"+strings.TrimSpace(ANAL_T2)+"'"
				} else {
					auxS=auxS+",'"+strings.TrimSpace(ANAL_T2)+"'"
				}
			}
			auxS = auxS + ") "
			query = query + auxS
			queryAnterior = queryAnterior + auxS
		}

		if Compare(ANAL_T3_START,"")!= 0  {
			queryDimension := "SELECT DISTINCT ANAL_T3 FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_A_SALFLDG] WHERE ANAL_T3 >= '"+ANAL_T3_START+"' AND ANAL_T3 <= '"+ANAL_T3_END+"' order by ANAL_T3 asc"
			rowsDimension, _ := conn.Query(queryDimension)
			primeraVez := true
			auxS := " AND ANAL_T3 IN ("
			for rowsDimension.Next()  {
				rowsDimension.Scan(&ANAL_T3)
				if primeraVez {
					primeraVez = false
					auxS=auxS+"'"+strings.TrimSpace(ANAL_T3)+"'"
				} else {
					auxS=auxS+",'"+strings.TrimSpace(ANAL_T3)+"'"
				}
			}
			auxS = auxS + ") "
			query = query + auxS
			queryAnterior = queryAnterior + auxS
		}

		if Compare(ANAL_T4_START,"")!= 0  {
			queryDimension := "SELECT DISTINCT ANAL_T4 FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_A_SALFLDG] WHERE ANAL_T4 >= '"+ANAL_T4_START+"' AND ANAL_T4 <= '"+ANAL_T4_END+"' order by ANAL_T4 asc"
			rowsDimension, _ := conn.Query(queryDimension)
			primeraVez := true
			auxS := " AND ANAL_T4 IN ("
			for rowsDimension.Next()  {
				rowsDimension.Scan(&ANAL_T4)
				if primeraVez {
					primeraVez = false
					auxS=auxS+"'"+strings.TrimSpace(ANAL_T4)+"'"
				} else {
					auxS=auxS+",'"+strings.TrimSpace(ANAL_T4)+"'"
				}
			}
			auxS = auxS + ") "
			query = query + auxS
			queryAnterior = queryAnterior + auxS
		}

		if Compare(ANAL_T5_START,"")!= 0  {
			queryDimension := "SELECT DISTINCT ANAL_T5 FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_A_SALFLDG] WHERE ANAL_T5 >= '"+ANAL_T5_START+"' AND ANAL_T5 <= '"+ANAL_T5_END+"' order by ANAL_T5 asc"
			rowsDimension, _ := conn.Query(queryDimension)
			primeraVez := true
			auxS := " AND ANAL_T5 IN ("
			for rowsDimension.Next()  {
				rowsDimension.Scan(&ANAL_T5)
				if primeraVez {
					primeraVez = false
					auxS=auxS+"'"+strings.TrimSpace(ANAL_T5)+"'"
				} else {
					auxS=auxS+",'"+strings.TrimSpace(ANAL_T5)+"'"
				}
			}
			auxS = auxS + ") "
			query = query + auxS
			queryAnterior = queryAnterior + auxS
		}

		if Compare(ANAL_T6_START,"")!= 0  {
			queryDimension := "SELECT DISTINCT ANAL_T6 FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_A_SALFLDG] WHERE ANAL_T6 >= '"+ANAL_T6_START+"' AND ANAL_T6 <= '"+ANAL_T6_END+"' order by ANAL_T6 asc"
			rowsDimension, _ := conn.Query(queryDimension)
			primeraVez := true
			auxS := " AND ANAL_T6 IN ("
			for rowsDimension.Next()  {
				rowsDimension.Scan(&ANAL_T6)
				if primeraVez {
					primeraVez = false
					auxS=auxS+"'"+strings.TrimSpace(ANAL_T6)+"'"
				} else {
					auxS=auxS+",'"+strings.TrimSpace(ANAL_T6)+"'"
				}
			}
			auxS = auxS + ") "
			query = query + auxS
			queryAnterior = queryAnterior + auxS
		}

		if Compare(ANAL_T7_START,"")!= 0  {
			queryDimension := "SELECT DISTINCT ANAL_T7 FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_A_SALFLDG] WHERE ANAL_T7 >= '"+ANAL_T7_START+"' AND ANAL_T7 <= '"+ANAL_T7_END+"' order by ANAL_T7 asc"
			rowsDimension, _ := conn.Query(queryDimension)
			primeraVez := true
			auxS := " AND ANAL_T7 IN ("
			for rowsDimension.Next()  {
				rowsDimension.Scan(&ANAL_T7)
				if primeraVez {
					primeraVez = false
					auxS=auxS+"'"+strings.TrimSpace(ANAL_T7)+"'"
				} else {
					auxS=auxS+",'"+strings.TrimSpace(ANAL_T7)+"'"
				}
			}
			auxS = auxS + ") "
			query = query + auxS
			queryAnterior = queryAnterior + auxS
		}

		if Compare(ANAL_T8_START,"")!= 0  {
			queryDimension := "SELECT DISTINCT ANAL_T8 FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_A_SALFLDG] WHERE ANAL_T8 >= '"+ANAL_T8_START+"' AND ANAL_T8 <= '"+ANAL_T8_END+"' order by ANAL_T8 asc"
			rowsDimension, _ := conn.Query(queryDimension)
			primeraVez := true
			auxS := " AND ANAL_T8 IN ("
			for rowsDimension.Next()  {
				rowsDimension.Scan(&ANAL_T8)
				if primeraVez {
					primeraVez = false
					auxS=auxS+"'"+strings.TrimSpace(ANAL_T8)+"'"
				} else {
					auxS=auxS+",'"+strings.TrimSpace(ANAL_T8)+"'"
				}
			}
			auxS = auxS + ") "
			query = query + auxS
			queryAnterior = queryAnterior + auxS
		}

		if Compare(ANAL_T9_START,"")!= 0  {
			queryDimension := "SELECT DISTINCT ANAL_T9 FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_A_SALFLDG] WHERE ANAL_T9 >= '"+ANAL_T9_START+"' AND ANAL_T9 <= '"+ANAL_T9_END+"' order by ANAL_T9 asc"
			rowsDimension, _ := conn.Query(queryDimension)
			primeraVez := true
			auxS := " AND ANAL_T9 IN ("
			for rowsDimension.Next()  {
				rowsDimension.Scan(&ANAL_T9)
				if primeraVez {
					primeraVez = false
					auxS=auxS+"'"+strings.TrimSpace(ANAL_T9)+"'"
				} else {
					auxS=auxS+",'"+strings.TrimSpace(ANAL_T9)+"'"
				}
			}
			auxS = auxS + ") "
			query = query + auxS
			queryAnterior = queryAnterior + auxS
		}
		var amountInical decimal.Dec
		rowsInicial, err := conn.Query(queryAnterior)
		var acumulador float64
		acumulador = 0.0
		for rowsInicial.Next()  {
			rowsInicial.Scan(&amountInical)
			acumulador = amountInical.Float64()*-1.0
		}
		query = query+" order by PERIOD, JRNAL_NO asc, JRNAL_LINE asc, TRANS_DATETIME asc"
		ac := accounting.Accounting{Symbol: "$", Precision: 2}
   		rows, err := conn.Query(query)
		var D_C string
		var DESCRIPTN string
		var diario int
		var linea int
		var ref2 string
		var PERIOD int
		var trans time.Time
		var allocation string
		var treference string	
		var amount decimal.Dec
		err = nil
		
		var periodoActual int
		var contador int
		periodoActual,_ = strconv.Atoi(delPeriodo)
		fmt.Println(periodoActual)
		contador = 0

		pdf.AddPage()
		//var nuevoNumero = 0
		yPrima := 45.0
		//var totalDelPeriodo = 0
	 	for rows.Next()  {

			rows.Scan(&D_C, &DESCRIPTN, &diario, &linea, &ref2, &PERIOD, &trans, &allocation, &treference,&amount)
			contador = contador + 1
			//nuevoNumero := contador%46
			//yPrima := float64((nuevoNumero*5.0)+45.0)
			if PERIOD != periodoActual {
				pdf.ClipPolygon([]gofpdf.PointType{{10, yPrima -2.5}, {200, yPrima-2.5},{200, yPrima-2.5 }}, true)
				pdf.ClipEnd()
				pdf.SetXY(10, yPrima)
				//if contador != 1 {
					checar := ac.FormatMoney(acumulador)
					if Compare(checar,"$-0.00") == 0 {
						acumulador = 0.00
					}
					
					pdf.CellFormat(0,0,tr(""+strconv.Itoa(periodoActual)+ " Total de Periodo: "+ ac.FormatMoney(acumulador)),"",1,"C",false,0,"")
					if yPrima < 270.0 {
						yPrima = yPrima + 5.0
					} else {
						yPrima = 45.0
						pdf.AddPage()
					}
					/*if nuevoNumero < 46 {
						nuevoNumero = nuevoNumero + 1	
					} else {
						nuevoNumero = 0
					}
					yPrima = float64((nuevoNumero*5.0)+45.0)*/
				//}
				periodoActual = PERIOD
			}

		/*	if contador%46==0{
				pdf.AddPage()
			}*/
			pdf.SetXY(10, yPrima)
			mes := int(trans.Month())
			dia :=  trans.Day()
			if mes > 9 && dia > 9 {
				pdf.Cell(0, 0, fmt.Sprintf("%d-%d-%d", trans.Year(), mes, dia))
			} else {
				if mes > 9 && dia < 10 {
					pdf.Cell(0, 0, fmt.Sprintf("%d-%d-0%d", trans.Year(), mes, dia))
				} else {
					if mes < 10 && dia > 9 {
						pdf.Cell(0, 0, fmt.Sprintf("%d-0%d-%d", trans.Year(), mes, dia))
					} else {
						pdf.Cell(0, 0, fmt.Sprintf("%d-0%d-0%d", trans.Year(), mes, dia))	
					}
				}
			}
			pdf.SetXY(30, yPrima)
			debito := 0.0
			credito := 0.0
			pdf.Cell(0, 0, strconv.Itoa(diario)+"-"+strconv.Itoa(linea))
			pdf.SetXY(45, yPrima)
			pdf.Cell(0, 0, ref2)
		
			amountPrima := math.Abs(amount.Float64())
			if Compare(D_C,"D")==0{
				debito:=amountPrima	
				acumulador = acumulador + amountPrima
				pdf.SetXY(145, yPrima)
				pdf.Cell(0, 0, ac.FormatMoney(debito))
				pdf.SetXY(165, yPrima)
				pdf.Cell(0, 0, ac.FormatMoney(credito))
			}else{
				credito:=amountPrima
				acumulador = acumulador - amountPrima
				pdf.SetXY(145, yPrima)
				pdf.Cell(0, 0, ac.FormatMoney(debito))
				pdf.SetXY(165, yPrima)
				pdf.Cell(0, 0, ac.FormatMoney(credito))
			}
			//totalDelPeriodo = totalDelPeriodo + amountPrima
			pdf.SetXY(70, yPrima)
			pdf.Cell(0, 0, treference)
			pdf.SetXY(95, yPrima)
			pdf.Cell(0, 0, DESCRIPTN)
			pdf.SetXY(187, yPrima)
			pdf.Cell(0, 0, ac.FormatMoney(acumulador))

			if yPrima < 270.0 {
				yPrima = yPrima + 5.0
			} else {
				yPrima = 45.0
				pdf.AddPage()
			}

			

		}
		var b bytes.Buffer
	    w := bufio.NewWriter(&b)
	    pdf.Output(w)
	    pdf.Close()
	    w.Flush()
	    c.Ctx.Output.ContentType("application/pdf")
	    c.Ctx.Output.Body(b.Bytes())
    }
}




func (c *ObtieneDimensionController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		numero := c.GetString("numero")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
		query := "SELECT ANL_CAT_ID, ENTRY_NUM FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_ENT_DEFN] WHERE ANL_ENT_ID = "+numero+" order by ENTRY_NUM asc"
		rows, err1 := conn.Query(query)
		if err1 != nil {
			log.Fatal("que paso:", err.Error())
		}
		var ANL_CAT_ID string
		//var S_HEAD string
		var ENTRY_NUM int64
		//err = nil
		models.ClearDimensionA()
		var dimen models.DimensionAsignada
        //_ = cuenta
     	for rows.Next()  {
			rows.Scan(&ANL_CAT_ID, &ENTRY_NUM)
			dimen = models.DimensionAsignada{ANL_CAT_ID,ENTRY_NUM}
			models.AddDimensionA(dimen, ANL_CAT_ID)
		}
		example := map[string]interface{}{ "dimensiones": models.GetAllDimensionA() }
		c.Data["json"] = &example
		c.ServeJSON()
	}
}
func (c *ClasificacionDimensionController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
		query := "SELECT ANL_CAT_ID, S_HEAD,DESCR FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_CAT] order by ANL_CAT_ID asc"
		rows, err := conn.Query(query)
		var ANL_CAT_ID string
		var S_HEAD string
		var DESCR string
		//err = nil
		var dimen models.Dimensiones
        //_ = cuenta
     	for rows.Next()  {
			rows.Scan(&ANL_CAT_ID, &S_HEAD, &DESCR)
			dimen = models.Dimensiones{ANL_CAT_ID,S_HEAD,DESCR}
			models.AddDimension(dimen, ANL_CAT_ID)
		}
		example := map[string]interface{}{ "dimensiones": models.GetAllDimensiones() }
		c.Data["json"] = &example
		c.ServeJSON()
	}
}
func (c *DimensionesController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
		query := "SELECT c.S_HEAD, a.ANL_CAT_ID, a.ANL_CODE, a.NAME, a.STATUS, a.PROHIBIT_POSTING FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_CODE] a INNER JOIN [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_CAT] c on c.ANL_CAT_ID = a.ANL_CAT_ID"
		rows, err := conn.Query(query)
		var ANL_CODE string
		var ANL_CAT_ID string
		var S_HEAD string
		var NAME string
		var STATUS int64
		var PROHIBIT_POSTING int64
		//err = nil
		var dimen models.Dimension
        //_ = cuenta
     	for rows.Next()  {
			rows.Scan(&S_HEAD, &ANL_CAT_ID, &ANL_CODE, &NAME, &STATUS, &PROHIBIT_POSTING)
			dimen = models.Dimension{ANL_CAT_ID,S_HEAD,ANL_CODE, NAME, STATUS, PROHIBIT_POSTING}
			models.AddDimensionU(dimen, ANL_CODE)
		}
		example := map[string]interface{}{ "dimensiones": models.GetAllDimension() }
		c.Data["json"] = &example
		c.ServeJSON()
	}
}
type Cedula struct {
	cedula string
	idCedula int64
}


func (c *GenerarCedulaPorLineaController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		idLinea := c.GetString("idLinea")
		delPeriodo := c.GetString("delPeriodo")
		alPeriodo := c.GetString("alPeriodo")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
		query := "SELECT cc.nombre, cl.cuenta, cl.D_C_Tipo, cl.ANAL_T0, cl.ANAL_T1, cl.ANAL_T2, cl.ANAL_T3, cl.ANAL_T4, cl.ANAL_T5, cl.ANAL_T6, cl.ANAL_T7, cl.ANAL_T8, cl.ANAL_T9 FROM [SUNPLUSADV].[dbo].[CedulasLinea] cl INNER JOIN [SUNPLUSADV].[dbo].[CedulasConceptos] cc on cc.idConcepto = cl.idConcepto WHERE cl.idLinea = "+idLinea+" order by cl.idLinea asc"
		rows, err := conn.Query(query)
		var nombre string
		var cuenta string
		var D_C_Tipo int64
		var ANAL_T0 string
		var ANAL_T1 string
		var ANAL_T2 string
		var ANAL_T3 string
		var ANAL_T4 string
		var ANAL_T5 string
		var ANAL_T6 string
		var ANAL_T7 string
		var ANAL_T8 string
		var ANAL_T9 string
		D_C := ""
		models.ClearCedulasConceptosLineas()
		for rows.Next()  {
			rows.Scan( &nombre, &cuenta, &D_C_Tipo, &ANAL_T0, &ANAL_T1, &ANAL_T2, &ANAL_T3, &ANAL_T4, &ANAL_T5, &ANAL_T6, &ANAL_T7, &ANAL_T8, &ANAL_T9)



			if D_C_Tipo == 1 {
				D_C = "'D'"
			}
			if D_C_Tipo == 2 {
				D_C = "'C'"
			}
			if D_C_Tipo == 4 {
				D_C = "'D','C'"
			}
			queryC := "SELECT JRNAL_NO, JRNAL_LINE, AMOUNT FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_A_SALFLDG] WHERE ACCNT_CODE = '"+cuenta+"' AND D_C in ("+D_C+") AND PERIOD >= "+delPeriodo+" AND PERIOD <= "+alPeriodo+" "
			if Compare(ANAL_T0,"") != 0 {
				queryC = queryC+"AND ANAL_T0 = '"+ANAL_T0+"' "
			}
			if Compare(ANAL_T1,"") != 0 {
				queryC = queryC+"AND ANAL_T1 = '"+ANAL_T1+"' "
			}
			if Compare(ANAL_T2,"") != 0 {
				queryC = queryC+"AND ANAL_T2 = '"+ANAL_T2+"' "
			}
			if Compare(ANAL_T3,"") != 0 {
				queryC = queryC+"AND ANAL_T3 = '"+ANAL_T3+"' "
			}
			if Compare(ANAL_T4,"") != 0 {
				queryC = queryC+"AND ANAL_T4 = '"+ANAL_T4+"' "
			}
			if Compare(ANAL_T5,"") != 0 {
				queryC = queryC+"AND ANAL_T5 = '"+ANAL_T5+"' "
			}
			if Compare(ANAL_T6,"") != 0 {
				queryC = queryC+"AND ANAL_T6 = '"+ANAL_T6+"' "
			}
			if Compare(ANAL_T6,"") != 0 {
				queryC = queryC+"AND ANAL_T6 = '"+ANAL_T7+"' "
			}
			if Compare(ANAL_T8,"") != 0 {
				queryC = queryC+"AND ANAL_T8 = '"+ANAL_T8+"' "
			}
			if Compare(ANAL_T9,"") != 0 {
				queryC = queryC+"AND ANAL_T9 = '"+ANAL_T9+"' "
			}
			queryC = queryC + " order by PERIOD asc, TRANS_DATETIME asc"
			rowsC, errC := conn.Query(queryC)
			if errC != nil {
				log.Fatal("Open connection failed:", err.Error())
			}
			var amount decimal.Dec
			var JRNAL_NO int
			var JRNAL_LINE int
			for rowsC.Next()  {
				rowsC.Scan(&JRNAL_NO, &JRNAL_LINE, &amount)
				var conceLinea models.CedulasConceptoLineas
				var diarioLinea = strconv.Itoa(JRNAL_NO)+"-"+strconv.Itoa(JRNAL_LINE)
        		conceLinea = models.CedulasConceptoLineas{1,diarioLinea, amount.Float64()}
				models.AddCedulasConceptosLineas(conceLinea, diarioLinea)
        	}
		}
		example := map[string]interface{}{ "cedulasAcumulador": models.GetAllCedulasConceptosLineas() }
		c.Data["json"] = &example
		c.ServeJSON()
	}
}
func (c *GenerarCedulaPorConceptoController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		idConcepto := c.GetString("idConcepto")
		delPeriodo := c.GetString("delPeriodo")
		alPeriodo := c.GetString("alPeriodo")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
		query := "SELECT cl.idLinea,  cc.nombre, cl.cuenta, cl.D_C_Tipo, cl.ANAL_T0, cl.ANAL_T1, cl.ANAL_T2, cl.ANAL_T3, cl.ANAL_T4, cl.ANAL_T5, cl.ANAL_T6, cl.ANAL_T7, cl.ANAL_T8, cl.ANAL_T9 FROM [SUNPLUSADV].[dbo].[CedulasLinea] cl INNER JOIN [SUNPLUSADV].[dbo].[CedulasConceptos] cc on cc.idConcepto = cl.idConcepto WHERE cc.idConcepto = "+idConcepto+" order by cl.idLinea asc"
		rows, err := conn.Query(query)
		var idLinea int64
		var nombre string
		var cuenta string
		var D_C_Tipo int64
		var ANAL_T0 string
		var ANAL_T1 string
		var ANAL_T2 string
		var ANAL_T3 string
		var ANAL_T4 string
		var ANAL_T5 string
		var ANAL_T6 string
		var ANAL_T7 string
		var ANAL_T8 string
		var ANAL_T9 string
		D_C := ""
		acumulador := 0.0
		var cuentaAnterior string
		cuentaAnterior = ""
		primeraVez := true
		var IdConceptoOLinea int64
		IdConceptoOLinea = -1
		models.ClearCedulasConceptosLineas()
		for rows.Next()  {
			rows.Scan(&idLinea, &nombre, &cuenta, &D_C_Tipo, &ANAL_T0, &ANAL_T1, &ANAL_T2, &ANAL_T3, &ANAL_T4, &ANAL_T5, &ANAL_T6, &ANAL_T7, &ANAL_T8, &ANAL_T9)
	        if Compare(cuentaAnterior, cuenta) !=0 {
	        	if primeraVez  {
	        		primeraVez = false
	        	} else {
	        		var conceLinea models.CedulasConceptoLineas
        			conceLinea = models.CedulasConceptoLineas{IdConceptoOLinea,cuentaAnterior,acumulador}
					models.AddCedulasConceptosLineas(conceLinea, cuentaAnterior)
	        	}
	        	acumulador = 0.0
	        	
	        }
			if D_C_Tipo == 1 {
				D_C = "'D'"
			}
			if D_C_Tipo == 2 {
				D_C = "'C'"
			}
			if D_C_Tipo == 4 {
				D_C = "'D','C'"
			}
			queryC := "SELECT SUM(AMOUNT) as amount FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_A_SALFLDG] WHERE ACCNT_CODE = '"+cuenta+"' AND D_C in ("+D_C+") AND PERIOD >= "+delPeriodo+" AND PERIOD <= "+alPeriodo+" "
			if Compare(ANAL_T0,"") != 0 {
				queryC = queryC+"AND ANAL_T0 = '"+ANAL_T0+"' "
			}
			if Compare(ANAL_T1,"") != 0 {
				queryC = queryC+"AND ANAL_T1 = '"+ANAL_T1+"' "
			}
			if Compare(ANAL_T2,"") != 0 {
				queryC = queryC+"AND ANAL_T2 = '"+ANAL_T2+"' "
			}
			if Compare(ANAL_T3,"") != 0 {
				queryC = queryC+"AND ANAL_T3 = '"+ANAL_T3+"' "
			}
			if Compare(ANAL_T4,"") != 0 {
				queryC = queryC+"AND ANAL_T4 = '"+ANAL_T4+"' "
			}
			if Compare(ANAL_T5,"") != 0 {
				queryC = queryC+"AND ANAL_T5 = '"+ANAL_T5+"' "
			}
			if Compare(ANAL_T6,"") != 0 {
				queryC = queryC+"AND ANAL_T6 = '"+ANAL_T6+"' "
			}
			if Compare(ANAL_T6,"") != 0 {
				queryC = queryC+"AND ANAL_T6 = '"+ANAL_T7+"' "
			}
			if Compare(ANAL_T8,"") != 0 {
				queryC = queryC+"AND ANAL_T8 = '"+ANAL_T8+"' "
			}
			if Compare(ANAL_T9,"") != 0 {
				queryC = queryC+"AND ANAL_T9 = '"+ANAL_T9+"' "
			}
			rowsC, errC := conn.Query(queryC)
			if errC != nil {
				log.Fatal("Open connection failed:", err.Error())
			}
			var amount decimal.Dec
			for rowsC.Next()  {
				rowsC.Scan(&amount)
        	}
        	acumulador+= amount.Float64()
			cuentaAnterior = cuenta
			IdConceptoOLinea = idLinea
		}
		var conceLinea models.CedulasConceptoLineas
        conceLinea = models.CedulasConceptoLineas{IdConceptoOLinea,cuentaAnterior,acumulador}
		models.AddCedulasConceptosLineas(conceLinea, cuentaAnterior)
		example := map[string]interface{}{ "cedulasAcumulador": models.GetAllCedulasConceptosLineas() }
		c.Data["json"] = &example
		c.ServeJSON()
	}
}
func (c *GenerarCedulasController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		idCedula := c.GetString("idCedula")
		delPeriodo := c.GetString("delPeriodo")
		alPeriodo := c.GetString("alPeriodo")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
		query := "SELECT cl.idLinea,  cc.nombre, cl.cuenta, cl.D_C_Tipo, cl.ANAL_T0, cl.ANAL_T1, cl.ANAL_T2, cl.ANAL_T3, cl.ANAL_T4, cl.ANAL_T5, cl.ANAL_T6, cl.ANAL_T7, cl.ANAL_T8, cl.ANAL_T9, cl.idConcepto FROM [SUNPLUSADV].[dbo].[CedulasLinea] cl INNER JOIN [SUNPLUSADV].[dbo].[CedulasConceptos] cc on cc.idConcepto = cl.idConcepto WHERE cc.idCedula = "+idCedula+" order by cl.idLinea asc"
		rows, err := conn.Query(query)
		var idLinea int64
		var idConcepto int64
		var nombre string
		var cuenta string
		var D_C_Tipo int64
		var ANAL_T0 string
		var ANAL_T1 string
		var ANAL_T2 string
		var ANAL_T3 string
		var ANAL_T4 string
		var ANAL_T5 string
		var ANAL_T6 string
		var ANAL_T7 string
		var ANAL_T8 string
		var ANAL_T9 string
		D_C := ""
		acumulador := 0.0
		var idConceptoAnterior int64
		idConceptoAnterior = -1
		primeraVez := true
		nombreAnterior := ""
		var IdConceptoOLinea int64
		IdConceptoOLinea = -1
		models.ClearCedulasConceptosLineas()
		for rows.Next()  {
			rows.Scan(&idLinea, &nombre, &cuenta, &D_C_Tipo, &ANAL_T0, &ANAL_T1, &ANAL_T2, &ANAL_T3, &ANAL_T4, &ANAL_T5, &ANAL_T6, &ANAL_T7, &ANAL_T8, &ANAL_T9, &idConcepto)
	        if idConceptoAnterior != idConcepto {
	        	if primeraVez  {
	        		primeraVez = false
	        	} else {
	        		var conceLinea models.CedulasConceptoLineas
        			conceLinea = models.CedulasConceptoLineas{IdConceptoOLinea,nombreAnterior,acumulador}
					models.AddCedulasConceptosLineas(conceLinea, nombreAnterior)
	        	}
	        	acumulador = 0.0
	        	idConceptoAnterior = idConcepto
	        }
			if D_C_Tipo == 1 {
				D_C = "'D'"
			}
			if D_C_Tipo == 2 {
				D_C = "'C'"
			}
			if D_C_Tipo == 4 {
				D_C = "'D','C'"
			}
			queryC := "SELECT SUM(AMOUNT) as amount FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_A_SALFLDG] WHERE ACCNT_CODE = '"+cuenta+"' AND D_C in ("+D_C+") AND PERIOD >= "+delPeriodo+" AND PERIOD <= "+alPeriodo+" "
			if Compare(ANAL_T0,"") != 0 {
				queryC = queryC+"AND ANAL_T0 = '"+ANAL_T0+"' "
			}
			if Compare(ANAL_T1,"") != 0 {
				queryC = queryC+"AND ANAL_T1 = '"+ANAL_T1+"' "
			}
			if Compare(ANAL_T2,"") != 0 {
				queryC = queryC+"AND ANAL_T2 = '"+ANAL_T2+"' "
			}
			if Compare(ANAL_T3,"") != 0 {
				queryC = queryC+"AND ANAL_T3 = '"+ANAL_T3+"' "
			}
			if Compare(ANAL_T4,"") != 0 {
				queryC = queryC+"AND ANAL_T4 = '"+ANAL_T4+"' "
			}
			if Compare(ANAL_T5,"") != 0 {
				queryC = queryC+"AND ANAL_T5 = '"+ANAL_T5+"' "
			}
			if Compare(ANAL_T6,"") != 0 {
				queryC = queryC+"AND ANAL_T6 = '"+ANAL_T6+"' "
			}
			if Compare(ANAL_T6,"") != 0 {
				queryC = queryC+"AND ANAL_T6 = '"+ANAL_T7+"' "
			}
			if Compare(ANAL_T8,"") != 0 {
				queryC = queryC+"AND ANAL_T8 = '"+ANAL_T8+"' "
			}
			if Compare(ANAL_T9,"") != 0 {
				queryC = queryC+"AND ANAL_T9 = '"+ANAL_T9+"' "
			}
			rowsC, errC := conn.Query(queryC)
			if errC != nil {
				log.Fatal("Open connection failed:", err.Error())
			}
			var amount decimal.Dec
			for rowsC.Next()  {
				rowsC.Scan(&amount)
        	}
        	acumulador+= amount.Float64()
			nombreAnterior = nombre
			IdConceptoOLinea = idConcepto
		}
		var conceLinea models.CedulasConceptoLineas
        conceLinea = models.CedulasConceptoLineas{IdConceptoOLinea,nombreAnterior,acumulador}
		models.AddCedulasConceptosLineas(conceLinea, nombreAnterior)
		example := map[string]interface{}{ "cedulasAcumulador": models.GetAllCedulasConceptosLineas() }
		c.Data["json"] = &example
		c.ServeJSON()
	}
}




func (c *ListaConceptosCedulasController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		idCedula := c.GetString("idCedula")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
		query := "SELECT idConcepto, idCedula, nombre FROM [SUNPLUSADV].[dbo].[CedulasConceptos] WHERE idCedula = "+idCedula+" order by idConcepto asc"
		rows, err := conn.Query(query)
		var idConcepto int64
		var idCedulaX int64
		var nombre string
		var cedul  models.CedulasConceptos
        for rows.Next()  {
			rows.Scan(&idConcepto, &idCedulaX, &nombre)
			cedul = models.CedulasConceptos{idConcepto,idCedulaX,nombre}
			models.AddCedulasConceptos(cedul, nombre)
		}
		example := map[string]interface{}{ "cedulasConceptos": models.GetAllCedulasConceptos() }
		c.Data["json"] = &example
		c.ServeJSON()
	}
}


func (c *DiariosReversiadosController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
		query := "SELECT DISTINCT JRNAL_NO, ALLOC_REF,PERIOD FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_A_SALFLDG] WHERE ALLOCATION ='C' order by ALLOC_REF asc, JRNAL_NO asc"
		rows, err := conn.Query(query)
		var JRNAL_NO int
		var ALLOC_REF int		
		var PERIOD int64
		models.ClearDiariosReversiados()
		var cedul  models.DiarioRe
        for rows.Next()  {
			rows.Scan(&JRNAL_NO, &ALLOC_REF, &PERIOD)
			cedul = models.DiarioRe{JRNAL_NO,ALLOC_REF,PERIOD}
			models.AddDiarioReversiado(cedul, strconv.Itoa(JRNAL_NO) )
		}
		example := map[string]interface{}{ "diariosReversiados": models.GetAllDiarioReversiados() }
		c.Data["json"] = &example
		c.ServeJSON()
	}
}


func (c *ListaLineasTiposDeDiarioController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		idTipoDeDiario := c.GetString("idTipoDeDiario")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}		
		query := "SELECT idLinea, Linea, deboFacturar, servicio, cliente, concepto, DESCRIPTN, ACNT_CODE , D_C, ANAL_T0, ANAL_T1, ANAL_T2, ANAL_T3, ANAL_T4, ANAL_T5, ANAL_T6, ANAL_T7, ANAL_T8, ANAL_T9 FROM [SUNPLUSADV].[dbo].[TiposDeDiarioLineas] WHERE idTipoDeDiario = "+idTipoDeDiario+" order by Linea asc"
		fmt.Println(query)
		rows, err := conn.Query(query)
		var idLinea int
		var Linea int
		var deboFacturar int
		var cuenta string
		var cliente string
		var servicio string
		var concepto string
		var DESCRIPTN string
		var D_C string
		var ANAL_T0 string
		var ANAL_T1 string
		var ANAL_T2 string
		var ANAL_T3 string
		var ANAL_T4 string
		var ANAL_T5 string
		var ANAL_T6 string
		var ANAL_T7 string
		var ANAL_T8 string
		var ANAL_T9 string
		models.ClearTiposDeDiarioLineas()
		var cedul  models.TipoDeDiariosLineas
        for rows.Next()  {
			rows.Scan(&idLinea, &Linea, &deboFacturar, &servicio, &cliente, &concepto, &DESCRIPTN, &cuenta, &D_C, &ANAL_T0, &ANAL_T1, &ANAL_T2, &ANAL_T3, &ANAL_T4, &ANAL_T5, &ANAL_T6, &ANAL_T7, &ANAL_T8, &ANAL_T9)
			cedul = models.TipoDeDiariosLineas{idLinea, Linea, deboFacturar, cuenta, cliente, servicio, concepto, DESCRIPTN, D_C, ANAL_T0,ANAL_T1,ANAL_T2,ANAL_T3,ANAL_T4,ANAL_T5,ANAL_T6,ANAL_T7,ANAL_T8,ANAL_T9}
			models.AddTiposDeDiarioLineas(cedul, strconv.Itoa(idLinea) )
		}
		example := map[string]interface{}{ "tiposDeDiarioLineas": models.GetAllTiposDeDiariosLineas() }
		c.Data["json"] = &example
		c.ServeJSON()
	}
}


func (c *ChecaOtrosCamposController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
		query := "SELECT idCampo, nombre, url FROM [SUNPLUSADV].[dbo].[zCampos] order by idCampo asc"
		rows, err := conn.Query(query)
		if err != nil {
			example := map[string]interface{}{ "success" : 0 }
			c.Data["json"] = &example
			c.ServeJSON()
		}
		var idCampo int
		var nombre string
		var url string
		models.ClearOtrosCampos()
		var cedul  models.OtrosCampos
        for rows.Next()  {
			rows.Scan(&idCampo, &nombre, &url)
			cedul = models.OtrosCampos{idCampo,nombre,url}
			models.AddOtrosCampos(cedul, strconv.Itoa(idCampo) )
		}
		currentTime := int64(time.Now().Unix())
		tm := time.Unix(currentTime, 0)
		dateString := tm.String() 
		substring := string(dateString[0:10])
	   	byteArray := []byte(substring)
		hasher := sha512.New()
	    hasher.Write(byteArray)
		cryptoText := base64.StdEncoding.EncodeToString(  []byte(hex.EncodeToString(hasher.Sum(nil))))
		example := map[string]interface{}{ "success" : 1 , "otrosCampos": models.GetAllOtrosCampos(), "hash" : cryptoText, "BUNIT" : BUNIT.(string) }
		c.Data["json"] = &example
		c.ServeJSON()
	}
}
func (c *ListaLineasCedulasController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		IdConcepto := c.GetString("IdConcepto")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
		query := "SELECT idLinea, idConcepto, cuenta, D_C_Tipo, ANAL_T0, ANAL_T1, ANAL_T2, ANAL_T3, ANAL_T4, ANAL_T5, ANAL_T6, ANAL_T7, ANAL_T8, ANAL_T9 FROM [SUNPLUSADV].[dbo].[CedulasLinea] WHERE idConcepto = "+IdConcepto+" order by idLinea asc"
		rows, err := conn.Query(query)
		var idLinea int
		var idConceptoX int64		
		var cuenta string
		var D_C_Tipo int64
		var ANAL_T0 string
		var ANAL_T1 string
		var ANAL_T2 string
		var ANAL_T3 string
		var ANAL_T4 string
		var ANAL_T5 string
		var ANAL_T6 string
		var ANAL_T7 string
		var ANAL_T8 string
		var ANAL_T9 string
		models.ClearCedulasLineas()
		var cedul  models.CedulasLineas
        for rows.Next()  {
			rows.Scan(&idLinea, &idConceptoX, &cuenta, &D_C_Tipo, &ANAL_T0, &ANAL_T1, &ANAL_T2, &ANAL_T3, &ANAL_T4, &ANAL_T5, &ANAL_T6, &ANAL_T7, &ANAL_T8, &ANAL_T9)
			cedul = models.CedulasLineas{idLinea,idConceptoX,cuenta,D_C_Tipo,ANAL_T0,ANAL_T1,ANAL_T2,ANAL_T3,ANAL_T4,ANAL_T5,ANAL_T6,ANAL_T7,ANAL_T8,ANAL_T9}
			models.AddCedulasLineas(cedul, strconv.Itoa(idLinea) )
		}
		example := map[string]interface{}{ "cedulasLineas": models.GetAllCedulasLineas() }
		c.Data["json"] = &example
		c.ServeJSON()
	}
}
func (c *VeDetallePrimerNivelController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		tipo := c.GetString("tipo")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
		PERIOD := c.GetString("PERIOD")
		anio:=string(PERIOD[0:4])	
		mes:=string(PERIOD[5:7])	
		PERIOD_ANTERIOR :=""
		mesNumero, _ := strconv.Atoi(mes)
		if mesNumero == 1 {
			mesS := "12"
			anioAnterior, _ := strconv.Atoi(anio)
			anioAnterior--
			PERIOD_ANTERIOR = strconv.Itoa(anioAnterior)+"0"+mesS
		} else {
			mesNumero--
			mesS := strconv.Itoa(mesNumero)
			if mesNumero < 10 {
				mesS = "0"+mesS
			}
			PERIOD_ANTERIOR = anio+"0"+mesS
		}
		
	

		if Compare(tipo,"1")==0 {//ACTIVOS CORRIENTES
		  	//1 - 109, fondo 10, CAJAS Y BANCOS
			models.ClearLineaDetalle()
			var cedul  models.LineaDetalle
       
			query1 := `SELECT ISNULL(SUM(AMOUNT),0) as amount
			  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
			  WHERE SUBSTRING(ACCNT_CODE,1,3) in('100','101','102','103','104','105','106','107','108','109')
			  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
			rows1, err1 := conn.Query(query1)
			if err1 != nil {
				fmt.Println(query1)
				panic(err1)
			}
			var AMOUNT_CAJA_BANCOS decimal.Dec
			if rows1.Next()  {
				rows1.Scan(&AMOUNT_CAJA_BANCOS)
				amountPrima := math.Abs(AMOUNT_CAJA_BANCOS.Float64())
				cedul = models.LineaDetalle{1,"Caja y bancos",amountPrima}
				models.AddLineaDetalle(cedul,"caja")
			}

			//11 - 119, fondo 10, INVERSIONES
		
			query2 := `SELECT ISNULL(SUM(AMOUNT),0) as amount
			  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
			  WHERE SUBSTRING(ACCNT_CODE,1,3) in('110','111','112','113','114','115','116','117','118','119')
			  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
			rows2, err2 := conn.Query(query2)
			if err2 != nil {
				fmt.Println(query2)
				panic(err2)
			}
			var AMOUNT_INVERSIONES decimal.Dec
			if rows2.Next()  {
				rows2.Scan(&AMOUNT_INVERSIONES)
				amountPrima := math.Abs(AMOUNT_INVERSIONES.Float64())
				cedul = models.LineaDetalle{2,"Inversiones",amountPrima}
				models.AddLineaDetalle(cedul,"inversiones")
			}

			//139, fondo 10, CUENTAS X COBRAR
			//12, 14, 131,132,133,134,135,136,137,138, CUENTAS X COBRAR
			//3 - 349, fondo 10, CUENTAS X COBRAR
			var AMOUNT_CUENTAS_POR_COBRAR decimal.Dec
			acumulador := 0.0
			
			
			query3_A := `SELECT ISNULL(SUM(AMOUNT),0) as amount
			  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
			  WHERE SUBSTRING(ACCNT_CODE,1,3) in('139')
			  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
			rows3_A, err3_A := conn.Query(query3_A)
			if err3_A != nil {
				fmt.Println(query3_A)
				panic(err3_A)
			}
			if rows3_A.Next()  {
				rows3_A.Scan(&AMOUNT_CUENTAS_POR_COBRAR)
				amountPrima := math.Abs(AMOUNT_CUENTAS_POR_COBRAR.Float64())
				acumulador=acumulador+amountPrima
			}
			var ACNT_CODE string

			query3_BB := `SELECT ACNT_CODE FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ACNT]
 			 WHERE ISNUMERIC( SUBSTRING(ACNT_CODE,1,1))<>1 OR (SUBSTRING(ACNT_CODE,1,2) in('12','14','30','31','32','33','34') OR SUBSTRING(ACNT_CODE,1,3) in('130','131','132','133','134','135','136','137','138'))`
			rows3_BB, err3_BB := conn.Query(query3_BB)
			if err3_BB != nil {
				fmt.Println(query3_BB)
				panic(err3_BB)
			}
			for rows3_BB.Next()  {
				rows3_BB.Scan(&ACNT_CODE)
				query3_B := `SELECT ISNULL(SUM(AMOUNT),0) as amount
				  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
				  WHERE ACCNT_CODE = '`+ACNT_CODE+`'
				  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
				rows3_B, err3_B := conn.Query(query3_B)
				if err3_B != nil {
					fmt.Println(query3_B)
					panic(err3_B)
				}
				if rows3_B.Next()  {
					rows3_B.Scan(&AMOUNT_CUENTAS_POR_COBRAR)
					amountPrima1 := AMOUNT_CUENTAS_POR_COBRAR.Float64()
					if amountPrima1 < 0 {
						amountPrima := math.Abs(AMOUNT_CUENTAS_POR_COBRAR.Float64())
						acumulador=acumulador+amountPrima	
					}
				}
			}

			cedul = models.LineaDetalle{3,"Cuentas por cobrar",acumulador}
			models.AddLineaDetalle(cedul,"Cuentas por cobrar")
			//16- 169, fondo10, INVENTARIOS
		

			query4 := `SELECT ISNULL(SUM(AMOUNT),0) as amount
			  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
			  WHERE SUBSTRING(ACCNT_CODE,1,2) in('16')
			  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
			rows4, err4 := conn.Query(query4)
			if err4 != nil {
				fmt.Println(query4)
				panic(err4)
			}
			var AMOUNT_INVENTARIO decimal.Dec
			if rows4.Next()  {
				rows4.Scan(&AMOUNT_INVENTARIO)
				amountPrima := math.Abs(AMOUNT_INVENTARIO.Float64())
				cedul = models.LineaDetalle{4,"Inventario",amountPrima}
				models.AddLineaDetalle(cedul,"Inventario")
			}
			//17- 189, fondo10, Pagos anticipados
			
			query5 := `SELECT ISNULL(SUM(AMOUNT),0) as amount
			  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
			  WHERE SUBSTRING(ACCNT_CODE,1,2) in('17','18')
			  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
			rows5, err5 := conn.Query(query5)
			if err5 != nil {
				fmt.Println(query5)
				panic(err5)
			}
			var AMOUNT_PAGOS_ANTICIPADOS decimal.Dec
			if rows5.Next()  {
				rows5.Scan(&AMOUNT_PAGOS_ANTICIPADOS)
				amountPrima := math.Abs(AMOUNT_PAGOS_ANTICIPADOS.Float64())
				cedul = models.LineaDetalle{5,"Pagos anticipados",amountPrima}
				models.AddLineaDetalle(cedul,"pagosAnticipados")
			}

			//19, fondo 10, cuentas por cobrar entre fondos
			
			query6 := `SELECT ISNULL(SUM(AMOUNT),0) as amount
			  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
			  WHERE SUBSTRING(ACCNT_CODE,1,2) in('19')
			  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
			rows6, err6 := conn.Query(query6)
			if err6 != nil {
				fmt.Println(query6)
				panic(err6)
			}
			var AMOUNT_ENTRE_FONDOS decimal.Dec
			if rows6.Next()  {
				rows6.Scan(&AMOUNT_ENTRE_FONDOS)
				amountPrima := math.Abs(AMOUNT_ENTRE_FONDOS.Float64())
				cedul = models.LineaDetalle{6,"Cuentas por cobrar entre fondos",amountPrima}
				models.AddLineaDetalle(cedul,"entreFondos")
			}

			//15- 159, fondo10, Documentos y prestamos x pagar
			

			query7 := `SELECT ISNULL(SUM(AMOUNT),0) as amount
			  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
			  WHERE SUBSTRING(ACCNT_CODE,1,2) in('15')
			  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
			rows7, err7 := conn.Query(query7)
			if err7 != nil {
				fmt.Println(query7)
				panic(err7)
			}
			var AMOUNT_DOCUMENTOS decimal.Dec
			if rows7.Next()  {
				rows7.Scan(&AMOUNT_DOCUMENTOS)
				amountPrima := math.Abs(AMOUNT_DOCUMENTOS.Float64())
				cedul = models.LineaDetalle{7,"Documentos",amountPrima}
				models.AddLineaDetalle(cedul,"Documentos y prestamos x pagar")
			}
		} else {
			if Compare(tipo,"2")==0 {//PASIVOS CORRIENTES
		  		//1 - 109, fondo 10, CAJAS Y BANCOS
				models.ClearLineaDetalle()
				var cedul  models.LineaDetalle
	       
				var AMOUNT_CUENTAS_POR_PAGAR decimal.Dec
				acumulador := 0.0
				var ACNT_CODE string

				query3_BB := `SELECT ACNT_CODE FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ACNT]
	 			 WHERE ISNUMERIC( SUBSTRING(ACNT_CODE,1,1))<>1 OR (SUBSTRING(ACNT_CODE,1,2) in('12','14','30','31','32','33','34') OR SUBSTRING(ACNT_CODE,1,3) in('130','131','132','133','134','135','136','137','138'))`
				rows3_BB, err3_BB := conn.Query(query3_BB)
				if err3_BB != nil {
					fmt.Println(query3_BB)
					panic(err3_BB)
				}
				for rows3_BB.Next()  {
					rows3_BB.Scan(&ACNT_CODE)
					query3_B := `SELECT ISNULL(SUM(AMOUNT),0) as amount
					  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
					  WHERE ACCNT_CODE = '`+ACNT_CODE+`'
					  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
					rows3_B, err3_B := conn.Query(query3_B)
					if err3_B != nil {
						fmt.Println(query3_B)
						panic(err3_B)
					}
					if rows3_B.Next()  {
						rows3_B.Scan(&AMOUNT_CUENTAS_POR_PAGAR)
						amountPrima1 := AMOUNT_CUENTAS_POR_PAGAR.Float64()
						if amountPrima1 > 0 {
							amountPrima := math.Abs(AMOUNT_CUENTAS_POR_PAGAR.Float64())
							acumulador=acumulador+amountPrima	
						}
					}
				}

				cedul = models.LineaDetalle{1,"Cuentas por pagar",acumulador}
				models.AddLineaDetalle(cedul,"Cuentas por pagar")
				
				query7 := `SELECT ISNULL(SUM(AMOUNT),0) as amount
				  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
				  WHERE SUBSTRING(ACCNT_CODE,1,2) in('35')
				  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
				rows7, err7 := conn.Query(query7)
				if err7 != nil {
					fmt.Println(query7)
					panic(err7)
				}
				var AMOUNT_DOCUMENTOS decimal.Dec
				if rows7.Next()  {
					rows7.Scan(&AMOUNT_DOCUMENTOS)
					amountPrima := math.Abs(AMOUNT_DOCUMENTOS.Float64())
					cedul = models.LineaDetalle{2,"Documentos",amountPrima}
					models.AddLineaDetalle(cedul,"Documentos x pagar")
				}

				query1 := `SELECT ISNULL(SUM(AMOUNT),0) as amount
				  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
				  WHERE SUBSTRING(ACCNT_CODE,1,2) in('36')
				  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
				rows1, err1 := conn.Query(query1)
				if err1 != nil {
					fmt.Println(query1)
					panic(err1)
				}
				var AMOUNT_FONDOS_CONFIADOS decimal.Dec
				if rows1.Next()  {
					rows1.Scan(&AMOUNT_FONDOS_CONFIADOS)
					amountPrima := math.Abs(AMOUNT_FONDOS_CONFIADOS.Float64())
					cedul = models.LineaDetalle{3,"Fondos",amountPrima}
					models.AddLineaDetalle(cedul,"Fondos confiados y de agencia")
				}

				query2 := `SELECT ISNULL(SUM(AMOUNT),0) as amount
				  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
				  WHERE SUBSTRING(ACCNT_CODE,1,2) in('37','38')
				  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
				rows2, err2 := conn.Query(query2)
				if err2 != nil {
					fmt.Println(query2)
					panic(err2)
				}
				var AMOUNT_OTROS_PASIVOS_CORRIENTES decimal.Dec
				if rows2.Next()  {
					rows2.Scan(&AMOUNT_OTROS_PASIVOS_CORRIENTES)
					amountPrima := math.Abs(AMOUNT_OTROS_PASIVOS_CORRIENTES.Float64())
					cedul = models.LineaDetalle{4,"Otros pasivos corrientes",amountPrima}
					models.AddLineaDetalle(cedul,"Otros pasivos corrientes")
				}

			
				acumulador = 0.0
				query4_BB := `SELECT ACNT_CODE FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ACNT]
	 			 WHERE SUBSTRING(ACNT_CODE,1,2) in('19')`
				rows4_BB, err4_BB := conn.Query(query4_BB)
				if err4_BB != nil {
					fmt.Println(query4_BB)
					panic(err4_BB)
				}
				for rows4_BB.Next()  {
					rows4_BB.Scan(&ACNT_CODE)
					query4_B := `SELECT ISNULL(SUM(AMOUNT),0) as amount
					  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
					  WHERE ACCNT_CODE = '`+ACNT_CODE+`'
					  AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 = '10'`
					rows4_B, err4_B := conn.Query(query4_B)
					if err4_B != nil {
						fmt.Println(query4_B)
						panic(err4_B)
					}
					if rows4_B.Next()  {
						rows4_B.Scan(&AMOUNT_CUENTAS_POR_PAGAR)
						amountPrima1 := AMOUNT_CUENTAS_POR_PAGAR.Float64()
						if amountPrima1 > 0 {
							amountPrima := math.Abs(AMOUNT_CUENTAS_POR_PAGAR.Float64())
							acumulador=acumulador+amountPrima	
						}
					}
				}
				cedul = models.LineaDetalle{5,"Cuentas por pagar entre fondos",acumulador}
				models.AddLineaDetalle(cedul,"Cuentas por pagar entre fondos")
			} else {
			if Compare(tipo,"3")==0 {//CGASTOS OPERATIVOS
		  		

						models.ClearLineaDetalle()
						var cedul  models.LineaDetalle

						esteAnio:=string(PERIOD[0:4])	
						PERIOD_INICIO := esteAnio+"001"

						anteriorAnio, _ := strconv.Atoi(esteAnio)
						anteriorAnio--
						anteriorAnioS := strconv.Itoa(anteriorAnio)

						//PERIOD_INICIO_ANTERIOR := anteriorAnioS+"001"
						PERIOD_ANTERIOR := anteriorAnioS+string(PERIOD[4:7])
						acumulador := 0.0
			       
						
						//AND  SUBSTRING( CAST(PERIOD AS NVARCHAR(11)),1,4)  = '`+esteAnio+`'

						query7 := `SELECT ISNULL(SUM(AMOUNT),0) as amount
						  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
						  WHERE (SUBSTRING(ACCNT_CODE,1,2) in('80','81','82','85','86','87','88','89','90','91','92','93','94','95','98') OR
						  SUBSTRING(ACCNT_CODE,1,3) in('978') )
						  AND PERIOD >= `+PERIOD_INICIO+` AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10','20') 
						  AND ANAL_T3 >= 'AFO'`
						rows7, err7 := conn.Query(query7)
						if err7 != nil {
							fmt.Println(query7)
							panic(err7)
						}
						var AMOUNT_ESTE_ANO decimal.Dec
						if rows7.Next()  {
							rows7.Scan(&AMOUNT_ESTE_ANO)
							amountPrima := math.Abs(AMOUNT_ESTE_ANO.Float64())
							acumulador = acumulador + amountPrima
							cedul = models.LineaDetalle{1,"Este año",amountPrima}
							models.AddLineaDetalle(cedul,"Este año")
						}
						
						query1 := `SELECT ISNULL(SUM(AMOUNT),0) as amount
						  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
						  WHERE (SUBSTRING(ACCNT_CODE,1,2) in('80','81','82','85','86','87','88','89','90','91','92','93','94','95','98') OR
						  SUBSTRING(ACCNT_CODE,1,3) in('978') )
							 AND PERIOD >= `+PERIOD_ANTERIOR+` AND PERIOD < `+PERIOD_INICIO+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10','20') 
						   AND ANAL_T3 >= 'AFO'`
							
						
						rows1, err1 := conn.Query(query1)
						if err1 != nil {
							fmt.Println(query1)
							panic(err1)
						}
						var AMOUNT_FONDOS_CONFIADOS decimal.Dec
						if rows1.Next()  {
							rows1.Scan(&AMOUNT_FONDOS_CONFIADOS)
							amountPrima := math.Abs(AMOUNT_FONDOS_CONFIADOS.Float64())
							acumulador = acumulador + amountPrima
							cedul = models.LineaDetalle{2,"Más el año anterior",amountPrima}
							models.AddLineaDetalle(cedul,"Más el año anterior")
						}

						query8 := `SELECT ISNULL(SUM(AMOUNT),0) as amount
						  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
						  WHERE ACCNT_CODE = '878777'
						  AND PERIOD >= `+PERIOD_INICIO+` AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10','20') 
						  AND ANAL_T3 >= 'AFO'`
						rows8, err8 := conn.Query(query8)
						if err8 != nil {
							fmt.Println(query8)
							panic(err8)
						}
						var AMOUNT_DIEZMO_UNION decimal.Dec
						if rows8.Next()  {
							rows8.Scan(&AMOUNT_DIEZMO_UNION)
							amountPrima := math.Abs(AMOUNT_DIEZMO_UNION.Float64())
							acumulador = acumulador - amountPrima
							cedul = models.LineaDetalle{3,"Menos diezmo a unión",amountPrima}
							models.AddLineaDetalle(cedul,"Menos diezmo a unión")
						}

						query88 := `SELECT ISNULL(SUM(AMOUNT),0) as amount
						  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
						  WHERE ACCNT_CODE = '878777'
						  AND PERIOD >= `+PERIOD_ANTERIOR+` AND PERIOD < `+PERIOD_INICIO+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10','20') 
						  AND ANAL_T3 >= 'AFO'`
						rows88, err88 := conn.Query(query88)
						if err88 != nil {
							fmt.Println(query88)
							panic(err88)
						}
						if rows88.Next()  {
							rows88.Scan(&AMOUNT_DIEZMO_UNION)
							amountPrima := math.Abs(AMOUNT_DIEZMO_UNION.Float64())
							acumulador = acumulador - amountPrima
							cedul = models.LineaDetalle{4,"Menos diezmo a unión - año pasado",amountPrima}
							models.AddLineaDetalle(cedul,"Menos diezmo a unión - año pasado")
						}

						query9 := `SELECT ISNULL(SUM(AMOUNT),0) as amount
						  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
						  WHERE ACCNT_CODE = '878555'
						  AND PERIOD >= `+PERIOD_INICIO+` AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10','20') 
						  AND ANAL_T3 >= 'AFO'`
						rows9, err9 := conn.Query(query9)
						if err9 != nil {
							fmt.Println(query9)
							panic(err9)
						}
						var AMOUNT_DIEZMO_TAM decimal.Dec
						if rows9.Next()  {
							rows9.Scan(&AMOUNT_DIEZMO_TAM)
							amountPrima := math.Abs(AMOUNT_DIEZMO_TAM.Float64())
							acumulador = acumulador - amountPrima
							cedul = models.LineaDetalle{5,"Menos diezmo a TAM",amountPrima}
							models.AddLineaDetalle(cedul,"Menos diezmo a TAM")
						}

						query99 := `SELECT ISNULL(SUM(AMOUNT),0) as amount
						  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
						  WHERE ACCNT_CODE = '878555'
						  AND PERIOD >= `+PERIOD_ANTERIOR+` AND PERIOD < `+PERIOD_INICIO+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10','20') 
						  AND ANAL_T3 >= 'AFO'`
						rows99, err99 := conn.Query(query99)
						if err99 != nil {
							fmt.Println(query99)
							panic(err99)
						}
						if rows99.Next()  {
							rows99.Scan(&AMOUNT_DIEZMO_TAM)
							amountPrima := math.Abs(AMOUNT_DIEZMO_TAM.Float64())
							acumulador = acumulador - amountPrima
							cedul = models.LineaDetalle{6,"Menos diezmo a TAM - año pasado",amountPrima}
							models.AddLineaDetalle(cedul,"Menos diezmo a TAM - año pasado")
						}

						
						queryA := `SELECT ISNULL(SUM(AMOUNT),0) as amount
						  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
						  WHERE (SUBSTRING(ACCNT_CODE,1,2) in('80','81','82','83','84','85','86','87','88','89','90','91','92','93','94','95','98') OR
						  SUBSTRING(ACCNT_CODE,1,3) in('978') )
							 AND PERIOD >= `+PERIOD_INICIO+` AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10','20') 
						   AND ANAL_T3 = 'AFOMISIO01'`
							
						rowsA, errA := conn.Query(queryA)
						if errA != nil {
							fmt.Println(queryA)
							panic(errA)
						}
						var AMOUNT_AFOMISIO01 decimal.Dec
						if rowsA.Next()  {
							rowsA.Scan(&AMOUNT_AFOMISIO01)
							amountPrima := math.Abs(AMOUNT_AFOMISIO01.Float64())
							acumulador = acumulador - amountPrima
							cedul = models.LineaDetalle{7,"Menos Ofrenda AFOMISIO01",amountPrima}
							models.AddLineaDetalle(cedul,"Menos Ofrenda AFOMISIO01")
						}

						queryAA := `SELECT ISNULL(SUM(AMOUNT),0) as amount
						  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
						  WHERE (SUBSTRING(ACCNT_CODE,1,2) in('80','81','82','83','84','85','86','87','88','89','90','91','92','93','94','95','98') OR
						  SUBSTRING(ACCNT_CODE,1,3) in('978') )
							 AND PERIOD >= `+PERIOD_ANTERIOR+` AND PERIOD < `+PERIOD_INICIO+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10','20') 
						   AND ANAL_T3 = 'AFOMISIO01'`
							
						rowsAA, errAA := conn.Query(queryAA)
						if errAA != nil {
							fmt.Println(queryAA)
							panic(errAA)
						}
						if rowsAA.Next()  {
							rowsAA.Scan(&AMOUNT_AFOMISIO01)
							amountPrima := math.Abs(AMOUNT_AFOMISIO01.Float64())
							acumulador = acumulador - amountPrima
							cedul = models.LineaDetalle{8,"Menos Ofrenda AFOMISIO01 - año pasado",amountPrima}
							models.AddLineaDetalle(cedul,"Menos Ofrenda AFOMISIO01 - año pasado")
						}

						queryB := `SELECT ISNULL(SUM(AMOUNT),0) as amount
						  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
						  WHERE (SUBSTRING(ACCNT_CODE,1,2) in('80','81','82','83','84','85','86','87','88','89','90','91','92','93','94','95','98') OR
						  SUBSTRING(ACCNT_CODE,1,3) in('978') )
							 AND PERIOD >= `+PERIOD_INICIO+` AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10','20') 
						   AND ANAL_T3 = 'AFOAG1111G'`
							
						rowsB, errB := conn.Query(queryB)
						if errB != nil {
							fmt.Println(queryB)
							panic(errB)
						}
						var AMOUNT_AFOAG1111G decimal.Dec
						if rowsB.Next()  {
							rowsB.Scan(&AMOUNT_AFOAG1111G)
							amountPrima := math.Abs(AMOUNT_AFOAG1111G.Float64())
							acumulador = acumulador - amountPrima
							cedul = models.LineaDetalle{9,"Menos Ofrenda AFOAG1111G",amountPrima}
							models.AddLineaDetalle(cedul,"Menos Ofrenda AFOAG1111G")
						}

						queryBB := `SELECT ISNULL(SUM(AMOUNT),0) as amount
						  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
						  WHERE (SUBSTRING(ACCNT_CODE,1,2) in('80','81','82','83','84','85','86','87','88','89','90','91','92','93','94','95','98') OR
						  SUBSTRING(ACCNT_CODE,1,3) in('978') )
							 AND PERIOD >= `+PERIOD_ANTERIOR+` AND PERIOD < `+PERIOD_INICIO+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10','20') 
						   AND ANAL_T3 = 'AFOAG1111G'`
							
						rowsBB, errBB := conn.Query(queryBB)
						if errBB != nil {
							fmt.Println(queryBB)
							panic(errBB)
						}
						if rowsBB.Next()  {
							rowsBB.Scan(&AMOUNT_AFOAG1111G)
							amountPrima := math.Abs(AMOUNT_AFOAG1111G.Float64())
							acumulador = acumulador - amountPrima
							cedul = models.LineaDetalle{10,"Menos Ofrenda AFOAG1111G - año pasado",amountPrima}
							models.AddLineaDetalle(cedul,"Menos Ofrenda AFOAG1111G - año pasado")
						}

						queryC := `SELECT ISNULL(SUM(AMOUNT),0) as amount
						  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
						  WHERE (SUBSTRING(ACCNT_CODE,1,2) in('80','81','82','83','84','85','86','87','88','89','90','91','92','93','94','95','98') OR
						  SUBSTRING(ACCNT_CODE,1,3) in('978') )
							 AND PERIOD >= `+PERIOD_INICIO+` AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10','20') 
						   AND ANAL_T3 = 'AFOAG1111D'`
							
						rowsC, errC := conn.Query(queryC)
						if errC != nil {
							fmt.Println(queryC)
							panic(errC)
						}
						var AMOUNT_AFOAG1111GD decimal.Dec
						if rowsC.Next()  {
							rowsC.Scan(&AMOUNT_AFOAG1111GD)
							amountPrima := math.Abs(AMOUNT_AFOAG1111GD.Float64())
							acumulador = acumulador - amountPrima
							cedul = models.LineaDetalle{11,"Menos Plan de Desarrollo AFOAG1111D",amountPrima}
							models.AddLineaDetalle(cedul,"Menos Plan de Desarrollo AFOAG1111D")
						}

						queryCC := `SELECT ISNULL(SUM(AMOUNT),0) as amount
						  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
						  WHERE (SUBSTRING(ACCNT_CODE,1,2) in('80','81','82','83','84','85','86','87','88','89','90','91','92','93','94','95','98') OR
						  SUBSTRING(ACCNT_CODE,1,3) in('978') )
							 AND PERIOD >= `+PERIOD_ANTERIOR+` AND PERIOD < `+PERIOD_INICIO+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10','20') 
						   AND ANAL_T3 = 'AFOAG1111D'`
							
						rowsCC, errCC := conn.Query(queryCC)
						if errCC != nil {
							fmt.Println(queryC)
							panic(errC)
						}
						if rowsCC.Next()  {
							rowsCC.Scan(&AMOUNT_AFOAG1111GD)
							amountPrima := math.Abs(AMOUNT_AFOAG1111GD.Float64())
							acumulador = acumulador - amountPrima
							cedul = models.LineaDetalle{12,"Menos Plan de Desarrollo AFOAG1111D - año pasado",amountPrima}
							models.AddLineaDetalle(cedul,"Menos Plan de Desarrollo AFOAG1111D - año pasado")
						}

						cedul = models.LineaDetalle{13,"Igual a gastos operativos para el cálculo",acumulador}
						models.AddLineaDetalle(cedul,"Igual1")
						pcent := 0.2
						if Compare(beego.AppConfig.String("TipoCampo"),"2")==0 {
							pcent = 0.3
						}
						if Compare(beego.AppConfig.String("TipoCampo"),"1")==0 {
							pcent = 0.2
						}
						numero := int(pcent*100)
						
						cedul = models.LineaDetalle{14,""+ strconv.Itoa(numero)+ "% ",acumulador*pcent}
						models.AddLineaDetalle(cedul,"Igual2")
				} else {
					if Compare(tipo,"4")==0 {//FONDOS ASIGNADOS

						models.ClearLineaDetalle()
						var cedul  models.LineaDetalle

						queryC := `SELECT ISNULL(SUM(AMOUNT),0) as amount
						  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG]
						  WHERE SUBSTRING(ACCNT_CODE,1,1) in('5','6','7','8','9')
							 AND PERIOD < `+PERIOD+` AND ALLOCATION != 'C' AND ANAL_T2 in ('10') 
						   AND SUBSTRING(ANAL_T3,1,2)  in ('AF')`
							
						rowsC, errC := conn.Query(queryC)
						if errC != nil {
							fmt.Println(queryC)
							panic(errC)
						}
						var AMOUNT_FONDOS_ASIGNADOS decimal.Dec
						if rowsC.Next()  {
							rowsC.Scan(&AMOUNT_FONDOS_ASIGNADOS)
							amountPrima := math.Abs(AMOUNT_FONDOS_ASIGNADOS.Float64())
							cedul = models.LineaDetalle{1,"Fondos asignados",amountPrima}
							models.AddLineaDetalle(cedul,"Fondos asignados")
						}
		  		
					}
				}//else fondos asignados
			}
		} 
		example := map[string]interface{}{ "success" : 1,"periodoAnterior":PERIOD_ANTERIOR, "lineas": models.GetAllLineaDetalle() }
		c.Data["json"] = &example
		c.ServeJSON()
	}
}



func (c *GetConfigController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
	

		query := `SELECT idConfig, Empresa, pass, tipoDimension FROM [SUNPLUSADV].[dbo].[zConfig] 
		  WHERE BUNIT = '`+BUNIT.(string)+`'`
		rows, err := conn.Query(query)
		if err != nil {
			fmt.Println(query)
			panic(err)
		}
		var idConfig int
		var Empresa string
		var pass string
		var tipoDimension int
		example := map[string]interface{}{ "success" : 1,"idConfig":0, "Empresa": "", "pass" : "", "tipoDimension": 4}
		if rows.Next()  {
			rows.Scan(&idConfig, &Empresa, &pass, &tipoDimension)
			uDec, errorx := base64.StdEncoding.DecodeString(pass)
			if errorx != nil {
				panic(errorx)
			}
			uDecS := string(uDec) 
			example = map[string]interface{}{ "success" : 1,"idConfig":idConfig, "Empresa": Empresa, "pass" : uDecS, "tipoDimension": tipoDimension}
		}
		c.Data["json"] = &example
		c.ServeJSON()
	}
}
func (c *VeDetalleController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		busqueda := c.GetString("busqueda")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
		currentTime := int64(time.Now().Unix())
		tm := time.Unix(currentTime, 0)
		dateString := tm.String() 
		substring := string(dateString[0:10])	
		anio:=string(substring[0:4])	
		mes:=string(substring[5:7])	
		PERIOD_ACTUAL := anio+"0"+mes
		anioAnterior, _ := strconv.Atoi(anio)
		anioAnterior--
		PERIOD_ANTERIOR := strconv.Itoa(anioAnterior)+"0"+mes
	

		query := `SELECT JRNAL_NO, JRNAL_LINE, AMOUNT, D_C, ACCNT_CODE, DESCRIPTN, PERIOD, JRNAL_SRCE, JRNAL_TYPE, TRANS_DATETIME FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG] 
		  WHERE SUBSTRING(ACCNT_CODE,1,1) in (`+busqueda+`)
		  AND PERIOD >= `+PERIOD_ANTERIOR+` AND PERIOD < `+PERIOD_ACTUAL+` AND ALLOCATION != 'C'
		  order by JRNAL_NO asc, JRNAL_LINE asc`
		rows, err := conn.Query(query)
		if err != nil {
			fmt.Println(query)
			panic(err)
		}
		var JRNAL_NO int
		var JRNAL_LINE int
		var AMOUNT decimal.Dec
		var D_C string
		var ACNT_CODE string
		var DESCRIPTN string
		var PERIOD int
		var JRNAL_SRCE string
		var JRNAL_TYPE string
		var TRANS_DATETIME time.Time
		models.ClearLineaDeUnDiario()
		var cedul  models.LineaDeUnDiario
        for rows.Next()  {
			rows.Scan(&JRNAL_NO, &JRNAL_LINE, &AMOUNT, &D_C, &ACNT_CODE, &DESCRIPTN, &PERIOD, &JRNAL_SRCE, &JRNAL_TYPE, &TRANS_DATETIME)
			amountPrima := math.Abs(AMOUNT.Float64())
			mes := int(TRANS_DATETIME.Month())
			dia :=  TRANS_DATETIME.Day()
			fechaReal := ""
			if mes > 9 && dia > 9 {
				fechaReal = fmt.Sprintf("%d-%d-%d", TRANS_DATETIME.Year(), mes, dia)
			} else {
				if mes > 9 && dia < 10 {
					fechaReal = fmt.Sprintf("%d-%d-0%d", TRANS_DATETIME.Year(), mes, dia)
				} else {
					if mes < 10 && dia > 9 {
						fechaReal = fmt.Sprintf("%d-0%d-%d", TRANS_DATETIME.Year(), mes, dia)
					} else {
						fechaReal = fmt.Sprintf("%d-0%d-0%d", TRANS_DATETIME.Year(), mes, dia)
					}
				}
			}
			cedul = models.LineaDeUnDiario{JRNAL_NO,JRNAL_LINE,amountPrima,D_C,ACNT_CODE, DESCRIPTN, PERIOD, JRNAL_SRCE, JRNAL_TYPE,fechaReal}
			index:= strconv.Itoa(JRNAL_NO)+"-"+ strconv.Itoa(JRNAL_LINE)
			models.AddLineaDeUnDiario(cedul,index)
		}
		example := map[string]interface{}{ "success" : 1,"periodoAnterior":PERIOD_ANTERIOR, "periodoActual": PERIOD_ACTUAL, "lineas": models.GetAllLineaDeUnDiario() }
		c.Data["json"] = &example
		c.ServeJSON()
	}
}
func (c *ListaActivosFijosController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
		query := "SELECT ASSET_CODE,STATUS,ASSET_STATUS,DESCR,START_PERD,END_PERD,LAST_PERD,DISPOSAL_PERD,DISPOSED,BASE_GROSS,BASE_DEP,BASE_NET,BASE_PCENT,TXN_GROSS,TXN_DEP,TXN_NET,TXN_PCENT FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ASSET] order by ASSET_CODE asc"
		rows, err := conn.Query(query)

		var ASSET_CODE string
		var STATUS int
		var ASSET_STATUS int
		var DESCR string
		var START_PERIOD int
		var END_PERIOD int
		var ULTIMO_PERIOD int
		var DISPOSAL_PERIOD int
		var DISPOSED int
		var BASE_GROSS float64
		var BASE_DEP float64
		var BASE_NET float64
		var BASE_PCENT float64
		var TXN_GROSS float64
		var TXN_DEP float64
		var TXN_NET float64
		var TXN_PCENT float64
		models.ClearActivosFijos()
		var cedul  models.ActivoFijo
        for rows.Next()  {
			rows.Scan(&ASSET_CODE, &STATUS, &ASSET_STATUS, &DESCR, &START_PERIOD, &END_PERIOD, &ULTIMO_PERIOD, &DISPOSAL_PERIOD, &DISPOSED, &BASE_GROSS, &BASE_DEP, &BASE_NET, &BASE_PCENT,&TXN_GROSS, &TXN_DEP, &TXN_NET, &TXN_PCENT)
			//fmt.Println(BASE_GROSS.Float64())
		//	BASE_G := math.Abs(BASE_GROSS.Float64())
			cedul = models.ActivoFijo{ASSET_CODE, STATUS, ASSET_STATUS, DESCR, START_PERIOD, END_PERIOD, ULTIMO_PERIOD, DISPOSAL_PERIOD, DISPOSED, BASE_GROSS, BASE_DEP, BASE_NET, BASE_PCENT,TXN_GROSS, TXN_DEP, TXN_NET, TXN_PCENT}
			models.AddActivoFijo(cedul, ASSET_CODE )
		}
		example := map[string]interface{}{"success":1, "activos": models.GetAllActivosFijos() }
		c.Data["json"] = &example
		c.ServeJSON()
	}
	example := map[string]interface{}{ "success":0}
	c.Data["json"] = &example
	c.ServeJSON()
}
func (c *GenerarReporteDeMATController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		delPeriodo := c.GetString("delPeriodo")
		alPeriodo := c.GetString("alPeriodo")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
		query := "SELECT NAME, ANL_CODE FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_CODE] WHERE SUBSTRING(ANL_CODE,1,2) = 'ER' AND LEN(RTRIM(ANL_CODE)) = 9"
		rows, err := conn.Query(query)
		var Nombre string
		var Codigo string
		//var PeriodosList map[string]*PorPeriodo
		models.ClearIglesias()
		var cedul  models.Iglesias
        for rows.Next()  {
			rows.Scan(&Nombre, &Codigo)
			Nombre = strings.TrimSpace(Nombre)
			Codigo = strings.TrimSpace(Codigo)
			cedul = models.Iglesias{Nombre,Codigo,make(map[string]*models.PorPeriodo)}
			models.AddIglesias(cedul, Codigo)
		}
		query2 := `SELECT DISTINCT a.ACCNT_CODE, b.DESCR
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG] a
		  INNER JOIN [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ACNT] b on b.ACNT_CODE = a.ACCNT_CODE
		  WHERE SUBSTRING(ANAL_T6,1,2) = 'ER' AND PERIOD >= `+delPeriodo+` AND PERIOD <= `+alPeriodo+`
		  order by a.ACCNT_CODE asc`
		rows2, err2 := conn.Query(query2)
		if err2 != nil {
			fmt.Println(query2)
			panic(err2)
		}
		var cuentaActual models.Cuenta
	
		var Cuenta string
		var NCuenta string
		models.ClearCuentas()
		for rows2.Next()  {
			rows2.Scan(&Cuenta, &NCuenta)
			Cuenta = strings.TrimSpace(Cuenta)
			NCuenta = strings.TrimSpace(NCuenta)
			cuentaActual = models.Cuenta{Cuenta,NCuenta}
			models.AddCuenta(cuentaActual, Cuenta)
		}
		models.ClearPeriodos()
		var PERIODActual models.Periodo
		

		query1 := `SELECT a.ANAL_T6, a.ANAL_T9, b.DESCR, a.ACCNT_CODE, a.PERIOD, a.AMOUNT
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG] a
		  INNER JOIN [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ACNT] b on b.ACNT_CODE = a.ACCNT_CODE
		  WHERE SUBSTRING(ANAL_T6,1,2) = 'ER' AND PERIOD >= `+delPeriodo+` AND PERIOD <= `+alPeriodo+`
		  order by PERIOD asc, a.ANAL_T3 asc`
		rows1, err1 := conn.Query(query1)
		if err1 != nil {
			fmt.Println(query1)
			panic(err1)
		}
		var ORG_ID string
		var FNCT string
		var DESCR string
		var ACNT_CODE string
		var PERIOD int
		var AMOUNT decimal.Dec
		var iglesiaActual models.Iglesias
		for rows1.Next()  {
			rows1.Scan(&ORG_ID, &FNCT, &DESCR, &ACNT_CODE, &PERIOD, &AMOUNT)
			ORG_ID = strings.TrimSpace(ORG_ID)
			FNCT = strings.TrimSpace(FNCT)
			DESCR = strings.TrimSpace(DESCR)
			ACNT_CODE = strings.TrimSpace(ACNT_CODE)
			amountPrima := AMOUNT.Float64()
			iglesiaActual = models.DameIglesiaPorCodigo(ORG_ID)
			PERIODString := strconv.Itoa(PERIOD)
			PERIODActual = models.Periodo{PERIODString}
			models.AddPeriodo(PERIODActual, PERIODString)
			if _, ok := iglesiaActual.PeriodosList[PERIODString]; ok {
				} else {
				var PERIODM models.PorPeriodo
				PERIODM = models.PorPeriodo{PERIOD, FNCT, make(map[string]*models.PorCuenta) }
				iglesiaActual.PeriodosList[PERIODString] = &PERIODM
			}
			iglesiaActual.PeriodosList[PERIODString].PERIOD = PERIOD
			if Compare(FNCT,"") != 0 {
				iglesiaActual.PeriodosList[PERIODString].Distrito = FNCT
			}
			if _, ok := iglesiaActual.PeriodosList[PERIODString].CuentasList[ACNT_CODE]; ok {
			} else {
				var Cuenta models.PorCuenta
				Cuenta = models.PorCuenta{ACNT_CODE,0.0,DESCR}
				iglesiaActual.PeriodosList[PERIODString].CuentasList[ACNT_CODE] = &Cuenta

			}
			iglesiaActual.PeriodosList[PERIODString].CuentasList[ACNT_CODE].Saldo += amountPrima	
		}
		example := map[string]interface{}{ "success" : 1 , "diezmos": models.GetAllIglesias(), "cuentas" : models.GetAllCuentas(), "periodos" : models.GetAllPeriodos() }
		c.Data["json"] = &example
		c.ServeJSON()
		return
	}
	example := map[string]interface{}{ "success" : 0 }
	c.Data["json"] = &example
	c.ServeJSON()
}
func (c *GenerarReporteDeIglesiasController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		delPeriodo := c.GetString("delPeriodo")
		alPeriodo := c.GetString("alPeriodo")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
		query := "SELECT NAME, ANL_CODE FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_CODE] WHERE SUBSTRING(ANL_CODE,1,1) = 'T'"
		rows, err := conn.Query(query)
		var Nombre string
		var Codigo string
		//var PeriodosList map[string]*PorPeriodo
		models.ClearIglesias()
		var cedul  models.Iglesias
        for rows.Next()  {
			rows.Scan(&Nombre, &Codigo)
			Nombre = strings.TrimSpace(Nombre)
			Codigo = strings.TrimSpace(Codigo)
			cedul = models.Iglesias{Nombre,Codigo,make(map[string]*models.PorPeriodo)}
			models.AddIglesias(cedul, Codigo)
		}
		query2 := `SELECT DISTINCT a.ACCNT_CODE, b.DESCR
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG] a
		  INNER JOIN [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ACNT] b on b.ACNT_CODE = a.ACCNT_CODE
		  WHERE SUBSTRING(ANAL_T5,1,1) = 'T' AND PERIOD >= `+delPeriodo+` AND PERIOD <= `+alPeriodo+`
		  order by a.ACCNT_CODE asc`
		rows2, err2 := conn.Query(query2)
		if err2 != nil {
			fmt.Println(query2)
			panic(err2)
		}
		var cuentaActual models.Cuenta
	
		var Cuenta string
		var NCuenta string
		models.ClearCuentas()
		for rows2.Next()  {
			rows2.Scan(&Cuenta, &NCuenta)
			Cuenta = strings.TrimSpace(Cuenta)
			NCuenta = strings.TrimSpace(NCuenta)
			cuentaActual = models.Cuenta{Cuenta,NCuenta}
			models.AddCuenta(cuentaActual, Cuenta)
		}
		models.ClearPeriodos()
		var PERIODActual models.Periodo
		

		query1 := `SELECT a.ANAL_T5, a.ANAL_T3, b.DESCR, a.ACCNT_CODE, a.PERIOD, a.AMOUNT
		  FROM [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_A_SALFLDG] a
		  INNER JOIN [SUNPLUSADV].[dbo].[`+BUNIT.(string)+`_ACNT] b on b.ACNT_CODE = a.ACCNT_CODE
		  WHERE SUBSTRING(ANAL_T5,1,1) = 'T' AND PERIOD >= `+delPeriodo+` AND PERIOD <= `+alPeriodo+`
		  order by PERIOD asc, a.ANAL_T3 asc, a.ANAL_T5 asc`
		rows1, err1 := conn.Query(query1)
		if err1 != nil {
			fmt.Println(query1)
			panic(err1)
		}
		var ORG_ID string
		var FNCT string
		var DESCR string
		var ACNT_CODE string
		var PERIOD int
		var AMOUNT decimal.Dec
		var iglesiaActual models.Iglesias
		for rows1.Next()  {
			rows1.Scan(&ORG_ID, &FNCT, &DESCR, &ACNT_CODE, &PERIOD, &AMOUNT)
			ORG_ID = strings.TrimSpace(ORG_ID)
			FNCT = strings.TrimSpace(FNCT)
			DESCR = strings.TrimSpace(DESCR)
			ACNT_CODE = strings.TrimSpace(ACNT_CODE)
			amountPrima := AMOUNT.Float64()
			iglesiaActual = models.DameIglesiaPorCodigo(ORG_ID)
			PERIODString := strconv.Itoa(PERIOD)
			PERIODActual = models.Periodo{PERIODString}
			models.AddPeriodo(PERIODActual, PERIODString)
			if _, ok := iglesiaActual.PeriodosList[PERIODString]; ok {
				} else {
				var PERIODM models.PorPeriodo
				PERIODM = models.PorPeriodo{PERIOD, FNCT, make(map[string]*models.PorCuenta) }
				iglesiaActual.PeriodosList[PERIODString] = &PERIODM
			}
			iglesiaActual.PeriodosList[PERIODString].PERIOD = PERIOD
			if Compare(FNCT,"") != 0 {
				iglesiaActual.PeriodosList[PERIODString].Distrito = FNCT
			}
			if _, ok := iglesiaActual.PeriodosList[PERIODString].CuentasList[ACNT_CODE]; ok {
			} else {
				var Cuenta models.PorCuenta
				Cuenta = models.PorCuenta{ACNT_CODE,0.0,DESCR}
				iglesiaActual.PeriodosList[PERIODString].CuentasList[ACNT_CODE] = &Cuenta

			}
			iglesiaActual.PeriodosList[PERIODString].CuentasList[ACNT_CODE].Saldo += amountPrima	
		}
		example := map[string]interface{}{ "success" : 1 , "diezmos": models.GetAllIglesias(), "cuentas" : models.GetAllCuentas(), "periodos" : models.GetAllPeriodos() }
		c.Data["json"] = &example
		c.ServeJSON()
		return
	}
	example := map[string]interface{}{ "success" : 0 }
	c.Data["json"] = &example
	c.ServeJSON()
}
func (c *ListaTiposDeDiarioController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
		query := "SELECT idTipoDeDiario, Codigo, nombre FROM [SUNPLUSADV].[dbo].[TiposDeDiario] WHERE BUNIT = '"+BUNIT.(string)+"' order by idTipoDeDiario asc"
		rows, err := conn.Query(query)
		var idTipoDeDiario int
		var Codigo string
		var Nombre string
		models.ClearTiposDeDiario()
		var cedul  models.TiposDD
        for rows.Next()  {
			rows.Scan(&idTipoDeDiario, &Codigo, &Nombre)
			cedul = models.TiposDD{idTipoDeDiario,Codigo,Nombre}
			models.AddTiposDeDiariop(cedul, strconv.Itoa(idTipoDeDiario) )
		}
		example := map[string]interface{}{ "tiposDeDiario": models.GetAllTiposDeDiario() }
		c.Data["json"] = &example
		c.ServeJSON()
	}
}
func (c *ListaCedulasController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		CedulaList := make(map[string]*Cedula)
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
		query := "SELECT nombre, idCedula FROM [SUNPLUSADV].[dbo].[Cedulas] WHERE BUNIT = '"+BUNIT.(string)+"' order by idCedula asc"
		rows, err := conn.Query(query)
		var nombre1 string
		var idCedula int64
		var cedul Cedula
       for rows.Next()  {
			rows.Scan(&nombre1, &idCedula)
			cedul = Cedula{nombre1,idCedula}
			//hardcode, checar despues!
			nombre := nombre1+"|"+strconv.FormatInt(idCedula,10)
			CedulaList[nombre] = &cedul
		}
		example := map[string]interface{}{ "cedulas": CedulaList }
		c.Data["json"] = &example
		c.ServeJSON()
	}
}
func (c *TipoDeDimensionesController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
		query := "SELECT ANL_CAT_ID, S_HEAD, DESCR FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_CAT] order by ANL_CAT_ID asc"
		rows, err := conn.Query(query)
		var ANL_CAT_ID string
		var S_HEAD string
		var DESCR string
		//err = nil
		var dimen models.Dimensiones
        //_ = cuenta
     	for rows.Next()  {
			rows.Scan(&ANL_CAT_ID, &S_HEAD, &DESCR)
			dimen = models.Dimensiones{ANL_CAT_ID,S_HEAD,DESCR}
			models.AddDimension(dimen, ANL_CAT_ID)
		}
		example := map[string]interface{}{ "dimensiones": models.GetAllDimensiones() }
		c.Data["json"] = &example
		c.ServeJSON()
	}
}
func (c *CambiarUnidadDeNegocioController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
		query := "SELECT BU_CODE FROM [SU_DOMAINDB].[dbo].[DOMN_BU_DSRCE_LINK] order by BU_CODE asc"
		rows, err := conn.Query(query)
		var BU_CODE string
		//err = nil
		var unit models.Unidad
        //_ = cuenta
     	for rows.Next()  {
			rows.Scan(&BU_CODE)
			unit = models.Unidad{BU_CODE}
			models.AddUnidadDeNegocio(unit, BU_CODE)
		}
		example := map[string]interface{}{ "unidades": models.GetAllUnidadDeNegocio(), "BUNIT" : BUNIT.(string)  }
		c.Data["json"] = &example
		c.ServeJSON()
	}
}


func (c *ConfigInicialLineasTiposDeDiarioController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}

		queryCedulas := "SELECT idTipoDeDiario, Codigo, nombre FROM [SUNPLUSADV].[dbo].[TiposDeDiario] WHERE BUNIT = '"+BUNIT.(string)+"' order by idTipoDeDiario asc"
		rowsCedulas, _ := conn.Query(queryCedulas)
		var idTipoDeDiario int
		var Codigo string
		var Nombre string
		models.ClearTiposDeDiario()
		var cedul  models.TiposDD
        for rowsCedulas.Next()  {
			rowsCedulas.Scan(&idTipoDeDiario, &Codigo, &Nombre)
			cedul = models.TiposDD{idTipoDeDiario,Codigo,Nombre}
			models.AddTiposDeDiariop(cedul, strconv.Itoa(idTipoDeDiario) )
		}
		






		query := "SELECT ACNT_CODE, DESCR FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ACNT] order by ACNT_CODE asc"
		rows, err := conn.Query(query)
		var ACNT_CODE string
		var DESCR string
		err = nil
		var cuenta models.Cuenta
        _ = cuenta
     	for rows.Next()  {
			rows.Scan(&ACNT_CODE, &DESCR)
			cuenta := models.Cuenta{ACNT_CODE, DESCR}
			models.AddCuenta(cuenta, ACNT_CODE)
		}
		queryPeriodos := "SELECT DISTINCT PERIOD FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_A_SALFLDG] order by PERIOD asc"
		rowsPeriodos, err := conn.Query(queryPeriodos)
		var PERIOD string
		err = nil
		var periodo models.Periodo
        _ = periodo
        for rowsPeriodos.Next()  {
			rowsPeriodos.Scan(&PERIOD)
			periodo := models.Periodo{PERIOD}
			models.AddPeriodo(periodo, PERIOD)
		}

		queryDimensiones := "SELECT ANL_CAT_ID, ENTRY_NUM, S_HEAD FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_ENT_DEFN] WHERE ANL_ENT_ID = 901 order by ENTRY_NUM asc"
		rowsDimensiones, errDimensiones := conn.Query(queryDimensiones)
		if errDimensiones != nil {
			log.Fatal("que paso:", errDimensiones.Error())
		}
		var ANL_CAT_ID string
		var ENTRY_NUM int64

		var STATUS int64
		var ANL_CODE string
		var S_HEAD string

		var S_HEAD1 string
		var S_HEAD2 string
		var S_HEAD3 string
		var S_HEAD4 string
		var S_HEAD5 string
		var S_HEAD6 string
		var S_HEAD7 string
		var S_HEAD8 string
		var S_HEAD9 string
		var S_HEAD10 string
		
		var PROHIBIT_POSTING int64
		var NAME string
		var dimen models.DimensionLite
        
		for rowsDimensiones.Next()  {
			rowsDimensiones.Scan(&ANL_CAT_ID, &ENTRY_NUM, &S_HEAD)
			queryD := "SELECT c.ANL_CODE, c.NAME, c.STATUS, c.PROHIBIT_POSTING, a.DESCR FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_CODE] c INNER JOIN [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_CAT] a on a.ANL_CAT_ID = c.ANL_CAT_ID WHERE c.ANL_CAT_ID = '"+ANL_CAT_ID+"' order by c.ANL_CODE asc"
			rowsD, errD := conn.Query(queryD)
			if errD != nil {
				log.Fatal("que paso:", errD.Error())
			}
			for rowsD.Next()  {
				rowsD.Scan(&ANL_CODE, &NAME, &STATUS, &PROHIBIT_POSTING, &DESCR)
				dimen = models.DimensionLite{ANL_CAT_ID,ANL_CODE,ENTRY_NUM, STATUS, PROHIBIT_POSTING,NAME, DESCR}
				if ENTRY_NUM == 1 {
					models.AddDimensionU1(dimen,ANL_CODE);
					S_HEAD1 = S_HEAD
				}
				if ENTRY_NUM == 2 {
					models.AddDimensionU2(dimen,ANL_CODE);
					S_HEAD2 = S_HEAD
				}
				if ENTRY_NUM == 3 {
					models.AddDimensionU3(dimen,ANL_CODE);
					S_HEAD3 = S_HEAD
				}
				if ENTRY_NUM == 4 {
					models.AddDimensionU4(dimen,ANL_CODE);
					S_HEAD4 = S_HEAD
				}
				if ENTRY_NUM == 5 {
					models.AddDimensionU5(dimen,ANL_CODE);
					S_HEAD5 = S_HEAD
				}
				if ENTRY_NUM == 6 {
					models.AddDimensionU6(dimen,ANL_CODE);
					S_HEAD6 = S_HEAD
				}
				if ENTRY_NUM == 7 {
					models.AddDimensionU7(dimen,ANL_CODE);
					S_HEAD7 = S_HEAD
				}
				if ENTRY_NUM == 8 {
					models.AddDimensionU8(dimen,ANL_CODE);
					S_HEAD8 = S_HEAD
				}
				if ENTRY_NUM == 9 {
					models.AddDimensionU9(dimen,ANL_CODE);
					S_HEAD9 = S_HEAD
				}
				if ENTRY_NUM == 10 {
					models.AddDimensionU10(dimen,ANL_CODE);
					S_HEAD10 = S_HEAD
				}
			}
		}
		currentTime := int64(time.Now().Unix())
		example := map[string]interface{}{ "time" : currentTime, "D1" : S_HEAD1,"D2" : S_HEAD2, "D3" : S_HEAD3, "D4" : S_HEAD4,"D5" : S_HEAD5,"D6" : S_HEAD6,"D7" : S_HEAD7,"D8" : S_HEAD8,"D9" : S_HEAD9,"D10" : S_HEAD10, "tiposDeDiario": models.GetAllTiposDeDiario(), "cuentas": models.GetAllCuentas(), "periodos": models.GetAllPeriodos(), "ANAL_T0": models.GetAllDimension1(), "ANAL_T1": models.GetAllDimension2(), "ANAL_T2": models.GetAllDimension3(), "ANAL_T3": models.GetAllDimension4(), "ANAL_T4": models.GetAllDimension5(), "ANAL_T5": models.GetAllDimension6(), "ANAL_T6": models.GetAllDimension7(), "ANAL_T7": models.GetAllDimension8(), "ANAL_T8": models.GetAllDimension9(), "ANAL_T9": models.GetAllDimension10()}
		c.Data["json"] = &example
		c.ServeJSON()	
	}					
}	
func (c *ConfigInicialLineasCedulasController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}

		CedulaList := make(map[string]*Cedula)
		queryCedulas := "SELECT nombre, idCedula FROM [SUNPLUSADV].[dbo].[Cedulas] WHERE BUNIT = '"+BUNIT.(string)+"' order by idCedula asc"
		rowsCedulas, _ := conn.Query(queryCedulas)
		var nombre1 string
		var idCedula int64
		var cedul Cedula
       	for rowsCedulas.Next()  {
			rowsCedulas.Scan(&nombre1, &idCedula)
			cedul = Cedula{nombre1,idCedula}
			//hardcode, checar despues!
			nombre := nombre1+"|"+strconv.FormatInt(idCedula,10)
			CedulaList[nombre] = &cedul
		}
		






		query := "SELECT ACNT_CODE, DESCR FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ACNT] order by ACNT_CODE asc"
		rows, err := conn.Query(query)
		var ACNT_CODE string
		var DESCR string
		err = nil
		var cuenta models.Cuenta
        _ = cuenta
     	for rows.Next()  {
			rows.Scan(&ACNT_CODE, &DESCR)
			cuenta := models.Cuenta{ACNT_CODE, DESCR}
			models.AddCuenta(cuenta, ACNT_CODE)
		}
		queryPeriodos := "SELECT DISTINCT PERIOD FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_A_SALFLDG] order by PERIOD asc"
		rowsPeriodos, err := conn.Query(queryPeriodos)
		var PERIOD string
		err = nil
		var periodo models.Periodo
        _ = periodo
        for rowsPeriodos.Next()  {
			rowsPeriodos.Scan(&PERIOD)
			periodo := models.Periodo{PERIOD}
			models.AddPeriodo(periodo, PERIOD)
		}

		queryDimensiones := "SELECT ANL_CAT_ID, ENTRY_NUM, S_HEAD FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_ENT_DEFN] WHERE ANL_ENT_ID = 901 order by ENTRY_NUM asc"
		rowsDimensiones, errDimensiones := conn.Query(queryDimensiones)
		if errDimensiones != nil {
			log.Fatal("que paso:", errDimensiones.Error())
		}
		var ANL_CAT_ID string
		var ENTRY_NUM int64

		var STATUS int64
		var ANL_CODE string
		var S_HEAD string

		var S_HEAD1 string
		var S_HEAD2 string
		var S_HEAD3 string
		var S_HEAD4 string
		var S_HEAD5 string
		var S_HEAD6 string
		var S_HEAD7 string
		var S_HEAD8 string
		var S_HEAD9 string
		var S_HEAD10 string
		
		var PROHIBIT_POSTING int64
		var NAME string
		var dimen models.DimensionLite
        
		for rowsDimensiones.Next()  {
			rowsDimensiones.Scan(&ANL_CAT_ID, &ENTRY_NUM, &S_HEAD)
			queryD := "SELECT c.ANL_CODE, c.NAME, c.STATUS, c.PROHIBIT_POSTING, a.DESCR FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_CODE] c INNER JOIN [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_CAT] a on a.ANL_CAT_ID = c.ANL_CAT_ID WHERE c.ANL_CAT_ID = '"+ANL_CAT_ID+"' order by c.ANL_CODE asc"
			rowsD, errD := conn.Query(queryD)
			if errD != nil {
				log.Fatal("que paso:", errD.Error())
			}
			for rowsD.Next()  {
				rowsD.Scan(&ANL_CODE, &NAME, &STATUS, &PROHIBIT_POSTING, &DESCR)
				dimen = models.DimensionLite{ANL_CAT_ID,ANL_CODE,ENTRY_NUM, STATUS, PROHIBIT_POSTING,NAME, DESCR}
				if ENTRY_NUM == 1 {
					models.AddDimensionU1(dimen,ANL_CODE);
					S_HEAD1 = S_HEAD
				}
				if ENTRY_NUM == 2 {
					models.AddDimensionU2(dimen,ANL_CODE);
					S_HEAD2 = S_HEAD
				}
				if ENTRY_NUM == 3 {
					models.AddDimensionU3(dimen,ANL_CODE);
					S_HEAD3 = S_HEAD
				}
				if ENTRY_NUM == 4 {
					models.AddDimensionU4(dimen,ANL_CODE);
					S_HEAD4 = S_HEAD
				}
				if ENTRY_NUM == 5 {
					models.AddDimensionU5(dimen,ANL_CODE);
					S_HEAD5 = S_HEAD
				}
				if ENTRY_NUM == 6 {
					models.AddDimensionU6(dimen,ANL_CODE);
					S_HEAD6 = S_HEAD
				}
				if ENTRY_NUM == 7 {
					models.AddDimensionU7(dimen,ANL_CODE);
					S_HEAD7 = S_HEAD
				}
				if ENTRY_NUM == 8 {
					models.AddDimensionU8(dimen,ANL_CODE);
					S_HEAD8 = S_HEAD
				}
				if ENTRY_NUM == 9 {
					models.AddDimensionU9(dimen,ANL_CODE);
					S_HEAD9 = S_HEAD
				}
				if ENTRY_NUM == 10 {
					models.AddDimensionU10(dimen,ANL_CODE);
					S_HEAD10 = S_HEAD
				}
			}
		}

		example := map[string]interface{}{ "D1" : S_HEAD1,"D2" : S_HEAD2, "D3" : S_HEAD3, "D4" : S_HEAD4,"D5" : S_HEAD5,"D6" : S_HEAD6,"D7" : S_HEAD7,"D8" : S_HEAD8,"D9" : S_HEAD9,"D10" : S_HEAD10, "cedulas": CedulaList, "cuentas": models.GetAllCuentas(), "periodos": models.GetAllPeriodos(), "ANAL_T0": models.GetAllDimension1(), "ANAL_T1": models.GetAllDimension2(), "ANAL_T2": models.GetAllDimension3(), "ANAL_T3": models.GetAllDimension4(), "ANAL_T4": models.GetAllDimension5(), "ANAL_T5": models.GetAllDimension6(), "ANAL_T6": models.GetAllDimension7(), "ANAL_T7": models.GetAllDimension8(), "ANAL_T8": models.GetAllDimension9(), "ANAL_T9": models.GetAllDimension10()}
		c.Data["json"] = &example
		c.ServeJSON()	
	}					
}	
func (c *EstadosDeCuentaController) Post() {
	alias := c.GetSession("alias")
	if alias == nil{
		return
	}
	tipoDeUsuario := c.GetSession("tipoDeUsuario")
	if tienePermisosContador(tipoDeUsuario.(int)) {
		BUNIT := c.GetSession("BUNIT")
		connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
		conn, err := sql.Open("mssql", connString2)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
		query := "SELECT ACNT_CODE, DESCR FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ACNT] order by ACNT_CODE asc"
		rows, err := conn.Query(query)
		var ACNT_CODE string
		var DESCR string
		err = nil
		var cuenta models.Cuenta
        _ = cuenta
     	for rows.Next()  {
			rows.Scan(&ACNT_CODE, &DESCR)
			cuenta := models.Cuenta{ACNT_CODE, DESCR}
			models.AddCuenta(cuenta, ACNT_CODE)
		}
		queryPeriodos := "SELECT DISTINCT PERIOD FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_A_SALFLDG] order by PERIOD asc"
		rowsPeriodos, err := conn.Query(queryPeriodos)
		var PERIOD string
		err = nil
		var periodo models.Periodo
        _ = periodo
        for rowsPeriodos.Next()  {
			rowsPeriodos.Scan(&PERIOD)
			periodo := models.Periodo{PERIOD}
			models.AddPeriodo(periodo, PERIOD)
		}

		queryDimensiones := "SELECT ANL_CAT_ID, ENTRY_NUM FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_ENT_DEFN] WHERE ANL_ENT_ID = 901 order by ENTRY_NUM asc"
		rowsDimensiones, errDimensiones := conn.Query(queryDimensiones)
		if errDimensiones != nil {
			log.Fatal("que paso:", errDimensiones.Error())
		}
		var ANL_CAT_ID string
		var ENTRY_NUM int64

		var STATUS int64
		var ANL_CODE string
		var PROHIBIT_POSTING int64
		var NAME string
		var dimen models.DimensionLite
        
		for rowsDimensiones.Next()  {
			rowsDimensiones.Scan(&ANL_CAT_ID, &ENTRY_NUM)
			queryD := "SELECT c.ANL_CODE, c.NAME, c.STATUS, c.PROHIBIT_POSTING, a.DESCR FROM [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_CODE] c INNER JOIN [SUNPLUSADV].[dbo].["+BUNIT.(string)+"_ANL_CAT] a on a.ANL_CAT_ID = c.ANL_CAT_ID WHERE c.ANL_CAT_ID = '"+ANL_CAT_ID+"' order by c.ANL_CODE asc"
			rowsD, errD := conn.Query(queryD)
			if errD != nil {
				log.Fatal("que paso:", errD.Error())
			}
			for rowsD.Next()  {
				rowsD.Scan(&ANL_CODE, &NAME, &STATUS, &PROHIBIT_POSTING, &DESCR)
				dimen = models.DimensionLite{ANL_CAT_ID,ANL_CODE,ENTRY_NUM, STATUS, PROHIBIT_POSTING,NAME, DESCR}
				if ENTRY_NUM == 1 {
					models.AddDimensionU1(dimen,ANL_CODE);
				}
				if ENTRY_NUM == 2 {
					models.AddDimensionU2(dimen,ANL_CODE);
				}
				if ENTRY_NUM == 3 {
					models.AddDimensionU3(dimen,ANL_CODE);
				}
				if ENTRY_NUM == 4 {
					models.AddDimensionU4(dimen,ANL_CODE);
				}
				if ENTRY_NUM == 5 {
					models.AddDimensionU5(dimen,ANL_CODE);
				}
				if ENTRY_NUM == 6 {
					models.AddDimensionU6(dimen,ANL_CODE);
				}
				if ENTRY_NUM == 7 {
					models.AddDimensionU7(dimen,ANL_CODE);
				}
				if ENTRY_NUM == 8 {
					models.AddDimensionU8(dimen,ANL_CODE);
				}
				if ENTRY_NUM == 9 {
					models.AddDimensionU9(dimen,ANL_CODE);
				}
				if ENTRY_NUM == 10 {
					models.AddDimensionU10(dimen,ANL_CODE);
				}
			}
		}

		example := map[string]interface{}{ "cuentas": models.GetAllCuentas(), "periodos": models.GetAllPeriodos(), "ANAL_T0": models.GetAllDimension1(), "ANAL_T1": models.GetAllDimension2(), "ANAL_T2": models.GetAllDimension3(), "ANAL_T3": models.GetAllDimension4(), "ANAL_T4": models.GetAllDimension5(), "ANAL_T5": models.GetAllDimension6(), "ANAL_T6": models.GetAllDimension7(), "ANAL_T7": models.GetAllDimension8(), "ANAL_T8": models.GetAllDimension9(), "ANAL_T9": models.GetAllDimension10()}
		c.Data["json"] = &example
		c.ServeJSON()	
	}					
}	
func (c *LoginController) Post() {
	correo := c.GetString("correo")
	password := c.GetString("password")
	c.Data["correo"] = correo
	c.Data["password"] = password
	passwordPrimo := []byte(password)
	hasher := sha512.New()
    hasher.Write(passwordPrimo)
	cryptoText := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	flag.Parse() // parse the command line args
	connString2 := "Database="+beego.AppConfig.String("mssqldb")+";Data Source="+beego.AppConfig.String("mssqlurls")+";Integrated Security=False;User ID="+beego.AppConfig.String("mssqluser")+";Password="+beego.AppConfig.String("mssqlpass")+";connect timeout = 1000; encrypt=disable;";
	conn, err := sql.Open("mssql", connString2)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()
	stmt, err := conn.Prepare("SELECT TOP 1 idUsuario, usuario, pass,activo, alias, BUNIT, tipoDeUsuario FROM [SUNPLUSADV].[dbo].[users] WHERE usuario = '"+correo+"'")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()
	row := stmt.QueryRow()
	var activo int
	var idUsuario int
	var usuario string
	var pass string
	var alias string
	var BUNIT string
	var tipoDeUsuario int
	err = row.Scan(&idUsuario, &usuario, &pass, &activo,&alias,&BUNIT,&tipoDeUsuario)
	if err != nil {
		//log.Fatal("Scan failed:", err.Error())
		c.Data["mensaje"] = "contraseña incorrecta bd"
		c.TplName = "index.tpl"
	}
		fmt.Println("lo que escribiste:"+cryptoText)
		
fmt.Println("bd: "+pass)
		
	if Compare(cryptoText,pass)==0{
		fmt.Println("password CORRECTO")
	    c.SetSession("tipoDeUsuario",tipoDeUsuario)
	    c.SetSession("alias", alias)
	    c.SetSession("usuario", usuario)
	    c.SetSession("idUsuario", idUsuario)
	    c.SetSession("BUNIT", BUNIT)
	    c.Data["alias"] = alias
	    c.Data["BUNIT"] = BUNIT
	  	ayudame := "menu"+strconv.Itoa(tipoDeUsuario)+".tpl"
		c.Layout = "seven.tpl"
	    c.TplName = "vacio.tpl"
	    c.LayoutSections = make(map[string]string)
	    c.LayoutSections["menu"] = ayudame
	    c.LayoutSections["contenido"] = "nada.tpl"
	}else{
		c.Data["mensaje"] = "contraseña incorrecta"
		c.TplName = "index.tpl"
		fmt.Println("password incorrecto")
	 }
}

func Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}