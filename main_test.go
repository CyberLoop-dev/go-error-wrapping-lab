package main

import (
	"errors"
	"testing"
)

func TestWrappedErrorCanBeInspected(t *testing.T) {
	err := loadSettings()
	if !errors.Is(err, ErrMissing) {
		t.Fatal("expected errors.Is to find ErrMissing through the chain")
	}
	var pathErr *PathError
	if !errors.As(err, &pathErr) {
		t.Fatal("expected errors.As to extract PathError")
	}
	if pathErr.Path != "settings.toml" {
		t.Fatalf("unexpected path: %s", pathErr.Path)
	}
}
