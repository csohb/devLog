package database

import (
	"gorm.io/gorm"
	"time"
)

const TBNameCareer = "CAREER"

type TBCareer struct {
	gorm.Model
	UserID      string    `gorm:"column:user_id;type:varchar(20);comment:유저 아이디"`
	CompanyName string    `gorm:"column:company_name;type:varchar(50);comment:회사이름"`
	StartDate   time.Time `gorm:"column:start_date;type:datetime;comment:입사일"`
	EndDate     time.Time `gorm:"column:end_date;type:datetime;comment:퇴사일"`
	JobTitle    string    `gorm:"column:job_title;type:varchar(30);comment:직무 타이틀"`
	JobDetail   string    `gorm:"column:job_title;type:varchar(400);comment:직무소개"`
}

func (t TBCareer) TableName() string {
	return TBNameCareer
}

func (t TBCareer) DeleteCareer(db *gorm.DB, ids []int) error {
	return db.Delete(&t, ids).Error
}

func (t TBCareer) UpdateCareer(db *gorm.DB, career TBCareer) error {
	return db.Model(&t).Updates(map[string]interface{}{
		"company_name": career.CompanyName,
		"start_date":   career.StartDate,
		"end_date":     career.EndDate,
		"job_title":    career.JobTitle,
		"job_detail":   career.JobDetail,
	}).Error
}
