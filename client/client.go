package client

import "github.com/go-idp/scheduler/job"

type Client interface {
	Create(job *job.Job) error
	Delete(id string) error
	//
	List() ([]*job.Job, error)
	Get(id string) (*job.Job, error)
	//
	Clear() error
}

type client struct {
	cfg *Config
}

func New(cfg *Config) Client {
	return &client{
		cfg: cfg,
	}
}
