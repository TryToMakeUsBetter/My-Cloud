package account

import (
	"context"

	v1 "my-cloud/api/account/v1"
	"my-cloud/internal/logic/users"
)

func (c *ControllerV1) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {
	user, err := users.Info(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.InfoRes{
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
	return
}
