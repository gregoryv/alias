// Package cli helps define sub commands
package cli

var DefaultCmds = NewCommandSet()

// Add a named command to the default command set
func Add(name string, cmd Command) error {
	return DefaultCmds.Add(name, cmd)
}

// List commands in the default command set
func List() []string {
	return DefaultCmds.List()
}
