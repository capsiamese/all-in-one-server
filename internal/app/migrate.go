package app

import (
	"context"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/urfave/cli/v2"
	"log"
	"notification/ent"
)

var MigrateCmd = &cli.Command{
	Name:    "migrate",
	Aliases: []string{"mig"},
	Usage:   "migrate data base",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "dsn",
			Aliases:  []string{"d"},
			Required: true,
			Usage:    "data source name",
			//Value:    "postgres://postgres:123456@localhost:5432/postgres?sslmode=disable",
		},
		&cli.StringFlag{
			Name:  "opt",
			Value: "up",
			Usage: "migrate option",
		},
		&cli.BoolFlag{
			Name:    "ent",
			Value:   false,
			Usage:   "only migrate ent",
			Aliases: []string{"e"},
		},
	},
	Before: func(ctx *cli.Context) error {
		return nil
	},
	Action: func(ctx *cli.Context) error {
		if ctx.Bool("ent") {
			return MigrateEnt(ctx.String("dsn"))
		}
		err := Migrate(ctx.String("dsn"), ctx.String("opt"))
		if err != nil {
			return err
		}
		return MigrateEnt(ctx.String("dsn"))
	},
	After: func(ctx *cli.Context) error {
		return nil
	},
}

func Migrate(dsn, opt string) error {
	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		return err
	}
	if opt == "up" {
		err = m.Up()
	} else if opt == "down" {
		err = m.Down()
	} else {
		return errors.New("invalid option")
	}
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	se, de := m.Close()
	if se != nil {
		return se
	}
	if de != nil {
		return de
	}
	return nil
}

func MigrateEnt(dsn string) error {
	c, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	return c.Schema.Create(context.Background())
}
