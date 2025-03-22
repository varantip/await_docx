package controllers

import (
	"Leafy_api/models"
	"encoding/json"
	"fmt"

	_ "github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// Единая структура ответа на запросы
type Response struct {
	Err  bool `json:"Err"`
	Data any  `json:"Data"`
}

type AllUsersResponseExample struct {
	Err  bool          `json:"Err" example:"false"`
	Data []models.User `json:"Data"`
}

type UserPostResponseExample struct {
	Err  bool `json:"Err" example:"false"`
	Data int  `json:"Data" example:"1"`
}

type LoginResponseExample struct {
	Err  bool   `json:"Err" example:"false"`
	Data string `json:"Data" example:"ABCDEFGH123456789.ADDBSJADSAJDN0123032.ASDNINASIDAID31213"`
}

type LogoutResponseExample struct {
	Err  bool   `json:"Err" example:"false"`
	Data string `json:"Data" example:"Logout success"`
}

type DeleteResponseExample struct {
	Err  bool   `json:"Err" example:"false"`
	Data string `json:"Data" example:"Delete success"`
}

var SecretKey = []byte("some-key")

// Проверка заголовков для всех запросов
func (u *UserController) HandlerFunc(rules string) bool {
	switch rules {
	case "GetAll", "Get", "Logout", "Put": // rules - в значении имеет название функции, которая выполняется при вызове метода
		arrayToken := u.Ctx.Request.Header["Authorization"]
		fmt.Println(arrayToken, "arrtok")
		token, _ := models.VerifyToken(arrayToken[0][7:]) // забираем данные из токена
		fmt.Println(token, token["id"])                   // Можем делать проверку
		if len(arrayToken) > 0 {                          // проверка токена, тут проверка по хедеру.
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
// @Param	body		body 	models.User	true		"все данные о новом пользователе(айди выберется автоматически, так что можно оставлять 0)"
// @Success 200 {object} controllers.UserPostResponseExample
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var User models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &User)
	uid, err := models.AddUser(User)
	if err != nil {
		u.Data["json"] = Response{Err: true, Data: "Такой логин уже занят!"}
	} else {
		u.Data["json"] = Response{Err: false, Data: uid}
	}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Param Authorization header true "хедер Авторизации. пример(фигурные скобки убрать): Bearer {token}"
// @Success 200 {object} controllers.AllUsersResponseExample
// @Failure 401 :not authorized
// @router /all [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = Response{Err: false, Data: users}
	u.ServeJSON()
}

// @Title Get
// @Description get user using jwt token
// @Param Authorization header true "хедер Авторизации. пример(фигурные скобки убрать): Bearer {token}"
// @Success 200 {object} models.User
// @Failure 404 no such user
// @router / [get]
func (u *UserController) Get() {
	arrayToken := u.Ctx.Request.Header["Authorization"]
	fmt.Println(arrayToken, "arrtok")
	token, _ := models.VerifyToken(arrayToken[0][7:])
	uid := int64(token["id"].(float64))
	user, err := models.GetUser(uid)
	if err != nil {
		u.Data["json"] = Response{Err: true, Data: err.Error()}
		u.Abort("404")
	} else {
		u.Data["json"] = Response{Err: false, Data: user}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param Authorization header true "хедер Авторизации. пример(фигурные скобки убрать): Bearer {token}"
// @Param	body		body 	models.User	true		"Все данные о пользователе"
// @Success 200 {object} models.User
// @Failure 403 data problems
// @router / [put]
func (u *UserController) Put() {
	arrayToken := u.Ctx.Request.Header["Authorization"]
	fmt.Println(arrayToken, "arrtok")
	token, _ := models.VerifyToken(arrayToken[0][7:])
	uid := int64(token["id"].(float64))
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	err := models.UpdateUser(&user)
	if err != nil {
		u.Data["json"] = Response{Err: true, Data: err.Error()}
	} else {
		uu, _ := models.GetUser(uid)
		u.Data["json"] = uu
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param Authorization header true "хедер Авторизации. пример(фигурные скобки убрать): Bearer {token}
// @Success 200 {object} controllers.DeleteResponseExample
// @Failure 403 header problems ??
// @router / [delete]
func (u *UserController) Delete() {
	arrayToken := u.Ctx.Request.Header["Authorization"]
	fmt.Println(arrayToken, "arrtok")
	token, _ := models.VerifyToken(arrayToken[0][7:])
	uid := int64(token["id"].(float64))
	del := models.DeleteUser(uid)
	if del {
		u.Data["json"] = Response{Err: false, Data: "Delete Success"}
	} else {
		u.Data["json"] = Response{Err: true, Data: "User not found!"}
	}
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	login			query 	string	true		"The login"
// @Param	password		query 	string	true		"The password"
// @Success 200 {object} controllers.LoginResponseExample
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	var user models.User
	user.Login = u.Ctx.Input.Query("login")
	user.Password = u.Ctx.Input.Query("password")
	token := models.Login(user)
	if token != "" {
		u.Data["json"] = Response{Err: false, Data: token}
		u.Ctx.Output.Header("Authorization", token)
		// установка значения сессии
		u.SetSession("accessToken", token)
		fmt.Println(u.Ctx.Output)
	} else {
		u.Data["json"] = Response{Err: true, Data: "Неверное имя пользователя или пароль!"}
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Param Authorization header true "Хедер авторизации. example: Bearer {token} "
// @Success 200 {object} controllers.LogoutResponseExample
// @router /logout [get]
func (u *UserController) Logout() {
	// получение значения сессии
	//u.GetSession("accessToken")
	// удаление значения сессии
	// u.DelSession("accessToken")
	// уничтожение сессии
	u.DestroySession()
	u.Data["json"] = Response{Err: false, Data: "Logout success"}
	u.ServeJSON()
}
