package cli

import (
	"fmt"

	"github.com/j3-n/gyrio/internal/db"
	"github.com/urfave/cli/v2"
)

func sqliteAction() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		file := ctx.String("file")

		conn := db.New(db.SQLite)
		db, err := conn.Conn(file)
		if err != nil {
			return err
		}

		tbls, err := db.Tables()
		if err != nil {
			return err
		}

		if len(tbls) == 0 {
			return nil
		}

		var data [][]map[string]any
		for _, tbl := range tbls {
			l, err := db.List(tbl)
			if err != nil {
				return err
			}

			data = append(data, l)
		}

		fmt.Println(data)

		return nil
	}
}
