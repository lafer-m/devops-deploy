package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/prometheus/common/log"
	"io/ioutil"
	"net/http"
)

type Inventory struct {
	beego.Controller
}

// @Title Generate Inventory
// @Description generate inventory
// @Param   info   body   struct  true   "generate all.yml and inventory.yml content"
// @Success 201 {code}
// @Failure 403 body is empty
// @router /generate [post]
func (i *Inventory) Generate() {
	defer i.ServeJSON()
	r := &struct {
		ProductName string `json:"product_name"`
		Version     string
	}{}
	body, _ := ioutil.ReadAll(i.Ctx.Request.Body)
	if err := json.Unmarshal(body, r); err != nil {
		log.Error(err)
		i.Ctx.Output.Status = http.StatusInternalServerError
		i.Data["json"] = map[string]interface{}{"code": 1, "msg": err.Error()}
		return
	}

	infos, err := GetProductServiceInfos(r.ProductName, r.Version)
	if err != nil {
		log.Error(err)
		i.Ctx.Output.Status = http.StatusInternalServerError
		i.Data["json"] = map[string]interface{}{"code": 1, "msg": err.Error()}
		return
	}

	log.Info(r, infos)

}

// @Title List Inventory
// @Description list inventory
// @Success 201 {int}
// @Failure 403 body is empty
// @router / [get]
func (i *Inventory) List() {
	defer i.ServeJSON()

}
