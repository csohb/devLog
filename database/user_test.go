package database

import (
	"devLog/database/conn"
	"fmt"
	"testing"
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
