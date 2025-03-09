package main

import (
	_ "Leafy_api/routers"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/beego/beego/v2/server/web/session/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	println(apppath)
	beego.TestBeegoInit(apppath)
	//Загрузка строки подключения к БД из конфига
	sqlconn, _ := beego.AppConfig.String("sqlconn")
	fmt.Println(sqlconn)
	//Подключаем драйвер
	orm.RegisterDriver("postgres", orm.DRPostgres)
	//Подключение к БД
	orm.RegisterDataBase("Leafy", "postgres", sqlconn)
}

// TestGetUser проверяет запрос /v1/User/:id .
func TestGetUser(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/v1/user/5", nil)
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	res := strconv.FormatInt(int64(w.Code), 10) // если нужно посмотреть код
	println(res)
	res = w.Body.String() // если нужно посмотреть вывод
	println(res)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\n  \"err\": false,\n  \"data\": {\n    \"User_id\": 5,\n    \"Name\": \"eggs benedict\",\n    \"Login\": \"john\",\n    \"Password\": \"doe\"\n  }\n}", w.Body.String())
}
