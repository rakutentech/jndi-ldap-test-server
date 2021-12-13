package logging

import (
	"github.com/jwalton/go-supportscolor"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/term"
	"os"
)

type Flags struct {
	Color string
	Level string
}

var originalLogger = log.Logger

func InitializeLogger() {
	log.Logger = determineLogger(os.Stderr, "")
}

func UpdateLoggerWithFlags(flags *Flags) {
	log.Logger = determineLogger(os.Stderr, flags.Color).
		Level(determineLevel(flags.Level))
}

func determineLogger(file *os.File, colorFlag string) zerolog.Logger {
	fd := file.Fd()
	if !term.IsTerminal(int(fd)) {
		return originalLogger // Use default logger (JSON)
	}

	// Use console logger with color depending on terminal
	return log.Output(
		zerolog.ConsoleWriter{Out: os.Stderr, NoColor: !shouldUseColor(fd, colorFlag)},
	)
}

func determineLevel(levelFlag string) zerolog.Level {
	switch levelFlag {
	case "debug":
		return zerolog.DebugLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	default:
		return zerolog.InfoLevel
	}
}

func shouldUseColor(fd uintptr, colorFlag string) bool {
	switch colorFlag {
	case "always":
		return true
	case "never":
		return false
	default:
		// auto
		return supportscolor.SupportsColor(
			fd,
			supportscolor.SniffFlagsOption(false),
		).SupportsColor
	}
}
