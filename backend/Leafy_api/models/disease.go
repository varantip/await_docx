package models

import (
	"errors"

	"github.com/beego/beego/v2/client/orm"
)

type Disease struct {
	Disease_id int64  `orm:"column(Disease_id);pk;auto"`
	Name       string `orm:"column(Name)"`
	Treatment  string `orm:"column(Treatment)"`
}

func (d *Disease) TableName() string {
	return "Disease"
}
func init() {
	orm.RegisterModel(new(Disease))
}

// GetDisease retrieves Disease by id. Returns error if
// id doesn't exist
func GetDisease(id int64) (d *Disease, err error) {
	o := orm.NewOrmUsingDB("Leafy")
	disease := Disease{Disease_id: id}
	err = o.Read(&disease)
	if err == nil {
		return &disease, nil
	}
	return nil, errors.New("заболевания с таким id не существует")

}

// GetAllDiseases retrieves all Diseases. Returns empty list if
// no records exist
func GetAllDiseases() (dss *[]Disease) {
	o := orm.NewOrmUsingDB("Leafy")
	var diseases []Disease
	qb, _ := orm.NewQueryBuilder("postgres")
	o.Raw(qb.Select("Disease_id", "Name", "Treatment").From("Disease").String()).QueryRows(&diseases)
	return &diseases
}
