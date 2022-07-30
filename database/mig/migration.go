package main

import (
	"devLog/database"
	"devLog/database/conn"
)

func main() {

	db, err := conn.ConnectForYJ()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&database.TBUser{}, &database.TBCareer{}, &database.TBProject{}, &database.TBTech{}, &database.TBIntroduce{})
	db.AutoMigrate(&database.TBBlog{})
	db.AutoMigrate(&database.TBTech{})
	db.AutoMigrate(&database.TBProject{})
	//db.AutoMigrate(&database.TBUser{})
	//db.AutoMigrate(&database.TBCareer{})
}
