package models

import (
	"errors"
	"time"
	
)

var _ = time.Thursday
//PriceList
type PriceList struct {
	
	Id   uint     `gorm:"column:id" form:"id" json:"id" comment:"" sql:"int,PRI"`
	PriceName   string     `gorm:"column:price_name" form:"price_name" json:"price_name" comment:"" sql:"varchar(20)"`
	ProductPrice   int     `gorm:"column:product_price" form:"product_price" json:"product_price" comment:"" sql:"int"`
	ProductLine   int     `gorm:"column:product_line" form:"product_line" json:"product_line" comment:"" sql:"int"`
}
//TableName
func (m *PriceList) TableName() string {
	return "price_list"
}
//One
func (m *PriceList) One() (one *PriceList, err error) {
	one = &PriceList{}
	err = crudOne(m, one)
	return
}
//All
func (m *PriceList) All(q *PaginationQuery) (list *[]PriceList, total uint, err error) {
	list = &[]PriceList{}
	total, err = crudAll(m, q, list)
	return
}
//Update
func (m *PriceList) Update() (err error) {
	where := PriceList{Id: m.Id}
	m.Id = 0
	
	return crudUpdate(m, where)
}
//Create
func (m *PriceList) Create() (err error) {
	m.Id = 0
    
	return mysqlDB.Create(m).Error
}
//Delete
func (m *PriceList) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
