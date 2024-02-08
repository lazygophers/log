package log

import "fmt"

func Trace(args ...interface{}) {
	std.Trace(args...)
}

func Log(level Level, args ...interface{}) {
	std.Log(level, args...)
}

func Info(args ...interface{}) {
	std.Info(args...)
}

func Warn(args ...interface{}) {
	std.Warn(args...)
}

func Error(args ...interface{}) {
	std.Error(args...)
}

func Panic(args ...interface{}) {
	std.Panic(args...)
}

func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

func Logf(level Level, format string, args ...interface{}) {
	std.Logf(level, format, args...)
}

func Infof(format string, args ...interface{}) {
	std.Info(fmt.Sprintf(format, args...))
}

func Warnf(format string, args ...interface{}) {
	std.Warn(fmt.Sprintf(format, args...))
}

func Errorf(format string, args ...interface{}) {
	std.Error(fmt.Sprintf(format, args...))
}

func Panicf(format string, args ...interface{}) {
	std.Panic(fmt.Sprintf(format, args...))
}

func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}

func StartMsg() {
	std.StartMsg()
}
