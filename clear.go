package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func commandClear(args []string) error {
	dir, err := baseDir()
	if err != nil {
		return err
	}

	// TODO: "9"以上の数字のファイルに対応
	if paths, err := filepath.Glob(filepath.Join(dir, "[0-9].go")); err == nil {
		for _, path := range paths {
			if err := os.Remove(path); err != nil {
				return fmt.Errorf("failed to remove file: %s: %w", path, err)
			}
		}
	}

	return nil
}
