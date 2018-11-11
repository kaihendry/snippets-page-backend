package model

//Collection name
const Collection = "users"

//User structure
type User struct {
	Db           base
	ID           string
	PasswordHash string
	Email        string
}

func (u *User) Exists(email string, password string) []byte {

	u.Db.FindBy()

	return []byte("sdfsdf")
}
