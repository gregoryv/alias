package cli

import (
	"fmt"
)

var DefaultRouter = NewRouter()

// Command interface at it's most basic level
type Command func() error
type Router map[string]Command

func NewRouter() Router {
	return make(map[string]Command)
}

// New registers a new named command, cannot be override existing named commands
func (r Router) Add(name string, cmd Command) error {
	if _, ok := r[name]; ok {
		return fmt.Errorf("%q already added", name)
	}
	r[name] = cmd
	return nil
}

// New registers the command with the default router
func Add(name string, cmd Command) error {
	return DefaultRouter.Add(name, cmd)
}
