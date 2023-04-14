package models

import "time"

// SocialMedia represents the model for an socialMedia
type SocialMedia struct {
	ID             uint   `gorm:"primaryKey" json:"id" form:"id" valid:"required~ID is required"`
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~Name is required"`
	SocialMediaUrl string `gorm:"not null;type:text" json:"socialMediaUrl" form:"socialMediaUrl" valid:"required~SocialMediaUrl is required"`
	UserID         uint   `gorm:"not null" json:"userId" form:"userId" valid:"required~UserId is required"`
	UpdatedAt      time.Time
	CreatedAt      time.Time
	User           User `gorm:"foreignKey:UserID"`
}
