package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-zoox/zoox"
)

func (s *server) Get() zoox.HandlerFunc {
	return func(ctx *zoox.Context) {
		id := ctx.Param().Get("id").String()
		if id == "" {
			ctx.Fail(errors.New("id is required"), http.StatusBadRequest, "id is required")
			return
		}

		job := jobStore.Get(id)
		if job == nil {
			ctx.Fail(errors.New("job not found"), http.StatusNotFound, fmt.Sprintf("job %s not found", id))
			return
		}

		ctx.Success(job)
	}
}
