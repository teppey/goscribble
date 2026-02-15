package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	numBackup = 9

	template = `package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
}
`
)

func newCommand(args []string) error {
	dir, err := baseDir()
	if err != nil {
		return err
	}

	if err := os.Mkdir(dir, 0700); err != nil && !os.IsExist(err) {
		return fmt.Errorf("failed to create directory: %s: %w", dir, err)
	}

	const num = 0
	if err := rotate(dir, num); err != nil {
		return fmt.Errorf("failed to rotate: %s: %w", dir, err)
	}

	newPath := filepath.Join(dir, fmt.Sprintf("%d.go", num))
	if err := os.WriteFile(newPath, []byte(template), 0600); err != nil {
		return fmt.Errorf("failed to write file: %s: %w", newPath, err)
	}

	if err := edit(newPath); err != nil {
		return err
	}

	if err := goimports(newPath); err != nil {
		return err
	}

	if err := run(newPath); err != nil {
		return err
	}

	return nil
}

func rotate(dir string, n int) error {
	if n > numBackup {
		msg := fmt.Sprintf("`n` should be `numBackup` or less: n=%d, numBackup:%d", n, numBackup)
		panic(msg)
	}

	cur := filepath.Join(dir, fmt.Sprintf("%d.go", n))
	if !exists(cur) {
		return nil
	}

	if n == numBackup {
		if err := os.Remove(cur); err != nil {
			return err
		}
		return nil
	}

	old := filepath.Join(dir, fmt.Sprintf("%d.go", n+1))
	if exists(old) {
		if err := rotate(dir, n+1); err != nil {
			return err
		}
	}

	if err := os.Rename(cur, old); err != nil {
		return err
	}

	return nil
}
