package database

import "gorm.io/gorm"

const TBNameStack = "STACK"

type TBStack struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(20);comment:기술스택이름"`
}

func (t TBStack) TableName() string {
	return TBNameStack
}

func (t TBStack) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}
