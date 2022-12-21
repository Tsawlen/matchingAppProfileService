package middleware

import (
	"app/matchingAppProfileService/common/security"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(context *gin.Context) {
	header := context.Request.Header.Get("Authorization")
	if header == "" {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	authorization := strings.TrimPrefix(header, "Bearer ")

	token, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		key, err := security.GetPublicToken()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return key, nil
	})

	if err != nil {
		_, ok := err.(*jwt.ValidationError)
		if ok {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			context.AbortWithStatus(http.StatusUnauthorized)
		}
	} else {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// Get the JWT
	context.Next()
}

func MicroServiceAuth(context *gin.Context) {
	header := context.Request.Header.Get("Authorization")
	if header == "" {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	authorization := strings.TrimPrefix(header, "Bearer ")

	token, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		key, err := security.GetPublicToken()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return key, nil
	})

	if err != nil {
		_, ok := err.(*jwt.ValidationError)
		if ok {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["microServiceAuth"].(bool) != true {
			context.AbortWithStatus(http.StatusUnauthorized)
		}
	} else {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// Get the JWT
	context.Next()
}
