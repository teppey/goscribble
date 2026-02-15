package main

import (
	"fmt"
	"path/filepath"
)

func listCommand(args []string) error {
	dir, err := baseDir()
	if err != nil {
		return err
	}

	if paths, err := filepath.Glob(filepath.Join(dir, "[0-9].go")); err == nil {
		for _, path := range paths {
			fmt.Println(path)
		}
	}

	return nil
}
