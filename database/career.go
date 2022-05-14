package database

import (
	"gorm.io/gorm"
	"time"
)

const TBNameCareer = "CAREER"

type TBCareer struct {
	gorm.Model
	CompanyName string    `gorm:"column:company_name;type:varchar(50);comment:회사이름"`
	StartDate   time.Time `gorm:"column:start_date;type:datetime;comment:입사일"`
	EndDate     time.Time `gorm:"column:end_date;type:datetime;comment:퇴사일"`
	JobTitle    string    `gorm:"column:job_title;type:varchar(30);comment:직무 타이틀"`
	JobDetail   string    `gorm:"column:JOB_DETAIL;type:varchar(400);comment:직무소개"`
}

func (t TBCareer) TableName() string {
	return TBNameCareer
}
