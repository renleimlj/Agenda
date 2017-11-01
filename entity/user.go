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
				f, _ := os.Create("CurUser")
    			defer f.Close()
    			_, err := f.WriteString(name)
    			if err != nil {}
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
	b, _ := PathExists("CurUser")
	if b == true {
		os.Remove("CurUser")
		return 0
	} else {
		return 1
	}
}

func UserExists(username string) bool {
	file, err := os.OpenFile("User", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	for decoder.More() {
		var users User
		decoder.Decode(&users)
		if users.Name == username {
			return true
		}
	}
	return false;
}

func Query(username string) (string, int) {
	b, _ := PathExists("CurUser")
	if b != true {
		return "falied", 2
	}
	file, err := os.OpenFile("User", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	for decoder.More() {
		var users User
		decoder.Decode(&users)
		if users.Name == username {
			var response string = "Name: " + users.Name + "\n" + "Email: " + users.Email + "\n" + "Phone: " + users.Phone + "\n"
			return response, 0
		}
	}
	return "failed", 1
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
