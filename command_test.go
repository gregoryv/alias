package shal_test

import (
	"github.com/gregoryv/shal"
	"testing"
)

var nop = func() error { return nil }

func TestNewCommandFunc(t *testing.T) {
	alias := "help"
	cmd, _ := shal.NewCommandFunc(alias, nop)
	if cmd == nil {
		t.Error("NewCommandFunc(...) should return a Command")
	}
	_, err := shal.NewCommandFunc(alias, nop)
	if err == nil {
		t.Errorf("NewCommandFunc() should fail if trying to register under same name")
	}
}

func TestCommand_Run(t *testing.T) {
	cmd, _ := shal.NewCommandFunc("testsrun", nop)
	if err := cmd.Run(); err != nil {
		t.Error("nop command should be runnable")
	}
}
