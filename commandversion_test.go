package main

import (
	"testing"
)

func Test_CLI(t *testing.T) {
	versionCommand := VersionCommand{}

	if err := versionCommand.Execute([]string{}); err != nil {
		t.Fail()
	}
}
