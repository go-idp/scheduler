package commands

import (
	"github.com/go-idp/scheduler/client"
	"github.com/go-idp/scheduler/job"
	"github.com/go-zoox/cli"
	"github.com/go-zoox/core-utils/fmt"
)

func RegisterClient(app *cli.MultipleProgram) {
	app.Register("client", &cli.Command{
		Name:  "client",
		Usage: "Run the scheduler client",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "server",
				Usage:    "scheduler server",
				Aliases:  []string{"s"},
				EnvVars:  []string{"SERVER"},
				Required: true,
			},
		},
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List all jobs",
				Action: func(ctx *cli.Context) error {
					s := client.New(&client.Config{
						Server: ctx.String("server"),
					})

					data, err := s.List()
					if err != nil {
						return err
					}

					fmt.PrintJSON(data)
					return nil
				},
			},
			{
				Name:  "create",
				Usage: "Create a new job",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "id",
						Usage:   "job id",
						Aliases: []string{"i"},
						EnvVars: []string{"ID"},
					},
					&cli.StringFlag{
						Name:     "cron",
						Usage:    "job cron, e.g. 0 10 * * * means every day at 10:00",
						Aliases:  []string{"c"},
						EnvVars:  []string{"CRON"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "command",
						Usage:    "job command",
						Aliases:  []string{"m"},
						EnvVars:  []string{"COMMAND"},
						Required: true,
					},
					&cli.BoolFlag{
						Name:    "run-right-now",
						Usage:   "run the job right now",
						Aliases: []string{"r"},
						EnvVars: []string{"RUN_RIGHT_NOW"},
					},
				},
				Action: func(ctx *cli.Context) error {
					s := client.New(&client.Config{
						Server: ctx.String("server"),
					})

					err := s.Create(&job.Job{
						ID:          ctx.String("id"),
						Cron:        ctx.String("cron"),
						Command:     ctx.String("command"),
						RunRightNow: ctx.Bool("run-right-now"),
					})
					if err != nil {
						return err
					}

					fmt.Println("job created")
					return nil
				},
			},
			{
				Name:  "clear",
				Usage: "Clear all jobs",
				Action: func(ctx *cli.Context) error {
					s := client.New(&client.Config{
						Server: ctx.String("server"),
					})

					err := s.Clear()
					if err != nil {
						return err
					}

					fmt.Println("all jobs cleared")
					return nil
				},
			},
			{
				Name:      "get",
				Usage:     "Get a job",
				ArgsUsage: "<id>",
				Action: func(ctx *cli.Context) error {
					s := client.New(&client.Config{
						Server: ctx.String("server"),
					})

					data, err := s.Get(ctx.Args().First())
					if err != nil {
						return err
					}

					fmt.PrintJSON(data)
					return nil
				},
			},
			{
				Name:      "delete",
				Usage:     "Delete a job",
				ArgsUsage: "<id>",
				Action: func(ctx *cli.Context) error {
					s := client.New(&client.Config{
						Server: ctx.String("server"),
					})

					err := s.Delete(ctx.Args().First())
					if err != nil {
						return err
					}

					fmt.Println("job deleted")
					return nil
				},
			},
		},
	})
}
