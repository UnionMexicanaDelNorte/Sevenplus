<!DOCTYPE html>
<html lang="en">

<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
  <!-- Meta, title, CSS, favicons, etc. -->
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">

  <title>Sevenplus </title>

  <!-- Bootstrap core CSS -->

  <link href="static/css/bootstrap.min.css" rel="stylesheet">
  <link rel="shortcut icon" href="static/favicon.ico">
  <link href="static/fonts/css/font-awesome.min.css" rel="stylesheet">
  <link href="static/css/animate.min.css" rel="stylesheet">

  <!-- Custom styling plus plugins -->
  <link href="static/css/custom.css" rel="stylesheet">
  <link rel="stylesheet" type="text/css" href="static/css/maps/jquery-jvectormap-2.0.3.css" />
  <link href="static/css/icheck/flat/green.css" rel="stylesheet" />
  <link href="static/css/floatexamples.css" rel="stylesheet" type="text/css" />

  <script src="static/js/jquery-2.2.4.min.js"></script>
  <script src="static/js/autocomplete/jquery.autocomplete.js"></script>
  
  <script src="static/js/nprogress.js"></script>

  <!--[if lt IE 9]>
        <script src="../assets/js/ie8-responsive-file-warning.js"></script>
        <![endif]-->

  <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
  <!--[if lt IE 9]>
          <script src="https://oss.maxcdn.com/html5shiv/3.7.2/html5shiv.min.js"></script>
          <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
        <![endif]-->

</head>


