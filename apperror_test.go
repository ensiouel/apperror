package apperror_test

import (
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
