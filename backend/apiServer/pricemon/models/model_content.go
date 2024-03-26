package models

import (
	"errors"
	"time"
	
)

var _ = time.Thursday
//Content
type Content struct {
	
	Id   uint     `gorm:"column:id" form:"id" json:"id" comment:"" sql:"int,PRI"`
	Title   string     `gorm:"column:title" form:"title" json:"title" comment:"" sql:"varchar(256)"`
}
//TableName
func (m *Content) TableName() string {
	return "content"
}
//One
func (m *Content) One() (one *Content, err error) {
	one = &Content{}
	err = crudOne(m, one)
	return
}
//All
func (m *Content) All(q *PaginationQuery) (list *[]Content, total uint, err error) {
	list = &[]Content{}
	total, err = crudAll(m, q, list)
	return
}
//Update
func (m *Content) Update() (err error) {
	where := Content{Id: m.Id}
	m.Id = 0
	
	return crudUpdate(m, where)
}
//Create
func (m *Content) Create() (err error) {
	m.Id = 0
    
	return mysqlDB.Create(m).Error
}
//Delete
func (m *Content) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