<body class="nav-md">

  <div class="container body">


    <div class="main_container">

      <div class="col-md-3 left_col">
        <div class="left_col scroll-view">

          <div class="navbar nav_title" style="border: 0;">
            <a href="" class="site_title"><span>Sevenplus</span></a>
             <style>
  .speech {border: 0px solid #DDD; width: 300px; padding: 0; margin: 0}
  .speech input {border: 0; width: 240px; display: inline-block; height: 30px;}
  .speech img {float: left; width: 40px }
</style>
 
<!-- Search Form -->
<form id="labnol" method="get" action="https://www.google.com/search">
  <div id="transcript2"></div>
  <div class="speech">
    <input type="text" name="q" id="transcript" style="display:none;" placeholder="Speak" />
    <img  style="display:none;" src="static/img/micro.png" />
  </div>
</form>
 

          </div>
          <div class="clearfix"></div>

          <!-- menu prile quick info -->
          <div class="profile">
            <div class="profile_pic">
              <img onclick="startDictation()" src="static/img/micro.png" alt="..." class="img-circle profile_img">
            </div>
            <div class="profile_info">
              <span>Bienvenido,</span>
              <h2>{{.alias}}, est&aacute;s en: <span id="BUNITACTUAL">{{.BUNIT}}</span></h2>
            </div>
          </div>
          <!-- /menu prile quick info -->

          <br />
{{.LayoutContent}}
          <!-- sidebar menu -->
          <div id="sidebar-menu" class="main_menu_side hidden-print main_menu">

            <div class="menu_section">
              {{.menu}}
            </div>
          </div>
          <!-- /sidebar menu -->
        </div>
      </div>

      <!-- top navigation -->
      <div class="top_nav">

        <div class="nav_menu">
          <nav class="" role="navigation">
            <div class="nav toggle">
              <a id="menu_toggle"><i class="fa fa-bars"></i></a>
            </div>

            <ul class="nav navbar-nav navbar-right">
              <li class="">
                <a href="javascript:;" class="user-profile dropdown-toggle" data-toggle="dropdown" aria-expanded="false">
                  <img src="static/images/img.jpg" alt="">{{.alias}}
                  <span class=" fa fa-angle-down"></span>
                </a>
                <ul class="dropdown-menu dropdown-usermenu animated fadeInDown pull-right">
                  <li><a href="javascript:;">  Profile</a>
                  </li>
                  <li>
                    <a href="javascript:;">
                      <span class="badge bg-red pull-right">50%</span>
                      <span>Settings</span>
                    </a>
                  </li>
                  <li>
                    <a href="javascript:;">Help</a>
                  </li>
                  <li><a href="login.html"><i class="fa fa-sign-out pull-right"></i> Log Out</a>
                  </li>
                </ul>
              </li>

              <li role="presentation" class="dropdown">
                <a href="javascript:;" class="dropdown-toggle info-number" data-toggle="dropdown" aria-expanded="false">
                  <i class="fa fa-envelope-o"></i>
                  <span class="badge bg-green">6</span>
                </a>
                <ul id="menu1" class="dropdown-menu list-unstyled msg_list animated fadeInDown" role="menu">
                  <li>
                    <a>
                      <span class="image">
                                        <img src="static/images/img.jpg" alt="Profile Image" />
                                    </span>
                      <span>
                                        <span>John Smith</span>
                      <span class="time">3 mins ago</span>
                      </span>
                      <span class="message">
                                        Film festivals used to be do-or-die moments for movie makers. They were where...
                                    </span>
                    </a>
                  </li>
                  <li>
                    <a>
                      <span class="image">
                                        <img src="static/images/img.jpg" alt="Profile Image" />
                                    </span>
                      <span>
                                        <span>John Smith</span>
                      <span class="time">3 mins ago</span>
                      </span>
                      <span class="message">
                                        Film festivals used to be do-or-die moments for movie makers. They were where...
                                    </span>
                    </a>
                  </li>
                  <li>
                    <a>
                      <span class="image">
                                        <img src="static/images/img.jpg" alt="Profile Image" />
                                    </span>
                      <span>
                                        <span>John Smith</span>
                      <span class="time">3 mins ago</span>
                      </span>
                      <span class="message">
                                        Film festivals used to be do-or-die moments for movie makers. They were where...
                                    </span>
                    </a>
                  </li>
                  <li>
                    <a>
                      <span class="image">
                                        <img src="static/images/img.jpg" alt="Profile Image" />
                                    </span>
                      <span>
                                        <span>John Smith</span>
                      <span class="time">3 mins ago</span>
                      </span>
                      <span class="message">
                                        Film festivals used to be do-or-die moments for movie makers. They were where...
                                    </span>
                    </a>
                  </li>
                  <li>
                    <div class="text-center">
                      <a href="inbox.html">
                        <strong>See All Alerts</strong>
                        <i class="fa fa-angle-right"></i>
                      </a>
                    </div>
                  </li>
                </ul>
              </li>

            </ul>
          </nav>
        </div>

      </div>
      <!-- /top navigation -->


      <!-- page content -->
      <div id="mainContent" class="right_col" role="main">

       
       {{.contenido}}
        <!-- footer content -->

        <footer>
          <div class="copyright-info">
            <p class="pull-right">Sevenplus</a>  
            </p>
          </div>
          <div class="clearfix"></div>
        </footer>
        <!-- /footer content -->
      </div>
      <!-- /page content -->

    </div>

  </div>

  <div id="custom_notifications" class="custom-notifications dsp_none">
    <ul class="list-unstyled notifications clearfix" data-tabbed_notifications="notif-group">
    </ul>
    <div class="clearfix"></div>
    <div id="notif-group" class="tabbed_notifications"></div>
  </div>

  <script src="static/js/bootstrap.min.js"></script>
  <div id="audios"></div>
 
  <!-- gauge js -->
  <script type="text/javascript" src="static/js/gauge/gauge.min.js"></script>
  <!--script type="text/javascript" src="static/js/gauge/gauge_demo.js"></script-->
  <!-- bootstrap progress js -->
  <script src="static/js/progressbar/bootstrap-progressbar.min.js"></script>
  <script src="static/js/nicescroll/jquery.nicescroll.min.js"></script>
  <!-- icheck -->
  <script src="static/js/icheck/icheck.min.js"></script>
  <!-- daterangepicker -->
  <script type="text/javascript" src="static/js/moment/moment.min.js"></script>
  <script type="text/javascript" src="static/js/datepicker/daterangepicker.js"></script>
  <!-- chart js -->
  <script src="static/js/chartjs/chart.min.js"></script>
  <script src="static/js/echart/echarts-all.js"></script>
  <script src="static/js/jquery.maskMoney.min.js"></script>
  <script src="static/js/custom.js"></script>

  <!-- flot js -->
  <!--[if lte IE 8]><script type="text/javascript" src="static/js/excanvas.min.js"></script><![endif]-->
  <script type="text/javascript" src="static/js/flot/jquery.flot.js"></script>
  <script type="text/javascript" src="static/js/flot/jquery.flot.pie.js"></script>
  <script type="text/javascript" src="static/js/flot/jquery.flot.orderBars.js"></script>
  <script type="text/javascript" src="static/js/flot/jquery.flot.time.min.js"></script>
  <script type="text/javascript" src="static/js/flot/date.js"></script>
  <script type="text/javascript" src="static/js/flot/jquery.flot.spline.js"></script>
  <script type="text/javascript" src="static/js/flot/jquery.flot.stack.js"></script>
  <script type="text/javascript" src="static/js/flot/curvedLines.js"></script>
  <script type="text/javascript" src="static/js/flot/jquery.flot.resize.js"></script>
  
  <!-- worldmap -->
  <script type="text/javascript" src="static/js/maps/jquery-jvectormap-2.0.3.min.js"></script>
  <script type="text/javascript" src="static/js/maps/gdp-data.js"></script>
  <script type="text/javascript" src="static/js/maps/jquery-jvectormap-world-mill-en.js"></script>
  <script type="text/javascript" src="static/js/maps/jquery-jvectormap-us-aea-en.js"></script>
  <!-- pace -->
  <script src="static/js/pace/pace.min.js"></script>
  <script src="static/js/papaparse.min.js"></script>
  <script src="static/js/sha512.min.js"></script>
    

  <!-- /dashbord linegraph -->
  <!-- datepicker -->
  <script type="text/javascript">
  var xhrPool = [];
  $(document).ajaxSend(function(e, jqXHR, options){
    xhrPool.push(jqXHR);
  });
  $(document).ajaxComplete(function(e, jqXHR, options) {
    xhrPool = $.grep(xhrPool, function(x){return x!=jqXHR});
  });
  var abort = function() {
    $.each(xhrPool, function(idx, jqXHR) {
      jqXHR.abort();
    });
  };

  var oldbeforeunload = window.onbeforeunload;
  window.onbeforeunload = function() {
    var r = oldbeforeunload ? oldbeforeunload() : undefined;
    if (r == undefined) {
      // only cancel requests if there is no prompt to stay on the page
      // if there is a prompt, it will likely give the requests enough time to finish
      abort();
    }
    return r;
  }




  var palabrasComunes = ["hello","dame","damela","hola",];
 function startDictation() {
  document.getElementById('transcript2').innerHTML='';
    if (window.hasOwnProperty('webkitSpeechRecognition')) {
 
      var recognition = new webkitSpeechRecognition();
 
      recognition.continuous = false;
      recognition.interimResults = false;
 
      recognition.lang = "es-MX";
      recognition.start();
 
      recognition.onresult = function(e) {
        document.getElementById('transcript').value
                                 = e.results[0][0].transcript;
        document.getElementById('transcript2').innerHTML
                                 = e.results[0][0].transcript;
        recognition.stop();
        var texto =  document.getElementById('transcript2').innerHTML;
        if(texto.indexOf("cuenta")>-1)
        {
          if(texto.indexOf("estado")>-1)//estado de cuenta
          {
            var palabras = texto.split(' ');
            //pensando que podría ser la ultima y/o penultima
            var ultima = palabras[palabras.length-1];
           // var cuenta = '';
            var penultima = palabras[palabras.length-2];
            /*if(ultima.length==6)
            {
              cuenta = ultima;
            }
            else
            {
              if(ultima.length==3 && penultima.length==3)
              {
                cuenta=penultima+ultima;
              }
            }
            if(cuenta.length>0)
            {*/
              $("#audios").html('<audio src="static/siri/generandoEstadoDeCuenta.m4a" autoplay="autoplay"></audio>');
              var delPeriodo = 2000001;
              var alPeriodo = 2100012;
              window.open(
                "/generarEstadoDeCuenta?cuenta=" + penultima.toUpperCase()+" "+ultima.toUpperCase()+"&delPeriodo="+delPeriodo+"&alPeriodo="+alPeriodo+"&busca=1",
                '_blank'
              );
            //}

          }  
        }
        else
        {
          if(texto.indexOf("poliza")>-1 || texto.indexOf("diario")>-1)
          {
            var palabras = texto.split(' ');
            var i;
            for(i=0;i<palabras.length;i++)
            {
              var talvesNumero = palabras[i];
              if(!isNaN(talvesNumero))
              {
                $("#audios").html('<audio src="static/siri/generandoDiario.m4a" autoplay="autoplay"></audio>');
                window.open(
                  "/generarDiario?diario=" + talvesNumero,
                  '_blank'
                );
                break;
              }
            }
          }
        }
       // document.getElementById('labnol').submit();
      };
 
      recognition.onerror = function(e) {
        recognition.stop();
      }
 
    }
  }
  

  $(document).on("change", "#cedulasSelect", function (){
    actualizaLineaCedulas();
  });
  var theme = {
          color: [
              "#26B99A", "#34495E", "#BDC3C7", "#3498DB",
              "#9B59B6", "#8abb6f", "#759c6a", "#bfd3b7"
          ],
          title: {
              itemGap: 8,
              textStyle: {
                  fontWeight: "normal",
                  color: "#408829"
              }
          },
          dataRange: {
              color: ["#1f610a", "#97b58d"]
          },
          toolbox: {
              color: ["#408829", "#408829", "#408829", "#408829"]
          },
          tooltip: {
              backgroundColor: "rgba(0,0,0,0.5)",
              axisPointer: {
                  type: "line",
                  lineStyle: {
                      color: "#408829",
                      type: "dashed"
                  },
                  crossStyle: {
                      color: "#408829"
                  },
                  shadowStyle: {
                      color: "rgba(200,200,200,0.3)"
                  }
              }
          },
          dataZoom: {
              dataBackgroundColor: '#eee',
              fillerColor: "rgba(64,136,41,0.2)",
              handleColor: "#408829"
          },
          grid: {
              borderWidth: 0
          },
          categoryAxis: {
              axisLine: {
                  lineStyle: {
                      color: "#408829"
                  }
              },
              splitLine: {
                  lineStyle: {
                      color: ["#eee"]
                  }
              }
          },
          valueAxis: {
              axisLine: {
                  lineStyle: {
                      color: "#408829"
                  }
              },
              splitArea: {
                  show: true,
                  areaStyle: {
                      color: ["rgba(250,250,250,0.1)", "rgba(200,200,200,0.1)"]
                  }
              },
              splitLine: {
                  lineStyle: {
                      color: ["#eee"]
                  }
              }
          },
          timeline: {
              lineStyle: {
                  color: "#408829"
              },
              controlStyle: {
                  normal: {color: "#408829"},
                  emphasis: {color: "#408829"}
              }
          },
          k: {
              itemStyle: {
                  normal: {
                      color: "#68a54a",
                      color0: "#a9cba2",
                      lineStyle: {
                          width: 1,
                          color: "#408829",
                          color0: "#86b379"
                      }
                  }
              }
          },
          map: {
              itemStyle: {
                  normal: {
                      areaStyle: {
                          color: "#ddd"
                      },
                      label: {
                          textStyle: {
                              color: "#c12e34"
                          }
                      }
                  },
                  emphasis: {
                      areaStyle: {
                          color: "#99d2dd"
                      },
                      label: {
                          textStyle: {
                              color: "#c12e34"
                          }
                      }
                  }
              }
          },
          force: {
              itemStyle: {
                  normal: {
                      linkStyle: {
                          strokeColor: "#408829"
                      }
                  }
              }
          },
          chord: {
              padding: 4,
              itemStyle: {
                  normal: {
                      lineStyle: {
                          width: 1,
                          color: "rgba(128, 128, 128, 0.5)"
                      },
                      chordStyle: {
                          lineStyle: {
                              width: 1,
                              color: "rgba(128, 128, 128, 0.5)"
                          }
                      }
                  },
                  emphasis: {
                      lineStyle: {
                          width: 1,
                          color: "rgba(128, 128, 128, 0.5)"
                      },
                      chordStyle: {
                          lineStyle: {
                              width: 1,
                              color: "rgba(128, 128, 128, 0.5)"
                          }
                      }
                  }
              }
          },
          gauge: {
              startAngle: 225,
              endAngle: -45,
              axisLine: {
                  show: true,
                  lineStyle: {
                      color: [[0.2, "#86b379"], [0.8, "#68a54a"], [1, "#408829"]],
                      width: 8
                  }
              },
              axisTick: {
                  splitNumber: 10,
                  length: 12,
                  lineStyle: {
                      color: "auto"
                  }
              },
              axisLabel: {
                  textStyle: {
                      color: "auto"
                  }
              },
              splitLine: {
                  length: 18,
                  lineStyle: {
                      color: "auto"
                  }
              },
              pointer: {
                  length: "90%",
                  color: "auto"
              },
              title: {
                  textStyle: {
                      color: "#333"
                  }
              },
              detail: {
                  textStyle: {
                      color: "auto"
                  }
              }
          },
          textStyle: {
              fontFamily: "Arial, Verdana, sans-serif"
          }
      };
  

  $(document).on("change", "#tipoDeDiarioSelect", function (){
    actualizaLaTablaDondeEstanLasLineasDelTipoDeDiario();
  });
  $(document).on("change", "#conceptoSelect", function (){
    actualizaLaTablaDondeEstanLasLineasDeLasCedulaSegunElConcepto();
  });

  $(document).on("change", "#cedulasSelectConceptos", function (){
    actualizaConceptosCedulas();
  });
  
  $(document).on("click", ".cancelaAjax", function (){
    abort();
  });

  $(document).on("click", "#generarDiarioEnExcel", function (){
    var diario = parseInt($("#diario").val());
    window.open(
      "/generarDiarioEnExcel?diario=" + diario,
      '_blank'
    );

  });

  $(document).on("click", "#generarDiarioEnPDF", function (){
    var diario = parseInt($("#diario").val());
    window.open(
      "/generarDiario?diario=" + diario,
      '_blank'
    );

  });

  

  $(document).on("click", "#generarEstadoDeCuenta", function (){
    var auxC = $("#autocomplete-custom-append").val();
    if (auxC.indexOf("|")>0)
    {
      var cuenta = auxC.split("|")[0].trim();
      var descr = auxC.split("|")[1].trim();
      var delPeriodo = $("#delPeriodo").val();
      var alPeriodo = $("#alPeriodo").val();
      var ANAL_T0_START = "";
      if($("#ANAL_T0_START").length>0)
      {
        ANAL_T0_START = $("#ANAL_T0_START").val();  
      }
      var ANAL_T0_END = "";
      if($("#ANAL_T0_END").length>0)
      {
        ANAL_T0_END = $("#ANAL_T0_END").val();
      }
      var ANAL_T1_START = "";
      if($("#ANAL_T1_START").length>0)
      {
        ANAL_T1_START = $("#ANAL_T1_START").val();  
      }
      var ANAL_T1_END = "";
      if($("#ANAL_T1_END").length>0)
      {
        ANAL_T1_END = $("#ANAL_T1_END").val();  
      }
      var ANAL_T2_START = "";
      if($("#ANAL_T2_START").length>0)
      {
        ANAL_T2_START = $("#ANAL_T2_START").val();  
      }
      var ANAL_T2_END = "";
      if($("#ANAL_T2_END").length>0)
      {
        ANAL_T2_END = $("#ANAL_T2_END").val();  
      }
      var ANAL_T3_START = "";
      if($("#ANAL_T3_START").length>0)
      {
        ANAL_T3_START = $("#ANAL_T3_START").val();  
      }
      var ANAL_T3_END = "";
      if($("#ANAL_T3_END").length>0)
      {
        ANAL_T3_END = $("#ANAL_T3_END").val();  
      }
      var ANAL_T4_START = "";
      if($("#ANAL_T4_START").length>0)
      {
        ANAL_T4_START = $("#ANAL_T4_START").val();  
      }
      var ANAL_T4_END = "";
      if($("#ANAL_T4_END").length>0)
      {
        ANAL_T4_END = $("#ANAL_T4_END").val();  
      }
      var ANAL_T5_START = "";
      if($("#ANAL_T5_START").length>0)
      {
        ANAL_T5_START = $("#ANAL_T5_START").val();  
      }
      var ANAL_T5_END = "";
      if($("#ANAL_T5_END").length>0)
      {
        ANAL_T5_END = $("#ANAL_T5_END").val();  
      }
      var ANAL_T6_START = "";
      if($("#ANAL_T6_START").length>0)
      {
        ANAL_T6_START = $("#ANAL_T6_START").val();  
      }
      var ANAL_T6_END = "";
      if($("#ANAL_T6_END").length>0)
      {
        ANAL_T6_END = $("#ANAL_T6_END").val();
      }

      var ANAL_T7_START = "";
      if($("#ANAL_T7_START").length>0)
      {
        ANAL_T7_START = $("#ANAL_T7_START").val();  
      }
      var ANAL_T7_END = "";
      if($("#ANAL_T7_START").length>0)
      {
      ANAL_T7_END = $("#ANAL_T7_END").val();
}
      var ANAL_T8_START = $("#ANAL_T8_START").val();
      var ANAL_T8_END = $("#ANAL_T8_END").val();

      var ANAL_T9_START = $("#ANAL_T9_START").val();
      var ANAL_T9_END = $("#ANAL_T9_END").val();



      window.open(
        "/generarEstadoDeCuenta?cuenta=" + cuenta+"&delPeriodo="+delPeriodo+"&alPeriodo="+alPeriodo+"&descr="+descr+"&d0s="+ANAL_T0_START+"&d0e="+ANAL_T0_END+"&d1s="+ANAL_T1_START+"&d1e="+ANAL_T1_END+"&d2s="+ANAL_T2_START+"&d2e="+ANAL_T2_END+"&d3s="+ANAL_T3_START+"&d3e="+ANAL_T3_END+"&d4s="+ANAL_T4_START+"&d4e="+ANAL_T4_END+"&d5s="+ANAL_T5_START+"&d5e="+ANAL_T5_END+"&d6s="+ANAL_T6_START+"&d6e="+ANAL_T6_END+"&d7s="+ANAL_T7_START+"&d7e="+ANAL_T7_END+"&d8s="+ANAL_T8_START+"&d8e="+ANAL_T8_END+"&d9s="+ANAL_T9_START+"&d9e="+ANAL_T9_END+"&busca=0",
        '_blank'
      );
    }
    else
    {
      alert("Tienes que seleccionar una opción del autocompletar.")
    }
  });


   
  function configInicialDeDiariosReversiados()
  {
     $.ajax({
        url: "/configInicialDiariosReversiados",
        dataType : "json",
        type : "post",
       // data: param, 
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp) {
            var cuentasArray = new Array();
            var descrArray = new Array();
            var aver2 = Object.keys(resp.diariosReversiados);
             var cad = '';
             var palabra = "Original";
            for(i=0;i<aver2.length;i++)
            {
              var diario = resp.diariosReversiados[aver2[i]]["Diario"];
              var ref = resp.diariosReversiados[aver2[i]]["AllocRef"];
              var PERIOD = resp.diariosReversiados[aver2[i]]["PERIOD"];
              //<a class="vesAlDiario" href="#" diario="10">10-1</a>
              if(i%2==0)
              {
                palabra = "Original";
              }
              else
              {
                palabra = "Reverseado";
              }
              cad=cad+'<tr><td><a class="vesAlDiario" href="#" diario="'+diario+'">'+diario+'</a></td><td>'+ref+'</td><td>'+PERIOD+'</td></tr>';
            }
            $("#tablaDiariosReversiados").html(cad);
          }
      });
  }
  function configInicialDeCedulasConceptos()
  {
     $.ajax({
        url: "/configInicialLineasCedulas",
        dataType : "json",
        type : "post",
       // data: param, 
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp) {
            var cuentasArray = new Array();
            var descrArray = new Array();
            var aver2 = Object.keys(resp.cuentas);
             var cedulas = Object.keys(resp.cedulas);
            var cad = '<option value="-1">Selecciona una opci&oacute;n</option>';
            for(i=0;i<cedulas.length;i++)
            {
              var todo = cedulas[i].split('|');
              var nombre = todo[0];
              var idCedula = todo[1];
              cad=cad+'<option value="'+idCedula+'">'+nombre+'</option>';
            }
            $("#cedulasSelectConceptos").html(cad);
          }
      });
  }
  function configInicialDeCedulas()
  {
     $.ajax({
        url: "/configInicialLineasCedulas",
        dataType : "json",
        type : "post",
       // data: param, 
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp) {
            var cuentasArray = new Array();
            var descrArray = new Array();
            var aver2 = Object.keys(resp.cuentas);
             var cedulas = Object.keys(resp.cedulas);
            var cad = '<option value="-1">Selecciona una opci&oacute;n</option>';
            for(i=0;i<cedulas.length;i++)
            {
              var todo = cedulas[i].split('|');
              var nombre = todo[0];
              var idCedula = todo[1];
              cad=cad+'<option value="'+idCedula+'">'+nombre+'</option>';
            }
            $("#cedulasSelect").html(cad);

             aver2 = Object.keys(resp.periodos);
              var cad = '';
              for(i=0;i<aver2.length;i++)
              {
                cad=cad+'<option value="'+aver2[i].toString().trim()+'">'+aver2[i].toString().trim()+'</option>';
              } 
              $("#delPeriodo").html(cad);
              $("#alPeriodo").html(cad);
              $("#alPeriodo option:last").attr("selected","selected");
              actualizaConceptosCedulas();
          }
      });
  }
   function configInicialDeCedulasParaVer()
  {
     $.ajax({
        url: "/configInicialLineasCedulas",
        dataType : "json",
        type : "post",
       // data: param, 
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp) {
            var cuentasArray = new Array();
            var descrArray = new Array();
            var aver2 = Object.keys(resp.cuentas);
             var cedulas = Object.keys(resp.cedulas);
            var cad = '<option value="-1">Selecciona una opci&oacute;n</option>';
            for(i=0;i<cedulas.length;i++)
            {
              var todo = cedulas[i].split('|');
              var nombre = todo[0];
              var idCedula = todo[1];
              cad=cad+'<option value="'+idCedula+'">'+nombre+'</option>';
            }
            $("#cedulasSelectVer").html(cad);

             aver2 = Object.keys(resp.periodos);
              var cad = '';
              for(i=0;i<aver2.length;i++)
              {
                cad=cad+'<option value="'+aver2[i].toString().trim()+'">'+aver2[i].toString().trim()+'</option>';
              } 
              $("#delPeriodo").html(cad);
              $("#alPeriodo").html(cad);
              $("#alPeriodo option:last").attr("selected","selected");
          }
      });
  }
  
   function configInicialDeLineasDeCapturaDeDiario()
  {
     $.ajax({
        url: "/configInicialDeLineasDeTiposDeDiario",
        dataType : "json",
        type : "post",
       // data: param, 
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp) {
            timestampActual = resp.time;
            var cuentasArray = new Array();
            var descrArray = new Array();
            var aver2 = Object.keys(resp.cuentas);
            var ANAL_T0 = Object.keys(resp.ANAL_T0);
            var ANAL_T1 = Object.keys(resp.ANAL_T1);
            var ANAL_T2 = Object.keys(resp.ANAL_T2);
            var ANAL_T3 = Object.keys(resp.ANAL_T3);
            var ANAL_T4 = Object.keys(resp.ANAL_T4);
            var ANAL_T5 = Object.keys(resp.ANAL_T5);
            var ANAL_T6 = Object.keys(resp.ANAL_T6);
            var ANAL_T7 = Object.keys(resp.ANAL_T7);
            var ANAL_T8 = Object.keys(resp.ANAL_T8);
            var ANAL_T9 = Object.keys(resp.ANAL_T9);
            var tiposDeDiario = Object.keys(resp.tiposDeDiario);
            var cad = '<option value="-1">Selecciona una opci&oacute;n</option>';
            for(i=0;i<tiposDeDiario.length;i++)
            {
              var IdTipoDeDiario = resp.tiposDeDiario[tiposDeDiario[i]]["IdTipoDeDiario"];
              var nombre = resp.tiposDeDiario[tiposDeDiario[i]]["Nombre"];
              cad=cad+'<option value="'+IdTipoDeDiario+'">'+nombre+'</option>';
            }
            $("#tipoDeDiarioSelectDeCaptura").html(cad);
            if(resp.D1=="")
            {
              resp.D1 = "sin nombre";
            }
            $("#ANAL_T0_LABEL").html(resp.D1);
            $("#ANAL_T1_LABEL").html(resp.D2);
            $("#ANAL_T2_LABEL").html(resp.D3);
            $("#ANAL_T3_LABEL").html(resp.D4);
            $("#ANAL_T4_LABEL").html(resp.D5);
            $("#ANAL_T5_LABEL").html(resp.D6);
            $("#ANAL_T6_LABEL").html(resp.D7);
            $("#ANAL_T7_LABEL").html(resp.D8);
            $("#ANAL_T8_LABEL").html(resp.D9);
            $("#ANAL_T9_LABEL").html(resp.D10);


            $("#tdD0").html(resp.D1);
            $("#tdD1").html(resp.D2);
            $("#tdD2").html(resp.D3);
            $("#tdD3").html(resp.D4);
            $("#tdD4").html(resp.D5);
            $("#tdD5").html(resp.D6);
            $("#tdD6").html(resp.D7);
            $("#tdD7").html(resp.D8);
            $("#tdD8").html(resp.D9);
            $("#tdD9").html(resp.D10);
            
            for(i=0;i<aver2.length;i++)
            {
              var aver3 = Object.keys(resp.cuentas[aver2[i]]);
              //var param = {value : aver2[i].toString(), data : resp.cuentas[aver2[i]]["DESCR"] }
              cuentasArray.push( aver2[i].toString().trim() + "|"+ resp.cuentas[aver2[i]]["DESCR"]);
            }
            $('#autocomplete-custom-append').autocomplete({
              lookup: cuentasArray,
              appendTo: '#autocomplete-container'
            }); 
               
            var ANAL_T0Array = new Array();
            for(i=0;i<ANAL_T0.length;i++)
            {
              var aver3 = Object.keys(resp.ANAL_T0[ANAL_T0[i]]);
              ANAL_T0Array.push( ANAL_T0[i].toString().trim() + "|"+ resp.ANAL_T0[ANAL_T0[i]]["NAME"]);
            }    
            $('#ANAL_T0').autocomplete({
              lookup: ANAL_T0Array,
              appendTo: '#ANAL_T0-container'
            }); 
            
            var ANAL_T1Array = new Array();
            for(i=0;i<ANAL_T1.length;i++)
            {
              ANAL_T1Array.push( ANAL_T1[i].toString().trim() + "|"+ resp.ANAL_T1[ANAL_T1[i]]["NAME"]);
            }    
            $('#ANAL_T1').autocomplete({
              lookup: ANAL_T1Array,
              appendTo: '#ANAL_T1-container'
            });
            
            var ANAL_T2Array = new Array();
            for(i=0;i<ANAL_T2.length;i++)
            {
              ANAL_T2Array.push( ANAL_T2[i].toString().trim() + "|"+ resp.ANAL_T2[ANAL_T2[i]]["NAME"]);
            }    
            $('#ANAL_T2').autocomplete({
              lookup: ANAL_T2Array,
              appendTo: '#ANAL_T2-container'
            });

            var ANAL_T3Array = new Array();
            for(i=0;i<ANAL_T3.length;i++)
            {
              ANAL_T3Array.push( ANAL_T3[i].toString().trim() + "|"+ resp.ANAL_T3[ANAL_T3[i]]["NAME"]);
            }    
            $('#ANAL_T3').autocomplete({
              lookup: ANAL_T3Array,
              appendTo: '#ANAL_T3-container'
            });

            var ANAL_T4Array = new Array();
            for(i=0;i<ANAL_T4.length;i++)
            {
              ANAL_T4Array.push( ANAL_T4[i].toString().trim() + "|"+ resp.ANAL_T4[ANAL_T4[i]]["NAME"]);
            }    
            $('#ANAL_T4').autocomplete({
              lookup: ANAL_T4Array,
              appendTo: '#ANAL_T4-container'
            });
            
            var ANAL_T5Array = new Array();
            for(i=0;i<ANAL_T5.length;i++)
            {
              ANAL_T5Array.push( ANAL_T5[i].toString().trim() + "|"+ resp.ANAL_T5[ANAL_T5[i]]["NAME"]);
            }    
            $('#ANAL_T5').autocomplete({
              lookup: ANAL_T5Array,
              appendTo: '#ANAL_T5-container'
            }); 


            var ANAL_T6Array = new Array();
            for(i=0;i<ANAL_T6.length;i++)
            {
              ANAL_T6Array.push( ANAL_T6[i].toString().trim() + "|"+ resp.ANAL_T6[ANAL_T6[i]]["NAME"]);
            }    
            $('#ANAL_T6').autocomplete({
              lookup: ANAL_T6Array,
              appendTo: '#ANAL_T6-container'
            });

             var ANAL_T7Array = new Array();
            for(i=0;i<ANAL_T7.length;i++)
            {
              ANAL_T7Array.push( ANAL_T7[i].toString().trim() + "|"+ resp.ANAL_T7[ANAL_T7[i]]["NAME"]);
            }    
            $('#ANAL_T7').autocomplete({
              lookup: ANAL_T7Array,
              appendTo: '#ANAL_T7-container'
            }); 

             var ANAL_T8Array = new Array();
            for(i=0;i<ANAL_T8.length;i++)
            {
              ANAL_T8Array.push( ANAL_T8[i].toString().trim() + "|"+ resp.ANAL_T8[ANAL_T8[i]]["NAME"]);
            }    
            $('#ANAL_T8').autocomplete({
              lookup: ANAL_T8Array,
              appendTo: '#ANAL_T8-container'
            }); 

             var ANAL_T9Array = new Array();
            for(i=0;i<ANAL_T9.length;i++)
            {
              ANAL_T9Array.push( ANAL_T9[i].toString().trim() + "|"+ resp.ANAL_T9[ANAL_T9[i]]["NAME"]);
            }    
            $('#ANAL_T9').autocomplete({
              lookup: ANAL_T9Array,
              appendTo: '#ANAL_T9-container'
            });  
          

          }
      });
  }

   function configInicialDeLineasDeTiposDeDiario()
  {
     $.ajax({
        url: "/configInicialDeLineasDeTiposDeDiario",
        dataType : "json",
        type : "post",
       // data: param, 
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp) {
            var cuentasArray = new Array();
            var descrArray = new Array();
            var aver2 = Object.keys(resp.cuentas);
            var ANAL_T0 = Object.keys(resp.ANAL_T0);
            var ANAL_T1 = Object.keys(resp.ANAL_T1);
            var ANAL_T2 = Object.keys(resp.ANAL_T2);
            var ANAL_T3 = Object.keys(resp.ANAL_T3);
            var ANAL_T4 = Object.keys(resp.ANAL_T4);
            var ANAL_T5 = Object.keys(resp.ANAL_T5);
            var ANAL_T6 = Object.keys(resp.ANAL_T6);
            var ANAL_T7 = Object.keys(resp.ANAL_T7);
            var ANAL_T8 = Object.keys(resp.ANAL_T8);
            var ANAL_T9 = Object.keys(resp.ANAL_T9);
            var tiposDeDiario = Object.keys(resp.tiposDeDiario);
            var cad = '<option value="-1">Selecciona una opci&oacute;n</option>';
            for(i=0;i<tiposDeDiario.length;i++)
            {
              var IdTipoDeDiario = resp.tiposDeDiario[tiposDeDiario[i]]["IdTipoDeDiario"];
              var nombre = resp.tiposDeDiario[tiposDeDiario[i]]["Nombre"];
              cad=cad+'<option value="'+IdTipoDeDiario+'">'+nombre+'</option>';
            }
            $("#tipoDeDiarioSelect").html(cad);
            if(resp.D1=="")
            {
              resp.D1 = "sin nombre";
            }
            $("#ANAL_T0_LABEL").html(resp.D1);
            $("#ANAL_T1_LABEL").html(resp.D2);
            $("#ANAL_T2_LABEL").html(resp.D3);
            $("#ANAL_T3_LABEL").html(resp.D4);
            $("#ANAL_T4_LABEL").html(resp.D5);
            $("#ANAL_T5_LABEL").html(resp.D6);
            $("#ANAL_T6_LABEL").html(resp.D7);
            $("#ANAL_T7_LABEL").html(resp.D8);
            $("#ANAL_T8_LABEL").html(resp.D9);
            $("#ANAL_T9_LABEL").html(resp.D10);


            $("#tdD0").html(resp.D1);
            $("#tdD1").html(resp.D2);
            $("#tdD2").html(resp.D3);
            $("#tdD3").html(resp.D4);
            $("#tdD4").html(resp.D5);
            $("#tdD5").html(resp.D6);
            $("#tdD6").html(resp.D7);
            $("#tdD7").html(resp.D8);
            $("#tdD8").html(resp.D9);
            $("#tdD9").html(resp.D10);
            
            for(i=0;i<aver2.length;i++)
            {
              var aver3 = Object.keys(resp.cuentas[aver2[i]]);
              //var param = {value : aver2[i].toString(), data : resp.cuentas[aver2[i]]["DESCR"] }
              cuentasArray.push( aver2[i].toString().trim() + "|"+ resp.cuentas[aver2[i]]["DESCR"]);
            }
            $('#autocomplete-custom-append').autocomplete({
              lookup: cuentasArray,
              appendTo: '#autocomplete-container'
            }); 
               
            var ANAL_T0Array = new Array();
            for(i=0;i<ANAL_T0.length;i++)
            {
              var aver3 = Object.keys(resp.ANAL_T0[ANAL_T0[i]]);
              ANAL_T0Array.push( ANAL_T0[i].toString().trim() + "|"+ resp.ANAL_T0[ANAL_T0[i]]["NAME"]);
            }    
            $('#ANAL_T0').autocomplete({
              lookup: ANAL_T0Array,
              appendTo: '#ANAL_T0-container'
            }); 
            
            var ANAL_T1Array = new Array();
            for(i=0;i<ANAL_T1.length;i++)
            {
              ANAL_T1Array.push( ANAL_T1[i].toString().trim() + "|"+ resp.ANAL_T1[ANAL_T1[i]]["NAME"]);
            }    
            $('#ANAL_T1').autocomplete({
              lookup: ANAL_T1Array,
              appendTo: '#ANAL_T1-container'
            });
            
            var ANAL_T2Array = new Array();
            for(i=0;i<ANAL_T2.length;i++)
            {
              ANAL_T2Array.push( ANAL_T2[i].toString().trim() + "|"+ resp.ANAL_T2[ANAL_T2[i]]["NAME"]);
            }    
            $('#ANAL_T2').autocomplete({
              lookup: ANAL_T2Array,
              appendTo: '#ANAL_T2-container'
            });

            var ANAL_T3Array = new Array();
            for(i=0;i<ANAL_T3.length;i++)
            {
              ANAL_T3Array.push( ANAL_T3[i].toString().trim() + "|"+ resp.ANAL_T3[ANAL_T3[i]]["NAME"]);
            }    
            $('#ANAL_T3').autocomplete({
              lookup: ANAL_T3Array,
              appendTo: '#ANAL_T3-container'
            });

            var ANAL_T4Array = new Array();
            for(i=0;i<ANAL_T4.length;i++)
            {
              ANAL_T4Array.push( ANAL_T4[i].toString().trim() + "|"+ resp.ANAL_T4[ANAL_T4[i]]["NAME"]);
            }    
            $('#ANAL_T4').autocomplete({
              lookup: ANAL_T4Array,
              appendTo: '#ANAL_T4-container'
            });
            
            var ANAL_T5Array = new Array();
            for(i=0;i<ANAL_T5.length;i++)
            {
              ANAL_T5Array.push( ANAL_T5[i].toString().trim() + "|"+ resp.ANAL_T5[ANAL_T5[i]]["NAME"]);
            }    
            $('#ANAL_T5').autocomplete({
              lookup: ANAL_T5Array,
              appendTo: '#ANAL_T5-container'
            }); 


            var ANAL_T6Array = new Array();
            for(i=0;i<ANAL_T6.length;i++)
            {
              ANAL_T6Array.push( ANAL_T6[i].toString().trim() + "|"+ resp.ANAL_T6[ANAL_T6[i]]["NAME"]);
            }    
            $('#ANAL_T6').autocomplete({
              lookup: ANAL_T6Array,
              appendTo: '#ANAL_T6-container'
            });

             var ANAL_T7Array = new Array();
            for(i=0;i<ANAL_T7.length;i++)
            {
              ANAL_T7Array.push( ANAL_T7[i].toString().trim() + "|"+ resp.ANAL_T7[ANAL_T7[i]]["NAME"]);
            }    
            $('#ANAL_T7').autocomplete({
              lookup: ANAL_T7Array,
              appendTo: '#ANAL_T7-container'
            }); 

             var ANAL_T8Array = new Array();
            for(i=0;i<ANAL_T8.length;i++)
            {
              ANAL_T8Array.push( ANAL_T8[i].toString().trim() + "|"+ resp.ANAL_T8[ANAL_T8[i]]["NAME"]);
            }    
            $('#ANAL_T8').autocomplete({
              lookup: ANAL_T8Array,
              appendTo: '#ANAL_T8-container'
            }); 

             var ANAL_T9Array = new Array();
            for(i=0;i<ANAL_T9.length;i++)
            {
              ANAL_T9Array.push( ANAL_T9[i].toString().trim() + "|"+ resp.ANAL_T9[ANAL_T9[i]]["NAME"]);
            }    
            $('#ANAL_T9').autocomplete({
              lookup: ANAL_T9Array,
              appendTo: '#ANAL_T9-container'
            });  
          

          }
      });
  }
  function configInicialDeLineasDeCedulas()
  {
     $.ajax({
        url: "/configInicialLineasCedulas",
        dataType : "json",
        type : "post",
       // data: param, 
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp) {
            var cuentasArray = new Array();
            var descrArray = new Array();
            var aver2 = Object.keys(resp.cuentas);
            var ANAL_T0 = Object.keys(resp.ANAL_T0);
            var ANAL_T1 = Object.keys(resp.ANAL_T1);
            var ANAL_T2 = Object.keys(resp.ANAL_T2);
            var ANAL_T3 = Object.keys(resp.ANAL_T3);
            var ANAL_T4 = Object.keys(resp.ANAL_T4);
            var ANAL_T5 = Object.keys(resp.ANAL_T5);
            var ANAL_T6 = Object.keys(resp.ANAL_T6);
            var ANAL_T7 = Object.keys(resp.ANAL_T7);
            var ANAL_T8 = Object.keys(resp.ANAL_T8);
            var ANAL_T9 = Object.keys(resp.ANAL_T9);
            var cedulas = Object.keys(resp.cedulas);
            var cad = '<option value="-1">Selecciona una opci&oacute;n</option>';
            for(i=0;i<cedulas.length;i++)
            {
              var todo = cedulas[i].split('|');
              var nombre = todo[0];
              var idCedula = todo[1];
              cad=cad+'<option value="'+idCedula+'">'+nombre+'</option>';
            }
            $("#cedulasSelect").html(cad);
            if(resp.D1=="")
            {
              resp.D1 = "sin nombre";
            }
            $("#ANAL_T0_LABEL").html(resp.D1);
            $("#ANAL_T1_LABEL").html(resp.D2);
            $("#ANAL_T2_LABEL").html(resp.D3);
            $("#ANAL_T3_LABEL").html(resp.D4);
            $("#ANAL_T4_LABEL").html(resp.D5);
            $("#ANAL_T5_LABEL").html(resp.D6);
            $("#ANAL_T6_LABEL").html(resp.D7);
            $("#ANAL_T7_LABEL").html(resp.D8);
            $("#ANAL_T8_LABEL").html(resp.D9);
            $("#ANAL_T9_LABEL").html(resp.D10);


            $("#tdD0").html(resp.D1);
            $("#tdD1").html(resp.D2);
            $("#tdD2").html(resp.D3);
            $("#tdD3").html(resp.D4);
            $("#tdD4").html(resp.D5);
            $("#tdD5").html(resp.D6);
            $("#tdD6").html(resp.D7);
            $("#tdD7").html(resp.D8);
            $("#tdD8").html(resp.D9);
            $("#tdD9").html(resp.D10);
            
            for(i=0;i<aver2.length;i++)
            {
              var aver3 = Object.keys(resp.cuentas[aver2[i]]);
              //var param = {value : aver2[i].toString(), data : resp.cuentas[aver2[i]]["DESCR"] }
              cuentasArray.push( aver2[i].toString().trim() + "|"+ resp.cuentas[aver2[i]]["DESCR"]);
            }
            $('#autocomplete-custom-append').autocomplete({
              lookup: cuentasArray,
              appendTo: '#autocomplete-container'
            }); 
               
            var ANAL_T0Array = new Array();
            for(i=0;i<ANAL_T0.length;i++)
            {
              var aver3 = Object.keys(resp.ANAL_T0[ANAL_T0[i]]);
              ANAL_T0Array.push( ANAL_T0[i].toString().trim() + "|"+ resp.ANAL_T0[ANAL_T0[i]]["NAME"]);
            }    
            $('#ANAL_T0').autocomplete({
              lookup: ANAL_T0Array,
              appendTo: '#ANAL_T0-container'
            }); 
            
            var ANAL_T1Array = new Array();
            for(i=0;i<ANAL_T1.length;i++)
            {
              ANAL_T1Array.push( ANAL_T1[i].toString().trim() + "|"+ resp.ANAL_T1[ANAL_T1[i]]["NAME"]);
            }    
            $('#ANAL_T1').autocomplete({
              lookup: ANAL_T1Array,
              appendTo: '#ANAL_T1-container'
            });
            
            var ANAL_T2Array = new Array();
            for(i=0;i<ANAL_T2.length;i++)
            {
              ANAL_T2Array.push( ANAL_T2[i].toString().trim() + "|"+ resp.ANAL_T2[ANAL_T2[i]]["NAME"]);
            }    
            $('#ANAL_T2').autocomplete({
              lookup: ANAL_T2Array,
              appendTo: '#ANAL_T2-container'
            });

            var ANAL_T3Array = new Array();
            for(i=0;i<ANAL_T3.length;i++)
            {
              ANAL_T3Array.push( ANAL_T3[i].toString().trim() + "|"+ resp.ANAL_T3[ANAL_T3[i]]["NAME"]);
            }    
            $('#ANAL_T3').autocomplete({
              lookup: ANAL_T3Array,
              appendTo: '#ANAL_T3-container'
            });

            var ANAL_T4Array = new Array();
            for(i=0;i<ANAL_T4.length;i++)
            {
              ANAL_T4Array.push( ANAL_T4[i].toString().trim() + "|"+ resp.ANAL_T4[ANAL_T4[i]]["NAME"]);
            }    
            $('#ANAL_T4').autocomplete({
              lookup: ANAL_T4Array,
              appendTo: '#ANAL_T4-container'
            });
            
            var ANAL_T5Array = new Array();
            for(i=0;i<ANAL_T5.length;i++)
            {
              ANAL_T5Array.push( ANAL_T5[i].toString().trim() + "|"+ resp.ANAL_T5[ANAL_T5[i]]["NAME"]);
            }    
            $('#ANAL_T5').autocomplete({
              lookup: ANAL_T5Array,
              appendTo: '#ANAL_T5-container'
            }); 


            var ANAL_T6Array = new Array();
            for(i=0;i<ANAL_T6.length;i++)
            {
              ANAL_T6Array.push( ANAL_T6[i].toString().trim() + "|"+ resp.ANAL_T6[ANAL_T6[i]]["NAME"]);
            }    
            $('#ANAL_T6').autocomplete({
              lookup: ANAL_T6Array,
              appendTo: '#ANAL_T6-container'
            });

             var ANAL_T7Array = new Array();
            for(i=0;i<ANAL_T7.length;i++)
            {
              ANAL_T7Array.push( ANAL_T7[i].toString().trim() + "|"+ resp.ANAL_T7[ANAL_T7[i]]["NAME"]);
            }    
            $('#ANAL_T7').autocomplete({
              lookup: ANAL_T7Array,
              appendTo: '#ANAL_T7-container'
            }); 

             var ANAL_T8Array = new Array();
            for(i=0;i<ANAL_T8.length;i++)
            {
              ANAL_T8Array.push( ANAL_T8[i].toString().trim() + "|"+ resp.ANAL_T8[ANAL_T8[i]]["NAME"]);
            }    
            $('#ANAL_T8').autocomplete({
              lookup: ANAL_T8Array,
              appendTo: '#ANAL_T8-container'
            }); 

             var ANAL_T9Array = new Array();
            for(i=0;i<ANAL_T9.length;i++)
            {
              ANAL_T9Array.push( ANAL_T9[i].toString().trim() + "|"+ resp.ANAL_T9[ANAL_T9[i]]["NAME"]);
            }    
            $('#ANAL_T9').autocomplete({
              lookup: ANAL_T9Array,
              appendTo: '#ANAL_T9-container'
            });  
          

          }
      });
  }
  
  

  function actualizaConceptosCedulas()
  {
    var idCedula = $("#cedulasSelectConceptos").val();
    var param = {idCedula : idCedula};
     $.ajax({
        url: "/listaConceptosCedulas",
        dataType : "json",
        type : "post",
        data: param, 
        async : true,
        beforeSend : function (){
         },
        complete : function (){
        }, 
        success : function (resp){
            var aver2 = Object.keys(resp.cedulasConceptos);
            var cad = '';
            for(i=0;i<aver2.length;i++)
            {
              cad=cad+'<tr><td>'+resp.cedulasConceptos[aver2[i]]["IdConcepto"]+'</td><td>'+resp.cedulasConceptos[aver2[i]]["Nombre"]+'</td><td><button type="button" IdConcepto="'+resp.cedulasConceptos[aver2[i]]["IdConcepto"]+'" nombre="'+resp.cedulasConceptos[aver2[i]]["Nombre"]+'" class="btn btn-round btn-warning editarConceptoCedula">Editar</button></td></tr>';
            } 
            $("#tablaCedulaConceptos").html(cad);
        }
      });
  }


  
  function actualizaLaTablaDondeEstanLasLineasDelTipoDeDiario()
  {
    var idTipoDeDiario = $("#tipoDeDiarioSelect").val();
    var param = {idTipoDeDiario : idTipoDeDiario};
     $.ajax({
        url: "/listaLineasTiposDeDiario",
        dataType : "json",
        type : "post",
        data: param, 
        async : true,
        beforeSend : function (){
         },
        complete : function (){
        }, 
        success : function (resp){
           
            var aver2 = Object.keys(resp.tiposDeDiarioLineas);
            var cad = '';
            for(i=0;i<aver2.length;i++)
            {
              var tipoDC = parseInt(resp.tiposDeDiarioLineas[aver2[i]]["D_C"]);
              var palabraDC ="";
              if(tipoDC=="D")
              {
                palabraDC = "D&eacute;bitos";
              }
              if(tipoDC=="C")
              {
                palabraDC = "Cr&eacute;ditos";
              }
              cad=cad+'<tr><td>'+resp.tiposDeDiarioLineas[aver2[i]]["Linea"]+'</td><td>'+resp.tiposDeDiarioLineas[aver2[i]]["Cuenta"]+'</td><td>'+palabraDC+'</td><td>'+resp.tiposDeDiarioLineas[aver2[i]]["DESCRIPTN"]+'</td><td>'+resp.tiposDeDiarioLineas[aver2[i]]["ANAL_T0"]+'</td><td>'+resp.tiposDeDiarioLineas[aver2[i]]["ANAL_T1"]+'</td><td>'+resp.tiposDeDiarioLineas[aver2[i]]["ANAL_T2"]+'</td><td>'+resp.tiposDeDiarioLineas[aver2[i]]["ANAL_T3"]+'</td><td>'+resp.tiposDeDiarioLineas[aver2[i]]["ANAL_T4"]+'</td><td>'+resp.tiposDeDiarioLineas[aver2[i]]["ANAL_T5"]+'</td><td>'+resp.tiposDeDiarioLineas[aver2[i]]["ANAL_T6"]+'</td><td>'+resp.tiposDeDiarioLineas[aver2[i]]["ANAL_T7"]+'</td><td>'+resp.tiposDeDiarioLineas[aver2[i]]["ANAL_T8"]+'</td><td>'+resp.tiposDeDiarioLineas[aver2[i]]["ANAL_T9"]+'</td><td>'+resp.tiposDeDiarioLineas[aver2[i]]["DeboFacturar"]+'</td><td>'+resp.tiposDeDiarioLineas[aver2[i]]["Cliente"]+'</td><td>'+resp.tiposDeDiarioLineas[aver2[i]]["Servicio"]+'</td><td>'+resp.tiposDeDiarioLineas[aver2[i]]["Concepto"]+'</td><td><button type="button" idLinea="'+resp.tiposDeDiarioLineas[aver2[i]]["IdLinea"]+'" class="btn btn-round btn-warning editarLineaTipoDeDiario">Editar</button></td></tr>';
            } 
            $("#tablaLineaTiposDeDiario").html(cad);
            
        }
      });
  }
  function actualizaLaTablaDondeEstanLasLineasDeLasCedulaSegunElConcepto()
  {
    var IdConcepto = $("#conceptoSelect").val();
    var param = {IdConcepto : IdConcepto};
     $.ajax({
        url: "/listaLineasCedulas",
        dataType : "json",
        type : "post",
        data: param, 
        async : true,
        beforeSend : function (){
         },
        complete : function (){
        }, 
        success : function (resp){
           
            var aver2 = Object.keys(resp.cedulasLineas);
            var cad = '';
            for(i=0;i<aver2.length;i++)
            {
              var tipoDC = parseInt(resp.cedulasLineas[aver2[i]]["TipoDC"]);
              var palabraDC ="";
              if(tipoDC==1)
              {
                palabraDC = "D&eacute;bitos";
              }
              if(tipoDC==2)
              {
                palabraDC = "Cr&eacute;ditos";
              }
              if(tipoDC==4)
              {
                palabraDC = "Ambos D y C";
              }
              cad=cad+'<tr><td>'+resp.cedulasLineas[aver2[i]]["IdLinea"]+'</td><td>'+resp.cedulasLineas[aver2[i]]["Cuenta"]+'</td><td>'+palabraDC+'</td><td>'+resp.cedulasLineas[aver2[i]]["ANAL_T0"]+'</td><td>'+resp.cedulasLineas[aver2[i]]["ANAL_T1"]+'</td><td>'+resp.cedulasLineas[aver2[i]]["ANAL_T2"]+'</td><td>'+resp.cedulasLineas[aver2[i]]["ANAL_T3"]+'</td><td>'+resp.cedulasLineas[aver2[i]]["ANAL_T4"]+'</td><td>'+resp.cedulasLineas[aver2[i]]["ANAL_T5"]+'</td><td>'+resp.cedulasLineas[aver2[i]]["ANAL_T6"]+'</td><td>'+resp.cedulasLineas[aver2[i]]["ANAL_T7"]+'</td><td>'+resp.cedulasLineas[aver2[i]]["ANAL_T8"]+'</td><td>'+resp.cedulasLineas[aver2[i]]["ANAL_T9"]+'</td><td><button type="button" idLinea="'+resp.cedulasLineas[aver2[i]]["IdLinea"]+'" nombre="'+resp.cedulasLineas[aver2[i]]["Nombre"]+'" class="btn btn-round btn-warning editarLineaCedula">Editar</button></td></tr>';
            } 
            $("#tablaLineaCedulas").html(cad);
            
        }
      });
  }
  function actualizaLineaCedulas()
  {
    var idCedula = $("#cedulasSelect").val();
    var param = {idCedula : idCedula};
//listaLineasCedulas
    




     $.ajax({
        url: "/listaConceptosCedulas",
        dataType : "json",
        type : "post",
        data: param, 
        async : true,
        beforeSend : function (){
         },
        complete : function (){
        }, 
        success : function (resp){
            var cedulas = Object.keys(resp.cedulasConceptos);
            var cad = '<option value="-1">Selecciona una opci&oacute;n</option>';
            for(i=0;i<cedulas.length;i++)
            {

              cad=cad+'<option value="'+resp.cedulasConceptos[cedulas[i]]["IdConcepto"]+'">'+resp.cedulasConceptos[cedulas[i]]["Nombre"]+'</option>';
            }
            $("#conceptoSelect").html(cad);

          /*
            var aver2 = Object.keys(resp.cedulasLineas);
            var cad = '';
            for(i=0;i<aver2.length;i++)
            {
              var tipoDC = parseInt(resp.cedulasLineas[aver2[i]]["TipoDC"]);
              var palabraDC ="";
              if(tipoDC==1)
              {
                palabraDC = "D&eacute;bitos";
              }
              if(tipoDC==2)
              {
                palabraDC = "Cr&eacute;ditos";
              }
              if(tipoDC==4)
              {
                palabraDC = "Ambos D y C";
              }
              cad=cad+'<tr><td>'+resp.cedulasLineas[aver2[i]]["IdLinea"]+'</td><td>'+resp.cedulasLineas[aver2[i]]["Nombre"]+'</td><td>'+resp.cedulasLineas[aver2[i]]["Cuenta"]+'</td><td>'+palabraDC+'</td><td>'+resp.cedulasLineas[aver2[i]]["ANAL_T0"]+'</td><td>'+resp.cedulasLineas[aver2[i]]["ANAL_T1"]+'</td><td>'+resp.cedulasLineas[aver2[i]]["ANAL_T2"]+'</td><td>'+resp.cedulasLineas[aver2[i]]["ANAL_T3"]+'</td><td>'+resp.cedulasLineas[aver2[i]]["ANAL_T4"]+'</td><td>'+resp.cedulasLineas[aver2[i]]["ANAL_T5"]+'</td><td>'+resp.cedulasLineas[aver2[i]]["ANAL_T6"]+'</td><td>'+resp.cedulasLineas[aver2[i]]["ANAL_T7"]+'</td><td>'+resp.cedulasLineas[aver2[i]]["ANAL_T8"]+'</td><td>'+resp.cedulasLineas[aver2[i]]["ANAL_T9"]+'</td><td><button type="button" idLinea="'+resp.cedulasLineas[aver2[i]]["IdLinea"]+'" nombre="'+resp.cedulasLineas[aver2[i]]["Nombre"]+'" class="btn btn-round btn-warning editarLineaCedula">Editar</button></td></tr>';
            } 
            $("#tablaLineaCedulas").html(cad);
            */
        }
      });
  }

  

  function actualizaTiposDeDiario()
  {
     $.ajax({
        url: "/listaTiposDeDiario",
        dataType : "json",
        type : "post",
        async : true,
        beforeSend : function (){
        },
        complete : function (){
        }, 
        success : function (resp){
            var aver2 = Object.keys(resp.tiposDeDiario);
            var cad = '';
            for(i=0;i<aver2.length;i++)
            {
              var nombre = resp.tiposDeDiario [aver2[i]]["Nombre"];
              var Codigo = resp.tiposDeDiario [aver2[i]]["Codigo"];
              var IdTipoDeDiario = resp.tiposDeDiario [aver2[i]]["IdTipoDeDiario"];
              cad=cad+'<tr><td>'+IdTipoDeDiario+'</td><td>'+nombre+'</td><td>'+Codigo+'</td><td><button type="button" IdTipoDeDiario="'+IdTipoDeDiario+'" nombre="'+nombre+'" Codigo="'+Codigo+'" class="btn btn-round btn-warning editarTipoDeDiario">Editar</button></td></tr>';
            } 
            $("#tablaTiposDeDiario").html(cad);
        }
      });
  }

  function actualizaCedulas()
  {
     $.ajax({
        url: "/listaCedulas",
        dataType : "json",
        type : "post",
       // data: param, 
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp){
            var aver2 = Object.keys(resp.cedulas);
            var cad = '';
            for(i=0;i<aver2.length;i++)
            {
              var todo = aver2[i].split('|');
              var nombre = todo[0];
              var idCedula = todo[1];
              cad=cad+'<tr><td>'+idCedula+'</td><td>'+nombre+'</td><td><button type="button" idCedula="'+idCedula+'" nombre="'+nombre+'" class="btn btn-round btn-warning editarCedula">Editar</button></td></tr>';
            } 
            $("#tablaCedulas").html(cad);
        }
      });
  }

  function actualizaTipoDeDimensiones()
  {
     $.ajax({
        url: "/tipoDeDimensiones",
        dataType : "json",
        type : "post",
       // data: param, 
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp){
            var aver2 = Object.keys(resp.dimensiones);
            var cad = '';
            for(i=0;i<aver2.length;i++)
            {
              cad=cad+'<tr><td>'+resp.dimensiones[aver2[i]]["ANL_CAT_ID"]+'</td><td>'+resp.dimensiones[aver2[i]]["S_HEAD"]+'</td><td>'+resp.dimensiones[aver2[i]]["DESCR"]+'</td><td><button type="button" ANL_CAT_ID="'+resp.dimensiones[aver2[i]]["ANL_CAT_ID"]+'" S_HEAD="'+resp.dimensiones[aver2[i]]["S_HEAD"]+'" DESCR="'+resp.dimensiones[aver2[i]]["DESCR"]+'" class="btn btn-round btn-warning editarTipoDeDimension">Editar</button></td></tr>';
            } 
            $("#tablaDimensiones").html(cad);
        }
      });
  }



  function actualizaClasificacionDimensiones()
  {
     $.ajax({
        url: "/clasificacionDeDimensiones",
        dataType : "json",
        type : "post",
       // data: param, 
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp){
            var aver2 = Object.keys(resp.dimensiones);
            var cad = '<option selected="selected" value="0">Seleccione una opci&oacute;n</option>';
            for(i=0;i<aver2.length;i++)
            {
              cad=cad+'<option S_HEAD="'+resp.dimensiones[aver2[i]]["S_HEAD"]+'" value="'+resp.dimensiones[aver2[i]]["ANL_CAT_ID"]+'">'+resp.dimensiones[aver2[i]]["S_HEAD"]+' - '+resp.dimensiones[aver2[i]]["DESCR"]+'</option>';
            } 
            $("#entry1").html(cad);
            $("#entry2").html(cad);
            $("#entry3").html(cad);
            $("#entry4").html(cad);
            $("#entry5").html(cad);
            $("#entry6").html(cad);
            $("#entry7").html(cad);
            $("#entry8").html(cad);
            $("#entry9").html(cad);
            $("#entry10").html(cad);
        }
      });
  }


  function actualizaDimensiones()
  {
     $.ajax({
        url: "/dimensiones",
        dataType : "json",
        type : "post",
       // data: param, 
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp){
            var aver2 = Object.keys(resp.dimensiones);
            var cad = '';
            for(i=0;i<aver2.length;i++)
            {
              cad=cad+'<tr><td>'+resp.dimensiones[aver2[i]]["S_HEAD"]+'</td><td>'+resp.dimensiones[aver2[i]]["ANL_CODE"]+'</td><td>'+resp.dimensiones[aver2[i]]["NAME"]+'</td><td>'+resp.dimensiones[aver2[i]]["STATUS"]+'</td><td>'+resp.dimensiones[aver2[i]]["PROHIBIT_POSTING"]+'</td><td><button type="button" ANL_CAT_ID="'+resp.dimensiones[aver2[i]]["ANL_CAT_ID"]+'" S_HEAD="'+resp.dimensiones[aver2[i]]["S_HEAD"]+'" ANL_CODE="'+resp.dimensiones[aver2[i]]["ANL_CODE"]+'"  NAME="'+resp.dimensiones[aver2[i]]["NAME"]+'" STATUS="'+resp.dimensiones[aver2[i]]["STATUS"]+'" PROHIBIT_POSTING="'+resp.dimensiones[aver2[i]]["PROHIBIT_POSTING"]+'" class="btn btn-round btn-warning editarDimension">Editar</button></td></tr>';
            } 
            $("#tablaDimensionesEditar").html(cad);
        }
      });
  }

  

  $(document).on("click", "#guardarNuevaDimension", function (){
    var ANL_CAT_ID = $(this).attr("ANL_CAT_ID");
    var ANL_CODE = $(this).attr("ANL_CODE");
    var NAME = $("#nombre").val();
    var STATUS = $("#status").val();
    var PROHIBIT_POSTING = $("#prohibido").val();
    
    if(ANL_CAT_ID=="-1" && ANL_CODE =="-1")
    {
      ANL_CAT_ID = $("#tipo").val();
      ANL_CODE = $("#codigo").val();

      var param = {
      tipo : ANL_CAT_ID,
      codigo : ANL_CODE,
      status : STATUS,
      nombre : NAME,
      prohibido : PROHIBIT_POSTING
    };
       $.ajax({
        url: "/nuevaDimension",
        dataType : "json",
        type : "post",
        data: param, 
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp){
          if(resp.success==1)
          {
            alert("Dimension creada exitosamente");
            actualizaDimensiones();
          }
          if(resp.success==2)
          {
            alert("Ya existe esa dimension!");
          }
           if(resp.success==3)
          {
            alert("Hubo un error desconocido! Perdon!");
          }
        }
      });
    }
    else
    {
      var param = {
        tipo : ANL_CAT_ID,
        codigo : ANL_CODE,
        status : STATUS,
        nombre : NAME,
        prohibido : PROHIBIT_POSTING
      };
      $.ajax({
        url: "/editarDimension",
        dataType : "json",
        type : "post",
        data: param, 
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp){
          if(resp.success==1)
          {
            alert("Tipo de dimension editado exitosamente");
            actualizaDimensiones();
          }
          if(resp.success==2)
          {
            alert("eh? no nos hackies por favor");
            actualizaDimensiones();
          }
          if(resp.success==3)
          {
            alert("Hubo un error!");
            actualizaDimensiones();
          }
        }
      });
    }
  });

  $(document).on("click", "#guardarConceptoCedula", function (){
    var idCedula = parseInt($("#cedulasSelectConceptos").val());
    var nombre = $("#nombreConcepto").val();
    if(idCedula>0)
    {
      var param = {nombre : nombre, idCedula : idCedula};
      $.ajax({
        url: "/guardarConceptoCedula",
        dataType : "json",
        type : "post",
        data: param, 
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp){
          if(resp.success==1)
          {
            alert("Linea cedula creada exitosamente");
            actualizaConceptosCedulas();
          }
          if(resp.success==0)
          {
            alert("Hubo un error!");
          }
        }
      });
    }
  });
  

  $(document).on("click", "#guardarNuevaLineaDelTipoDeDiario", function (){
    var IdTipoDeDiario = parseInt($("#tipoDeDiarioSelect").val());
    var linea = parseInt($("#lineaDelTipoDeDiario").val());
    var servicio = $("#servicio").val().trim();
    var cliente = $("#cliente").val().trim();
    var concepto = $("#concepto").val().trim();
    var descripcion = $("#descripcion").val().trim();
    var debofacturar = $('input[name=gender]:checked').val();
    var cuenta = $("#autocomplete-custom-append").val().trim();
    var tipoDC = $("#tipoDC").val();
    var ANAL_T0 = $("#ANAL_T0").val();
    var ANAL_T1 = $("#ANAL_T1").val();
    var ANAL_T2 = $("#ANAL_T2").val();
    var ANAL_T3 = $("#ANAL_T3").val();
    var ANAL_T4 = $("#ANAL_T4").val();
    var ANAL_T5 = $("#ANAL_T5").val();
    var ANAL_T6 = $("#ANAL_T6").val();
    var ANAL_T7 = $("#ANAL_T7").val();
    var ANAL_T8 = $("#ANAL_T8").val();
    var ANAL_T9 = $("#ANAL_T9").val();

      
    if(cuenta.indexOf('|')>-1)
    {
      cuenta = cuenta.split('|')[0];
    }
    if(ANAL_T0.indexOf('|')>-1)
    {
      ANAL_T0 = ANAL_T0.split('|')[0];
    }
    if(ANAL_T1.indexOf('|')>-1)
    {
      ANAL_T1 = ANAL_T1.split('|')[0];
    }
    if(ANAL_T2.indexOf('|')>-1)
    {
      ANAL_T2 = ANAL_T2.split('|')[0];
    }
    if(ANAL_T3.indexOf('|')>-1)
    {
      ANAL_T3 = ANAL_T3.split('|')[0];
    }
    if(ANAL_T4.indexOf('|')>-1)
    {
      ANAL_T4 = ANAL_T4.split('|')[0];
    }
    if(ANAL_T5.indexOf('|')>-1)
    {
      ANAL_T5 = ANAL_T5.split('|')[0];
    }
    if(ANAL_T6.indexOf('|')>-1)
    {
      ANAL_T6 = ANAL_T6.split('|')[0];
    }
    if(ANAL_T7.indexOf('|')>-1)
    {
      ANAL_T7 = ANAL_T7.split('|')[0];
    }
    if(ANAL_T8.indexOf('|')>-1)
    {
      ANAL_T8 = ANAL_T8.split('|')[0];
    }
    if(ANAL_T9.indexOf('|')>-1)
    {
      ANAL_T9 = ANAL_T9.split('|')[0];
    }
    if(cuenta=="")
    {
      alert("Tienes que seleccionar una cuenta primero!");
      return;
    }
    if(IdTipoDeDiario>0)
    {
        var param = {debofacturar: debofacturar, cliente: cliente, servicio: servicio, concepto: concepto, descripcion: descripcion, linea: linea, IdTipoDeDiario : IdTipoDeDiario, cuenta : cuenta , tipoDC : tipoDC, ANAL_T0 : ANAL_T0, ANAL_T1 : ANAL_T1, ANAL_T2 : ANAL_T2, ANAL_T3 : ANAL_T3, ANAL_T4 : ANAL_T4, ANAL_T5 : ANAL_T5, ANAL_T6 : ANAL_T6, ANAL_T7 : ANAL_T7, ANAL_T8 : ANAL_T8, ANAL_T9 : ANAL_T9};
       $.ajax({
        url: "/nuevoOModificaLineaTipoDeDiario",
        dataType : "json",
        type : "post",
        data: param, 
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp){
          if(resp.success==1)
          {
            alert("Linea de tipo de diario creada exitosamente");
            actualizaLaTablaDondeEstanLasLineasDelTipoDeDiario();
          }
          if(resp.success==0)
          {
            alert("Hubo un error!");
          }
        }
      });
    }
  });
  $(document).on("click", "#guardarNuevaLineaCedula", function (){
    var IdConcepto = parseInt($("#conceptoSelect").val());
    var nombre = $("#nombreLinea").val();
    var cuenta = $("#autocomplete-custom-append").val();
    var tipoDC = $("#tipoDC").val();
    var ANAL_T0 = $("#ANAL_T0").val();
    var ANAL_T1 = $("#ANAL_T1").val();
    var ANAL_T2 = $("#ANAL_T2").val();
    var ANAL_T3 = $("#ANAL_T3").val();
    var ANAL_T4 = $("#ANAL_T4").val();
    var ANAL_T5 = $("#ANAL_T5").val();
    var ANAL_T6 = $("#ANAL_T6").val();
    var ANAL_T7 = $("#ANAL_T7").val();
    var ANAL_T8 = $("#ANAL_T8").val();
    var ANAL_T9 = $("#ANAL_T9").val();

    if(cuenta.indexOf('|')>-1)
    {
      cuenta = cuenta.split('|')[0];
    }


    if(ANAL_T0.indexOf('|')>-1)
    {
      ANAL_T0 = ANAL_T0.split('|')[0];
    }


    if(ANAL_T1.indexOf('|')>-1)
    {
      ANAL_T1 = ANAL_T1.split('|')[0];
    }


    if(ANAL_T2.indexOf('|')>-1)
    {
      ANAL_T2 = ANAL_T2.split('|')[0];
    }


    if(ANAL_T3.indexOf('|')>-1)
    {
      ANAL_T3 = ANAL_T3.split('|')[0];
    }


    if(ANAL_T4.indexOf('|')>-1)
    {
      ANAL_T4 = ANAL_T4.split('|')[0];
    }


    if(ANAL_T5.indexOf('|')>-1)
    {
      ANAL_T5 = ANAL_T5.split('|')[0];
    }


    if(ANAL_T6.indexOf('|')>-1)
    {
      ANAL_T6 = ANAL_T6.split('|')[0];
    }


    if(ANAL_T7.indexOf('|')>-1)
    {
      ANAL_T7 = ANAL_T7.split('|')[0];
    }


    if(ANAL_T8.indexOf('|')>-1)
    {
      ANAL_T8 = ANAL_T8.split('|')[0];
    }


    if(ANAL_T9.indexOf('|')>-1)
    {
      ANAL_T9 = ANAL_T9.split('|')[0];
    }
    if(IdConcepto>0)
    {
      var param = {nombre : nombre, IdConcepto : IdConcepto, cuenta : cuenta , tipoDC : tipoDC, ANAL_T0 : ANAL_T0, ANAL_T1 : ANAL_T1, ANAL_T2 : ANAL_T2, ANAL_T3 : ANAL_T3, ANAL_T4 : ANAL_T4, ANAL_T5 : ANAL_T5, ANAL_T6 : ANAL_T6, ANAL_T7 : ANAL_T7, ANAL_T8 : ANAL_T8, ANAL_T9 : ANAL_T9};
      $.ajax({
        url: "/nuevoLineaCedula",
        dataType : "json",
        type : "post",
        data: param, 
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp){
          if(resp.success==1)
          {
            alert("Linea cedula creada exitosamente");
            actualizaLaTablaDondeEstanLasLineasDeLasCedulaSegunElConcepto();
          }
          if(resp.success==0)
          {
            alert("Hubo un error!");
          }
        }
      });
    }
    else
    {
     /* var param = {ANL_CAT_ID: ANL_CAT_ID, S_HEAD : S_HEAD, DESCR : DESCR};
      $.ajax({
        url: "/editarTipoDeDimensiones",
        dataType : "json",
        type : "post",
        data: param, 
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp){
          if(resp.success==1)
          {
            alert("Tipo de dimension editado exitosamente");
            actualizaTipoDeDimensiones();
          }
          
          if(resp.success==2)
          {
            alert("Hubo un error!");
            actualizaTipoDeDimensiones();
          }
        }
      });*/
    }
  });


  
    $(document).on("click", "#guardarNuevoTipoDeDiario", function (){
    var idTipoDeDiario = $(this).attr("idTipoDeDiario");
    var nombre = $("#nombre").val();
    var codigo = $("#codigo").val();
    
    if(idTipoDeDiario=="-1")
    {
      var param = {nombre : nombre, codigo: codigo};
      $.ajax({
        url: "/nuevoTipoDeDiario",
        dataType : "json",
        type : "post",
        data: param, 
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp){
          if(resp.success==1)
          {
            alert("Tipo de diario creado exitosamente");
            actualizaTiposDeDiario();
          }
          if(resp.success==0)
          {
            alert("Hubo un error! perdon!");
          }
        }
      });
    }
    else
    {
     /* var param = {ANL_CAT_ID: ANL_CAT_ID, S_HEAD : S_HEAD, DESCR : DESCR};
      $.ajax({
        url: "/editarTipoDeDimensiones",
        dataType : "json",
        type : "post",
        data: param, 
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp){
          if(resp.success==1)
          {
            alert("Tipo de dimension editado exitosamente");
            actualizaTipoDeDimensiones();
          }
          
          if(resp.success==2)
          {
            alert("Hubo un error!");
            actualizaTipoDeDimensiones();
          }
        }
      });*/
    }
  });
    
    $(document).on("change", "#tipoDeDiarioSelectDeCaptura", function (){
      $("#referencia").select();
      });
  $(document).on("click", "#guardarNuevaCedula", function (){
    var idCedula = $(this).attr("idCedula");
    var nombre = $("#nombre").val();
    if(idCedula=="-1")
    {
      var param = {nombre : nombre};
      $.ajax({
        url: "/nuevoCedula",
        dataType : "json",
        type : "post",
        data: param, 
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp){
          if(resp.success==1)
          {
            alert("Cedula creada exitosamente");
            actualizaCedulas();
          }
          if(resp.success==0)
          {
            alert("Hubo un error!");
          }
        }
      });
    }
    else
    {
     /* var param = {ANL_CAT_ID: ANL_CAT_ID, S_HEAD : S_HEAD, DESCR : DESCR};
      $.ajax({
        url: "/editarTipoDeDimensiones",
        dataType : "json",
        type : "post",
        data: param, 
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp){
          if(resp.success==1)
          {
            alert("Tipo de dimension editado exitosamente");
            actualizaTipoDeDimensiones();
          }
          
          if(resp.success==2)
          {
            alert("Hubo un error!");
            actualizaTipoDeDimensiones();
          }
        }
      });*/
    }
  });

  function toFixed(value, precision) {
    var power = Math.pow(10, precision || 0);
    return String(Math.round(value * power) / power);
}
function format2(n, currency) {
    return currency + " " + toFixed(n,2).replace(/(\d)(?=(\d{3})+\.)/g, "$1,");
}

 $(document).on("click", ".vesAlDiario", function (){
    var diario = parseInt($(this).attr("diario"));
    window.open(
      "/generarDiario?diario=" + diario,
      '_blank'
    );
  });
 $(document).on("click", ".vesADetallePorDiario", function (){
    var idLinea = parseInt($(this).attr("idLinea"));
    var delPeriodo = $("#delPeriodo").val();
    var alPeriodo = $("#alPeriodo").val();
    if(idLinea>-1)
    {
      var param = {idLinea : idLinea, delPeriodo: delPeriodo, alPeriodo : alPeriodo};
      $.ajax({
        url: "/generaCedulaPorLinea",
        dataType : "json",
        type : "post",
        data: param, 
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp){
          var aver2 = resp.cedulasAcumulador;
          var cad='';
          for (var p in aver2) {
            var objeto=resp.cedulasAcumulador[p];
            var diario = p.split('-')[0];
               cad=cad+'<tr><td><a class="vesAlDiario" href="#" diario="'+diario+'" >'+p+'</a></td><td>'+format2(parseFloat(objeto["Cantidad"]),"$")+'</td></tr>';
          }
          $("#tablaConcentradoDiarios").html(cad);
        }
      });
    }
  });
 $(document).on("click", ".vesADetallePorCuenta", function (){
    var idConcepto = parseInt($(this).attr("idConcepto"));
    var delPeriodo = $("#delPeriodo").val();
    var alPeriodo = $("#alPeriodo").val();
    if(idConcepto>-1)
    {
      var param = {idConcepto : idConcepto, delPeriodo: delPeriodo, alPeriodo : alPeriodo};
      $.ajax({
        url: "/generaCedulaPorConcepto",
        dataType : "json",
        type : "post",
        data: param, 
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp){
          var aver2 = resp.cedulasAcumulador;
          var cad='';
          for (var p in aver2) {
            var objeto=resp.cedulasAcumulador[p];
               cad=cad+'<tr><td><a class="vesADetallePorDiario" href="#" idLinea="'+objeto["IdConceptoOLinea"]+'" >'+p+'</a></td><td>'+format2(parseFloat(objeto["Cantidad"]),"$")+'</td></tr>';
          }
          $("#tablaConcentradoCuenta").html(cad);
        }
      });
    }
  });
  $(document).on("click", "#verCedulaButton", function (){
    var idCedula = parseInt($("#cedulasSelectVer").val());
    var delPeriodo = $("#delPeriodo").val();
    var alPeriodo = $("#alPeriodo").val();
    if(idCedula>-1)
    {
      var param = {idCedula : idCedula, delPeriodo: delPeriodo, alPeriodo : alPeriodo};
      $.ajax({
        url: "/generaCedula",
        dataType : "json",
        type : "post",
        data: param, 
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp){
          var aver2 = resp.cedulasAcumulador;
          var cad='';
          for (var p in aver2) {
            var objeto=resp.cedulasAcumulador[p];
               cad=cad+'<tr><td><a class="vesADetallePorCuenta" href="#" idConcepto="'+objeto["IdConceptoOLinea"]+'" >'+p+'</a></td><td>'+format2(parseFloat(objeto["Cantidad"]),"$")+'</td></tr>';
          }
          $("#tablaConcentradoCedula").html(cad);
        }
      });
    }
  });
  $(document).on("click", "#guardarNuevoTipoDeDimension", function (){
    var ANL_CAT_ID = $(this).attr("ANL_CAT_ID");
    var S_HEAD = $("#codigo").val();
    var DESCR = $("#descr").val();
    if(ANL_CAT_ID=="-1")
    {
      var param = {S_HEAD : S_HEAD, DESCR : DESCR};
      $.ajax({
        url: "/nuevoTipoDeDimensiones",
        dataType : "json",
        type : "post",
        data: param, 
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp){
          if(resp.success==1)
          {
            alert("Tipo de dimension creada exitosamente");
            actualizaTipoDeDimensiones();
          }
          if(resp.success==2)
          {
            alert("Hubo un error!");
            actualizaTipoDeDimensiones();
          }
        }
      });
    }
    else
    {
      var param = {ANL_CAT_ID: ANL_CAT_ID, S_HEAD : S_HEAD, DESCR : DESCR};
      $.ajax({
        url: "/editarTipoDeDimensiones",
        dataType : "json",
        type : "post",
        data: param, 
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp){
          if(resp.success==1)
          {
            alert("Tipo de dimension editado exitosamente");
            actualizaTipoDeDimensiones();
          }
          
          if(resp.success==2)
          {
            alert("Hubo un error!");
            actualizaTipoDeDimensiones();
          }
        }
      });
    }
  });
  

