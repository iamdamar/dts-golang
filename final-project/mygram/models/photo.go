package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	ID			uint		`gorm:"PrimaryKey" json:"id"`
	Title		string		`gorm:"not null" json:"title" form:"title" valid:"required~Title is required"`
	Caption		string		`json:"caption" form:"caption"`
	PhotoUrl	string		`gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Photo URL is required"`

	UserID		uint		`json:"user_id"`
	User		*User		`json:"user,omitempty"`

	CreatedAt	*time.Time	`json:"created_at,omitempty"`
	UpdatedAt	*time.Time	`json:"updated_at,omitempty"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}