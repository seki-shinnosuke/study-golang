package logger

import (
	"fmt"
	"log"
)

var logger *log.Logger

func init() {
	logger = log.Default()
}

func Debug(format string, a ...any) {
	text := logText("[DEBUG]", format, a...)
	logger.Print(text)
}

func Info(format string, a ...any) {
	text := logText("[INFO]", format, a...)
	logger.Print(text)
}

func Warn(format string, a ...any) {
	text := logText("[WARN]", format, a...)
	logger.Print(text)
}

func Error(format string, a ...any) {
	text := logText("[ERROR]", format, a...)
	logger.Printf(text)
}

func Fatal(format string, a ...any) {
	text := logText("[FATAL]", format, a...)
	logger.Fatal(text)
}

func logText(level, format string, a ...any) string {
	text := level + format
	if len(a) != 0 {
		text = fmt.Sprintf(text, a...)
	}
	return text
}
