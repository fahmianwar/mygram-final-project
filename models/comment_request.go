package models

//  CommentRequest represents the model for an commentRequest
type CommentRequest struct {
	PhotoID uint   `gorm:"not null" json:"photoId" form:"photoId" valid:"required~PhotoId is required"`
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~Message is required"`
}
