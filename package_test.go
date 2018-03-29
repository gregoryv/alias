package cli_test

import (
	"fmt"
	"github.com/gregoryv/cli"
	"testing"
)

func init() {
	cli.Add("help", nop)
}

func TestCall(t *testing.T) {
	cli.Call("help")
}

func TestList(t *testing.T) {
	if res := cli.List(); len(res) == 0 {
		t.Fail()
	}
}

func ExampleAdd() {
	cmds := cli.NewActionSet()
	cmds.Add("help", cli.ActionFunc(func() error { return nil }))
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
