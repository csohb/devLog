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
	Title       string  `gorm:"column:title;type:varchar(100);comment:글 제목"`
	Content     string  `gorm:"column:content;type:text;comment:"`
	Description string  `gorm:"column:description;type:varchar(100);comment:설명"`
	Writer      string  `gorm:"column:writer;type:varchar(20);comment:작성자"`
	View        int     `gorm:"column:view;type:int(10);comment:조회수"`
	Heart       int     `gorm:"column:heart;type:int(10);comment:하트수"`
	Tags        []TBTag `gorm:"many2many:BLOG_TAGS"`
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
	if err := engine.Order("created_at Desc").Offset(page * count).Limit(count).Find(&list).Error; err != nil {
		logrus.WithError(err).Error("list find err")
		return nil, 0, err
	}
	return list, total, nil
}

func (t TBBlog) SearchTags(db *gorm.DB, tag string, page, count int) ([]TBBlog, int64, error) {
	var total int64
	var list []TBBlog
	var err error

	var tags []TBTag
	if err = db.Model(&TBTag{}).Where("tag = ?", tag).Find(&tags).Error; err != nil {
		return list, total, err
	}

	tagIdList := make([]uint, len(tags))
	for i, v := range tags {
		tagIdList[i] = v.ID
	}

	blogIdList := make([]uint, len(tagIdList))
	if err = db.Model("blog_tags").Select("tb_blog_id", "tb_blog_id = ?", tagIdList).Error; err != nil {
		return nil, 0, err
	}

	if err = db.Model(&t).Where("id = ?", blogIdList).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err = db.Model(&t).Where("id = ?", blogIdList).
		Offset(page * count).Limit(count).
		Find(&list).Error; err != nil {

		return nil, 0, err
	}

	return list, total, nil
}

func (t *TBBlog) Get(db *gorm.DB, id string) error {
	bid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	fmt.Println("bid = ", bid)
	return db.Model(&t).Preload("Tags").Take(t, "id = ?", bid).Error
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

func (t *TBBlog) UpdateHeart(db *gorm.DB, id string, count int) error {
	bid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	t.ID = uint(bid)

	if err = db.Model(&t).Updates(map[string]interface{}{
		"heart": gorm.Expr("heart + ?", count),
	}).Error; err != nil {
		return err
	}

	return nil
}

func (t *TBBlog) UpdateView(db *gorm.DB, id string, count int) error {
	bid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	t.ID = uint(bid)

	if err = db.Model(&t).Updates(map[string]interface{}{
		"view": gorm.Expr("view + ?", count),
	}).Error; err != nil {
		return err
	}

	return nil
}

func (t *TBBlog) Update(db *gorm.DB, tagList []TBTag) error {

	engine := db.Model(&t).Session(&gorm.Session{FullSaveAssociations: true})

	if err := engine.Updates(map[string]interface{}{
		"title":   t.Title,
		"content": t.Content,
	}).Error; err != nil {
		return err
	}

	if err := engine.Association("Tags").Replace(tagList); err != nil {
		return err
	}
	return nil
}

func (t *TBBlog) BlogForMain(db *gorm.DB) (ret []TBBlog, err error) {
	if err = db.Model(&t).Preload("Tags").
		Order("created_at DESC").
		Limit(5).Find(&ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
}
