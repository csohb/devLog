package database

const TBNameUser = "USER"

type TBUser struct {
	ID       string `gorm:"column:id;type:varchar(20);comment:아이디"`
	NickName string `gorm:"column:nick_name;type:varchar(20);comment:닉네임"`
	Password string `gorm:"column:password;type:varchar(100);comment:패스워드"`
}

func (t TBUser) TableName() string {
	return TBNameUser
}
