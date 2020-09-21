package domain

type User struct {
	uuid string
	email string
	password string
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
