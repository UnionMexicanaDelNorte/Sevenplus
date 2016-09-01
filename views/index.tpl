<!DOCTYPE html>
<html>
  
<head>
  <title>Sevenplus</title>
    <meta charset="utf-8">

    <!--link href="css/style.css" rel='stylesheet' type='text/css' /-->

    <style>

html,body,div,span,applet,object,iframe,h1,h2,h3,h4,h5,h6,p,blockquote,pre,a,abbr,acronym,address,big,cite,code,del,dfn,em,img,ins,kbd,q,s,samp,small,strike,strong,sub,sup,tt,var,b,u,i,dl,dt,dd,ol,nav ul,nav li,fieldset,form,label,legend,table,caption,tbody,tfoot,thead,tr,th,td,article,aside,canvas,details,embed,figure,figcaption,footer,header,hgroup,menu,nav,output,ruby,section,summary,time,mark,audio,video{margin:0;padding:0;border:0;font-size:100%;font:inherit;vertical-align:baseline;}
article, aside, details, figcaption, figure,footer, header, hgroup, menu, nav, section {display: block;}
ol,ul{list-style:none;margin:0px;padding:0px;}
blockquote,q{quotes:none;}
blockquote:before,blockquote:after,q:before,q:after{content:'';content:none;}
table{border-collapse:collapse;border-spacing:0;}
/* start editing from here */
a{text-decoration:none;}
.txt-rt{text-align:right;}/* text align right */
.txt-lt{text-align:left;}/* text align left */
.txt-center{text-align:center;}/* text align center */
.float-rt{float:right;}/* float right */
.float-lt{float:left;}/* float left */
.clear{clear:both;}/* clear float */
.pos-relative{position:relative;}/* Position Relative */
.pos-absolute{position:absolute;}/* Position Absolute */
.vertical-base{ vertical-align:baseline;}/* vertical align baseline */
.vertical-top{  vertical-align:top;}/* vertical align top */
nav.vertical ul li{ display:block;}/* vertical menu */
nav.horizontal ul li{ display: inline-block;}/* horizontal menu */
img{max-width:100%;}
/*end reset
background:url(../images/bg.jpg)  no-repeat center fixed;
*/
/****-----start-body----****/
body{
  background:url(static/images/bg.jpg)  no-repeat center fixed;
  -webkit-background-size: cover;
  -moz-background-size: cover;
  -o-background-size: cover;
  background-size: cover;
  position:relative;
  font-family: 'Open Sans', sans-serif;

}
.login-form {
  background: url(static/images/bgf.png) no-repeat 100% 0%;
  background-size: 100%;
  width:37%;
  margin:10% auto 4% auto;
  position: relative;
}
.login-form h1{
  text-align:left;
  color: #fff;
  font-size:1.8em;
  font-weight:500;
  font-family: 'Open Sans', sans-serif;
  position:absolute;
  top:7%;
  left:11%;
  margin:34px 0;
} 
.login-form h2 a{
  text-align:right;
  color: #fff;
  font-size:1.15em;
  font-weight:500;
  font-family: 'Open Sans', sans-serif;
  position:absolute;
  top:9.5%;
  right:11%;
  margin:34px 0;
} 
form {
  padding:32% 3em;
}
form li{
  border:1px solid #FFC600;
  list-style:none;
  margin-bottom:16px;
  width:100%;
  background: #FFC600;

}

.icon{
  background:url(static/images/icons.png)  no-repeat 0px 0px;
  height:40px;
  width:40px;
  display: block;
  float:left;
  margin:5px -11px 0px 0px
}

.user{
  background: url(static/images/icons.png) no-repeat -173px 5px;
  
}
.lock{
  background: url(static/images/icons.png) no-repeat -173px 5px;
}
.arrow{
  background: url(static/images/icons.png) no-repeat -200px 5px;
  position: absolute;
}
input[type="text"], input[type="password"] {
  font-family: 'Open Sans', sans-serif;
  width:60%;
  padding: 0.5em 2em 0.5em 1em;
  color: #fff;
  font-size:20px;
  outline: none;
  background: none;
  border:none;
}