$(document).on("click", ".editarDimension", function (){
  var ANL_CAT_ID = $(this).attr("ANL_CAT_ID").trim();
  var S_HEAD = $(this).attr("S_HEAD").trim();
  var NAME = $(this).attr("NAME").trim();
  var ANL_CODE = $(this).attr("ANL_CODE").trim();
  var STATUS = $(this).attr("STATUS").trim();
  var PROHIBIT_POSTING = $(this).attr("PROHIBIT_POSTING").trim();
  $("#guardarNuevaDimension").attr("ANL_CAT_ID",ANL_CAT_ID);
  $("#guardarNuevaDimension").attr("ANL_CODE",ANL_CODE);
  $("#codigo").val(ANL_CODE);
  $("#tipo").val(S_HEAD);
  $("#nombre").val(NAME);
  $("#status").val(STATUS);
  $("#prohibido").val(PROHIBIT_POSTING);
  $("#tipo").attr("readonly", true);
  $("#codigo").attr("readonly", true);
});

  $(document).on("click", ".editarTipoDeDimension", function (){
    var ANL_CAT_ID = $(this).attr("ANL_CAT_ID");
    var S_HEAD = $(this).attr("S_HEAD");
    var DESCR = $(this).attr("DESCR");
    $("#guardarNuevoTipoDeDimension").attr("ANL_CAT_ID",ANL_CAT_ID);
    $("#codigo").val(S_HEAD);
    $("#descr").val(DESCR);
  });

  
  $(document).on("click", "#limpiar", function (){
    $("#guardarNuevoTipoDeDimension").attr("ANL_CAT_ID","-1");
    $("#codigo").val("");
    $("#descr").val("");
  });

$(document).on("click", "#limpiarD", function (){
    $("#guardarNuevaDimension").attr("ANL_CAT_ID","-1");
    $("#guardarNuevaDimension").attr("ANL_CODE","-1");
    $("#tipo").val("");
    $("#codigo").val("");
    $("#tipo").attr("readonly", false);
    $("#codigo").attr("readonly", false);
    $("#nombre").val("");
    $("#status").val("");
    $("#prohibido").val("");
  });

$(document).on("click", "#subidaInicialCuentas", function (){
    var cad = '<div class="">'+
          '<div class="page-title">'+
            '<div class="title_left">'+
              '<h3>Subida inicial de cuentas</h3>'+
            '</div>'+
          '</div>'+
          '<div class="clearfix"></div>'+
          '<div class="row">'+
            '<div class="col-md-12 col-sm-12 col-xs-12">'+
              '<div class="x_panel">'+
                '<div class="x_title">'+
                  '<h2>Cuentas <small></small></h2>'+
                  '<ul class="nav navbar-right panel_toolbox">'+
                    '<li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>'+
                    '</li>'+
                    '<li><a class="close-link"><i class="fa fa-close"></i></a>'+
                    '</li>'+
                  '</ul>'+
                  '<div class="clearfix"></div>'+
                '</div>'+
                '<div class="x_content">'+
                    '<p>Descarga esta lista de ejemplo de las cuentas, llenala y conviertela en archivo CSV, despu&eacute;s subela aqu&iacute;.</p>'+
                    '<a href="static/files/FormatoParaSubirCuentas.xlsx" target="_blank"><button type="button" class="btn btn-primary">Descargar lista de ejemplo</button></a>'+
                    '<input type="file"  id="archivoCSVCuentas" />'+
                    '<br>'+
                    '<br>'+
                    '<br>'+
                    '<br>'+
                '</div>'+
                '<div class="clearfix"></div>'+
              '</div>'+
            '</div>'+
          '</div>'+
        '</div>';     
      $("#mainContent").html(cad);
 });
  $(document).on("click", "#subidaInicialDimensiones", function (){
    var cad = '<div class="">'+
          '<div class="page-title">'+
            '<div class="title_left">'+
              '<h3>Tipo de dimensiones</h3>'+
            '</div>'+
          '</div>'+
          '<div class="clearfix"></div>'+
          '<div class="row">'+
            '<div class="col-md-12 col-sm-12 col-xs-12">'+
              '<div class="x_panel">'+
                '<div class="x_title">'+
                  '<h2>Tipo de dimensiones <small></small></h2>'+
                  '<ul class="nav navbar-right panel_toolbox">'+
                    '<li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>'+
                    '</li>'+
                    '<li><a class="close-link"><i class="fa fa-close"></i></a>'+
                    '</li>'+
                  '</ul>'+
                  '<div class="clearfix"></div>'+
                '</div>'+
                '<div class="x_content">'+
                    '<p>Descarga esta lista de ejemplo, llenala y conviertela en archivo CSV, despu&eacute;s subela aqu&iacute;.</p>'+
                    '<a href="static/files/FormatoParaSubirDimensiones.xlsx" target="_blank"><button type="button" class="btn btn-primary">Descargar lista de ejemplo</button></a>'+
                    '<input type="file"  id="archivoCSV" />'+
                    '<br>'+
                    '<br>'+
                    '<br>'+
                    '<br>'+
                '</div>'+
                '<div class="clearfix"></div>'+
              '</div>'+
            '</div>'+
          '</div>'+
        '</div>';     
      $("#mainContent").html(cad);
});
  

$(document).on("change", "#TipoDondeEstaranLasDimensiones", function (){  
  var numeroTipo = parseInt($(this).val());
  $('#entry1').val(0);
  $('#entry2').val(0);
  $('#entry3').val(0);
  $('#entry4').val(0);
  $('#entry5').val(0);
  $('#entry6').val(0);
  $('#entry7').val(0);
  $('#entry8').val(0);
  $('#entry9').val(0);
  $('#entry10').val(0);
  if(numeroTipo==0)
  {
//clasificacionDeDimensiones
  }
  else
  {
    var param = {numero : numeroTipo}
    $.ajax({
      url: "/obtieneDimensionesAsignadasATipo",
      dataType : "json",
      type : "post",
      data: param,
      async : true,
      beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
      },
      complete : function (){
      }, 
      success : function (resp){
        var i;
        var aver2 = Object.keys(resp.dimensiones);
        var cad = '';
        var numero, S_HEAD, ANL_CAT_ID;
         for(i=0;i<aver2.length;i++)
        {
          numero = parseInt(resp.dimensiones[aver2[i]]["ENTRY_NUM"]);
          //S_HEAD = resp.dimensiones[aver2[i]]["S_HEAD"];
          ANL_CAT_ID = resp.dimensiones[aver2[i]]["ANL_CAT_ID"];
          $('#entry'+numero).val(ANL_CAT_ID);
          //$('#startTime option[value=17:00:00]').prop('selected', 'selected');
        } 
      }
    });  
  }

});
$(document).on("change", "#archivoCSVCuentas", function (){  
  var fileName = $('#archivoCSVCuentas').val();
  var extension = fileName.substr(fileName.length-3);
  if(extension == 'csv' || extension == 'txt' )
  {
    var results = Papa.parse($('#archivoCSVCuentas')[0].files[0], {
      complete: function(filas) {
        var i;
        var exitosos = 0;
        var yaExistian = 0;
        var conErrores = 0;
        var cuantosTotales = filas.data.length-2;
        for(i=1;i<filas.data.length;i++)
        {
          var unaFila = filas.data[i];
          if(unaFila[0]=="D"){unaFila[0]=0;}
          if(unaFila[0]=="C"){unaFila[0]=1;}
          if(unaFila[0]=="T"){unaFila[0]=2;}
          if(unaFila[0]=="P"){unaFila[0]=3;}
          if(unaFila[0]=="B"){unaFila[0]=4;}
          if(unaFila[0]=="M"){unaFila[0]=5;}

          if(unaFila[6]=="B"){unaFila[6]=0;}
          if(unaFila[6]=="O"){unaFila[6]=1;}

          if(unaFila[7]=="O"){unaFila[7]=0;}
          if(unaFila[7]=="C"){unaFila[7]=3;}
         
         var j=0;
         for(j=15;j<=24;j++)
         {
          if(unaFila[j]=="P"){unaFila[j]=3;}
          if(unaFila[j]=="O"){unaFila[j]=2;}
          if(unaFila[j]=="M"){unaFila[j]=1;}
         }
          

         
         //cuacua
            var param = {
              ACNT_CODE : unaFila[1].trim(),
              DESCR : unaFila[2],
              ACNT_TYPE : unaFila[0],
              BAL_TYPE : unaFila[6],
              STATUS : unaFila[7],
              LINK_ACNT : unaFila[12],
              ENTER_ANL_1 : unaFila[15],
              ENTER_ANL_2 : unaFila[16],
              ENTER_ANL_3 : unaFila[17],
              ENTER_ANL_4 : unaFila[18],
              ENTER_ANL_5 : unaFila[19],
              ENTER_ANL_6 : unaFila[20],
              ENTER_ANL_7 : unaFila[21],
              ENTER_ANL_8 : unaFila[22],
              ENTER_ANL_9 : unaFila[23],
              ENTER_ANL_10 : unaFila[24],
              D1 : unaFila[29],
              D2 : unaFila[30],
              D3 : unaFila[31],
              D4 : unaFila[32],
              D5 : unaFila[33],
              D6 : unaFila[34],
              D7 : unaFila[35],
              D8 : unaFila[36],
              D9 : unaFila[37],
              D10 : unaFila[38]
            };
             $.ajax({
                url: "/nuevaCuenta",
                dataType : "json",
                type : "post",
                data: param,
                async : false,
                beforeSend : function (){
            //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
                },
                complete : function (){
                }, 
                success : function (resp){
                  if(resp.success==1)
                  {
                    exitosos++;
                  }
                  if(resp.success==2)
                  {
                    yaExistian++;
                  }
                  if(resp.success==3)
                  {
                    conErrores++;
                  }
                }
              });  
          
        }
        alert("Exitosos: "+exitosos+", Ya existian: "+yaExistian+", con error: "+conErrores+", de un total de: "+cuantosTotales);  
      }
    });
  }
  else
  {
    alert('El formato de archivo permitido es csv o txt.');
  }
});

$(document).on("change", "#archivoCSV", function (){  
  var fileName = $('#archivoCSV').val();
  var extension = fileName.substr(fileName.length-3);
  if(extension == 'csv' || extension == 'txt' )
  {
    var results = Papa.parse($('#archivoCSV')[0].files[0], {
      complete: function(filas) {
        var i;
        var exitosos = 0;
        var yaExistian = 0;
        var conErrores = 0;
        var cuantosTotales = filas.data.length-2;
        for(i=1;i<filas.data.length;i++)
        {
          var unaFila = filas.data[i];
         
            var param = {
              tipo : unaFila[0],
              codigo : unaFila[1],
              status : unaFila[2],
              nombre : unaFila[3],
              prohibido : unaFila[4]
            };
             $.ajax({
                url: "/nuevaDimension",
                dataType : "json",
                type : "post",
                data: param,
                async : false,
                beforeSend : function (){
            //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
                },
                complete : function (){
                }, 
                success : function (resp){
                  if(resp.success==1)
                  {
                    exitosos++;
                  }
                  if(resp.success==2)
                  {
                    yaExistian++;
                  }
                  if(resp.success==3)
                  {
                    conErrores++;
                  }
                }
              });  
          
        }
        alert("Exitosos: "+exitosos+", Ya existian: "+yaExistian+", con error: "+conErrores+", de un total de: "+cuantosTotales);  
      }
    });
  }
  else
  {
    alert('El formato de archivo permitido es csv o txt.');
  }
});


