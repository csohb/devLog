package database

import "gorm.io/gorm"

const TBNameIntroduce = "INTRODUCE"

type TBIntroduce struct {
	gorm.Model
	UserID     string `gorm:"column:user_id;type:varchar(20);comment:유저 아이디"`
	Developer  string `gorm:"column:developer;type:char(1);comment:F-front developer, B-backend developer"`
	Intro      string `gorm:"column:intro;type:varchar(500);comment:자기소개멘트"`
	ProfileUrl string `gorm:"column:profile_url;type:varchar(200);comment:프로필사진"`
}

func (t TBIntroduce) TableName() string {
	return TBNameIntroduce
}
