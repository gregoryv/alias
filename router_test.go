package ui_test

import (
	"github.com/gregoryv/ui"
	"testing"
)

var nop = func() error { return nil }

func TestNew(t *testing.T) {
	name := "help"
	cmd, _ := ui.NewAlias(name, nop)
	if cmd == nil {
		t.Error("should return a Command")
	}
	_, err := ui.NewAlias(name, nop)
	if err == nil {
		t.Errorf("should fail if trying to register under same name")
	}
}

func Test_runAliasedCommand(t *testing.T) {
	cmd, _ := ui.NewAlias("testsrun", nop)
	if err := cmd(); err != nil {
		t.Error("nop command should be runnable")
	}
}
