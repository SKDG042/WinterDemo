package middleware

import (
	"WinterDemo/api/types"
	"WinterDemo/configs"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 生成token
func GenerateToken(username string) (string, error) {
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Minute)), // 有效时间1分钟
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 生效时间
			Issuer:    "042",                          // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(configs.GlobalConfig.Server.JWTSecret))
}

// 解析token
func ParseToken(tokenString string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.GlobalConfig.Server.JWTSecret), nil
	}) //解析token

	if err != nil {
		return "", fmt.Errorf("token解析错误")	
	}

	if !token.Valid {
		return "", fmt.Errorf("token无效")
	}//检验token是否有效

	return claims.Username, nil
}

func JWTauth() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		authHeader := string(ctx.GetHeader("Authorization"))
		if authHeader == "" {
			ctx.JSON(consts.StatusUnauthorized, types.ErrorResponse(10006,"请先登录"))
			ctx.Abort()
			return
		}

		//JWT的token格式为Bearer token
		//通过检测authHeader是否以Bearer开头判断是否为Bearer token
		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(consts.StatusUnauthorized, types.ErrorResponse(10006,"无效的token格式"))
			ctx.Abort()
			return 
		}

		token := authHeader[7:] //Bearer带空格一共7个字符

		username, err := ParseToken(token)
		if err != nil {
			ctx.JSON(consts.StatusUnauthorized, types.ErrorResponse(10006, "token解析错误"))
			ctx.Abort()
			return
		}

		ctx.Set("username", username) //将username设置到*app.RequestContext中
		ctx.Next(c) //调用下一个处理函数，这个请求周期没有结束
	}
}