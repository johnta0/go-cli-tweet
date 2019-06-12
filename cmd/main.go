package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "Go CLI Tweet"
	app.Usage = "cli app to tweet without browsing timeline. Browsing timeline is dangerous."
	app.Version = "1.0"

	// app.Action = func (context *cli.Context) error {
	// 	return nil
	// }

	app.Action = printArg

	app.Run(os.Args)
}

func printArg(c *cli.Context) error {
	if c.NArg() == 1 {
		fmt.Println(c.Args().Get(0))
	}
	return nil
}
