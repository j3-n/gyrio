package cli

import (
	"github.com/urfave/cli/v2"
)

func New() cli.App {
	return cli.App{
		Name:     "gyrio",
		Usage:    "view your databases in the terminal",
		Commands: cmds,
	}
}
