package IndexController

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (p *IndexController) Index() {
	p.TplName = "index/index.html"
}
