package model

import "time"

type Type struct {
	ID              uint      `json:"id" gorm:"column:id;primarykey"`
	Name            string    `json:"name" gorm:"column:name;type:varchar(255);not null"`
	Description     string    `json:"description" gorm:"column:description;type:text;not null"`
	BackgroundColor string    `json:"background_color" gorm:"column:background_color;type:varchar(255);not null"`
	TextColor       string    `json:"text_color" gorm:"column:text_color;type:varchar(255);not null"`
	CreatedAt       time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Type) TableName() string {
	return "types"
}
