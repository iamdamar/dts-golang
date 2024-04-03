package models

import (
	"time"
	"errors"
	"fmt"
	
	"gorm.io/gorm"
)

type Item struct {
	// ID			uint		`gorm:"primaryKey"`
	// Name		string		`gorm:"not null;type:varchar(191)"`
	// Brand	string		`gorm:"not null;type:varchar(191)"`
	// UserID		uint
	ID			uint		`gorm:"primaryKey"`
	Code		string		`gorm:"not null;type:varchar(191)"`
	Description	string		`gorm:"not null;type:varchar(191)"`
	Quantity	string		`gorm:"not null;type:varchar(191)"`
	OrderID		uint
	OrderedAt	time.Time
}

func (p *Item) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Item Before Create()")

	if len(p.Code) < 4 {
		err = errors.New("item code is too short")
	}

	return
}