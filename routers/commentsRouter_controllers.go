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
            Router: "/:iduser/contacts",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatOne/controllers:ContactController"] = append(beego.GlobalControllerRouter["chatOne/controllers:ContactController"],
        beego.ControllerComments{
            Method: "CreateContact",
            Router: "/:iduser/contacts",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatOne/controllers:ContactController"] = append(beego.GlobalControllerRouter["chatOne/controllers:ContactController"],
        beego.ControllerComments{
            Method: "OneContact",
            Router: "/:iduser/contacts/:idcontact",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatOne/controllers:ContactController"] = append(beego.GlobalControllerRouter["chatOne/controllers:ContactController"],
        beego.ControllerComments{
            Method: "DeleteContact",
            Router: "/:iduser/contacts/:idcontact",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatOne/controllers:ContactController"] = append(beego.GlobalControllerRouter["chatOne/controllers:ContactController"],
        beego.ControllerComments{
            Method: "UpdateContact",
            Router: "/:iduser/contacts/:idcontact",
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

    beego.GlobalControllerRouter["chatOne/controllers:MessageController"] = append(beego.GlobalControllerRouter["chatOne/controllers:MessageController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/:iduser/message",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatOne/controllers:MessageController"] = append(beego.GlobalControllerRouter["chatOne/controllers:MessageController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/:iduser/message",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatOne/controllers:MessageController"] = append(beego.GlobalControllerRouter["chatOne/controllers:MessageController"],
        beego.ControllerComments{
            Method: "OneMessage",
            Router: "/:iduser/message/:idmessage",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatOne/controllers:UserController"] = append(beego.GlobalControllerRouter["chatOne/controllers:UserController"],
        beego.ControllerComments{
            Method: "CreateUser",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatOne/controllers:UserController"] = append(beego.GlobalControllerRouter["chatOne/controllers:UserController"],
        beego.ControllerComments{
            Method: "AllUser",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatOne/controllers:UserController"] = append(beego.GlobalControllerRouter["chatOne/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserOne",
            Router: "/:iduser",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatOne/controllers:UserController"] = append(beego.GlobalControllerRouter["chatOne/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateUser",
            Router: "/:iduser",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatOne/controllers:UserController"] = append(beego.GlobalControllerRouter["chatOne/controllers:UserController"],
        beego.ControllerComments{
            Method: "DeleteUser",
            Router: "/:iduser",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
