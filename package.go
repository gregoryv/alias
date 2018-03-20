// Package cli implements router for subcommand cli applications
package cli

var DefaultRouter = NewRouter()

// New registers the command with the default router
func Add(name string, cmd Command) error {
	return DefaultRouter.Add(name, cmd)
}
