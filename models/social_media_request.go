package models

// SocialMediaRequest represents the model for an socialMediaRequest
type SocialMediaRequest struct {
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~Name is required"`
	SocialMediaUrl string `gorm:"not null;type:text" json:"socialMediaUrl" form:"socialMediaUrl" valid:"required~SocialMediaUrl is required"`
}
