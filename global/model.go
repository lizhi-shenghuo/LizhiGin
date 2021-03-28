package global

import (
	"gorm.io/gorm"
	"time"
)

type LizhiModel struct {
	ID uint `gorm:"primarykey"`
	CreateAt time.Time
	UpdateAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
