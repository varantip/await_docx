package models

import (
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
)

type Notification struct { //потребовалось создать новую табличку в бд, но она скорее всего не пойдёт в финальный проект
	Notif_id int64     `orm:"column(Notif_id);pk;auto"`
	User_id  int64     `orm:"column(User_id)"`
	Data     string    `orm:"size:(256);column(Data)"`
	Time     time.Time `orm:"column(Time);type(timestamp with time zone)"`
}

type ParsedNotification struct {
	Notif_id int64
	User_id  int64
	Data     string
	Time     string
}

func (n *Notification) TableName() string {
	return "Notification"
}

func init() {
	orm.RegisterModel(new(Notification))
}

func GetAllNotifsById(id int64) (Notifs []Notification) {
	o := orm.NewOrmUsingDB("Leafy")
	qb, _ := orm.NewQueryBuilder("postgres")
	o.Raw(qb.Select("Notif_id", "User_id", "Data", "Time").From("Notification").Where(fmt.Sprintf(`"User_id" = %d`, id)).String()).QueryRows(&Notifs)
	return Notifs
}

func GetAllExpiredNotifsById(id int64) *[]Notification {
	var Notifs []Notification
	o := orm.NewOrmUsingDB("Leafy")
	qb, _ := orm.NewQueryBuilder("postgres")
	o.Raw(qb.Select("Notif_id", "User_id", "Data", "Time").From("Notification").Where(fmt.Sprintf(`"Time" <= 'NOW' AND "User_id" = %d`, id)).String()).QueryRows(&Notifs)
	return &Notifs
}

func DeleteAllExpiredNotifsById(id int64) {
	o := orm.NewOrmUsingDB("Leafy")
	toDel := GetAllExpiredNotifsById(id)
	for _, elem := range *toDel {
		o.Delete(&elem)
	}
}

func AddNotif(n Notification) (id int64, err error) {
	o := orm.NewOrmUsingDB("Leafy")
	id, err = o.Insert(&n)
	return id, err
}
