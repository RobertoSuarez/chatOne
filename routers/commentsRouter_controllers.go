package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["chatOne/controllers:ChatController"] = append(beego.GlobalControllerRouter["chatOne/controllers:ChatController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/websocket",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatOne/controllers:ContactController"] = append(beego.GlobalControllerRouter["chatOne/controllers:ContactController"],
        beego.ControllerComments{
            Method: "AllContacts",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatOne/controllers:ContactController"] = append(beego.GlobalControllerRouter["chatOne/controllers:ContactController"],
        beego.ControllerComments{
            Method: "CreateContact",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatOne/controllers:ContactController"] = append(beego.GlobalControllerRouter["chatOne/controllers:ContactController"],
        beego.ControllerComments{
            Method: "OneContact",
            Router: "/:idcontact",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatOne/controllers:ContactController"] = append(beego.GlobalControllerRouter["chatOne/controllers:ContactController"],
        beego.ControllerComments{
            Method: "DeleteContact",
            Router: "/:idcontact",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatOne/controllers:ContactController"] = append(beego.GlobalControllerRouter["chatOne/controllers:ContactController"],
        beego.ControllerComments{
            Method: "UpdateContact",
            Router: "/:idcontact",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatOne/controllers:LoginController"] = append(beego.GlobalControllerRouter["chatOne/controllers:LoginController"],
        beego.ControllerComments{
            Method: "IniciarSession",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatOne/controllers:UserController"] = append(beego.GlobalControllerRouter["chatOne/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserOne",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatOne/controllers:UserController"] = append(beego.GlobalControllerRouter["chatOne/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateUser",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatOne/controllers:UserController"] = append(beego.GlobalControllerRouter["chatOne/controllers:UserController"],
        beego.ControllerComments{
            Method: "DeleteUser",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
