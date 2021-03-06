package config

import "github.com/beego/beego/v2/core/logs"

func init() {
	err := logs.SetLogger(logs.AdapterFile, `{"filename":"logs/project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	if err != nil {
		panic(err)
	}
}