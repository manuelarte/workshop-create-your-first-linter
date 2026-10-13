// Package main with an example on how to use embeddedswagger
package main

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

const defaultAddress = ":8080"

func main() {
	logger := slog.Default()

	if err := run(logger); err != nil {
		logger.Error("Application error", slog.Any("err", err))
	}
}

func run(logger *slog.Logger) error {
	ctx := context.Background()

	handler := http.NewServeMux()
	s := &http.Server{
		Addr:           defaultAddress,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	logger.InfoContext(
		ctx,
		"server listening",
		slog.String("addr", s.Addr),
	)

	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("server error: %w", err)
	}

	return nil
}
