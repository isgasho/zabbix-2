package zabbix

import "log"

// Logger is a reusable interface for logging
// errors and events. This is used by the API class
// and can be replaced with your own implementation
type Logger interface {
	Print(...interface{})
	Warning(...interface{})
	Fatal(...interface{})
}

// StdOutLogger implements the logger interface and
// wrap the go stdlib log functionality.
type StdOutLogger struct {
}

// Print ouptuts a general message to std out
func (sol *StdOutLogger) Print(args ...interface{}) {
	log.Print(args)
}

// Warning acts the same as Print
func (sol *StdOutLogger) Warning(args ...interface{}) {
	log.Print(args)
}

// Fatal outputs a fatal error message to std out.
func (sol *StdOutLogger) Fatal(args ...interface{}) {
	log.Fatal(args)
}
