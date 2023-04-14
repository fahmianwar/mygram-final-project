package models

type UserLogin struct {
	Username string `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
}
