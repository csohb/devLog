package database

import (
	"gorm.io/gorm"
	"time"
)

const TBNameProject = "PROJECT"

type TBProject struct {
	gorm.Model
	UserID      string    `gorm:"column:user_id;type:varchar(20);comment:아이디"`
	Name        string    `gorm:"column:name;type:varchar(20);comment:프로젝트명"`
	IsPersonal  bool      `gorm:"column:is_personal;type:boolean;comment:개인프로젝트 여부"`
	StartDate   time.Time `gorm:"column:start_date;type:datetime;comment:프로젝트 시작일"`
	EndDate     time.Time `gorm:"column:end_date;type:datetime;comment:프로젝트 완료일"`
	Description string    `gorm:"column:description;type:text;comment:프로젝트 설명"`
	//Stack       []TBTech  `gorm:"foreignKey:project_id;references:id"`
}

func (t TBProject) TableName() string {
	return TBNameProject
}

func (t *TBProject) Save(db *gorm.DB) error {
	return db.Create(&t).Error
}
