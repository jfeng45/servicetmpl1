// Package logger represents a generic logging interface
package logger

import "github.com/jfeng45/glogger"

// Log is a package level variable, every program should access logging function through "Log"
var Log glogger.Logger

// SetLogger is the setter for log variable, it should be the only way to assign value to log
func SetLogger(newLogger glogger.Logger) {
	Log = newLogger
}
