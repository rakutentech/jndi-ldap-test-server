package logging

import (
	"fmt"
	"github.com/rs/zerolog"
)

type StdLogger interface {
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})

	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Panicln(v ...interface{})

	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

func NewStdAdapter(logger *zerolog.Logger) StdLogger {
	return &stdLoggerAdapter{logger}
}

type stdLoggerAdapter struct {
	logger *zerolog.Logger
}

// Verify that interface is implemented
var _ StdLogger = &stdLoggerAdapter{}

func (a stdLoggerAdapter) Fatal(v ...interface{}) {
	a.logger.Fatal()
	a.logger.Fatal().Msg(fmt.Sprint(v...))
}

func (a stdLoggerAdapter) Fatalf(format string, v ...interface{}) {
	a.logger.Fatal().Msgf(format, v...)
}

func (a stdLoggerAdapter) Fatalln(v ...interface{}) {
	a.logger.Fatal().Msg(fmt.Sprint(v...))
}

func (a stdLoggerAdapter) Panic(v ...interface{}) {
	a.logger.Panic().Msg(fmt.Sprint(v...))
}

func (a stdLoggerAdapter) Panicf(format string, v ...interface{}) {
	a.logger.Panic().Msgf(format, v...)
}

func (a stdLoggerAdapter) Panicln(v ...interface{}) {
	a.logger.Panic().Msg(fmt.Sprint(v...))
}

func (a stdLoggerAdapter) Print(v ...interface{}) {
	a.logger.Print(v...)
}

func (a stdLoggerAdapter) Printf(format string, v ...interface{}) {
	a.logger.Printf(format, v...)
}

func (a stdLoggerAdapter) Println(v ...interface{}) {
	a.logger.Print(v...)
}

