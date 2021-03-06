// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/Go-Docker-Hackathon/team-ZJU-FT/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/index",
			beego.NSInclude(
				&controllers.HomeController{},
			),
		),
		beego.NSNamespace("/template",
			beego.NSInclude(
				&controllers.TemplateController{},
			),
		),
		beego.NSNamespace("/terminal",
			beego.NSInclude(
				&controllers.TerminalController{},
			),
		),
		beego.NSNamespace("/page",
			beego.NSInclude(
				&controllers.PageController{},
			),
		),
		beego.NSNamespace("/testbuild",
			beego.NSInclude(
				&controllers.BuildController{},
			),
		),
		beego.NSNamespace("/tutorial",
			beego.NSInclude(
				&controllers.TutorialController{},
			),
		),
		beego.NSNamespace("/test",
			beego.NSInclude(
				&controllers.TestController{},
			),
		),
		beego.NSNamespace("/name",
			beego.NSInclude(
				&controllers.TmpController{},
			),
		),
		beego.NSNamespace("/cdf",
			beego.NSInclude(
				&controllers.CdfController{},
			),
		),
		beego.NSNamespace("/term",
			beego.NSInclude(
				&controllers.TermController{},
			),
		),
	)
	beego.AddNamespace(ns)
	//beego.Router("/", &controllers.PageController{})
	//beego.SetStaticPath("/static", "static")
	//beego.SetStaticPath("/views", "views")

}
