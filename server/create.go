package server

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-idp/scheduler/job"
	"github.com/go-zoox/command"
	"github.com/go-zoox/cron"
	"github.com/go-zoox/uuid"
	"github.com/go-zoox/zoox"
)

func (s *server) Create() zoox.HandlerFunc {
	return func(ctx *zoox.Context) {
		job := &job.Job{}
		if err := ctx.BindJSON(job); err != nil {
			ctx.Fail(err, http.StatusBadRequest, "invalid json")
			return
		}
		if job.Cron == "" {
			ctx.Fail(errors.New("cron is required"), http.StatusBadRequest, "cron is required")
			return
		}
		if job.Command == "" {
			ctx.Fail(errors.New("command is required"), http.StatusBadRequest, "command is required")
			return
		}

		if job.ID == "" {
			job.ID = uuid.V4()
		}

		if jobStore.Has(job.ID) {
			ctx.Fail(fmt.Errorf("job %s already exists", job.ID), http.StatusBadRequest, fmt.Sprintf("job %s already exists", job.ID))
			return
		}

		err := s.engine.AddJob(job.ID, job.Cron, func() error {
			cmd, err := command.New(&command.Config{
				Command: job.Command,
				// Engine:  "docker",
				// Image:   "whatwewant/zmicro:v1",
			})
			if err != nil {
				return err
			}

			return cmd.Run()
		}, func(cfg *cron.AddJobConfig) {
			if job.RunRightNow {
				cfg.RunRightNow = true
			}
		})
		if err != nil {
			ctx.Fail(err, http.StatusInternalServerError, fmt.Sprintf("failed to add job %s (err: %s)", job.ID, err.Error()))
			return
		}

		job.CreatedAt = time.Now()
		jobStore.Set(job.ID, job)

		ctx.JSON(200, job)
	}
}
