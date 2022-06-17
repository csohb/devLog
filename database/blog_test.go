package database

import (
	"devLog/database/conn"
	"fmt"
	"testing"
)

func TestBlog(t *testing.T) {
	db, err := conn.ConnectForYJ()
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
	tb := TBBlog{}
	/*tb.Get(db, "1")
	fmt.Println(tb.Title)*/
	list, total, err := tb.Find(db, 0, 10)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("list : ", list)
	fmt.Println("total : ", total)
}
