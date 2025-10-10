package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shama3541/simplebank/token"
)

const (
	authHeaderKey  = "Authorization"
	authType       = "bearer"
	authpayloadkey = "authpayloadkey"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader(authHeaderKey)
		if len(authorization) == 0 {
			err := errors.New("the header is missing authorization")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
		fields := strings.Fields(authorization)
		if len(fields) < 2 {
			err := errors.New("the header is missing authorization")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}

		authtype := strings.ToLower(fields[0])
		if authtype != authType {
			err := errors.New("this is an invalid auth type")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
		token := fields[1]
		result, err := tokenMaker.VerifyToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.Set(authpayloadkey, result)
		ctx.Next()

	}
}
