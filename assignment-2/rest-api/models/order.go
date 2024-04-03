package models

import "time"

type Order struct {
	OrderId			uint	`gorm:"PrimaryKey"`
	CustomerName	string	`gorm:"not null;unique;type:varchar(191)"`
	Items			[]Item
	OrderedAt		time.Time
}

// i want the output of GET orders to be like this:
// [
//     {
//         "OrderId": 1,
//         "CustomerName": "Ahmad Kasim",
//         "Items":
// 			[{
// 				"ItemId": 2,
// 				"ItemCode": "123",
// 				"Description": "Macbook 12",
// 				"Quantity": "4",
// 				"OrderId": 1
// 			}]
//         "OrderedAt": "0001-01-01T07:00:00+07:00"
//     }
// ]

