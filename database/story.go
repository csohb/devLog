package database

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
)

const TBNameStory = "STORY"

type TBStory struct {
	gorm.Model
	Title       string `gorm:"column:title;type:varchar(100);comment:글 제목"`
	Content     string `gorm:"column:content;type:text;comment:content"`
	Type        string `gorm:"column:type;type:varchar(1);comment:F-frontEnd, B-backEnd"`
	Image       string `gorm:"column:image;type:text;comment:이미지 url":`
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

	engine := db.Model(&t)
	if err := engine.Count(&total).Error; err != nil {
		logrus.WithError(err).Error("list count err")
		return nil, 0, err
	}

	if err := engine.Order("created_at DESC").Offset(page * count).Limit(count).Find(&list).Error; err != nil {
		logrus.WithError(err).Error("list find err")
		return nil, 0, err
	}

	fmt.Println("list : ", list)
	return list, total, nil
}

func (t *TBStory) Get(db *gorm.DB, id string) error {
	bid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return db.Model(&t).Take(t, "id = ?", bid).Error
}

func (t *TBStory) Save(db *gorm.DB) error {
	return db.Model(&t).Create(&t).Error
}

func (t *TBStory) Delete(db *gorm.DB, id string) error {
	bid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return db.Model(&t).Delete("id = ?", bid).Error
}

func (t *TBStory) Update(db *gorm.DB) error {
	return db.Model(&t).Updates(map[string]interface{}{
		"title":       t.Title,
		"content":     t.Content,
		"type":        t.Type,
		"description": t.Description,
		"image":       t.Image,
	}).Error
}

func (t *TBStory) UpdateView(db *gorm.DB, id string) error {
	bid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	t.ID = uint(bid)

	if err = db.Model(&t).Updates(map[string]interface{}{
		"view": gorm.Expr("view + ?", 1),
	}).Error; err != nil {
		return err
	}

	return nil
}
