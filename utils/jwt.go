package utils

import (
	"errors"
	"net/http"
	"time"

	"columbus/models"
	. "columbus/log"
	_ "columbus/config"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)


const (
	ExpireTime    time.Duration = time.Minute * 60 * 24 * 30
	Realm         string        = "jwt auth"
)

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 一些常量
var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
)

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	ID    string `json:"userId"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	jwt.StandardClaims
}

// 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// 获取signKey
func GetSignKey() string {
	sign_key := viper.GetString("jwt.sign_key")
	return sign_key
}

// 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		Error.Println(err)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}


// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) == 0 {
			NewError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
			c.Abort()
		}
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			NewError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
			c.Abort()
		} else {
			j := NewJWT()
			// parseToken 解析token包含的信息
			claims, err := j.ParseToken(token)
			// 塞入用户信息
			c.Set("user", claims)
			if err != nil {
				Error.Println(err)
				NewError(c, http.StatusUnauthorized, errors.New("Authorization is error"))
				c.Abort()
			}
		}
		c.Next()
	}
}


// 生成JWT令牌
func GenerateToken(user models.Account) string {

	expire := time.Now().Add(ExpireTime)

	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	claims := make(jwt.MapClaims)
	claims["id"] = user.Uid
	claims["exp"] = expire.Unix()
	token.Claims = claims

	// Sign and get the complete encoded token as a string
	sign_key := GetSignKey()
	tokenString, err := token.SignedString([]byte(sign_key))


	if err != nil {
		Error.Println(err)
		return ""
	}

	Info.Println(tokenString)
	return tokenString
}