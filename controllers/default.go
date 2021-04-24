package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Email"] = "electrosonix12@gmail.com"
	if c.GetSession("id") == nil {
		c.SetSession("id", "electrosonix12@gmail.com")
		c.Data["session"] = "Establecimos una session"
	}
	c.TplName = "home.html"
}
