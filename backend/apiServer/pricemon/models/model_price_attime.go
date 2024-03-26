package models

import (
	"errors"
	"time"
	
)

var _ = time.Thursday
//PriceAttime
type PriceAttime struct {
	
	Id   uint     `gorm:"column:id" form:"id" json:"id" comment:"" sql:"int,PRI"`
	ContentId   int     `gorm:"column:content_id" form:"content_id" json:"content_id" comment:"" sql:"int,MUL"`
	PriceId   int     `gorm:"column:price_id" form:"price_id" json:"price_id" comment:"" sql:"int,MUL"`
	StartDate   *time.Time     `gorm:"column:start_date" form:"start_date" json:"start_date,omitempty" comment:"" sql:"datetime"`
	EndDate   *time.Time     `gorm:"column:end_date" form:"end_date" json:"end_date,omitempty" comment:"" sql:"datetime"`
}
//TableName
func (m *PriceAttime) TableName() string {
	return "price_attime"
}
//One
func (m *PriceAttime) One() (one *PriceAttime, err error) {
	one = &PriceAttime{}
	err = crudOne(m, one)
	return
}
//All
func (m *PriceAttime) All(q *PaginationQuery) (list *[]PriceAttime, total uint, err error) {
	list = &[]PriceAttime{}
	total, err = crudAll(m, q, list)
	return
}
//Update
func (m *PriceAttime) Update() (err error) {
	where := PriceAttime{Id: m.Id}
	m.Id = 0
	
	return crudUpdate(m, where)
}
//Create
func (m *PriceAttime) Create() (err error) {
	m.Id = 0
    
	return mysqlDB.Create(m).Error
}
//Delete
func (m *PriceAttime) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
