package tangent

type User struct {
	ID        int    `json:"id`
	Name      string `json:"name"`
	Email     string `json:"email`
	Password  string `json:"password`
	CreatedAt string `json:"createdAt`
	UpdatedAt string `json:"createdAt`
}

type UserService interface {
	User(id int) (*User, error)
	Users() ([]*User, error)
	CreateUser(*User) error
	DeleteUser(id int) error
}
