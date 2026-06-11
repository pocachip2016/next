package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

// check_list.status 값 상수 — 정수 리터럴 직접 사용 금지
const (
	StatusDetected       = 0 // 탐지=승인대기
	StatusNotified       = 1 // CP 통보 완료
	StatusDeleteConfirmed = 2 // 삭제 확인
	StatusClosed         = 3 // 종결
)

//CheckList
type CheckList struct {
	Id                 uint       `gorm:"column:id" form:"id" json:"id" comment:"" sql:"int,PRI"`
	ContentId          uint       `gorm:"column:content_id" form:"content_id" json:"content_id" comment:"" sql:"int,MUL"`
	PostId             string     `gorm:"column:post_id" form:"post_id" json:"post_id" comment:"" sql:"varchar(256)"`
	PostIdx            string     `gorm:"column:post_idx" form:"post_idx" json:"post_idx" comment:"" sql:"varchar(256)"`
	PostTxt            string     `gorm:"column:post_txt" form:"post_txt" json:"post_txt" comment:"" sql:"varchar(256)"`
	FirstTime          *time.Time `gorm:"column:first_time" form:"first_time" json:"first_time,omitempty" comment:"" sql:"timestamp"`
	UpdateTime         *time.Time `gorm:"column:update_time" form:"update_time" json:"update_time,omitempty" comment:"" sql:"timestamp"`
	Status             int        `gorm:"column:status" form:"status" json:"status" comment:"0=detected_pending,1=notified,2=delete_confirmed,3=closed" sql:"tinyint"`
	NotifiedAt         *time.Time `gorm:"column:notified_at" form:"notified_at" json:"notified_at,omitempty" comment:"" sql:"timestamp"`
	DeleteConfirmedAt  *time.Time `gorm:"column:delete_confirmed_at" form:"delete_confirmed_at" json:"delete_confirmed_at,omitempty" comment:"" sql:"timestamp"`
	ClosedAt           *time.Time `gorm:"column:closed_at" form:"closed_at" json:"closed_at,omitempty" comment:"" sql:"timestamp"`
}

//TableName
func (m *CheckList) TableName() string {
	return "check_list"
}

//One
func (m *CheckList) One() (one *CheckList, err error) {
	one = &CheckList{}
	err = crudOne(m, one)
	return
}

//All
func (m *CheckList) All(q *PaginationQuery) (list *[]CheckList, total uint, err error) {
	list = &[]CheckList{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *CheckList) Update() (err error) {
	where := CheckList{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

//Create
func (m *CheckList) Create() (err error) {
	m.Id = 0

	return mysqlDB.Create(m).Error
}

//Delete
func (m *CheckList) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
