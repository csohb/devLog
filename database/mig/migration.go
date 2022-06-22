package main

import (
	"devLog/database"
	"devLog/database/conn"
)

func main() {

	db, err := conn.ConnectForTest()
	if err != nil {
		panic(err)
	}

	//db.AutoMigrate(&database.TBUser{}, &database.TBCareer{}, &database.TBProject{}, &database.TBTech{}, &database.TBIntroduce{})
	db.AutoMigrate(&database.TBBlog{})
	//db.AutoMigrate(&database.TBTech{})
}