$(document).on("change", "#fechaTrans", function (){
  $("#cantidad").select();
  $("#cantidad").focus();
  $("#cantidad").click();
});  


  $(document).on("click", "#depre", function (){
    var cad = '<div class="">'+
          '<div class="page-title">'+
            '<div class="title_left">'+
              '<h3>Depreciaci&oacute;n de activos fijos</h3>'+
            '</div>'+
          '</div>'+
          '<div class="clearfix"></div>'+
          '<div class="row">'+
            '<div class="col-md-12 col-sm-12 col-xs-12">'+
              '<div class="x_panel">'+
                '<div class="x_title">'+
                  '<h2>Depreciaci&oacute;n <small></small></h2>'+
                  '<ul class="nav navbar-right panel_toolbox">'+
                    '<li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>'+
                    '</li>'+
                    '<li><a class="close-link"><i class="fa fa-close"></i></a>'+
                    '</li>'+
                  '</ul>'+
                  '<div class="clearfix"></div>'+
                '</div>'+
                '<div class="x_content">'+
                  '<br />'+
                    '<div class="form-group">'+
                      '<div class="col-md-6 col-sm-6 col-xs-6 col-md-offset-6">'+
                        '<button type="submit" id="depreciarTodaLaBanda"  class="btn btn-success">Depreciar todos</button>'+
                      '</div>'+
                    '</div>'+
                '</div>'+
                '<div class="clearfix"></div>'+
                '<div class="col-md-12 col-sm-12 col-xs-12">'+
                  '<div class="x_panel">'+
                    '<div class="x_title">'+
                      '<h2>Activos <small>fijos</small></h2>'+
                      '<div class="clearfix"></div>'+
                    '</div>'+
                    '<div class="x_content">'+
                      '<table class="table table-hover">'+
                        '<thead><tr><th>C&oacute;digo</th><th>Nombre</th><th>Status</th><th>Status</th><th>Periodo inicial</th><th>Periodo final</th><th>&Uacute;ltimo periodo</th><th>Periodo disposal</th><th>Disposed</th><th>Cantidad inicial base</th><th>Cantidad depreciado base</th><th>Cantidad neta base</th><th>Porcentaje de depreciaci&oacute;n base</th><th>Cantidad inicial transacci&oacute;n</th><th>Cantidad depreciado transacci&oacute;n</th><th>Cantidad neta transacci&oacute;n</th><th>Porcentaje de depreciaci&oacute;n transacci&oacute;n</th></thead>'+
                        '<tbody id="tablaActivosFijosDepre"></tbody>'+
                      '</table>'+
                    '</div>'+
                  '</div>'+
                '</div>'+
              '</div>'+
            '</div>'+
          '</div>'+
        '</div>';     
      $("#mainContent").html(cad);
      actualizaActivosFijos();
 });

   $(document).on("click", "#modificarDimensiones", function (){
    var cad = '<div class="">'+
          '<div class="page-title">'+
            '<div class="title_left">'+
              '<h3>Tipo de dimensiones</h3>'+
            '</div>'+
          '</div>'+
          '<div class="clearfix"></div>'+
          '<div class="row">'+
            '<div class="col-md-12 col-sm-12 col-xs-12">'+
              '<div class="x_panel">'+
                '<div class="x_title">'+
                  '<h2>Tipo de dimensiones <small></small></h2>'+
                  '<ul class="nav navbar-right panel_toolbox">'+
                    '<li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>'+
                    '</li>'+
                    '<li><a class="close-link"><i class="fa fa-close"></i></a>'+
                    '</li>'+
                  '</ul>'+
                  '<div class="clearfix"></div>'+
                '</div>'+
                '<div class="x_content">'+
                  '<br />'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Tipo:</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text"  name="tipo" id="tipo" class="form-control col-md-10" style="float: left;" />'+
                       '</div>'+
                    '</div>'+
                     '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">C&oacute;digo:</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text"  name="codigo" id="codigo" class="form-control col-md-10" style="float: left;" />'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Nombre:</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" name="nombre" id="nombre" class="form-control col-md-10" style="float: left;" />'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Status: (0 normal, 3 cerrado)</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" name="status" maxlength="1"  id="status" class="form-control col-md-10" style="float: left;" />'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Prohibido contabilizar: (0 no, 1 prohibido contabilizar)</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" name="prohibido" maxlength="1" id="prohibido" class="form-control col-md-10" style="float: left;" />'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<div class="col-md-6 col-sm-6 col-xs-6 col-md-offset-6">'+
                        '<button type="submit" id="guardarNuevaDimension" ANL_CAT_ID="-1" ANL_CODE="-1"  class="btn btn-success">Guardar nuevo o modificacion</button>'+
                      '</div>'+
                      '<div class="col-md-6 col-sm-6 col-xs-6 col-md-offset-6">'+
                        '<button type="submit" id="limpiarD" class="btn btn-warning">Limpiar</button>'+
                      '</div>'+
                    '</div>'+
                '</div>'+
                '<div class="clearfix"></div>'+
                '<div class="col-md-12 col-sm-12 col-xs-12">'+
                  '<div class="x_panel">'+
                    '<div class="x_title">'+
                      '<h2>Dimensiones <small>Tipos de dimensiones</small></h2>'+
                      '<div class="clearfix"></div>'+
                    '</div>'+
                    '<div class="x_content">'+
                      '<table class="table table-hover">'+
                        '<thead><tr><th>Tipo</th><th>C&oacute;digo</th><th>Nombre</th><th>Status</th><th>Prohibido contabilizar</th><th>Editar</th></thead>'+
                        '<tbody id="tablaDimensionesEditar"></tbody>'+
                      '</table>'+
                    '</div>'+
                  '</div>'+
                '</div>'+
              '</div>'+
            '</div>'+
          '</div>'+
        '</div>';     
      $("#mainContent").html(cad);
      actualizaDimensiones();
 });

$(document).on("click", "#asignarDimension", function (){

    var cad = '<div class="">'+
          '<div class="page-title">'+
            '<div class="title_left">'+
              '<h3>Tipo de dimensiones</h3>'+
            '</div>'+
          '</div>'+
          '<div class="clearfix"></div>'+
          '<div class="row">'+
            '<div class="col-md-12 col-sm-12 col-xs-12">'+
              '<div class="x_panel">'+
                '<div class="x_title">'+
                  '<h2>Tipo de dimensiones <small></small></h2>'+
                  '<ul class="nav navbar-right panel_toolbox">'+
                    '<li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>'+
                    '</li>'+
                    '<li><a class="close-link"><i class="fa fa-close"></i></a>'+
                    '</li>'+
                  '</ul>'+
                  '<div class="clearfix"></div>'+
                '</div>'+
                '<div class="x_content">'+
                  '<br />'+
                    '<div class="form-group">'+
                      '<p><label for="TipoDondeEstaranLasDimensiones">Tipo:</label>'+
                          '<select id="TipoDondeEstaranLasDimensiones" class="form-control" required="" >'+
                          '<option value="0">Seleccione una opci&oacute;n</option>'+
                          '<option value="901">Dimensiones para diarios</option>'+
                          '<option value="204">Dimensiones para activos</option>'+
                          '<option value="152">Dimensiones para cuentas</option>'+
                          '</select></p>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<p><label for="entry1">Dimensi&oacute;n 1:</label>'+
                          '<select id="entry1" class="form-control" required="" ></select></p>'+
                    '</div>'+
                     '<div class="form-group">'+
                      '<p><label for="entry2">Dimensi&oacute;n 2:</label>'+
                          '<select id="entry2" class="form-control" required="" ></select></p>'+
                    '</div>'+
                     '<div class="form-group">'+
                      '<p><label for="entry3">Dimensi&oacute;n 3:</label>'+
                          '<select id="entry3" class="form-control" required="" ></select></p>'+
                    '</div>'+
                     '<div class="form-group">'+
                      '<p><label for="entry4">Dimensi&oacute;n 4:</label>'+
                          '<select id="entry4" class="form-control" required="" ></select></p>'+
                    '</div>'+
                     '<div class="form-group">'+
                      '<p><label for="entry5">Dimensi&oacute;n 5:</label>'+
                          '<select id="entry5" class="form-control" required="" ></select></p>'+
                    '</div>'+
                     '<div class="form-group">'+
                      '<p><label for="entry6">Dimensi&oacute;n 6:</label>'+
                          '<select id="entry6" class="form-control" required="" ></select></p>'+
                    '</div>'+
                     '<div class="form-group">'+
                      '<p><label for="entry7">Dimensi&oacute;n 7:</label>'+
                          '<select id="entry7" class="form-control" required="" ></select></p>'+
                    '</div>'+
                     '<div class="form-group">'+
                      '<p><label for="entry8">Dimensi&oacute;n 8:</label>'+
                          '<select id="entry8" class="form-control" required="" ></select></p>'+
                    '</div>'+
                     '<div class="form-group">'+
                      '<p><label for="entry9">Dimensi&oacute;n 9:</label>'+
                          '<select id="entry9" class="form-control" required="" ></select></p>'+
                    '</div>'+
                     '<div class="form-group">'+
                      '<p><label for="entry10">Dimensi&oacute;n 10:</label>'+
                          '<select id="entry10" class="form-control" required="" ></select></p>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<div class="col-md-6 col-sm-6 col-xs-6 col-md-offset-6">'+
                        '<button type="submit" id="guardarClasificacionDimensiones" ANL_CAT_ID="-1" class="btn btn-success">Guardar</button>'+
                      '</div>'+
                    '</div>'+
                '</div>'+
              '</div>'+
            '</div>'+
          '</div>'+
        '</div>';     
      $("#mainContent").html(cad);
      actualizaClasificacionDimensiones();
 });




$(document).on("click", "#verDiariosReverseados", function (){
    var cad = '<div class="">'+
          '<div class="page-title">'+
            '<div class="title_left">'+
              '<h3>Ver diarios reverseados</h3>'+
            '</div>'+
          '</div>'+
          '<div class="clearfix"></div>'+
          '<div class="row">'+
            '<div class="col-md-12 col-sm-12 col-xs-12">'+
              '<div class="x_panel">'+
                '<div class="clearfix"></div>'+
                '<div class="col-md-12 col-sm-12 col-xs-12">'+
                  '<div class="x_panel">'+
                    '<div class="x_title">'+
                      '<h2>Ver <small>Diarios reversiados</small></h2>'+
                      '<div class="clearfix"></div>'+
                    '</div>'+
                    '<div class="x_content">'+
                      '<table class="table table-hover">'+
                        '<thead><tr><th>Diario</th><th>Referencia</th><th>Periodo</th></thead>'+
                        '<tbody id="tablaDiariosReversiados"></tbody>'+
                      '</table>'+
                    '</div>'+
                  '</div>'+
                '</div>'+
              '</div>'+
            '</div>'+
          '</div>'+
        '</div>';     
      $("#mainContent").html(cad);
      configInicialDeDiariosReversiados();
 });
$(document).on("click", "#crearConceptos", function (){
    var cad = '<div class="">'+
          '<div class="page-title">'+
            '<div class="title_left">'+
              '<h3>Crear conceptos de c&eacute;dula</h3>'+
            '</div>'+
          '</div>'+
          '<div class="clearfix"></div>'+
          '<div class="row">'+
            '<div class="col-md-12 col-sm-12 col-xs-12">'+
              '<div class="x_panel">'+
                '<div class="x_title">'+
                  '<h2>Crear conceptos para c&eacute;dulas<small></small></h2>'+
                  '<ul class="nav navbar-right panel_toolbox">'+
                    '<li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>'+
                    '</li>'+
                    '<li><a class="close-link"><i class="fa fa-close"></i></a>'+
                    '</li>'+
                  '</ul>'+
                  '<div class="clearfix"></div>'+
                '</div>'+
                '<div class="x_content">'+
                  '<br />'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Elige cedula:</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<select style="width: 100%;font-size: 20px;" id="cedulasSelectConceptos"></select>'+
                       '</div>'+
                    '</div>'+
                   '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Nombre:</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" name="nombreConcepto" id="nombreConcepto" class="form-control col-md-10" style="float: left;" />'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<div class="col-md-6 col-sm-6 col-xs-6 col-md-offset-6">'+
                        '<button type="submit" id="guardarConceptoCedula" idCedula="-1" class="btn btn-success">Guardar concepto</button>'+
                      '</div>'+
                    '</div>'+
                '</div>'+
                '<div class="clearfix"></div>'+
                '<div class="col-md-12 col-sm-12 col-xs-12">'+
                  '<div class="x_panel">'+
                    '<div class="x_title">'+
                      '<h2>Ver <small>C&eacute;dulas</small></h2>'+
                      '<div class="clearfix"></div>'+
                    '</div>'+
                    '<div class="x_content">'+
                      '<table class="table table-hover">'+
                        '<thead><tr><th>#</th><th>Nombre</th></thead>'+
                        '<tbody id="tablaCedulaConceptos"></tbody>'+
                      '</table>'+
                    '</div>'+
                  '</div>'+
                '</div>'+
              '</div>'+
            '</div>'+
          '</div>'+
        '</div>';     
      $("#mainContent").html(cad);
      configInicialDeCedulasConceptos();
 });
$(document).on("click", "#verCedula", function (){
    var cad = '<div class="">'+
          '<div class="page-title">'+
            '<div class="title_left">'+
              '<h3>Ver c&eacute;dula</h3>'+
            '</div>'+
          '</div>'+
          '<div class="clearfix"></div>'+
          '<div class="row">'+
            '<div class="col-md-12 col-sm-12 col-xs-12">'+
              '<div class="x_panel">'+
                '<div class="x_title">'+
                  '<h2>Ver c&eacute;dulas<small></small></h2>'+
                  '<ul class="nav navbar-right panel_toolbox">'+
                    '<li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>'+
                    '</li>'+
                    '<li><a class="close-link"><i class="fa fa-close"></i></a>'+
                    '</li>'+
                  '</ul>'+
                  '<div class="clearfix"></div>'+
                '</div>'+
                '<div class="x_content">'+
                  '<br />'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Elige cedula:</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<select style="width: 100%;font-size: 20px;" id="cedulasSelectVer"></select>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<p><label for="delPeriodo">Del periodo:</label>'+
                          '<select id="delPeriodo" class="form-control" required="" >'+
                            '</select></p>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<p><label for="alPeriodo">Al periodo:</label>'+
                          '<select id="alPeriodo" class="form-control" required="">'+
                            '</select></p>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<div class="col-md-6 col-sm-6 col-xs-6 col-md-offset-6">'+
                        '<button type="submit" id="verCedulaButton" idCedula="-1" class="btn btn-success">Generar C&eacute;dula</button>'+
                      '</div>'+
                    '</div>'+
                '</div>'+
                '<div class="clearfix"></div>'+
                '<div class="col-md-4 col-sm-4 col-xs-4">'+
                  '<div class="x_panel">'+
                    '<div class="x_title">'+
                      '<h2>Ver <small>C&eacute;dulas</small></h2>'+
                      '<div class="clearfix"></div>'+
                    '</div>'+
                    '<div class="x_content">'+
                      '<table class="table table-hover">'+
                        '<thead><tr><th>Nombre</th><th>Cantidad</th></thead>'+
                        '<tbody id="tablaConcentradoCedula"></tbody>'+
                      '</table>'+
                    '</div>'+
                  '</div>'+
                '</div>'+
                '<div class="col-md-4 col-sm-4 col-xs-4">'+
                  '<div class="x_panel">'+
                    '<div class="x_title">'+
                      '<h2>Ver <small>Cuentas</small></h2>'+
                      '<div class="clearfix"></div>'+
                    '</div>'+
                    '<div class="x_content">'+
                      '<table class="table table-hover">'+
                        '<thead><tr><th>Nombre</th><th>Cantidad</th></thead>'+
                        '<tbody id="tablaConcentradoCuenta"></tbody>'+
                      '</table>'+
                    '</div>'+
                  '</div>'+
                '</div>'+
                '<div class="col-md-4 col-sm-4 col-xs-4">'+
                  '<div class="x_panel">'+
                    '<div class="x_title">'+
                      '<h2>Ver <small>Diarios</small></h2>'+
                      '<div class="clearfix"></div>'+
                    '</div>'+
                    '<div class="x_content">'+
                      '<table class="table table-hover">'+
                        '<thead><tr><th>Nombre</th><th>Cantidad</th></thead>'+
                        '<tbody id="tablaConcentradoDiarios"></tbody>'+
                      '</table>'+
                    '</div>'+
                  '</div>'+
                '</div>'+
              '</div>'+
            '</div>'+
          '</div>'+
        '</div>';     
      $("#mainContent").html(cad);
      configInicialDeCedulasParaVer();
 });


var lineaEnLaQueVoy=1;
var lineaGlobal = 1;
var referenciaActual="";
var debitosReferencia=0;
var creditosReferencia=0;
var cuentaAnterior="";
$(document).on('keydown', '#referencia', function(ev) {
    if(ev.which === 13) {
        empiezaReferencia();
        return false;
    }
});

function empiezaReferencia()
{
  referenciaActual = $("#referencia").val().trim();
  if(referenciaActual=="")
  {

  }
  else
  {
    $("#referencia").attr("disabled","disabled");
    $("#tipoDeDiarioSelectDeCaptura").attr("disabled","disabled");
    $("#infoLabel").html("D&eacute;bitos: "+format2(debitosReferencia,"$")+" Cr&eacute;ditos: "+format2(creditosReferencia,"$")+" Linea: "+lineaEnLaQueVoy);
    idTipoDeDiarioGlobal=$("#tipoDeDiarioSelectDeCaptura").val();
    
    ponLineaSegunTipoDeDiario(idTipoDeDiarioGlobal , lineaEnLaQueVoy);
  }
}
function ponDimensionesSegunCuenta()
{
  var cuenta = $("#autocomplete-custom-append").val().trim();
  if(cuenta.indexOf('|')>-1)
  {
    cuenta = cuenta.split('|')[0];
  }
  if(cuenta=="")
  {
    return;
  }
  if(cuenta==cuentaAnterior)
  {
    return;
  }
  var param = {cuenta : cuenta};
  $.ajax({
    url: "/dameDimensionesDisponiblesSegunLaCuenta",
    dataType : "json",
    type : "post",
    data: param,
    async : true,
    beforeSend : function (){
    },
    complete : function (){
    }, 
    success : function (resp){
      if(resp.success==1)
      {
        if(resp.status!=0)
        {
          alert("En la cuenta "+cuenta+" no se permite contabilizar.");
          $("#autocomplete-custom-append").val("");
          return;
        }
        cuentaAnterior=cuenta;
        $("#divD0").css("display","none");
        $("#divD1").css("display","none");
        $("#divD2").css("display","none");
        $("#divD3").css("display","none");
        $("#divD4").css("display","none");
        $("#divD5").css("display","none");
        $("#divD6").css("display","none");
        $("#divD7").css("display","none");
        $("#divD8").css("display","none");
        $("#divD9").css("display","none");
       /* $("#ANAL_T0").val("");
        $("#ANAL_T1").val("");
        $("#ANAL_T2").val("");
        $("#ANAL_T3").val("");
        $("#ANAL_T4").val("");
        $("#ANAL_T5").val("");
        $("#ANAL_T6").val("");
        $("#ANAL_T7").val("");
        $("#ANAL_T8").val("");
        $("#ANAL_T9").val("");*/
        var limite = resp.tipoDimension;
        //1 obligatorio
        //2 opcional
        //3 prohibido (ninguno!)
        //4 todos!
        var prohibido=3;
        if(resp.D0<=limite && limite != prohibido)
        {
          $("#divD0").css("display","block");
        }
        $("#ANAL_T0").attr("status",resp.D0);
        if(resp.D1<=limite && limite != prohibido)
        {
          $("#divD1").css("display","block");
        }
        $("#ANAL_T1").attr("status",resp.D1);
        if(resp.D2<=limite && limite != prohibido)
        {
          $("#divD2").css("display","block");          
        }
        $("#ANAL_T2").attr("status",resp.D2);
        if(resp.D3<=limite && limite != prohibido)
        {
          $("#divD3").css("display","block");
        }
        $("#ANAL_T3").attr("status",resp.D3);
        if(resp.D4<=limite && limite != prohibido)
        {
          $("#divD4").css("display","block");
        }
        $("#ANAL_T4").attr("status",resp.D4);
        if(resp.D5<=limite && limite != prohibido)
        {
          $("#divD5").css("display","block");
        }
        $("#ANAL_T5").attr("status",resp.D5);
        if(resp.D6<=limite && limite != prohibido)
        {
          $("#divD6").css("display","block");
        }
        $("#ANAL_T6").attr("status",resp.D6);
        if(resp.D7<=limite && limite != prohibido)
        {
          $("#divD7").css("display","block");
        }
        $("#ANAL_T7").attr("status",resp.D7);
        if(resp.D8<=limite && limite != prohibido)
        {
          $("#divD8").css("display","block");
        }
        $("#ANAL_T8").attr("status",resp.D8);
        if(resp.D9<=limite && limite != prohibido)
        {
          $("#divD9").css("display","block");
        }
        $("#ANAL_T9").attr("status",resp.D9);

        
      }
      if(resp.success==0)
      {
        //esa cuenta no existe!
      }
    }
  });
}


$(document).on("click", "#nuevaLinea", function (){
  grabaLineaActual(1);
});

$(document).on("click", "#ok", function (){
  grabaLineaActual(0);
});
var idTipoDeDiarioGlobal;


$(document).on("click", "#nuevaReferencia", function (){
  grabaLineaActual(3);
});
$(document).on("click", "#contabilizar", function (){
  $("body").css("cursor","wait");
  grabaLineaActual(2);
});


