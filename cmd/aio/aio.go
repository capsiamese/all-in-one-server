package main

import (
	"aio/config"
	"aio/internal/app"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

func main() {
	xxx := &cli.App{
		Name:     "all in one BOOM",
		Version:  "2021.5.9",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{Name: "siamese", Email: "cat@siamese.galaxy"},
		},
		Copyright: "All Right Received",
		Usage:     "I don't know how to use this, i'm just a cat",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Value: "config/config.yml",
				Usage: "app config file",
			},
		},
		Commands: []*cli.Command{app.MigrateCmd},
		Before:   Before,
		Action:   Action,
		After:    After,
	}
	if err := xxx.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}

func Before(ctx *cli.Context) error {
	return nil
}

func After(ctx *cli.Context) error {
	return nil
}

func Action(ctx *cli.Context) error {
	cfg, err := config.NewConfig(ctx.String("config"))
	if err != nil {
		return err
	}
	return app.Run(cfg)
}
