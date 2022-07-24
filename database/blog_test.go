package database

import (
	"devLog/database/conn"
	"fmt"
	"gorm.io/gorm"
	"testing"
)

func TestBlog(t *testing.T) {
	db, err := conn.ConnectForTest()
	if err != nil {
		t.Error(err)
	}
	//db.AutoMigrate(&TBBlog{}, &TBTag{})
	/*var Tag = []TBTag{
		{Tag: "golang"}, {Tag: "programming"},
	}

	tb := TBBlog{
		Title:   "첫번째 블로깅",
		Content: "golang 개발자들은 코로나도 피해간다는데 사실인가요?",
		Writer:  "csohb",
		View:    0,
		Heart:   0,
		Tags:    Tag,
	}
	if err = tb.Save(db); err != nil {
		t.Error(err)
	}
	*/
	tb := TBBlog{
		Model: gorm.Model{
			ID: 1,
		},
	}
	/*tb.Get(db, "1")
	fmt.Println(tb)*/

	if err = db.Model(&tb).Updates(map[string]interface{}{
		"heart": gorm.Expr("heart + ?", -1),
	}).Error; err != nil {
		t.Error(err)
	}

	fmt.Println(tb.Heart)

	/*if err = tb.UpdateHeart(db, "1", 1); err != nil {
		t.Error(err)
	}
	fmt.Println(tb.Heart)*/
	/*list, total, err := tb.Find(db, 0, 10)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("list : ", list)
	fmt.Println("total : ", total)*/
}

func TestBlogUpdate(t *testing.T) {
	db, err := conn.ConnectForTest()
	if err != nil {
		t.Error(err)
	}
	/*tags := []TBTag{
		{
			Tag: "test",
		}, {
			Tag: "golang",
		}, {
			Tag: "backend",
		},
	}

	tb := TBBlog{
		Title:   "test",
		Content: "test",
		Writer:  "csohb",
		View:    0,
		Heart:   0,
		Tags:    tags,
	}

	if err = tb.Save(db); err != nil {
		t.Error(err)
	}*/
	tags := []TBTag{
		{
			Tag: "u-test",
		}, {
			Tag: "u-golang",
		}, {
			Tag: "u-backend",
		},
	}
	tb := TBBlog{
		Model: gorm.Model{
			ID: 1,
		},
		Title:   "update-test",
		Content: "update-test",
	}

	if err = tb.Update(db, tags); err != nil {
		t.Error(err)
	}
}

func TestBlogGet(t *testing.T) {
	db, err := conn.ConnectForTest()
	if err != nil {
		t.Error(err)
	}
	tb := TBBlog{}
	if err := tb.Get(db, "1"); err != nil {
		t.Error(err)
	}

	fmt.Println(tb.Tags)
}

func TestBlogTag(t *testing.T) {
	db, err := conn.ConnectForYJ()
	if err != nil {
		t.Error(err)
	}

	tb := TBBlog{}
	list, count, err := tb.SearchTags(db, "golang", 0, 10)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("list : ", list)
	fmt.Println("count : ", count)

}
