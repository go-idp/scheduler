package commands

import (
	"github.com/go-idp/scheduler/server"
	"github.com/go-zoox/cli"
)

func RegisterServer(app *cli.MultipleProgram) {
	app.Register("server", &cli.Command{
		Name:  "server",
		Usage: "Run the scheduler server",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Usage:   "server port",
				Aliases: []string{"p"},
				EnvVars: []string{"PORT"},
				Value:   8080,
			},
		},
		Action: func(ctx *cli.Context) error {
			s := server.New()
			return s.Run(&server.Config{
				Port: ctx.Int("port"),
			})
		},
	})
}
