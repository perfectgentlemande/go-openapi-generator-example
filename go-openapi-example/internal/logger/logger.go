package logger

import (
	"github.com/rs/zerolog"
)

func NewRequestLogger(base *zerolog.Logger, method, path, id string) *zerolog.Logger {
	childLogger := base.With().Fields(
		map[string]interface{}{
			"method":     method,
			"request_id": id,
			"path":       path,
		}).Logger()

	return &childLogger
}
