package client

import (
	"fmt"

	"github.com/go-zoox/fetch"
)

func (c *client) Clear() error {
	response, err := fetch.Post(fmt.Sprintf("%s/jobs/clear", c.cfg.Server), &fetch.Config{})
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