function grabaLineaActual(queHagoAlFinal)
{
  //primero validamos
  $("#autocomplete-custom-append").css("background-color","#FFFFFF");
  $("#fechaTrans").css("background-color","#FFFFFF");
  $("#descripcion").css("background-color","#FFFFFF");

  $("#ANAL_T0").css("background-color","#FFFFFF");
  $("#ANAL_T1").css("background-color","#FFFFFF");
  $("#ANAL_T2").css("background-color","#FFFFFF");
  $("#ANAL_T3").css("background-color","#FFFFFF");
  $("#ANAL_T4").css("background-color","#FFFFFF");
  $("#ANAL_T5").css("background-color","#FFFFFF");
  $("#ANAL_T6").css("background-color","#FFFFFF");
  $("#ANAL_T7").css("background-color","#FFFFFF");
  $("#ANAL_T8").css("background-color","#FFFFFF");
  $("#ANAL_T9").css("background-color","#FFFFFF");
 
  

  var cuenta = $("#autocomplete-custom-append").val().trim();
  if(cuenta=="")
  {
    $("#autocomplete-custom-append").css("background-color","#FFFC00");
    $("body").css("cursor","default");
    return;
  }
  var fechaTrans = $("#fechaTrans").val().trim();
  if(fechaTrans=="")
  {
    $("#fechaTrans").css("background-color","#FFFC00");
    $("body").css("cursor","default");
    return;
  }
  var descripcion = $("#descripcion").val().trim();
  if(descripcion=="")
  {
    $("#descripcion").css("background-color","#FFFC00");
    $("body").css("cursor","default");
    return;
  }
  var S0 = parseInt($("#ANAL_T0").attr("status"));
  var S1 = parseInt($("#ANAL_T1").attr("status"));
  var S2 = parseInt($("#ANAL_T2").attr("status"));
  var S3 = parseInt($("#ANAL_T3").attr("status"));
  var S4 = parseInt($("#ANAL_T4").attr("status"));
  var S5 = parseInt($("#ANAL_T5").attr("status"));
  var S6 = parseInt($("#ANAL_T6").attr("status"));
  var S7 = parseInt($("#ANAL_T7").attr("status"));
  var S8 = parseInt($("#ANAL_T8").attr("status"));
  var S9 = parseInt($("#ANAL_T9").attr("status"));
  var ANAL_T0 = $("#ANAL_T0").val().trim();
  var ANAL_T1 = $("#ANAL_T1").val().trim();
  var ANAL_T2 = $("#ANAL_T2").val().trim();
  var ANAL_T3 = $("#ANAL_T3").val().trim();
  var ANAL_T4 = $("#ANAL_T4").val().trim();
  var ANAL_T5 = $("#ANAL_T5").val().trim();
  var ANAL_T6 = $("#ANAL_T6").val().trim();
  var ANAL_T7 = $("#ANAL_T7").val().trim();
  var ANAL_T8 = $("#ANAL_T8").val().trim();
  var ANAL_T9 = $("#ANAL_T9").val().trim();
  var obligatorio=1;
  if(S0==obligatorio && ANAL_T0=="")
  {
    $("#ANAL_T0").css("background-color","#FFFC00");
    $("body").css("cursor","default");
    return;
  }
  if(S1==obligatorio && ANAL_T1=="")
  {
    $("#ANAL_T1").css("background-color","#FFFC00");
    $("body").css("cursor","default");
    return;
  }
  if(S2==obligatorio && ANAL_T2=="")
  {
    $("#ANAL_T2").css("background-color","#FFFC00");
    $("body").css("cursor","default");
    return;
  }
  if(S3==obligatorio && ANAL_T3=="")
  {
    $("#ANAL_T3").css("background-color","#FFFC00");
    $("body").css("cursor","default");
    return;
  }
  if(S4==obligatorio && ANAL_T4=="")
  {
    $("#ANAL_T4").css("background-color","#FFFC00");
    $("body").css("cursor","default");
    return;
  }
  if(S5==obligatorio && ANAL_T5=="")
  {
    $("#ANAL_T5").css("background-color","#FFFC00");
    $("body").css("cursor","default");
    return;
  }
  if(S6==obligatorio && ANAL_T6=="")
  {
    $("#ANAL_T6").css("background-color","#FFFC00");
    $("body").css("cursor","default");
    return;
  }
  if(S7==obligatorio && ANAL_T7=="")
  {
    $("#ANAL_T7").css("background-color","#FFFC00");
    $("body").css("cursor","default");
    return;
  }
  if(S8==obligatorio && ANAL_T8=="")
  {
    $("#ANAL_T8").css("background-color","#FFFC00");
    $("body").css("cursor","default");
    return;
  }
  if(S9==obligatorio && ANAL_T9=="")
  {
    $("#ANAL_T9").css("background-color","#FFFC00");
    $("body").css("cursor","default");
    return;
  }
  if(ANAL_T0.indexOf('|')>-1)
  {
    ANAL_T0 = ANAL_T0.split('|')[0];
  }
  if(ANAL_T1.indexOf('|')>-1)
  {
    ANAL_T1 = ANAL_T1.split('|')[0];
  }
  if(ANAL_T2.indexOf('|')>-1)
  {
    ANAL_T2 = ANAL_T2.split('|')[0];
  }
  if(ANAL_T3.indexOf('|')>-1)
  {
    ANAL_T3 = ANAL_T3.split('|')[0];
  }
  if(ANAL_T4.indexOf('|')>-1)
  {
    ANAL_T4 = ANAL_T4.split('|')[0];
  }
  if(ANAL_T5.indexOf('|')>-1)
  {
    ANAL_T5 = ANAL_T5.split('|')[0];
  }
  if(ANAL_T6.indexOf('|')>-1)
  {
    ANAL_T6 = ANAL_T6.split('|')[0];
  }
  if(ANAL_T7.indexOf('|')>-1)
  {
    ANAL_T7 = ANAL_T7.split('|')[0];
  }
  if(ANAL_T8.indexOf('|')>-1)
  {
    ANAL_T8 = ANAL_T8.split('|')[0];
  }
  if(ANAL_T9.indexOf('|')>-1)
  {
    ANAL_T9 = ANAL_T9.split('|')[0];
  }

  var cantidad = Math.abs(parseFloat($("#cantidad").val().trim().replace("$","")));
  var D_C = $("#tipoDC").val();
  var param = {
    timestamp : timestampActual,
    lineaG : lineaGlobal,
    lineaT : lineaEnLaQueVoy,
    ref : referenciaActual,
    ACNT_CODE : cuenta,
    fecha : fechaTrans,
    descripcion : descripcion,
    D_C : D_C,
    AMOUNT : cantidad,
    ANAL_T0 : ANAL_T0,
    ANAL_T1 : ANAL_T1,
    ANAL_T2 : ANAL_T2,
    ANAL_T3 : ANAL_T3,
    ANAL_T4 : ANAL_T4,
    ANAL_T5 : ANAL_T5,
    ANAL_T6 : ANAL_T6,
    ANAL_T7 : ANAL_T7,
    ANAL_T8 : ANAL_T8,
    ANAL_T9 : ANAL_T9,
    idTipoDeDiario : idTipoDeDiarioGlobal
  };
  $.ajax({
    url: "/guardaLineaEnTemporalesDiario",
    dataType : "json",
    type : "post",
    data: param,
    async : true,
    beforeSend : function (){
    },
    complete : function (){
    }, 
    success : function (resp){
      if(resp.success==1)
      {
        debitosReferencia=resp.debitos;
        creditosReferencia=resp.creditos;
        if(queHagoAlFinal==3)//nueva referencia
        {
          lineaGlobal++;
          lineaEnLaQueVoy=1;
          referenciaActual="";
          debitosReferencia=0;
          creditosReferencia=0;
          cuentaAnterior="";
          $("#divDescr").css("display","none");
          $("#divCantidad").css("display","none");
          $("#cantidad").val("");
          $("#divD_C").css("display","none");
          $("#divCuenta").css("display","none");
          $("#divFecha").css("display","none");
          $("#divD0").css("display","none");
          $("#divD1").css("display","none");
          $("#divD2").css("display","none");
          $("#divD3").css("display","none");
          $("#divD4").css("display","none");
          $("#divD5").css("display","none");
          $("#divD6").css("display","none");
          $("#divD7").css("display","none");
          $("#divD8").css("display","none");
          $("#divD9").css("display","none");
       
          //ponLineaSegunTipoDeDiario(idTipoDeDiarioGlobal, lineaEnLaQueVoy);
          $("#referencia").removeAttr("disabled");
          $("#referencia").select();
        }
        if(queHagoAlFinal==1)
        {
          lineaGlobal++;
          lineaEnLaQueVoy++;
          ponLineaSegunTipoDeDiario(idTipoDeDiarioGlobal, lineaEnLaQueVoy);
        }
        $("#infoLabel").html("D&eacute;bitos: "+format2(debitosReferencia,"$")+" Cr&eacute;ditos: "+format2(creditosReferencia,"$")+"  Linea: "+lineaEnLaQueVoy);
        if(queHagoAlFinal==2)//coontabiliza
        {
          var param = {
            timestamp : timestampActual,
            idTipoDeDiario : idTipoDeDiarioGlobal
          };
          $.ajax({
            url: "/contabilizaPorTimeStamp",
            dataType : "json",
            type : "post",
            data: param,
            async : true,
            beforeSend : function (){
            },
            complete : function (){
            }, 
            success : function (resp){
              $("body").css("cursor","default");
              if(resp.success==1)
              {   
                window.open(
                  "/generarDiario?diario=" + resp.diario,
                  '_blank'
                );
                empiezaCapturaDeDiarios();
              }
              if(resp.success==2)
              {   
                alert("Los debitos suman: "+resp.debitos+" , los creditos suman: "+resp.creditos);
              }
              if(resp.success==0)
              {   
                alert("Hubo un error! perdon");
              }
            }
          });
        }
      }
    }
  });
}
function actualizaActivosFijos()
{
  $.ajax({
    url: "/dameActivosFijos",
    dataType : "json",
    type : "post",
    async : true,
    beforeSend : function (){
    },
    complete : function (){
    }, 
    success : function (resp){
      if(resp.success==1)
      {
        var i;
        var cad='';
        var aver2 = Object.keys(resp.activos);
        for(i=0;i<aver2.length;i++)
        {
          cad=cad+'<tr><td>'+resp.activos[aver2[i]]["ASSET_CODE"]+'</td><td>'+resp.activos[aver2[i]]["DESCR"]+'</td><td>'+resp.activos[aver2[i]]["STATUS"]+'</td><td>'+resp.activos[aver2[i]]["ASSET_STATUS"]+'</td><td>'+resp.activos[aver2[i]]["START_PERIOD"]+'</td><td>'+resp.activos[aver2[i]]["END_PERIOD"]+'</td><td>'+resp.activos[aver2[i]]["ULTIMO_PERIOD"]+'</td><td>'+resp.activos[aver2[i]]["DISPOSAL_PERIOD"]+'</td><td>'+resp.activos[aver2[i]]["DISPOSED"]+'</td><td>'+format2(resp.activos[aver2[i]]["BASE_GROSS"],"$")+'</td><td>'+format2(resp.activos[aver2[i]]["BASE_DEP"],"$")+'</td><td>'+format2(resp.activos[aver2[i]]["BASE_NET"],"$")+'</td><td>'+resp.activos[aver2[i]]["BASE_PCENT"]+'</td><td>'+format2(resp.activos[aver2[i]]["TXN_GROSS"],"$")+'</td><td>'+format2(resp.activos[aver2[i]]["TXN_DEP"],"$")+'</td><td>'+format2(resp.activos[aver2[i]]["TXN_NET"],"$")+'</td><td>'+resp.activos[aver2[i]]["TXN_PCENT"]+'</td></tr>';
        } 
        $("#tablaActivosFijosDepre").html(cad);
      }
    }
  });
}
function ponLineaSegunTipoDeDiario(idTipoDeDiario, linea)
{
  var param = {
    idTipoDeDiario : idTipoDeDiario,
    linea : linea
  };
  $.ajax({
    url: "/dameLineaDelTipoDeDiario",
    dataType : "json",
    type : "post",
    data: param,
    async : true,
    beforeSend : function (){
    },
    complete : function (){
    }, 
    success : function (resp){
      if(resp.success==1)
      {
        $("#divDescr").css("display","block");
        if(resp.DESCRIPTN!="..")
        {
          $("#descripcion").val(resp.DESCRIPTN);  
        }
        

        $("#divCantidad").css("display","block");


        
        $("#divD_C").css("display","block");
        $("#tipoDC").val(resp.D_C);

        $("#divCuenta").css("display","block");
        $("#autocomplete-custom-append").val(resp.cuenta);
        ponDimensionesSegunCuenta();
        if(resp.D0!="..")
        {
          $("#ANAL_T0").val(resp.D0);
        }
        if(resp.D1!="..")
        {
          $("#ANAL_T1").val(resp.D1);
        }
        if(resp.D2!="..")
        {
          $("#ANAL_T2").val(resp.D2);
        }
        if(resp.D3!="..")
        {
          $("#ANAL_T3").val(resp.D3);
        }
        if(resp.D4!="..")
        {
          $("#ANAL_T4").val(resp.D4);
        }
        if(resp.D5!="..")
        {
          $("#ANAL_T5").val(resp.D5);
        }
        if(resp.D6!="..")
        {
          $("#ANAL_T6").val(resp.D6);
        }
        if(resp.D7!="..")
        {
          $("#ANAL_T7").val(resp.D7);
        }
        if(resp.D8!="..")
        {
          $("#ANAL_T8").val(resp.D8);
        }
        if(resp.D9!="..")
        {
          $("#ANAL_T9").val(resp.D9);
        }
        $("#divFecha").css("display","block");
        if(linea==1)
        {
          $("#fechaTrans").click();
        }
      }
    }
  });
}
var timestampActual;
function empiezaCapturaDeDiarios()
{
 
  lineaEnLaQueVoy=1;
  referenciaActual="";
  debitosReferencia=0;
  creditosReferencia=0;
  cuentaAnterior="";
  lineaGlobal = 1;
idTipoDeDiarioGlobal=-1;
    var cad = '<div class="">'+
          '<div class="page-title">'+
            '<div class="title_left">'+
              '<h3>Capturar Diario</h3>'+
            '</div>'+
          '</div>'+
          '<div class="clearfix"></div>'+
          '<div class="row">'+
            '<div class="col-md-12 col-sm-12 col-xs-12">'+
              '<div class="x_panel">'+
                '<div class="x_title">'+
                  '<h2>Captura de diario<small></small></h2>'+
                  '<ul class="nav navbar-right panel_toolbox">'+
                    '<li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>'+
                    '</li>'+
                    '<li><a class="close-link"><i class="fa fa-close"></i></a>'+
                    '</li>'+
                  '</ul>'+
                  '<div class="clearfix"></div>'+
                '</div>'+
                '<div class="x_content">'+
                  '<br />'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Elige tipo de diario:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<select style="width: 100%;font-size: 20px;" id="tipoDeDiarioSelectDeCaptura"></select>'+
                       '</div>'+
                    '</div>'+
                    '<br><br>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Referencia:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="referencia" class="form-control col-md-10" style="float: left;" />'+
                        '</div>'+
                    '</div>'+
                    '<br><br>'+
                    '<div class="form-group">'+
                      '<div class="col-md-9 col-sm-9 col-xs-9 col-md-offset-6">'+
                        '<button type="button" id="ok" class="btn btn-success">OK</button>'+
                        '<button type="button" id="nuevaLinea" class="btn btn-success">Nueva linea</button>'+
                        '<button type="button" id="nuevaReferencia" class="btn btn-success">Nueva referencia</button>'+
                        '<button type="button" id="contabilizar" class="btn btn-success">Contabilizar</button>'+
                      '</div>'+
                    '</div>'+
                    '<br><br>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-9 col-sm-9 col-xs-12" id="infoLabel"></label>'+
                    '</div>'+
                    '<br><br>'+ 
                    '<div class="form-group" style="display:none;" id="divFecha">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Fecha de la transacci&oacute;n:</label>'+
                      '<div class="col-md-3">'+
                        '<fieldset>'+
                          '<div class="control-group">'+
                            '<div class="controls">'+
                              '<div class="col-md-11 xdisplay_inputx form-group has-feedback">'+
                                '<input type="text" class="form-control has-feedback-left" id="fechaTrans" placeholder="Fecha de transaccion" aria-describedby="inputSuccess2Status">'+
                                '<span class="fa fa-calendar-o form-control-feedback left" id="calendarioFecha" aria-hidden="true"></span>'+
                                '<span id="inputSuccess2Status" class="sr-only">(success)</span>'+
                              '</div>'+
                            '</div>'+
                          '</div>'+
                        '</fieldset>'+
                      '</div>'+
                    '</div>'+
                    '<br><br>'+
                    '<div class="form-group" style="display:none;" id="divCuenta">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Cuenta:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="autocomplete-custom-append" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="autocomplete-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<br><br>'+
                     '<div class="form-group" style="display:none;" id="divDescr">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Descripci&oacute;n:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="descripcion" class="form-control col-md-10" style="float: left;" />'+
                        '</div>'+
                    '</div>'+
                    '<br><br>'+
                     '<div class="form-group" style="display:none;" id="divCantidad">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Cantidad:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="cantidad" class="form-control col-md-10" style="float: left;" />'+
                        '</div>'+
                    '</div>'+
                    '<br><br>'+
                    '<div class="form-group" style="display:none;" id="divD_C">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Elige tipo de movimiento:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<select style="width: 100%;font-size: 20px;" id="tipoDC"><option value="D">D&eacute;bito</option><option value="C">Cr&eacute;dito</option></select>'+
                       '</div>'+
                    '</div>'+
                    '<br><br>'+
                    '<div class="form-group" style="display:none;" id="divD0">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T0_LABEL">D0:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T0" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T0-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group" style="display:none;" id="divD1">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T1_LABEL">D1:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T1" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T1-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group"  style="display:none;" id="divD2">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T2_LABEL">D2:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                       '<input type="text" id="ANAL_T2" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T2-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group"  style="display:none;" id="divD3">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T3_LABEL">D3:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T3" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T3-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group"  style="display:none;" id="divD4">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T4_LABEL">D4:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T4" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T4-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group"  style="display:none;" id="divD5">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T5_LABEL">D5:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T5" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T5-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group"  style="display:none;" id="divD6">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T6_LABEL">D6:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T6" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T6-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group"  style="display:none;" id="divD7">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T7_LABEL">D7:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T7" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T7-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group"  style="display:none;" id="divD8">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T8_LABEL">D8:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T8" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T8-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group"  style="display:none;" id="divD9">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T9_LABEL">D9:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T9" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T9-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                '</div>'+
                '<div class="clearfix"></div>'+
                '<div class="col-md-12 col-sm-12 col-xs-12">'+
                  '<div class="x_panel">'+
                    '<div class="x_title">'+
                      '<h2>Lineas <small>del diario</small></h2>'+
                      '<div class="clearfix"></div>'+
                    '</div>'+
                    '<div class="x_content">'+
                      '<table class="table table-hover">'+
                        '<thead><tr><th>#</th></th><th>Cuenta</th><th>Tipo</th><th>Descripci&oacute;n</th><th id="tdD0">D0</th><th id="tdD1">D1</th><th id="tdD2">D2</th><th id="tdD3">D3</th><th id="tdD4">D4</th><th id="tdD5">D5</th><th id="tdD6">D6</th><th id="tdD7">D7</th><th id="tdD8">D8</th><th id="tdD9">D9</th><th>Facturar</th><th>Cliente</th><th>Servicio</th><th>Concepto</th><th>Editar</th></thead>'+
                        '<tbody id="tablaLineasDelDiario"></tbody>'+
                      '</table>'+
                    '</div>'+
                  '</div>'+
                '</div>'+
              '</div>'+
            '</div>'+
          '</div>'+
        '</div>';     
      $("#mainContent").html(cad);

      $("#cantidad").maskMoney({thousands:'', decimal:'.', allowZero:true, prefix: '$ '});
      $('#fechaTrans').daterangepicker({
          singleDatePicker: true,
          calender_style: "picker_1"
        }, function(start, end, label) {
          console.log(start.toISOString(), end.toISOString(), label);
        });
      $("#referencia").focusout(function(){
          empiezaReferencia();
      });

      $("#autocomplete-custom-append").focusout(function(){
          ponDimensionesSegunCuenta();
      });

     $("#autocomplete-custom-append").keydown(function(ev){
       if(ev.which === 13) {
        ponDimensionesSegunCuenta();
        return false;
        }
      });

      configInicialDeLineasDeCapturaDeDiario();
}
$(document).on("click", "#capturaDeDiarios", function (){
  empiezaCapturaDeDiarios();
 });
$(document).on("click", "#lineasTiposDeDiario", function (){
    var cad = '<div class="">'+
          '<div class="page-title">'+
            '<div class="title_left">'+
              '<h3>Crear lineas de Tipos de Diario</h3>'+
            '</div>'+
          '</div>'+
          '<div class="clearfix"></div>'+
          '<div class="row">'+
            '<div class="col-md-12 col-sm-12 col-xs-12">'+
              '<div class="x_panel">'+
                '<div class="x_title">'+
                  '<h2>Crear lineas de tipos de diario<small></small></h2>'+
                  '<ul class="nav navbar-right panel_toolbox">'+
                    '<li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>'+
                    '</li>'+
                    '<li><a class="close-link"><i class="fa fa-close"></i></a>'+
                    '</li>'+
                  '</ul>'+
                  '<div class="clearfix"></div>'+
                '</div>'+
                '<div class="x_content">'+
                  '<br />'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Elige tipo de diario:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<select style="width: 100%;font-size: 20px;" id="tipoDeDiarioSelect"></select>'+
                       '</div>'+
                    '</div>'+
                    '<br><br>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Linea:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="lineaDelTipoDeDiario" class="form-control col-md-10" style="float: left;" />'+
                        '</div>'+
                    '</div>'+
                    '<br><br>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Debo facturar</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<div id="gender" class="btn-group" data-toggle="buttons">'+
                          '<label class="btn btn-default" data-toggle-class="btn-primary" data-toggle-passive-class="btn-default">'+
                            '<input type="radio" name="gender" value="1" data-parsley-multiple="gender" data-parsley-id="12"> &nbsp; Si &nbsp;'+
                          '</label>'+
                          '<label class="btn btn-default active" data-toggle-class="btn-primary" data-toggle-passive-class="btn-default">'+
                            '<input type="radio" name="gender" value="0" data-parsley-multiple="gender"> No'+
                          '</label>'+
                        '</div>'+
                      '</div>'+
                    '</div>'+
                    '<br><br>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Servicio:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="servicio" class="form-control col-md-10" style="float: left;" />'+
                        '</div>'+
                    '</div>'+
                    '<br><br>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Cliente:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="cliente" class="form-control col-md-10" style="float: left;" />'+
                        '</div>'+
                    '</div>'+
                    '<br><br>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Concepto:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="concepto" class="form-control col-md-10" style="float: left;" />'+
                        '</div>'+
                    '</div>'+
                    '<br><br>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Cuenta:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="autocomplete-custom-append" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="autocomplete-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<br><br>'+
                     '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Descripci&oacute;n:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="descripcion" class="form-control col-md-10" style="float: left;" />'+
                        '</div>'+
                    '</div>'+
                    '<br><br>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Elige tipo de movimiento:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<select style="width: 100%;font-size: 20px;" id="tipoDC"><option value="D">D&eacute;bito</option><option value="C">Cr&eacute;dito</option></select>'+
                       '</div>'+
                    '</div>'+
                    '<br><br>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T0_LABEL">D0:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T0" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T0-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T1_LABEL">D1:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T1" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T1-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T2_LABEL">D2:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                       '<input type="text" id="ANAL_T2" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T2-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T3_LABEL">D3:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T3" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T3-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T4_LABEL">D4:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T4" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T4-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T5_LABEL">D5:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T5" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T5-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T6_LABEL">D6:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T6" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T6-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T7_LABEL">D7:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T7" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T7-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T8_LABEL">D8:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T8" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T8-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T9_LABEL">D9:</label>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T9" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T9-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<div class="col-md-6 col-sm-6 col-xs-6 col-md-offset-6">'+
                        '<button type="submit" id="guardarNuevaLineaDelTipoDeDiario" idTipoDeDiario="-1" class="btn btn-success">Guardar nueva linea</button>'+
                      '</div>'+
                    '</div>'+
                '</div>'+
                '<div class="clearfix"></div>'+
                '<div class="col-md-12 col-sm-12 col-xs-12">'+
                  '<div class="x_panel">'+
                    '<div class="x_title">'+
                      '<h2>Crear <small>Lineas de tipo de diario</small></h2>'+
                      '<div class="clearfix"></div>'+
                    '</div>'+
                    '<div class="x_content">'+
                      '<table class="table table-hover">'+
                        '<thead><tr><th>#</th></th><th>Cuenta</th><th>Tipo</th><th>Descripci&oacute;n</th><th id="tdD0">D0</th><th id="tdD1">D1</th><th id="tdD2">D2</th><th id="tdD3">D3</th><th id="tdD4">D4</th><th id="tdD5">D5</th><th id="tdD6">D6</th><th id="tdD7">D7</th><th id="tdD8">D8</th><th id="tdD9">D9</th><th>Facturar</th><th>Cliente</th><th>Servicio</th><th>Concepto</th><th>Editar</th></thead>'+
                        '<tbody id="tablaLineaTiposDeDiario"></tbody>'+
                      '</table>'+
                    '</div>'+
                  '</div>'+
                '</div>'+
              '</div>'+
            '</div>'+
          '</div>'+
        '</div>';     
      $("#mainContent").html(cad);
      //configInicialDeLineasDeCedulas();
      configInicialDeLineasDeTiposDeDiario();
 });

$(document).on("click", "#lineasDeCedula", function (){
    var cad = '<div class="">'+
          '<div class="page-title">'+
            '<div class="title_left">'+
              '<h3>Crear lineas de c&eacute;dula</h3>'+
            '</div>'+
          '</div>'+
          '<div class="clearfix"></div>'+
          '<div class="row">'+
            '<div class="col-md-12 col-sm-12 col-xs-12">'+
              '<div class="x_panel">'+
                '<div class="x_title">'+
                  '<h2>Crear lineas de c&eacute;dulas<small></small></h2>'+
                  '<ul class="nav navbar-right panel_toolbox">'+
                    '<li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>'+
                    '</li>'+
                    '<li><a class="close-link"><i class="fa fa-close"></i></a>'+
                    '</li>'+
                  '</ul>'+
                  '<div class="clearfix"></div>'+
                '</div>'+
                '<div class="x_content">'+
                  '<br />'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Elige cedula:</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<select style="width: 100%;font-size: 20px;" id="cedulasSelect"></select>'+
                       '</div>'+
                    '</div>'+
                     '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Elige Concepto:</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<select style="width: 100%;font-size: 20px;" id="conceptoSelect"></select>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Cuenta:</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="autocomplete-custom-append" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="autocomplete-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Elige tipo de movimientos:</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<select style="width: 100%;font-size: 20px;" id="tipoDC"><option value="1">D&eacute;bitos</option><option value="2">Cr&eacute;ditos</option><option value="4">Ambos</option></select>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T0_LABEL">D0:</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T0" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T0-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T1_LABEL">D1:</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T1" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T1-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T2_LABEL">D2:</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                       '<input type="text" id="ANAL_T2" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T2-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T3_LABEL">D3:</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T3" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T3-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T4_LABEL">D4:</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T4" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T4-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T5_LABEL">D5:</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T5" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T5-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T6_LABEL">D6:</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T6" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T6-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T7_LABEL">D7:</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T7" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T7-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T8_LABEL">D8:</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T8" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T8-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12" id="ANAL_T9_LABEL">D9:</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" id="ANAL_T9" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="ANAL_T9-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<div class="col-md-6 col-sm-6 col-xs-6 col-md-offset-6">'+
                        '<button type="submit" id="guardarNuevaLineaCedula" idCedula="-1" class="btn btn-success">Guardar nueva linea</button>'+
                      '</div>'+
                    '</div>'+
                '</div>'+
                '<div class="clearfix"></div>'+
                '<div class="col-md-12 col-sm-12 col-xs-12">'+
                  '<div class="x_panel">'+
                    '<div class="x_title">'+
                      '<h2>Crear <small>C&eacute;dulas</small></h2>'+
                      '<div class="clearfix"></div>'+
                    '</div>'+
                    '<div class="x_content">'+
                      '<table class="table table-hover">'+
                        '<thead><tr><th>#</th></th><th>Cuenta</th><th>Tipo</th><th id="tdD0">D0</th><th id="tdD1">D1</th><th id="tdD2">D2</th><th id="tdD3">D3</th><th id="tdD4">D4</th><th id="tdD5">D5</th><th id="tdD6">D6</th><th id="tdD7">D7</th><th id="tdD8">D8</th><th id="tdD9">D9</th><th>Editar</th></thead>'+
                        '<tbody id="tablaLineaCedulas"></tbody>'+
                      '</table>'+
                    '</div>'+
                  '</div>'+
                '</div>'+
              '</div>'+
            '</div>'+
          '</div>'+
        '</div>';     
      $("#mainContent").html(cad);
      configInicialDeLineasDeCedulas();
 });


$(document).on("click", "#tiposDeDiario", function (){
    var cad = '<div class="">'+
          '<div class="page-title">'+
            '<div class="title_left">'+
              '<h3>Crear Tipos de diario</h3>'+
            '</div>'+
          '</div>'+
          '<div class="clearfix"></div>'+
          '<div class="row">'+
            '<div class="col-md-12 col-sm-12 col-xs-12">'+
              '<div class="x_panel">'+
                '<div class="x_title">'+
                  '<h2>Crear Tipos de diario<small></small></h2>'+
                  '<ul class="nav navbar-right panel_toolbox">'+
                    '<li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>'+
                    '</li>'+
                    '<li><a class="close-link"><i class="fa fa-close"></i></a>'+
                    '</li>'+
                  '</ul>'+
                  '<div class="clearfix"></div>'+
                '</div>'+
                '<div class="x_content">'+
                  '<br />'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Nombre:</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" name="nombre" id="nombre" class="form-control col-md-10" style="float: left;" />'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">C&oacute;digo:</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" name="codigo" id="codigo" class="form-control col-md-10" style="float: left;" />'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<div class="col-md-6 col-sm-6 col-xs-6 col-md-offset-6">'+
                        '<button type="submit" id="guardarNuevoTipoDeDiario" idTipoDeDiario="-1" class="btn btn-success">Guardar nuevo</button>'+
                      '</div>'+
                    '</div>'+
                '</div>'+
                '<div class="clearfix"></div>'+
                '<div class="col-md-12 col-sm-12 col-xs-12">'+
                  '<div class="x_panel">'+
                    '<div class="x_title">'+
                      '<h2>Crear <small>Tipos de Diario</small></h2>'+
                      '<div class="clearfix"></div>'+
                    '</div>'+
                    '<div class="x_content">'+
                      '<table class="table table-hover">'+
                        '<thead><tr><th>#</th><th>Nombre</th><th>C&oacute;digo</th><th>Editar</th></thead>'+
                        '<tbody id="tablaTiposDeDiario"></tbody>'+
                      '</table>'+
                    '</div>'+
                  '</div>'+
                '</div>'+
              '</div>'+
            '</div>'+
          '</div>'+
        '</div>';     
      $("#mainContent").html(cad);
      actualizaTiposDeDiario();
 });
$(document).on("click", "#crearCedula", function (){
    var cad = '<div class="">'+
          '<div class="page-title">'+
            '<div class="title_left">'+
              '<h3>Crear c&eacute;dula</h3>'+
            '</div>'+
          '</div>'+
          '<div class="clearfix"></div>'+
          '<div class="row">'+
            '<div class="col-md-12 col-sm-12 col-xs-12">'+
              '<div class="x_panel">'+
                '<div class="x_title">'+
                  '<h2>Crear c&eacute;dulas<small></small></h2>'+
                  '<ul class="nav navbar-right panel_toolbox">'+
                    '<li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>'+
                    '</li>'+
                    '<li><a class="close-link"><i class="fa fa-close"></i></a>'+
                    '</li>'+
                  '</ul>'+
                  '<div class="clearfix"></div>'+
                '</div>'+
                '<div class="x_content">'+
                  '<br />'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Nombre:</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" name="nombre" id="nombre" class="form-control col-md-10" style="float: left;" />'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<div class="col-md-6 col-sm-6 col-xs-6 col-md-offset-6">'+
                        '<button type="submit" id="guardarNuevaCedula" idCedula="-1" class="btn btn-success">Guardar nueva</button>'+
                      '</div>'+
                    '</div>'+
                '</div>'+
                '<div class="clearfix"></div>'+
                '<div class="col-md-12 col-sm-12 col-xs-12">'+
                  '<div class="x_panel">'+
                    '<div class="x_title">'+
                      '<h2>Crear <small>C&eacute;dulas</small></h2>'+
                      '<div class="clearfix"></div>'+
                    '</div>'+
                    '<div class="x_content">'+
                      '<table class="table table-hover">'+
                        '<thead><tr><th>#</th><th>Nombre</th><th>Editar</th></thead>'+
                        '<tbody id="tablaCedulas"></tbody>'+
                      '</table>'+
                    '</div>'+
                  '</div>'+
                '</div>'+
              '</div>'+
            '</div>'+
          '</div>'+
        '</div>';     
      $("#mainContent").html(cad);
      actualizaCedulas();
 });
   $(document).on("click", "#tipoDeDimensiones", function (){
    var cad = '<div class="">'+
          '<div class="page-title">'+
            '<div class="title_left">'+
              '<h3>Tipo de dimensiones</h3>'+
            '</div>'+
          '</div>'+
          '<div class="clearfix"></div>'+
          '<div class="row">'+
            '<div class="col-md-12 col-sm-12 col-xs-12">'+
              '<div class="x_panel">'+
                '<div class="x_title">'+
                  '<h2>Tipo de dimensiones <small></small></h2>'+
                  '<ul class="nav navbar-right panel_toolbox">'+
                    '<li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>'+
                    '</li>'+
                    '<li><a class="close-link"><i class="fa fa-close"></i></a>'+
                    '</li>'+
                  '</ul>'+
                  '<div class="clearfix"></div>'+
                '</div>'+
                '<div class="x_content">'+
                  '<br />'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">C&oacute;digo:</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" name="codigo" id="codigo" class="form-control col-md-10" style="float: left;" />'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Descripci&oacute;n:</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" name="descr" id="descr" class="form-control col-md-10" style="float: left;" />'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<div class="col-md-6 col-sm-6 col-xs-6 col-md-offset-6">'+
                        '<button type="submit" id="guardarNuevoTipoDeDimension" ANL_CAT_ID="-1" class="btn btn-success">Guardar nuevo o modificacion</button>'+
                      '</div>'+
                      '<div class="col-md-6 col-sm-6 col-xs-6 col-md-offset-6">'+
                        '<button type="submit" id="limpiar" class="btn btn-warning">Limpiar</button>'+
                      '</div>'+
                    '</div>'+
                '</div>'+
                '<div class="clearfix"></div>'+
                '<div class="col-md-12 col-sm-12 col-xs-12">'+
                  '<div class="x_panel">'+
                    '<div class="x_title">'+
                      '<h2>Dimensiones <small>Tipos de dimensiones</small></h2>'+
                      '<div class="clearfix"></div>'+
                    '</div>'+
                    '<div class="x_content">'+
                      '<table class="table table-hover">'+
                        '<thead><tr><th>#</th><th>Nombre</th><th>Descripci&oacute;n</th><th>Editar</th></thead>'+
                        '<tbody id="tablaDimensiones"></tbody>'+
                      '</table>'+
                    '</div>'+
                  '</div>'+
                '</div>'+
              '</div>'+
            '</div>'+
          '</div>'+
        '</div>';     
      $("#mainContent").html(cad);
      actualizaTipoDeDimensiones();
 });
   $(document).on("click", "#diarioReporte", function (){
    var cad = '<div class="">'+
          '<div class="page-title">'+
            '<div class="title_left">'+
              '<h3>Diario en excel</h3>'+
            '</div>'+
          '</div>'+
          '<div class="clearfix"></div>'+
          '<div class="row">'+
            '<div class="col-md-12 col-sm-12 col-xs-12">'+
              '<div class="x_panel">'+
                '<div class="x_title">'+
                  '<h2>Diario en excel <small></small></h2>'+
                  '<ul class="nav navbar-right panel_toolbox">'+
                    '<li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>'+
                    '</li>'+
                    '<li><a class="close-link"><i class="fa fa-close"></i></a>'+
                    '</li>'+
                  '</ul>'+
                  '<div class="clearfix"></div>'+
                '</div>'+
                '<div class="x_content">'+
                  '<br />'+
                     '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Diario</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" name="country" id="diario" class="form-control col-md-10" style="float: left;" />'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<div class="col-md-6 col-sm-6 col-xs-12 col-md-offset-3">'+
                        '<button type="submit" id="generarDiarioEnPDF" class="btn btn-success">Generar</button>'+
                      '</div>'+
                    '</div>'+
                '</div>'+
              '</div>'+
            '</div>'+
          '</div>'+
        '</div>';     
      $("#mainContent").html(cad);
 });
   $(document).on("click", "#diarioEnExcel", function (){
    var cad = '<div class="">'+
          '<div class="page-title">'+
            '<div class="title_left">'+
              '<h3>Diario en excel</h3>'+
            '</div>'+
          '</div>'+
          '<div class="clearfix"></div>'+
          '<div class="row">'+
            '<div class="col-md-12 col-sm-12 col-xs-12">'+
              '<div class="x_panel">'+
                '<div class="x_title">'+
                  '<h2>Diario en excel <small></small></h2>'+
                  '<ul class="nav navbar-right panel_toolbox">'+
                    '<li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>'+
                    '</li>'+
                    '<li><a class="close-link"><i class="fa fa-close"></i></a>'+
                    '</li>'+
                  '</ul>'+
                  '<div class="clearfix"></div>'+
                '</div>'+
                '<div class="x_content">'+
                  '<br />'+
                     '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Diario</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" name="country" id="diario" class="form-control col-md-10" style="float: left;" />'+
                       '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<div class="col-md-6 col-sm-6 col-xs-12 col-md-offset-3">'+
                        '<button type="submit" id="generarDiarioEnExcel" class="btn btn-success">Generar</button>'+
                      '</div>'+
                    '</div>'+
                '</div>'+
              '</div>'+
            '</div>'+
          '</div>'+
        '</div>';     
      $("#mainContent").html(cad);
 });
   


$(document).on("click", "#guardarClasificacionDimensiones", function (){
  var TipoDondeEstaranLasDimensiones = parseInt($("#TipoDondeEstaranLasDimensiones").val());
  if(TipoDondeEstaranLasDimensiones==0)
  {
    alert("Tienes que seleccionar un tipo antes de guardar!");
    return;
  }
  var S_HEAD1 = $("#entry1").attr("S_HEAD");
  var S_HEAD2 = $("#entry2").attr("S_HEAD");
  var S_HEAD3 = $("#entry3").attr("S_HEAD");
  var S_HEAD4 = $("#entry4").attr("S_HEAD");
  var S_HEAD5 = $("#entry5").attr("S_HEAD");
  var S_HEAD6 = $("#entry6").attr("S_HEAD");
  var S_HEAD7 = $("#entry7").attr("S_HEAD");
  var S_HEAD8 = $("#entry8").attr("S_HEAD");
  var S_HEAD9 = $("#entry9").attr("S_HEAD");
  var S_HEAD10 = $("#entry10").attr("S_HEAD");
  var entry1 = $("#entry1").val();
  var entry2 = $("#entry2").val();
  var entry3 = $("#entry3").val();
  var entry4 = $("#entry4").val();
  var entry5 = $("#entry5").val();
  var entry6 = $("#entry6").val();
  var entry7 = $("#entry7").val();
  var entry8 = $("#entry8").val();
  var entry9 = $("#entry9").val();
  var entry10 = $("#entry10").val();  
  var param = {TipoDondeEstaranLasDimensiones : TipoDondeEstaranLasDimensiones,
    entry1 : entry1,
    entry2 : entry2,
    entry3 : entry3,
    entry4 : entry4,
    entry5 : entry5,
    entry6 : entry6,
    entry7 : entry7,
    entry8 : entry8,
    entry9 : entry9,
    entry10 : entry10,
    S_HEAD1 : S_HEAD1,
    S_HEAD2 : S_HEAD2,
    S_HEAD3 : S_HEAD3,
    S_HEAD4 : S_HEAD4,
    S_HEAD5 : S_HEAD5,
    S_HEAD6 : S_HEAD6,
    S_HEAD7 : S_HEAD7,
    S_HEAD8 : S_HEAD8,
    S_HEAD9 : S_HEAD9,
    S_HEAD10 : S_HEAD10
  };
  $.ajax({
    url: "/guardarClasificacionDimensiones",
    dataType : "json",
    type : "post",
    data: param,
    async : true,
    beforeSend : function (){
//      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
    },
    complete : function (){
    }, 
    success : function (resp){
      if(resp.success==1)
      {
        alert("Los cambios en la asignacion de las dimensiones se hicieron correctamente!");
      }
       if(resp.success==2)
      {
        alert("Hubi un error!");
      }
    }
  });  
});

