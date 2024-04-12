package api

import (
	"context"
	"net/http"

	"github.com/perfectgentlemande/go-openapi-generator-example/internal/logger"
	"github.com/perfectgentlemande/go-openapi-generator-example/internal/service"
	openapi "github.com/perfectgentlemande/go-openapi-generator-example/openapi"
	"github.com/rs/zerolog"
)

type Controller struct {
	log  *zerolog.Logger
	srvc *service.Service
}

func New(srvc *service.Service, log *zerolog.Logger) *Controller {
	return &Controller{
		srvc: srvc,
		log:  log,
	}
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
		return openapi.ImplResponse{
			Body: openapi.ApiError{
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Code: http.StatusInternalServerError,
		}, nil
	}

	return openapi.ImplResponse{
		Body: usersFromService(users),
		Code: http.StatusOK,
	}, nil
}

func (c *Controller) UserIdDelete(ctx context.Context, id string) (openapi.ImplResponse, error) {
	reqData := openapi.RequestDataFromContext(ctx)
	log := logger.NewRequestLogger(c.log, reqData.Method, reqData.Path, reqData.ID)

	err := c.srvc.DeleteUserByID(ctx, id)
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("cannot delete user by id")
		return openapi.ImplResponse{
			Body: openapi.ApiError{
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Code: http.StatusInternalServerError,
		}, nil
	}
	return openapi.ImplResponse{
		Body: openapi.CreatedItem{
			Id: id,
		},
		Code: http.StatusOK,
	}, nil
}

func (c *Controller) UserIdGet(ctx context.Context, id string) (openapi.ImplResponse, error) {
	reqData := openapi.RequestDataFromContext(ctx)
	log := logger.NewRequestLogger(c.log, reqData.Method, reqData.Path, reqData.ID)

	user, err := c.srvc.GetUserByID(ctx, id)
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("cannot get user by id")
		return openapi.ImplResponse{
			Body: openapi.ApiError{
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Code: http.StatusInternalServerError,
		}, nil
	}

	return openapi.ImplResponse{
		Body: userFromService(&user),
		Code: http.StatusOK,
	}, nil
}

func (c *Controller) UserIdPut(ctx context.Context, id string, user openapi.User) (openapi.ImplResponse, error) {
	reqData := openapi.RequestDataFromContext(ctx)
	log := logger.NewRequestLogger(c.log, reqData.Method, reqData.Path, reqData.ID)

	updatedUser, err := c.srvc.UpdateUserByID(ctx, id, userToService(&user))
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("cannot update user by id")
		return openapi.ImplResponse{
			Body: openapi.ApiError{
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Code: http.StatusInternalServerError,
		}, nil
	}

	return openapi.ImplResponse{
		Body: userFromService(&updatedUser),
		Code: http.StatusOK,
	}, nil
}

func (c *Controller) UserPost(ctx context.Context, user openapi.User) (openapi.ImplResponse, error) {
	reqData := openapi.RequestDataFromContext(ctx)
	log := logger.NewRequestLogger(c.log, reqData.Method, reqData.Path, reqData.ID)

	id, err := c.srvc.CreateUser(ctx, userToService(&user))
	if err != nil {
		log.Error().Err(err).Msg("cannot create user")
		return openapi.ImplResponse{
			Body: openapi.ApiError{
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Code: http.StatusInternalServerError,
		}, nil
	}

	return openapi.ImplResponse{
		Body: openapi.CreatedItem{
			Id: id,
		},
		Code: http.StatusOK,
	}, nil
}
