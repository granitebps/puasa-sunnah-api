package model

import "time"

type Category struct {
	ID        uint      `json:"id" gorm:"column:id;primarykey"`
	Name      string    `json:"name" gorm:"column:name;type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Category) TableName() string {
	return "categories"
}
