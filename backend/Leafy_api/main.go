package main

import (
	_ "Leafy_api/routers"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	_ "github.com/beego/beego/v2/server/web/session/postgres"
	_ "github.com/lib/pq"
)

func main() {
	//Загрузка строки подключения к БД из конфига
	sqlconn, _ := beego.AppConfig.String("sqlconn")
	//Подключаем драйвер
	orm.RegisterDriver("postgres", orm.DRPostgres)
	//Подключение к БД
	orm.RegisterDataBase("Leafy", "postgres", sqlconn)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//Настройка CORS
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},                                                //Доступ с любого адреса
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"}, //Допустимые запросы
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	beego.Run()
}
