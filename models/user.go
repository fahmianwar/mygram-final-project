package models

import (
	"mygram-final-project/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// User represents the model for an user
type User struct {
	ID        uint       `gorm:"primaryKey"`
	Username  string     `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
	Email     string     `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password  string     `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age       string     `gorm:"not null" json:"age" form:"age" valid:"required~Age user is required"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(u)
	u.Password = helpers.HashPass(u.Password)

	return err
}
