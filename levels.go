package golog

import (
	"fmt"
	"os"
)

type logLevel int

// Log levels for controlling whether or not logs of certain types will be
// emitted by your logger.
const (
	DebugLevel logLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

// Map log levels to human readable labels.
var levelNames = map[logLevel]string{
	DebugLevel: "DEBUG",
	InfoLevel:  "INFO",
	WarnLevel:  "WARN",
	ErrorLevel: "ERROR",
}

// emit is the general purpose log line emitter.
func (l *Logger) emit(level logLevel, tmpl string, args ...interface{}) {
	message := l.Format(level, tmpl, args...)

	// If we have a log writer, send it there.
	if l.Config.Writer != nil {
		// l.Config.Writer.Write(message)
	} else {
		// No writer given so we default to standard out/error.
		if level <= InfoLevel {
			fmt.Fprintln(os.Stdout, message)
		} else {
			fmt.Fprintln(os.Stderr, message)
		}
	}
}

// Debug emits a debug-level message from the logger.
func (l *Logger) Debug(tmpl string, args ...interface{}) {
	if l.Config.Level <= DebugLevel {
		l.emit(DebugLevel, tmpl, args...)
	}
}

// Info emits an informational message.
func (l *Logger) Info(tmpl string, args ...interface{}) {
	if l.Config.Level <= InfoLevel {
		l.emit(InfoLevel, tmpl, args...)
	}
}

// Warn emits a warning message.
func (l *Logger) Warn(tmpl string, args ...interface{}) {
	if l.Config.Level <= WarnLevel {
		l.emit(WarnLevel, tmpl, args...)
	}
}

// Error emits an error message.
func (l *Logger) Error(tmpl string, args ...interface{}) {
	if l.Config.Level <= ErrorLevel {
		l.emit(ErrorLevel, tmpl, args...)
	}
}
