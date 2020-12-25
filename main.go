package main

import (
	"github.com/jessevdk/go-flags"
	"log"
	"os"
)

func main() {
	os.Exit(run())
}

func run() int {
	var options struct{}
	var parser = flags.NewParser(&options, flags.Default)

	if _, err := parser.AddCommand("new", "Create a new memo", "", &NewCommand{}); err != nil {
		log.Fatal(err)
	}
	if _, err := parser.AddCommand("version", "Show version", "", &VersionCommand{}); err != nil {
		log.Fatal(err)
	}

	if _, err := parser.Parse(); err != nil {
		switch err.(type) {
		case *flags.Error:
			fe, _ := err.(*flags.Error)
			if fe.Type == flags.ErrHelp {
				return 0
			}
			return 1
		default:
			return 1
		}
	}

	return 0
}
