package meta

type User struct {
	Username string
	Password string
}

type UserGroup struct {
	Name  string
	Users []*User
}
