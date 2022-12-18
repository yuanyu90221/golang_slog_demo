package main

import (
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	// slog.Info("Test slog")
	// textHandler := slog.NewTextHandler(os.Stdout)
	// logger := slog.New(textHandler)
	// logger.Info("Go slog test")
	opts := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	jsonHandler := opts.NewJSONHandler(os.Stdout)
	jsonLogger := slog.New(jsonHandler)
	// jsonLogger.Info("Go slog test")
	// logger.Info("Usage Statistic", slog.Int("current-memory", 50))
	jsonLogger.Info("Usage Statistics",
		slog.Group("memory",
			slog.Int("current", 50),
			slog.Int("min", 20),
			slog.Int("max", 80)),
		slog.Int("cpu", 10),
		slog.String("app-version", "v0.0.1-beta"),
	)
	jsonLogger.Debug("Debug")
	jsonLogger.Warn("Warn")
}
