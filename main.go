package main

import (
	"fmt"
	"github.com/Rudolph-Miller/minified-diff/minified_diff"
	"github.com/codegangsta/cli"
	"os"
)

func notExists(name string) bool {
	_, err := os.Stat(name)
	return os.IsNotExist(err)
}

func main() {
	app := cli.NewApp()
	app.Name = "minified-diff"
	app.Usage = "minified-diff file1.min.js file2.min.js"
	app.Version = "0.0.1"

	app.Action = func(c *cli.Context) {
		args := c.Args()
		if len(args) < 2 {
			fmt.Println("Please specify two files.")
			os.Exit(1)
		}

		file1 := args[0]
		file2 := args[1]

		if notExists(file1) {
			fmt.Println("First file does not exist")
			os.Exit(1)
		}

		if notExists(file2) {
			fmt.Println("Second file does not exist")
			os.Exit(1)
		}

		result := minified_diff.MinifiedDiff(file1, file2)
		fmt.Print(*result)

	}

	app.Run(os.Args)
}
