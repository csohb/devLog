package database

import (
	"devLog/database/conn"
	"fmt"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	db, err := conn.ConnectForYJ()
	if err != nil {
		t.Error(err)
	}
	tb := TBUser{}
	/*tb := TBCareer{}

	err = db.Model(&tb).Delete(&tb, "company_name IS NULL").Error
	if err != nil {
		t.Error(err)
	}*/
	tb.GetCareer(db, "yujin")
	fmt.Println(tb)

	/*	career := make([]TBCareer, 2)
		career[0] = TBCareer{
			UserID:      "yujin",
			CompanyName: "TUV Rheinland Korea",
			StartDate:   time.Date(2019, 11, 11, 00, 00, 00, 00, time.UTC),
			EndDate:     time.Date(2020, 07, 30, 00, 00, 00, 00, time.UTC),
			JobTitle:    "Technical Support",
			JobDetail:   "GA, Annual Fee Handling",
		}
		fmt.Println(career[0])

		/*user := TBUser{
			ID:       "yujin",
			NickName: "csohb",
			Password: "Kawaii",
			Career:   career,
		}*/
	/*err = db.Model(&tb).Create(user).Error
	if err != nil {
		t.Error(err)
	}*/
}

func TestUserIntroduce(t *testing.T) {
	db, err := conn.ConnectForTest()
	if err != nil {
		t.Error(err)
	}
	tb := TBUser{}

	data, err := tb.GetIntroduce(db, "yujin")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("data : %+v", data)
}

func TestUpdateIntroduce(t *testing.T) {
	db, err := conn.ConnectForTest()
	if err != nil {
		t.Error(err)
	}
	tb := TBUser{
		ID: "yujin",
		Introduce: TBIntroduce{
			UserID:     "yujin",
			Intro:      "",
			ProfileUrl: "",
		},
		Career: []TBCareer{
			{
				Model:       gorm.Model{},
				UserID:      "",
				CompanyName: "",
				StartDate:   time.Time{},
				EndDate:     time.Time{},
				JobTitle:    "",
				JobDetail:   "",
			}, {
				Model:       gorm.Model{},
				UserID:      "",
				CompanyName: "",
				StartDate:   time.Time{},
				EndDate:     time.Time{},
				JobTitle:    "",
				JobDetail:   "",
			},
		},
		Project: []TBProject{
			{
				Model:       gorm.Model{},
				UserID:      "",
				Name:        "",
				IsPersonal:  false,
				StartDate:   time.Time{},
				EndDate:     time.Time{},
				Description: "",
			}, {
				Model:       gorm.Model{},
				UserID:      "",
				Name:        "",
				IsPersonal:  false,
				StartDate:   time.Time{},
				EndDate:     time.Time{},
				Description: "",
			},
		},
	}

	if err = tb.UpdateIntroduce(db, tb.ID); err != nil {
		t.Error(err)
	}

}
