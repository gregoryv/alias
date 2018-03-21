package cli

import (
	"fmt"
	"sort"
)

// Command interface at it's most basic level
type Command func() error
type CommandSet map[string]Command

func NewCommandSet() CommandSet {
	return make(map[string]Command)
}

// Add a named command to the set, cannot be override existing named commands
func (cmds CommandSet) Add(name string, cmd Command) error {
	if _, ok := cmds[name]; ok {
		return fmt.Errorf("%q already added", name)
	}
	cmds[name] = cmd
	return nil
}

// List returns a sorted list of named commands
func (cmds CommandSet) List() (names []string) {
	names = make([]string, 0)
	for name, _ := range cmds {
		names = append(names, name)
	}
	sort.Sort(sort.StringSlice(names))
	return
}

func (cmds CommandSet) Call(name string) error {
	if cmd, ok := cmds[name]; ok {
		return cmd()
	}
	return fmt.Errorf("%q no such command")
}
