package controllers

import (
	"github.com/astaxie/beego"
	"github.com/lafer-m/devops-deploy/conf"
	"github.com/lafer-m/devops-deploy/models"
	"github.com/prometheus/common/log"
)

type Product struct {
	beego.Controller
}

// @Title List product infos
// @Description list infos
// @Success 201 {int}
// @Failure 403 body is empty
// @router /  [get]
func (p *Product) List() {
	defer p.ServeJSON()

	result := []*models.ProductInfo{}

	products := conf.ProductsCfg

	for product, versions := range products.Products {
		var pro *models.ProductInfo
		for version := range versions.Versions {
			if pro == nil {
				pro = &models.ProductInfo{
					Name:    product,
					Version: []string{},
				}
			}
			pro.Version = append(pro.Version, version)
		}
		if pro != nil {
			result = append(result, pro)
		}
	}

	log.Info(result)
	p.Data["json"] = map[string][]*models.ProductInfo{
		"products": result,
	}
}

// @Title List product services info
// @Description list services infos
// @Param   productName        path    string  true        "the productName and version 's svc you want to get"
// @Param   productVersion        path    string  true        "the productName and version 's svc you want to get"
// @Success 201 {[]models.ServiceInfo}
// @Failure 403 body is empty
// @router /:productName/:productVersion/service  [get]
func (p *Product) Service() {
	defer p.ServeJSON()
	result := []*models.ServiceInfo{}
	products := conf.ProductsCfg.Products
	pro, ver := p.GetString(":productName"), p.GetString(":productVersion")
	for product, versions := range products {
		if product != pro {
			continue
		}
		for version, svcs := range versions.Versions {
			if version != ver {
				continue
			}
			for svc, replicas := range svcs {
				var enabled bool = false
				if replicas > 0 {
					enabled = true
				}
				result = append(result, &models.ServiceInfo{
					Name:     svc,
					Tag:      svc,
					Replicas: replicas,
					Enabled:  enabled,
				})
			}
		}
	}

	p.Data["json"] = result
}
