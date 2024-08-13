package server

import (
	"sort"

	"github.com/go-idp/scheduler/job"
	"github.com/go-zoox/zoox"
)

func (s *server) List() zoox.HandlerFunc {
	return func(ctx *zoox.Context) {
		jobs := []*job.Job{}
		jobStore.ForEach(func(s string, j *job.Job) (stop bool) {
			jobs = append(jobs, j)
			return
		})

		sort.Slice(jobs, func(i, j int) bool {
			return jobs[i].CreatedAt.After(jobs[j].CreatedAt)
		})

		ctx.Success(zoox.H{
			"total": len(jobs),
			"data":  jobs,
		})
	}
}
