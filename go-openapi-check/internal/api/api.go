package api

import (
	"context"

	"github.com/google/uuid"
	"github.com/perfectgentlemande/go-openapi-generator-example/internal/service"
	openapi "github.com/perfectgentlemande/go-openapi-generator-example/openapi"
)

type Controller struct {
	srvc *service.Service
}

func (c *Controller) UserGet(ctx context.Context, limit int64, offset int64) (openapi.ImplResponse, error) {
	users, err := c.srvc.ListUsers(ctx, &service.ListUsersParams{
		Limit:  &limit,
		Offset: &offset,
	})
	if err != nil {

	}

	return openapi.ImplResponse{
		Body: usersFromService(users),
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
	user, err := c.srvc.GetUserByID(ctx, id)
	if err != nil {

	}

	return openapi.ImplResponse{
		Body: userFromService(&user),
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
