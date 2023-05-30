package apperror_test

import (
	"errors"
	"github.com/ensiouel/apperror"
	"testing"
)

func TestNew(t *testing.T) {
	err := apperror.New(apperror.NotFound)

	if err.Message != "not found" {
		t.Errorf("expected %s, got %s", "not found", err.Message)
	}

	if err.Code != apperror.NotFound {
		t.Errorf("expected %d, got %d", apperror.NotFound, err.Code)
	}
}

func TestError_WithMessage(t *testing.T) {
	err := apperror.New(apperror.NotFound).WithMessage("user not found")

	if err.Message != "user not found" {
		t.Errorf("expected %s, got %s", "user not found", err.Message)
	}
}

func TestError_WithError(t *testing.T) {
	err := apperror.New(apperror.Internal).WithError(errors.New("timeout"))

	if err.Error() != "timeout" {
		t.Errorf("expected %s, got %s", "timeout", err.Message)
	}
}
