package main

import (
	"context"
	"log"
	"ns/cmd"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli/v3"
)

func main() {
	rootCmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:    "lookup",
				Aliases: []string{"l"},
				Usage:   color.BlueString("custom tool to make dns checkup on domains"),
				Action:  cmd.Lookup,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "type",
						Value: "A",
						Usage: "type of dns record (A, MX, NS, etc)",
					},
				},
			},
			{
				Name:    "scan",
				Aliases: []string{"s"},
				Usage:   color.BlueString("custom tool to make port checkup on domains"),
				Action:  cmd.ScanPort,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "port",
						Value: "80",
						Usage: "port to check (default: 80)",
					},
				},
			},
		},
	}
	if err := rootCmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
