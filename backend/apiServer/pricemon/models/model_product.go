package models

import (
	"errors"
	"time"
	
)

var _ = time.Thursday
//Product
type Product struct {
	
	Id   uint     `gorm:"column:id" form:"id" json:"id" comment:"" sql:"int,PRI"`
	Name   string     `gorm:"column:name" form:"name" json:"name" comment:"" sql:"varchar(256)"`
	Description   string     `gorm:"column:description" form:"description" json:"description" comment:"" sql:"varchar(256)"`
}
//TableName
func (m *Product) TableName() string {
	return "product"
}
//One
func (m *Product) One() (one *Product, err error) {
	one = &Product{}
	err = crudOne(m, one)
	return
}
//All
func (m *Product) All(q *PaginationQuery) (list *[]Product, total uint, err error) {
	list = &[]Product{}
	total, err = crudAll(m, q, list)
	return
}
//Update
func (m *Product) Update() (err error) {
	where := Product{Id: m.Id}
	m.Id = 0
	
	return crudUpdate(m, where)
}
//Create
func (m *Product) Create() (err error) {
	m.Id = 0
    
	return mysqlDB.Create(m).Error
}
//Delete
func (m *Product) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
