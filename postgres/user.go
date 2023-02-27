package postgres

import (
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
	return nil, nil
}

func (s *UserServiceDB) Users() ([]*tangent.User, error) {
	return nil, nil
}

func (s *UserServiceDB) CreateUser(user *tangent.User) error {
	return nil
}
func (s *UserServiceDB) DeleteUser(id int) error {
	return nil
}
