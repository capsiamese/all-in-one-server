package app

import (
	"context"
	"errors"
	"fmt"
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
	Before: func(ctx *cli.Context) error {
		fmt.Println("before")
		return nil
	},
	Action: func(ctx *cli.Context) error {
		fmt.Println("action")
		return nil
	},
	After: func(ctx *cli.Context) error {
		fmt.Println("after")
		return nil
	},
}

func Migrate(dsn, opt string) {
	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	if opt == "up" {
		err = m.Up()
	} else if opt == "down" {
		err = m.Down()
	} else {
		err = errors.New("invalid option")
	}
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalln(err)
	}
	se, de := m.Close()
	if se != nil {
		log.Fatalln(se)
	}
	if de != nil {
		log.Fatalln(de)
	}
}

func MigrateEnt(dsn string) {
	cli, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	err = cli.Schema.Create(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
}
