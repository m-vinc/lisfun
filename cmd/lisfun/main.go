package main

import (
	"log"

	"github.com/coder/serpent"
	"github.com/pkg/errors"
)

type RootCmd struct {
	serpent.Command
}

func (rootCmd *RootCmd) Run() error {
	return errors.WithStack(rootCmd.Command.
		Invoke().
		WithOS().
		Run())
}

func main() {
	cmd := &RootCmd{
		Command: serpent.Command{
			Use:   "lisfun",
			Short: "Listen spotify and have fun",
		},
	}
	cmd.AddSubcommands(
		cmd.run(),
	)

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
