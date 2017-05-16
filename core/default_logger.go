package core

import "fmt"

// DefaultLogger implements the api.Logger interface, you should inject your own logger in
// this default logger just does simple fmt.Print commands
type DefaultLogger struct {
}


// Log lets you pass in a severity and error to report back
func (l *DefaultLogger) Log(severity string, err error) {
	fmt.Printf("[%s] - %v\n", severity, err)
}
