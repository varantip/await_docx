package controllers

import (
	"Leafy_api/models"

	_ "github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about User's plants(Unfinished)
type UserPlantController struct {
	beego.Controller
}

// func (up *UserPlantController) HandlerFunc(rules string) bool {
// 	fmt.Println("bruhh")
// 	fmt.Println(up.GetSession("accessToken"))
// 	switch rules {
// 	default:
// 		up.Abort("401")
// 		return true
// 	}
// }

// @Title Get
// @Description get UserPlant by User_id
// @Param	uid		path 	string	true		"user id"
// @Success 200 {object} models.Disease
// @Failure 403 :uid is empty
// @router /:uid [get]
func (upt *UserPlantController) Get() {
	id, _ := upt.GetInt64(":id")
	UPlants := models.GetAllUsersPlants(id)
	upt.Data["json"] = Response{Err: false, Data: UPlants}
	upt.ServeJSON()
}
