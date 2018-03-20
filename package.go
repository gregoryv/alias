// Package cli helps define sub commands
package cli

var DefaultCmds = NewCommandSet()

// New registers the command with the default router
func Add(name string, cmd Command) error {
	return DefaultCmds.Add(name, cmd)
}
