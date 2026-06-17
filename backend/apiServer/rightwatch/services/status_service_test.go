package services

import (
	"errors"
	"testing"
)

func TestIsAllowedTransition(t *testing.T) {
	valid := []struct{ from, to int }{
		{0, 1}, // detected_pending → notified
		{1, 2}, // notified → delete_confirmed
		{2, 3}, // delete_confirmed → closed
	}
	for _, c := range valid {
		if !IsAllowedTransition(c.from, c.to) {
			t.Errorf("expected %d→%d to be allowed", c.from, c.to)
		}
	}

	invalid := []struct{ from, to int }{
		{0, 2}, {0, 3},
		{1, 0}, {2, 0}, {2, 1},
		{3, 0}, {3, 1}, {3, 2},
		{99, 0},
	}
	for _, c := range invalid {
		if IsAllowedTransition(c.from, c.to) {
			t.Errorf("expected %d→%d to be forbidden", c.from, c.to)
		}
	}
}

func TestTransitionError(t *testing.T) {
	cases := []struct{ from, to int }{
		{0, 2}, {0, 3}, {3, 0}, {99, 0},
	}
	for _, c := range cases {
		err := TransitionError(c.from, c.to)
		if err == nil {
			t.Errorf("TransitionError(%d, %d) returned nil", c.from, c.to)
			continue
		}
		if !errors.Is(err, ErrInvalidTransition) {
			t.Errorf("TransitionError(%d, %d) = %v, want ErrInvalidTransition", c.from, c.to, err)
		}
	}
}
