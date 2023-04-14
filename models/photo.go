package models

import "time"

// Photo represents the model for an photo
type Photo struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null" json:"title" form:"title" valid:"required~Title is required"`
	Caption   string `json:"caption" form:"caption"`
	PhotoUrl  string `gorm:"not null" json:"photoUrl" form:"photoUrl" valid:"required~PhotoUrl is required"`
	UserID    uint   `gorm:"not null" json:"userId" form:"userId" valid:"required~UserId is required"`
	User      User   `gorm:"foreignKey:UserID"`
	UpdatedAt time.Time
	CreatedAt time.Time
}
