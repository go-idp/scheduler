package client

import (
	"encoding/json"
	"fmt"

	"github.com/go-idp/scheduler/job"
	"github.com/go-zoox/fetch"
)

func (c *client) Get(id string) (*job.Job, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	response, err := fetch.Get(fmt.Sprintf("%s/jobs/%s", c.cfg.Server, id), &fetch.Config{})
	if err != nil {
		return nil, err
	}

	if ok := response.Ok(); !ok {
		return nil, fmt.Errorf("failed to retrieve job: %s", response.String())
	}

	if code := response.Get("code").Uint(); code != 200 {
		return nil, fmt.Errorf("failed to retrieve job: %s", response.String())
	}

	var j job.Job
	if err := json.Unmarshal([]byte(response.Get("result").String()), &j); err != nil {
		return nil, err
	}

	return &j, nil
}
