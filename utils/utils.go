package utils

import (
	"errors"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	jwt "gopkg.in/dgrijalva/jwt-go.v2"
)

// JWTSecret is the encryption string
var JWTSecret string

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-=()*&ˆ%$#@!"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// RandomPassword is a function for returning a random password
func RandomPassword(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// InitiateTokenParams initiate the token parameters
func InitiateTokenParams() {
	JWTSecret = viper.GetString("jwt.secret")
}

// CheckJWTToken is ...
func CheckJWTToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := jwt.ParseFromRequest(c.Request, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(JWTSecret))
			return b, nil
		})
		if err == nil && token.Valid {
			c.Set("username", token.Claims["username"])
			c.Next()
			return
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				c.AbortWithError(400, errors.New("Invalid token"))
				return
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				c.AbortWithError(400, errors.New("Expired token"))
				return
			} else {
				c.AbortWithError(400, errors.New("Couldn't handle this token"))
				return
			}
		} else {
			c.AbortWithError(401, err)
			return
		}
	}
}
