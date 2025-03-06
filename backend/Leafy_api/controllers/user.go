package controllers

import (
	"Leafy_api/models"
	"encoding/json"
	"fmt"
	"slices"

	_ "github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// Единая структура ответа на запросы
type Response struct {
	Err  bool `json:"err"`
	Data any  `json:"data"`
}

var SecretKey = []byte("some-key")

// Проверка заголовков для всех запросов
func (u *UserController) HandlerFunc(rules string) bool {
	fmt.Println(u.Ctx.Request)
	fmt.Println(u.Ctx.Request.Header["Authorization"])
	fmt.Println(u.GetSession("accessToken"), "====")
	switch rules {
	case "GetAll", "Logout": // rules - в значении имеет название функции, которая выполняется при вызове метода
		if u.GetSession("accessToken") == nil { // GetSession возвращает nil если нет ключа, поэтому приходится проверять

			break
		}
		accessToken := u.GetSession("accessToken").(string)
		fmt.Println(accessToken, "acctok")
		arrayToken := u.Ctx.Request.Header["Authorization"]
		fmt.Println(arrayToken, "arrtok")
		token, _ := models.VerifyToken(accessToken)                                    // забираем данные из токена
		fmt.Println(token, token["id"])                                                // Можем делать проверку
		if len(arrayToken) > 0 && slices.Contains(arrayToken, "Bearer "+accessToken) { //проверка токена, тут проверка через сессию
			return false
		}
	default: //все не указанные методы будут выполняться без авторизации
		return false
	}
	u.Abort("401") // выдаем ошибку авторизации
	return true
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.User_id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var User models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &User)
	uid := models.AddUser(User)
	u.Data["json"] = Response{Err: false, Data: uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Param Authorization header true "Authorization header. example: Bearer {token} "
// @Success 200 {object} models.User
// @Failure 401 :not authorized
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = Response{Err: false, Data: users}
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The uid"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid, err := u.GetInt64(":uid")
	if err == nil {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = Response{Err: true, Data: err.Error()}
		} else {
			u.Data["json"] = Response{Err: false, Data: user}
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid, err := u.GetInt64(":uid")
	if err == nil {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		err := models.UpdateUser(&user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			uu, _ := models.GetUser(uid)
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid, err := u.GetInt64(":uid")
	if err == nil {
		del := models.DeleteUser(uid)
		if del {
			u.Data["json"] = Response{Err: false, Data: "Пользователь удален"}
		} else {
			u.Data["json"] = Response{Err: true, Data: "Пользователь не найден"}
		}
	}
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	login			query 	string	true		"The login"
// @Param	password		query 	string	true		"The password"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	var user models.User
	user.Login = u.Ctx.Input.Query("login")
	user.Password = u.Ctx.Input.Query("password")
	token := models.Login(user)
	if token != "" {
		u.Data["json"] = Response{Err: false, Data: token}
		// установка значения сессии
		u.SetSession("accessToken", token)
		fmt.Println(u.Ctx.Output)
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Param Authorization header true "Authorization header. example: Bearer {token} "
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	// получение значения сессии
	//u.GetSession("accessToken")
	// удаление значения сессии
	// u.DelSession("accessToken")
	// уничтожение сессии
	u.DestroySession()
	u.Data["json"] = Response{Err: false, Data: "Вышли из сессии"}
	u.ServeJSON()
}
