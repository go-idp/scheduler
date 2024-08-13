package server

import (
	"github.com/go-idp/scheduler"
	"github.com/go-zoox/datetime"
	"github.com/go-zoox/zoox"
)

func (s *server) Info() zoox.HandlerFunc {
	runningAt := datetime.Now().Format("YYYY-MM-DD HH:mm:ss")

	return func(ctx *zoox.Context) {
		ctx.JSON(200, zoox.H{
			"name":    "scheduler service for idp",
			"version": scheduler.Version,
			"status": map[string]any{
				"total":      s.engine.Length(),
				"running_at": runningAt,
			},
		})
	}
}
