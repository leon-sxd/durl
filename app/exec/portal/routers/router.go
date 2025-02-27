package routers

import (
	"html/template"
	"net/http"

	"durl/app/exec/portal/controllers"

	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/server/web"
)

// RouterHandler 路由跳转
func RouterHandler() {

	// 初始化短链池
	controllers.InitCon()

	// 首页
	web.Router("/", &controllers.Controller{}, "get:Index")

	// 获取xsrfToken
	web.Router("/xsrf-token", &controllers.Controller{}, "get:GetXsrfToken")

	// 设置短链
	web.Router("/shortUrl", &controllers.Controller{}, "post:SetShortUrl")

	web.ErrorHandler("404", pageNotFound)

}

// 定义404页面
func pageNotFound(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.html").ParseFiles(web.BConfig.WebConfig.ViewsPath + "/404.html")
	data := make(map[string]interface{})
	data["content"] = "page not found"
	// 百度统计key
	sConf, _ := config.String("Baidu")
	if sConf != "" {
		data["Statistical_Baidu_Key"] = sConf
	}
	_ = t.Execute(rw, data)
}
