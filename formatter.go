package golog

import (
	"bytes"
	"fmt"
	"text/template"
	"time"

	"github.com/kirsle/golog/ansi"
)

// Convenient log formats to use in your logger.
const (
	// DefaultFormat: shows the date in the secondary (dark) color, the label
	// in the bright color, and the message text in the normal color.
	DefaultFormat = `{{.Secondary}}{{.Time}}{{.Reset}} {{.Primary}}[{{.Level}}]{{.Reset}} {{.Message}}`

	// ColorfulFormat: like the DefaultFormat, but the message itself is also
	// colored using the secondary color.
	ColorfulFormat = `{{.Secondary}}{{.Time}}{{.Reset}} {{.Primary}}[{{.Level}}]{{.Reset}} {{.Secondary}}{{.Message}}{{.Reset}}`
)

// Convenient time formats to use in your logger.
const (
	// DefaultTime is the default, in `yyyy-mm-dd hh:mm:ss` format.
	DefaultTime = `2006-01-02 15:04:05`

	// FriendlyTime is a human readable `Jan 2 15:04:05 2006` format.
	FriendlyTime = `Jan 2 15:04:05 2006`
)

// formatter provides the variables that can be used in the log format.
type formatter struct {
	Time      string
	Level     string
	Message   string
	Primary   string
	Secondary string
	Reset     string
}

// Format and return a log message.
func (l *Logger) Format(level logLevel, tmpl string, args ...interface{}) string {
	// Prepare the variables to apply to the log message format.
	format := formatter{
		Time:    time.Now().Format(l.Config.TimeFormat),
		Level:   levelNames[level],
		Message: fmt.Sprintf(tmpl, args...),
	}

	// Find the theme color to use.
	if interactive && l.Config.Colors != NoColor {
		var (
			primary   ThemeColor
			secondary ThemeColor
		)

		switch level {
		case DebugLevel:
			primary = l.Config.Theme.Debug
			secondary = l.Config.Theme.DebugSecondary
		case InfoLevel:
			primary = l.Config.Theme.Info
			secondary = l.Config.Theme.InfoSecondary
		case WarnLevel:
			primary = l.Config.Theme.Warn
			secondary = l.Config.Theme.WarnSecondary
		case ErrorLevel:
			primary = l.Config.Theme.Error
			secondary = l.Config.Theme.ErrorSecondary
		}

		// What color level are we supporting?
		if l.Config.Colors == ANSIColor {
			format.Primary = fmt.Sprintf("\x1B[%sm", primary.ANSI)
			format.Secondary = fmt.Sprintf("\x1B[%sm", secondary.ANSI)
			format.Reset = fmt.Sprintf("\x1B[%sm", ansi.Reset)
		} else if l.Config.Colors == ExtendedColor {
			format.Primary = fmt.Sprintf("\x1B[%sm", primary.Extended)
			format.Secondary = fmt.Sprintf("\x1B[%sm", secondary.Extended)
			format.Reset = fmt.Sprintf("\x1B[%sm", ansi.Reset)
		}
	}

	// Do we have the template cached?
	if l.template == nil {
		template, err := template.New("golog").Parse(l.Config.Format)
		if err != nil {
			return fmt.Sprintf("[GoLog format error: %s]", err)
		}
		l.template = template
	}

	// Evaluate the template.
	var buf bytes.Buffer
	err := l.template.Execute(&buf, format)
	if err != nil {
		return fmt.Sprintf("[GoLog template error: %s]", err)
	}

	return buf.String()
}
