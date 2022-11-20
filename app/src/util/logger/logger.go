package logger

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
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

func CustomGinLogger(paths []string) gin.LoggerConfig {
	conf := gin.LoggerConfig{
		SkipPaths: paths,
	}
	conf.Formatter = customGinLogFormatter
	return conf
}

var customGinLogFormatter = func(param gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.Latency > time.Minute {
		param.Latency = param.Latency.Truncate(time.Second)
	}

	return fmt.Sprintf("%v [INFO]|%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
		param.TimeStamp.Format("2006/01/02 15:04:05"),
		statusColor,
		param.StatusCode,
		resetColor,
		param.Latency,
		param.ClientIP,
		methodColor,
		param.Method,
		resetColor,
		param.Path,
		param.ErrorMessage,
	)
}
