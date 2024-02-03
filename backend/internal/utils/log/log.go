package log

import (
	"fmt"
	"maps"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
)

var blue = color.New(color.FgBlue).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()

// var orange = color.New(color.Fgor).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()

var TheLogger Logger

type Logger struct {
	tags        map[string]any
	color       bool
	skip        int
	testCapture *string
}

func (l *Logger) SetTestCapture(ptr *string) {
	l.testCapture = ptr
}

func (l Logger) Debug(format string, a ...any) {
	print(blue("[DEBUG]"), fmt.Sprintf(format, a...), l.tags, l.skip)
}

func (l Logger) Info(format string, a ...any) {
	print(green("[INFO ]"), fmt.Sprintf(format, a...), l.tags, l.skip)
}

func (l Logger) Warn(format string, a ...any) {
	print(yellow("[WARN ]"), fmt.Sprintf(format, a...), l.tags, l.skip)
}

func (l Logger) Error(format string, a ...any) {
	print(red("[ERROR]"), fmt.Sprintf(format, a...), l.tags, l.skip)
}

func (l Logger) Alert(format string, a ...any) {
	print(red("[ALERT]"), fmt.Sprintf(format, a...), l.tags, l.skip)
}

func Init(color bool) error {
	TheLogger = Logger{
		tags:  make(map[string]any),
		color: color,
		skip:  2,
	}

	return nil
}

func print(level, text string, tags map[string]any, skip int) {
	_, file, line, _ := runtime.Caller(skip)

	file = strings.Replace(file, "/app/", "", 1)

	tagText := ""
	if len(tags) > 0 {
		for k, v := range tags {
			tagText += fmt.Sprintf("%s=%v ", k, v)
		}

		tagText = fmt.Sprintf("{ %s}", tagText)
	}

	output := fmt.Sprintf("%s %s %s:%v > %s %s", time.Now().Format(time.RFC3339Nano), level, file, line, text, tagText)
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
	newLogger := TheLogger
	newLogger.tags[blue("request_id")] = reqid

	return newLogger
}

func (l Logger) WithRequestId(reqid string) Logger {
	newLogger := l
	maps.Copy(newLogger.tags, l.tags)
	newLogger.tags[blue("request_id")] = reqid

	return newLogger
}

func WithErr(err error) Logger {
	newLogger := TheLogger
	newLogger.tags[red("error")] = fmt.Sprintf("%v", err)

	return newLogger
}

func (l Logger) WithErr(err error) Logger {
	newLogger := l
	maps.Copy(newLogger.tags, l.tags)
	fmt.Println("In WithErr:", err)
	newLogger.tags[red("error")] = fmt.Sprintf("%v", err)

	return newLogger
}

func (l Logger) withSkipCount(skip int) Logger {
	newLogger := l
	maps.Copy(newLogger.tags, l.tags)

	newLogger.skip = skip

	return newLogger
}
