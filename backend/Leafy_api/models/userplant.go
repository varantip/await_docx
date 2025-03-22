package models

import (
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterModel(new(UserPlant))
}

type UserPlant struct {
	UserPlant_id       int64     `orm:"column(UserPlant_id);pk;auto"`
	User_id            int64     `orm:"column(User_id)"`
	PlantType_id       int64     `orm:"column(PlantType_id)"`
	Disease_id         int64     `orm:"column(Disease_id)"`
	DateOfLastWatering time.Time `orm:"column(DateOfLastWatering);type(date)"`
	DateOfPlanting     time.Time `orm:"column(DateOfPlanting);type(date)"`
	Nickname           string    `orm:"size(256);column(Nickname)"`
}

type ParsedUserPlant struct {
	UserPlant_id       int64  `orm:"column(UserPlant_id);pk;auto"`
	User_id            int64  `orm:"column(User_id)"`
	PlantType_id       int64  `orm:"column(PlantType_id)"`
	Disease_id         int64  `orm:"column(Disease_id)"`
	DateOfLastWatering string `orm:"column(DateOfLastWatering)"`
	DateOfPlanting     string `orm:"column(DateOfPlanting)"`
	Nickname           string `orm:"size(256);column(Nickname)"`
}

func (up *UserPlant) TableName() string {
	return "UserPlant"
}

func AddUserPlant(up UserPlant) int64 {
	o := orm.NewOrmUsingDB("Leafy")
	fmt.Println(up)
	id, err := o.Insert(&up)
	if err != nil {
		fmt.Println(err)
	}
	return id
}

func UpdateUserPlant(up UserPlant) error {
	o := orm.NewOrmUsingDB("Leafy")
	fmt.Println(&up)
	_, err := o.Update(&up, "PlantType_id", "Disease_id", "DateOfLastWatering", "DateOfPlanting", "Nickname")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetAllUsersPlants(id int64) *[]UserPlant {
	var UPlants []UserPlant
	o := orm.NewOrmUsingDB("Leafy")
	qb, _ := orm.NewQueryBuilder("postgres")
	o.Raw(qb.Select("UserPlant_id", "User_id", "PlantType_id", "Disease_id", "DateOfLastWatering", "DateOfPlanting", "Nickname").From("UserPlant").Where(fmt.Sprintf(`"User_id" = %d`, id)).String()).QueryRows(&UPlants)
	return &UPlants
}

func GetAllUserPlants() *[]UserPlant {
	var UPlants []UserPlant
	o := orm.NewOrmUsingDB("Leafy")
	qb, _ := orm.NewQueryBuilder("postgres")
	o.Raw(qb.Select("UserPlant_id", "User_id", "PlantType_id", "Disease_id", "DateOfLastWatering", "DateOfPlanting", "Nickname").From("UserPlant").String()).QueryRows(&UPlants)
	return &UPlants
}
