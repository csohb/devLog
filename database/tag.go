package database

import "gorm.io/gorm"

const TBNameTag = "TAG"

type TBTag struct {
	gorm.Model
	Tag string `gorm:"column:tag;type:varchar(20);comment:태그"`
}

func (TBTag) TableName() string {
	return TBNameTag
}
