package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/config/yaml"
	"github.com/lafer-m/devops-deploy/conf"
	_ "github.com/lafer-m/devops-deploy/models"
	_ "github.com/lafer-m/devops-deploy/routers"
	"github.com/prometheus/common/log"
)

func main() {
	log.Info("Products configs ", conf.ProductsCfg)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.SetViewsPath("web-src/dist")
	beego.SetStaticPath("static", "web-src/dist/static")
	beego.Run()
}
