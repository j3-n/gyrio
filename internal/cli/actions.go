package cli

import (
	"fmt"

	"github.com/j3-n/gyrio/internal/db"
	"github.com/urfave/cli/v2"
)

// Fetches all the data from the database, every item from every table.
// Returns an errro if any occur during the fetching.
func fetch(db *db.DB) ([][]map[string]any, error) {
	tbls, err := db.Tables()
	if err != nil {
		return nil, err
	}

	if len(tbls) == 0 {
		return nil, nil
	}

	var data [][]map[string]any
	for _, tbl := range tbls {
		l, err := db.List(tbl)
		if err != nil {
			return nil, err
		}

		data = append(data, l)
	}

	return data, nil
}

func sqliteAction() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		file := ctx.String("file")

		conn := db.New(db.SQLite)
		db, err := conn.Conn(file)
		if err != nil {
			return err
		}
		defer db.Close()

		data, err := fetch(db)
		if err != nil {
			return err
		}

		fmt.Println(data)

		return nil
	}
}

func postgresAction() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		addr := ctx.String("address")
		user := ctx.String("username")
		pass := ctx.String("password")
		name := ctx.String("database")
		port := ctx.String("port")

		conn := db.New(db.Postgres)
		db, err := conn.Conn(addr, user, pass, name, port)
		if err != nil {
			return err
		}
		defer db.Close()

		data, err := fetch(db)
		if err != nil {
			return err
		}

		fmt.Println(data)

		return nil
	}
}
