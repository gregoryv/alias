package cli_test

import (
	"github.com/gregoryv/cli"
	"testing"
)

var nop = func() error { return nil }

func TestNewRouter(t *testing.T) {
	if r := cli.NewRouter(); r == nil {
		t.Error("NewRouter() returned nil")
	}
}

func TestAdd(t *testing.T) {
	name := "help"
	if err := cli.Add(name, nop); err != nil {
		t.Error("should return a Command")
	}
	if err := cli.Add(name, nop); err == nil {
		t.Errorf("should fail if trying to register under same name")
	}
}
