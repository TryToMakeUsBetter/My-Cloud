package users

import (
	"context"

	v1 "my-cloud/api/users/v1"
	"my-cloud/internal/logic/users"
	"my-cloud/internal/model"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	err = users.CheckUser(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	err = users.Register(ctx, model.UserInput{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})

	return nil, err
}
