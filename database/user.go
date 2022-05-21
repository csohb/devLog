package database

import "gorm.io/gorm"

const TBNameUser = "USER"

type TBUser struct {
	ID       string     `gorm:"column:id;type:varchar(20);comment:아이디"`
	NickName string     `gorm:"column:nick_name;type:varchar(20);comment:닉네임"`
	Password string     `gorm:"column:password;type:varchar(100);comment:패스워드"`
	Career   []TBCareer `gorm:"foreignKey:user_id;references:id"`
}

func (t TBUser) TableName() string {
	return TBNameUser
}

func (t *TBUser) GetCareer(db *gorm.DB, userId string) error {
	return db.Preload("Career").Take(&t, "id = ?", userId).Error
}
