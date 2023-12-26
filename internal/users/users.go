package user

type User struct {
	Id  string
	New bool
}

func NewUser(id string, new bool) *User {
	return &User{
		Id:  id,
		New: new,
	}
}
