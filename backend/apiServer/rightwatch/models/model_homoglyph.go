package models

import "errors"

// Homoglyph 변형 문자 정규화 사전 엔트리
type Homoglyph struct {
	Id       uint   `gorm:"column:id" form:"id" json:"id" sql:"int,PRI"`
	FromChar string `gorm:"column:from_char" form:"from_char" json:"from_char" sql:"varchar(16)"`
	ToChar   string `gorm:"column:to_char" form:"to_char" json:"to_char" sql:"varchar(16)"`
}

// TableName 테이블명
func (m *Homoglyph) TableName() string {
	return "homoglyph_map"
}

// One 단건 조회
func (m *Homoglyph) One() (one *Homoglyph, err error) {
	one = &Homoglyph{}
	err = crudOne(m, one)
	return
}

// All 목록 조회
func (m *Homoglyph) All(q *PaginationQuery) (list *[]Homoglyph, total uint, err error) {
	list = &[]Homoglyph{}
	total, err = crudAll(m, q, list)
	return
}

// Create 생성
func (m *Homoglyph) Create() (err error) {
	m.Id = 0
	return mysqlDB.Create(m).Error
}

// Update 수정
func (m *Homoglyph) Update() (err error) {
	where := Homoglyph{Id: m.Id}
	m.Id = 0
	return crudUpdate(m, where)
}

// Delete 삭제
func (m *Homoglyph) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
