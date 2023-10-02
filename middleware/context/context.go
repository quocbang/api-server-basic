package context

import (
	"github.com/labstack/echo"
	"github.com/quocbang/api-server-basic/middleware/authorization"
	"go.uber.org/zap"
)

const (
	Logger string = "logger"
)

func GetPrincipal(ctx echo.Context) *authorization.Principal {
	pricipal := ctx.Get(authorization.UserPrincipalKey)
	if values, ok := pricipal.(*authorization.Principal); ok {
		return values
	}
	return &authorization.Principal{}
}

// SetLogger is set logger to conext.
func SetLogger(ctx echo.Context, logger *zap.Logger) {
	ctx.Set(Logger, logger)
}

// GetLogger is get logger from conext.
func GetLogger(ctx echo.Context) *zap.Logger {
	logger := ctx.Get(Logger)
	if log, ok := logger.(*zap.Logger); ok {
		return log
	}
	return zap.L()
}
