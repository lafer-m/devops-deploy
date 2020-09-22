package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/lafer-m/devops-deploy/controllers"
	"time"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowOrigins:    []string{"*"},
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "UserToken", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		MaxAge:          5 * time.Minute,
	}))
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/inventory", beego.NSInclude(&controllers.Inventory{})),
		beego.NSNamespace("/deploy", beego.NSInclude(&controllers.Deploy{})),
		beego.NSNamespace("/products", beego.NSInclude(&controllers.Product{})),
	)
	beego.AddNamespace(ns)
	beego.Router("/", &controllers.MainController{})
}
