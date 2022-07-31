package database

import (
	"gorm.io/gorm"
	"strconv"
	"time"
)

const TBNameProject = "PROJECT"

type TBProject struct {
	gorm.Model
	UserID      string    `gorm:"column:user_id;type:varchar(20);comment:아이디"`
	Name        string    `gorm:"column:name;type:varchar(20);comment:프로젝트명"`
	IsPersonal  bool      `gorm:"column:is_personal;type:boolean;comment:개인프로젝트 여부"`
	StartDate   time.Time `gorm:"column:start_date;type:datetime;comment:프로젝트 시작일"`
	EndDate     time.Time `gorm:"column:end_date;type:datetime;comment:프로젝트 완료일"`
	Description string    `gorm:"column:description;type:text;comment:프로젝트 설명"`
	Stack       []TBStack `gorm:"many2many:PROJECT_STACK"`
}

func (t TBProject) TableName() string {
	return TBNameProject
}

func (t *TBProject) Save(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t *TBProject) Get(db *gorm.DB, pid uint) error {
	return db.Model(&t).Preload("Stack").Take(&t, "id = ?", pid).Error
}

func (t *TBProject) Delete(db *gorm.DB, pid string) error {
	id, err := strconv.Atoi(pid)
	if err != nil {
		return err
	}
	return db.Model(&t).Delete("id = ?", id).Error
}

func (t *TBProject) Update(db *gorm.DB, stackList []TBStack) error {
	engine := db.Model(&t).Session(&gorm.Session{FullSaveAssociations: true})
	if err := engine.Updates(map[string]interface{}{
		"name":        t.Name,
		"is_personal": t.IsPersonal,
		"start_date":  t.StartDate,
		"end_date":    t.EndDate,
		"description": t.Description,
	}).Error; err != nil {
		return err
	}

	if err := engine.Association("Stack").Replace(stackList); err != nil {
		return err
	}
	return nil
}
