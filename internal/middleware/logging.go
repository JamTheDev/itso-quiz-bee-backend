package middleware

import (
	"net/http"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

type Model struct {
	Log zerolog.Logger
}

func (m Model) RequestLogger(next http.Handler) http.Handler {
	h := hlog.NewHandler(m.Log)

	access := hlog.AccessHandler(
		func(r *http.Request, status, size int, duration time.Duration) {
			hlog.FromRequest(r).Info().
				Str("method", r.Method).
				Str("url", r.URL.RequestURI()).
				Int("status_code", status).
				Str("user_agent", r.UserAgent()).
				Dur("elapsed_ms", duration).
				Msg("Incoming request.")
		},
	)

	userAgent := hlog.UserAgentHandler("user_agent")

	return h(access(userAgent(next)))
}
