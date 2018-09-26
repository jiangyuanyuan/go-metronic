package router

import (
	"github.com/astaxie/beego"
	"workspace/go-metronic/controller/IndexController"
)

func init() {
	beego.Router("/index", &IndexController.IndexController{}, "*:Index")
}
