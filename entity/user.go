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

func Login(name, password string) int {
	file, err := os.OpenFile("User", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	for decoder.More() {
		var users User
		decoder.Decode(&users)
		if users.Name == name { // 存在该用户名
			if users.Password == password {
				os.Remove("CurUser")
				file2, err2 := os.OpenFile("CurUser", os.O_RDWR|os.O_CREATE, 0666)
				if err2 != nil {
					panic(err2)
				}
				encoder := json.NewEncoder(file2)
				encoder.Encode(users.Name)
				file2.Close()
				file.Close()
				return 0
			} else {
				return 1
			}
		}
	}
	return 1
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Logout() int {
	_, e := PathExists("CurUser")
	if e == nil {
		os.Remove("CurUser")
		return 0
	} else {
		return 1
	}
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
