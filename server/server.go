package server

import (
	"github.com/go-idp/scheduler"
	"github.com/go-idp/scheduler/job"
	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/core-utils/safe"
	"github.com/go-zoox/logger"
	"github.com/go-zoox/zoox"
	"github.com/go-zoox/zoox/defaults"
	"github.com/go-zoox/zoox/middleware"

	"github.com/go-zoox/cron"
)

var jobStore = safe.NewMap[string, *job.Job]()

type Server interface {
	Run(cfg *Config) error
	//
	Info() zoox.HandlerFunc
}

type server struct {
	engine *cron.Cron
}

func New() Server {
	c, err := cron.New(&cron.Config{
		TimeZone: "Asia/Shanghai",
	})
	if err != nil {
		panic(err)
	}

	if err := c.Start(); err != nil {
		panic(err)
	}

	return &server{
		engine: c,
	}
}

func (s *server) Run(cfg *Config) error {
	logger.Infof("server config: %+v", cfg)

	app := defaults.Default()

	app.SetBanner(fmt.Sprintf(`
  _____       _______  ___    ____    __          __     __       
 / ___/__    /  _/ _ \/ _ \  / __/___/ /  ___ ___/ /_ __/ /__ ____
/ (_ / _ \  _/ // // / ___/ _\ \/ __/ _ \/ -_) _  / // / / -_) __/
\___/\___/ /___/____/_/    /___/\__/_//_/\__/\_,_/\_,_/_/\__/_/   
                                                                 %s

Scheduler Service for Go IDP

____________________________________O/_______
                                    O\
`, scheduler.Version))

	if cfg.Username != "" || cfg.Password != "" {
		app.Use(middleware.BasicAuth("scheduler service for idp", map[string]string{
			cfg.Username: cfg.Password,
		}))
	}

	app.Post("/jobs", s.Create())
	//
	app.Post("/jobs/clear", s.Clear())
	//
	app.Get("/jobs", s.List())
	app.Get("/jobs/:id", s.Get())
	app.Delete("/jobs/:id", s.Delete())

	app.Get("/", s.Info())

	return app.Run(fmt.Sprintf(":%d", cfg.Port))
}
