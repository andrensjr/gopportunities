package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

func (l *Logger) Debug(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *Logger) Info(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) Warn(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}

func (l *Logger) Err(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
