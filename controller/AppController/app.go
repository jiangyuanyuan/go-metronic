package AppController

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type AppController struct {
	beego.Controller
}

func (p *AppController) Index() {
	logs.Debug("到了index")
	p.TplName = "index/index.html"
	//m := make(map[string]interface{})
	//m["code"] = 200
	//m["message"] = "success"
	//p.Data["json"] = m
	//p.ServeJSON(true)
}
