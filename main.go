package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "minified-diff"
	app.Usage = "minified-diff file1.js file2.js"
	app.Version = "0.0.1"

	app.Action = func(c *cli.Context) {
		if len(c.Args()) < 2 {
			println("Please specify files")
			os.Exit(1)
		}

		println("Hello")
	}

	app.Run(os.Args)
}
