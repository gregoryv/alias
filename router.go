package ui

import (
	"fmt"
)

var DefaultRouter = NewRouter()

type Command func() error
type Router map[string]Command

func NewRouter() Router {
	return make(map[string]Command)
}

// New registers a new alias, existing alias cannot be overriden
func (r Router) New(ui string, cmd Command) error {
	if _, ok := r[ui]; ok {
		return fmt.Errorf("%q already registered", ui)
	}
	r[ui] = cmd
	return nil
}

// New registers the command with the default router
func New(ui string, cmd Command) (Command, error) {
	return cmd, DefaultRouter.New(ui, cmd)
}
