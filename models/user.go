package models

import (
	"golang-mygram/helpers"
	"strconv"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// User represents the model of a user
type User struct {
	GormModel
	Username string `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password string `gorm:"not null" json:"-" form:"password" valid:"required~Your password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age      int    `gorm:"not null" json:"age" form:"age" valid:"required~Your age is required, numeric~Age must be numeric, age~Age has to more than 8"`
	// Admin        bool          `gorm:"not null;default:0" json:"admin" form:"admin"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Socialmedias []Socialmedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	govalidator.TagMap["age"] = govalidator.Validator(func(str string) bool {
		age, _ := strconv.Atoi(str)
		return age > 8
	})
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return err
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return err
}
