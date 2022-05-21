package database

import (
	"devLog/database/conn"
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	db, err := conn.ConnectForTest()
	if err != nil {
		t.Error(err)
	}
	tb := TBUser{}
	tb.GetCareer(db, "yujin")
	fmt.Println(tb)
}
