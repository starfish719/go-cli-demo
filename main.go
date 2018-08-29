package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "sampleApp"
	app.Usage = "This app echo input arguments"
	app.Version = "0.0.1"

	app.Action = func(context *cli.Context) error {
		if context.Bool("command") {
			fmt.Println(context.Args().Get(0))
		} else {
			fmt.Println("nothing")
		}
		return nil
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "command, c",
			Usage: "option",
		},
	}

	app.Run(os.Args)
}
