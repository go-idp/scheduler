package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-zoox/zoox"
)

func (s *server) Delete() zoox.HandlerFunc {
	return func(ctx *zoox.Context) {
		id := ctx.Param().Get("id").String()
		if id == "" {
			ctx.Fail(errors.New("id is required"), http.StatusBadRequest, "id is required")
			return
		}

		err := s.engine.RemoveJob(id)
		if err != nil {
			ctx.Fail(err, http.StatusInternalServerError, fmt.Sprintf("failed to remove job %s(err: %s)", id, err.Error()))
			return
		}

		jobStore.Del(id)

		ctx.Success(nil)
	}
}
