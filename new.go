package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	maxNum = 3

	template = `package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
`
)

func commandNew(args []string) error {
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
	file, err := os.Create(newPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %s: %w", newPath, err)
	}
	defer file.Close()

	io.WriteString(file, template)

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
	if n > maxNum {
		panic("over limit")
	}

	cur := filepath.Join(dir, fmt.Sprintf("%d.go", n))
	if !exists(cur) {
		return nil
	}

	if n == maxNum {
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

func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			panic(err)
		}
		return false
	}

	return true
}
