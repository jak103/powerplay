package log

import (
	"fmt"
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
	tags  map[string]any
	color bool
}

func (l Logger) Debug(format string, a ...any) {
	print(blue("DEBUG"), fmt.Sprintf(format, a...))
}

func (l Logger) Info(format string, a ...any) {
	print(green("INFO "), fmt.Sprintf(format, a...))
}

func (l Logger) Warn(format string, a ...any) {
	print(yellow("WARN "), fmt.Sprintf(format, a...))
}

func (l Logger) Error(format string, a ...any) {
	print(red("ERROR"), fmt.Sprintf(format, a...))
}

func (l Logger) Alert(format string, a ...any) {
	print(red("ALERT"), fmt.Sprintf(format, a...))
}

func Init(color bool) error {
	TheLogger = Logger{
		tags:  make(map[string]any),
		color: color,
	}

	return nil
}

func print(level, text string) {
	_, file, line, _ := runtime.Caller(2)

	file = strings.Replace(file, "/app/", "", 1)

	fmt.Printf("%s %s:%v > [%s] %s\n", time.Now().Format(time.RFC3339Nano), file, line, level, text)
}

func Debug(format string, a ...any) {
	TheLogger.Debug(format, a...)
}

func Info(format string, a ...any) {
	TheLogger.Info(format, a...)
}

func Warn(format string, a ...any) {
	TheLogger.Warn(format, a...)
}

func Error(format string, a ...any) {
	TheLogger.Error(format, a...)
}

func Alert(format string, a ...any) {
	TheLogger.Alert(format, a...)
}

func WithRequestId(reqid string) Logger {
	newLogger := TheLogger

	newLogger.tags["request_id"] = reqid

	return newLogger
}

func (l Logger) WithRequestId(reqid string) Logger {
	newLogger := l

	newLogger.tags["request_id"] = reqid

	return newLogger
}

func WithErr(err error) Logger {
	newLogger := TheLogger

	newLogger.tags["error"] = fmt.Sprintf("%v", err)

	return newLogger
}

func (l Logger) WithErr(err error) Logger {
	newLogger := l

	newLogger.tags["error"] = fmt.Sprintf("%v", err)

	return newLogger
}
