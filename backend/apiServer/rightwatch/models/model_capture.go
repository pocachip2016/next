package models

import (
	"time"
)

// Capture: check_list 항목별 화면 캡처 히스토리 (M7)
// 레거시 Screenshots(url_md5 PK)는 단일 캡처용이라 분리한 신규 모델.
type Capture struct {
	Id          uint       `gorm:"column:id" form:"id" json:"id" comment:"" sql:"int,PRI"`
	CheckListId uint       `gorm:"column:check_list_id" form:"check_list_id" json:"check_list_id" comment:"" sql:"int,MUL"`
	CaptureType string     `gorm:"column:capture_type" form:"capture_type" json:"capture_type" comment:"list|detail|takedown" sql:"varchar(16)"`
	FilePath    string     `gorm:"column:file_path" form:"file_path" json:"file_path" comment:"" sql:"varchar(512)"`
	PageUrl     string     `gorm:"column:page_url" form:"page_url" json:"page_url" comment:"" sql:"varchar(512)"`
	StatusAt    int        `gorm:"column:status_at" form:"status_at" json:"status_at" comment:"" sql:"tinyint"`
	TriggerBy   string     `gorm:"column:trigger_by" form:"trigger_by" json:"trigger_by" comment:"auto|manual" sql:"varchar(8)"`
	CapturedAt  *time.Time `gorm:"column:captured_at" form:"captured_at" json:"captured_at,omitempty" comment:"" sql:"timestamp"`
}

// TableName
func (m *Capture) TableName() string {
	return "capture"
}

// Create inserts a new capture row and populates m.Id with the new auto-increment id.
func (m *Capture) Create() error {
	m.Id = 0
	return mysqlDB.Create(m).Error
}

// One loads a capture by its set fields (typically Id).
func (m *Capture) One() (one *Capture, err error) {
	one = &Capture{}
	err = crudOne(m, one)
	return
}

// SaveFilePath updates only the file_path column for this capture id.
// 캡처 row를 먼저 insert해 id를 채번한 뒤, {id}.jpg 경로를 역으로 기록할 때 사용.
func (m *Capture) SaveFilePath(path string) error {
	return mysqlDB.Model(&Capture{}).Where("id = ?", m.Id).Update("file_path", path).Error
}

// CapturesByCheckList returns all captures for a check_list item, oldest first.
func CapturesByCheckList(checkListId uint) ([]Capture, error) {
	var list []Capture
	err := mysqlDB.Where("check_list_id = ?", checkListId).Order("captured_at asc").Find(&list).Error
	return list, err
}

// LoadCaptureTarget joins check_list → post (post.idx = check_list.post_idx) to
// obtain the metadata needed to build ondisk capture URLs. itemURL/cat1/cat2 may
// be empty when no matching post row exists — the caller falls back to an
// idx-based detail URL.
func LoadCaptureTarget(checkListId uint) (postIdx, itemURL, cat1, cat2 string, err error) {
	cl := &CheckList{Id: checkListId}
	clRow, err := cl.One()
	if err != nil {
		return
	}
	postIdx = clRow.PostIdx

	var post Post
	if e := mysqlDB.Where("idx = ?", clRow.PostIdx).First(&post).Error; e == nil {
		itemURL, cat1, cat2 = post.ItemUrl, post.Cat1Code, post.Cat2Code
	}
	return
}
