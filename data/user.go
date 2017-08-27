package data

import "github.com/pkg/errors"

func (e *Engine) CreateUser(username, password string) (User, error) {
	user := User{
		Username: username,
		Password: password,
	}
	err := e.db.Create(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (e *Engine) AuthorizeUser(username, password string) (user User, err error) {
	err = e.db.Where(&User{Username: username}).First(&user).Error
	if err != nil {
		return User{}, err
	}
	if user.Password == password {
		return user, nil
	}
	return User{}, errors.New("Invalid user or password")
}
