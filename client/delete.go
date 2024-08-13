package client

import (
	"fmt"

	"github.com/go-zoox/fetch"
)

func (c *client) Delete(id string) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}

	response, err := fetch.Delete(fmt.Sprintf("%s/jobs/%s", c.cfg.Server, id), &fetch.Config{})
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
