package users

import (
	"context"

	"my-cloud/internal/dao"
	"my-cloud/internal/model"
	"my-cloud/internal/model/do"

	"github.com/gogf/gf/v2/errors/gerror"
)

func Register(ctx context.Context, in model.UserInput) error {
	_, err := dao.Users.Ctx(ctx).Data(do.Users{
		Username: in.Username,
		Password: encryptPassword(in.Password),
		Email:    in.Email,
	}).Insert()
	if err != nil {
		return err
	}
	return nil
}

func CheckUser(ctx context.Context, username string) error {
	cnt, err := dao.Users.Ctx(ctx).Where("username =", username).Count()
	if err != nil {
		return err
	}

	if cnt > 0 {
		return gerror.New("重复的Username")
	}

	return nil
}
