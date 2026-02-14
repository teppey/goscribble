package main

import (
	"fmt"
	"os"
)

type exitCode int

const (
	exitOK    exitCode = 0
	exitError exitCode = 1
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("goscrawl is a tool for edit and run Go code instantly on console.\n")
		fmt.Println("Usage:\n")
		fmt.Println("\tgoscrawl <command> [arguments]\n")
		fmt.Println("The commands are:\n")
		fmt.Println("\tclear  clear all files")
		fmt.Println("\tedit   edit and run file")
		fmt.Println("\tlist   list all files")
		fmt.Println("\tnew    create and edit file and run it")
		fmt.Println("\trun    run file")
		fmt.Println("\tshow   display file\n")
		os.Exit(int(exitError))
	}

	command := os.Args[1]
	subArgs := os.Args[2:]
	var err error
	switch command {
	case "clear":
		err = commandClear(subArgs)
	case "edit":
		err = commandEdit(subArgs)
	case "list":
		err = commandList(subArgs)
	case "new":
		err = commandNew(subArgs)
	case "run":
		err = commandRun(subArgs)
	case "show":
		err = commandShow(subArgs)
	default:
		err = fmt.Errorf("unknown command: %s", command)
	}

	code := exitOK
	if err != nil {
		fmt.Printf("error: %s", err)
		code = exitError
	}

	os.Exit(int(code))
}
