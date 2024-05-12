package api

import (
	"context"
	"net/http"

	"github.com/perfectgentlemande/go-openapi-generator-example/internal/logger"
	"github.com/perfectgentlemande/go-openapi-generator-example/internal/openapi"
	"github.com/perfectgentlemande/go-openapi-generator-example/internal/service"
	"github.com/rs/zerolog"
)

type Controller struct {
	log  *zerolog.Logger
	srvc *service.Service
}

type Config struct {
	Addr string
}

func New(srvc *service.Service, log *zerolog.Logger) *Controller {
	return &Controller{
		srvc: srvc,
		log:  log,
	}
}

func respond(payload interface{}, code int) (openapi.ImplResponse, error) {
	return openapi.ImplResponse{
		Body: payload,
		Code: code,
	}, nil
}
func respondWithError(msg string, code int) (openapi.ImplResponse, error) {
	return respond(openapi.ApiError{Message: msg}, code)
}
func respondWithSuccess(payload interface{}) (openapi.ImplResponse, error) {
	return respond(payload, http.StatusOK)
}

func (c *Controller) UserGet(ctx context.Context, limit int64, offset int64) (openapi.ImplResponse, error) {
	reqData := openapi.RequestDataFromContext(ctx)
	log := logger.NewRequestLogger(c.log, reqData.Method, reqData.Path, reqData.ID)

	users, err := c.srvc.ListUsers(ctx, &service.ListUsersParams{
		Limit:  &limit,
		Offset: &offset,
	})
	if err != nil {
		log.Error().Err(err).Msg("cannot list users")
		return respondWithError(
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}

	return respondWithSuccess(users)
}

func (c *Controller) UserIdDelete(ctx context.Context, id string) (openapi.ImplResponse, error) {
	reqData := openapi.RequestDataFromContext(ctx)
	log := logger.NewRequestLogger(c.log, reqData.Method, reqData.Path, reqData.ID)

	err := c.srvc.DeleteUserByID(ctx, id)
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("cannot delete user by id")
		return respondWithError(
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}
	return respondWithSuccess(openapi.CreatedItem{Id: id})
}

func (c *Controller) UserIdGet(ctx context.Context, id string) (openapi.ImplResponse, error) {
	reqData := openapi.RequestDataFromContext(ctx)
	log := logger.NewRequestLogger(c.log, reqData.Method, reqData.Path, reqData.ID)

	user, err := c.srvc.GetUserByID(ctx, id)
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("cannot get user by id")
		return respondWithError(
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}

	return respondWithSuccess(userFromService(&user))
}

func (c *Controller) UserIdPut(ctx context.Context, id string, user openapi.User) (openapi.ImplResponse, error) {
	reqData := openapi.RequestDataFromContext(ctx)
	log := logger.NewRequestLogger(c.log, reqData.Method, reqData.Path, reqData.ID)

	updatedUser, err := c.srvc.UpdateUserByID(ctx, id, userToService(&user))
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("cannot update user by id")
		return respondWithError(
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}

	return respondWithSuccess(userFromService(&updatedUser))
}

func (c *Controller) UserPost(ctx context.Context, user openapi.User) (openapi.ImplResponse, error) {
	reqData := openapi.RequestDataFromContext(ctx)
	log := logger.NewRequestLogger(c.log, reqData.Method, reqData.Path, reqData.ID)

	id, err := c.srvc.CreateUser(ctx, userToService(&user))
	if err != nil {
		log.Error().Err(err).Msg("cannot create user")
		return respondWithError(
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}

	return respondWithSuccess(openapi.CreatedItem{Id: id})
}
