package routers

import (
	"durl/app/exec/openapi/controllers"

	"github.com/beego/beego/v2/server/web"
)

// RouterHandler 路由跳转
func RouterHandler() {

	// openapi初始化
	controllers.InitCon()

	// 设置短链
	web.Router("/shortUrl", &controllers.Controller{}, "post:SetShortUrl")

	// 修改短链
	web.Router("/shortUrl", &controllers.Controller{}, "put:UpdateShortUrl")

	// 删除短链
	web.Router("/shortUrl", &controllers.Controller{}, "delete:DelShortKey")

}
