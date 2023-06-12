package pocketlog

import "io"

// Option defines a functional option to our logger
type Option func(*Logger)

func WithOuput(output io.Writer) Option {
	return func(lgr *Logger) {
		lgr.output = output
	}
}
