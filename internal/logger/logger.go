package logger

type Logger struct {
	Level string
}

func New() *Logger {
	return &Logger{}
}
