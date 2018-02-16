package alias_test

import (
	"github.com/gregoryv/alias"
	"testing"
)

var nop = func() error { return nil }

func TestNewBasic(t *testing.T) {
	name := "help"
	cmd, _ := alias.NewBasic(name, nop)
	if cmd == nil {
		t.Error("should return a Command")
	}
	_, err := alias.NewBasic(name, nop)
	if err == nil {
		t.Errorf("should fail if trying to register under same name")
	}
}

func TestCommand_Run(t *testing.T) {
	cmd, _ := alias.NewBasic("testsrun", nop)
	if err := cmd.Run(); err != nil {
		t.Error("nop command should be runnable")
	}
}
