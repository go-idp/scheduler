package main

import (
	"github.com/go-idp/scheduler"
	"github.com/go-idp/scheduler/cmd/scheduler/commands"
	"github.com/go-zoox/cli"
)

func main() {
	app := cli.NewMultipleProgram(&cli.MultipleProgramConfig{
		Name:    "scheduler",
		Usage:   "Scheduler Service for IDP",
		Version: scheduler.Version,
	})

	// server
	commands.RegisterServer(app)
	// client
	commands.RegisterClient(app)

	app.Run()
}
