package model

import "time"

type Source struct {
	ID        uint      `json:"id" gorm:"column:id;primarykey"`
	Url       string    `json:"url" gorm:"column:url;type:text;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Source) TableName() string {
	return "sources"
}
