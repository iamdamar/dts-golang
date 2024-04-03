package models

import (
	"mygram/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	ID			uint		`gorm:"PrimaryKey"`
	Username	string		`gorm:"not null" json:"username" form:"username" valid:"required~Your username is required"`
	Email		string		`gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password	string		`gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age			int			`gorm:"not null" json:"age" form:"age" valid:"required~Minimum age is 9"`
	CreatedAt	*time.Time	`json:"created_at,omitempty"`
	UpdatedAt	*time.Time	`json:"updated_at,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}