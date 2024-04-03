package models

type Item struct {
	ItemId			uint	`gorm:"PrimaryKey"`
	ItemCode		string	`gorm:"not null;unique;type:varchar(191)"`
	Description		string	`gorm:"not null;unique;type:varchar(191)"`
	Quantity		string	`gorm:"not null;unique;type:varchar(191)"`
	OrderId			uint
}