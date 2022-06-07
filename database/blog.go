package database

import "gorm.io/gorm"

const TBNameBlog = "BLOG"

type TBBlog struct {
	gorm.Model
	Title   string  `gorm:"column:title;type:varchar(100);comment:글 제목"`
	Content string  `gorm:"column:content;type:varchar(65535);comment:"`
	Writer  string  `gorm:"column:writer;type:varchar(20);comment:작성자"`
	View    int     `gorm:"column:view;type:int(10);comment:조회수"`
	Heart   int     `gorm:"column:heart;type:int(10);comment:하트수"`
	Tags    []TBTag `gorm:"many2many:blog_tags"`
}

func (TBBlog) TableName() string {
	return TBNameBlog
}
