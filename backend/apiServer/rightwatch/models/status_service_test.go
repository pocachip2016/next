package models

import (
	"errors"
	"testing"

	"rightwatch/services"
)

// TestIsAllowedTransition: IsAllowedTransition is in services/ — tested there.
// Here we verify that the Transition wrapper correctly rejects invalid pairs.
func TestTransitionInvalidReturnsError(t *testing.T) {
	cases := []struct{ from, to int }{
		{StatusDetected, StatusDeleteConfirmed},
		{StatusDetected, StatusClosed},
		{StatusClosed, StatusDetected},
		{99, 0},
	}
	for _, c := range cases {
		err := Transition(1, c.from, c.to)
		if err == nil {
			t.Errorf("Transition(1, %d, %d) expected error, got nil", c.from, c.to)
			continue
		}
		if !errors.Is(err, services.ErrInvalidTransition) {
			t.Errorf("Transition(1, %d, %d) error = %v, want ErrInvalidTransition", c.from, c.to, err)
		}
	}
}