var contains = function(needle) {
    // Per spec, the way to identify NaN is that it is not equal to itself
    var findNaN = needle !== needle;
    var indexOf;

    if(!findNaN && typeof Array.prototype.indexOf === 'function') {
        indexOf = Array.prototype.indexOf;
    } else {
        indexOf = function(needle) {
            var i = -1, index = -1;

            for(i = 0; i < this.length; i++) {
                var item = this[i];

                if((findNaN && item !== item) || item === needle) {
                    index = i;
                    break;
                }
            }

            return index;
        };
    }

    return indexOf.call(this, needle) > -1;
};


$(document).on("click", "#generarReporteDeMAT", function (){
  var delPeriodo = $("#delPeriodo").val();
  var alPeriodo = $("#alPeriodo").val();

  var param = {delPeriodo : delPeriodo, alPeriodo : alPeriodo}
  $.ajax({
    url: "/generarReporteDeMAT",
    dataType : "json",
    type : "post",
    data: param,
    async : true,
    beforeSend : function (){
    },
    complete : function (){
    }, 
    success : function (resp){
      if(resp.success==1)
      {
        Chart.defaults.global.legend = {
          enabled: true
        };
         Chart.defaults.global.title = {
          enabled: true
        };
        var iglesias = Object.keys(resp.diezmos);
        var cuentas = Object.keys(resp.cuentas);
        var periodos = Object.keys(resp.periodos);
        var iglesiasArray = new Array();
        var i, j, k;
        var height = 27 * iglesias.length;
        for(i=0;i<iglesias.length;i++)
        {
          iglesiasArray.push(resp.diezmos[iglesias[i]].Nombre);
        }
        var cadJS='<script type="text/javascript">';
        for(i=0;i<cuentas.length;i++)
        {
          var cuenta = resp.cuentas[cuentas[i]].ACNT_CODE;
          var ncuenta = resp.cuentas[cuentas[i]].DESCR;
          /*for(j=0;j<periodos.length;j++) 
          {
            var PERIOD = periodos[j];*/
            var datos = new Array();
            for(k=0;k<iglesias.length;k++)  
            {
              var valor = 0;
              //var PERIODOS = Object.keys();
              for(j=0;j<periodos.length;j++) 
              {
                var PERIOD = periodos[j];
                if(PERIOD in resp.diezmos[iglesias[k]].PeriodosList)
                {
                  if(cuenta in resp.diezmos[iglesias[k]].PeriodosList[PERIOD].CuentasList)
                  {
                    valor += resp.diezmos[iglesias[k]].PeriodosList[PERIOD].CuentasList[cuenta].Saldo;
                  } 
                }
              }
              if(valor!=0)
              {
                valor = Math.round(valor * 100) / 100  
              }
              datos.push(valor);
            }
            var cad = '<div class="col-md-12 col-sm-12 col-xs-12">'+
                '<div class="x_panel">'+
                  '<div class="x_title">'+
                    '<h2>'+ncuenta+' '+delPeriodo+' al '+alPeriodo+'<small>Personas</small></h2>'+
                    '<div class="clearfix"></div>'+
                  '</div>'+
                  '<div class="x_content">'+
                    '<div id="echart_bar_horizontal'+cuenta+'" style="height: '+height+'px; -webkit-tap-highlight-color: transparent; -webkit-user-select: none; position: relative; background-color: transparent;"><div style="position: relative; overflow: hidden; width: 245px; height: '+height+'px; cursor: default;"><canvas width="490" height="'+(height+370)+'" data-zr-dom-id="zr_0" style="position: absolute; left: 0px; top: 0px; width: 245px; height: '+height+'px; -webkit-user-select: none; -webkit-tap-highlight-color: rgba(0, 0, 0, 0);"></canvas></div><div style="position: absolute; display: none; border: 0px solid rgb(51, 51, 51); white-space: nowrap; z-index: 9999999; transition: left 0.4s cubic-bezier(0.23, 1, 0.32, 1), top 0.4s cubic-bezier(0.23, 1, 0.32, 1); border-radius: 4px; color: rgb(255, 255, 255); font-style: normal; font-variant: normal; font-weight: normal; font-stretch: normal; font-size: 14px; font-family: Arial, Verdana, sans-serif; line-height: 21px; padding: 5px; left: 46px; top: 184.167px; background-color: rgba(0, 0, 0, 0.498039);">Apr<br><span style="display:inline-block;margin-right:5px;border-radius:10px;width:9px;height:9px;background-color:#26B99A"></span>2015 : 104,970<br><span style="display:inline-block;margin-right:5px;border-radius:10px;width:9px;height:9px;background-color:#34495E"></span>2016 : 121,594</div></div>'+
                  '</div>'+
                '</div>'+
              '</div>';

          $("#graficas").append(cad);


           var echartBar = echarts.init(document.getElementById('echart_bar_horizontal'+cuenta), theme);
            echartBar.setOption({
              title: {
                text: ncuenta,
                subtext: cuenta
              },
              tooltip: {
                trigger: 'axis'
              },
              legend: {
                x: 100,
                data: [ncuenta]
              },
              toolbox: {
                show: true,
                feature: {
                  saveAsImage: {
                    show: true,
                    title: "Save Image"
                  }
                }
              },
              calculable: true,
              xAxis: [{
                type: 'value',
                boundaryGap: [0, 0.01]
              }],
              yAxis: [{
                type: 'category',
                data: iglesiasArray
              }],
              series: [{
                name: ncuenta,
                type: 'bar',
                data: datos
              }]
            });
         // }//for J
         }// for I
        //cadJS = cadJS + '<\/script>';
        //$("#graficasJS").html(cadJS);
      }
       if(resp.success==0)
      {
        alert("Hubo un error, perdon");
      }
    }
  });  
});
$(document).on("click", "#generarReporteDeIglesias", function (){
  var delPeriodo = $("#delPeriodo").val();
  var alPeriodo = $("#alPeriodo").val();

  var param = {delPeriodo : delPeriodo, alPeriodo : alPeriodo}
  $.ajax({
    url: "/generarReporteDeIglesias",
    dataType : "json",
    type : "post",
    data: param,
    async : true,
    beforeSend : function (){
    },
    complete : function (){
    }, 
    success : function (resp){
      if(resp.success==1)
      {
        Chart.defaults.global.legend = {
          enabled: true
        };
         Chart.defaults.global.title = {
          enabled: true
        };
        var iglesias = Object.keys(resp.diezmos);
        var cuentas = Object.keys(resp.cuentas);
        var periodos = Object.keys(resp.periodos);
        var iglesiasArray = new Array();
        var i, j, k;
        var height = 27 * iglesias.length;
        for(i=0;i<iglesias.length;i++)
        {
          iglesiasArray.push(resp.diezmos[iglesias[i]].Nombre);
        }
        var cadJS='<script type="text/javascript">';
        for(i=0;i<cuentas.length;i++)
        {
          var cuenta = resp.cuentas[cuentas[i]].ACNT_CODE;
          var ncuenta = resp.cuentas[cuentas[i]].DESCR;
          /*for(j=0;j<periodos.length;j++) 
          {
            var PERIOD = periodos[j];*/
            var datos = new Array();
            for(k=0;k<iglesias.length;k++)  
            {
              var valor = 0;
              //var PERIODOS = Object.keys();
              for(j=0;j<periodos.length;j++) 
              {
                var PERIOD = periodos[j];
                if(PERIOD in resp.diezmos[iglesias[k]].PeriodosList)
                {
                  if(cuenta in resp.diezmos[iglesias[k]].PeriodosList[PERIOD].CuentasList)
                  {
                    valor += resp.diezmos[iglesias[k]].PeriodosList[PERIOD].CuentasList[cuenta].Saldo;
                  } 
                }
              }
              if(valor!=0)
              {
                valor = Math.round(valor * 100) / 100  
              }
              datos.push(valor);
            }
            var cad = '<div class="col-md-12 col-sm-12 col-xs-12">'+
                '<div class="x_panel">'+
                  '<div class="x_title">'+
                    '<h2>'+ncuenta+' '+delPeriodo+' al '+alPeriodo+'<small>Iglesias</small></h2>'+
                    '<div class="clearfix"></div>'+
                  '</div>'+
                  '<div class="x_content">'+
                    '<div id="echart_bar_horizontal'+cuenta+'" style="height: '+height+'px; -webkit-tap-highlight-color: transparent; -webkit-user-select: none; position: relative; background-color: transparent;"><div style="position: relative; overflow: hidden; width: 245px; height: '+height+'px; cursor: default;"><canvas width="490" height="'+(height+370)+'" data-zr-dom-id="zr_0" style="position: absolute; left: 0px; top: 0px; width: 245px; height: '+height+'px; -webkit-user-select: none; -webkit-tap-highlight-color: rgba(0, 0, 0, 0);"></canvas></div><div style="position: absolute; display: none; border: 0px solid rgb(51, 51, 51); white-space: nowrap; z-index: 9999999; transition: left 0.4s cubic-bezier(0.23, 1, 0.32, 1), top 0.4s cubic-bezier(0.23, 1, 0.32, 1); border-radius: 4px; color: rgb(255, 255, 255); font-style: normal; font-variant: normal; font-weight: normal; font-stretch: normal; font-size: 14px; font-family: Arial, Verdana, sans-serif; line-height: 21px; padding: 5px; left: 46px; top: 184.167px; background-color: rgba(0, 0, 0, 0.498039);">Apr<br><span style="display:inline-block;margin-right:5px;border-radius:10px;width:9px;height:9px;background-color:#26B99A"></span>2015 : 104,970<br><span style="display:inline-block;margin-right:5px;border-radius:10px;width:9px;height:9px;background-color:#34495E"></span>2016 : 121,594</div></div>'+
                  '</div>'+
                '</div>'+
              '</div>';

          $("#graficas").append(cad);


           var echartBar = echarts.init(document.getElementById('echart_bar_horizontal'+cuenta), theme);
            echartBar.setOption({
              title: {
                text: ncuenta,
                subtext: cuenta
              },
              tooltip: {
                trigger: 'axis'
              },
              legend: {
                x: 100,
                data: [ncuenta]
              },
              toolbox: {
                show: true,
                feature: {
                  saveAsImage: {
                    show: true,
                    title: "Save Image"
                  }
                }
              },
              calculable: true,
              xAxis: [{
                type: 'value',
                boundaryGap: [0, 0.01]
              }],
              yAxis: [{
                type: 'category',
                data: iglesiasArray
              }],
              series: [{
                name: ncuenta,
                type: 'bar',
                data: datos
              }]
            });
         // }//for J
         }// for I
        //cadJS = cadJS + '<\/script>';
        //$("#graficasJS").html(cadJS);
      }
       if(resp.success==0)
      {
        alert("Hubo un error, perdon");
      }
    }
  });  
});

   $(document).on("click", "#generarBUNIT", function (){
    var BUNIT = $("#BUNIT").val();
    var cuadre = $("#cuadre").val();
    var descr = $("#descr").val();
    if(BUNIT.length==3)
    {
      if(cuadre=="")
      {
        alert("Tienes que escribir la cuenta de cuadre");
        return;
      }
      if(descr=="")
      {
        alert("Tienes que escribir la descripcion de la unidad de negocio");
        return;
      }
      var param = {BUNIT : BUNIT, cuadre : cuadre, descr : descr}
      $.ajax({
        url: "/generarBUNIT",
        dataType : "json",
        type : "post",
        data: param,
        async : true,
        beforeSend : function (){
    //      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
        },
        complete : function (){
        }, 
        success : function (resp){
          if(resp.success==1)
          {
            alert("Unidad de negocio creada!");
          }
           if(resp.success==2)
          {
            alert("Ya existe esa Unidad de negocio!");
          }
        }
      });  
    }
    else
    {
      alert("La Unidad de Negocio debe de tener 3 letras");
    }
    
  });
   
var Base64={_keyStr:"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=",encode:function(e){var t="";var n,r,i,s,o,u,a;var f=0;e=Base64._utf8_encode(e);while(f<e.length){n=e.charCodeAt(f++);r=e.charCodeAt(f++);i=e.charCodeAt(f++);s=n>>2;o=(n&3)<<4|r>>4;u=(r&15)<<2|i>>6;a=i&63;if(isNaN(r)){u=a=64}else if(isNaN(i)){a=64}t=t+this._keyStr.charAt(s)+this._keyStr.charAt(o)+this._keyStr.charAt(u)+this._keyStr.charAt(a)}return t},decode:function(e){var t="";var n,r,i;var s,o,u,a;var f=0;e=e.replace(/[^A-Za-z0-9+/=]/g,"");while(f<e.length){s=this._keyStr.indexOf(e.charAt(f++));o=this._keyStr.indexOf(e.charAt(f++));u=this._keyStr.indexOf(e.charAt(f++));a=this._keyStr.indexOf(e.charAt(f++));n=s<<2|o>>4;r=(o&15)<<4|u>>2;i=(u&3)<<6|a;t=t+String.fromCharCode(n);if(u!=64){t=t+String.fromCharCode(r)}if(a!=64){t=t+String.fromCharCode(i)}}t=Base64._utf8_decode(t);return t},_utf8_encode:function(e){e=e.replace(/rn/g,"n");var t="";for(var n=0;n<e.length;n++){var r=e.charCodeAt(n);if(r<128){t+=String.fromCharCode(r)}else if(r>127&&r<2048){t+=String.fromCharCode(r>>6|192);t+=String.fromCharCode(r&63|128)}else{t+=String.fromCharCode(r>>12|224);t+=String.fromCharCode(r>>6&63|128);t+=String.fromCharCode(r&63|128)}}return t},_utf8_decode:function(e){var t="";var n=0;var r=c1=c2=0;while(n<e.length){r=e.charCodeAt(n);if(r<128){t+=String.fromCharCode(r);n++}else if(r>191&&r<224){c2=e.charCodeAt(n+1);t+=String.fromCharCode((r&31)<<6|c2&63);n+=2}else{c2=e.charCodeAt(n+1);c3=e.charCodeAt(n+2);t+=String.fromCharCode((r&15)<<12|(c2&63)<<6|c3&63);n+=3}}return t}}

$(document).on("click", "#guardarConfigOpciones", function (){
    var tipoDimension = $("#selectTipoD").val();
    var Empresa = $("#Empresa").val().trim();
    var pass = Base64.encode($("#passwordDeSAT").val());
    var param = {pass : pass, Empresa : Empresa, tipoDimension : tipoDimension}
    $.ajax({
      url: "/guardarConfigOpciones",
      dataType : "json",
      type : "post",
      data: param,
      async : true,
      beforeSend : function (){
      },
      complete : function (){
      }, 
      success : function (resp){
        if(resp.success==1)
        {
          alert("Cambios guardados!");
        }
      }
    });  
  });
   
   $(document).on("click", "#guardarBUNIT", function (){
    var BUNIT = $("#selectBUNIT").val();
    var param = {BUNIT : BUNIT}
    $.ajax({
      url: "/guardarBUNIT",
      dataType : "json",
      type : "post",
      data: param,
      async : true,
      beforeSend : function (){
      },
      complete : function (){
      }, 
      success : function (resp){
        if(resp.success==1)
        {
          $("#BUNITACTUAL").html(BUNIT);
          alert("Unidad de negocio cambiada");
        }
      }
    });  
  });
   

$(document).on("click", "#configOpciones", function (){
  $.ajax({
    url: "/configOpciones",
    dataType : "json",
    type : "post",
    async : true,
    beforeSend : function (){
//      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
    },
    complete : function (){
    }, 
    success : function (resp){
      var cad = '<div class="">'+
          '<div class="page-title">'+
            '<div class="title_left">'+
              '<h3>Opciones de la unidad de negocio</h3>'+
            '</div>'+
          '</div>'+
          '<div class="clearfix"></div>'+
          '<div class="row">'+
            '<div class="col-md-12 col-sm-12 col-xs-12">'+
              '<div class="x_panel">'+
                '<div class="x_title">'+
                  '<h2>Opciones de la unidad de negocio <small></small></h2>'+
                  '<ul class="nav navbar-right panel_toolbox">'+
                    '<li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>'+
                    '</li>'+
                    '<li><a class="close-link"><i class="fa fa-close"></i></a>'+
                    '</li>'+
                  '</ul>'+
                  '<div class="clearfix"></div>'+
                '</div>'+
                '<div class="x_content">'+
                  '<br />'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Nombre de la empresa: (carpeta en C:/Compacw/Empresas/)</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" maxlength="50" name="Empresa" id="Empresa" class="form-control col-md-10" style="float: left;" value="'+resp.Empresa+'" />'+
                      '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Password de : (carpeta en C:/Compacw/Empresas/)</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" maxlength="100" id="passwordDeSAT" class="form-control col-md-10" style="float: left;" value="'+Base64.decode(resp.pass)+'" />'+
                      '</div>'+
                    '</div>'+
                     '<div class="form-group">'+
                      '<p><label for="selectTipoD">Seleccionar uso de unidades de negocio:</label>'+
                          '<select id="selectTipoD" class="form-control" required="" >'+
                            '<option value="1">Las obligatorias</option>'+
                            '<option value="2">Las opcionales y las obligatorias</option>'+
                            '<option value="3">Ninguna</option>'+
                            '<option value="4">Todas</option>'+
                          '</select></p>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<div class="col-md-6 col-sm-6 col-xs-12 col-md-offset-3">'+
                        '<button type="submit" id="guardarConfigOpciones" class="btn btn-success">Guardar</button>'+
                      '</div>'+
                    '</div>'+
                '</div>'+
              '</div>'+
            '</div>'+
          '</div>'+
        '</div>';
         $("#mainContent").html(cad);
      $('#selectTipoD option[value="'+resp.tipoDimension+'"]').attr('selected', 'selected');
    }});  
  });

$(document).on("click", "#cambiarUnidadDeNegocio", function (){
  $.ajax({
    url: "/cambiarUnidadDeNegocio",
    dataType : "json",
    type : "post",
    async : true,
    beforeSend : function (){
//      myMsgSpinner = msgSpinner("Obteniendo la informacion...");
    },
    complete : function (){
    }, 
    success : function (resp){
      var bunit = resp.BUNIT;
      var aver2 = Object.keys(resp.unidades);
      var cad = '<div class="">'+
          '<div class="page-title">'+
            '<div class="title_left">'+
              '<h3>Cambiar de unidad de negocio</h3>'+
            '</div>'+
          '</div>'+
          '<div class="clearfix"></div>'+
          '<div class="row">'+
            '<div class="col-md-12 col-sm-12 col-xs-12">'+
              '<div class="x_panel">'+
                '<div class="x_title">'+
                  '<h2>Cambiar unidad de negocio <small></small></h2>'+
                  '<ul class="nav navbar-right panel_toolbox">'+
                    '<li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>'+
                    '</li>'+
                    '<li><a class="close-link"><i class="fa fa-close"></i></a>'+
                    '</li>'+
                  '</ul>'+
                  '<div class="clearfix"></div>'+
                '</div>'+
                '<div class="x_content">'+
                  '<br />'+
                    '<div class="form-group">'+
                      '<p><label for="selectBUNIT">Seleccionar unidad de negocio:</label>'+
                          '<select id="selectBUNIT" class="form-control" required="" >'+
                            '</select></p>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<div class="col-md-6 col-sm-6 col-xs-12 col-md-offset-3">'+
                        '<button type="submit" id="guardarBUNIT" class="btn btn-success">Guardar</button>'+
                      '</div>'+
                    '</div>'+
                '</div>'+
              '</div>'+
            '</div>'+
          '</div>'+
        '</div>';
         $("#mainContent").html(cad);
       cad = '';
      for(i=0;i<aver2.length;i++)
      {
        cad=cad+'<option value="'+aver2[i].toString()+'">'+aver2[i].toString()+'</option>';
      }          
      $("#selectBUNIT").html(cad);
      $('#selectBUNIT option[value="'+bunit+'"]').attr('selected', 'selected');
    }});  
  });



$(document).on("click", "#iglesiasReporte", function (){
 
      var cad = '<div class="">'+
          '<div class="page-title">'+
            '<div class="title_left">'+
              '<h3>Diezmos de iglesias</h3>'+
            '</div>'+
          '</div>'+
          '<div class="clearfix"></div>'+
          '<div class="row">'+
            '<div class="col-md-12 col-sm-12 col-xs-12">'+
              '<div class="x_panel">'+
                '<div class="x_title">'+
                  '<h2>Diezmos de iglesias<small></small></h2>'+
                  '<ul class="nav navbar-right panel_toolbox">'+
                    '<li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>'+
                    '</li>'+
                    '<li><a class="close-link"><i class="fa fa-close"></i></a>'+
                    '</li>'+
                  '</ul>'+
                  '<div class="clearfix"></div>'+
                '</div>'+
                '<div class="x_content">'+
                  '<br />'+
                  '<div class="form-group">'+
                      '<p><label for="delPeriodo">Del periodo:</label>'+
                          '<select id="delPeriodo" class="form-control" required="" >'+
                            '</select></p>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<p><label for="alPeriodo">Al periodo:</label>'+
                          '<select id="alPeriodo" class="form-control" required="">'+
                            '</select></p>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<div class="col-md-6 col-sm-6 col-xs-12 col-md-offset-3">'+
                        '<button type="submit" id="generarReporteDeIglesias" class="btn btn-success">Generar</button>'+
                      '</div>'+
                    '</div>'+
                '</div>'+
              '</div>'+
            '</div>'+
          '</div>'+
          '<div class="clearfix"></div>'+
          '<div class="row" id="graficas"></div>'+
          '<div stye="display:none" id="graficasJS"></div>'+
        '</div>';
      $("#mainContent").html(cad);
      $.ajax({
        url: "/estadosDeCuenta",
        dataType : "json",
        type : "post",
        async : true,
        beforeSend : function (){
        },
        complete : function (){
        }, 
        success : function (resp){
          var aver2 = Object.keys(resp.periodos);
          cad = '';
          for(i=0;i<aver2.length;i++)
          {
            cad=cad+'<option value="'+aver2[i].toString().trim()+'">'+aver2[i].toString().trim()+'</option>';
          } 
          $("#delPeriodo").html(cad);
          $("#alPeriodo").html(cad);
          $("#alPeriodo option:last").attr("selected","selected");
        }
      });
  });


$(document).on("click", "#gastosPersonasReporte", function (){
 
      var cad = '<div class="">'+
          '<div class="page-title">'+
            '<div class="title_left">'+
              '<h3>Gastos de Ministros/Asociados</h3>'+
            '</div>'+
          '</div>'+
          '<div class="clearfix"></div>'+
          '<div class="row">'+
            '<div class="col-md-12 col-sm-12 col-xs-12">'+
              '<div class="x_panel">'+
                '<div class="x_title">'+
                  '<h2>Gastos de Ministros/Asociados<small></small></h2>'+
                  '<ul class="nav navbar-right panel_toolbox">'+
                    '<li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>'+
                    '</li>'+
                    '<li><a class="close-link"><i class="fa fa-close"></i></a>'+
                    '</li>'+
                  '</ul>'+
                  '<div class="clearfix"></div>'+
                '</div>'+
                '<div class="x_content">'+
                  '<br />'+
                  '<div class="form-group">'+
                      '<p><label for="delPeriodo">Del periodo:</label>'+
                          '<select id="delPeriodo" class="form-control" required="" >'+
                            '</select></p>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<p><label for="alPeriodo">Al periodo:</label>'+
                          '<select id="alPeriodo" class="form-control" required="">'+
                            '</select></p>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<div class="col-md-6 col-sm-6 col-xs-12 col-md-offset-3">'+
                        '<button type="submit" id="generarReporteDeMAT" class="btn btn-success">Generar</button>'+
                      '</div>'+
                    '</div>'+
                '</div>'+
              '</div>'+
            '</div>'+
          '</div>'+
          '<div class="clearfix"></div>'+
          '<div class="row" id="graficas"></div>'+
          '<div stye="display:none" id="graficasJS"></div>'+
        '</div>';
      $("#mainContent").html(cad);
      $.ajax({
        url: "/estadosDeCuenta",
        dataType : "json",
        type : "post",
        async : true,
        beforeSend : function (){
        },
        complete : function (){
        }, 
        success : function (resp){
          var aver2 = Object.keys(resp.periodos);
          cad = '';
          for(i=0;i<aver2.length;i++)
          {
            cad=cad+'<option value="'+aver2[i].toString().trim()+'">'+aver2[i].toString().trim()+'</option>';
          } 
          $("#delPeriodo").html(cad);
          $("#alPeriodo").html(cad);
          $("#alPeriodo option:last").attr("selected","selected");
        }
      });
  });

$(document).on("click", "#crearUnidadDeNegocio", function (){
 
      var cad = '<div class="">'+
          '<div class="page-title">'+
            '<div class="title_left">'+
              '<h3>Estado de cuenta</h3>'+
            '</div>'+
          '</div>'+
          '<div class="clearfix"></div>'+
          '<div class="row">'+
            '<div class="col-md-12 col-sm-12 col-xs-12">'+
              '<div class="x_panel">'+
                '<div class="x_title">'+
                  '<h2>Estado de cuenta <small></small></h2>'+
                  '<ul class="nav navbar-right panel_toolbox">'+
                    '<li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>'+
                    '</li>'+
                    '<li><a class="close-link"><i class="fa fa-close"></i></a>'+
                    '</li>'+
                  '</ul>'+
                  '<div class="clearfix"></div>'+
                '</div>'+
                '<div class="x_content">'+
                  '<br />'+
                     '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Nombre de la nueva unidad de negocio: (3 letras)</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" maxlength="3" name="BUNIT" id="BUNIT" class="form-control col-md-10" style="float: left;" />'+
                      '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Descripci&oacute;n de la unidad de negocio</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text"  name="descr" id="descr" class="form-control col-md-10" style="float: left;" />'+
                      '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Cuenta de cuadre:</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text"  name="cuadre" id="cuadre" class="form-control col-md-10" style="float: left;" />'+
                      '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<div class="col-md-6 col-sm-6 col-xs-12 col-md-offset-3">'+
                        '<button type="submit" id="generarBUNIT" class="btn btn-success">Generar</button>'+
                      '</div>'+
                    '</div>'+
                '</div>'+
              '</div>'+
            '</div>'+
          '</div>'+
        '</div>';
      $("#mainContent").html(cad);
  });

