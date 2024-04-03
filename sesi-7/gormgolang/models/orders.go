package models

import "time"

type Order struct {
	// ID			uint		`gorm:"primaryKey"`
	// Email		string		`gorm:"not null;unique;type:varchar(191)"`
	// Products	[]Product
	// CreatedAt	time.Time
	ID				uint		`gorm:"primaryKey"`
	CustomerName	string		`gorm:"not null;unique;type:varchar(191)"`
	Items			[]Item
	OrderedAt		time.Time
}