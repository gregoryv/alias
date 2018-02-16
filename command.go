package alias

import "fmt"

var DefaultRouter Router

type Router map[string]Command

func NewRouter() Router {
	return make(map[string]Command)
}

type Command interface {
	Run() error
}

type BasicFunc func() error

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
	command BasicFunc
}

func (fc *funcCommand) Run() error {
	return fc.command()
}

func NewBasic(alias string, fn BasicFunc) (Command, error) {
	cmd := &funcCommand{command: fn}
	return cmd, newCommand(alias, cmd)
}
