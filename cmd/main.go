package main

import (
	"anksus/cache-system/command"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Name:    "app",
		Version: "1.0.0",
		Usage:   "Cache Service",
		Commands: []*cli.Command{
			{
				Name:   "start",
				Usage:  "starts the cache system",
				Action: command.StartCacheSystem,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
