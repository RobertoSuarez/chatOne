package controllers

import (
	"chatOne/models"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"
)


type LoginController struct {
	web.Controller
}


// TODO: /login debe de resivir un json con las creadenciales,
// y retornar un json con un token para que el js cree la cookies

// @router /login [post]
func (l *LoginController) IniciarSession() {
	db := models.GetDatabase()

	msg := models.NewInfo()
	l.Data["json"] = msg
	// nos aseguramos que se enviara un json
	defer l.ServeJSON()

	// Las credenciales que envia el cliente
	var authDetails models.Authentication

	err := json.Unmarshal(l.Ctx.Input.RequestBody, &authDetails)
	if err != nil {
		msg.SetInfo(nil, true, err.Error())
		return
	}

	// Consultamos a la base de datos el usuario, por medio del email
	// enviado del cliente
	var authUser models.Usuario
	db.Where("email = ?", authDetails.Email).First(&authUser)
	logs.Info("Login authUser: ", authUser)
	if authUser.Email == "" {
		msg.SetInfo(nil, true, "usuario o contraseña incorrecto")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(authUser.Password), []byte(authDetails.Password))
	if err != nil {
		msg.SetInfo(nil, true, "usuario o contraseña incorrecto")
		return
	}

	validToken, err := GenerateJWT(authUser.ID, authUser.Email, authUser.Role)
	if err != nil {
		msg.SetInfo(nil, true, "Fallo al generar el token")
		l.ServeJSON()
		return
	}

	var token models.Token
	token.Email = authUser.Email
	token.Role = authUser.Role
	token.TokenString = validToken

	msg.SetInfo(token, false, "Login exitoso")
}

func GenerateJWT(id uint, email, role string) (string, error) {
	var mySigningKey = []byte(models.Secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iduser"] = id
	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 50).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		logs.Error("Algo salio mal")
		return "", err
	}

	return tokenString, nil
}

// Filtrar usuarios
func FiltrarUsuario(ctx *context.Context) {
	logs.Info("Filtro de usuario ", ctx.Input.URL())
	if strings.Contains(ctx.Input.URL(), "/login") {
		return
	}
	// si no se tiene la cookies id, entonces
	// se redireciona a login, para iniciar session
	id := ctx.GetCookie("id")
	if id == "" {
		ctx.Redirect(302, "/login")
	}

}

// filtrar Consultas con el token
func IsAuthorized(ctx *context.Context) {
	logs.Info("Filtro del token: ", ctx.Input.URL())
	// si se quiere logear se da permiso
	if strings.Contains(ctx.Input.URL(), "/login") {
		logs.Info("Pasamos el filtro de autorización")
		return
	}

	// si quiere crear un nuevo usuario no nesecita token
	if strings.Contains(ctx.Input.URL(), "/users") && ctx.Input.Method() == "POST" {
		return
	}

	if ctx.Request.Header.Get("Token") == "" {
		ctx.WriteString("Error: Falta el token")
		return
	}

	var Myfirma = []byte(models.Secretkey)

	token, err := jwt.Parse(ctx.Request.Header.Get("Token"), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Error al parsin del token")
		}
		return Myfirma, nil
	})

	if err != nil {
		ctx.WriteString("Tu token a expirado")
		return
	}

	if claimns, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// if claimns["role"] == "admin" {
		// 	//ctx.Request.Header.Set("Role", "admin")
		// 	ctx.Input.SetData("Role", "admin")
		// 	ctx.Input.SetData("iduser", claimns["iduser"])
		// 	return
		// }else if claimns["role"] == "user" {
		// 	ctx.Input.SetData("Role", "user")
		// 	ctx.Input.SetData("iduser", claimns["iduser"])
		// 	return
		// }
		ctx.Input.SetData("iduser", claimns["iduser"])
		ctx.Input.SetData("email", claimns["email"])
		ctx.Input.SetData("Role", claimns["role"])
		return
	}

	ctx.WriteString("Not Authorized")
	return

}
