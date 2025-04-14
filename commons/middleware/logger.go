package middleware

import (
	"strings"
	"trendtracker/constants"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL
		reqID := getRequestID(ctx)

		ctx.Set(constants.PathLogParam, path)
		ctx.Set(constants.HeaderRequestID, reqID)

		ctx.Next()
	}
}

func getRequestID(ctx *gin.Context) string {
	reqID := ctx.Request.Header.Get(constants.HeaderRequestID)
	if strings.EqualFold(reqID, constants.Empty) {
		reqID = uuid.NewString()
	}

	return reqID
}
