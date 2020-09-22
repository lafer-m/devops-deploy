package controllers

import "github.com/astaxie/beego"

type Deploy struct {
	beego.Controller
}

// @router /deploy [post]
func (d *Deploy) Deploy() {

}
