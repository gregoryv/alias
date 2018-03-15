package alias_test

import (
	"github.com/gregoryv/alias"
	"testing"
)

var nop = func() error { return nil }

func TestNew(t *testing.T) {
	name := "help"
	cmd, _ := alias.New(name, nop)
	if cmd == nil {
		t.Error("should return a Command")
	}
	_, err := alias.New(name, nop)
	if err == nil {
		t.Errorf("should fail if trying to register under same name")
	}
}

func TestCommand(t *testing.T) {
	cmd, _ := alias.New("testsrun", nop)
	if err := cmd(); err != nil {
		t.Error("nop command should be runnable")
	}
}
