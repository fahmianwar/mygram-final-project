package models

// PhotoRequest represents the model for an photoRequest
type PhotoRequest struct {
	Title    string `gorm:"not null" json:"title" form:"title" valid:"required~Title is required"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `gorm:"not null" json:"photoUrl" form:"photoUrl" valid:"required~PhotoUrl is required"`
}
