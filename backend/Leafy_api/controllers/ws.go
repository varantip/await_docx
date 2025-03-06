package controllers

import (
	"Leafy_api/models"
	"fmt"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

var hub = models.NewHub()

func init() {
	go hub.Run()
}

type WebSocketController struct {
	beego.Controller
}

// @Title Get
// @Description Initiate Websocket connection
// @Success 200
// @router / [get]
func (ws *WebSocketController) Get() {
	if ws.GetSession("accessToken") != nil {
		accessToken := ws.GetSession("accessToken").(string)
		token, _ := models.VerifyToken(accessToken)
		fmt.Println(strconv.FormatInt(int64(token["id"].(float64)), 10))
		models.ServeWs(hub, ws.Ctx.ResponseWriter, ws.Ctx.Request, strconv.FormatInt(int64(token["id"].(float64)), 10))
	}
}

func (ws *WebSocketController) ViewNotifs() {
	if ws.GetSession("accessToken") != nil {
		accessToken := ws.GetSession("accessToken").(string)
		token, _ := models.VerifyToken(accessToken)
		if len(token) > 0 {
			ws.TplName = "home.html"
			ws.Render()
		} else {
			ws.Abort("401")
		}
	} else {
		ws.Abort("401")
	}

}
