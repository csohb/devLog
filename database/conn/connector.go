package conn

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectForMariaDB(dsn string) (*gorm.DB, error) {
	mysql, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		logrus.WithError(err).Error("DB Connect error")
		panic(err)
	}
	return mysql, nil
}

func ConnectForTest() (*gorm.DB, error) {
	dsn := "root:Jenoayo12!@tcp(127.0.0.1:3306)/devlog?charset=utf8mb4&parseTime=True&loc=Local"
	mysql, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		logrus.WithError(err).Error("DB Connect error")
		panic(err)
	}
	return mysql, nil
}

func ConnectForYJ() (*gorm.DB, error) {
	dsn := "rladudwl210:kk5371826@tcp(uws7-111.cafe24.com:3306)/rladudwl210?charset=utf8mb4&parseTime=True&loc=Local"
	mysql, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		logrus.WithError(err).Error("DB Connect error")
		panic(err)
	}
	return mysql, nil
}
