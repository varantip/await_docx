package controllers

import (
	"Leafy_api/models"

	_ "github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

// Операции с типами растений
type PlantTypeController struct {
	beego.Controller
}

type AllPlantTypesResponseExample struct {
	Err  bool               `json:"Err" example:"false"`
	Data []models.PlantType `json:"Data"`
}

// @Title GetAll
// @Description get all PlantTypes
// @Success 200 {object} controllers.AllPlantTypesResponseExample
// @router / [get]
func (pt *PlantTypeController) GetAll() {
	ptypes := models.GetAllPlantTypes()
	pt.Data["json"] = Response{Err: false, Data: ptypes}
	pt.ServeJSON()
}

// @Title Get
// @Description get plant_type by id
// @Param	id		path 	string	true		"айди типа растения"
// @Success 200 {object} models.PlantType
// @Failure 403 :id is empty
// @router /:id [get]
func (pt *PlantTypeController) Get() {
	id, err := pt.GetInt64(":id")
	if err == nil {
		ptype, err := models.GetPlantType(id)
		if err == nil {
			pt.Data["json"] = Response{Err: false, Data: &ptype}
		} else {
			pt.Data["json"] = Response{Err: true, Data: "Растения с таким id нет в каталоге"}
		}
	}
	pt.ServeJSON()
}
