package client

import (
	"fmt"

	"github.com/go-idp/scheduler/job"
	"github.com/go-zoox/fetch"
)

func (c *client) Create(job *job.Job) error {
	response, err := fetch.Post(fmt.Sprintf("%s/jobs", c.cfg.Server), &fetch.Config{
		Headers: fetch.Headers{
			"Content-Type": "application/json",
		},
		Body: job,
	})
	if err != nil {
		return err
	}

	if ok := response.Ok(); !ok {
		return fmt.Errorf("failed to create job: %s", response.String())
	}

	if code := response.Get("code").Uint(); code != 200 {
		return fmt.Errorf("failed to create job: %s", response.String())
	}

	return nil
}
