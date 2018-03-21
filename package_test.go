package cli_test

import (
	"fmt"
	"github.com/gregoryv/cli"
	"testing"
)

func TestAdd(t *testing.T) {
	name := "help"
	if err := cli.Add(name, nop); err != nil {
		t.Error("should return a Command")
	}
	if err := cli.Add(name, nop); err == nil {
		t.Errorf("should fail if trying to register under same name")
	}
}

func TestList(t *testing.T) {
	cli.Add("help", nop)
	if res := cli.List(); len(res) == 0 {
		t.Fail()
	}
}

func ExampleAdd() {
	cmds := cli.NewCommandSet()
	cmds.Add("help", nop)
	cmds.Add("add", nop)
	cmds.Add("list", nop)
	for _, name := range cmds.List() {
		fmt.Println(name)
	}
	// output:
	// add
	// help
	// list
}
