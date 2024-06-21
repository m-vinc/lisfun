package main

import (
	"lisfun/internal/app"

	"github.com/coder/serpent"
	"github.com/pkg/errors"
)

func (rootCmd *RootCmd) run() *serpent.Command {
	var (
		port     string
		loglevel string
		env      string
	)

	cmd := &serpent.Command{
		Use:   "run",
		Short: "Run lisfun",
		Handler: func(inv *serpent.Invocation) error {
			app, err := app.New(&app.Config{
				Env:      env,
				Port:     port,
				LogLevel: loglevel,
			})
			if err != nil {
				return errors.WithStack(err)
			}

			if err := app.Run(inv.Context()); err != nil {
				return errors.WithStack(err)
			}

			return nil
		},
	}

	cmd.Options = serpent.OptionSet{{
		Flag:        "env",
		Env:         "LISFUN_ENV",
		Description: "Current environment lisfun is running: [dev production].",
		Value:       serpent.StringOf(&env),
		Default:     "dev",
	}, {
		Flag:        "port",
		Env:         "LISFUN_PORT",
		Description: "Specify a lisfun port number.",
		Value:       serpent.StringOf(&port),
	}, {
		Flag:        "log_level",
		Env:         "LISFUN_LOG_LEVEL",
		Description: "Log level of the global logger.",
		Value:       serpent.StringOf(&loglevel),
		Default:     "info",
	}}

	return cmd
}
