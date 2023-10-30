package log

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/jak103/leaguemanager/internal/config"
)

var blue = color.New(color.FgBlue).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()

// var orange = color.New(color.Fgor).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()

func Init(config *config.Config) error {
	return nil
}

func print(level, text string) {
	fmt.Printf("%s [%s] %s\n", time.Now().Format(time.RFC3339Nano), level, text)
}

func Debug(format string, a ...any) {
	print(blue("DEBUG"), fmt.Sprintf(format, a...))
}

func Info(format string, a ...any) {
	print(green("INFO "), fmt.Sprintf(format, a...))
}

func Warn(format string, a ...any) {
	print(yellow("WARN "), fmt.Sprintf(format, a...))
}

func Error(format string, a ...any) {
	print(red("ERROR"), fmt.Sprintf(format, a...))
}
