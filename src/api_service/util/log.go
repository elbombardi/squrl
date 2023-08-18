package util

import (
	"fmt"
	"log/slog"
)

func Debugf(message string, args ...any) {
	slog.Debug(fmt.Sprintf(message, args...))
}

func Debug(args ...any) {
	slog.Debug(fmt.Sprint(args...))
}

func Infof(message string, args ...any) {
	slog.Info(fmt.Sprintf(message, args...))
}

func Info(args ...any) {
	slog.Info(fmt.Sprint(args...))
}

func Errorf(message string, args ...any) {
	slog.Error(fmt.Sprintf(message, args...))
}

func Error(args ...any) {
	slog.Error(fmt.Sprint(args...))
}
