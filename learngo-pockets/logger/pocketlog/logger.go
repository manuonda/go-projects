package pocketlog

import (
	"fmt"
	"io"
	"os"
)

type Logger struct {
	threshold Level
	output    io.Writer
}

// Debug formats and prints a message if the log level is debug or
// higher
func (l *Logger) Debugf(format string, args ...any) {

	if l.threshold > LevelDebug {
		return
	}
	l.logf(format, args...)

}

func (l *Logger) Infof(format string, args ...any) {

}

// logf prints the message to the output
// Add decorations here, if any. #A
func (l *Logger) logf(format string, args ...any) {
	_, _ = fmt.Fprintf(l.output, format+"", args...)
}

func New(threshold Level, opts ...Option) *Logger {
	lgr := &Logger{threshold: threshold, output: os.Stdout}

	for _, configFunc := range opts {
		configFunc(lgr)
	}

	return lgr
}
