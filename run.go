package main

import (
	"path/filepath"
)

func runCommand(args []string) error {
	dir, err := baseDir()
	if err != nil {
		return err
	}

	path := filepath.Join(dir, "0.go")

	if err := run(path); err != nil {
		return err
	}

	return nil
}
