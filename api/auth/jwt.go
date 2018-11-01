package auth

import (
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt"
)

var AuthMiddleware = &jwt.GinJWTMiddleware{
	Realm:         os.Getenv("JWT_REALM"),
	Key:           []byte(os.Getenv("JWT_SECRET")),
	Timeout:       time.Hour,
	MaxRefresh:    time.Hour * 24,
	Authenticator: Authenticator,
	Authorizator:  Authorizator,
	PayloadFunc:   PayloadFunc,
	Unauthorized:  Unauthorized,
	TokenLookup:   "header:Authorization",
	TokenHeadName: "Bearer",
	TimeFunc:      time.Now,
	SendCookie:    true,
	SecureCookie:  false,
	//CookieHTTPOnly: true,
	//CookieDomain: "localhost:5555",
	//CookieName:   "token",
}