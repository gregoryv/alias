package alias

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
func (r Router) New(alias string, cmd Command) error {
	if _, ok := r[alias]; ok {
		return fmt.Errorf("%q already registered", alias)
	}
	r[alias] = cmd
	return nil
}

// New registers the command with the default router
func New(alias string, cmd Command) (Command, error) {
	return cmd, DefaultRouter.New(alias, cmd)
}
