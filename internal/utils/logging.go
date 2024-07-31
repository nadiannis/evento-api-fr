package utils

import (
	"github.com/gin-gonic/gin"
)

const LogMessageCtxKey string = "logMessage"

func SetLogMessage(ctx *gin.Context, message any) {
	ctx.Set(LogMessageCtxKey, message)
}

func GetLogMessage(ctx *gin.Context) any {
	return ctx.Value(LogMessageCtxKey)
}
