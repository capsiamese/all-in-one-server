package app

import (
	"aio/config"
	"aio/ent"
	v1 "aio/internal/controller/http/v1"
	"aio/internal/usecase"
	"aio/internal/usecase/repo/bark"
	"aio/internal/usecase/repo/pushdeer"
	"aio/internal/usecase/webapi"
	"aio/pkg/httpserver"
	"aio/pkg/logger"
	"aio/pkg/postgres"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func Run(cfg *config.Config) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	wd = filepath.ToSlash(wd)
	// **** log
	l, err := logger.New(cfg.Log.Level)
	if err != nil {
		return err
	}
	// **** log
	// *** cache
	/*
		r := rdb.New(&redis.Options{
			Addr:     cfg.RDB.Addr,
			Password: cfg.RDB.Password,
		}, rdb.WithKeyPrefix("app.notify"))
		err = r.Cli.Ping(context.Background()).Err()
		if err != nil {
			l.Fatalln(err)
		}
	*/
	// *** cache

	// **** db
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable",
		cfg.PG.User, cfg.PG.Password, cfg.PG.URL, cfg.PG.Port)
	pg, err := postgres.New(dsn) //postgres.WithCache(r))
	if err != nil {
		return err
	}
	defer pg.Close()
	// **** db

	entCli, err := ent.Open("postgres", dsn)
	if err != nil {
		return err
	}

	barkApns, err := webapi.NewBarkAPNs(l)
	if err != nil {
		return err
	}
	pdApns, err := webapi.NewPushDeerAPNs(l)
	if err != nil {
		return err
	}

	// **** service
	barkUseCase := usecase.NewBark(bark.New(pg, l), barkApns, l)
	pushDeerUseCase := usecase.NewPushDeer(pushdeer.New(pg, l), pdApns)
	extensionUseCase := usecase.NewExtension(entCli, l)
	// **** service

	e := gin.New()
	v1.NewRouter(e, l, &v1.UseCase{
		B: barkUseCase,
		P: pushDeerUseCase,
		E: extensionUseCase,
	})
	ctx := context.WithValue(context.Background(), "logger", l)
	ctx = context.WithValue(ctx, "rootDir", wd)
	httpServer := httpserver.New(e,
		httpserver.WithContext(ctx),
		httpserver.WithAddr(cfg.Http.Addr),
	)

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-sig:
		l.Warningln("signal received:", s)

	case s := <-httpServer.Notify():
		l.Warningln("server notify:", s)
	}

	return httpServer.Shutdown()
}
