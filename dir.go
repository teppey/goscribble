package main

import (
	"fmt"
)

func dirCommand(args []string) error {
	dir, err := baseDir()
	if err != nil {
		return err
	}

	fmt.Println(dir)
	return nil
}
