// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"Leafy_api/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/base_plants",
			beego.NSInclude(
				&controllers.PlantTypeController{},
			),
		),
		beego.NSNamespace("/diseases",
			beego.NSInclude(
				&controllers.DiseaseController{},
			),
		),
		beego.NSNamespace("/user_plants",
			beego.NSInclude(
				&controllers.UserPlantController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
