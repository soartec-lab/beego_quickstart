package routers

import (
	"github.com/soartec-lab/beego_quickstart/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
