package database

import (
	"devLog/database/conn"
	"testing"
)

func TestTechCreate(t *testing.T) {
	db, err := conn.ConnectForYJ()
	if err != nil {
		t.Error(err)
	}
	tb := TBTech{
		Name:       "GO",
		UserID:     "yujin",
		Percentage: 80,
	}
	if err = tb.Create(db); err != nil {
		t.Error(err)
	}
}
