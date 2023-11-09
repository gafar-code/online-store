package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationKey        = "authorization"
	authorizationBearer     = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(c *gin.Context) {
	useValidation := false
	mustBeValidated := []string{
		"/api/v1/product",
		"/api/v1/product/:id",
		"/api/v1/cart",
		"/api/v1/cart/delete-all",
		"/api/v1/order",
		"/api/v1/order/:id",
		"/api/v1/transaction",
	}

	for _, item := range mustBeValidated {
		if item == c.Request.URL.Path {
			useValidation = true
		}
	}

	if useValidation {
		authorizationHeader := c.GetHeader(authorizationKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			c.AbortWithStatusJSON(http.StatusUnauthorized, ResponseErr{
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
			})
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			c.AbortWithStatusJSON(http.StatusUnauthorized, ResponseErr{
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
			})
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			c.AbortWithStatusJSON(http.StatusUnauthorized, ResponseErr{
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
			})
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ResponseErr{
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
			})
			return
		}

		c.Set(authorizationPayloadKey, payload)
	}

	c.Next()
}
