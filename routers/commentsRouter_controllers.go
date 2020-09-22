package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/lafer-m/devops-deploy/controllers:Deploy"] = append(beego.GlobalControllerRouter["github.com/lafer-m/devops-deploy/controllers:Deploy"],
		beego.ControllerComments{
			Method:           "Deploy",
			Router:           "/deploy",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/lafer-m/devops-deploy/controllers:Inventory"] = append(beego.GlobalControllerRouter["github.com/lafer-m/devops-deploy/controllers:Inventory"],
		beego.ControllerComments{
			Method:           "List",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/lafer-m/devops-deploy/controllers:Inventory"] = append(beego.GlobalControllerRouter["github.com/lafer-m/devops-deploy/controllers:Inventory"],
		beego.ControllerComments{
			Method:           "Generate",
			Router:           "/generate",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/lafer-m/devops-deploy/controllers:Product"] = append(beego.GlobalControllerRouter["github.com/lafer-m/devops-deploy/controllers:Product"],
		beego.ControllerComments{
			Method:           "List",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/lafer-m/devops-deploy/controllers:Product"] = append(beego.GlobalControllerRouter["github.com/lafer-m/devops-deploy/controllers:Product"],
		beego.ControllerComments{
			Method:           "Service",
			Router:           "/:productName/:productVersion/service",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
