package main

import (
	_ "chatOne/routers"
	"os"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	// beego.BConfig.WebConfig.Session.SessionOn = true
	// beego.BConfig.WebConfig.Session.SessionProvider = "file"
	// beego.BConfig.WebConfig.Session.SessionProviderConfig = "tmp/"

	port := os.Getenv("PORT")
	portint, err := strconv.Atoi(port)
	if err == nil {
		beego.BConfig.Listen.HTTPPort = portint
		beego.BConfig.Listen.HTTPSPort = portint
	}
	beego.Run()
}

