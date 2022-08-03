package database

import "gorm.io/gorm"

const TBNameStory = "STORY"

type TBStory struct {
	gorm.Model
}
