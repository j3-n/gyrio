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

// Format the data from a tables data into an []string of keys and
// a [][]string of data.
// Will convert the generic any type into a string.
func format(data []map[string]any) ([]string, [][]string) {
	keys := make([]string, len(data[0]))
	i := 0
	for key := range data[0] {
		keys[i] = key
		i++
	}

	tbl := make([][]string, len(data))
	i = 0
	for _, chunk := range data {
		j := 0
		tbl[i] = make([]string, len(chunk))

		for _, val := range chunk {
			str := fmt.Sprintf("%v", val)
			tbl[i][j] = str
			j++
		}

		i++
	}

	return keys, tbl
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
		keys, tbl := format(data[0])

		fmt.Println(keys)
		fmt.Println(tbl)

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

func mysqlAction() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		user := ctx.String("username")
		pass := ctx.String("password")
		addr := ctx.String("address")
		port := ctx.String("port")
		name := ctx.String("database")

		conn := db.New(db.MySQL)
		db, err := conn.Conn(user, pass, addr, port, name)
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

func mariaAction() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		user := ctx.String("username")
		pass := ctx.String("password")
		addr := ctx.String("address")
		port := ctx.String("port")
		name := ctx.String("database")

		conn := db.New(db.Maria)
		db, err := conn.Conn(user, pass, addr, port, name)
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
