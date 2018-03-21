// Package cli helps define sub commands
package cli

var DefaultCmds = NewCommandSet()

// Add a named command to the default command set
func Add(name string, cmd Command) error {
	return DefaultCmds.Add(name, cmd)
}

// Call a registered command in the default command set
func Call(name string) error {
	return DefaultCmds.Call(name)
}

// List commands in the default command set
func List() []string {
	return DefaultCmds.List()
}
