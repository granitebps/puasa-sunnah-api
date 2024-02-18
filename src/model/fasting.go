package model

import (
	"time"

	"gorm.io/datatypes"
)

type Fasting struct {
	ID         uint           `json:"id" gorm:"column:id;primarykey"`
	CategoryID uint           `json:"category_id" gorm:"column:category_id;not null"`
	TypeID     uint           `json:"type_id" gorm:"column:type_id;not null"`
	Date       datatypes.Date `json:"date" gorm:"column:date;not null"`
	Year       uint32         `json:"year" gorm:"column:year;not null"`
	Month      uint32         `json:"month" gorm:"column:month;not null"`
	Day        uint32         `json:"day" gorm:"column:day;not null"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

func (Fasting) TableName() string {
	return "fastings"
}
