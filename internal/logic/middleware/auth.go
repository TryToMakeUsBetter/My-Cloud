package middleware

import (
	"my-cloud/utility"
	"net/http"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(g *ghttp.Request) {
	var (
		jwtKey      = utility.JwtKey
		tokenString = g.Header.Get("Authorization")
	)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		g.Response.WriteStatus(http.StatusForbidden)
		g.Exit()
	}

	g.Middleware.Next()
}
