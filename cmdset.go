package cli

import (
	"fmt"
	"sort"
)

type Action interface {
	Run() error
}

// Command interface at it's most basic level
type ActionFunc func() error

func (cmd ActionFunc) Run() error {
	return cmd()
}

type ActionSet map[string]Action

func NewActionSet() ActionSet {
	return make(map[string]Action)
}

// Add a named command to the set, cannot be override existing named commands
func (cmds ActionSet) Add(name string, cmd Action) error {
	if _, ok := cmds[name]; ok {
		return fmt.Errorf("%q already added", name)
	}
	cmds[name] = cmd
	return nil
}

// List returns a sorted list of named commands
func (cmds ActionSet) List() (names []string) {
	names = make([]string, 0)
	for name, _ := range cmds {
		names = append(names, name)
	}
	sort.Sort(sort.StringSlice(names))
	return
}

func (cmds ActionSet) Call(name string) error {
	if cmd, ok := cmds[name]; ok {
		return cmd.Run()
	}
	return fmt.Errorf("%q no such command")
}
