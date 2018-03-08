package service

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type configStruct struct {
	Db struct {
		Driver string `yaml:"driver"`
		Dsn    string `yaml:"dsn"`
	}

	Hosts []string `yaml:"hosts"`
}

var Config configStruct

func init() {
	parseConfig()
}

func parseConfig() {

	data, ferr := ioutil.ReadFile("conf/conf.yaml")
	if ferr != nil {
		panic(ferr)
	}

	perr := yaml.Unmarshal([]byte(data), &Config)
	if perr != nil {
		panic(perr)
	}
	return
}
