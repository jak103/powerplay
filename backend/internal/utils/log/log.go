package log

import (
	"fmt"
	"maps"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
)

const (
	DEBUG int = 0
	INFO  int = 1
	WARN  int = 2
	ERROR int = 3
	ALERT int = 4
)

var blue = color.New(color.FgBlue).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()

var magenta = color.New(color.FgMagenta).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()

var TheLogger Logger

type Logger struct {
	tags        map[string]any
	color       bool
	skip        int
	testCapture *string
	level       int
	requestId   string
}

func (l *Logger) SetLevel(level string) {
	switch level {
	case "DEBUG":
		l.level = DEBUG
	case "INFO":
		l.level = INFO
	case "WARN":
		l.level = WARN
	case "ERROR":
		l.level = ERROR
	case "ALERT":
		l.level = ALERT
	}
}

func (l *Logger) SetTestCapture(ptr *string) {
	l.testCapture = ptr
}

func (l Logger) Debug(format string, a ...any) {
	if l.level == DEBUG {
		print(magenta("[DEBUG]"), fmt.Sprintf(format, a...), l.requestId, l.tags, l.skip)
	}
}

func (l Logger) Info(format string, a ...any) {
	if l.level <= INFO {
		print(green("[INFO ]"), fmt.Sprintf(format, a...), l.requestId, l.tags, l.skip)
	}
}

func (l Logger) Warn(format string, a ...any) {
	if l.level <= WARN {
		print(yellow("[WARN ]"), fmt.Sprintf(format, a...), l.requestId, l.tags, l.skip)
	}
}

func (l Logger) Error(format string, a ...any) {
	if l.level <= ERROR {
		print(red("[ERROR]"), fmt.Sprintf(format, a...), l.requestId, l.tags, l.skip)
	}
}

func (l Logger) Alert(format string, a ...any) {
	if l.level <= ALERT {
		print(red("[ALERT]"), fmt.Sprintf(format, a...), l.requestId, l.tags, l.skip)
	}
}

func Init(level string, color bool) error {
	TheLogger = Logger{
		tags:  make(map[string]any),
		color: color,
		skip:  2,
	}

	TheLogger.SetLevel(level)

	return nil
}

func newLogger(oldLogger Logger) Logger {
	newLogger := Logger{
		tags:      make(map[string]any),
		color:     oldLogger.color,
		skip:      oldLogger.skip,
		level:     oldLogger.level,
		requestId: oldLogger.requestId,
	}

	maps.Copy(newLogger.tags, oldLogger.tags)

	return newLogger
}

func print(level, text, requestId string, tags map[string]any, skip int) {
	_, file, line, _ := runtime.Caller(skip)

	file = strings.TrimPrefix(file, "/app/")

	tagText := ""
	if len(tags) > 0 {
		for k, v := range tags {
			tagText += fmt.Sprintf("%s=%v ", k, v)
		}

		tagText = fmt.Sprintf("{ %s}", tagText)
	}

	output := fmt.Sprintf("%s %s [%s] %s %s < %s:%v", time.Now().Format("2006-01-02T15:04:05.0000000Z07:00"), level, requestId, text, tagText, file, line)
	fmt.Println(output)

	if TheLogger.testCapture != nil {
		*TheLogger.testCapture = fmt.Sprintf("%s%s\n", *TheLogger.testCapture, output)
	}
}

func Debug(format string, a ...any) {
	TheLogger.withSkipCount(3).Debug(format, a...)
}

func Info(format string, a ...any) {
	TheLogger.withSkipCount(3).Info(format, a...)
}
func Warn(format string, a ...any) {
	TheLogger.withSkipCount(3).Warn(format, a...)
}

func Error(format string, a ...any) {
	TheLogger.withSkipCount(3).Error(format, a...)
}

func Alert(format string, a ...any) {
	TheLogger.withSkipCount(3).Alert(format, a...)
}

func WithRequestId(reqid string) Logger {
	newLogger := newLogger(TheLogger)
	newLogger.requestId = reqid

	return newLogger
}

func (l Logger) WithRequestId(reqid string) Logger {
	newLogger := newLogger(l)
	newLogger.requestId = reqid

	return newLogger
}

func WithErr(err error) Logger {
	newLogger := newLogger(TheLogger)
	newLogger.tags[red("error")] = fmt.Sprintf("%v", err)

	return newLogger
}

func (l Logger) WithErr(err error) Logger {
	newLogger := newLogger(l)
	newLogger.tags[red("error")] = fmt.Sprintf("%v", err)

	return newLogger
}

func (l Logger) withSkipCount(skip int) Logger {
	newLogger := newLogger(l)
	newLogger.skip = skip

	return newLogger
}
