package domain

type User struct {
	uuid string `json:"uuid" binding:"-"`
	email string `json:"email" binding:"required"`
	password string `json:"password" binding:"required"`
}

func NewUser(email string, password string) *User {
	return &User{
		email: email,
		password: password,
	}
}

func (u *User) SetUUID(uuid string) {
	u.uuid = uuid
}

func (u User) GetUUID() string {
	return u.uuid
}

func (u User) GetEmail() string {
	return u.email
}

func (u User) GetPassword() string {
	return u.password
}
