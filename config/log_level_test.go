package config

import (
	"log/slog"
	"testing"
)

func TestParseLogLevel(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  slog.Level
		ok    bool
	}{
		{name: "debug", input: "debug", want: slog.LevelDebug, ok: true},
		{name: "info", input: "info", want: slog.LevelInfo, ok: true},
		{name: "warn", input: "warn", want: slog.LevelWarn, ok: true},
		{name: "warning", input: "warning", want: slog.LevelWarn, ok: true},
		{name: "error", input: "error", want: slog.LevelError, ok: true},
		{name: "uppercase and spaced", input: "  DEBUG  ", want: slog.LevelDebug, ok: true},
		{name: "unknown", input: "trace", want: slog.LevelInfo, ok: false},
		{name: "empty", input: "", want: slog.LevelInfo, ok: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := ParseLogLevel(tt.input)
			if got != tt.want {
				t.Fatalf("ParseLogLevel(%q) level = %v, want %v", tt.input, got, tt.want)
			}
			if ok != tt.ok {
				t.Fatalf("ParseLogLevel(%q) ok = %v, want %v", tt.input, ok, tt.ok)
			}
		})
	}
}
