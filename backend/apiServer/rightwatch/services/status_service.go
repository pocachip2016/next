package services

import (
	"errors"
	"fmt"
)

var ErrInvalidTransition = errors.New("invalid status transition")

// AllowedTransitions: from → only allowed next state
// 0=detected_pending → 1=notified → 2=delete_confirmed → 3=closed
var AllowedTransitions = map[int]int{0: 1, 1: 2, 2: 3}

// IsAllowedTransition from→to 전이가 허용되는지 검사 (pure, DB 없음)
func IsAllowedTransition(from, to int) bool {
	expected, ok := AllowedTransitions[from]
	return ok && expected == to
}

// TransitionError 허용되지 않은 전이에 대한 에러를 생성한다
func TransitionError(from, to int) error {
	return fmt.Errorf("%w: %d→%d", ErrInvalidTransition, from, to)
}
