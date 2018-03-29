// Package cli helps define actions for commandline applications
package cli

var DefaultActions = NewActionSet()

// Add a named command to the default command set
func Add(name string, cmd ActionFunc) error {
	return DefaultActions.Add(name, cmd)
}

// Call a registered command in the default command set
func Call(name string) error {
	return DefaultActions.Call(name)
}

// List commands in the default command set
func List() []string {
	return DefaultActions.List()
}
