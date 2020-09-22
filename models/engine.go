package models

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-xorm/xorm"
	"github.com/lafer-m/devops-deploy/conf"
	_ "github.com/mattn/go-sqlite3"
	"github.com/prometheus/common/log"
	"hash/fnv"
)

// use sqlite as store
const sqlFile = "store.db"

var Engine *xorm.Engine

func initTables() {
	if err := Engine.CreateTables(&ProductRuntimeConfig{}, &DeployStatus{}); err != nil {
		panic(err)
	}
}

func initProductsRuntimeInfos() {
	products := conf.ProductsCfg
	log.Info("sqlite init", products)
	for product, versions := range products.Products {
		log.Info(product, versions.Name)
		for version, svcs := range versions.Versions {
			log.Info(version)
			for service, replicas := range svcs {
				var enabled bool = false
				if replicas > 0 {
					enabled = true
				}
				//this value of ukey must be uniq, exclude password
				ukey := GetUkey(product, version, service, "")

				has, err := Engine.Where("ukey = ?", ukey).Get(&ProductRuntimeConfig{})
				if err != nil {
					panic(err)
				}
				if !has {
					info := &ProductRuntimeConfig{
						ProductName:    product,
						ProductVersion: version,
						ServiceName:    service,
						Replicas:       replicas,
						Enabled:        enabled,
						Ip:             "",
						Password:       "",
						Ukey:           ukey,
					}
					Engine.Insert(info)
				}

			}
		}
	}
}

func HasAlreadyExist() bool {

	return false
}

func newEngine() {
	engine, err := xorm.NewEngine("sqlite3", sqlFile)
	if err != nil {
		panic(err)
	}
	Engine = engine
	initTables()
	initProductsRuntimeInfos()
}

func init() {
	newEngine()
}

// HashObject writes the specified object to a hash using the spew library
// which follows pointers and prints actual values of the nested objects
// ensuring the hash does not change when a pointer changes.
// The returned hash can be used for object comparisons.
//
// This is inspired by controller revisions in StatefulSets:
// https://github.com/kubernetes/kubernetes/blob/8de1569ddae62e8fab559fe6bd210a5d6100a277/pkg/controller/history/controller_history.go#L89-L101
func HashObject(object interface{}) string {
	hf := fnv.New32()
	printer := spew.ConfigState{
		Indent:         " ",
		SortKeys:       true,
		DisableMethods: true,
		SpewKeys:       true,
	}
	_, _ = printer.Fprintf(hf, "%#v", object)
	return fmt.Sprint(hf.Sum32())
}

func GetUkey(product, version, service, ip string) string {
	return HashObject(fmt.Sprintf("%s%s%s%s", product, version, service, ip))
}
