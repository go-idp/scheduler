package client

import (
	"encoding/json"
	"fmt"

	"github.com/go-idp/scheduler/job"
	"github.com/go-zoox/fetch"
)

func (c *client) List() ([]*job.Job, error) {
	response, err := fetch.Get(fmt.Sprintf("%s/jobs", c.cfg.Server), &fetch.Config{})
	if err != nil {
		return nil, err
	}

	if ok := response.Ok(); !ok {
		return nil, fmt.Errorf("failed to list jobs: %s", response.String())
	}

	if code := response.Get("code").Uint(); code != 200 {
		return nil, fmt.Errorf("failed to list jobs: %s", response.String())
	}

	jobs := []*job.Job{}
	if err := json.Unmarshal([]byte(response.Get("result.data").String()), &jobs); err != nil {
		return nil, err
	}

	return jobs, nil
}
