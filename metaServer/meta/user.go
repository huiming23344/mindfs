package meta

type User struct {
	Name     string
	Password string
}

type UserGroup struct {
	Name  string
	Users []*User
}
