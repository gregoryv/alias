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

func TestCommandSet_Add(t *testing.T) {
	cmds := cli.NewCommandSet()
	if len(cmds) != 0 {
		t.Errorf("should be empty by default")
	}
	if cmds.Add("n", nop); len(cmds) != 1 {
		t.Errorf("should increase size of set")
	}
	if err := cmds.Add("n", nop); err == nil {
		t.Errorf("should fail to add same name")
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

func TestCommandSet_Call(t *testing.T) {
	cmds := cli.NewCommandSet()
	called := false
	cmds.Add("help", func() error {
		called = true
		return nil
	})
	if err := cmds.Call("help"); err != nil {
		t.Fail()
	}
	if err := cmds.Call("not"); err == nil {
		t.Fail()
	}
}
