package database

import (
	"gorm.io/gorm"
	"strconv"
)

const TBNameTech = "TECH"

type TBTech struct {
	gorm.Model
	Name       string `gorm:"column:name;type:varchar(20);comment:기술스택이름"`
	UserID     string `gorm:"column:user_id;type:varchar(20);comment:유저 아이디"`
	Percentage int    `gorm:"column:percentage;type:int(3);comment:숙련도"`
}

func (t TBTech) TableName() string {
	return TBNameTech
}

func (t TBTech) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t TBTech) Update(db *gorm.DB) error {
	err := db.Model(&t).Where("id = ?", t.ID).Updates(map[string]interface{}{
		"name":       t.Name,
		"percentage": t.Percentage,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (t TBTech) Delete(db *gorm.DB, id string) error {
	tid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return db.Model(&t).Delete("id = ?", tid).Error
}
