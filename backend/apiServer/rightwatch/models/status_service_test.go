package models

import (
	"errors"
	"testing"
)

func TestIsAllowedTransition(t *testing.T) {
	valid := []struct{ from, to int }{
		{StatusDetected, StatusNotified},
		{StatusNotified, StatusDeleteConfirmed},
		{StatusDeleteConfirmed, StatusClosed},
	}
	for _, c := range valid {
		if !IsAllowedTransition(c.from, c.to) {
			t.Errorf("expected %d→%d to be allowed", c.from, c.to)
		}
	}

	invalid := []struct{ from, to int }{
		{StatusDetected, StatusDeleteConfirmed},
		{StatusDetected, StatusClosed},
		{StatusNotified, StatusDetected},
		{StatusDeleteConfirmed, StatusDetected},
		{StatusDeleteConfirmed, StatusNotified},
		{StatusClosed, StatusDetected},
		{StatusClosed, StatusNotified},
		{99, 0},
	}
	for _, c := range invalid {
		if IsAllowedTransition(c.from, c.to) {
			t.Errorf("expected %d→%d to be forbidden", c.from, c.to)
		}
	}
}

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
		if !errors.Is(err, ErrInvalidTransition) {
			t.Errorf("Transition(1, %d, %d) error = %v, want ErrInvalidTransition", c.from, c.to, err)
		}
	}
}
