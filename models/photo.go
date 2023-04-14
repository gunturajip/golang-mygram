package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Photo represents the model of a photo
type Photo struct {
	GormModel
	Title    string    `gorm:"not null" json:"title" form:"title" valid:"required~Title of your photo is required"`
	Caption  string    `json:"caption" form:"caption"`
	PhotoURL string    `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~URL of your photo is required"`
	UserID   uint      `gorm:"not null" json:"user_id"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return err
	}

	err = nil
	return err
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return err
	}

	err = nil
	return err
}
