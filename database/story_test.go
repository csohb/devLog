package database

import (
	"devLog/database/conn"
	"testing"
)

func TestStory(t *testing.T) {
	db, err := conn.ConnectForYJ()
	if err != nil {
		t.Error(err)
	}

	tb := TBStory{
		Title:       "개발 블로그를 만들게된 스토리",
		Content:     "블라블라",
		Description: "디스크립션입니다",
		Writer:      "yujin",
		View:        0,
	}

	if err = tb.Save(db); err != nil {
		t.Error(err)
	}
}
