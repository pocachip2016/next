package models

import (
	"time"

	"rightwatch/services"
)

// LoadNotifyInfo 통보에 필요한 CP/콘텐츠 정보를 로드한다.
// check_list.id → kta_contents → cp 순서로 조회.
// CpId 없거나 이메일 없으면 빈 값 반환 — mailer가 dry-run으로 폴백.
func LoadNotifyInfo(checkListId uint) (cpName, cpEmail, contentTitle, postURL string, err error) {
	cl := &CheckList{Id: checkListId}
	clRow, err := cl.One()
	if err != nil {
		return
	}
	postURL = clRow.PostTxt

	content := &KtaContent{Id: clRow.ContentId}
	contentRow, err := content.One()
	if err != nil {
		return
	}
	contentTitle = contentRow.Title

	if contentRow.CpId == nil {
		return
	}
	cp := &Cp{Id: *contentRow.CpId}
	cpRow, err := cp.One()
	if err != nil {
		return
	}
	cpName, cpEmail = cpRow.Name, cpRow.Email
	return
}

// PostExists post_idx로 post 테이블을 조회. false면 게시물이 삭제된 것으로 판정.
func PostExists(postIdx string) bool {
	var count int
	mysqlDB.Model(&Post{}).Where("idx = ?", postIdx).Count(&count)
	return count > 0
}

// Transition check_list 행의 status를 from→to로 전이한다.
// 허용되지 않은 전이는 services.ErrInvalidTransition 반환.
// WHERE "id=? AND status=?" 조건으로 race condition 방어.
func Transition(id uint, from, to int) error {
	if !services.IsAllowedTransition(from, to) {
		return services.TransitionError(from, to)
	}
	now := time.Now()
	updates := map[string]interface{}{"status": to}
	switch to {
	case StatusNotified:
		updates["notified_at"] = now
	case StatusDeleteConfirmed:
		updates["delete_confirmed_at"] = now
	case StatusClosed:
		updates["closed_at"] = now
	}
	return mysqlDB.Model(&CheckList{}).
		Where("id = ? AND status = ?", id, from).
		Updates(updates).Error
}
