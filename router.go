package cli

import (
	"fmt"
	"sort"
)

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

// List returns a sorted list of commands
func (r Router) List() []string {
	result := make([]string, 0)
	for name, _ := range r {
		result = append(result, name)
	}
	sort.Sort(sort.StringSlice(result))
	return result
}
