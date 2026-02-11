package main

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
)

func baseDir() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", errors.New("failed to get current user")
	}

	return filepath.Join(os.TempDir(), "gosketch_"+u.Username), nil
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
