package controllers

import (
	"Leafy_api/models"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Notifications (as of now, only used in WebSockets, probably wont go in the final app)
type NotificationController struct {
	beego.Controller
}

// @Title Get
// @Description Gather Notifs by Id
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Response
// @router /:uid [get]
func (nc *NotificationController) Get() {
	id, err := nc.GetInt64(":uid")
	if err != nil {
		nc.Data["json"] = Response{Err: true, Data: "такого uid нет"}
	} else {
		Notifs := models.GetAllNotifsById(id)
		nc.Data["json"] = Response{Err: false, Data: Notifs}
	}
	nc.ServeJSON()
}

// @Title Delete
// @Description Delete all expired notifications
// @Param	uid		path 	string	true		"The user_id for which you check and delete notifications"
// @Success 200  {string} success
// @router /:uid [delete]
func (nc *NotificationController) Delete() {
	id, err := nc.GetInt64(":uid")
	fmt.Println(id)
	if err != nil {
		nc.Data["json"] = Response{Err: true, Data: "такого uid нет"}
	} else {
		models.DeleteAllExpiredNotifsById(id)
		nc.Data["json"] = Response{Err: false, Data: "success"}
	}
	nc.ServeJSON()
}

const layout = "2006-01-02T15:04:05-0700"

// @Title Post
// @Description Add a Notification
// @Param body body models.Notification true "The notfication, date specified like so: YYYY-MM-DDThh:mm:sstz (for kaliningrad tz = +0200)"
// @Success 200 {object} int64
// @router / [post]
func (nc *NotificationController) Post() {
	var parsedNotif models.ParsedNotification
	json.Unmarshal(nc.Ctx.Input.RequestBody, &parsedNotif)
	parsedTime, err := time.Parse(layout, parsedNotif.Time)
	fmt.Println(parsedTime)
	if err == nil {
		Notif := models.Notification{Notif_id: parsedNotif.Notif_id, User_id: parsedNotif.User_id, Data: parsedNotif.Data, Time: parsedTime}
		id, err := models.AddNotif(Notif)
		if err != nil {
			nc.Data["json"] = Response{Err: true, Data: err}
		} else {
			nc.Data["json"] = Response{Err: false, Data: id}
		}
	} else {
		nc.Data["json"] = Response{Err: true, Data: err}
	}

	nc.ServeJSON()
}
