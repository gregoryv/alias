package shal

import "fmt"

var DefaultRouter Router

type Router map[string]Command

func NewRouter() Router {
	return make(map[string]Command)
}

type Command interface {
	Run() error
}

func newCommand(alias string, cmd Command) error {
	if DefaultRouter == nil {
		DefaultRouter = NewRouter()
	}
	if _, ok := DefaultRouter[alias]; ok {
		return fmt.Errorf("%q already registered", alias)
	}
	DefaultRouter[alias] = cmd
	return nil
}

type funcCommand struct {
	command func() error
}

func (fc *funcCommand) Run() error {
	return fc.command()
}

func NewCommandFunc(alias string, fn func() error) (Command, error) {
	cmd := &funcCommand{command: fn}
	return cmd, newCommand(alias, cmd)
}
