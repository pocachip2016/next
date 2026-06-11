package models

import (
	"time"

	"rightwatch/services"
)

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
