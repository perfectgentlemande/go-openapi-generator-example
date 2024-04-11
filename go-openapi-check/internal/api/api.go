package api

import (
	"context"

	"github.com/google/uuid"
	openapi "github.com/perfectgentlemande/go-openapi-generator-example/openapi"
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