$(document).on("click", ".selectEmpiezaDimension", function (){
  var dimen = $(this).attr("dimen");
  var thisValue = $(this).val();
  var nombreEnd = "ANAL_T"+dimen+"_END";
  var bandera = 1;
  $('#'+nombreEnd).children('option').each(function () {
    if(this.value==thisValue)
    {
      bandera = 0;
    }
    if(bandera==1)
    {
      $(this).css("display","none");  
    }
    else
    {
      $(this).css("display","block");   
    }    
  });
});
$(document).on("change", ".selectEmpiezaDimension", function (){
  var dimen = $(this).attr("dimen");
  var thisValue = $(this).val();
  var nombreEnd = "ANAL_T"+dimen+"_END";
  $("#"+nombreEnd).val(thisValue);
});
$(document).on("click", "#estadosDeCuenta", function (){
  $.ajax({
    url: "/estadosDeCuenta",
    dataType : "json",
    type : "post",
    async : true,
    beforeSend : function (){
    },
    complete : function (){
    }, 
    success : function (resp){
      var cuentasArray = new Array();
      
      var descrArray = new Array();
      var aver2 = Object.keys(resp.cuentas);
      var ANAL_T0 = Object.keys(resp.ANAL_T0);
      var ANAL_T1 = Object.keys(resp.ANAL_T1);
      var ANAL_T2 = Object.keys(resp.ANAL_T2);
      var ANAL_T3 = Object.keys(resp.ANAL_T3);
      var ANAL_T4 = Object.keys(resp.ANAL_T4);
      var ANAL_T5 = Object.keys(resp.ANAL_T5);
      var ANAL_T6 = Object.keys(resp.ANAL_T6);
      var ANAL_T7 = Object.keys(resp.ANAL_T7);
      var ANAL_T8 = Object.keys(resp.ANAL_T8);
      var ANAL_T9 = Object.keys(resp.ANAL_T9);

      
      var cad = '<div class="">'+
          '<div class="page-title">'+
            '<div class="title_left">'+
              '<h3>Estado de cuenta</h3>'+
            '</div>'+
          '</div>'+
          '<div class="clearfix"></div>'+
          '<div class="row">'+
            '<div class="col-md-12 col-sm-12 col-xs-12">'+
              '<div class="x_panel">'+
                '<div class="x_title">'+
                  '<h2>Estado de cuenta <small></small></h2>'+
                  '<ul class="nav navbar-right panel_toolbox">'+
                    '<li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>'+
                    '</li>'+
                    '<li><a class="close-link"><i class="fa fa-close"></i></a>'+
                    '</li>'+
                  '</ul>'+
                  '<div class="clearfix"></div>'+
                '</div>'+
                '<div class="x_content">'+
                  '<br />'+
                     '<div class="form-group">'+
                      '<label class="control-label col-md-3 col-sm-3 col-xs-12">Cuenta</label><span class="required">*</span>'+
                      '<div class="col-md-9 col-sm-9 col-xs-12">'+
                        '<input type="text" name="country" id="autocomplete-custom-append" class="form-control col-md-10" style="float: left;" />'+
                        '<div id="autocomplete-container" style="position: relative; float: left; width: 400px; margin: 10px;"></div>'+
                      '</div>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<p><label for="delPeriodo">Del periodo:</label>'+
                          '<select id="delPeriodo" class="form-control" required="" >'+
                            '</select></p>'+
                    '</div>'+
                    '<div class="form-group">'+
                      '<p><label for="alPeriodo">Al periodo:</label>'+
                          '<select id="alPeriodo" class="form-control" required="">'+
                            '</select></p>'+
                    '</div>';
                    cad=cad+'<div class="form-group">'+
                      '<div class="col-md-6 col-sm-6 col-xs-12 col-md-offset-3">'+
                        '<button type="submit" id="generarEstadoDeCuenta" class="btn btn-success">Generar</button>'+
                      '</div>'+
                    '</div>';
                    if(ANAL_T0.length>0)
                    {
                      var nombre = resp.ANAL_T0[ANAL_T0[0]]["DESCR"];
                      cad=cad+'<br><div class="form-group">'+
                                  '<p><label style="width: 100%; border-top: 2px solid #E6E9ED;" for="ANAL_T0_START">'+nombre+'</label>'+
                                      '<select id="ANAL_T0_START" class="form-control selectEmpiezaDimension" dimen="0" required="" >'+
                                        '</select> <label for="ANAL_T0_END">A:</label>'+
                                      '<select id="ANAL_T0_END" class="form-control" required="">'+
                                        '</select></p>'+
                                '</div>';
                    }
                    if(ANAL_T1.length>0)
                    {
                      var nombre = resp.ANAL_T1[ANAL_T1[0]]["DESCR"];
                      cad=cad+ '<br><div class="form-group">'+
                                  '<p><label style="width: 100%; border-top: 2px solid #E6E9ED;" for="ANAL_T1_START">'+nombre+'</label>'+
                                      '<select id="ANAL_T1_START" class="form-control selectEmpiezaDimension" dimen="1" required="" >'+
                                        '</select> <label for="ANAL_T1_END">A:</label>'+
                                      '<select id="ANAL_T1_END" class="form-control" required="">'+
                                        '</select></p>'+
                                '</div>';
                    }
                     if(ANAL_T2.length>0)
                    {
                      var nombre = resp.ANAL_T2[ANAL_T2[0]]["DESCR"];
                      cad=cad+ '<br><div class="form-group">'+
                                  '<p><label style="width: 100%; border-top: 2px solid #E6E9ED;" for="ANAL_T2_START">'+nombre+'</label>'+
                                      '<select id="ANAL_T2_START" class="form-control selectEmpiezaDimension" dimen="2" required="" >'+
                                        '</select> <label for="ANAL_T2_END">A:</label>'+
                                      '<select id="ANAL_T2_END" class="form-control" required="">'+
                                        '</select></p>'+
                                '</div>';
                    }
                     if(ANAL_T3.length>0)
                    {
                      var nombre = resp.ANAL_T3[ANAL_T3[0]]["DESCR"];
                      cad=cad+ '<br><div class="form-group">'+
                                  '<p><label style="width: 100%; border-top: 2px solid #E6E9ED;" for="ANAL_T3_START">'+nombre+'</label>'+
                                      '<select id="ANAL_T3_START" class="form-control selectEmpiezaDimension" dimen="3" required="" >'+
                                        '</select> <label for="ANAL_T3_END">A:</label>'+
                                      '<select id="ANAL_T3_END" class="form-control" required="">'+
                                        '</select></p>'+
                                '</div>';
                    }
                     if(ANAL_T4.length>0)
                    {
                      var nombre = resp.ANAL_T4[ANAL_T4[0]]["DESCR"];
                      cad=cad+ '<br><div class="form-group">'+
                                  '<p><label style="width: 100%; border-top: 2px solid #E6E9ED;" for="ANAL_T4_START">'+nombre+'</label>'+
                                      '<select id="ANAL_T4_START" class="form-control selectEmpiezaDimension" dimen="4" required="" >'+
                                        '</select> <label for="ANAL_T4_END">A:</label>'+
                                      '<select id="ANAL_T4_END" class="form-control" required="">'+
                                        '</select></p>'+
                                '</div>';
                    }
                     if(ANAL_T5.length>0)
                    {
                      var nombre = resp.ANAL_T5[ANAL_T5[0]]["DESCR"];
                      cad=cad+ '<br><div class="form-group">'+
                                  '<p><label style="width: 100%; border-top: 2px solid #E6E9ED;" for="ANAL_T5_START">'+nombre+'</label>'+
                                      '<select id="ANAL_T5_START" class="form-control selectEmpiezaDimension" dimen="5" required="" >'+
                                        '</select> <label for="ANAL_T5_END">A:</label>'+
                                      '<select id="ANAL_T5_END" class="form-control" required="">'+
                                        '</select></p>'+
                                '</div>';
                    }
                     if(ANAL_T6.length>0)
                    {
                      var nombre = resp.ANAL_T6[ANAL_T6[0]]["DESCR"];
                      cad=cad+ '<br><div class="form-group">'+
                                  '<p><label style="width: 100%; border-top: 2px solid #E6E9ED;" for="ANAL_T6_START">'+nombre+'</label>'+
                                      '<select id="ANAL_T6_START" class="form-control selectEmpiezaDimension" dimen="6" required="" >'+
                                        '</select> <label for="ANAL_T6_END">A:</label>'+
                                      '<select id="ANAL_T6_END" class="form-control" required="">'+
                                        '</select></p>'+
                                '</div>';
                    }
                     if(ANAL_T7.length>0)
                    {
                      var nombre = resp.ANAL_T7[ANAL_T7[0]]["DESCR"];
                      cad=cad+ '<br><div class="form-group">'+
                                  '<p><label style="width: 100%; border-top: 2px solid #E6E9ED;" for="ANAL_T7_START">'+nombre+'</label>'+
                                      '<select id="ANAL_T7_START" class="form-control selectEmpiezaDimension" dimen="7" required="" >'+
                                        '</select> <label for="ANAL_T7_END">A:</label>'+
                                      '<select id="ANAL_T7_END" class="form-control" required="">'+
                                        '</select></p>'+
                                '</div>';
                    }
                     if(ANAL_T8.length>0)
                    {
                      var nombre = resp.ANAL_T8[ANAL_T8[0]]["DESCR"];
                      cad=cad+ '<br><div class="form-group">'+
                                  '<p><label style="width: 100%; border-top: 2px solid #E6E9ED;" for="ANAL_T8_START">'+nombre+'</label>'+
                                      '<select id="ANAL_T8_START" class="form-control selectEmpiezaDimension" dimen="8" required="" >'+
                                        '</select> <label for="ANAL_T8_END">A:</label>'+
                                      '<select id="ANAL_T8_END" class="form-control" required="">'+
                                        '</select></p>'+
                                '</div>';
                    }
                     if(ANAL_T9.length>0)
                    {
                      var nombre = resp.ANAL_T9[ANAL_T9[0]]["DESCR"];
                      cad=cad+ '<br><div class="form-group">'+
                                  '<p><label style="width: 100%; border-top: 2px solid #E6E9ED;" for="ANAL_T9_START">'+nombre+'</label>'+
                                      '<select id="ANAL_T9_START" class="form-control selectEmpiezaDimension" dimen="9" required="" >'+
                                        '</select> <label for="ANAL_T9_END">A:</label>'+
                                      '<select id="ANAL_T9_END" class="form-control" required="">'+
                                        '</select></p>'+
                                '</div>';
                    }

                    cad=cad+'<div class="form-group">'+
                      '<div class="col-md-6 col-sm-6 col-xs-12 col-md-offset-3">'+
                        '<button type="submit" id="generarEstadoDeCuenta" class="btn btn-success">Generar</button>'+
                      '</div>'+
                    '</div>'+
                '</div>'+
              '</div>'+
            '</div>'+
          '</div>'+
        '</div>';
      for(i=0;i<aver2.length;i++)
      {
        var aver3 = Object.keys(resp.cuentas[aver2[i]]);
        var param = {value : aver2[i].toString(), data : resp.cuentas[aver2[i]]["DESCR"] }
        cuentasArray.push( aver2[i].toString().trim() + "|"+ resp.cuentas[aver2[i]]["DESCR"]);
        //cuentasArray.push(param);
        //descrArray.push(resp.cuentas[aver2[i]]["DESCR"]);
      }          
      $("#mainContent").html(cad);
      aver2 = Object.keys(resp.periodos);
      cad = '';
      for(i=0;i<aver2.length;i++)
      {
        cad=cad+'<option value="'+aver2[i].toString().trim()+'">'+aver2[i].toString().trim()+'</option>';
      } 
      $("#delPeriodo").html(cad);
      $("#alPeriodo").html(cad);

      cad = '<option value="">Selecciona una opci&oacute;n</option>';
      aver2 = Object.keys(resp.ANAL_T0);
      for(i=0;i<aver2.length;i++)
      {
        cad=cad+'<option value="'+aver2[i].toString().trim()+'">'+resp.ANAL_T0[aver2[i]]["NAME"].toString().trim()+'</option>';
      } 
      $("#ANAL_T0_START").html(cad);
      $("#ANAL_T0_END").html(cad);


      cad = '<option value="">Selecciona una opci&oacute;n</option>';
      aver2 = Object.keys(resp.ANAL_T1);
      for(i=0;i<aver2.length;i++)
      {
        cad=cad+'<option value="'+aver2[i].toString().trim()+'">'+resp.ANAL_T1[aver2[i]]["NAME"].toString().trim()+'</option>';
      } 
      $("#ANAL_T1_START").html(cad);
      $("#ANAL_T1_END").html(cad);

      cad = '<option value="">Selecciona una opci&oacute;n</option>';
      aver2 = Object.keys(resp.ANAL_T2);
      for(i=0;i<aver2.length;i++)
      {
        cad=cad+'<option value="'+aver2[i].toString().trim()+'">'+resp.ANAL_T2[aver2[i]]["NAME"].toString().trim()+'</option>';
      } 
      $("#ANAL_T2_START").html(cad);
      $("#ANAL_T2_END").html(cad);


      cad = '<option value="">Selecciona una opci&oacute;n</option>';
      aver2 = Object.keys(resp.ANAL_T3);
      for(i=0;i<aver2.length;i++)
      {
        cad=cad+'<option value="'+aver2[i].toString().trim()+'">'+resp.ANAL_T3[aver2[i]]["NAME"].toString().trim()+'</option>';
      } 
      $("#ANAL_T3_START").html(cad);
      $("#ANAL_T3_END").html(cad);
      
      cad = '<option value="">Selecciona una opci&oacute;n</option>';
      aver2 = Object.keys(resp.ANAL_T4);
      for(i=0;i<aver2.length;i++)
      {
        cad=cad+'<option value="'+aver2[i].toString().trim()+'">'+resp.ANAL_T4[aver2[i]]["NAME"].toString().trim()+'</option>';
      } 
      $("#ANAL_T4_START").html(cad);
      $("#ANAL_T4_END").html(cad);
      
      cad = '<option value="">Selecciona una opci&oacute;n</option>';
      aver2 = Object.keys(resp.ANAL_T5);
      for(i=0;i<aver2.length;i++)
      {
        cad=cad+'<option value="'+aver2[i].toString().trim()+'">'+resp.ANAL_T5[aver2[i]]["NAME"].toString().trim()+'</option>';
      } 
      $("#ANAL_T5_START").html(cad);
      $("#ANAL_T5_END").html(cad);
      
      cad = '<option value="">Selecciona una opci&oacute;n</option>';
      aver2 = Object.keys(resp.ANAL_T6);
      for(i=0;i<aver2.length;i++)
      {
        cad=cad+'<option value="'+aver2[i].toString().trim()+'">'+resp.ANAL_T6[aver2[i]]["NAME"].toString().trim()+'</option>';
      } 
      $("#ANAL_T6_START").html(cad);
      $("#ANAL_T6_END").html(cad);
      
      cad = '<option value="">Selecciona una opci&oacute;n</option>';
      aver2 = Object.keys(resp.ANAL_T7);
      for(i=0;i<aver2.length;i++)
      {
        cad=cad+'<option value="'+aver2[i].toString().trim()+'">'+resp.ANAL_T7[aver2[i]]["NAME"].toString().trim()+'</option>';
      } 
      $("#ANAL_T7_START").html(cad);
      $("#ANAL_T7_END").html(cad);
      
      cad = '<option value="">Selecciona una opci&oacute;n</option>';
      aver2 = Object.keys(resp.ANAL_T8);
      for(i=0;i<aver2.length;i++)
      {
        cad=cad+'<option value="'+aver2[i].toString().trim()+'">'+resp.ANAL_T8[aver2[i]]["NAME"].toString().trim()+'</option>';
      } 
      $("#ANAL_T8_START").html(cad);
      $("#ANAL_T8_END").html(cad);
      
      cad = '<option value="">Selecciona una opci&oacute;n</option>';
      aver2 = Object.keys(resp.ANAL_T9);
      for(i=0;i<aver2.length;i++)
      {
        cad=cad+'<option value="'+aver2[i].toString().trim()+'">'+resp.ANAL_T9[aver2[i]]["NAME"].toString().trim()+'</option>';
      } 
      $("#ANAL_T9_START").html(cad);
      $("#ANAL_T9_END").html(cad);
      
      
      

      $("#alPeriodo option:last").attr("selected","selected");
      $('#autocomplete-custom-append').autocomplete({
        lookup: cuentasArray,
        appendTo: '#autocomplete-container'
      });        
    }});  
  });

