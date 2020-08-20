package colorlog

import (
	"os"

	"github.com/fatih/color"
	"github.com/tangx/colorlog/internal/fmtter"
)

func Println(format string, args ...interface{}) {
	format = fmtter.Formatter(format, "")
	color.White(format, args...)
}

func Trace(format string, args ...interface{}) {
	format = fmtter.Formatter(format, "trace")
	color.White(format, args...)
}

func Info(format string, args ...interface{}) {
	format = fmtter.Formatter(format, "info")
	color.Blue(format, args...)
}

func Warn(format string, args ...interface{}) {
	format = fmtter.Formatter(format, "warn")
	color.Yellow(format, args...)
}

func Error(format string, args ...interface{}) {
	format = fmtter.Formatter(format, "error")
	color.Red(format, args...)
}

func Fatal(format string, args ...interface{}) {
	format = fmtter.Formatter(format, "fatal")
	color.HiRed(format, args...)
	os.Exit(1)
}

func Panic(format string, args ...interface{}) {
	format = fmtter.Formatter(format, "panic")
	format = color.HiRedString(format, args...)
	panic(format)
}
