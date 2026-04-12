package main

import (
	"log/slog"
	"os"

	"backedn_go/internal/app"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	if err := app.Run(); err != nil {
		slog.Error("app failed", "error", err)
		os.Exit(1)
	}
}
