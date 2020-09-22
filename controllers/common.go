package controllers

import (
	"github.com/prometheus/common/log"
	"github.com/lafer-m/devops-deploy/conf"
	"github.com/lafer-m/devops-deploy/models"
)

// get the infos from sqlite db
func GetProductServiceInfos(product, version string) (*models.ProductServiceInfo, error) {
	result := &models.ProductServiceInfo{
		ProductName:    product,
		ProductVersion: version,
		DeployMod:      conf.GetDeployMod(product),
		Services:       []*models.ServiceInfo{},
		IpInfos:        []*models.IpInfo{},
	}
	engine := models.Engine

	var runtimeInfos []*models.ProductRuntimeConfig

	if err := engine.Table("product_runtime_config").Where("product_name = ? and product_version = ?", product, version).Find(&runtimeInfos)
		err != nil {
		return result, err
	}

	for _, info := range runtimeInfos {
		log.Info(info)
		if info.Ip == "" && info.ServiceName != "" {
			var enabled bool = false
			if info.Replicas > 0 {
				enabled = true
			}
			result.Services = append(result.Services, &models.ServiceInfo{
				Name:     info.ServiceName,
				Tag:      info.ServiceName,
				Replicas: info.Replicas,
				Enabled:  enabled,
			})
		}
		if info.Ip != "" && info.ServiceName == "" {
			result.IpInfos = append(result.IpInfos, &models.IpInfo{
				Ip:       info.Ip,
				Password: info.Password,
			})
		}
	}

	return result, nil
}

func ParseServiceName() {}


// generate inventory for ansible deploy
func GenerateInventory(product *models.ProductServiceInfo) {

}

// generate all.yml for ansible deploy
func GenerateAll(product *models.ProductServiceInfo) {}
