package routers

import (
	"github.com/beego/beego/v2/server/web"
	"webapp/controllers"
)

func init() {

	web.Router("/", &controllers.MainController{})
}
