package user

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/password"
	"goblog/pkg/types"
)

func (user *User) Create() (err error) {
	if err = model.DB.Create(&user).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

func Get(idStr string) (User, error) {
	var user User
	id := types.StringToUint64(idStr)
	if err := model.DB.First(&user,id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func GetByEmail(email string) (User, error) {
	var user User
	if err := model.DB.Where("email=?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (user *User) ComparePassword(_password string) bool {
	return password.CheckHash(_password, user.Password)
}

func (user *User) Update() (rowsAffected int64, err error) {
	result := model.DB.Save(&user)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, nil
	}
	return result.RowsAffected, nil
}
