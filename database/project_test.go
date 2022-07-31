package database

import (
	"devLog/database/conn"
	"fmt"
	"testing"
	"time"
)

func TestProject(t *testing.T) {
	db, err := conn.ConnectForTest()
	if err != nil {
		t.Error(err)
	}

	tb := TBProject{
		UserID:      "yeong",
		Name:        "개발자의 개발 log, devLog",
		IsPersonal:  true,
		StartDate:   time.Date(2022, 4, 1, 0, 0, 0, 0, time.Local),
		EndDate:     time.Now().Add(time.Hour * 24 * 30),
		Description: "개발자의 매일을 기록하다. devLog",
	}

	if err = tb.Save(db); err != nil {
		t.Error(t)
	}
}

func TestProjectGet(t *testing.T) {
	db, err := conn.ConnectForYJ()
	if err != nil {
		t.Error(err)
	}
	tb := TBProject{}
	tb.Get(db, 4)
	fmt.Println("tb : ", tb)
}

func TestTech(t *testing.T) {
	db, err := conn.ConnectForYJ()
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&TBTech{})
	/*var Tech = []TBTech{
		{
			Name: "Go",
		}, {
			Name: "gorm",
		}, {
			Name: "echo framework",
		}, {
			Name: "svelte",
		}, {
			Name: "typescript",
		}, {
			Name: "javascript",
		}, {
			Name: "Python",
		}, {
			Name: "Django",
		}, {
			Name: "Vue.js",
		},
	}

	for _, v := range Tech {
		tb := TBTech{
			Name: v.Name,
		}
		if err = tb.Create(db); err != nil {
			t.Error(err)
		}
	}*/

}
