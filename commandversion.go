package main

import "fmt"

type VersionCommand struct {}

func (x VersionCommand) Execute(args []string) error {
	fmt.Printf("Version: %s\n", version)

	return nil
}
