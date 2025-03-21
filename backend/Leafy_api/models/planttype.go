package models

import (
	"errors"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterModel(new(PlantType))
}

type PlantType struct {
	PlantType_id          int64  `orm:"column(PlantType_id);pk;auto"`
	Name                  string `orm:"column(Name)"`
	WateringFrequency     int64  `orm:"column(WateringFrequency)"`
	TemperaturePreference string `orm:"size(256); column(TemperaturePreferece)"`
	LightPreference       string `orm:"size(256); column(LightPreference)"`
	Description           string `orm:"column(Description)"`
	Image_Link            string `orm:"column(Image_Link)"`
	Bio_Name              string `orm:"size(256); column(Bio_Name)"`
}

func (pt *PlantType) TableName() string {
	// db table name
	return "PlantType"
}

func GetPlantType(id int64) (pt *PlantType, err error) {
	o := orm.NewOrmUsingDB("Leafy")
	ptype := PlantType{PlantType_id: id}
	err = o.Read(&ptype)
	fmt.Println(ptype.PlantType_id)
	if err == orm.ErrNoRows {
		return nil, errors.New("в каталоге не существует растения с таким id")
	}
	return &ptype, nil
}

func GetAllPlantTypes() (pts *[]PlantType) {
	o := orm.NewOrmUsingDB("Leafy")
	var ptypes []PlantType
	qb, _ := orm.NewQueryBuilder("postgres")
	o.Raw(qb.Select("PlantType_id", "WateringFrequency", "Name", "TemperaturePreference", "LightPreference", "Description", "Image_Link", "Bio_Name").From("PlantType").Limit(20).String()).QueryRows(&ptypes) //лимит 20, т.к. в MVP у нас будет ограниченное количество типов растений
	return &ptypes
}
