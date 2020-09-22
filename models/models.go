package models

import "time"

type ProductInfo struct {
	Name    string   `json:"name"`
	Version []string `json:"version"`
}

type DownloadFiles struct {
	Inventory string `json:"inventory"`
	All       string `json:"all"`
}

type ServiceInfo struct {
	Name     string `json:"name"`
	Tag      string `json:"tag"`
	Replicas int    `json:"replicas"`
	Enabled  bool   `json:"enabled"`
}

type ProductServiceInfo struct {
	ProductName    string         `json:"product_name"`
	ProductVersion string         `json:"product_version"`
	DeployMod      string         `json:"deploy_mod"`
	Services       []*ServiceInfo `json:"services"`
	IpInfos        []*IpInfo      `json:"ip_infos"`
}

type IpInfo struct {
	Ip       string `json:"ip"`
	Password string `json:"password"`
}

// runtime config from web
type ProductRuntimeConfig struct {
	Id             int    `json:"id" xorm:"id notnull pk autoincr INT(11)"`
	Ukey           string `json:"ukey" xorm:"ukey unique varchar(128)"`
	ProductName    string `json:"product_name" xorm:"product_name varchar(64)"`
	ProductVersion string `json:"product_version" xorm:"product_version varchar(64)"`
	ServiceName    string `json:"service_name,omitempty" xorm:"service_name varchar(64)"`
	Replicas       int    `json:"replicas,omitempty" xorm:"INT(11)"`
	Enabled        bool   `json:"enabled,omitempty" `
	Ip             string `json:"ip,omitempty" xorm:"varchar(64)"`
	Password       string `json:"password,omitempty" xorm:"varchar(64)"`
}

// deploy status
type DeployStatus struct {
	Id             int       `json:"id" xorm:"id notnull pk autoincr INT(11)"`
	ProductName    string    `json:"product_name" xorm:"product_name unique varchar(64)"`
	ProductVersion string    `json:"product_version" xorm:"product_version unique varchar(64)"`
	Started        time.Time `json:"started" xorm:"notnull created datetime"`
	Status         string    `json:"status" xorm:"varchar(64)"`
}
