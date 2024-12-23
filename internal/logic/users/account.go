package users

import (
	"context"
	"errors"
	"time"

	"github.com/gogf/gf/v2/frame/g"
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

func Info(ctx context.Context) (user *entity.Users, err error) {
	user = new(entity.Users)
	tokenString := g.RequestFromCtx(ctx).Request.Header.Get("Authorization")

	tokenClaims, _ := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return utility.JwtKey, nil
	})

	if claims, ok := tokenClaims.Claims.(*UserClaims); ok && tokenClaims.Valid {
		err = dao.Users.Ctx(ctx).Where("id", claims.Id).Scan(&user)
	}
	return
}

func GetUid(ctx context.Context) (uint, error) {
	user, err := Info(ctx)
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}
