package object

type User struct {
	Name, Password, Email, Phone string
}

func (u User) Init(name, password, email, phone string) {
	u.Name= name
	u.Password= password
	u.Email= email
	u.Phone= phone
}

func (u User) GetName() string {
	return u.Name;
}

func (u User) SetName(name string) {
	u.Name = name;
}

func (u User) GetPassword() string {
	return u.Password;
}

func (u User) SetPassword(password string) {
	u.Password = password;
}

func (u User) GetEmail() string {
	return u.Email;
}

func (u User) SetEmail(email string) {
	u.Email = email;
}

func (u User) GetPhone() string {
	return u.Phone;
}

func (u User) SetPhone(phone string) {
	u.Phone = phone;
}
