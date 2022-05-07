package database

const TBNameTech = "TECH"

type TBTech struct {
	Name      string `gorm:"column:name;type:varchar(20);comment:기술스택이름"`
	ProjectID uint   `gorm:"column:project_id;comment:프로젝트 아이디"`
}

func (t TBTech) TableName() string {
	return TBNameTech
}
