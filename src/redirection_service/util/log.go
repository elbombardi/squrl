package util

import (
	"log/slog"
	"os"
)

func NewLogger(config *Config) *slog.Logger {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: LogLevel(config.LogLevel),
	})
	logger := slog.New(logHandler)
	logger = logger.With(
		slog.Group("program_info",
			slog.Int("pid", os.Getpid()),
			slog.String("component", "squrl.redirection"),
			slog.String("version", VERSION),
			slog.String("environment", config.Environment),
		),
	)
	slog.SetDefault(logger)
	LogConfig(config)
	return logger
}
