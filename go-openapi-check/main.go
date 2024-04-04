/*
 * go-openapi-generator-example
 *
 * Example of REST API. Includes such things as MongoDB, slog, gorilla mux etc... Some kind of a sample project API for educational purposes.
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	openapi "github.com/perfectgentlemande/go-openapi-generator-example/openapi"
	"github.com/rs/zerolog"
)

type Controller struct {
}

func (c *Controller) UserGet(ctx context.Context, limit int64, offset int64) (openapi.ImplResponse, error) {
	return openapi.ImplResponse{
		Body: []openapi.User{
			{
				Id:       uuid.NewString(),
				Username: "alice01",
			},
			{
				Id:       uuid.NewString(),
				Username: "bob01",
			},
		},
		Code: 200,
	}, nil
}
func (c *Controller) UserIdDelete(ctx context.Context, id string) (openapi.ImplResponse, error) {
	return openapi.ImplResponse{
		Body: openapi.User{
			Id:       uuid.NewString(),
			Username: "alice01",
		},
		Code: 200,
	}, nil
}
func (c *Controller) UserIdGet(ctx context.Context, id string) (openapi.ImplResponse, error) {
	return openapi.ImplResponse{
		Body: openapi.User{
			Id:       uuid.NewString(),
			Username: "alice01",
		},
		Code: 200,
	}, nil
}
func (c *Controller) UserIdPut(ctx context.Context, id string) (openapi.ImplResponse, error) {
	return openapi.ImplResponse{
		Body: openapi.User{
			Id:       uuid.NewString(),
			Username: "alice01",
		},
		Code: 200,
	}, nil
}
func (c *Controller) UserPost(ctx context.Context) (openapi.ImplResponse, error) {
	return openapi.ImplResponse{
		Body: openapi.User{
			Id:       uuid.NewString(),
			Username: "alice01",
		},
		Code: 200,
	}, nil
}

func main() {
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	UserAPIService := &Controller{}
	UserAPIController := openapi.NewUserAPIController(UserAPIService)

	router := openapi.NewRouter(UserAPIController)

	go func() {
		log.Info().Str("addr", ":8000").Msg("Server starting")
		log.Fatal().Err(http.ListenAndServe(":8080", router)).Msg("failed to listen")
	}()

	<-ctx.Done()
	log.Info().Msg("caught os signal")
}
