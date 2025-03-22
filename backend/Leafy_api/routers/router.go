// @APIVersion 1.0.0
// @Title Leafy API
// @Description API Documentation for Leafy app
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
			beego.NSRouter("/", &controllers.UserController{}, "get:Get"),
			beego.NSRouter("/all", &controllers.UserController{}, "get:GetAll"),
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
			beego.NSRouter("/", &controllers.UserPlantController{}, "get:Get"),
			beego.NSRouter("/all", &controllers.UserPlantController{}, "get:GetAll"),
		),
		beego.NSNamespace("/notifications",
			beego.NSInclude(
				&controllers.NotificationController{},
			),
		),
		beego.NSNamespace("/identify",
			beego.NSInclude(
				&controllers.IdentController{},
			),
		),
	)
	beego.AddNamespace(ns)
	beego.SetStaticPath("/v1/assets", "assets")
	beego.Router("/ws", &controllers.WebSocketController{})
	beego.Router("/notifs", &controllers.WebSocketController{}, "get:ViewNotifs")
}