input[type="submit"]{
  float:left;
  font-size:18px;
  display: inline-block;
  font-weight:600;
  color: #615F5F;
  transition: 0.1s all;
  -webkit-transition: 0.1s all;
  -moz-transition: 0.1s all;
  -o-transition: 0.1s all;
  cursor: pointer;
  outline: none;
  padding:15px 10px;
  margin-top:3px;
  font-family: 'Open Sans', sans-serif;
  background: #222;
  width:45%;
  border:1px solid #222;
}
input[type="submit"]:hover{
  background: #FFC600;
  border:1px solid #FFC600;
}
.forgot{
  margin:30px 7px;
}
.forgot h3 a{
  color:#fff;
  font-size:15px;
  font-family: 'Open Sans', sans-serif; 
  float:right;
  font-weight:600;
  margin-top:19px;
}
.forgot h3 a:hover{
  color:#FFC600;
}
/****************/
.copy-right {
  position: absolute;
  bottom:-16%;
  left: 43.7%;
}
.copy-right p {
  color: #fff;
  font-size: 1em;
  font-family: 'Open Sans', sans-serif; 
  font-weight: 600;
}
.copy-right p a {
  font-family: 'Open Sans', sans-serif; 
  font-size: 1em;
  color:#170829;
  -webkit-transition: all 0.3s ease-out;
  -moz-transition: all 0.3s ease-out;
  -ms-transition: all 0.3s ease-out;
  -o-transition: all 0.3s ease-out;
  transition: all 0.3s ease-out;
}
.copy-right p a:hover {
  color:#fff;
}
/*-----start-responsive-design------*/
@media (max-width:1024px){
.login-form  {
    margin:8% auto 0;
    width:45%;
  }
  .copy-right {
  left: 41%;
  bottom:-18%;
  }
}
@media (max-width:768px){
  .login-form  {
    margin:12% auto 0;
    width:57%;
  }
  .copy-right {
  left:38%;
  bottom:-14%;
  }
}
@media (max-width:640px){                                  
  .login-form  {
    margin:13% auto 0;
    width:70%;
  }
  .copy-right {
  left:36%;
  bottom:-18%;
  }
}
@media (max-width:480px){                                  
  .login-form  {
    margin:20% auto 0;
    width:90%;
  }
  .copy-right {
  left:30%;
  bottom:-17%;
  }
  
}
@media (max-width:320px){                                  
  .login-form  {
    margin:20% auto 0;
    width:100%;
  }
  .icon{
    margin:2px -11px 0px 0px;
  }
  form{
    padding:38% 2.5em; 
  }
  .login-form h1{
    font-size:1.7em;
    left:35%;
    top:0%;
  }
  .login-form h2{
    font-size:1em;
    right:29%;
    top:10%;
  }
  input[type="text"], input[type="password"]{
    font-size:18px;
    padding: 0.4em 1em 0.45em 1em;15px 2px
  }
  .copy-right {
  left:20%;
  bottom:9%;
  }
  .forgot{
    text-align:center;
    margin:5px 2px;
  }
  .forgot h3 a{
    float:none;
    font-size:14px;
  }
  input[type="submit"]{
    font-size:15px;
    padding: 10px 10px;
    width:56%;
    float:none;
    margin-top: 5px;
  }
  .arrow {
    background: url(../images/icons.png) no-repeat -62px -6px;
    left: 29%;
    top: 62.99%;
  }

}

</style>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <script type="application/x-javascript"> addEventListener("load", function() { setTimeout(hideURLbar, 0); }, false); function hideURLbar(){ window.scrollTo(0,1); } </script>
    <link rel="shortcut icon" href="static/favicon.ico">
    <!--webfonts-->
    <link href='http://fonts.googleapis.com/css?family=Open+Sans:600italic,400,300,600,700' rel='stylesheet' type='text/css'>
    <!--//webfonts-->
</head>
<body>
        <div class="login-form">
            <h1> {{.mensaje}}</h1>
            <h2><a href="#"></a></h2>
        <form method="POST" action="/app" onsubmit="DoSubmit();">
        <input type="hidden" name="servicio"  value="login"/>
          <li>
            <input type="text" class="text" id="correo" name="correo" value="correo@dominio.org" onfocus="this.value = '';" onblur="if (this.value == '') {this.value = 'User Name';}" ><a href="#" class=" icon user"></a>
          </li>
          <li>
            <input type="password" id="password" name="password" value="Password" onfocus="this.value = '';" onblur="if (this.value == '') {this.value = 'Password';}"><a href="#" class=" icon lock"></a>
          </li>
          
           <div class ="forgot">
            <h3><a href="#">&iquest;Se le olvido de su password?</a></h3>
            <input type="submit" id="botonEntrar" value="Entrar" > <a href="#" class=" icon arrow"></a>                                                                                                                                                                                                                                 </h4>
          </div>
        </form>
      </div>
      <script src="static/js/sha512.min.js"></script>
      <script type="text/javascript">
       function DoSubmit() {
            var pass = document.getElementById("password").value;
            pass = sha512(pass);
            document.getElementById("password").value = pass;
            return true;
        }

      </script>
  
          
</body>
</html>