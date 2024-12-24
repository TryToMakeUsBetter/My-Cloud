package users

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"my-cloud/internal/dao"
	"my-cloud/internal/model/entity"
	"my-cloud/utility"
)

type UserClaims struct {
	Id       uint
	Username string
	jwt.RegisteredClaims
}

func Login(ctx context.Context, username, password string) (tokenString string, err error) {
	var user entity.Users
	err = dao.Users.Ctx(ctx).Where("username", username).Scan(&user)
	if err != nil {
		return "", errors.New("用户名或密码错误")
	}

	if user.Id == 0 {
		return "", errors.New("用户不存在")
	}

	// 将密码加密后与数据库中的密码进行比对
	if user.Password != encryptPassword(password) {
		return "", errors.New("用户名或密码错误")
	}

	// 生成token
	userClaims := &UserClaims{
		Id:       user.Id,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(6 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	return token.SignedString(utility.JwtKey)
}
