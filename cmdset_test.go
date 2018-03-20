package cli_test

import (
	"fmt"
	"github.com/gregoryv/cli"
	"testing"
)

var nop = func() error { return nil }

func TestNewCommandSet(t *testing.T) {
	if r := cli.NewCommandSet(); r == nil {
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

func TestCommandSet_List(t *testing.T) {
	cmds := cli.NewCommandSet()
	if result := cmds.List(); len(result) != 0 {
		t.Errorf("should be empty by default")
	}
	cmds.Add("b", nop)
	cmds.Add("c", nop)
	cmds.Add("a", nop)
	exp := []string{"a", "b", "c"}
	res := cmds.List()
	if fmt.Sprintf("%v", res) != fmt.Sprintf("%v", exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}
