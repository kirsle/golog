package golog

import (
	"os"
	"sync"
	"text/template"

	"golang.org/x/crypto/ssh/terminal"
)

// An internal map of named loggers. This allows for GetLogger() to be called
// many times from anywhere in your code base, but for only one logger instance
// to be created for it.
var (
	loggers     map[string]*Logger
	loggerMutex sync.Mutex

	interactive bool // is an interactive shell
)

func init() {
	loggers = map[string]*Logger{}

	// Detect if we're running in an interactive shell, so we can globally
	// disable colors when redirecting to a log file.
	interactive = terminal.IsTerminal(int(os.Stdout.Fd()))
}

// Logger stores the configuration for a named logger instance.
type Logger struct {
	Name   string
	Config *Config

	// Private cached text/template, the first time the formatter is used.
	template *template.Template
}

// GetLogger initializes and returns a new Logger.
func GetLogger(name string) *Logger {
	loggerMutex.Lock()
	defer loggerMutex.Unlock()

	// Initialize the logger the first time we ask for it.
	if _, ok := loggers[name]; !ok {
		loggers[name] = &Logger{
			Name:   name,
			Config: DefaultConfig(),
		}
	}

	return loggers[name]
}
