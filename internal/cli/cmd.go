package cli

import (
	"errors"
	"strconv"

	"github.com/urfave/cli/v2"
)

var cmds = []*cli.Command{
	// Misc cmds
	{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "view ways to use gyrio",
	},

	{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "get the current version of gyrio",
		Action: func(ctx *cli.Context) error {
			return nil
		},
	},

	// Database viewing commands
	{
		Name:    "sqlite",
		Aliases: []string{"sl"},
		Usage:   "view an sqlite database with gyrio",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "file",
				Aliases:  []string{"f"},
				Usage:    "sqlite database file",
				Required: true,
			},
		},
		Action: sqliteAction(),
	},

	{
		Name:    "postgres",
		Aliases: []string{"pg"},
		Usage:   "view a postgres database with gyrio",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "address",
				Aliases:  []string{"a"},
				Usage:    "postgres host address, not including the port",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "username",
				Aliases:  []string{"U"},
				Usage:    "postgres username",
				Required: true,
			},
			// TODO: this can be a potential security weakness, may want to remove this in the future.
			&cli.StringFlag{
				Name:     "password",
				Aliases:  []string{"P"},
				Usage:    "postgres password",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "database",
				Aliases:  []string{"d"},
				Usage:    "postgres database name",
				Required: true,
			},
			&cli.StringFlag{
				Name:        "port",
				Aliases:     []string{"p"},
				Usage:       "postgres port",
				Value:       "5432",
				DefaultText: "5432",
				Required:    false,
				Action: func(ctx *cli.Context, val string) error {
					i, err := strconv.Atoi(val)
					if err != nil {
						return err
					}

					if i > 65535 {
						return errors.New("error, port has exceeded maximum value")
					}

					return nil
				},
			},
		},
		Action: postgresAction(),
	},

	{
		// TODO: add mysql support
		Name:    "mysql",
		Aliases: []string{"ms"},
		Usage:   "view a mysql database with gyrio",
		Action: func(ctx *cli.Context) error {
			return nil
		},
	},
}
