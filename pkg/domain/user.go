package domain

type User struct {
	uuid     string
	email    string
	password string
}

func NewUser(email string, password string) *User {
	return &User{
		email:    email,
		password: password,
	}
}

func (u *User) SetUUID(uuid string) {
	u.uuid = uuid
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u User) UUID() string {
	return u.uuid
}

func (u User) Email() string {
	return u.email
}

func (u User) Password() string {
	return u.password
}
