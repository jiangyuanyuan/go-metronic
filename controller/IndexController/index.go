package IndexController

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type IndexController struct {
	beego.Controller
}

func (p *IndexController) Index() {
	logs.Debug("到了index")
	p.Layout = "layout/layout.html"
	p.TplName = "index/index.html"
	//m := make(map[string]interface{})
	//m["code"] = 200
	//m["message"] = "success"
	////p.Data["Result"] = 200
	////p.Data["Message"] = "success"
	//p.Data["json"] = m
	//p.ServeJSON(true)
}
