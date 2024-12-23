package users

import (
	"context"

	v1 "my-cloud/api/users/v1"
	"my-cloud/internal/logic/users"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	err = users.Register(ctx, req.Username, req.Password, req.Email)
	return nil, err
}
