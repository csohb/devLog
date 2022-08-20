package database

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const TBNameStory = "STORY"

type TBStory struct {
	gorm.Model
	Title       string `gorm:"column:title;type:varchar(100);comment:글 제목"`
	Content     string `gorm:"column:content;type:text;comment:"`
	Description string `gorm:"column:description;type:varchar(100);comment:설명"`
	Writer      string `gorm:"column:writer;type:varchar(20);comment:작성자"`
	View        int    `gorm:"column:view;type:int(10);comment:조회수"`
}

func (TBStory) TableName() string {
	return TBNameStory
}

func (t TBStory) Find(db *gorm.DB, page, count int) ([]TBStory, int64, error) {
	var total int64
	var list []TBStory

	if err := db.Model(&t).Count(&total).Error; err != nil {
		logrus.WithError(err).Error("list count err")
		return nil, 0, err
	}

	if err := db.Model(&t).Order("created_at DESC").
		Offset(page * count).Limit(count).Find(&list).Error; err != nil {
		logrus.WithError(err).Error("list find err")
		return nil, 0, err
	}

	return list, total, nil
}

