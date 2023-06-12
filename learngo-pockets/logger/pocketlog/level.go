package pocketlog

// Level represents an available logging level.
type Level byte

const (
	// LevelDebug represents the lowest level of log , mostly used for
	LevelDebug = iota
	//LevelInfo represents a logging level that contains information deemed valuable.
	LevelInfo
	//Level Error represents the highest logging level, only to be used to trace
	LevelError
)
