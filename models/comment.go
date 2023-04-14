package models

import "time"

//  Comment represents the model for an comment
type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null" json:"userId" form:"userId" valid:"required~UserId is required"`
	PhotoID   uint   `gorm:"not null" json:"photoId" form:"photoId" valid:"required~PhotoId is required"`
	Message   string `gorm:"not null" json:"message" form:"message" valid:"required~Message is required"`
	UpdatedAt time.Time
	CreatedAt time.Time
	User      User  `gorm:"foreignKey:UserID"`
	Photo     Photo `gorm:"foreignKey:PhotoID"`
}
