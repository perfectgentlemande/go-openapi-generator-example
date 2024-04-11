package api

import (
	"github.com/perfectgentlemande/go-openapi-generator-example/internal/service"
	"github.com/perfectgentlemande/go-openapi-generator-example/openapi"
)

func userFromService(u *service.User) openapi.User {
	if u == nil {
		return openapi.User{}
	}

	return openapi.User{
		Id:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func usersFromService(us []service.User) []openapi.User {
	res := make([]openapi.User, 0, len(us))

	for i := range us {
		res = append(res, userFromService(&us[i]))
	}

	return res
}
