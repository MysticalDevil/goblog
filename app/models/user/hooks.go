package user

import (
	"goblog/pkg/password"
	"gorm.io/gorm"
)

//// BeforeCreate Gorm 的模型钩子，创建模型前调用
//func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
//	user.Password = password.Hash(user.Password)
//	return
//}
//
//// BeforeUpdate Gorm 的模型钩子，更新模型前调用
//func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
//	if !password.IsHashed(user.Password) {
//		user.Password = password.Hash(user.Password)
//	}
//	return
//}

// BeforeSave Gorm 的模型钩子函数，在保存和更新前调用
func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	if !password.IsHashed(user.Password) {
		user.Password = password.Hash(user.Password)
	}
	return
}
