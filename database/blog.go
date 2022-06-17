package database

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
)

const TBNameBlog = "BLOG"

type TBBlog struct {
	gorm.Model
	Title   string  `gorm:"column:title;type:varchar(100);comment:글 제목"`
	Content string  `gorm:"column:content;type:varchar(65535);comment:"`
	Writer  string  `gorm:"column:writer;type:varchar(20);comment:작성자"`
	View    int     `gorm:"column:view;type:int(10);comment:조회수"`
	Heart   int     `gorm:"column:heart;type:int(10);comment:하트수"`
	Tags    []TBTag `gorm:"many2many:BLOG_TAGS"`
}

func (TBBlog) TableName() string {
	return TBNameBlog
}

func (t TBBlog) Find(db *gorm.DB, page, count int) ([]TBBlog, int64, error) {
	var total int64
	var list []TBBlog

	engine := db.Model(&t).Preload("Tags")
	if err := engine.Count(&total).Error; err != nil {
		logrus.WithError(err).Error("list count err")
		return nil, 0, err
	}
	if err := engine.Offset(page * count).Limit(count).Find(&list).Error; err != nil {
		logrus.WithError(err).Error("list find err")
		return nil, 0, err
	}
	return list, total, nil
}

func (t TBBlog) Get(db *gorm.DB, id string) error {
	bid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	fmt.Println("bid = ", bid)
	return db.Model(&t).Take(&t, "writer = ?", "csohb").Error
}

func (t *TBBlog) Save(db *gorm.DB) error {
	return db.Model(&t).Create(&t).Error
}

func (t *TBBlog) Delete(db *gorm.DB, id string) error {
	bid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return db.Model(&t).Delete("id = ?", bid).Error
}

func (t *TBBlog) Update(db *gorm.DB) error {
	return db.Model(&t).Where("id = ?", t.ID).Updates(map[string]interface{}{
		"title":   t.Title,
		"content": t.Content,
		"tag":     t.Tags,
	}).Error
}
