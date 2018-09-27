package router

import (
	"github.com/astaxie/beego"
	"workspace/go-metronic/controller/AppController"
)

func init() {
	beego.Router("/index", &AppController.AppController{}, "*:Index")
}