function factura(rfc,importe,referencia,fecha,saldo,referencia,leyenda1,correo)
{
   var api ="http://umn.redirectme.net:755/factura";
      var param = {
      Empresa : 'IASD_UMN',
      servicio : '2',
      cliente : rfc,
      concepto : '4',
      precio : importe,
      contra : 'SUFTRDkyNjAzOVNBSQ==',
      correoReceptor : correo,
      _xsrf : '61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o'
    };
    $.ajax({
      url: api,
      data : param,
      dataType : "json",
      type : "GET",
      async : true,
      beforeSend : function (){
      },
      complete : function (){
 var todasLasLineas = {};
          //debito a bancos
          //credito a ingreso x diezmos!
          var anio = fecha.substring(0,4);
          var mes = fecha.substring(5,2);
          var PERIOD = anio+"0"+mes;
          PERIOD = PERIOD.substring(0,7);
          fecha.replace("/","-");
            todasLasLineas[0] = {};
            todasLasLineas[0]["PERIOD"] = PERIOD;
            todasLasLineas[0]["treference"] = "ING "+fecha;
            todasLasLineas[0]["fecha"] = fecha;
            todasLasLineas[0]["cuenta"] = "102110"
            todasLasLineas[0]["fondo"] = "10"
            todasLasLineas[0]["funcion"] = "";
            todasLasLineas[0]["reestriccion"] = "";
            todasLasLineas[0]["orgid"] = "";
            todasLasLineas[0]["who"] = "";
            todasLasLineas[0]["detalle"] = "";
            todasLasLineas[0]["proyecto"] = "";
            todasLasLineas[0]["descripcion"] = "Ingreso automatico";
            todasLasLineas[0]["cantidad"] = -1*parseFloat(importe);


            todasLasLineas[1] = {};
            todasLasLineas[1]["PERIOD"] = PERIOD;
            todasLasLineas[1]["treference"] = "ING "+fecha;
            todasLasLineas[1]["fecha"] = fecha;
            todasLasLineas[1]["cuenta"] = "611110"
            todasLasLineas[1]["fondo"] = "10"
            todasLasLineas[1]["funcion"] = "";
            todasLasLineas[1]["reestriccion"] = "";
            todasLasLineas[1]["orgid"] = "";
            todasLasLineas[1]["who"] = "";
            todasLasLineas[1]["detalle"] = "";
            todasLasLineas[1]["proyecto"] = "";
            todasLasLineas[1]["descripcion"] = "Ingreso automatico";
            todasLasLineas[1]["cantidad"] = parseFloat(importe);
          
          var lineasJSON = JSON.stringify(todasLasLineas);
          var param = { 
            lineas : lineasJSON
          }
          $.ajax({
            url: "/registraDiario",
           dataType : "json",
            data: param,
            type : "post",
            crossDomain: true,
            //jsonpCallback : "requestLineaDelPresupuesto",
            async : true,
            beforeSend : function (){
            },
            error : function (jqXHR, textStatus, errorThrown) {
                alert('Error cos: ' + textStatus + " " + errorThrown);
            },
            complete : function (algo){
              
            }, 
            success : function (resp){
                if(resp.success==1)
                {
                  window.open(
                    "/generarDiario?diario=" + resp.diario,
                    '_blank'
                  );
                }
              }
          });
      }, 
      success : function (resp){
        if(resp.success==1)
        {

        }
      }
    });
setTimeout(function(){
    
        },3000);
}
    function veAVerBancos()
    {
      var api ="http://juntas.adventistasumn.org/ban.php";
      var param = {
        servicio : 'app',
        accion : 'damePendientesPorFacturar',
        rfc : "IAS9306298H0"
      };
    /*$.ajax({
      url: api,
      data : param,
        crossDomain:true,
      dataType : "jsonp",
      type : "post",
      async : true,
      beforeSend : function (){
      },
      complete : function (){
      }, 
      success : function (resp){

        if(resp.success==1)
        {*/
          var i;
          var pendientes = new Array();
         // pendientes.push({fecha:'2016/06/01', saldo:6080954.9, importe: 3.21, referencia:1, leyenda1:'DEPOSITO SALVO BUEN COBRO', leyenda2:'',leyenda3:''});
          pendientes.push({fecha:'2016/06/03', saldo:6081924.9, importe: 1.23, referencia:253000154, leyenda1:'AGMK11 ENVIO REMESAS JUNIO', leyenda2:'',leyenda3:''});
          
          for(i=0;i<pendientes.length;i++)
          {
            var fecha = pendientes[i].fecha;
            var saldo = pendientes[i].saldo;
            var importe = pendientes[i].importe;
            var referencia = pendientes[i].referencia;
            var leyenda1 = pendientes[i].leyenda1;
            var leyenda2 = pendientes[i].leyenda2;
            var leyenda3 = pendientes[i].leyenda3;
            //busca AGM111
            var todos1 = leyenda1.split(" ");
            var j;
            for(j=0;j<todos1.length;j++)
            {
              var palabra = todos1[j];
              var rfc="138CLIEN01";
              var correo = "agm111@unav.edu.mx";
              if(palabra.length==6)
              {
                if(palabra.substring(0,2)=="AG")
                {
                  rfc=palabra;
                  correo=rfc.toLowerCase()+"@unav.edu.mx";
                }
              }
            }
            factura(rfc,importe,referencia,fecha,saldo,referencia,leyenda1,correo);
          }
     /*   }
      }
    });*/

    }
  //  veAVerBancos();
  function requestCampo(resp2)
  {
     if(resp2.success==1)
                {
                  var t1 = Date.now()+parseInt(resp2.idCampo);
                  var t2 = t1+1; 
                  var cadGlobal='';

                   cadGlobal=cadGlobal+' <div class="row">\
          <div class="col-md-12 col-sm-12 col-xs-12">\
                <div class="x_panel">\
                  <div class="x_title">\
                    <h2>'+resp2.nombre+'</h2>\
                    <ul class="nav navbar-right panel_toolbox">\
                      <li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>\
                      </li>\
                      <li class="dropdown">\
                        <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false"><i class="fa fa-wrench"></i></a>\
                        <ul class="dropdown-menu" role="menu">\
                          <li><a href="#">Settings 1</a>\
                          </li>\
                          <li><a href="#">Settings 2</a>\
                          </li>\
                        </ul>\
                      </li>\
                      <li><a class="close-link"><i class="fa fa-close"></i></a>\
                      </li>\
                    </ul>\
                    <div class="clearfix"></div>\
                  </div>\
                  <div class="x_content">\
                    <form class="form-horizontal form-label-left">\
                      <div class="form-group">\
                        <label class="control-label col-md-12 col-sm-12 col-xs-12" style="text-align: center;font-size: 40px;">'+resp2.nombre+'</label>\
                      </div>\
                    </form>\
                    </div>\
                  </div>\
                </div>\
              </div>\
          </div>\
          <div class="clearfix"></div>\
          <div class="row">\
          <div class="col-md-4 col-sm-4 col-xs-12">\
                <div class="x_panel">\
                  <div class="x_title">\
                    <h2>Capital operativo</h2>\
                    <ul class="nav navbar-right panel_toolbox">\
                      <li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>\
                      </li>\
                      <li class="dropdown">\
                        <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false"><i class="fa fa-wrench"></i></a>\
                        <ul class="dropdown-menu" role="menu">\
                          <li><a href="#">Settings 1</a>\
                          </li>\
                          <li><a href="#">Settings 2</a>\
                          </li>\
                        </ul>\
                      </li>\
                      <li><a class="close-link"><i class="fa fa-close"></i></a>\
                      </li>\
                    </ul>\
                    <div class="clearfix"></div>\
                  </div>\
                  <div class="x_content">\
                    <div id="echart_guage'+resp2.idCampo+'" style="height: 370px; -webkit-tap-highlight-color: transparent; -webkit-user-select: none; position: relative; background-color: transparent;" ><div style="position: relative; overflow: hidden; width: 269px; height: 370px; cursor: default;"><canvas width="538" height="740" data-zr-dom-id="zr_0" style="position: absolute; left: 0px; top: 0px; width: 269px; height: 370px; -webkit-user-select: none; -webkit-tap-highlight-color: rgba(0, 0, 0, 0);"></canvas></div><div style="position: absolute; display: none; border: 0px solid rgb(51, 51, 51); white-space: nowrap; z-index: 9999999; transition: left 0.4s cubic-bezier(0.23, 1, 0.32, 1), top 0.4s cubic-bezier(0.23, 1, 0.32, 1); border-radius: 4px; color: rgb(255, 255, 255); font-style: normal; font-variant: normal; font-weight: normal; font-stretch: normal; font-size: 14px; font-family: Arial, Verdana, sans-serif; line-height: 21px; padding: 5px; left: 28.3438px; top: 208px; background-color: rgba(0, 0, 0, 0.498039);">Performance <br>Performance : 50%</div></div>\
                  </div>\
                </div>\
              </div>\
              <div class="col-md-8 col-sm-8 col-xs-12">\
                <div class="x_panel">\
                  <div class="x_title">\
                    <h2>Calculado al periodo: <small>'+resp2.periodoAnterior+' </small></h2>\
                    <ul class="nav navbar-right panel_toolbox">\
                      <li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>\
                      </li>\
                      <li class="dropdown">\
                        <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false"><i class="fa fa-wrench"></i></a>\
                        <ul class="dropdown-menu" role="menu">\
                          <li><a href="#">Settings 1</a>\
                          </li>\
                          <li><a href="#">Settings 2</a>\
                          </li>\
                        </ul>\
                      </li>\
                      <li><a class="close-link"><i class="fa fa-close"></i></a>\
                      </li>\
                    </ul>\
                    <div class="clearfix"></div>\
                  </div>\
                  <div class="x_content">\
                    <table class="table table-striped">\
                      <thead>\
                        <tr>\
                          <th>Concepto</th>\
                          <th>Monto</th>\
                        </tr>\
                      </thead>\
                      <tbody id="tablaCapitalOperativo'+resp2.idCampo+'">';
                        cadGlobal = cadGlobal+ '<tr><td><a href="#" busqueda="\'1\'" tipo="1" class="veDetalle">Activos corrientes</a></td><td>'+format2(resp2.corrientesActivos,"$")+'</td></tr>';
                        cadGlobal=cadGlobal + '<tr><td><a href="#" busqueda="\'3\'" tipo="2" class="veDetalle">Pasivos corrientes</td><td>'+format2(resp2.corrientesPasivos,"$")+'</a></td></tr>';
                        var capitalTrabajoActual = resp2.corrientesActivos - resp2.corrientesPasivos;
                        cadGlobal=cadGlobal + '<tr><td>Capital de trabajo actual</td><td>'+format2(capitalTrabajoActual,"$")+'</td></tr>';
                        cadGlobal=cadGlobal + '<tr><td colspan="2">Capital de trabajo recomendado por el reglamento</td></tr>';
                        var gastosOperativos = resp2.gastosOperativos;
                        cadGlobal=cadGlobal + '<tr><td><a href="#" busqueda="\'8\',\'9\'" tipo="3" class="veDetalle">'+(resp2.pcent*100) +'% de los gastos operativos</td><td>'+format2(gastosOperativos,"$")+'</a></td></tr>';            
                        cadGlobal=cadGlobal + '<tr><td><a href="#" busqueda="\'2\'" tipo="4" class="veDetalle">Saldo fondos asignados</td><td>'+format2(resp2.activosNetosAsignados,"$")+'</a></td></tr>';
                        var capitalDeTrabajoRecomendado = gastosOperativos+resp2.activosNetosAsignados;
                        cadGlobal=cadGlobal + '<tr><td>Capital de trabajo recomendado</td><td>'+format2(capitalDeTrabajoRecomendado,"$")+'</td></tr>';
                        var excedenteODeficit = capitalTrabajoActual - capitalDeTrabajoRecomendado;
                        cadGlobal=cadGlobal + '<tr><td>Excedente (D&eacute;ficit) de real sobre recomendado</td><td>'+format2(excedenteODeficit,"$")+'</td></tr>';
                        var porcentaje = (capitalTrabajoActual/capitalDeTrabajoRecomendado)*100;
                        porcentaje = Math.round(porcentaje * 100) / 100;
                         cadGlobal=cadGlobal + '<tr><td>Porcentaje</td><td>'+porcentaje+' %</td></tr>';
                      cadGlobal=cadGlobal+'</tbody>\
                    </table>\
                  </div>\
                </div>\
              </div>\
              </div>\
              <div class="clearfix"></div>\
               <div class="row">\
              <div class="col-md-4 col-sm-4 col-xs-12">\
                <div class="x_panel">\
                  <div class="x_title">\
                    <h2>Liquidez</h2>\
                    <ul class="nav navbar-right panel_toolbox">\
                      <li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>\
                      </li>\
                      <li class="dropdown">\
                        <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false"><i class="fa fa-wrench"></i></a>\
                        <ul class="dropdown-menu" role="menu">\
                          <li><a href="#">Settings 1</a>\
                          </li>\
                          <li><a href="#">Settings 2</a>\
                          </li>\
                        </ul>\
                      </li>\
                      <li><a class="close-link"><i class="fa fa-close"></i></a>\
                      </li>\
                    </ul>\
                    <div class="clearfix"></div>\
                  </div>\
                  <div class="x_content">\
                    <div id="liquidez_chart'+resp2.idCampo+'" style="height: 370px; -webkit-tap-highlight-color: transparent; -webkit-user-select: none; position: relative; background-color: transparent;"><div style="position: relative; overflow: hidden; width: 269px; height: 370px; cursor: default;"><canvas width="538" height="740" data-zr-dom-id="zr_0" style="position: absolute; left: 0px; top: 0px; width: 269px; height: 370px; -webkit-user-select: none; -webkit-tap-highlight-color: rgba(0, 0, 0, 0);"></canvas></div><div style="position: absolute; display: none; border: 0px solid rgb(51, 51, 51); white-space: nowrap; z-index: 9999999; transition: left 0.4s cubic-bezier(0.23, 1, 0.32, 1), top 0.4s cubic-bezier(0.23, 1, 0.32, 1); border-radius: 4px; color: rgb(255, 255, 255); font-style: normal; font-variant: normal; font-weight: normal; font-stretch: normal; font-size: 14px; font-family: Arial, Verdana, sans-serif; line-height: 21px; padding: 5px; left: 28.3438px; top: 208px; background-color: rgba(0, 0, 0, 0.498039);">Performance <br>Performance : 50%</div></div>\
                  </div>\
                </div>\
              </div>\
              <div class="col-md-8 col-sm-8 col-xs-12">\
                <div class="x_panel">\
                  <div class="x_title">\
                    <h2>Calculado al periodo: <small>'+resp2.periodoAnterior+' </small></h2>\
                    <ul class="nav navbar-right panel_toolbox">\
                      <li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>\
                      </li>\
                      <li class="dropdown">\
                        <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false"><i class="fa fa-wrench"></i></a>\
                        <ul class="dropdown-menu" role="menu">\
                          <li><a href="#">Settings 1</a>\
                          </li>\
                          <li><a href="#">Settings 2</a>\
                          </li>\
                        </ul>\
                      </li>\
                      <li><a class="close-link"><i class="fa fa-close"></i></a>\
                      </li>\
                    </ul>\
                    <div class="clearfix"></div>\
                  </div>\
                  <div class="x_content">\
                    <table class="table table-striped">\
                      <thead>\
                        <tr>\
                          <th>Concepto</th>\
                          <th>Monto</th>\
                        </tr>\
                      </thead>\
                      <tbody id="tablaLiquidez'+resp2.idCampo+'">';
                        cadGlobal = cadGlobal+ '<tr><td><a href="#" busqueda="\'1\'" tipo="5" class="veDetalle">Efectivo</a></td><td>'+format2(resp2.efectivo,"$")+'</td></tr>';
                       cadGlobal=cadGlobal + '<tr><td><a href="#" busqueda="\'3\'" tipo="6" class="veDetalle">Inversiones</td><td>'+format2(resp2.inversiones,"$")+'</a></td></tr>';
                       cadGlobal=cadGlobal + '<tr><td><a href="#" busqueda="\'3\'" tipo="7" class="veDetalle">Cuentas x cobrar a org. superiores</td><td>'+format2(resp2.cuentaCobrarOrgSuperior,"$")+'</a></td></tr>';
                       var totalDeActivosLiquidos = parseFloat(resp2.efectivo) + parseFloat(resp2.inversiones) + parseFloat(resp2.cuentaCobrarOrgSuperior);
                        cadGlobal=cadGlobal + '<tr><td>Total de activos liquidos</td><td>'+format2(totalDeActivosLiquidos,"$")+'</td></tr>';
                        cadGlobal=cadGlobal + '<tr><td colspan="2">Menos compromisos</td></tr>';
                        cadGlobal=cadGlobal + '<tr><td><a href="#" busqueda="\'3\'" tipo="2" class="veDetalle">Pasivos corrientes</td><td>'+format2(resp2.corrientesPasivos,"$")+'</a></td></tr>';
                        cacadGlobald=cadGlobal + '<tr><td><a href="#" busqueda="\'2\'" tipo="4" class="veDetalle">Saldo fondos asignados</td><td>'+format2(resp2.activosNetosAsignados,"$")+'</a></td></tr>';
                        var totalDeCompromisos = resp2.activosNetosAsignados + resp2.corrientesPasivos;
                        cadGlobal=cadGlobal + '<tr><td>Total de compromisos</td><td>'+format2(totalDeCompromisos,"$")+'</td></tr>';
                        var activosLiquidosNetos = totalDeActivosLiquidos - totalDeCompromisos;
                        var porcentajeLiquidez = (totalDeActivosLiquidos/totalDeCompromisos)*100;
                        porcentajeLiquidez = Math.round(porcentajeLiquidez * 100) / 100;
                      cadGlobal=cadGlobal + '<tr><td>Porcentaje</td><td>'+porcentajeLiquidez+' %</td></tr>';
                      cadGlobal=cadGlobal+'</tbody>\
                    </table>\
                  </div>\
                </div>\
              </div>';
           //   console.log(resp.idCampo+" "+porcentaje+" "+porcentajeLiquidez);
              
                $("#contenidoOtrosCampos").append(cadGlobal);  
              
              
             
      var echartGauge = echarts.init(document.getElementById("echart_guage"+resp2.idCampo), theme);
      echartGauge.setOption({
        tooltip: {
          formatter: "{a} <br/>{b} : {c}%"
        },
        toolbox: {
          show: true,
          feature: {
            restore: {
              show: true,
              title: "Restore"
            },
            saveAsImage: {
              show: true,
              title: "Save Image"
            }
          }
        },
        series: [{
          name: "Capital Operativo",
          type: "gauge",
          center: ["50%", "50%"],
          startAngle: 140,
          endAngle: -140,
          min: 0,
          max: 100,
          precision: 0,
          splitNumber: 10,
          axisLine: {
            show: true,
            lineStyle: {
              color: [
                [0.4, "#ff4500"],
                [0.8, "orange"],
                [0.9, "skyblue"],
                [1, "lightgreen"]
              ],
              width: 30
            }
          },
          axisTick: {
            show: true,
            splitNumber: 5,
            length: 8,
            lineStyle: {
              color: "#eee",
              width: 1,
              type: "solid"
            }
          },
          axisLabel: {
            show: true,
            formatter: function(v) {
              switch (v + "") {
                case "40":
                  return "40";
                case "80":
                  return "80";
                case "90":
                  return "90";
                case "100":
                  return "100";
                default:
                  return "";
              }
            },
            textStyle: {
              color: "#333"
            }
          },
          splitLine: {
            show: true,
            length: 30,
            lineStyle: {
              color: "#eee",
              width: 2,
              type: "solid"
            }
          },
          pointer: {
            length: "80%",
            width: 8,
            color: "auto"
          },
          title: {
            show: true,
            offsetCenter: ["-65%", -10],
            textStyle: {
              color: "#333",
              fontSize: 15
            }
          },
          detail: {
            show: true,
            backgroundColor: "rgba(0,0,0,0)",
            borderWidth: 0,
            borderColor: "#ccc",
            width: 100,
            height: 40,
            offsetCenter: ["-60%", 10],
            formatter: "{value}%",
            textStyle: {
              color: "auto",
              fontSize: 30
            }
          },
          data: [{
            value: porcentaje,
            name: "Capital operativo"
          }]
        }]
      });
      var liquidezChart = echarts.init(document.getElementById("liquidez_chart"+resp2.idCampo), theme);
      liquidezChart.setOption({
        tooltip: {
          formatter: "{a} <br/>{b} : {c}%"
        },
        toolbox: {
          show: true,
          feature: {
            restore: {
              show: true,
              title: "Restore"
            },
            saveAsImage: {
              show: true,
              title: "Guardar"
            }
          }
        },
        series: [{
          name: "Liquidez",
          type: "gauge",
          center: ["50%", "50%"],
          startAngle: 140,
          endAngle: -140,
          min: 0,
          max: 100,
          precision: 0,
          splitNumber: 10,
          axisLine: {
            show: true,
            lineStyle: {
              color: [
              
                [0.4, "#ff4500"],
                [0.8, "orange"],
                [0.9, "skyblue"],
                [1, "lightgreen"]
              ],
              width: 30
            }
          },
          axisTick: {
            show: true,
            splitNumber: 5,
            length: 8,
            lineStyle: {
              color: "#eee",
              width: 1,
              type: "solid"
            }
          },
          axisLabel: {
            show: true,
            formatter: function(v) {
              switch (v + "") {
                case "40":
                  return "40";
                case "80":
                  return "80";
                case "90":
                  return "90";
                case "100":
                  return "100";
                default:
                  return "";
              }
            },
            textStyle: {
              color: "#333"
            }
          },
          splitLine: {
            show: true,
            length: 30,
            lineStyle: {
              color: "#eee",
              width: 2,
              type: "solid"
            }
          },
          pointer: {
            length: "80%",
            width: 8,
            color: "auto"
          },
          title: {
            show: true,
            offsetCenter: ["-65%", -10],
            textStyle: {
              color: "#333",
              fontSize: 15
            }
          },
          detail: {
            show: true,
            backgroundColor: "rgba(0,0,0,0)",
            borderWidth: 0,
            borderColor: "#ccc",
            width: 100,
            height: 40,
            offsetCenter: ["-60%", 10],
            formatter: "{value}%",
            textStyle: {
              color: "auto",
              fontSize: 30
            }
          },
          data: [{
            value: porcentajeLiquidez,
            name: "Liquidez"
          }]
        }]
      });
      indexGlobal++;
      if(indexGlobal<listaParam.length)
      {
        mandaAjaxChecarCampos();
      }
      else
      {
        $("#loadingCampos").html('');
      }
     }//if ajax
  }
  var  URL_GLOBAL, NOMBRE_GLOBAL, idCampo_GLOBAL,primeraVezGlobal;
    var listaParam, indexGlobal;
         
  function mandaAjaxChecarCampos()
  {
      $.ajax({
        url: listaParam[indexGlobal].url+"dashboard",
        dataType : "jsonp",
        type : "GET",
        cache: false,
        crossDomain: true,
        jsonpCallback : "requestCampo",
        data : listaParam[indexGlobal],
        async : false,
       error: function (request, status, error) {
        //error: function () {
            //console.log(request.responseText+"   "+status);
            var cadGlobal='';
            var nombre = listaParam[indexGlobal].nombre;

                   cadGlobal=cadGlobal+' <div class="row">\
          <div class="col-md-12 col-sm-12 col-xs-12">\
                <div class="x_panel">\
                  <div class="x_title">\
                    <h2>'+nombre+' </h2>\
                    <ul class="nav navbar-right panel_toolbox">\
                      <li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>\
                      </li>\
                      <li class="dropdown">\
                        <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false"><i class="fa fa-wrench"></i></a>\
                        <ul class="dropdown-menu" role="menu">\
                          <li><a href="#">Settings 1</a>\
                          </li>\
                          <li><a href="#">Settings 2</a>\
                          </li>\
                        </ul>\
                      </li>\
                      <li><a class="close-link"><i class="fa fa-close"></i></a>\
                      </li>\
                    </ul>\
                    <div class="clearfix"></div>\
                  </div>\
                  <div class="x_content">\
                    <form class="form-horizontal form-label-left">\
                      <div class="form-group">\
                        <label class="control-label col-md-12 col-sm-12 col-xs-12" style="text-align: center;font-size: 40px;">'+nombre+' est&aacute; offline</label>\
                      </div>\
                    </form>\
                    </div>\
                  </div>\
                </div>\
              </div>\
          </div>\
          <div class="clearfix"></div>\
          </div>';   
          $("#contenidoOtrosCampos").append(cadGlobal);  
          indexGlobal++;
          if(indexGlobal<listaParam.length)
          {
            mandaAjaxChecarCampos();
          }
          else
          {
            $("#loadingCampos").html('');
          }
        },
        beforeSend : function (){
        },
        complete : function (){
        }, 
        success : function (resp){
        }
      });//ajax dashboard

  }
  function checaOtrosCampos(PERIOD)
  {
     $.ajax({
      url: "/checaOtrosCampos",
      dataType : "json",
      type : "post",
      async : true,
      beforeSend : function (){
      },
      complete : function (){
      }, 
      success : function (resp1){
        if(resp1.success==1)
        {
          primeraVezGlobal=1;
          $("#loadingCampos").html('<center><img id="loading" src="static/images/loading2.gif" width="30%"  /><center>');
          var campos = Object.keys(resp1.otrosCampos);
          var i = 0;
          listaParam = new Array();
          indexGlobal=0;
          for(i=0;i<campos.length;i++)
          {
            URL_GLOBAL  = resp1.otrosCampos[campos[i]]["URL"];
            NOMBRE_GLOBAL  = resp1.otrosCampos[campos[i]]["Nombre"];
            idCampo_GLOBAL = resp1.otrosCampos[campos[i]]["IdCampo"];
            //hacer una lista con param, y recursivamente cargar esa lista
            listaParam.push({PERIOD : PERIOD, hash : resp1.hash, BUNIT : resp1.BUNIT, nombre: NOMBRE_GLOBAL, idCampo : idCampo_GLOBAL, url :resp1.otrosCampos[campos[i]]["URL"] });
          }
          mandaAjaxChecarCampos();
        }
      }
    });

  }
  function calculaDashboard1()
  {
    var today = new Date();
    var dd = today.getDate();
    var mm = today.getMonth()+1; //January is 0!
    var yyyy = today.getFullYear();
    if(dd<10) {
        dd='0'+dd
    } 

    if(mm<10) {
        mm='0'+mm
    } 

    today = mm+'/'+dd+'/'+yyyy;
    var PERIOD = yyyy+"0"+mm;
    capitalOperativoAlPeriodo(PERIOD);
  }

  function capitalOperativoAlPeriodo(PERIOD){
     $.ajax({
      url: "/dashboard",
      dataType : "json",
      type : "post",
      async : true,
      data : {PERIOD : PERIOD},
      beforeSend : function (){
      },
      complete : function (){
      }, 
      success : function (resp){
        if(resp.success==1)
        {
         var cad = '<div class="">\
          <div class="page-title">\
            <div class="title_left">\
              <h3>Dashboard</h3>\
            </div>\
          </div>\
          <iv class="clearfix"></div>\
          <div class="row">\
          <div class="col-md-12 col-sm-12 col-xs-12">\
                <div class="x_panel">\
                  <div class="x_title">\
                    <h2>Selecciona un periodo</h2>\
                    <ul class="nav navbar-right panel_toolbox">\
                      <li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>\
                      </li>\
                      <li class="dropdown">\
                        <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false"><i class="fa fa-wrench"></i></a>\
                        <ul class="dropdown-menu" role="menu">\
                          <li><a href="#">Settings 1</a>\
                          </li>\
                          <li><a href="#">Settings 2</a>\
                          </li>\
                        </ul>\
                      </li>\
                      <li><a class="close-link"><i class="fa fa-close"></i></a>\
                      </li>\
                    </ul>\
                    <div class="clearfix"></div>\
                  </div>\
                  <div class="x_content">\
                    <form class="form-horizontal form-label-left">\
                      <div class="form-group">\
                        <label class="control-label col-md-3 col-sm-3 col-xs-12">Periodo:</label>\
                        <div class="col-md-9 col-sm-9 col-xs-12">\
                          <select class="form-control" id="periodosSelect">';
                          var periodos = Object.keys(resp.periodos);
                          var j;
                          for(j=0;j<periodos.length;j++)
                          {
                            cad=cad+'<option value="'+periodos[j]+'">'+periodos[j]+'</option>';
                          }
                            cad=cad+'</select>\
                        </div>\
                      </div>\
                    </form>\
                    </div>\
                  </div>\
                </div>\
              </div>\
          </div>\
          <div class="clearfix"></div>\
          <div class="row">\
          <div class="col-md-4 col-sm-4 col-xs-12">\
                <div class="x_panel">\
                  <div class="x_title">\
                    <h2>Capital operativo</h2>\
                    <ul class="nav navbar-right panel_toolbox">\
                      <li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>\
                      </li>\
                      <li class="dropdown">\
                        <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false"><i class="fa fa-wrench"></i></a>\
                        <ul class="dropdown-menu" role="menu">\
                          <li><a href="#">Settings 1</a>\
                          </li>\
                          <li><a href="#">Settings 2</a>\
                          </li>\
                        </ul>\
                      </li>\
                      <li><a class="close-link"><i class="fa fa-close"></i></a>\
                      </li>\
                    </ul>\
                    <div class="clearfix"></div>\
                  </div>\
                  <div class="x_content">\
                    <div id="echart_guage" style="height: 370px; -webkit-tap-highlight-color: transparent; -webkit-user-select: none; position: relative; background-color: transparent;" _echarts_instance_="ec_1469722676878"><div style="position: relative; overflow: hidden; width: 269px; height: 370px; cursor: default;"><canvas width="538" height="740" data-zr-dom-id="zr_0" style="position: absolute; left: 0px; top: 0px; width: 269px; height: 370px; -webkit-user-select: none; -webkit-tap-highlight-color: rgba(0, 0, 0, 0);"></canvas></div><div style="position: absolute; display: none; border: 0px solid rgb(51, 51, 51); white-space: nowrap; z-index: 9999999; transition: left 0.4s cubic-bezier(0.23, 1, 0.32, 1), top 0.4s cubic-bezier(0.23, 1, 0.32, 1); border-radius: 4px; color: rgb(255, 255, 255); font-style: normal; font-variant: normal; font-weight: normal; font-stretch: normal; font-size: 14px; font-family: Arial, Verdana, sans-serif; line-height: 21px; padding: 5px; left: 28.3438px; top: 208px; background-color: rgba(0, 0, 0, 0.498039);">Performance <br>Performance : 50%</div></div>\
                  </div>\
                </div>\
              </div>\
              <div class="col-md-8 col-sm-8 col-xs-12">\
                <div class="x_panel">\
                  <div class="x_title">\
                    <h2>Calculado al periodo: <small>'+resp.periodoAnterior+' </small></h2>\
                    <ul class="nav navbar-right panel_toolbox">\
                      <li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>\
                      </li>\
                      <li class="dropdown">\
                        <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false"><i class="fa fa-wrench"></i></a>\
                        <ul class="dropdown-menu" role="menu">\
                          <li><a href="#">Settings 1</a>\
                          </li>\
                          <li><a href="#">Settings 2</a>\
                          </li>\
                        </ul>\
                      </li>\
                      <li><a class="close-link"><i class="fa fa-close"></i></a>\
                      </li>\
                    </ul>\
                    <div class="clearfix"></div>\
                  </div>\
                  <div class="x_content">\
                    <table class="table table-striped">\
                      <thead>\
                        <tr>\
                          <th>Concepto</th>\
                          <th>Monto</th>\
                        </tr>\
                      </thead>\
                      <tbody id="tablaCapitalOperativo">\
                      </tbody>\
                    </table>\
                  </div>\
                </div>\
              </div>\
              </div>\
              <div class="clearfix"></div>\
               <div class="row">\
              <div class="col-md-4 col-sm-4 col-xs-12">\
                <div class="x_panel">\
                  <div class="x_title">\
                    <h2>Liquidez</h2>\
                    <ul class="nav navbar-right panel_toolbox">\
                      <li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>\
                      </li>\
                      <li class="dropdown">\
                        <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false"><i class="fa fa-wrench"></i></a>\
                        <ul class="dropdown-menu" role="menu">\
                          <li><a href="#">Settings 1</a>\
                          </li>\
                          <li><a href="#">Settings 2</a>\
                          </li>\
                        </ul>\
                      </li>\
                      <li><a class="close-link"><i class="fa fa-close"></i></a>\
                      </li>\
                    </ul>\
                    <div class="clearfix"></div>\
                  </div>\
                  <div class="x_content">\
                    <div id="liquidez_chart" style="height: 370px; -webkit-tap-highlight-color: transparent; -webkit-user-select: none; position: relative; background-color: transparent;" _echarts_instance_="ec_1469722676879"><div style="position: relative; overflow: hidden; width: 269px; height: 370px; cursor: default;"><canvas width="538" height="740" data-zr-dom-id="zr_0" style="position: absolute; left: 0px; top: 0px; width: 269px; height: 370px; -webkit-user-select: none; -webkit-tap-highlight-color: rgba(0, 0, 0, 0);"></canvas></div><div style="position: absolute; display: none; border: 0px solid rgb(51, 51, 51); white-space: nowrap; z-index: 9999999; transition: left 0.4s cubic-bezier(0.23, 1, 0.32, 1), top 0.4s cubic-bezier(0.23, 1, 0.32, 1); border-radius: 4px; color: rgb(255, 255, 255); font-style: normal; font-variant: normal; font-weight: normal; font-stretch: normal; font-size: 14px; font-family: Arial, Verdana, sans-serif; line-height: 21px; padding: 5px; left: 28.3438px; top: 208px; background-color: rgba(0, 0, 0, 0.498039);">Performance <br>Performance : 50%</div></div>\
                  </div>\
                </div>\
              </div>\
              <div class="col-md-8 col-sm-8 col-xs-12">\
                <div class="x_panel">\
                  <div class="x_title">\
                    <h2>Calculado al periodo: <small>'+resp.periodoAnterior+' </small></h2>\
                    <ul class="nav navbar-right panel_toolbox">\
                      <li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>\
                      </li>\
                      <li class="dropdown">\
                        <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false"><i class="fa fa-wrench"></i></a>\
                        <ul class="dropdown-menu" role="menu">\
                          <li><a href="#">Settings 1</a>\
                          </li>\
                          <li><a href="#">Settings 2</a>\
                          </li>\
                        </ul>\
                      </li>\
                      <li><a class="close-link"><i class="fa fa-close"></i></a>\
                      </li>\
                    </ul>\
                    <div class="clearfix"></div>\
                  </div>\
                  <div class="x_content">\
                    <table class="table table-striped">\
                      <thead>\
                        <tr>\
                          <th>Concepto</th>\
                          <th>Monto</th>\
                        </tr>\
                      </thead>\
                      <tbody id="tablaLiquidez">\
                      </tbody>\
                    </table>\
                  </div>\
                </div>\
              </div>\
              <div id="contenidoOtrosCampos"></div>\
              <div id="loadingCampos"></div>\
              </div>';
            $("#mainContent").html(cad);
            checaOtrosCampos(PERIOD);
            $('#periodosSelect option[value="'+PERIOD+'"]').attr("selected","selected");



            $("#periodosSelect").change( function (){
               $("#mainContent").html('<center><img id="loading" src="static/images/loading2.gif" width="30%"  /><center>');
              var PERIOD = $(this).val();
              capitalOperativoAlPeriodo(PERIOD);
              
            });

            var cad = '<tr><td><a href="#" busqueda="\'1\'" tipo="1" class="veDetalle">Activos corrientes</a></td><td>'+format2(resp.corrientesActivos,"$")+'</td></tr>';
            cad=cad + '<tr><td><a href="#" busqueda="\'3\'" tipo="2" class="veDetalle">Pasivos corrientes</td><td>'+format2(resp.corrientesPasivos,"$")+'</a></td></tr>';
            var capitalTrabajoActual = resp.corrientesActivos - resp.corrientesPasivos;
            cad=cad + '<tr><td>Capital de trabajo actual</td><td>'+format2(capitalTrabajoActual,"$")+'</td></tr>';
            cad=cad + '<tr><td colspan="2">Capital de trabajo recomendado por el reglamento</td></tr>';
            var gastosOperativos = resp.gastosOperativos;
            cad=cad + '<tr><td><a href="#" busqueda="\'8\',\'9\'" tipo="3" class="veDetalle">'+(resp.pcent*100) +'% de los gastos operativos</td><td>'+format2(gastosOperativos,"$")+'</a></td></tr>';            
            cad=cad + '<tr><td><a href="#" busqueda="\'2\'" tipo="4" class="veDetalle">Saldo fondos asignados</td><td>'+format2(resp.activosNetosAsignados,"$")+'</a></td></tr>';
            var capitalDeTrabajoRecomendado = gastosOperativos+resp.activosNetosAsignados;
            cad=cad + '<tr><td>Capital de trabajo recomendado</td><td>'+format2(capitalDeTrabajoRecomendado,"$")+'</td></tr>';
            var excedenteODeficit = capitalTrabajoActual - capitalDeTrabajoRecomendado;
            cad=cad + '<tr><td>Excedente (D&eacute;ficit) de real sobre recomendado</td><td>'+format2(excedenteODeficit,"$")+'</td></tr>';
            var porcentaje = (capitalTrabajoActual/capitalDeTrabajoRecomendado)*100;
            porcentaje = Math.round(porcentaje * 100) / 100;
           // porcentaje=Math.round(porcentaje*100);
            cad=cad + '<tr><td>Porcentaje</td><td>'+porcentaje+' %</td></tr>';
            $("#tablaCapitalOperativo").html(cad);


            cad = '<tr><td><a href="#" busqueda="\'1\'" tipo="5" class="veDetalle">Efectivo</a></td><td>'+format2(resp.efectivo,"$")+'</td></tr>';
             cad=cad + '<tr><td><a href="#" busqueda="\'3\'" tipo="6" class="veDetalle">Inversiones</td><td>'+format2(resp.inversiones,"$")+'</a></td></tr>';
             cad=cad + '<tr><td><a href="#" busqueda="\'3\'" tipo="7" class="veDetalle">Cuentas x cobrar a org. superiores</td><td>'+format2(resp.cuentaCobrarOrgSuperior,"$")+'</a></td></tr>';
             var totalDeActivosLiquidos = parseFloat(resp.efectivo) + parseFloat(resp.inversiones) + parseFloat(resp.cuentaCobrarOrgSuperior);
              cad=cad + '<tr><td>Total de activos liquidos</td><td>'+format2(totalDeActivosLiquidos,"$")+'</td></tr>';
              cad=cad + '<tr><td colspan="2">Menos compromisos</td></tr>';
              cad=cad + '<tr><td><a href="#" busqueda="\'3\'" tipo="2" class="veDetalle">Pasivos corrientes</td><td>'+format2(resp.corrientesPasivos,"$")+'</a></td></tr>';
              cad=cad + '<tr><td><a href="#" busqueda="\'2\'" tipo="4" class="veDetalle">Saldo fondos asignados</td><td>'+format2(resp.activosNetosAsignados,"$")+'</a></td></tr>';
              var totalDeCompromisos = resp.activosNetosAsignados + resp.corrientesPasivos;
              cad=cad + '<tr><td>Total de compromisos</td><td>'+format2(totalDeCompromisos,"$")+'</td></tr>';
              var activosLiquidosNetos = totalDeActivosLiquidos - totalDeCompromisos;
              var porcentajeLiquidez = (totalDeActivosLiquidos/totalDeCompromisos)*100;
              porcentajeLiquidez = Math.round(porcentajeLiquidez * 100) / 100;
            cad=cad + '<tr><td>Porcentaje</td><td>'+porcentajeLiquidez+' %</td></tr>';
           
            $("#tablaLiquidez").html(cad);

     
      var echartGauge = echarts.init(document.getElementById("echart_guage"), theme);
      echartGauge.setOption({
        tooltip: {
          formatter: "{a} <br/>{b} : {c}%"
        },
        toolbox: {
          show: true,
          feature: {
            restore: {
              show: true,
              title: "Restore"
            },
            saveAsImage: {
              show: true,
              title: "Save Image"
            }
          }
        },
        series: [{
          name: "Capital Operativo",
          type: "gauge",
          center: ["50%", "50%"],
          startAngle: 140,
          endAngle: -140,
          min: 0,
          max: 100,
          precision: 0,
          splitNumber: 10,
          axisLine: {
            show: true,
            lineStyle: {
              color: [
                [0.4, "#ff4500"],
                [0.8, "orange"],
                [0.9, "skyblue"],
                [1, "lightgreen"]
              ],
              width: 30
            }
          },
          axisTick: {
            show: true,
            splitNumber: 5,
            length: 8,
            lineStyle: {
              color: "#eee",
              width: 1,
              type: "solid"
            }
          },
          axisLabel: {
            show: true,
            formatter: function(v) {
              switch (v + "") {
                case "40":
                  return "40";
                case "80":
                  return "80";
                case "90":
                  return "90";
                case "100":
                  return "100";
                default:
                  return "";
              }
            },
            textStyle: {
              color: "#333"
            }
          },
          splitLine: {
            show: true,
            length: 30,
            lineStyle: {
              color: "#eee",
              width: 2,
              type: "solid"
            }
          },
          pointer: {
            length: "80%",
            width: 8,
            color: "auto"
          },
          title: {
            show: true,
            offsetCenter: ["-65%", -10],
            textStyle: {
              color: "#333",
              fontSize: 15
            }
          },
          detail: {
            show: true,
            backgroundColor: "rgba(0,0,0,0)",
            borderWidth: 0,
            borderColor: "#ccc",
            width: 100,
            height: 40,
            offsetCenter: ["-60%", 10],
            formatter: "{value}%",
            textStyle: {
              color: "auto",
              fontSize: 30
            }
          },
          data: [{
            value: porcentaje,
            name: "Capital operativo"
          }]
        }]
      });
      var liquidezChart = echarts.init(document.getElementById("liquidez_chart"), theme);
      liquidezChart.setOption({
        tooltip: {
          formatter: "{a} <br/>{b} : {c}%"
        },
        toolbox: {
          show: true,
          feature: {
            restore: {
              show: true,
              title: "Restore"
            },
            saveAsImage: {
              show: true,
              title: "Guardar"
            }
          }
        },
        series: [{
          name: "Liquidez",
          type: "gauge",
          center: ["50%", "50%"],
          startAngle: 140,
          endAngle: -140,
          min: 0,
          max: 100,
          precision: 0,
          splitNumber: 10,
          axisLine: {
            show: true,
            lineStyle: {
              color: [
              
                [0.4, "#ff4500"],
                [0.8, "orange"],
                [0.9, "skyblue"],
                [1, "lightgreen"]
              ],
              width: 30
            }
          },
          axisTick: {
            show: true,
            splitNumber: 5,
            length: 8,
            lineStyle: {
              color: "#eee",
              width: 1,
              type: "solid"
            }
          },
          axisLabel: {
            show: true,
            formatter: function(v) {
              switch (v + "") {
                case "40":
                  return "40";
                case "80":
                  return "80";
                case "90":
                  return "90";
                case "100":
                  return "100";
                default:
                  return "";
              }
            },
            textStyle: {
              color: "#333"
            }
          },
          splitLine: {
            show: true,
            length: 30,
            lineStyle: {
              color: "#eee",
              width: 2,
              type: "solid"
            }
          },
          pointer: {
            length: "80%",
            width: 8,
            color: "auto"
          },
          title: {
            show: true,
            offsetCenter: ["-65%", -10],
            textStyle: {
              color: "#333",
              fontSize: 15
            }
          },
          detail: {
            show: true,
            backgroundColor: "rgba(0,0,0,0)",
            borderWidth: 0,
            borderColor: "#ccc",
            width: 100,
            height: 40,
            offsetCenter: ["-60%", 10],
            formatter: "{value}%",
            textStyle: {
              color: "auto",
              fontSize: 30
            }
          },
          data: [{
            value: porcentajeLiquidez,
            name: "Liquidez"
          }]
        }]
      });




     
    
    
        }
      }
    });
  }
  
  $(document).on("click", ".veDetalle", function (){
     $("body").css("cursor","wait");
      var busqueda = $(this).attr("busqueda");
      var titulo = $(this).html();
      var tipo = $(this).attr("tipo");
      var today = new Date();
      var dd = today.getDate();
      var mm = today.getMonth()+1; //January is 0!
      var yyyy = today.getFullYear();
      if(dd<10) {
          dd='0'+dd
      } 

      if(mm<10) {
          mm='0'+mm
      } 

      today = mm+'/'+dd+'/'+yyyy;
      var PERIOD = yyyy+"0"+mm;
      var param = {tipo : tipo, PERIOD : PERIOD}
      $.ajax({
        url: "/veDetallePrimerNivel",
        dataType : "json",
        type : "post",
        data: param,
        async : true,
        beforeSend : function (){
        },
        complete : function (){
        }, 
        success : function (resp){
          if(resp.success==1)
          {
            var cad = '<div class="">\
          <div class="page-title">\
            <div class="title_left">\
              <h3>Calculo de '+titulo+'</h3>\
            </div>\
          </div>\
          <div class="clearfix"></div>\
          <div class="row">\
              <div class="col-md-12 col-sm-12 col-xs-12">\
                <div class="x_panel">\
                  <div class="x_title">\
                    <h2>'+titulo+' calculado al periodo: <small>'+resp.periodoAnterior+'</small></h2>\
                    <ul class="nav navbar-right panel_toolbox">\
                      <li><a class="collapse-link"><i class="fa fa-chevron-up"></i></a>\
                      </li>\
                      <li class="dropdown">\
                        <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false"><i class="fa fa-wrench"></i></a>\
                        <ul class="dropdown-menu" role="menu">\
                          <li><a href="#">Settings 1</a>\
                          </li>\
                          <li><a href="#">Settings 2</a>\
                          </li>\
                        </ul>\
                      </li>\
                      <li><a class="close-link"><i class="fa fa-close"></i></a>\
                      </li>\
                    </ul>\
                    <div class="clearfix"></div>\
                  </div>\
                  <div class="x_content">\
                    <table class="table table-striped">\
                      <thead>\
                        <tr>\
                          <th>Nombre</th>\
                          <th>Cantidad</th>\
                        </tr>\
                      </thead>\
                      <tbody id="tablaVeDetalles">\
                      </tbody>\
                    </table>\
                  </div>\
                </div>\
              </div>\
              </div></div>';
            $("#mainContent").html(cad);
            var aver2 = Object.keys(resp.lineas);
            cad = '';
            var empiezo = 1;
            while(empiezo<=aver2.length)
            {
              for(i=0;i<aver2.length;i++)
              {
                if(parseInt(resp.lineas[aver2[i]]["Subdetalle"]) ==empiezo)
                {
                  cad=cad+'<tr><td><a href="#" class="vesAlOtroDettale" subdetalle="'+resp.lineas[aver2[i]]["Subdetalle"]+'">'+resp.lineas[aver2[i]]["Nombre"]+'</a></td><td>'+format2(resp.lineas[aver2[i]]["AMOUNT"],"$")+'</td></tr>';
                  empiezo++;
                }
              }  
            }
            
           /* var sumaTotal=0.0;
            var sumaDebitos=0;
            var sumaCreditos=0;
            for(i=0;i<aver2.length;i++)
            {
              if(resp.lineas[aver2[i]]["D_C"]=="D")
              {
                sumaDebitos=sumaDebitos+parseFloat(resp.lineas[aver2[i]]["AMOUNT"]);  
              }
              else
              {
                sumaCreditos=sumaCreditos-parseFloat(resp.lineas[aver2[i]]["AMOUNT"]);
              }
              
              cad=cad+'<tr><td><a href="#" class="vesAlDiario" diario="'+resp.lineas[aver2[i]]["JRNAL_NO"]+'">'+resp.lineas[aver2[i]]["JRNAL_NO"]+'</a></td><td>'+resp.lineas[aver2[i]]["JRNAL_LINE"]+'</td><td>'+format2(resp.lineas[aver2[i]]["AMOUNT"],"$")+'</td><td>'+resp.lineas[aver2[i]]["D_C"]+'</td><td>'+resp.lineas[aver2[i]]["ACNT_CODE"]+'</td><td>'+resp.lineas[aver2[i]]["DESCRIPTN"]+'</td><td>'+resp.lineas[aver2[i]]["PERIOD"]+'</td><td>'+resp.lineas[aver2[i]]["JRNAL_SRCE"]+'</td><td>'+resp.lineas[aver2[i]]["JRNAL_TYPE"]+'</td><td>'+resp.lineas[aver2[i]]["TRANS_DATETIME"]+'</td></tr>';
            }    
            sumaTotal=sumaDebitos+sumaCreditos;
            sumaTotal=sumaTotal*-1;
            cad=cad+'<tr><td colspan="3">Sumas:</td><td colspan="2">D: '+format2(sumaDebitos,"$")+'</td><td colspan="2">C: '+format2(sumaCreditos,"$")+'</td><td colspan="3">Total: '+format2(sumaTotal,"$")+'</td></tr>';
            */
            $("#tablaVeDetalles").html(cad);
             $("body").css("cursor","default");
          }
        }
      });  
    });
  
  $(document).on("click", "#inicio", function (){
      calculaDashboard1();
    });
   
    $(document).ready(function() {
      $.ajaxSetup({ cache: false });
      calculaDashboard1();  
    });
  </script>
  <script>
    NProgress.done();
  </script>
  <!-- /datepicker -->
  <!-- /footer content -->
</body>
</html>