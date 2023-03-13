package tangent

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name""`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Timestamp string `json:"timestamp"`
}

type UserService interface {
	User(id int) (*User, error)
	Users() ([]*User, error)
	CreateUser(*User) (int, error)
	DeleteUser(id int) error
}
