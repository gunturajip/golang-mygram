package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Comment represents the model of a comment
type Comment struct {
	GormModel
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~Message of your comment is required"`
	UserID  uint   `gorm:"not null" json:"user_id"`
	PhotoID uint   `gorm:"not null" json:"photo_id" form:"photo_id" valid:"required~ID of your photo is required"`
}

func (p *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return err
	}

	err = nil
	return err
}

func (p *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return err
	}

	err = nil
	return err
}
