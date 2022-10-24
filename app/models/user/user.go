package user

import "goblog/app/models"

type User struct {
	models.BaseModel

	Name     string `gorm:"type:varchar(255);not null;unique" valid:"name"`
	Email    string `gorm:"type:varchar(255);unique;" valid:"email"`
	Password string `gorm:"type:varchar(255)" valid:"password"`
	// gorm:"-" -- 设置 Gorm 在读写时略过此字段
	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}

func (user *User) Link() string {
	return ""
}