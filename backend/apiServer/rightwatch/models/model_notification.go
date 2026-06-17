package models

import "time"

// Notification 통보 발송 로그 (notification_log 테이블)
type Notification struct {
	Id          uint      `gorm:"column:id" form:"id" json:"id" sql:"int,PRI"`
	CheckListId uint      `gorm:"column:check_list_id" form:"check_list_id" json:"check_list_id" sql:"int,MUL"`
	CpId        *uint     `gorm:"column:cp_id" form:"cp_id" json:"cp_id,omitempty" sql:"int"`
	ToEmail     string    `gorm:"column:to_email" form:"to_email" json:"to_email" sql:"varchar(256)"`
	Subject     string    `gorm:"column:subject" form:"subject" json:"subject" sql:"varchar(256)"`
	Body        string    `gorm:"column:body" form:"body" json:"body" sql:"text"`
	DryRun      int       `gorm:"column:dry_run" form:"dry_run" json:"dry_run" sql:"tinyint"`
	Result      string    `gorm:"column:result" form:"result" json:"result" sql:"varchar(256)"`
	SentAt      time.Time `gorm:"column:sent_at" form:"sent_at" json:"sent_at" sql:"timestamp"`
}

// TableName 테이블명
func (m *Notification) TableName() string {
	return "notification_log"
}

// Create 발송 로그 적재
func (m *Notification) Create() error {
	m.Id = 0
	return mysqlDB.Create(m).Error
}
