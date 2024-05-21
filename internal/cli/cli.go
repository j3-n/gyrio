package cli

import (
	"fmt"
	"io"

	"github.com/urfave/cli/v2"
)

func init() {
	cli.HelpPrinter = func(writer io.Writer, templ string, data interface{}) {
		fmt.Fprintf(writer, "help\n")
	}
}

func New() cli.App {
	return cli.App{
		Name:  "gyrio",
		Usage: "view your databases in the terminal",
		Action: func(*cli.Context) error {
			fmt.Println("run gyrio help for more info")
			return nil
		},
	}
}
