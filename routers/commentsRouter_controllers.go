package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["chatOne/controllers:ChatController"] = append(beego.GlobalControllerRouter["chatOne/controllers:ChatController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["chatOne/controllers:ChatController"] = append(beego.GlobalControllerRouter["chatOne/controllers:ChatController"],
        beego.ControllerComments{
            Method: "AllUser",
            Router: "/users",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
