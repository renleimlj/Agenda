package entity

import (
	"encoding/json"
	"os"
)

type User struct {
	Name, Password, Email, Phone string
}

func (u *User) Init(name, password, email, phone string) {
	u.Name= name
	u.Password= password
	u.Email= email
	u.Phone= phone
}

func Register(name, password, email, phone string) int {
	file, err := os.OpenFile("User", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	for decoder.More() {
		var users User
		decoder.Decode(&users)
		if users.Name == name { // 创建失败，已存在该用户
			file.Close()
			return 1
		}
	}
	userinfo := make(map[string]string)
	userinfo["Name"], userinfo["Password"], userinfo["Email"], userinfo["Phone"] = name, password, email, phone
	encoder := json.NewEncoder(file)
	encoder.Encode(userinfo)
	file.Close()
	return 0
}

func (u User) GetName() string {
	return u.Name;
}

func (u *User) SetName(name string) {
	u.Name = name;
}

func (u User) GetPassword() string {
	return u.Password;
}

func (u *User) SetPassword(password string) {
	u.Password = password;
}

func (u User) GetEmail() string {
	return u.Email;
}

func (u *User) SetEmail(email string) {
	u.Email = email;
}

func (u User) GetPhone() string {
	return u.Phone;
}

func (u *User) SetPhone(phone string) {
	u.Phone = phone;
}
