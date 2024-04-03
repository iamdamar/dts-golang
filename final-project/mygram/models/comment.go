package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	ID			uint		`gorm:"PrimaryKey" json:"id"`
	Message		string		`gorm:"not null" json:"message" form:"message" valid:"required~Title is required"`

	PhotoID		uint		`json:"photo_id" form:"photo_id"`

	UserID		uint		`json:"user_id"`
	User		*User		`json:"user,omitempty"`

	CreatedAt	*time.Time	`json:"created_at,omitempty"`
	UpdatedAt	*time.Time	`json:"updated_at,omitempty"`
}

func (p *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}