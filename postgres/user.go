package postgres

import (
	"fmt"

	tangent "github.com/dfsantos-source/Tangent"
)

var _ tangent.UserService = (*UserServiceDB)(nil)

type UserServiceDB struct {
	db *DB
}

func CreateUserServiceDB(db *DB) *UserServiceDB {
	return &UserServiceDB{db: db}
}

func (s *UserServiceDB) User(id int) (*tangent.User, error) {
	query := `
		SELECT id, name, email, timestamp FROM Users WHERE ID = $1
	`

	user := &tangent.User{}

	row := s.db.sql.QueryRow(query, id)

	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Timestamp); err != nil {
		fmt.Println(err)
		return user, err
	}

	return user, nil
}

func (s *UserServiceDB) Users() ([]*tangent.User, error) {
	query := `
		SELECT id, name, email, timestamp FROM Users
	`

	rows, err := s.db.sql.Query(query)

	if err != nil {
		panic(err)
	}

	users := []*tangent.User{}

	for rows.Next() {
		user := &tangent.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Timestamp); err != nil {
			fmt.Println(err)
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *UserServiceDB) CreateUser(user *tangent.User) (int, error) {
	query := `
		INSERT INTO USERS (name, email, password) VALUES($1, $2, $3)
		RETURNING ID
	`

	id := -1

	err := s.db.sql.QueryRow(
		query,
		user.Name,
		user.Email,
		user.Password,
	).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	return id, nil
}
func (s *UserServiceDB) DeleteUser(id int) error {
	query := `
		DELETE FROM Users WHERE ID = $1
	`
	_, err := s.db.sql.Exec(query, id)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
