package fmtter

import (
	"fmt"
	"path/filepath"
	"runtime"
)

const (
	depth = 3
)

type Logger struct {
	prefix string
	flags  int
}

var logger Logger

func (l *Logger) SetPrefix(prefix string) {
	l.prefix = prefix
}
func (l *Logger) SetFlags(flags int) {
	l.flags = flags
}

func Formatter(format string, level string) string {
	format = fmt.Sprintf("%s: %s", level, format)
	return formatter(format, depth)
}
func formatter(str string, dep int) string {
	pc, file, line, _ := runtime.Caller(dep)
	file = filepath.Base(file)
	fc := runtime.FuncForPC(pc).Name()

	f := fmt.Sprintf("%s:%d %s() ", file, line, fc)
	return f + str
}

func init() {
	prefix := formatter("", depth)
	logger.SetPrefix(prefix)

	logger.SetFlags(0)
}
