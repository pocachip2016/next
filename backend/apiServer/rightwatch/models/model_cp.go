package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

// Cp 저작권자 모델
type Cp struct {
	Id        uint      `gorm:"column:id" form:"id" json:"id" comment:"" sql:"int,PRI"`
	Name      string    `gorm:"column:name" form:"name" json:"name" comment:"" sql:"varchar(256)"`
	Email     string    `gorm:"column:email" form:"email" json:"email" comment:"" sql:"varchar(256)"`
	Phone     string    `gorm:"column:phone" form:"phone" json:"phone" comment:"" sql:"varchar(64)"`
	CreatedAt time.Time `gorm:"column:created_at" form:"created_at" json:"created_at,omitempty" comment:"" sql:"timestamp"`
}

// TableName 테이블명
func (m *Cp) TableName() string {
	return "cp"
}

// One 단건 조회
func (m *Cp) One() (one *Cp, err error) {
	one = &Cp{}
	err = crudOne(m, one)
	return
}

// All 목록 조회
func (m *Cp) All(q *PaginationQuery) (list *[]Cp, total uint, err error) {
	list = &[]Cp{}
	total, err = crudAll(m, q, list)
	return
}

// Update 수정
func (m *Cp) Update() (err error) {
	where := Cp{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

// Create 생성
func (m *Cp) Create() (err error) {
	m.Id = 0

	return mysqlDB.Create(m).Error
}

// Delete 삭제
func (m *Cp) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
