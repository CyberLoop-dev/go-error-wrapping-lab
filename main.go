package main

import (
	"errors"
	"flag"
	"fmt"
)

var ErrMissing = errors.New("resource missing")

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	return fmt.Sprintf("%s %q: %v", e.Op, e.Path, e.Err)
}

func (e *PathError) Unwrap() error { return e.Err }

func loadSettings() error {
	pathErr := &PathError{Op: "open", Path: "settings.toml", Err: ErrMissing}
	return fmt.Errorf("load application settings: %w", pathErr)
}

func main() {
	show := flag.Bool("show-chain", false, "print the full wrapped error")
	flag.Parse()

	err := loadSettings()
	var pathErr *PathError
	fmt.Println("errors.Is(err, ErrMissing):", errors.Is(err, ErrMissing))
	fmt.Println("errors.As(err, &pathErr):", errors.As(err, &pathErr))
	if pathErr != nil {
		fmt.Printf("operation: %s\npath: %s\n", pathErr.Op, pathErr.Path)
	}
	if *show {
		fmt.Println("error:", err)
	}
}
