package database

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const TBNameUser = "USER"

type TBUser struct {
	ID        string      `gorm:"column:id;type:varchar(20);comment:아이디"`
	NickName  string      `gorm:"column:nick_name;type:varchar(20);comment:닉네임"`
	Password  string      `gorm:"column:password;type:varchar(100);comment:패스워드"`
	Introduce TBIntroduce `gorm:"foreignKey:user_id;references:id"`
	Career    []TBCareer  `gorm:"foreignKey:user_id;references:id"`
	Project   []TBProject `gorm:"foreignKey:user_id;references:id"`
}

func (t TBUser) TableName() string {
	return TBNameUser
}

func (t *TBUser) GetCareer(db *gorm.DB, userId string) error {
	return db.Preload("Career").Take(&t, "id = ?", userId).Error
}

func (t *TBUser) GetIntroduceForMainPage(db *gorm.DB) (ret []TBUser, err error) {
	if err = db.Preload("Introduce").Find(&ret).Error; err != nil {
		logrus.WithError(err).Error("get introduce for mainpage err")
		return nil, err
	}
	return ret, nil
}

func (t *TBUser) GetIntroduce(db *gorm.DB, userID string) (ret TBUser, err error) {
	if err = db.Model(&t).Preload("Project").
		Preload("Career").Preload("Introduce").
		Take(&ret, "id = ?", userID).Error; err != nil {
		return ret, err
	}

	return ret, nil
}
