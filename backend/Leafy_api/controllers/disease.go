package controllers

import (
	"Leafy_api/models"

	_ "github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

// Операции с заболеваниями
type DiseaseController struct {
	beego.Controller
}

type AllDiseasesResponseExample struct {
	Err  bool             `json:"Err" example:"false"`
	Data []models.Disease `json:"Data"`
}

// @Title GetAll
// @Description get all Diseases
// @Success 200 {object} controllers.AllDiseasesResponseExample
// @router / [get]
func (ds *DiseaseController) GetAll() {
	dss := models.GetAllDiseases()
	ds.Data["json"] = Response{Err: false, Data: dss}
	ds.ServeJSON()
}

// @Title Get
// @Description get Disease by id
// @Param	id		path 	string	true		"Айди заболевания"
// @Success 200 {object} models.Disease
// @Failure 403 :id is empty
// @router /:id [get]
func (ds *DiseaseController) Get() {
	id, err := ds.GetInt64(":id")
	if err == nil {
		dds, err := models.GetDisease(id)
		if err == nil {
			ds.Data["json"] = Response{Err: false, Data: &dds}
		} else {
			ds.Data["json"] = Response{Err: true, Data: "заболевания с таким id не существует"}
		}
	}
	ds.ServeJSON()
}
