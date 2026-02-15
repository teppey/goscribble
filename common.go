package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
)

func baseDir() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", errors.New("failed to get current user")
	}

	return filepath.Join(os.TempDir(), "goscribble_"+u.Username), nil
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

func edit(path string) error {
	cmd := exec.Command("vim", path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to edit file: %s: %w", cmd, err)
	}

	return nil
}

func goimports(path string) error {
	cmd := exec.Command("goimports", "-w", path)
	out, err := cmd.CombinedOutput()
	if len(out) > 0 {
		fmt.Print(string(out))
	}

	return err
}

func run(path string) error {
	cmd := exec.Command("go", "run", path)
	out, err := cmd.CombinedOutput()
	if len(out) > 0 {
		fmt.Print(string(out))
	}

	return err
}
