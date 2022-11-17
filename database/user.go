package database

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

const TBNameUser = "USER"

type TBUser struct {
	ID        string      `gorm:"column:id;type:varchar(20);comment:아이디"`
	Name      string      `gorm:"column:name;type:varchar(10);comment:이름"`
	NickName  string      `gorm:"column:nick_name;type:varchar(20);comment:닉네임"`
	Password  string      `gorm:"column:password;type:varchar(100);comment:패스워드"`
	Email     string      `gorm:"column:email;type:varchar(100);comment:이메일"`
	Addr      string      `gorm:"column:addr;type:varchar(100);comment:주소"`
	Birthday  time.Time   `gorm:"column:birthday;type:datetime;comment:생년월일"`
	Introduce TBIntroduce `gorm:"foreignKey:user_id;references:id"`
	Career    []TBCareer  `gorm:"foreignKey:user_id;references:id"`
	Project   []TBProject `gorm:"foreignKey:user_id;references:id"`
	Tech      []TBTech    `gorm:"foreignKey:user_id;references:id"`
}

func (t TBUser) TableName() string {
	return TBNameUser
}

func (t TBUser) CheckID(db *gorm.DB, userID string) error {
	if err := db.Model(&t).Take(&t, "id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}
	return nil
}

func (t TBUser) CheckPwd(db *gorm.DB, userID, pwd string) error {
	if err := db.Model(&t).Take(&t, "id = ?", userID).Error; err != nil {
		logrus.WithError(err)
	}
	if t.Password != pwd {
		return gorm.ErrInvalidData
	} else {
		return nil
	}
}

func (t *TBUser) UpdateIntroduce(db *gorm.DB, userId string) error {
	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&t).Where("id = ?", t.ID).Updates(map[string]interface{}{
			"email": t.Email,
			"addr":  t.Addr,
		}).Error; err != nil {
			return err
		}

		if err := tx.Model(&t.Introduce).Where("user_id = ?", userId).Updates(map[string]interface{}{
			"intro":       t.Introduce.Intro,
			"profile_url": t.Introduce.ProfileUrl,
		}).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}
	return nil
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
	if err = db.Model(&t).Preload("Introduce").
		Preload("Project").
		Preload("Career").
		Preload("Tech").
		Take(&ret, "id = ?", userID).Error; err != nil {
		return ret, err
	}

	return ret, nil
}
