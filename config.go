package golog

import "io"

// Config stores settings that control the logger's behavior.
type Config struct {
	// Level is one of DebugLevel, InfoLevel, WarnLevel, ErrorLevel or FatalLevel.
	// Messages emitted by the logger must be 'at least' this level to be logged.
	Level logLevel

	// What colors are supported? Default is NoColor. Use ANSIColor to support
	// legacy terminal emulators, or ExtendedColor for modern 256-color support.
	Colors colorLevel

	// Which color theme are you using? The default is DarkTheme.
	Theme Theme

	// Where to write the log messages to? If not defined with a custom io.Writer,
	// the default goes to standard output for Debug and Info messages and
	// standard error for warnings, errors, and fatal messages.
	Writer *io.Writer

	// How do you want to format your log lines? This should be a Go text format
	// string, with the following variable placeholders:
	//
	//    {{.Time}} inserts the date/time stamp for the log message.
	//    {{.Level}} inserts a label for the log level, e.g. "INFO" or "WARN"
	//    {{.Message}} inserts the text of the log message itself.
	//    {{.Primary}} inserts the color sequence for the primary color based
	//        on the log level for the message.
	//    {{.Secondary}} inserts the color sequence for the secondary color.
	//    {{.Reset}} inserts the 'reset' color sequence to stop coloring
	//        the rest of the text that follows.
	//
	// The default log format is as follows:
	//
	//    {{.Secondary}}{{.Time}}{{.Reset}} {{.Primary}}[{{.Level}}]{{.Reset}} {{.Message}}
	Format string

	// How do you want to format your time stamps? (The `{{.Time}}`). This uses
	// the Go `time` module, so the TimeFormat should use their reference date/time.
	// The default TimeFormat is: `2006-01-02 15:04:05`
	TimeFormat string
}

// DefaultConfig returns a Config with the default values filled in.
func DefaultConfig() *Config {
	return &Config{
		Theme:      DarkTheme,
		Format:     DefaultFormat,
		TimeFormat: DefaultTime,
	}
}

// Configure applies the configuration to the logger. If any of the following
// keys are not defined (or have zero-values), the default value for the key will
// be used instead:
//
//   Format
//   TimeFormat
func (l *Logger) Configure(cfg *Config) {
	// Important keys and their defaults.
	if cfg.Format == "" {
		cfg.Format = DefaultFormat
	}
	if cfg.TimeFormat == "" {
		cfg.TimeFormat = DefaultTime
	}

	l.Config = cfg
	l.template = nil
}
