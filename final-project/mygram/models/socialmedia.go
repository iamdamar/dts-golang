package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Socialmedia struct {
	ID				uint		`gorm:"PrimaryKey" json:"id"`
	Name			string		`gorm:"not null" json:"name" form:"name" valid:"required~Your name is required"`
	SocialMediaUrl	string		`gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~Social Media URL is required"`

	UserID			uint		`json:"user_id"`
	User			*User		`json:"user,omitempty"`

	CreatedAt		*time.Time	`json:"created_at,omitempty"`
	UpdatedAt		*time.Time	`json:"updated_at,omitempty"`
}

func (p *Socialmedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Socialmedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}