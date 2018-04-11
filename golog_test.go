package golog

import "testing"

func TestColors(t *testing.T) {
	interactive = true // fake an interactive TTY to test color outputs

	log := GetLogger("test")

	// Helper function to emit all the log types.
	emitLogs := func(message string) {
		log.Debug(message)
		log.Info(message)
		log.Warn(message)
		log.Error(message)
	}

	log.Configure(&Config{
		Theme:  DarkTheme,
		Colors: ANSIColor,
	})
	emitLogs("With standard 16-color ANSI codes.")

	log.Configure(&Config{
		Theme:  DarkTheme,
		Colors: ExtendedColor,
	})
	emitLogs("With xterm-256color codes.")

	log.Configure(&Config{
		Theme:  DarkTheme,
		Colors: ExtendedColor,
		Format: ColorfulFormat,
	})
	emitLogs("Colorful format.")
}

func TestLogLevels(t *testing.T) {
	log := GetLogger("levels")

	// Helper function to emit all the log types.
	emitLogs := func(message string) {
		log.Debug(message)
		log.Info(message)
		log.Warn(message)
		log.Error(message)
	}

	emitLogs("Default log level=debug")
	log.Config.Level = InfoLevel
	emitLogs("With Level=Info")
	log.Config.Level = WarnLevel
	emitLogs("With Level=Warn")
	log.Config.Level = ErrorLevel
	emitLogs("With Level=Error")
}
