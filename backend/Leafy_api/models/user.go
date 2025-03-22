package models

import (
	"errors"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterModel(new(User))
}

type User struct {
	User_id  int64  `orm:"column(User_id);pk;auto"`
	Name     string `orm:"size(256); column(Name)"`
	Login    string `orm:"size(256); column(Login)"`
	Password string `orm:"size(256); column(Password)"`
}

func (u *User) TableName() string {
	// db table name
	return "User"
}

var SecretKey = []byte("some-key")

func CreateToken(u User) (string, error) {
	// создаем заявку
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    u.User_id,
		"login": u.Login,
	})
	// генерируем токен
	tokenString, err := claims.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	// Проверка на ошибки
	if err != nil {
		return nil, err
	}

	// Проверка валидности токена
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Возврат данных токена
	return token.Claims.(jwt.MapClaims), nil
}

func AddUser(u User) (int64, error) {
	o := orm.NewOrmUsingDB("Leafy")
	qb, _ := orm.NewQueryBuilder("postgres")
	fmt.Println(u.User_id, u.Login, u.Name, u.Password)
	var checkUser User //будем проверять существование пользователя
	o.Raw(qb.Select("User_id", "Name", "Login", "Password").From("User").Where(fmt.Sprintf(`"Login" = '%s'`, u.Login)).Limit(1).String()).QueryRow(&checkUser)
	if checkUser.Password != "" {
		return 0, errors.New("пользователь с таким логином уже существует")
	}
	id, err := o.Insert(&u)
	if err != nil {
		fmt.Println(err)
	}

	return id, nil
}

func GetUser(uid int64) (u *User, err error) {
	o := orm.NewOrmUsingDB("Leafy")
	user := User{User_id: uid}
	err = o.Read(&user)
	if err == orm.ErrNoRows {
		return nil, errors.New("пользователь с таким id не найден")
	}
	return &user, nil
}

func GetAllUsers() *[]User {
	var users []User
	o := orm.NewOrmUsingDB("Leafy")
	qb, _ := orm.NewQueryBuilder("postgres")
	o.Raw(qb.Select("User_id", "Name", "Login", "Password").From("User").String()).QueryRows(&users)
	return &users
}

func UpdateUser(uu *User) (err error) {
	o := orm.NewOrmUsingDB("Leafy")
	_, err = o.Update(uu, "Name", "Login", "Password")
	if err != nil {
		return errors.New("пользователь не найден")
	}
	return nil
}

func Login(u User) string {
	o := orm.NewOrmUsingDB("Leafy")
	err := o.Read(&u, "Login", "Password")
	fmt.Println(u)
	if err != orm.ErrNoRows {
		tokenString, _ := CreateToken(u)
		return tokenString
	}
	return ""
}

func DeleteUser(uid int64) bool {
	o := orm.NewOrmUsingDB("Leafy")
	user := User{User_id: uid}
	_, err := o.Delete(&user)
	return err == nil
}
