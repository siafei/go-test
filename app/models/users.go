package models

import(
	."go-test/bootstrap"
	"golang.org/x/crypto/bcrypt"
)

type User struct{
	Model
	Mobile 		string  `json:"created_at"`
	Uname		string  `json:"created_at"`
	Password	string  `json:"created_at"`
}

func (User) TableName() string {
	return "users"
}

func (*User) CreateUser(data map[string]interface{}) error {
	password , _ := bcrypt.GenerateFromPassword([]byte(data["password"].(string)), bcrypt.DefaultCost)
	user := User{
		Uname:       data["uname"].(string),
		Password:     string(password),
		Mobile:         data["mobile"].(string),
	}
	if err := Db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}