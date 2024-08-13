package server

import (
	"fmt"
	"net/http"

	"github.com/go-zoox/zoox"
)

func (s *server) Clear() zoox.HandlerFunc {
	return func(ctx *zoox.Context) {
		if err := jobStore.Clear(); err != nil {
			ctx.Fail(err, http.StatusInternalServerError, fmt.Sprintf("failedt to clear job(1): job store (err: %s)", err.Error()))
			return
		}

		if err := s.engine.Clear(); err != nil {
			ctx.Fail(err, http.StatusInternalServerError, fmt.Sprintf("failedt to clear job(2): cron engine (err: %s)", err.Error()))
			return
		}

		ctx.Success(nil)
	}
}
