package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func showCommand(args []string) error {
	dir, err := baseDir()
	if err != nil {
		return err
	}

	path := filepath.Join(dir, "0.go")
	if exists(path) {
		contents, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		fmt.Print(string(contents))
	}

	return nil
}
