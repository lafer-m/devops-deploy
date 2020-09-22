package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"

	"os"
	"path/filepath"
)

var ProductsCfg = &Config{}

type Config struct {
	Products    map[string]Versions
	DeployPath  string `yaml:"deploy_path"`
	PackagePath string `yaml:"package_path"`
}

type Versions struct {
	Name     string
	Versions map[string]map[string]int
}

func GetCfgFile(configFile string) string {
	WorkPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cfg := filepath.Join(WorkPath, configFile)
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(cfg); err != nil {
		if os.IsNotExist(err) {
			cfg = filepath.Join(dir, configFile)
		}
	}

	return cfg
}

func GetDeployMod(product string) string {
	for pro, info := range ProductsCfg.Products {
		if pro == product {
			return info.Name
		}
	}
	return ""
}

// load global products configs
func LoadCfg() {
	products := GetCfgFile("conf/products.yml")

	fl, err := ioutil.ReadFile(products)
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(fl, ProductsCfg); err != nil {
		panic(err)
	}
}

func init() {
	LoadCfg()
}
