package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/perfectgentlemande/go-openapi-generator-example/internal/api"
	dbuser "github.com/perfectgentlemande/go-openapi-generator-example/internal/database"
	"github.com/perfectgentlemande/go-openapi-generator-example/internal/openapi"
	"github.com/perfectgentlemande/go-openapi-generator-example/internal/service"
	"github.com/rs/zerolog"
)

func main() {
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	signalCtx, _ := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	conf, err := readConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot read config")
	}

	dbUser, err := dbuser.NewDatabase(signalCtx, &dbuser.Config{
		DBName:  conf.DBUser.DBName,
		ConnStr: conf.DBUser.ConnStr,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to database")
	}
	defer dbUser.Close(signalCtx)

	err = dbUser.Ping(signalCtx)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot ping db")
	}

	UserAPIService := api.New(service.NewService(dbUser), &log)

	UserAPIController := openapi.NewUserAPIController(UserAPIService)

	router := openapi.NewRouter(UserAPIController)

	srv := &http.Server{
		Addr:    conf.Server.Addr,
		Handler: router,
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)
	timeoutCtx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	go func() {
		defer wg.Done()

		log.Info().Str("addr", ":80").Msg("server starting")
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Err(err).Msg("caught server listenAndServe error")
		} else {
			log.Info().Msg("listenAndServe successfully aborted")
		}
	}()
	go func() {
		defer wg.Done()

		<-signalCtx.Done()
		log.Info().Msg("caught os signal: server will be shut down")

		err := srv.Shutdown(timeoutCtx)
		if err != nil {
			log.Err(err).Msg("caught server shutdown error")
		} else {
			log.Info().Msg("server shutdown successful")
		}
	}()

	wg.Wait()
}
