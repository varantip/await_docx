package controllers

import (
	"Leafy_api/models"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type UserPlantsResponseExample struct {
	Err  bool `json:"Err" example:"false"`
	Data []models.UserPlant
}

type UserPlantPutResponseExample struct {
	Err  bool   `json:"Err" example:"false"`
	Data string `json:"Data" example:"Успешно изменено"`
}

// Operations about User's plants(Unfinished)
type UserPlantController struct {
	beego.Controller
}

// Проверка заголовков для всех запросов
func (upt *UserPlantController) HandlerFunc(rules string) bool {
	switch rules {
	case "Post", "Put", "Get": // rules - в значении имеет название функции, которая выполняется при вызове метода
		arrayToken := upt.Ctx.Request.Header["Authorization"]
		if len(arrayToken) == 0 {
			println("first")
			break
		}
		if len(arrayToken[0]) < 7 {
			println("second")
			break
		}
		// fmt.Println(arrayToken, "arrtok")
		token, err := models.VerifyToken(arrayToken[0][7:]) // забираем данные из токена
		if err != nil {
			println("third")
			break
		}
		fmt.Println(token, token["id"]) // Можем делать проверку
		// проверка токена, тут проверка по хедеру.
		if len(arrayToken) > 0 {
			return false
		}

	default:
		return false
	}
	upt.Abort("401") // выдаем ошибку авторизации
	return true
}

// @Title Get
// @Description По JWT токену узнаётся пользователь и получаются все его растения
// @Param Authorization header true "хедер Авторизации. пример(фигурные скобки убрать): Bearer {token}"
// @Success 200 {object} controllers.UserPlantsResponseExample
// @router / [get]
func (upt *UserPlantController) Get() {
	cryptToken := upt.Ctx.Request.Header["Authorization"][0][7:]
	verifiedToken, _ := models.VerifyToken(cryptToken)
	id := int64(verifiedToken["id"].(float64))
	var PUPlants []models.ParsedUserPlant
	UPlants := models.GetAllUsersPlants(id)
	for _, el := range *UPlants {
		PUPlants = append(PUPlants, models.ParsedUserPlant{UserPlant_id: el.UserPlant_id, User_id: el.User_id, PlantType_id: el.PlantType_id, Disease_id: el.Disease_id, DateOfLastWatering: el.DateOfLastWatering.Format(layout), DateOfPlanting: el.DateOfPlanting.Format(layout), Nickname: el.Nickname})
	}
	if len(PUPlants) > 0 {
		upt.Data["json"] = Response{Err: false, Data: PUPlants}
	} else {
		upt.Data["json"] = Response{Err: true, Data: "Растений нет"}
	}
	upt.ServeJSON()
}

// @Title GetAll
// @Description Подаёшь в ссылку айди юзера вместо :uid и получаешь на выходе все растения этого пользователя
// @Success 200 {object} controllers.UserPlantsResponseExample
// @Failure 403 :uid is empty
// @router /all [get]
func (upt *UserPlantController) GetAll() {
	UPlants := models.GetAllUserPlants()
	var PUPlants []models.ParsedUserPlant
	for _, el := range *UPlants {
		PUPlants = append(PUPlants, models.ParsedUserPlant{UserPlant_id: el.UserPlant_id, User_id: el.User_id, PlantType_id: el.PlantType_id, Disease_id: el.Disease_id, DateOfLastWatering: el.DateOfLastWatering.Format(layout), DateOfPlanting: el.DateOfPlanting.Format(layout), Nickname: el.Nickname})
	}
	upt.Data["json"] = Response{Err: false, Data: PUPlants}
	upt.ServeJSON()
}

// @Title CreateUserPlant
// @Description Создание userplant'ов
// @Param	body		body 	models.UserPlant		"Данные о новом растении(айди юзера выбирается из хедера авторизации, дата последнего полива - автоматом: 01.01.01 00:00:00, остальное нужно указать ).<br> <b>Время обязательно должно быть строго форматировано</b>: YYYY-MM-DDThh:mm:ssTZ, T - просто буква, TZ - часовой пояс, для калининграда вместо TZ нужно +0200"
// @Param Authorization header true "хедер Авторизации. пример(фигурные скобки убрать): Bearer {token}"
// @Success 200 {object} controllers.UserPostResponseExample
// @Failure 403 some of data is incorrect
// @router / [post]
func (upt *UserPlantController) Post() {
	arrayToken := upt.Ctx.Request.Header["Authorization"]
	token, _ := models.VerifyToken(arrayToken[0][7:])
	//println(token) //вывел верифицированный токен для проверки
	var parsed_UPlant models.ParsedUserPlant
	json.Unmarshal(upt.Ctx.Input.RequestBody, &parsed_UPlant)
	parsedPlantingDate, err := time.Parse(layout, parsed_UPlant.DateOfPlanting)
	// UserPlant_id       int64  `orm:"column(UserPlant_id);pk;auto"`
	// User_id            int64  `orm:"column(User_id)"`
	// PlantType_id       int64  `orm:"column(PlantType_id)"`
	// Disease_id         int64  `orm:"column(Disease_id)"`
	// DateOfLastWatering string `orm:"column(DateOfLastWatering)"`
	// DateOfPlanting     string `orm:"column(DateOfPlanting)"`
	// Nickname           string `orm:"size(256);column(Nickname)"`
	if err != nil {
		println("incorrect plant date")
		upt.Abort("403")
		return
	}

	token_id := int64(token["id"].(float64))
	var UPlant models.UserPlant = models.UserPlant{UserPlant_id: parsed_UPlant.UserPlant_id,
		User_id:            token_id,
		PlantType_id:       parsed_UPlant.PlantType_id,
		Disease_id:         parsed_UPlant.Disease_id,
		DateOfLastWatering: time.Time{},
		DateOfPlanting:     parsedPlantingDate,
		Nickname:           parsed_UPlant.Nickname,
	}
	fmt.Println(UPlant.User_id)

	parsed_user, err := models.GetUser(int64(token_id))
	if err != nil {
		print("some id problem ", parsed_user.User_id, " ", UPlant.User_id)
		upt.Abort("403")
		return
	}

	_, err = models.GetDisease(UPlant.Disease_id)
	if err != nil {
		print("nonexistent disease")
		upt.Abort("403")
		return
	}

	_, err = models.GetPlantType(UPlant.PlantType_id)
	if err != nil {
		print("nonexistent planttype")
		upt.Abort("403")
		return
	}
	uptid := models.AddUserPlant(UPlant)
	if uptid > 0 {
		upt.Data["json"] = Response{Err: false, Data: uptid}
	} else {
		upt.Data["json"] = Response{Err: true, Data: "ошибка при создании"}
	}
	upt.ServeJSON()
}

// @Title UpdateUserPlant
// @Description Обновление userplant'ов
// @Param	body		body 	models.UserPlant		"<b>Все</b> данные о новом растении (даже айди юзера).<br> <b>Время обязательно должно быть строго форматировано</b>: YYYY-MM-DDThh:mm:ssTZ, T - просто буква, TZ - часовой пояс, для калининграда вместо TZ нужно +0200"
// @Param Authorization header true "хедер Авторизации. пример(фигурные скобки убрать): Bearer {token}"
// @Success 200 {object} controllers.UserPlantPutResponseExample
// @Failure 403 some of data is incorrect
// @router / [put]
func (upt *UserPlantController) Put() {
	arrayToken := upt.Ctx.Request.Header["Authorization"]
	token, _ := models.VerifyToken(arrayToken[0][7:])
	//println(token) //вывел верифицированный токен для проверки
	var parsed_UPlant models.ParsedUserPlant
	json.Unmarshal(upt.Ctx.Input.RequestBody, &parsed_UPlant)
	parsedPlantingDate, err := time.Parse(layout, parsed_UPlant.DateOfPlanting)
	if err != nil {
		println("incorrect plant date")
		upt.Abort("403")
		return
	}
	parsedWateringDate, err := time.Parse(layout, parsed_UPlant.DateOfLastWatering)
	if err != nil {
		println("incorrect watering date")
		upt.Abort("403")
		return
	}
	token_id := int64(token["id"].(float64))
	var UPlant models.UserPlant = models.UserPlant{UserPlant_id: token_id,
		User_id:            parsed_UPlant.User_id,
		PlantType_id:       parsed_UPlant.PlantType_id,
		Disease_id:         parsed_UPlant.Disease_id,
		DateOfLastWatering: parsedWateringDate,
		DateOfPlanting:     parsedPlantingDate,
		Nickname:           parsed_UPlant.Nickname,
	}
	fmt.Println(UPlant.User_id)

	parsed_user, err := models.GetUser(int64(token_id))
	if err != nil {
		print("id mismatch", parsed_user.User_id)
		upt.Abort("403")
		return
	}

	_, err = models.GetDisease(UPlant.Disease_id)
	if err != nil {
		print("nonexistent disease")
		upt.Abort("403")
		return
	}

	_, err = models.GetPlantType(UPlant.PlantType_id)
	if err != nil {
		print("nonexistent planttype")
		upt.Abort("403")
		return
	}
	err = models.UpdateUserPlant(UPlant)
	if err == nil {
		upt.Data["json"] = Response{Err: false, Data: "Успешно изменено"}
	} else {
		upt.Data["json"] = Response{Err: true, Data: "ошибка при создании"}
	}
	upt.ServeJSON()
}
