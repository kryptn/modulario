package data

import "github.com/pkg/errors"

func (e *Engine) Register(username, password string) (User, error) {

	var count = 0
	e.db.Model(&User{}).Where(&User{Username: username}).Count(&count)
	if count == 0 {
		user := User{
			Username: username,
			Password: password,
		}
		return user, e.db.Create(&user).Error
	}

	return User{}, errors.New("User Exists")
}

func (e *Engine) Login(username, password string) (user User, err error) {
	err = e.db.Where(&User{Username: username, Password: password}).First(&user).Error
	return
}
